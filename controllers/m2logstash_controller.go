/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"github.com/go-logr/logr"
	loggerv1 "github.com/m2-oss/logstash-operator/api/v1"
	"github.com/redhat-cop/operator-utils/pkg/util"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const (
	label              = "m2/logger"
	ilmLabel           = "m2/logger-ilm-policy"
	typeLabel          = "m2/logger-type"
	replicasAnnotation = "logger.m2.ru/logstash-replicas"
)

// M2LogstashReconciler reconciles a M2Logstash object
type M2LogstashReconciler struct {
	// client.Client
	Log logr.Logger
	util.ReconcilerBase
}

var controllerLog = ctrl.Log.WithName("controller").WithName("M2Logstash")

// +kubebuilder:rbac:groups=logger.m2.ru,resources=m2logstashes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=logger.m2.ru,resources=m2logstashes/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=logger.m2.ru,resources=m2logstashes/finalizers,verbs=update
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch
// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=extensions,resources=podsecuritypolicies,verbs=use;
// +kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;update;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the M2Logstash object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.0/pkg/reconcile
func (r *M2LogstashReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithName("controller").WithName("m2logstash")
	controllerLog.Info("start reconcile", "request", req.NamespacedName)

	if req.NamespacedName.Namespace == "" {
		namespace := &corev1.Namespace{}
		if err := r.GetClient().Get(ctx, req.NamespacedName, namespace); err != nil {
			if apierrors.IsNotFound(err) {
				return ctrl.Result{}, nil
			}
			controllerLog.Error(err, "Cannot get object Namespace")
			return ctrl.Result{}, err
		}

		instanceList := &loggerv1.M2LogstashList{}
		err := r.GetClient().List(ctx, instanceList)
		if err != nil {
			return ctrl.Result{}, nil
		}

		for _, instance := range instanceList.Items {
			if namespace.GetLabels()[label] == "true" {
				err, response := r.syncResources(ctx, instance, req.NamespacedName.Name)
				if err != nil {
					controllerLog.Error(err, response)
					return r.ManageError(ctx, &instance, err)
				}
			} else {
				err, response := r.deleteResources(ctx, instance, req.NamespacedName.Name)
				if err != nil {
					controllerLog.Error(err, response)
					return r.ManageError(ctx, &instance, err)
				}
			}
		}
	}

	instance := &loggerv1.M2Logstash{}
	if err := r.GetClient().Get(ctx, req.NamespacedName, instance); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		controllerLog.Error(err, "cannot get object M2Logstash")
		return ctrl.Result{}, err
	} else {
		nameSpaceList := &corev1.NamespaceList{}
		opts := []client.ListOption{
			client.MatchingLabels{"m2/logger": "true"},
		}
		if err := r.GetClient().List(ctx, nameSpaceList, opts...); err != nil {
			return r.ManageError(ctx, instance, err)
		}
		for _, item := range nameSpaceList.Items {
			err, response := r.syncResources(ctx, *instance, item.GetName())
			if err != nil {
				log.Error(err, response)
				return r.ManageError(ctx, instance, err)
			}
		}
	}

	controllerLog.Info("finish reconcile", "request", req.NamespacedName)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *M2LogstashReconciler) SetupWithManager(mgr ctrl.Manager) error {

	return ctrl.NewControllerManagedBy(mgr).
		For(&loggerv1.M2Logstash{}).
		Watches(
			&source.Kind{Type: &corev1.Namespace{}},
			&handler.EnqueueRequestForObject{}).
		WithEventFilter(namespaceFilter()).
		WithOptions(controller.Options{MaxConcurrentReconciles: 2}).
		Complete(r)
}

// SyncResources creating or updating resources related to M2Logstash
func (r *M2LogstashReconciler) syncResources(ctx context.Context, instance loggerv1.M2Logstash, namespace string) (error, string) {
	controllerLog.Info("Sync instance resources", "instance", instance.GetName())
	pipelineList := &loggerv1.M2LogstashPipelineList{}
	r.GetClient().List(context.TODO(), pipelineList)
	err := r.CreateResourceIfNotExists(
		ctx,
		&instance,
		instance.GetNamespace(),
		r.serviceAccountFromCR(&instance, namespace),
	)
	if err != nil {
		return err, "cannot create serviceAccount"
	}

	err = r.CreateResourceIfNotExists(
		ctx,
		&instance,
		instance.GetNamespace(),
		r.roleFromCR(&instance, namespace),
	)
	if err != nil {
		return err, "cannot create role"
	}

	err = r.CreateOrUpdateResource(
		ctx,
		&instance,
		instance.GetNamespace(),
		r.roleBindingFromCR(&instance, namespace),
	)
	if err != nil {
		return err, "cannot create rolebinding"
	}

	err = r.CreateOrUpdateResource(
		ctx,
		&instance,
		instance.GetNamespace(),
		r.statefulSetFromCR(&instance, namespace, pipelineList),
	)

	if err != nil {
		return err, "cannot create statefulset"
	}
	err = r.CreateResourceIfNotExists(
		ctx,
		&instance,
		instance.GetNamespace(),
		r.clusterIPServiceFromCR(&instance, namespace, pipelineList),
	)

	if err != nil {
		return err, "cannot create clusterIP service"
	}

	err = r.CreateResourceIfNotExists(ctx,
		&instance,
		instance.GetNamespace(),
		headLessServiceFromCR(&instance, namespace),
	)
	if err != nil {
		return err, "cannot create headless service"
	}

	controllerLog.Info("Resources synced", "instance", instance.GetName())
	return nil, "Success"
}

// deleteResources deleting resources related to M2Logstash
func (r *M2LogstashReconciler) deleteResources(ctx context.Context, instance loggerv1.M2Logstash, namespace string) (error, string) {
	controllerLog.Info("Delete instance resources", "instance", instance.GetName())
	err := r.DeleteResourceIfExists(
		ctx,
		r.serviceAccountFromCR(&instance, namespace),
	)
	if err != nil {
		return err, "cannot delete serviceAccount"
	}

	err = r.DeleteResourceIfExists(
		ctx,
		r.roleFromCR(&instance, namespace),
	)
	if err != nil {
		return err, "cannot delete role"
	}

	err = r.DeleteResourceIfExists(
		ctx,
		r.roleBindingFromCR(&instance, namespace),
	)
	if err != nil {
		return err, "cannot delete rolebinding"
	}

	err = r.DeleteResourceIfExists(
		ctx,
		r.statefulSetFromCR(&instance, namespace, &loggerv1.M2LogstashPipelineList{}),
	)
	if err != nil {
		return err, "cannot delete statefulset"
	}
	err = r.DeleteResourceIfExists(
		ctx,
		r.clusterIPServiceFromCR(&instance, namespace, &loggerv1.M2LogstashPipelineList{}),
	)
	if err != nil {
		return err, "cannot delete clusterIP service"
	}

	err = r.DeleteResourceIfExists(
		ctx,
		headLessServiceFromCR(&instance, namespace),
	)
	if err != nil {
		return err, "cannot delete headless service"
	}

	controllerLog.Info("Resources deleted", "instance", instance.GetName())
	return nil, "Success"
}
