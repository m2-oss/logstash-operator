package controllers

import (
	"bytes"
	"context"
	"fmt"
	loggerv1 "github.com/m2-oss/logstash-operator/api/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"text/template"
)

type PipelineInputKafka struct {
	Topic   string
	GroupID string
}

type PipelineInputS3 struct {
	Bucket string
}

type PipelineInput struct {
	Kafka PipelineInputKafka
	S3    PipelineInputS3
}

type PipelineOutputElasticSearch struct {
	RolloverAlias string
	Policy        string
}

type PipelineOutputS3 struct {
	Bucket string
}
type PipelineOutput struct {
	Elasticsearch PipelineOutputElasticSearch
	S3            PipelineOutputS3
}

type LogstashPipeline struct {
	CRD    loggerv1.M2LogstashPipelineSpec
	Input  PipelineInput
	Output PipelineOutput
}

type LogstashDLQ struct {
	DLQ string
}

type LogstashPipelines struct {
	Pipelines []string
}

// generalConfigMapFromCR create ConfigMap with general configuration
func (r *M2LogstashPipelineReconciler) generalConfigMapFromCR(instance *loggerv1.M2LogstashPipeline, nameSpace string) *corev1.ConfigMap {
	data, err := r.getGeneralConfigData(instance)
	if err != nil {
		return nil
	}

	return &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        fmt.Sprint(generalConfigMapName, "-", nameSpace),
			Namespace:   instance.GetNamespace(),
			Labels:      getLabels(nameSpace),
			Annotations: getAnnotations(),
		},
		Data: data,
	}
}

// generalPipelineMapFromCR create ConfigMap for general pipeline configuration
func (r *M2LogstashPipelineReconciler) generalPipelineMapFromCR(pipelines []string, instNamespace string, nameSpace string) *corev1.ConfigMap {
	data, err := r.getGeneralPipelineData(pipelines)
	if err != nil {
		return nil
	}

	return &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        fmt.Sprint(generalPipelineName, "-", nameSpace),
			Namespace:   instNamespace,
			Labels:      getLabels(nameSpace),
			Annotations: getAnnotations(),
		},
		Data: data,
	}
}

// pipelineConfigMapFromCR create ConfgMap pipelines
func (r *M2LogstashPipelineReconciler) pipelineConfigMapFromCR(instance *loggerv1.M2LogstashPipeline, instNamespace string, nameSpace string) *corev1.ConfigMap {
	data, err := r.getPipelineConfigData(instance, nameSpace)
	if err != nil {
		return nil
	}

	return &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        fmt.Sprint(pipelineName, "-", nameSpace, "-", instance.GetName()),
			Namespace:   instNamespace,
			Labels:      getLabels(nameSpace),
			Annotations: getAnnotations(),
		},
		Data: data,
	}
}

// getGeneralConfigData create general Logstash configuration
func (r *M2LogstashPipelineReconciler) getGeneralConfigData(instance *loggerv1.M2LogstashPipeline) (map[string]string, error) {
	log := r.Log.WithName("creator").WithName("configmap")
	controllerLog.Info("Get general configmap data")

	var data = make(map[string]string)
	var logstashTpl bytes.Buffer

	templ, err := template.ParseFiles("templates/logstash.yml")
	if err != nil {
		log.Error(err, "failed parse config template", "template", "logstash.yml")
		return nil, err
	}
	dlq := LogstashDLQ{"false"}
	if instance.Spec.Input.DLQ.Path != "" {
		dlq = LogstashDLQ{"true"}
	}

	if err := templ.Execute(&logstashTpl, dlq); err != nil {
		log.Error(err, "failed generate config file", "template", "logstash.yml")
		return nil, err
	}
	data["logstash.yml"] = logstashTpl.String()

	return data, nil
}

// getGeneralPipelineData create general Logstash pipeline configuration
func (r *M2LogstashPipelineReconciler) getGeneralPipelineData(pipelines []string) (map[string]string, error) {
	log := r.Log.WithName("creator").WithName("configmap")
	controllerLog.Info("Get general pipeline data")
	var data = make(map[string]string)
	var pipelinesTpl bytes.Buffer

	templ, err := template.ParseFiles("templates/pipelines.yml")
	if err != nil {
		log.Error(err, "failed parse config template", "template", "pipelines.yml")
		return nil, err
	}

	pipelinesList := LogstashPipelines{pipelines}
	if err := templ.Execute(&pipelinesTpl, pipelinesList); err != nil {
		log.Error(err, "failed generate config file", "template", "pipelines.yml")
		return nil, err
	}
	data["pipelines.yml"] = pipelinesTpl.String()

	return data, nil
}

// getPipelineConfigData create Logstash pipeline
func (r *M2LogstashPipelineReconciler) getPipelineConfigData(instance *loggerv1.M2LogstashPipeline, namespaceName string) (map[string]string, error) {
	log := r.Log.WithName("creator").WithName("configmap")
	controllerLog.Info("Get general configmap data", "instance", instance)
	var data = make(map[string]string)
	var logstashTpl bytes.Buffer

	templ, err := template.ParseFiles("templates/logstash.conf")
	if err != nil {
		log.Error(err, "failed parse config template", "template", "logstash.conf")
		return nil, err
	}

	namespace := &corev1.Namespace{}
	if err := r.GetClient().Get(context.TODO(), types.NamespacedName{
		Namespace: "",
		Name:      namespaceName,
	}, namespace); err != nil {
		log.Error(err, "failed get namespace object", namespaceName)
		return nil, err
	}

	config := LogstashPipeline{
		CRD: instance.Spec,
		Input: PipelineInput{
			Kafka: PipelineInputKafka{
				Topic:   getTopic(namespace),
				GroupID: fmt.Sprint("logstash_", namespaceName),
			},
			S3: PipelineInputS3{
				Bucket: fmt.Sprint(instance.GetName(), "-", namespaceName),
			},
		},
		Output: PipelineOutput{
			Elasticsearch: PipelineOutputElasticSearch{
				RolloverAlias: getAlias(namespace),
				Policy:        getPolicy(namespace),
			},
			S3: PipelineOutputS3{
				Bucket: fmt.Sprint(instance.GetName(), "-", namespaceName),
			},
		},
	}

	// controllerLog.Info("Rendering Pipeline", "config", config)
	if err := templ.Execute(&logstashTpl, config); err != nil {
		log.Error(err, "failed generate config file", "template", "logstash.conf")
		return nil, err
	}
	data[fmt.Sprint(instance.GetName(), ".conf")] = logstashTpl.String()

	return data, nil
}
