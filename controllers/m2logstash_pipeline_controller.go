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

// M2LogstashPipelineReconciler reconciles a M2LogstashPipeline object
type M2LogstashPipelineReconciler struct {
	// client.Client
	Log logr.Logger
	util.ReconcilerBase
}

var controllerPipelineLog = ctrl.Log.WithName("controller").WithName("M2LogstashPipeline")

// +kubebuilder:rbac:groups=logger.m2.ru,resources=m2logstashpipelines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=logger.m2.ru,resources=m2logstashpipelines/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=logger.m2.ru,resources=m2logstashpipelines/finalizers,verbs=update

func (r *M2LogstashPipelineReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	controllerPipelineLog.Info("start reconcile", "request", req.NamespacedName)

	if req.NamespacedName.Namespace == "" {
		namespace := &corev1.Namespace{}
		if err := r.GetClient().Get(ctx, req.NamespacedName, namespace); err != nil {
			if apierrors.IsNotFound(err) {
				return ctrl.Result{}, nil
			}
			controllerPipelineLog.Error(err, "Cannot get object Namespace")
			return ctrl.Result{}, err
		}

		instanceList := &loggerv1.M2LogstashPipelineList{}
		err := r.GetClient().List(ctx, instanceList)
		if err != nil {
			return ctrl.Result{}, nil
		}

		for _, instance := range instanceList.Items {
			if namespace.GetLabels()[label] == "true" {
				err, response := r.syncPipelineResources(ctx, instance, req.NamespacedName.Name)
				if err != nil {
					controllerPipelineLog.Error(err, response)
					return r.ManageError(ctx, &instance, err)
				}
			} else {
				err, response := r.deletePipelineResources(ctx, instance, req.NamespacedName.Name)
				if err != nil {
					controllerPipelineLog.Error(err, response)
					return r.ManageError(ctx, &instance, err)
				}
			}
		}
	}
	instance := &loggerv1.M2LogstashPipeline{}
	if err := r.GetClient().Get(ctx, req.NamespacedName, instance); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		controllerPipelineLog.Error(err, "cannot get object M2LogstashPipeline")
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
			err, response := r.syncPipelineResources(ctx, *instance, item.GetName())
			if err != nil {
				controllerPipelineLog.Error(err, response)
				return r.ManageError(ctx, instance, err)
			}
		}
	}

	controllerPipelineLog.Info("finish reconcile", "request", req.NamespacedName)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *M2LogstashPipelineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&loggerv1.M2LogstashPipeline{}).
		Watches(
			&source.Kind{Type: &corev1.Namespace{}},
			&handler.EnqueueRequestForObject{}).
		WithEventFilter(namespaceFilter()).
		WithOptions(controller.Options{MaxConcurrentReconciles: 2}).
		Complete(r)
}

// SyncPipelineResources creating or updating resources related to M2LogstashPipeline
func (r *M2LogstashPipelineReconciler) syncPipelineResources(ctx context.Context, instance loggerv1.M2LogstashPipeline, namespace string) (error, string) {
	controllerPipelineLog.Info("Sync instance resources", "instance", instance.GetName())
	loggerList := &loggerv1.M2LogstashList{}
	r.GetClient().List(context.TODO(), loggerList)

	pipelineList := &loggerv1.M2LogstashPipelineList{}
	r.GetClient().List(context.TODO(), pipelineList)

	err := r.CreateOrUpdateResource(
		ctx,
		&instance,
		instance.GetNamespace(),
		r.generalConfigMapFromCR(&instance, namespace),
	)
	if err != nil {
		return err, "cannot create General configmap"
	}

	var pipelines = []string{
		instance.GetName(),
	}
	err = r.CreateOrUpdateResource(
		ctx,
		&instance,
		instance.GetNamespace(),
		r.generalPipelineMapFromCR(pipelines, instance.GetNamespace(), namespace),
	)
	if err != nil {
		return err, "cannot create general pipelines configmap"
	}

	err = r.CreateOrUpdateResource(
		ctx,
		&instance,
		instance.GetNamespace(),
		r.pipelineConfigMapFromCR(&instance, instance.GetNamespace(), namespace),
	)
	if err != nil {
		return err, "cannot create pipeline configmap"
	}

	err = r.CreateOrUpdateResource(
		ctx,
		&instance,
		instance.GetNamespace(),
		r.statefulSetFromCR(&loggerList.Items[0], namespace, pipelineList),
	)
	if err != nil {
		return err, "cannot update statefulset"
	}

	err = r.CreateOrUpdateResource(
		ctx,
		&instance,
		instance.GetNamespace(),
		r.clusterIPServiceFromCR(&loggerList.Items[0], namespace, pipelineList),
	)
	controllerPipelineLog.Info("Resources synced", "instance", instance.GetName())
	return nil, "Success"
}

// deletePipelineResources deleting resources related to M2LogstashPipeline
func (r *M2LogstashPipelineReconciler) deletePipelineResources(ctx context.Context, instance loggerv1.M2LogstashPipeline, namespace string) (error, string) {
	err := r.DeleteResourceIfExists(
		ctx,
		r.generalConfigMapFromCR(&instance, namespace),
	)
	if err != nil {
		return err, "cannot delete general configmap"
	}

	var pipelines = []string{
		instance.GetName(),
	}
	err = r.DeleteResourceIfExists(
		ctx,
		r.generalPipelineMapFromCR(pipelines, instance.GetNamespace(), namespace),
	)
	if err != nil {
		return err, "cannot delete general pipeline configmap"
	}

	err = r.DeleteResourceIfExists(
		ctx,
		r.pipelineConfigMapFromCR(&instance, instance.GetNamespace(), namespace),
	)
	if err != nil {
		return err, "cannot delete pipeline configmap"
	}

	return nil, "Success"
}
