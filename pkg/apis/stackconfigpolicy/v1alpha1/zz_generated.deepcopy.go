//go:build !ignore_autogenerated

// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/elastic/cloud-on-k8s/v2/pkg/apis/common/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchConfigPolicySpec) DeepCopyInto(out *ElasticsearchConfigPolicySpec) {
	*out = *in
	if in.ClusterSettings != nil {
		in, out := &in.ClusterSettings, &out.ClusterSettings
		*out = (*in).DeepCopy()
	}
	if in.SnapshotRepositories != nil {
		in, out := &in.SnapshotRepositories, &out.SnapshotRepositories
		*out = (*in).DeepCopy()
	}
	if in.SnapshotLifecyclePolicies != nil {
		in, out := &in.SnapshotLifecyclePolicies, &out.SnapshotLifecyclePolicies
		*out = (*in).DeepCopy()
	}
	if in.SecurityRoleMappings != nil {
		in, out := &in.SecurityRoleMappings, &out.SecurityRoleMappings
		*out = (*in).DeepCopy()
	}
	if in.IndexLifecyclePolicies != nil {
		in, out := &in.IndexLifecyclePolicies, &out.IndexLifecyclePolicies
		*out = (*in).DeepCopy()
	}
	if in.IngestPipelines != nil {
		in, out := &in.IngestPipelines, &out.IngestPipelines
		*out = (*in).DeepCopy()
	}
	in.IndexTemplates.DeepCopyInto(&out.IndexTemplates)
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = (*in).DeepCopy()
	}
	if in.SecretMounts != nil {
		in, out := &in.SecretMounts, &out.SecretMounts
		*out = make([]SecretMount, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchConfigPolicySpec.
func (in *ElasticsearchConfigPolicySpec) DeepCopy() *ElasticsearchConfigPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchConfigPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchPolicyStatus) DeepCopyInto(out *ElasticsearchPolicyStatus) {
	*out = *in
	out.Error = in.Error
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchPolicyStatus.
func (in *ElasticsearchPolicyStatus) DeepCopy() *ElasticsearchPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexTemplates) DeepCopyInto(out *IndexTemplates) {
	*out = *in
	if in.ComponentTemplates != nil {
		in, out := &in.ComponentTemplates, &out.ComponentTemplates
		*out = (*in).DeepCopy()
	}
	if in.ComposableIndexTemplates != nil {
		in, out := &in.ComposableIndexTemplates, &out.ComposableIndexTemplates
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexTemplates.
func (in *IndexTemplates) DeepCopy() *IndexTemplates {
	if in == nil {
		return nil
	}
	out := new(IndexTemplates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KibanaConfigPolicySpec) DeepCopyInto(out *KibanaConfigPolicySpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KibanaConfigPolicySpec.
func (in *KibanaConfigPolicySpec) DeepCopy() *KibanaConfigPolicySpec {
	if in == nil {
		return nil
	}
	out := new(KibanaConfigPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KibanaPolicyStatus) DeepCopyInto(out *KibanaPolicyStatus) {
	*out = *in
	out.Error = in.Error
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KibanaPolicyStatus.
func (in *KibanaPolicyStatus) DeepCopy() *KibanaPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(KibanaPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyStatusError) DeepCopyInto(out *PolicyStatusError) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyStatusError.
func (in *PolicyStatusError) DeepCopy() *PolicyStatusError {
	if in == nil {
		return nil
	}
	out := new(PolicyStatusError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePolicyStatus) DeepCopyInto(out *ResourcePolicyStatus) {
	*out = *in
	out.ElasticsearchStatus = in.ElasticsearchStatus
	out.KibanaStatus = in.KibanaStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePolicyStatus.
func (in *ResourcePolicyStatus) DeepCopy() *ResourcePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ResourcePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretMount) DeepCopyInto(out *SecretMount) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretMount.
func (in *SecretMount) DeepCopy() *SecretMount {
	if in == nil {
		return nil
	}
	out := new(SecretMount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackConfigPolicy) DeepCopyInto(out *StackConfigPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackConfigPolicy.
func (in *StackConfigPolicy) DeepCopy() *StackConfigPolicy {
	if in == nil {
		return nil
	}
	out := new(StackConfigPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StackConfigPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackConfigPolicyList) DeepCopyInto(out *StackConfigPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StackConfigPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackConfigPolicyList.
func (in *StackConfigPolicyList) DeepCopy() *StackConfigPolicyList {
	if in == nil {
		return nil
	}
	out := new(StackConfigPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StackConfigPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackConfigPolicySpec) DeepCopyInto(out *StackConfigPolicySpec) {
	*out = *in
	in.ResourceSelector.DeepCopyInto(&out.ResourceSelector)
	if in.SecureSettings != nil {
		in, out := &in.SecureSettings, &out.SecureSettings
		*out = make([]v1.SecretSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Elasticsearch.DeepCopyInto(&out.Elasticsearch)
	in.Kibana.DeepCopyInto(&out.Kibana)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackConfigPolicySpec.
func (in *StackConfigPolicySpec) DeepCopy() *StackConfigPolicySpec {
	if in == nil {
		return nil
	}
	out := new(StackConfigPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackConfigPolicyStatus) DeepCopyInto(out *StackConfigPolicyStatus) {
	*out = *in
	if in.ResourcesStatuses != nil {
		in, out := &in.ResourcesStatuses, &out.ResourcesStatuses
		*out = make(map[string]ResourcePolicyStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackConfigPolicyStatus.
func (in *StackConfigPolicyStatus) DeepCopy() *StackConfigPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(StackConfigPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
