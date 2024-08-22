//go:build !ignore_autogenerated

/*
Copyright 2021 Kong, Inc.

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

package v1beta1

import (
	"github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
	konnectv1alpha1 "github.com/kong/kubernetes-configuration/api/konnect/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressBackend) DeepCopyInto(out *IngressBackend) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressBackend.
func (in *IngressBackend) DeepCopy() *IngressBackend {
	if in == nil {
		return nil
	}
	out := new(IngressBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRule) DeepCopyInto(out *IngressRule) {
	*out = *in
	out.Backend = in.Backend
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRule.
func (in *IngressRule) DeepCopy() *IngressRule {
	if in == nil {
		return nil
	}
	out := new(IngressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressTLS) DeepCopyInto(out *IngressTLS) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressTLS.
func (in *IngressTLS) DeepCopy() *IngressTLS {
	if in == nil {
		return nil
	}
	out := new(IngressTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongConsumerGroup) DeepCopyInto(out *KongConsumerGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongConsumerGroup.
func (in *KongConsumerGroup) DeepCopy() *KongConsumerGroup {
	if in == nil {
		return nil
	}
	out := new(KongConsumerGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongConsumerGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongConsumerGroupList) DeepCopyInto(out *KongConsumerGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KongConsumerGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongConsumerGroupList.
func (in *KongConsumerGroupList) DeepCopy() *KongConsumerGroupList {
	if in == nil {
		return nil
	}
	out := new(KongConsumerGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongConsumerGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongConsumerGroupSpec) DeepCopyInto(out *KongConsumerGroupSpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ControlPlaneRef != nil {
		in, out := &in.ControlPlaneRef, &out.ControlPlaneRef
		*out = new(v1alpha1.ControlPlaneRef)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongConsumerGroupSpec.
func (in *KongConsumerGroupSpec) DeepCopy() *KongConsumerGroupSpec {
	if in == nil {
		return nil
	}
	out := new(KongConsumerGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongConsumerGroupStatus) DeepCopyInto(out *KongConsumerGroupStatus) {
	*out = *in
	if in.Konnect != nil {
		in, out := &in.Konnect, &out.Konnect
		*out = new(konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongConsumerGroupStatus.
func (in *KongConsumerGroupStatus) DeepCopy() *KongConsumerGroupStatus {
	if in == nil {
		return nil
	}
	out := new(KongConsumerGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongUpstreamActiveHealthcheck) DeepCopyInto(out *KongUpstreamActiveHealthcheck) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Concurrency != nil {
		in, out := &in.Concurrency, &out.Concurrency
		*out = new(int)
		**out = **in
	}
	if in.Healthy != nil {
		in, out := &in.Healthy, &out.Healthy
		*out = new(KongUpstreamHealthcheckHealthy)
		(*in).DeepCopyInto(*out)
	}
	if in.Unhealthy != nil {
		in, out := &in.Unhealthy, &out.Unhealthy
		*out = new(KongUpstreamHealthcheckUnhealthy)
		(*in).DeepCopyInto(*out)
	}
	if in.HTTPPath != nil {
		in, out := &in.HTTPPath, &out.HTTPPath
		*out = new(string)
		**out = **in
	}
	if in.HTTPSSNI != nil {
		in, out := &in.HTTPSSNI, &out.HTTPSSNI
		*out = new(string)
		**out = **in
	}
	if in.HTTPSVerifyCertificate != nil {
		in, out := &in.HTTPSVerifyCertificate, &out.HTTPSVerifyCertificate
		*out = new(bool)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongUpstreamActiveHealthcheck.
func (in *KongUpstreamActiveHealthcheck) DeepCopy() *KongUpstreamActiveHealthcheck {
	if in == nil {
		return nil
	}
	out := new(KongUpstreamActiveHealthcheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongUpstreamHash) DeepCopyInto(out *KongUpstreamHash) {
	*out = *in
	if in.Input != nil {
		in, out := &in.Input, &out.Input
		*out = new(HashInput)
		**out = **in
	}
	if in.Header != nil {
		in, out := &in.Header, &out.Header
		*out = new(string)
		**out = **in
	}
	if in.Cookie != nil {
		in, out := &in.Cookie, &out.Cookie
		*out = new(string)
		**out = **in
	}
	if in.CookiePath != nil {
		in, out := &in.CookiePath, &out.CookiePath
		*out = new(string)
		**out = **in
	}
	if in.QueryArg != nil {
		in, out := &in.QueryArg, &out.QueryArg
		*out = new(string)
		**out = **in
	}
	if in.URICapture != nil {
		in, out := &in.URICapture, &out.URICapture
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongUpstreamHash.
func (in *KongUpstreamHash) DeepCopy() *KongUpstreamHash {
	if in == nil {
		return nil
	}
	out := new(KongUpstreamHash)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongUpstreamHealthcheck) DeepCopyInto(out *KongUpstreamHealthcheck) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = new(KongUpstreamActiveHealthcheck)
		(*in).DeepCopyInto(*out)
	}
	if in.Passive != nil {
		in, out := &in.Passive, &out.Passive
		*out = new(KongUpstreamPassiveHealthcheck)
		(*in).DeepCopyInto(*out)
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongUpstreamHealthcheck.
func (in *KongUpstreamHealthcheck) DeepCopy() *KongUpstreamHealthcheck {
	if in == nil {
		return nil
	}
	out := new(KongUpstreamHealthcheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongUpstreamHealthcheckHealthy) DeepCopyInto(out *KongUpstreamHealthcheckHealthy) {
	*out = *in
	if in.HTTPStatuses != nil {
		in, out := &in.HTTPStatuses, &out.HTTPStatuses
		*out = make([]HTTPStatus, len(*in))
		copy(*out, *in)
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(int)
		**out = **in
	}
	if in.Successes != nil {
		in, out := &in.Successes, &out.Successes
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongUpstreamHealthcheckHealthy.
func (in *KongUpstreamHealthcheckHealthy) DeepCopy() *KongUpstreamHealthcheckHealthy {
	if in == nil {
		return nil
	}
	out := new(KongUpstreamHealthcheckHealthy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongUpstreamHealthcheckUnhealthy) DeepCopyInto(out *KongUpstreamHealthcheckUnhealthy) {
	*out = *in
	if in.HTTPFailures != nil {
		in, out := &in.HTTPFailures, &out.HTTPFailures
		*out = new(int)
		**out = **in
	}
	if in.HTTPStatuses != nil {
		in, out := &in.HTTPStatuses, &out.HTTPStatuses
		*out = make([]HTTPStatus, len(*in))
		copy(*out, *in)
	}
	if in.TCPFailures != nil {
		in, out := &in.TCPFailures, &out.TCPFailures
		*out = new(int)
		**out = **in
	}
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(int)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongUpstreamHealthcheckUnhealthy.
func (in *KongUpstreamHealthcheckUnhealthy) DeepCopy() *KongUpstreamHealthcheckUnhealthy {
	if in == nil {
		return nil
	}
	out := new(KongUpstreamHealthcheckUnhealthy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongUpstreamPassiveHealthcheck) DeepCopyInto(out *KongUpstreamPassiveHealthcheck) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Healthy != nil {
		in, out := &in.Healthy, &out.Healthy
		*out = new(KongUpstreamHealthcheckHealthy)
		(*in).DeepCopyInto(*out)
	}
	if in.Unhealthy != nil {
		in, out := &in.Unhealthy, &out.Unhealthy
		*out = new(KongUpstreamHealthcheckUnhealthy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongUpstreamPassiveHealthcheck.
func (in *KongUpstreamPassiveHealthcheck) DeepCopy() *KongUpstreamPassiveHealthcheck {
	if in == nil {
		return nil
	}
	out := new(KongUpstreamPassiveHealthcheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongUpstreamPolicy) DeepCopyInto(out *KongUpstreamPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongUpstreamPolicy.
func (in *KongUpstreamPolicy) DeepCopy() *KongUpstreamPolicy {
	if in == nil {
		return nil
	}
	out := new(KongUpstreamPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongUpstreamPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongUpstreamPolicyList) DeepCopyInto(out *KongUpstreamPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KongUpstreamPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongUpstreamPolicyList.
func (in *KongUpstreamPolicyList) DeepCopy() *KongUpstreamPolicyList {
	if in == nil {
		return nil
	}
	out := new(KongUpstreamPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KongUpstreamPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KongUpstreamPolicySpec) DeepCopyInto(out *KongUpstreamPolicySpec) {
	*out = *in
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(string)
		**out = **in
	}
	if in.Slots != nil {
		in, out := &in.Slots, &out.Slots
		*out = new(int)
		**out = **in
	}
	if in.HashOn != nil {
		in, out := &in.HashOn, &out.HashOn
		*out = new(KongUpstreamHash)
		(*in).DeepCopyInto(*out)
	}
	if in.HashOnFallback != nil {
		in, out := &in.HashOnFallback, &out.HashOnFallback
		*out = new(KongUpstreamHash)
		(*in).DeepCopyInto(*out)
	}
	if in.Healthchecks != nil {
		in, out := &in.Healthchecks, &out.Healthchecks
		*out = new(KongUpstreamHealthcheck)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KongUpstreamPolicySpec.
func (in *KongUpstreamPolicySpec) DeepCopy() *KongUpstreamPolicySpec {
	if in == nil {
		return nil
	}
	out := new(KongUpstreamPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPIngress) DeepCopyInto(out *TCPIngress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPIngress.
func (in *TCPIngress) DeepCopy() *TCPIngress {
	if in == nil {
		return nil
	}
	out := new(TCPIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TCPIngress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPIngressList) DeepCopyInto(out *TCPIngressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TCPIngress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPIngressList.
func (in *TCPIngressList) DeepCopy() *TCPIngressList {
	if in == nil {
		return nil
	}
	out := new(TCPIngressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TCPIngressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPIngressSpec) DeepCopyInto(out *TCPIngressSpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]IngressRule, len(*in))
		copy(*out, *in)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = make([]IngressTLS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPIngressSpec.
func (in *TCPIngressSpec) DeepCopy() *TCPIngressSpec {
	if in == nil {
		return nil
	}
	out := new(TCPIngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPIngressStatus) DeepCopyInto(out *TCPIngressStatus) {
	*out = *in
	in.LoadBalancer.DeepCopyInto(&out.LoadBalancer)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPIngressStatus.
func (in *TCPIngressStatus) DeepCopy() *TCPIngressStatus {
	if in == nil {
		return nil
	}
	out := new(TCPIngressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UDPIngress) DeepCopyInto(out *UDPIngress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UDPIngress.
func (in *UDPIngress) DeepCopy() *UDPIngress {
	if in == nil {
		return nil
	}
	out := new(UDPIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UDPIngress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UDPIngressList) DeepCopyInto(out *UDPIngressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UDPIngress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UDPIngressList.
func (in *UDPIngressList) DeepCopy() *UDPIngressList {
	if in == nil {
		return nil
	}
	out := new(UDPIngressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UDPIngressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UDPIngressRule) DeepCopyInto(out *UDPIngressRule) {
	*out = *in
	out.Backend = in.Backend
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UDPIngressRule.
func (in *UDPIngressRule) DeepCopy() *UDPIngressRule {
	if in == nil {
		return nil
	}
	out := new(UDPIngressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UDPIngressSpec) DeepCopyInto(out *UDPIngressSpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]UDPIngressRule, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UDPIngressSpec.
func (in *UDPIngressSpec) DeepCopy() *UDPIngressSpec {
	if in == nil {
		return nil
	}
	out := new(UDPIngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UDPIngressStatus) DeepCopyInto(out *UDPIngressStatus) {
	*out = *in
	in.LoadBalancer.DeepCopyInto(&out.LoadBalancer)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UDPIngressStatus.
func (in *UDPIngressStatus) DeepCopy() *UDPIngressStatus {
	if in == nil {
		return nil
	}
	out := new(UDPIngressStatus)
	in.DeepCopyInto(out)
	return out
}
