package v1

import (
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var log = logf.Log.WithName("m2logstash-resource")

func (r *M2Logstash) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-logger-m2-ru-v1-m2logstash,mutating=true,failurePolicy=fail,sideEffects=None,groups=logger.m2.ru,resources=m2logstashes,verbs=create;update,versions=v1,name=mm2logstash.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Defaulter = &M2Logstash{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *M2Logstash) Default() {
	log.Info("default", "name", r.Name)
	if r.Spec.Name == "" {
		r.Spec.Name = "logstash"
	}
	if r.Spec.Image == "" {
		r.Spec.Image = "logstash"
	}
	if r.Spec.Tag == "" {
		r.Spec.Tag = "7.16.2"
	}
	if r.Spec.Resources.Limit.Cpu == "" {
		r.Spec.Resources.Limit.Cpu = "2000m"
	}
	if r.Spec.Resources.Limit.Memory == "" {
		r.Spec.Resources.Limit.Memory = "1536Mi"
	}
	if r.Spec.Resources.Requests.Cpu == "" {
		r.Spec.Resources.Requests.Cpu = "100m"
	}
	if r.Spec.Resources.Requests.Memory == "" {
		r.Spec.Resources.Requests.Memory = "1536Mi"
	}
	if r.Spec.JavaOPTS == "" {
		r.Spec.JavaOPTS = "-Xmx1g -Xms1g"
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-logger-m2-ru-v1-m2logstash,mutating=false,failurePolicy=fail,sideEffects=None,groups=logger.m2.ru,resources=m2logstashes,verbs=create;update,versions=v1,name=vm2logstash.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &M2Logstash{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *M2Logstash) ValidateCreate() error {
	log.Info("validate create", "name", r.Name)
	return r.validateM2Logstash()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *M2Logstash) ValidateUpdate(old runtime.Object) error {
	log.Info("validate update", "name", r.Name)

	return r.validateM2Logstash()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *M2Logstash) ValidateDelete() error {
	log.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

func (r *M2Logstash) validateM2Logstash() error {
	var allErrs field.ErrorList
	if err := r.validateM2LogstashName(); err != nil {
		allErrs = append(allErrs, err)
	}
	//todo: validation Spec
	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "logger.m2.ru", Kind: "M2Logstash"},
		r.Name, allErrs)
}

func (r *M2Logstash) validateM2LogstashName() *field.Error {
	if len(r.ObjectMeta.Name) > validation.DNS1123LabelMaxLength {
		return field.Invalid(field.NewPath("metadata").Child("name"), r.Name, "must be no more than 63 characters")
	}
	return nil
}
