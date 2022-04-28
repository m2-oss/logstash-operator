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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type M2LogstashSpecLimit struct {
	Cpu    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

type M2LogstashSpecRequests struct {
	Cpu    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

type M2LogstashSpecResources struct {
	Limit    M2LogstashSpecLimit    `json:"limits,omitempty"`
	Requests M2LogstashSpecRequests `json:"requests,omitempty"`
}

// M2LogstashSpec defines the desired state of M2Logstash
type M2LogstashSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of M2Logstash. Edit M2Logstash_types.go to remove/update
	Image     string                  `json:"image,omitempty"`
	Tag       string                  `json:"tag,omitempty"`
	Name      string                  `json:"name,omitempty"`
	JavaOPTS  string                  `json:"java_opts,omitempty"`
	Resources M2LogstashSpecResources `json:"resources,omitempty"`
}

// +patchMergeKey=type
// +patchStrategy=merge
// +listType=map
// +listMapKey=type
// M2LogstashStatus defines the observed state of M2Logstash

type M2LogstashStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// M2Logstash is the Schema for the m2logstashes API
type M2Logstash struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   M2LogstashSpec   `json:"spec,omitempty"`
	Status M2LogstashStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// M2LogstashList contains a list of M2Logstash
type M2LogstashList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []M2Logstash `json:"items"`
}

func init() {
	SchemeBuilder.Register(&M2Logstash{}, &M2LogstashList{})
	SchemeBuilder.Register(&M2LogstashPipeline{}, &M2LogstashPipelineList{})
}
