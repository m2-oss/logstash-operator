package v1

import (
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *M2LogstashPipeline) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-logger-m2-ru-v1-m2logstashpipeline,mutating=true,failurePolicy=fail,sideEffects=None,groups=logger.m2.ru,resources=m2logstashpipelines,verbs=create;update,versions=v1,name=mm2logstashpipeline.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Defaulter = &M2LogstashPipeline{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *M2LogstashPipeline) Default() {
	log.Info("default", "name", r.Name)
	if r.Spec.Input.Beats.ID == "" {
		r.Spec.Input.Beats.ID = "beats"
	}
	if r.Spec.Input.TCP.ID == "" {
		r.Spec.Input.TCP.ID = "tcp"
	}
	if r.Spec.Input.UDP.ID == "" {
		r.Spec.Input.UDP.ID = "udp"
	}
	if r.Spec.Input.HTTP.ID == "" {
		r.Spec.Input.HTTP.ID = "http"
	}
	if r.Spec.Input.Kafka.ID == "" {
		r.Spec.Input.Kafka.ID = "kafka"
	}
	if r.Spec.Input.S3.ID == "" {
		r.Spec.Input.S3.ID = "s3"
	}
	if r.Spec.Input.DLQ.ID == "" {
		r.Spec.Input.DLQ.ID = "dlq"
	}
	if r.Spec.Input.Kafka.DecorateEvents == "" {
		r.Spec.Input.Kafka.DecorateEvents = "none"
	}
	if r.Spec.Input.Kafka.AutoOffsetReset == "" {
		r.Spec.Input.Kafka.AutoOffsetReset = "earliest"
	}
	if r.Spec.Input.Kafka.ConsumerThreads == 0 {
		r.Spec.Input.Kafka.ConsumerThreads = 1
	}
	if r.Spec.Input.Kafka.SaslMechanism == "" {
		r.Spec.Input.Kafka.SaslMechanism = "GSSAPI"
	}
	if r.Spec.Input.Kafka.SecurityProtocol == "" {
		r.Spec.Input.Kafka.SecurityProtocol = "PLAINTEXT"
	}
	if r.Spec.Input.S3.ExcludePattern == "" {
		r.Spec.Input.S3.ExcludePattern = "nil"
	}
	if r.Spec.Input.S3.GzipPattern == "" {
		r.Spec.Input.S3.GzipPattern = "\\.gz(ip)?$"
	}
	if r.Spec.Input.S3.Prefix == "" {
		r.Spec.Input.S3.GzipPattern = "nil"
	}
	if r.Spec.Input.S3.Region == "" {
		r.Spec.Input.S3.Region = "us-east-1"
	}
	if r.Spec.Input.DLQ.PipelineID == "" {
		r.Spec.Input.DLQ.PipelineID = "main"
	}
	if r.Spec.Output.S3.Region == "" {
		r.Spec.Output.S3.Region = "us-east-1"
	}
	if r.Spec.Output.S3.CannedACL == "" {
		r.Spec.Output.S3.CannedACL = "private"
	}
	if r.Spec.Output.S3.Encoding == "" {
		r.Spec.Output.S3.Encoding = "none"
	}
	if r.Spec.Output.S3.RotationStrategy == "" {
		r.Spec.Output.S3.RotationStrategy = "size_and_time"
	}
	if r.Spec.Output.S3.SizeFile == 0 {
		r.Spec.Output.S3.SizeFile = 5242880
	}
	if r.Spec.Output.S3.TimeFile == 0 {
		r.Spec.Output.S3.TimeFile = 15
	}
	if r.Spec.Output.S3.UploadWorkerCount == 0 {
		r.Spec.Output.S3.UploadWorkerCount = 4
	}
	if r.Spec.Output.Graphite.ReconnectInterval == 0 {
		r.Spec.Output.Graphite.ReconnectInterval = 2
	}
	if r.Spec.Output.Graphite.TimestampField == "" {
		r.Spec.Output.Graphite.TimestampField = "@timestamp"
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-logger-m2-ru-v1-m2logstashpipeline,mutating=false,failurePolicy=fail,sideEffects=None,groups=logger.m2.ru,resources=m2logstashpipelines,verbs=create;update,versions=v1,name=vm2logstashpipeline.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &M2LogstashPipeline{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *M2LogstashPipeline) ValidateCreate() error {
	log.Info("validate create", "name", r.Name)
	return r.validateM2LogstashPipeline()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *M2LogstashPipeline) ValidateUpdate(old runtime.Object) error {
	log.Info("validate update", "name", r.Name)

	return r.validateM2LogstashPipeline()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *M2LogstashPipeline) ValidateDelete() error {
	log.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

func (r *M2LogstashPipeline) validateM2LogstashPipeline() error {
	var allErrs field.ErrorList
	if err := r.validateM2LogstashPipelineName(); err != nil {
		allErrs = append(allErrs, err)
	}
	//todo: validation Spec
	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "logger.m2.ru", Kind: "M2LogstashPipeline"},
		r.Name, allErrs)
}

func (r *M2LogstashPipeline) validateM2LogstashPipelineName() *field.Error {
	if len(r.ObjectMeta.Name) > validation.DNS1123LabelMaxLength {
		return field.Invalid(field.NewPath("metadata").Child("name"), r.Name, "must be no more than 63 characters")
	}
	return nil
}
