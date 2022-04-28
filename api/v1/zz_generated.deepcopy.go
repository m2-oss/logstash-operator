//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2Logstash) DeepCopyInto(out *M2Logstash) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2Logstash.
func (in *M2Logstash) DeepCopy() *M2Logstash {
	if in == nil {
		return nil
	}
	out := new(M2Logstash)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *M2Logstash) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashInputSpec) DeepCopyInto(out *M2LogstashInputSpec) {
	*out = *in
	in.Kafka.DeepCopyInto(&out.Kafka)
	out.Beats = in.Beats
	out.HTTP = in.HTTP
	out.TCP = in.TCP
	out.UDP = in.UDP
	out.S3 = in.S3
	out.DLQ = in.DLQ
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashInputSpec.
func (in *M2LogstashInputSpec) DeepCopy() *M2LogstashInputSpec {
	if in == nil {
		return nil
	}
	out := new(M2LogstashInputSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashInputSpecBeats) DeepCopyInto(out *M2LogstashInputSpecBeats) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashInputSpecBeats.
func (in *M2LogstashInputSpecBeats) DeepCopy() *M2LogstashInputSpecBeats {
	if in == nil {
		return nil
	}
	out := new(M2LogstashInputSpecBeats)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashInputSpecDLQ) DeepCopyInto(out *M2LogstashInputSpecDLQ) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashInputSpecDLQ.
func (in *M2LogstashInputSpecDLQ) DeepCopy() *M2LogstashInputSpecDLQ {
	if in == nil {
		return nil
	}
	out := new(M2LogstashInputSpecDLQ)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashInputSpecHTTP) DeepCopyInto(out *M2LogstashInputSpecHTTP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashInputSpecHTTP.
func (in *M2LogstashInputSpecHTTP) DeepCopy() *M2LogstashInputSpecHTTP {
	if in == nil {
		return nil
	}
	out := new(M2LogstashInputSpecHTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashInputSpecKafka) DeepCopyInto(out *M2LogstashInputSpecKafka) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashInputSpecKafka.
func (in *M2LogstashInputSpecKafka) DeepCopy() *M2LogstashInputSpecKafka {
	if in == nil {
		return nil
	}
	out := new(M2LogstashInputSpecKafka)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashInputSpecS3) DeepCopyInto(out *M2LogstashInputSpecS3) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashInputSpecS3.
func (in *M2LogstashInputSpecS3) DeepCopy() *M2LogstashInputSpecS3 {
	if in == nil {
		return nil
	}
	out := new(M2LogstashInputSpecS3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashInputSpecTCP) DeepCopyInto(out *M2LogstashInputSpecTCP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashInputSpecTCP.
func (in *M2LogstashInputSpecTCP) DeepCopy() *M2LogstashInputSpecTCP {
	if in == nil {
		return nil
	}
	out := new(M2LogstashInputSpecTCP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashInputSpecUDP) DeepCopyInto(out *M2LogstashInputSpecUDP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashInputSpecUDP.
func (in *M2LogstashInputSpecUDP) DeepCopy() *M2LogstashInputSpecUDP {
	if in == nil {
		return nil
	}
	out := new(M2LogstashInputSpecUDP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashList) DeepCopyInto(out *M2LogstashList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]M2Logstash, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashList.
func (in *M2LogstashList) DeepCopy() *M2LogstashList {
	if in == nil {
		return nil
	}
	out := new(M2LogstashList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *M2LogstashList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashOutputSpec) DeepCopyInto(out *M2LogstashOutputSpec) {
	*out = *in
	in.Elasticsearch.DeepCopyInto(&out.Elasticsearch)
	out.TCP = in.TCP
	out.UDP = in.UDP
	out.Graphite = in.Graphite
	out.S3 = in.S3
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashOutputSpec.
func (in *M2LogstashOutputSpec) DeepCopy() *M2LogstashOutputSpec {
	if in == nil {
		return nil
	}
	out := new(M2LogstashOutputSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashOutputSpecElasticsearch) DeepCopyInto(out *M2LogstashOutputSpecElasticsearch) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashOutputSpecElasticsearch.
func (in *M2LogstashOutputSpecElasticsearch) DeepCopy() *M2LogstashOutputSpecElasticsearch {
	if in == nil {
		return nil
	}
	out := new(M2LogstashOutputSpecElasticsearch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashOutputSpecGraphite) DeepCopyInto(out *M2LogstashOutputSpecGraphite) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashOutputSpecGraphite.
func (in *M2LogstashOutputSpecGraphite) DeepCopy() *M2LogstashOutputSpecGraphite {
	if in == nil {
		return nil
	}
	out := new(M2LogstashOutputSpecGraphite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashOutputSpecS3) DeepCopyInto(out *M2LogstashOutputSpecS3) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashOutputSpecS3.
func (in *M2LogstashOutputSpecS3) DeepCopy() *M2LogstashOutputSpecS3 {
	if in == nil {
		return nil
	}
	out := new(M2LogstashOutputSpecS3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashOutputSpecTCP) DeepCopyInto(out *M2LogstashOutputSpecTCP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashOutputSpecTCP.
func (in *M2LogstashOutputSpecTCP) DeepCopy() *M2LogstashOutputSpecTCP {
	if in == nil {
		return nil
	}
	out := new(M2LogstashOutputSpecTCP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashOutputSpecUDP) DeepCopyInto(out *M2LogstashOutputSpecUDP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashOutputSpecUDP.
func (in *M2LogstashOutputSpecUDP) DeepCopy() *M2LogstashOutputSpecUDP {
	if in == nil {
		return nil
	}
	out := new(M2LogstashOutputSpecUDP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashPipeline) DeepCopyInto(out *M2LogstashPipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashPipeline.
func (in *M2LogstashPipeline) DeepCopy() *M2LogstashPipeline {
	if in == nil {
		return nil
	}
	out := new(M2LogstashPipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *M2LogstashPipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashPipelineList) DeepCopyInto(out *M2LogstashPipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]M2LogstashPipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashPipelineList.
func (in *M2LogstashPipelineList) DeepCopy() *M2LogstashPipelineList {
	if in == nil {
		return nil
	}
	out := new(M2LogstashPipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *M2LogstashPipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashPipelineSpec) DeepCopyInto(out *M2LogstashPipelineSpec) {
	*out = *in
	in.Input.DeepCopyInto(&out.Input)
	in.Output.DeepCopyInto(&out.Output)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashPipelineSpec.
func (in *M2LogstashPipelineSpec) DeepCopy() *M2LogstashPipelineSpec {
	if in == nil {
		return nil
	}
	out := new(M2LogstashPipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashPipelineStatus) DeepCopyInto(out *M2LogstashPipelineStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashPipelineStatus.
func (in *M2LogstashPipelineStatus) DeepCopy() *M2LogstashPipelineStatus {
	if in == nil {
		return nil
	}
	out := new(M2LogstashPipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashSpec) DeepCopyInto(out *M2LogstashSpec) {
	*out = *in
	out.Resources = in.Resources
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashSpec.
func (in *M2LogstashSpec) DeepCopy() *M2LogstashSpec {
	if in == nil {
		return nil
	}
	out := new(M2LogstashSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashSpecLimit) DeepCopyInto(out *M2LogstashSpecLimit) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashSpecLimit.
func (in *M2LogstashSpecLimit) DeepCopy() *M2LogstashSpecLimit {
	if in == nil {
		return nil
	}
	out := new(M2LogstashSpecLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashSpecRequests) DeepCopyInto(out *M2LogstashSpecRequests) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashSpecRequests.
func (in *M2LogstashSpecRequests) DeepCopy() *M2LogstashSpecRequests {
	if in == nil {
		return nil
	}
	out := new(M2LogstashSpecRequests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashSpecResources) DeepCopyInto(out *M2LogstashSpecResources) {
	*out = *in
	out.Limit = in.Limit
	out.Requests = in.Requests
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashSpecResources.
func (in *M2LogstashSpecResources) DeepCopy() *M2LogstashSpecResources {
	if in == nil {
		return nil
	}
	out := new(M2LogstashSpecResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M2LogstashStatus) DeepCopyInto(out *M2LogstashStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M2LogstashStatus.
func (in *M2LogstashStatus) DeepCopy() *M2LogstashStatus {
	if in == nil {
		return nil
	}
	out := new(M2LogstashStatus)
	in.DeepCopyInto(out)
	return out
}