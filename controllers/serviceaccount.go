package controllers

import (
	"fmt"
	loggerv1 "github.com/m2-oss/logstash-operator/api/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// serviceAccountFromCR create ServiceAccount
func (r *M2LogstashReconciler) serviceAccountFromCR(instance *loggerv1.M2Logstash, nameSpace string) *corev1.ServiceAccount {
	return &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        fmt.Sprint(instance.Spec.Name, "-", nameSpace),
			Namespace:   instance.GetNamespace(),
			Labels:      getLabels(nameSpace),
			Annotations: getAnnotations(),
		},
	}
}

// roleFromCR create Role
func (r *M2LogstashReconciler) roleFromCR(instance *loggerv1.M2Logstash, nameSpace string) *rbacv1.Role {
	return &rbacv1.Role{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Role",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        fmt.Sprint(instance.Spec.Name, "-", nameSpace),
			Namespace:   instance.GetNamespace(),
			Labels:      getLabels(nameSpace),
			Annotations: getAnnotations(),
		},
		Rules: []rbacv1.PolicyRule{{
			Verbs:         []string{"use"},
			APIGroups:     []string{"extensions"},
			Resources:     []string{"podsecuritypolicies"},
			ResourceNames: []string{fmt.Sprint(instance.Spec.Name, "-", nameSpace)},
		}},
	}
}

// roleBindingFromCR create RoleBinding
func (r *M2LogstashReconciler) roleBindingFromCR(instance *loggerv1.M2Logstash, nameSpace string) *rbacv1.RoleBinding {
	return &rbacv1.RoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "RoleBinding",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        fmt.Sprint(instance.Spec.Name, "-", nameSpace),
			Namespace:   instance.GetNamespace(),
			Labels:      getLabels(nameSpace),
			Annotations: getAnnotations(),
		},
		Subjects: []rbacv1.Subject{{
			Kind:      "ServiceAccount",
			Name:      fmt.Sprint(instance.Spec.Name, "-", nameSpace),
			Namespace: instance.GetNamespace(),
		}},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     fmt.Sprint(instance.Spec.Name, "-", nameSpace),
		},
	}
}
