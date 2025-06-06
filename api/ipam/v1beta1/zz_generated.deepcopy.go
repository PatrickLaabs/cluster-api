//go:build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	corev1beta1 "sigs.k8s.io/cluster-api/api/core/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddress) DeepCopyInto(out *IPAddress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddress.
func (in *IPAddress) DeepCopy() *IPAddress {
	if in == nil {
		return nil
	}
	out := new(IPAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPAddress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddressClaim) DeepCopyInto(out *IPAddressClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddressClaim.
func (in *IPAddressClaim) DeepCopy() *IPAddressClaim {
	if in == nil {
		return nil
	}
	out := new(IPAddressClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPAddressClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddressClaimList) DeepCopyInto(out *IPAddressClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPAddressClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddressClaimList.
func (in *IPAddressClaimList) DeepCopy() *IPAddressClaimList {
	if in == nil {
		return nil
	}
	out := new(IPAddressClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPAddressClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddressClaimSpec) DeepCopyInto(out *IPAddressClaimSpec) {
	*out = *in
	in.PoolRef.DeepCopyInto(&out.PoolRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddressClaimSpec.
func (in *IPAddressClaimSpec) DeepCopy() *IPAddressClaimSpec {
	if in == nil {
		return nil
	}
	out := new(IPAddressClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddressClaimStatus) DeepCopyInto(out *IPAddressClaimStatus) {
	*out = *in
	out.AddressRef = in.AddressRef
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(corev1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.V1Beta2 != nil {
		in, out := &in.V1Beta2, &out.V1Beta2
		*out = new(IPAddressClaimV1Beta2Status)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddressClaimStatus.
func (in *IPAddressClaimStatus) DeepCopy() *IPAddressClaimStatus {
	if in == nil {
		return nil
	}
	out := new(IPAddressClaimStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddressClaimV1Beta2Status) DeepCopyInto(out *IPAddressClaimV1Beta2Status) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddressClaimV1Beta2Status.
func (in *IPAddressClaimV1Beta2Status) DeepCopy() *IPAddressClaimV1Beta2Status {
	if in == nil {
		return nil
	}
	out := new(IPAddressClaimV1Beta2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddressList) DeepCopyInto(out *IPAddressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPAddress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddressList.
func (in *IPAddressList) DeepCopy() *IPAddressList {
	if in == nil {
		return nil
	}
	out := new(IPAddressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPAddressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAddressSpec) DeepCopyInto(out *IPAddressSpec) {
	*out = *in
	out.ClaimRef = in.ClaimRef
	in.PoolRef.DeepCopyInto(&out.PoolRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAddressSpec.
func (in *IPAddressSpec) DeepCopy() *IPAddressSpec {
	if in == nil {
		return nil
	}
	out := new(IPAddressSpec)
	in.DeepCopyInto(out)
	return out
}
