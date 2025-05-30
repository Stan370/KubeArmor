//go:build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 Authors of KubeArmor

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	security_kubearmor_comv1 "github.com/kubearmor/KubeArmor/pkg/KubeArmorController/api/security.kubearmor.com/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Adapters) DeepCopyInto(out *Adapters) {
	*out = *in
	out.ElasticSearch = in.ElasticSearch
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Adapters.
func (in *Adapters) DeepCopy() *Adapters {
	if in == nil {
		return nil
	}
	out := new(Adapters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticSearchAdapter) DeepCopyInto(out *ElasticSearchAdapter) {
	*out = *in
	out.Auth = in.Auth
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticSearchAdapter.
func (in *ElasticSearchAdapter) DeepCopy() *ElasticSearchAdapter {
	if in == nil {
		return nil
	}
	out := new(ElasticSearchAdapter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticSearchAuth) DeepCopyInto(out *ElasticSearchAuth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticSearchAuth.
func (in *ElasticSearchAuth) DeepCopy() *ElasticSearchAuth {
	if in == nil {
		return nil
	}
	out := new(ElasticSearchAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorConfig) DeepCopyInto(out *KubeArmorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorConfig.
func (in *KubeArmorConfig) DeepCopy() *KubeArmorConfig {
	if in == nil {
		return nil
	}
	out := new(KubeArmorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeArmorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorConfigList) DeepCopyInto(out *KubeArmorConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeArmorConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorConfigList.
func (in *KubeArmorConfigList) DeepCopy() *KubeArmorConfigList {
	if in == nil {
		return nil
	}
	out := new(KubeArmorConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeArmorConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorConfigSpec) DeepCopyInto(out *KubeArmorConfigSpec) {
	*out = *in
	in.RecommendedPolicies.DeepCopyInto(&out.RecommendedPolicies)
	if in.GloabalImagePullSecrets != nil {
		in, out := &in.GloabalImagePullSecrets, &out.GloabalImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.GlobalTolerations != nil {
		in, out := &in.GlobalTolerations, &out.GlobalTolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.KubeArmorImage.DeepCopyInto(&out.KubeArmorImage)
	in.KubeArmorInitImage.DeepCopyInto(&out.KubeArmorInitImage)
	in.KubeArmorRelayImage.DeepCopyInto(&out.KubeArmorRelayImage)
	in.KubeArmorControllerImage.DeepCopyInto(&out.KubeArmorControllerImage)
	in.KubeRbacProxyImage.DeepCopyInto(&out.KubeRbacProxyImage)
	in.Tls.DeepCopyInto(&out.Tls)
	out.Adapters = in.Adapters
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorConfigSpec.
func (in *KubeArmorConfigSpec) DeepCopy() *KubeArmorConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KubeArmorConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeArmorConfigStatus) DeepCopyInto(out *KubeArmorConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeArmorConfigStatus.
func (in *KubeArmorConfigStatus) DeepCopy() *KubeArmorConfigStatus {
	if in == nil {
		return nil
	}
	out := new(KubeArmorConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecommendedPolicies) DeepCopyInto(out *RecommendedPolicies) {
	*out = *in
	if in.MatchExpressions != nil {
		in, out := &in.MatchExpressions, &out.MatchExpressions
		*out = make([]security_kubearmor_comv1.ClusterMatchExpressionsType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExcludePolicy != nil {
		in, out := &in.ExcludePolicy, &out.ExcludePolicy
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecommendedPolicies.
func (in *RecommendedPolicies) DeepCopy() *RecommendedPolicies {
	if in == nil {
		return nil
	}
	out := new(RecommendedPolicies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tls) DeepCopyInto(out *Tls) {
	*out = *in
	if in.RelayExtraDnsNames != nil {
		in, out := &in.RelayExtraDnsNames, &out.RelayExtraDnsNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RelayExtraIpAddresses != nil {
		in, out := &in.RelayExtraIpAddresses, &out.RelayExtraIpAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tls.
func (in *Tls) DeepCopy() *Tls {
	if in == nil {
		return nil
	}
	out := new(Tls)
	in.DeepCopyInto(out)
	return out
}
