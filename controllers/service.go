package controllers

import (
	"context"
	"fmt"
	loggerv1 "github.com/m2-oss/logstash-operator/api/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// clusterIPServiceFromCR create ClusterIP Service
func (r *M2LogstashReconciler) clusterIPServiceFromCR(instance *loggerv1.M2Logstash, nameSpace string, pipelineList *loggerv1.M2LogstashPipelineList) *corev1.Service {

	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprint(instance.Spec.Name, "-", nameSpace),
			Namespace: instance.GetNamespace(),
			Labels:    getLabels(nameSpace),
		},
		Spec: corev1.ServiceSpec{
			Selector: getLabels(nameSpace),
			Ports:    getServicePorts(pipelineList),
			Type:     corev1.ServiceTypeClusterIP,
		},
	}
}

// clusterIPServiceFromCR create or update ClusterIP Service
func (r *M2LogstashPipelineReconciler) clusterIPServiceFromCR(instance *loggerv1.M2Logstash, nameSpace string, pipelineList *loggerv1.M2LogstashPipelineList) *corev1.Service {
	service := &corev1.Service{}

	desired := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprint(instance.Spec.Name, "-", nameSpace),
			Namespace: instance.GetNamespace(),
			Labels:    getLabels(nameSpace),
		},
		Spec: corev1.ServiceSpec{
			Selector: getLabels(nameSpace),
			Ports:    getServicePorts(pipelineList),
			Type:     corev1.ServiceTypeClusterIP,
		},
	}

	if err := r.GetClient().Get(
		context.TODO(),
		types.NamespacedName{
			Namespace: instance.GetNamespace(),
			Name:      fmt.Sprint(instance.Spec.Name, "-", nameSpace),
		},
		service,
	); err == nil {
		desired.Spec.ClusterIP = service.Spec.ClusterIP
	}

	return desired
}

// headLessServiceFromCR create Headless Service
func headLessServiceFromCR(instance *loggerv1.M2Logstash, nameSpace string) *corev1.Service {
	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprint(instance.Spec.Name, "-", nameSpace, "-headless"),
			Namespace: instance.GetNamespace(),
		},
		Spec: corev1.ServiceSpec{
			Selector: getLabels(nameSpace),
			Ports: []corev1.ServicePort{{
				Port:       9600,
				Protocol:   corev1.ProtocolTCP,
				TargetPort: intstr.FromInt(9600),
				Name:       "http",
			}},
			Type:      corev1.ServiceTypeClusterIP,
			ClusterIP: "None",
		},
		Status: corev1.ServiceStatus{},
	}
}
