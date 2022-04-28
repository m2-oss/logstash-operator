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

type M2LogstashInputSpecBeats struct {
	ID   string `json:"id,omitempty"`
	Port int32  `json:"port"`
}

type M2LogstashInputSpecHTTP struct {
	ID     string `json:"id,omitempty"`
	Port   int32  `json:"port"`
	Secret string `json:"secret,omitempty"`
}

type M2LogstashInputSpecTCP struct {
	ID   string `json:"id,omitempty"`
	Port int32  `json:"port"`
}

type M2LogstashInputSpecUDP struct {
	ID        string `json:"id,omitempty"`
	Port      int32  `json:"port"`
	QueueSize int    `json:"queue_size,omitempty"`
}

type M2LogstashInputSpecKafka struct {
	ID               string   `json:"id,omitempty"`
	Hosts            []string `json:"hosts,omitempty"`
	DecorateEvents   string   `json:"decorate_events,omitempty"`
	AutoOffsetReset  string   `json:"auto_offset_reset,omitempty"`
	ConsumerThreads  int      `json:"consumer_threads,omitempty"`
	SecurityProtocol string   `json:"security_protocol,omitempty"`
	SaslMechanism    string   `json:"sasl_mechanism,omitempty"`
	Secret           string   `json:"secret,omitempty"`
	Topic            string   `json:"topic,omitempty"`
}

type M2LogstashInputSpecS3 struct {
	ID             string `json:"id,omitempty"`
	Bucket         string `json:"bucket,omitempty"`
	Endpoint       string `json:"endpoint,omitempty"`
	Region         string `json:"region,omitempty"`
	GzipPattern    string `json:"gzip_pattern,omitempty"`
	Delete         bool   `json:"delete,omitempty"`
	ExcludePattern string `json:"exclude_pattern,omitempty"`
	Secret         string `json:"secret"`
	Prefix         string `json:"prefix,omitempty"`
}

type M2LogstashInputSpecDLQ struct {
	ID            string `json:"id,omitempty"`
	Path          string `json:"path"`
	CommitOffsets bool   `json:"commit_offsets,omitempty"`
	PipelineID    string `json:"pipeline_id,omitempty"`
}

// M2LogstashInputSpec defines the desired state of M2LogstashPipelineSpec
type M2LogstashInputSpec struct {
	Kafka   M2LogstashInputSpecKafka `json:"kafka,omitempty"`
	Beats   M2LogstashInputSpecBeats `json:"beats,omitempty"`
	HTTP    M2LogstashInputSpecHTTP  `json:"http,omitempty"`
	TCP     M2LogstashInputSpecTCP   `json:"tcp,omitempty"`
	UDP     M2LogstashInputSpecUDP   `json:"udp,omitempty"`
	S3      M2LogstashInputSpecS3    `json:"s3,omitempty"`
	DLQ     M2LogstashInputSpecDLQ   `json:"dlq,omitempty"`
	Workers int                      `json:"workers,omitempty"`
}

type M2LogstashOutputSpecElasticsearch struct {
	Hosts                      []string `json:"hosts,omitempty"`
	SSL                        bool     `json:"ssl,omitempty"`
	SSLCertificateVerification bool     `json:"ssl_certificate_verification,omitempty"`
	Cacert                     string   `json:"cacert,omitempty"`
	Secret                     string   `json:"secret,omitempty"`
	ILM                        bool     `json:"ilm,omitempty"`
	Index                      string   `json:"index,omitempty"`
}

type M2LogstashOutputSpecGraphite struct {
	Host              string `json:"host"`
	Port              int    `json:"port,omitempty"`
	ReconnectInterval int    `json:"reconnect_interval,omitempty"`
	ResendOnFailure   bool   `json:"resend_on_failure,omitempty"`
	TimestampField    string `json:"timestamp_field,omitempty"`
}

type M2LogstashOutputSpecS3 struct {
	Bucket            string `json:"bucket,omitempty"`
	CannedACL         string `json:"canned_acl,omitempty"`
	Encoding          string `json:"encoding,omitempty"`
	Endpoint          string `json:"endpoint,omitempty"`
	Region            string `json:"region,omitempty"`
	RotationStrategy  string `json:"rotation_strategy,omitempty"`
	SizeFile          int    `json:"size_file,omitempty"`
	TimeFile          int    `json:"time_file,omitempty"`
	UploadWorkerCount int    `json:"upload_worker_count,omitempty"`
	Secret            string `json:"secret"`
}

type M2LogstashOutputSpecTCP struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

type M2LogstashOutputSpecUDP struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

// +patchMergeKey=type
// +patchStrategy=merge
// +listType=map
// +listMapKey=type
// M2LogstashStatus defines the observed state of M2Logstash

// M2LogstashOutputSpec defines the desired state of M2LogstashPipelineSpec
type M2LogstashOutputSpec struct {
	Elasticsearch M2LogstashOutputSpecElasticsearch `json:"elasticsearch,omitempty"`
	TCP           M2LogstashOutputSpecTCP           `json:"tcp,omitempty"`
	UDP           M2LogstashOutputSpecUDP           `json:"udp,omitempty"`
	Graphite      M2LogstashOutputSpecGraphite      `json:"graphite,omitempty"`
	S3            M2LogstashOutputSpecS3            `json:"s3,omitempty"`
}

type M2LogstashPipelineSpec struct {
	Input  M2LogstashInputSpec  `json:"input"`
	Output M2LogstashOutputSpec `json:"output"`
	Filter string               `json:"filter,omitempty"`
}

// +patchMergeKey=type
// +patchStrategy=merge
// +listType=map
// +listMapKey=type
// M2LogstashPipelineStatus defines the observed state of M2LogstashPipeline

// M2LogstashPipelineStatus defines the observed state of M2LogstashPipeline
type M2LogstashPipelineStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// M2LogstashPipeline is the Schema for the m2logstashpipelines API
type M2LogstashPipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   M2LogstashPipelineSpec   `json:"spec,omitempty"`
	Status M2LogstashPipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// M2LogstashPipelineList contains a list of M2LogstashPipeline
type M2LogstashPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []M2LogstashPipeline `json:"items"`
}
