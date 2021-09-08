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
func (in *AdditionalAuthenticationProviderObservation) DeepCopyInto(out *AdditionalAuthenticationProviderObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalAuthenticationProviderObservation.
func (in *AdditionalAuthenticationProviderObservation) DeepCopy() *AdditionalAuthenticationProviderObservation {
	if in == nil {
		return nil
	}
	out := new(AdditionalAuthenticationProviderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalAuthenticationProviderParameters) DeepCopyInto(out *AdditionalAuthenticationProviderParameters) {
	*out = *in
	if in.OpenidConnectConfig != nil {
		in, out := &in.OpenidConnectConfig, &out.OpenidConnectConfig
		*out = make([]OpenidConnectConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UserPoolConfig != nil {
		in, out := &in.UserPoolConfig, &out.UserPoolConfig
		*out = make([]UserPoolConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalAuthenticationProviderParameters.
func (in *AdditionalAuthenticationProviderParameters) DeepCopy() *AdditionalAuthenticationProviderParameters {
	if in == nil {
		return nil
	}
	out := new(AdditionalAuthenticationProviderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncApiKey) DeepCopyInto(out *AppsyncApiKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncApiKey.
func (in *AppsyncApiKey) DeepCopy() *AppsyncApiKey {
	if in == nil {
		return nil
	}
	out := new(AppsyncApiKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppsyncApiKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncApiKeyList) DeepCopyInto(out *AppsyncApiKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppsyncApiKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncApiKeyList.
func (in *AppsyncApiKeyList) DeepCopy() *AppsyncApiKeyList {
	if in == nil {
		return nil
	}
	out := new(AppsyncApiKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppsyncApiKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncApiKeyObservation) DeepCopyInto(out *AppsyncApiKeyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncApiKeyObservation.
func (in *AppsyncApiKeyObservation) DeepCopy() *AppsyncApiKeyObservation {
	if in == nil {
		return nil
	}
	out := new(AppsyncApiKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncApiKeyParameters) DeepCopyInto(out *AppsyncApiKeyParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Expires != nil {
		in, out := &in.Expires, &out.Expires
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncApiKeyParameters.
func (in *AppsyncApiKeyParameters) DeepCopy() *AppsyncApiKeyParameters {
	if in == nil {
		return nil
	}
	out := new(AppsyncApiKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncApiKeySpec) DeepCopyInto(out *AppsyncApiKeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncApiKeySpec.
func (in *AppsyncApiKeySpec) DeepCopy() *AppsyncApiKeySpec {
	if in == nil {
		return nil
	}
	out := new(AppsyncApiKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncApiKeyStatus) DeepCopyInto(out *AppsyncApiKeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncApiKeyStatus.
func (in *AppsyncApiKeyStatus) DeepCopy() *AppsyncApiKeyStatus {
	if in == nil {
		return nil
	}
	out := new(AppsyncApiKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncDatasource) DeepCopyInto(out *AppsyncDatasource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncDatasource.
func (in *AppsyncDatasource) DeepCopy() *AppsyncDatasource {
	if in == nil {
		return nil
	}
	out := new(AppsyncDatasource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppsyncDatasource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncDatasourceList) DeepCopyInto(out *AppsyncDatasourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppsyncDatasource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncDatasourceList.
func (in *AppsyncDatasourceList) DeepCopy() *AppsyncDatasourceList {
	if in == nil {
		return nil
	}
	out := new(AppsyncDatasourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppsyncDatasourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncDatasourceObservation) DeepCopyInto(out *AppsyncDatasourceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncDatasourceObservation.
func (in *AppsyncDatasourceObservation) DeepCopy() *AppsyncDatasourceObservation {
	if in == nil {
		return nil
	}
	out := new(AppsyncDatasourceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncDatasourceParameters) DeepCopyInto(out *AppsyncDatasourceParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DynamodbConfig != nil {
		in, out := &in.DynamodbConfig, &out.DynamodbConfig
		*out = make([]DynamodbConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ElasticsearchConfig != nil {
		in, out := &in.ElasticsearchConfig, &out.ElasticsearchConfig
		*out = make([]ElasticsearchConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HttpConfig != nil {
		in, out := &in.HttpConfig, &out.HttpConfig
		*out = make([]HttpConfigParameters, len(*in))
		copy(*out, *in)
	}
	if in.LambdaConfig != nil {
		in, out := &in.LambdaConfig, &out.LambdaConfig
		*out = make([]LambdaConfigParameters, len(*in))
		copy(*out, *in)
	}
	if in.ServiceRoleArn != nil {
		in, out := &in.ServiceRoleArn, &out.ServiceRoleArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncDatasourceParameters.
func (in *AppsyncDatasourceParameters) DeepCopy() *AppsyncDatasourceParameters {
	if in == nil {
		return nil
	}
	out := new(AppsyncDatasourceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncDatasourceSpec) DeepCopyInto(out *AppsyncDatasourceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncDatasourceSpec.
func (in *AppsyncDatasourceSpec) DeepCopy() *AppsyncDatasourceSpec {
	if in == nil {
		return nil
	}
	out := new(AppsyncDatasourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncDatasourceStatus) DeepCopyInto(out *AppsyncDatasourceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncDatasourceStatus.
func (in *AppsyncDatasourceStatus) DeepCopy() *AppsyncDatasourceStatus {
	if in == nil {
		return nil
	}
	out := new(AppsyncDatasourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncFunction) DeepCopyInto(out *AppsyncFunction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncFunction.
func (in *AppsyncFunction) DeepCopy() *AppsyncFunction {
	if in == nil {
		return nil
	}
	out := new(AppsyncFunction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppsyncFunction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncFunctionList) DeepCopyInto(out *AppsyncFunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppsyncFunction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncFunctionList.
func (in *AppsyncFunctionList) DeepCopy() *AppsyncFunctionList {
	if in == nil {
		return nil
	}
	out := new(AppsyncFunctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppsyncFunctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncFunctionObservation) DeepCopyInto(out *AppsyncFunctionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncFunctionObservation.
func (in *AppsyncFunctionObservation) DeepCopy() *AppsyncFunctionObservation {
	if in == nil {
		return nil
	}
	out := new(AppsyncFunctionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncFunctionParameters) DeepCopyInto(out *AppsyncFunctionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FunctionVersion != nil {
		in, out := &in.FunctionVersion, &out.FunctionVersion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncFunctionParameters.
func (in *AppsyncFunctionParameters) DeepCopy() *AppsyncFunctionParameters {
	if in == nil {
		return nil
	}
	out := new(AppsyncFunctionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncFunctionSpec) DeepCopyInto(out *AppsyncFunctionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncFunctionSpec.
func (in *AppsyncFunctionSpec) DeepCopy() *AppsyncFunctionSpec {
	if in == nil {
		return nil
	}
	out := new(AppsyncFunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncFunctionStatus) DeepCopyInto(out *AppsyncFunctionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncFunctionStatus.
func (in *AppsyncFunctionStatus) DeepCopy() *AppsyncFunctionStatus {
	if in == nil {
		return nil
	}
	out := new(AppsyncFunctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncGraphqlApi) DeepCopyInto(out *AppsyncGraphqlApi) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncGraphqlApi.
func (in *AppsyncGraphqlApi) DeepCopy() *AppsyncGraphqlApi {
	if in == nil {
		return nil
	}
	out := new(AppsyncGraphqlApi)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppsyncGraphqlApi) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncGraphqlApiList) DeepCopyInto(out *AppsyncGraphqlApiList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppsyncGraphqlApi, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncGraphqlApiList.
func (in *AppsyncGraphqlApiList) DeepCopy() *AppsyncGraphqlApiList {
	if in == nil {
		return nil
	}
	out := new(AppsyncGraphqlApiList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppsyncGraphqlApiList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncGraphqlApiObservation) DeepCopyInto(out *AppsyncGraphqlApiObservation) {
	*out = *in
	if in.Uris != nil {
		in, out := &in.Uris, &out.Uris
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncGraphqlApiObservation.
func (in *AppsyncGraphqlApiObservation) DeepCopy() *AppsyncGraphqlApiObservation {
	if in == nil {
		return nil
	}
	out := new(AppsyncGraphqlApiObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncGraphqlApiOpenidConnectConfigObservation) DeepCopyInto(out *AppsyncGraphqlApiOpenidConnectConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncGraphqlApiOpenidConnectConfigObservation.
func (in *AppsyncGraphqlApiOpenidConnectConfigObservation) DeepCopy() *AppsyncGraphqlApiOpenidConnectConfigObservation {
	if in == nil {
		return nil
	}
	out := new(AppsyncGraphqlApiOpenidConnectConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncGraphqlApiOpenidConnectConfigParameters) DeepCopyInto(out *AppsyncGraphqlApiOpenidConnectConfigParameters) {
	*out = *in
	if in.AuthTtl != nil {
		in, out := &in.AuthTtl, &out.AuthTtl
		*out = new(int64)
		**out = **in
	}
	if in.ClientId != nil {
		in, out := &in.ClientId, &out.ClientId
		*out = new(string)
		**out = **in
	}
	if in.IatTtl != nil {
		in, out := &in.IatTtl, &out.IatTtl
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncGraphqlApiOpenidConnectConfigParameters.
func (in *AppsyncGraphqlApiOpenidConnectConfigParameters) DeepCopy() *AppsyncGraphqlApiOpenidConnectConfigParameters {
	if in == nil {
		return nil
	}
	out := new(AppsyncGraphqlApiOpenidConnectConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncGraphqlApiParameters) DeepCopyInto(out *AppsyncGraphqlApiParameters) {
	*out = *in
	if in.AdditionalAuthenticationProvider != nil {
		in, out := &in.AdditionalAuthenticationProvider, &out.AdditionalAuthenticationProvider
		*out = make([]AdditionalAuthenticationProviderParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LogConfig != nil {
		in, out := &in.LogConfig, &out.LogConfig
		*out = make([]LogConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OpenidConnectConfig != nil {
		in, out := &in.OpenidConnectConfig, &out.OpenidConnectConfig
		*out = make([]AppsyncGraphqlApiOpenidConnectConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = new(string)
		**out = **in
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
	if in.UserPoolConfig != nil {
		in, out := &in.UserPoolConfig, &out.UserPoolConfig
		*out = make([]AppsyncGraphqlApiUserPoolConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.XrayEnabled != nil {
		in, out := &in.XrayEnabled, &out.XrayEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncGraphqlApiParameters.
func (in *AppsyncGraphqlApiParameters) DeepCopy() *AppsyncGraphqlApiParameters {
	if in == nil {
		return nil
	}
	out := new(AppsyncGraphqlApiParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncGraphqlApiSpec) DeepCopyInto(out *AppsyncGraphqlApiSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncGraphqlApiSpec.
func (in *AppsyncGraphqlApiSpec) DeepCopy() *AppsyncGraphqlApiSpec {
	if in == nil {
		return nil
	}
	out := new(AppsyncGraphqlApiSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncGraphqlApiStatus) DeepCopyInto(out *AppsyncGraphqlApiStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncGraphqlApiStatus.
func (in *AppsyncGraphqlApiStatus) DeepCopy() *AppsyncGraphqlApiStatus {
	if in == nil {
		return nil
	}
	out := new(AppsyncGraphqlApiStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncGraphqlApiUserPoolConfigObservation) DeepCopyInto(out *AppsyncGraphqlApiUserPoolConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncGraphqlApiUserPoolConfigObservation.
func (in *AppsyncGraphqlApiUserPoolConfigObservation) DeepCopy() *AppsyncGraphqlApiUserPoolConfigObservation {
	if in == nil {
		return nil
	}
	out := new(AppsyncGraphqlApiUserPoolConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncGraphqlApiUserPoolConfigParameters) DeepCopyInto(out *AppsyncGraphqlApiUserPoolConfigParameters) {
	*out = *in
	if in.AppIdClientRegex != nil {
		in, out := &in.AppIdClientRegex, &out.AppIdClientRegex
		*out = new(string)
		**out = **in
	}
	if in.AwsRegion != nil {
		in, out := &in.AwsRegion, &out.AwsRegion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncGraphqlApiUserPoolConfigParameters.
func (in *AppsyncGraphqlApiUserPoolConfigParameters) DeepCopy() *AppsyncGraphqlApiUserPoolConfigParameters {
	if in == nil {
		return nil
	}
	out := new(AppsyncGraphqlApiUserPoolConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncResolver) DeepCopyInto(out *AppsyncResolver) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncResolver.
func (in *AppsyncResolver) DeepCopy() *AppsyncResolver {
	if in == nil {
		return nil
	}
	out := new(AppsyncResolver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppsyncResolver) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncResolverList) DeepCopyInto(out *AppsyncResolverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppsyncResolver, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncResolverList.
func (in *AppsyncResolverList) DeepCopy() *AppsyncResolverList {
	if in == nil {
		return nil
	}
	out := new(AppsyncResolverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppsyncResolverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncResolverObservation) DeepCopyInto(out *AppsyncResolverObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncResolverObservation.
func (in *AppsyncResolverObservation) DeepCopy() *AppsyncResolverObservation {
	if in == nil {
		return nil
	}
	out := new(AppsyncResolverObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncResolverParameters) DeepCopyInto(out *AppsyncResolverParameters) {
	*out = *in
	if in.CachingConfig != nil {
		in, out := &in.CachingConfig, &out.CachingConfig
		*out = make([]CachingConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DataSource != nil {
		in, out := &in.DataSource, &out.DataSource
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.PipelineConfig != nil {
		in, out := &in.PipelineConfig, &out.PipelineConfig
		*out = make([]PipelineConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RequestTemplate != nil {
		in, out := &in.RequestTemplate, &out.RequestTemplate
		*out = new(string)
		**out = **in
	}
	if in.ResponseTemplate != nil {
		in, out := &in.ResponseTemplate, &out.ResponseTemplate
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncResolverParameters.
func (in *AppsyncResolverParameters) DeepCopy() *AppsyncResolverParameters {
	if in == nil {
		return nil
	}
	out := new(AppsyncResolverParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncResolverSpec) DeepCopyInto(out *AppsyncResolverSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncResolverSpec.
func (in *AppsyncResolverSpec) DeepCopy() *AppsyncResolverSpec {
	if in == nil {
		return nil
	}
	out := new(AppsyncResolverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsyncResolverStatus) DeepCopyInto(out *AppsyncResolverStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsyncResolverStatus.
func (in *AppsyncResolverStatus) DeepCopy() *AppsyncResolverStatus {
	if in == nil {
		return nil
	}
	out := new(AppsyncResolverStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CachingConfigObservation) DeepCopyInto(out *CachingConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CachingConfigObservation.
func (in *CachingConfigObservation) DeepCopy() *CachingConfigObservation {
	if in == nil {
		return nil
	}
	out := new(CachingConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CachingConfigParameters) DeepCopyInto(out *CachingConfigParameters) {
	*out = *in
	if in.CachingKeys != nil {
		in, out := &in.CachingKeys, &out.CachingKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ttl != nil {
		in, out := &in.Ttl, &out.Ttl
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CachingConfigParameters.
func (in *CachingConfigParameters) DeepCopy() *CachingConfigParameters {
	if in == nil {
		return nil
	}
	out := new(CachingConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamodbConfigObservation) DeepCopyInto(out *DynamodbConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamodbConfigObservation.
func (in *DynamodbConfigObservation) DeepCopy() *DynamodbConfigObservation {
	if in == nil {
		return nil
	}
	out := new(DynamodbConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamodbConfigParameters) DeepCopyInto(out *DynamodbConfigParameters) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.UseCallerCredentials != nil {
		in, out := &in.UseCallerCredentials, &out.UseCallerCredentials
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamodbConfigParameters.
func (in *DynamodbConfigParameters) DeepCopy() *DynamodbConfigParameters {
	if in == nil {
		return nil
	}
	out := new(DynamodbConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchConfigObservation) DeepCopyInto(out *ElasticsearchConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchConfigObservation.
func (in *ElasticsearchConfigObservation) DeepCopy() *ElasticsearchConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchConfigParameters) DeepCopyInto(out *ElasticsearchConfigParameters) {
	*out = *in
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchConfigParameters.
func (in *ElasticsearchConfigParameters) DeepCopy() *ElasticsearchConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpConfigObservation) DeepCopyInto(out *HttpConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpConfigObservation.
func (in *HttpConfigObservation) DeepCopy() *HttpConfigObservation {
	if in == nil {
		return nil
	}
	out := new(HttpConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpConfigParameters) DeepCopyInto(out *HttpConfigParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpConfigParameters.
func (in *HttpConfigParameters) DeepCopy() *HttpConfigParameters {
	if in == nil {
		return nil
	}
	out := new(HttpConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaConfigObservation) DeepCopyInto(out *LambdaConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaConfigObservation.
func (in *LambdaConfigObservation) DeepCopy() *LambdaConfigObservation {
	if in == nil {
		return nil
	}
	out := new(LambdaConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LambdaConfigParameters) DeepCopyInto(out *LambdaConfigParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LambdaConfigParameters.
func (in *LambdaConfigParameters) DeepCopy() *LambdaConfigParameters {
	if in == nil {
		return nil
	}
	out := new(LambdaConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogConfigObservation) DeepCopyInto(out *LogConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogConfigObservation.
func (in *LogConfigObservation) DeepCopy() *LogConfigObservation {
	if in == nil {
		return nil
	}
	out := new(LogConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogConfigParameters) DeepCopyInto(out *LogConfigParameters) {
	*out = *in
	if in.ExcludeVerboseContent != nil {
		in, out := &in.ExcludeVerboseContent, &out.ExcludeVerboseContent
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogConfigParameters.
func (in *LogConfigParameters) DeepCopy() *LogConfigParameters {
	if in == nil {
		return nil
	}
	out := new(LogConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenidConnectConfigObservation) DeepCopyInto(out *OpenidConnectConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenidConnectConfigObservation.
func (in *OpenidConnectConfigObservation) DeepCopy() *OpenidConnectConfigObservation {
	if in == nil {
		return nil
	}
	out := new(OpenidConnectConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenidConnectConfigParameters) DeepCopyInto(out *OpenidConnectConfigParameters) {
	*out = *in
	if in.AuthTtl != nil {
		in, out := &in.AuthTtl, &out.AuthTtl
		*out = new(int64)
		**out = **in
	}
	if in.ClientId != nil {
		in, out := &in.ClientId, &out.ClientId
		*out = new(string)
		**out = **in
	}
	if in.IatTtl != nil {
		in, out := &in.IatTtl, &out.IatTtl
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenidConnectConfigParameters.
func (in *OpenidConnectConfigParameters) DeepCopy() *OpenidConnectConfigParameters {
	if in == nil {
		return nil
	}
	out := new(OpenidConnectConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineConfigObservation) DeepCopyInto(out *PipelineConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineConfigObservation.
func (in *PipelineConfigObservation) DeepCopy() *PipelineConfigObservation {
	if in == nil {
		return nil
	}
	out := new(PipelineConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineConfigParameters) DeepCopyInto(out *PipelineConfigParameters) {
	*out = *in
	if in.Functions != nil {
		in, out := &in.Functions, &out.Functions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineConfigParameters.
func (in *PipelineConfigParameters) DeepCopy() *PipelineConfigParameters {
	if in == nil {
		return nil
	}
	out := new(PipelineConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolConfigObservation) DeepCopyInto(out *UserPoolConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolConfigObservation.
func (in *UserPoolConfigObservation) DeepCopy() *UserPoolConfigObservation {
	if in == nil {
		return nil
	}
	out := new(UserPoolConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolConfigParameters) DeepCopyInto(out *UserPoolConfigParameters) {
	*out = *in
	if in.AppIdClientRegex != nil {
		in, out := &in.AppIdClientRegex, &out.AppIdClientRegex
		*out = new(string)
		**out = **in
	}
	if in.AwsRegion != nil {
		in, out := &in.AwsRegion, &out.AwsRegion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolConfigParameters.
func (in *UserPoolConfigParameters) DeepCopy() *UserPoolConfigParameters {
	if in == nil {
		return nil
	}
	out := new(UserPoolConfigParameters)
	in.DeepCopyInto(out)
	return out
}
