package controllers

import (
	"fmt"
	loggerv1 "github.com/m2-oss/logstash-operator/api/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"path/filepath"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// getPolicy get Elasticsearch ILM policy from namespace label
func getPolicy(namespace *corev1.Namespace) string {
	var Policy = "7-days-rollout"
	if label, ok := namespace.GetLabels()[ilmLabel]; ok {
		Policy = label
	}

	return Policy
}

// getAlias get Elasticsearch index alias from namespace label
func getAlias(namespace *corev1.Namespace) string {
	var Alias = namespace.GetName()
	if label, ok := namespace.GetLabels()[typeLabel]; ok {
		Alias = fmt.Sprint(label, "-", namespace.GetName())
	}

	return Alias
}

// getTopic get Kafka topic name from namespace name
func getTopic(namespace *corev1.Namespace) string {
	var Topic = namespace.GetName()

	return Topic
}

// getLabels get common labels
func getLabels(namespace string) map[string]string {
	return map[string]string{
		"operator":                   "logstash",
		"control-plane":              "logstash-operator",
		"app.kubernetes.io/instance": fmt.Sprint("logstash-", namespace),
		"app.kubernetes.io/name":     "logstash",
	}
}

// getAnnotations get common annotations
func getAnnotations() map[string]string {
	return map[string]string{
		"reloader.stakater.com/match": "true",
	}
}

// getVolumes get list of K8S Volumes
func getVolumes(instanceName string, pipelineList *loggerv1.M2LogstashPipelineList, namespace string) []corev1.Volume {
	var volumes = []corev1.Volume{}

	if len(pipelineList.Items) > 0 {
		volumes = append(volumes, corev1.Volume{
			Name: fmt.Sprint(generalConfigMapName, "-", namespace),
			VolumeSource: corev1.VolumeSource{ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: fmt.Sprint(generalConfigMapName, "-", namespace),
				},
			}},
		})
		volumes = append(volumes, corev1.Volume{
			Name: fmt.Sprint(generalPipelineName, "-", namespace),
			VolumeSource: corev1.VolumeSource{ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: fmt.Sprint(generalPipelineName, "-", namespace),
				},
			}},
		})
	}
	for _, item := range pipelineList.Items {
		volumes = append(volumes, corev1.Volume{
			Name: fmt.Sprint(instanceName, "-pipeline-", item.GetName()),
			VolumeSource: corev1.VolumeSource{ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: fmt.Sprint(pipelineName, "-", namespace, "-", item.GetName()),
				},
			}},
		})
		if item.Spec.Output.Elasticsearch.Cacert != "" {
			volumes = append(volumes, corev1.Volume{
				Name: fmt.Sprint(instanceName, "-certificate-", item.GetName()),
				VolumeSource: corev1.VolumeSource{Secret: &corev1.SecretVolumeSource{
					SecretName: fmt.Sprint(elasticCacert, "-", item.GetName()),
				}},
			})
		}
	}

	return volumes
}

// getSecrets get list of K8S Secrets
func getSecrets(pipelineList *loggerv1.M2LogstashPipelineList) []corev1.EnvFromSource {
	var secrets = []corev1.EnvFromSource{}

	for _, item := range pipelineList.Items {
		var secretName = ""
		if item.Spec.Input.HTTP.Secret != "" {
			secretName = item.Spec.Input.HTTP.Secret
		} else if item.Spec.Input.Kafka.Secret != "" {
			secretName = item.Spec.Input.Kafka.Secret
		} else if item.Spec.Input.S3.Secret != "" {
			secretName = item.Spec.Input.S3.Secret
		}
		if secretName != "" {
			secrets = append(secrets, getSecretRef(secretName))
		}

		secretName = ""
		if item.Spec.Output.Elasticsearch.Secret != "" {
			secretName = item.Spec.Output.Elasticsearch.Secret
		} else if item.Spec.Output.S3.Secret != "" {
			secretName = item.Spec.Output.S3.Secret
		}

		if secretName != "" {
			secrets = append(secrets, getSecretRef(secretName))
		}
	}
	return secrets
}

// getSecretRef get K8S secret ref from logstash pipeline
func getSecretRef(secretName string) corev1.EnvFromSource {
	var secret = corev1.EnvFromSource{}
	var optionalSecret = true

	secret = corev1.EnvFromSource{
		SecretRef: &corev1.SecretEnvSource{
			LocalObjectReference: corev1.LocalObjectReference{
				Name: secretName,
			},
			Optional: &optionalSecret,
		},
	}

	return secret
}

// getVolumeMounts get list of K8S STS volume mounts
func getVolumeMounts(instanceName string, pipelineList *loggerv1.M2LogstashPipelineList, namespace string) []corev1.VolumeMount {
	var mounts = []corev1.VolumeMount{}

	if len(pipelineList.Items) > 0 {
		mounts = []corev1.VolumeMount{{
			Name:      fmt.Sprint(generalConfigMapName, "-", namespace),
			MountPath: "/usr/share/logstash/config/logstash.yml",
			SubPath:   "logstash.yml",
		}, {
			Name:      fmt.Sprint(generalPipelineName, "-", namespace),
			MountPath: "/usr/share/logstash/config/pipelines.yml",
			SubPath:   "pipelines.yml",
		}}
	}
	for _, item := range pipelineList.Items {
		mounts = append(mounts, corev1.VolumeMount{
			Name:      fmt.Sprint(instanceName, "-pipeline-", item.GetName()),
			MountPath: fmt.Sprint("/usr/share/logstash/pipeline/", item.GetName(), ".conf"),
			SubPath:   fmt.Sprint(item.GetName(), ".conf"),
		})
		if item.Spec.Output.Elasticsearch.Cacert != "" {
			mounts = append(mounts, corev1.VolumeMount{
				Name:      fmt.Sprint(instanceName, "-certificate-", item.GetName()),
				MountPath: filepath.Dir(item.Spec.Output.Elasticsearch.Cacert),
				ReadOnly:  true,
			})
		}
	}
	return mounts
}

// getServicePorts - get list of K8S Service ports
func getServicePorts(pipelineList *loggerv1.M2LogstashPipelineList) []corev1.ServicePort {
	var ports = []corev1.ServicePort{{
		Port:       9304,
		Protocol:   corev1.ProtocolTCP,
		TargetPort: intstr.FromInt(9304),
		Name:       "exporter",
	}}

	for _, item := range pipelineList.Items {
		if item.Spec.Input.Beats.Port != 0 {
			ports = append(ports, corev1.ServicePort{
				Port:       item.Spec.Input.Beats.Port,
				Protocol:   corev1.ProtocolTCP,
				TargetPort: intstr.FromInt(int(item.Spec.Input.Beats.Port)),
				Name:       "beats",
			})
		} else if item.Spec.Input.HTTP.Port != 0 {
			ports = append(ports, corev1.ServicePort{
				Port:       item.Spec.Input.HTTP.Port,
				Protocol:   corev1.ProtocolTCP,
				TargetPort: intstr.FromInt(int(item.Spec.Input.HTTP.Port)),
				Name:       "http",
			})
		} else if item.Spec.Input.TCP.Port != 0 {
			ports = append(ports, corev1.ServicePort{
				Port:       item.Spec.Input.TCP.Port,
				Protocol:   corev1.ProtocolTCP,
				TargetPort: intstr.FromInt(int(item.Spec.Input.TCP.Port)),
				Name:       "tcp",
			})
		} else if item.Spec.Input.UDP.Port != 0 {
			ports = append(ports, corev1.ServicePort{
				Port:       item.Spec.Input.UDP.Port,
				Protocol:   corev1.ProtocolUDP,
				TargetPort: intstr.FromInt(int(item.Spec.Input.UDP.Port)),
				Name:       "udp",
			})
		}
	}
	return ports
}

// getContainerPorts get list of container ports
func getContainerPorts(pipelineList *loggerv1.M2LogstashPipelineList) []corev1.ContainerPort {
	var ports = []corev1.ContainerPort{{
		ContainerPort: 9600,
		Name:          "http",
		Protocol:      corev1.ProtocolTCP,
	}}

	for _, item := range pipelineList.Items {
		if item.Spec.Input.Beats.Port != 0 {
			ports = append(ports, corev1.ContainerPort{
				ContainerPort: item.Spec.Input.Beats.Port,
				Protocol:      corev1.ProtocolTCP,
				Name:          "beats",
			})
		} else if item.Spec.Input.HTTP.Port != 0 {
			ports = append(ports, corev1.ContainerPort{
				ContainerPort: item.Spec.Input.HTTP.Port,
				Protocol:      corev1.ProtocolTCP,
				Name:          "http",
			})
		} else if item.Spec.Input.TCP.Port != 0 {
			ports = append(ports, corev1.ContainerPort{
				ContainerPort: item.Spec.Input.TCP.Port,
				Protocol:      corev1.ProtocolTCP,
				Name:          "tcp",
			})
		} else if item.Spec.Input.UDP.Port != 0 {
			ports = append(ports, corev1.ContainerPort{
				ContainerPort: item.Spec.Input.UDP.Port,
				Protocol:      corev1.ProtocolUDP,
				Name:          "udp",
			})
		}
	}
	return ports
}

// namespaceFilter event filter
func namespaceFilter() predicate.Predicate {
	var response bool

	return predicate.Funcs{
		CreateFunc: func(event event.CreateEvent) bool {
			response = false
			if _, ok := event.Object.(*corev1.Namespace); ok {
				if val, ok := event.Object.GetLabels()[label]; ok {
					if val == "true" {
						response = true
					}
				}
			}
			if _, ok := event.Object.(*loggerv1.M2Logstash); ok {
				response = true
			}
			if _, ok := event.Object.(*loggerv1.M2LogstashPipeline); ok {
				response = true
			}
			return response
		},
		UpdateFunc: func(updateEvent event.UpdateEvent) bool {
			response = false

			_, oldObject := updateEvent.ObjectOld.(*corev1.Namespace)
			_, newObject := updateEvent.ObjectNew.(*corev1.Namespace)
			if oldObject && newObject {
				oldValue, _ := updateEvent.ObjectOld.GetLabels()[label]
				newValue, _ := updateEvent.ObjectNew.GetLabels()[label]

				old := oldValue == "true"
				new := newValue == "true"

				response = old != new

				if oldValue == "true" && newValue == "true" {
					response = true
				}
			}

			_, oldObject = updateEvent.ObjectOld.(*loggerv1.M2Logstash)
			_, newObject = updateEvent.ObjectNew.(*loggerv1.M2Logstash)
			if oldObject && newObject {
				response = true
			}

			_, oldObject = updateEvent.ObjectOld.(*loggerv1.M2LogstashPipeline)
			_, newObject = updateEvent.ObjectNew.(*loggerv1.M2LogstashPipeline)
			if oldObject && newObject {
				response = true
			}
			return response
		},

		DeleteFunc: func(deleteEvent event.DeleteEvent) bool {
			response = false
			_, ok := deleteEvent.Object.(*corev1.Namespace)
			if ok {
				response = true
			}
			_, ok = deleteEvent.Object.(*loggerv1.M2Logstash)
			if ok {
				response = true
			}
			_, ok = deleteEvent.Object.(*loggerv1.M2LogstashPipeline)
			if ok {
				response = true
			}
			return response
		},
	}
}
