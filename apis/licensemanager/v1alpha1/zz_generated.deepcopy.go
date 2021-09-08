// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicensemanagerAssociation) DeepCopyInto(out *LicensemanagerAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicensemanagerAssociation.
func (in *LicensemanagerAssociation) DeepCopy() *LicensemanagerAssociation {
	if in == nil {
		return nil
	}
	out := new(LicensemanagerAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicensemanagerAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicensemanagerAssociationList) DeepCopyInto(out *LicensemanagerAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LicensemanagerAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicensemanagerAssociationList.
func (in *LicensemanagerAssociationList) DeepCopy() *LicensemanagerAssociationList {
	if in == nil {
		return nil
	}
	out := new(LicensemanagerAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicensemanagerAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicensemanagerAssociationObservation) DeepCopyInto(out *LicensemanagerAssociationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicensemanagerAssociationObservation.
func (in *LicensemanagerAssociationObservation) DeepCopy() *LicensemanagerAssociationObservation {
	if in == nil {
		return nil
	}
	out := new(LicensemanagerAssociationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicensemanagerAssociationParameters) DeepCopyInto(out *LicensemanagerAssociationParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicensemanagerAssociationParameters.
func (in *LicensemanagerAssociationParameters) DeepCopy() *LicensemanagerAssociationParameters {
	if in == nil {
		return nil
	}
	out := new(LicensemanagerAssociationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicensemanagerAssociationSpec) DeepCopyInto(out *LicensemanagerAssociationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicensemanagerAssociationSpec.
func (in *LicensemanagerAssociationSpec) DeepCopy() *LicensemanagerAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(LicensemanagerAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicensemanagerAssociationStatus) DeepCopyInto(out *LicensemanagerAssociationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicensemanagerAssociationStatus.
func (in *LicensemanagerAssociationStatus) DeepCopy() *LicensemanagerAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(LicensemanagerAssociationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicensemanagerLicenseConfiguration) DeepCopyInto(out *LicensemanagerLicenseConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicensemanagerLicenseConfiguration.
func (in *LicensemanagerLicenseConfiguration) DeepCopy() *LicensemanagerLicenseConfiguration {
	if in == nil {
		return nil
	}
	out := new(LicensemanagerLicenseConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicensemanagerLicenseConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicensemanagerLicenseConfigurationList) DeepCopyInto(out *LicensemanagerLicenseConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LicensemanagerLicenseConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicensemanagerLicenseConfigurationList.
func (in *LicensemanagerLicenseConfigurationList) DeepCopy() *LicensemanagerLicenseConfigurationList {
	if in == nil {
		return nil
	}
	out := new(LicensemanagerLicenseConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicensemanagerLicenseConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicensemanagerLicenseConfigurationObservation) DeepCopyInto(out *LicensemanagerLicenseConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicensemanagerLicenseConfigurationObservation.
func (in *LicensemanagerLicenseConfigurationObservation) DeepCopy() *LicensemanagerLicenseConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(LicensemanagerLicenseConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicensemanagerLicenseConfigurationParameters) DeepCopyInto(out *LicensemanagerLicenseConfigurationParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.LicenseCount != nil {
		in, out := &in.LicenseCount, &out.LicenseCount
		*out = new(int64)
		**out = **in
	}
	if in.LicenseCountHardLimit != nil {
		in, out := &in.LicenseCountHardLimit, &out.LicenseCountHardLimit
		*out = new(bool)
		**out = **in
	}
	if in.LicenseRules != nil {
		in, out := &in.LicenseRules, &out.LicenseRules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicensemanagerLicenseConfigurationParameters.
func (in *LicensemanagerLicenseConfigurationParameters) DeepCopy() *LicensemanagerLicenseConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(LicensemanagerLicenseConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicensemanagerLicenseConfigurationSpec) DeepCopyInto(out *LicensemanagerLicenseConfigurationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicensemanagerLicenseConfigurationSpec.
func (in *LicensemanagerLicenseConfigurationSpec) DeepCopy() *LicensemanagerLicenseConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(LicensemanagerLicenseConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicensemanagerLicenseConfigurationStatus) DeepCopyInto(out *LicensemanagerLicenseConfigurationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicensemanagerLicenseConfigurationStatus.
func (in *LicensemanagerLicenseConfigurationStatus) DeepCopy() *LicensemanagerLicenseConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(LicensemanagerLicenseConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}
