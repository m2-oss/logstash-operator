package controllers

import (
	"context"
	"fmt"
	loggerv1 "github.com/m2-oss/logstash-operator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"strconv"
)

const (
	generalConfigMapName = "general-logstash-config"
	generalPipelineName  = "general-logstash-pipeline"
	pipelineName         = "logstash-pipeline"
	elasticCacert        = "logstash-elasticsearch-ca"
)

// statefulSetFromCR create StatefulSet
func (r *M2LogstashReconciler) statefulSetFromCR(instance *loggerv1.M2Logstash, nameSpace string, pipelineList *loggerv1.M2LogstashPipelineList) *appsv1.StatefulSet {
	replicas := r.getReplicaCount(nameSpace)

	return &appsv1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "StatefulSet",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        fmt.Sprint(instance.Spec.Name, "-", nameSpace),
			Namespace:   instance.GetNamespace(),
			Labels:      getLabels(nameSpace),
			Annotations: getAnnotations(),
		},
		Spec: appsv1.StatefulSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: getLabels(nameSpace),
			},
			Replicas: &replicas,
			UpdateStrategy: appsv1.StatefulSetUpdateStrategy{
				Type: appsv1.RollingUpdateStatefulSetStrategyType,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: getLabels(nameSpace),
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:  instance.Spec.Name,
						Image: fmt.Sprint(instance.Spec.Image, ":", instance.Spec.Tag),
						Ports: getContainerPorts(pipelineList),
						Resources: corev1.ResourceRequirements{
							Limits: corev1.ResourceList{
								"cpu":    resource.MustParse(instance.Spec.Resources.Limit.Cpu),
								"memory": resource.MustParse(instance.Spec.Resources.Limit.Memory),
							},
							Requests: corev1.ResourceList{
								"cpu":    resource.MustParse(instance.Spec.Resources.Requests.Cpu),
								"memory": resource.MustParse(instance.Spec.Resources.Requests.Memory),
							},
						},
						VolumeMounts: getVolumeMounts(instance.Spec.Name, pipelineList, nameSpace),
						ReadinessProbe: &corev1.Probe{
							Handler: corev1.Handler{Exec: &corev1.ExecAction{
								Command: []string{
									"curl",
									"localhost:9600",
								},
							}},
						},
						EnvFrom: getSecrets(pipelineList),
						Env: []corev1.EnvVar{{
							Name:  "LS_JAVA_OPTS",
							Value: instance.Spec.JavaOPTS,
						}},
					}, {
						Name:  fmt.Sprint(instance.Spec.Name, "-exporter"),
						Image: "cr.yandex/crp29rd1alarj2e8jmp5/vtblife/devops/logstash_exporter:v0.1.2",
						Ports: []corev1.ContainerPort{{
							ContainerPort: 9304,
							Name:          "exporter",
							Protocol:      corev1.ProtocolTCP,
						}},
						Resources: corev1.ResourceRequirements{
							Limits: corev1.ResourceList{
								"cpu":    resource.MustParse("300m"),
								"memory": resource.MustParse("200Mi"),
							},
							Requests: corev1.ResourceList{
								"cpu":    resource.MustParse("100m"),
								"memory": resource.MustParse("200Mi"),
							},
						},
						Command: []string{
							"/logstash_exporter",
							"--logstash.endpoint",
							"http://127.0.0.1:9600",
							"--web.listen-address",
							":9304",
						},
					}},
					Volumes:            getVolumes(instance.Spec.Name, pipelineList, nameSpace),
					ServiceAccountName: fmt.Sprint(instance.Spec.Name, "-", nameSpace),
				},
			},
		},
	}
}

// statefulSetFromCR create StatefulSet with pipelines
func (r *M2LogstashPipelineReconciler) statefulSetFromCR(instance *loggerv1.M2Logstash, nameSpace string, pipelineList *loggerv1.M2LogstashPipelineList) *appsv1.StatefulSet {
	replicas := r.getReplicaCount(nameSpace)

	return &appsv1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "StatefulSet",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        fmt.Sprint(instance.Spec.Name, "-", nameSpace),
			Namespace:   instance.GetNamespace(),
			Labels:      getLabels(nameSpace),
			Annotations: getAnnotations(),
		},
		Spec: appsv1.StatefulSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: getLabels(nameSpace),
			},
			Replicas: &replicas,
			UpdateStrategy: appsv1.StatefulSetUpdateStrategy{
				Type: appsv1.RollingUpdateStatefulSetStrategyType,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: getLabels(nameSpace),
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:  instance.Spec.Name,
						Image: fmt.Sprint(instance.Spec.Image, ":", instance.Spec.Tag),
						Ports: getContainerPorts(pipelineList),
						Resources: corev1.ResourceRequirements{
							Limits: corev1.ResourceList{
								"cpu":    resource.MustParse(instance.Spec.Resources.Limit.Cpu),
								"memory": resource.MustParse(instance.Spec.Resources.Limit.Memory),
							},
							Requests: corev1.ResourceList{
								"cpu":    resource.MustParse(instance.Spec.Resources.Requests.Cpu),
								"memory": resource.MustParse(instance.Spec.Resources.Requests.Memory),
							},
						},
						VolumeMounts: getVolumeMounts(instance.Spec.Name, pipelineList, nameSpace),
						ReadinessProbe: &corev1.Probe{
							Handler: corev1.Handler{Exec: &corev1.ExecAction{
								Command: []string{
									"curl",
									"localhost:9600",
								},
							}},
						},
						EnvFrom: getSecrets(pipelineList),
						Env: []corev1.EnvVar{{
							Name:  "LS_JAVA_OPTS",
							Value: instance.Spec.JavaOPTS,
						}},
					}, {
						Name:  fmt.Sprint(instance.Spec.Name, "-exporter"),
						Image: "cr.yandex/crp29rd1alarj2e8jmp5/vtblife/devops/logstash_exporter:v0.1.2",
						Ports: []corev1.ContainerPort{{
							ContainerPort: 9304,
							Name:          "exporter",
							Protocol:      corev1.ProtocolTCP,
						}},
						Resources: corev1.ResourceRequirements{
							Limits: corev1.ResourceList{
								"cpu":    resource.MustParse("300m"),
								"memory": resource.MustParse("200Mi"),
							},
							Requests: corev1.ResourceList{
								"cpu":    resource.MustParse("100m"),
								"memory": resource.MustParse("200Mi"),
							},
						},
						Command: []string{
							"/logstash_exporter",
							"--logstash.endpoint",
							"http://127.0.0.1:9600",
							"--web.listen-address",
							":9304",
						},
					}},
					Volumes:            getVolumes(instance.Spec.Name, pipelineList, nameSpace),
					ServiceAccountName: fmt.Sprint(instance.Spec.Name, "-", nameSpace),
				},
			},
		},
	}
}

// getReplicaCount get Replicas from namespace label
func (r *M2LogstashReconciler) getReplicaCount(namespaceName string) int32 {
	var Count = 1
	namespace := &corev1.Namespace{}

	if err := r.GetClient().Get(context.TODO(), types.NamespacedName{
		Namespace: "",
		Name:      namespaceName,
	}, namespace); err != nil {
		r.Log.Error(err, "failed get namespace object", namespaceName)
	}

	if annotation, ok := namespace.GetAnnotations()[replicasAnnotation]; ok {
		Count, _ = strconv.Atoi(annotation)
	}

	return int32(Count)
}

// getReplicaCount get Replicas from namespace label
func (r *M2LogstashPipelineReconciler) getReplicaCount(namespaceName string) int32 {
	var Count = 1
	namespace := &corev1.Namespace{}

	if err := r.GetClient().Get(context.TODO(), types.NamespacedName{
		Namespace: "",
		Name:      namespaceName,
	}, namespace); err != nil {
		r.Log.Error(err, "failed get namespace object", namespaceName)
	}

	if annotation, ok := namespace.GetAnnotations()[replicasAnnotation]; ok {
		Count, _ = strconv.Atoi(annotation)
	}

	return int32(Count)
}
