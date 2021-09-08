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
func (in *ActionThresholdObservation) DeepCopyInto(out *ActionThresholdObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionThresholdObservation.
func (in *ActionThresholdObservation) DeepCopy() *ActionThresholdObservation {
	if in == nil {
		return nil
	}
	out := new(ActionThresholdObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionThresholdParameters) DeepCopyInto(out *ActionThresholdParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionThresholdParameters.
func (in *ActionThresholdParameters) DeepCopy() *ActionThresholdParameters {
	if in == nil {
		return nil
	}
	out := new(ActionThresholdParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetsBudget) DeepCopyInto(out *BudgetsBudget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetsBudget.
func (in *BudgetsBudget) DeepCopy() *BudgetsBudget {
	if in == nil {
		return nil
	}
	out := new(BudgetsBudget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BudgetsBudget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetsBudgetAction) DeepCopyInto(out *BudgetsBudgetAction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetsBudgetAction.
func (in *BudgetsBudgetAction) DeepCopy() *BudgetsBudgetAction {
	if in == nil {
		return nil
	}
	out := new(BudgetsBudgetAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BudgetsBudgetAction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetsBudgetActionList) DeepCopyInto(out *BudgetsBudgetActionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BudgetsBudgetAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetsBudgetActionList.
func (in *BudgetsBudgetActionList) DeepCopy() *BudgetsBudgetActionList {
	if in == nil {
		return nil
	}
	out := new(BudgetsBudgetActionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BudgetsBudgetActionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetsBudgetActionObservation) DeepCopyInto(out *BudgetsBudgetActionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetsBudgetActionObservation.
func (in *BudgetsBudgetActionObservation) DeepCopy() *BudgetsBudgetActionObservation {
	if in == nil {
		return nil
	}
	out := new(BudgetsBudgetActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetsBudgetActionParameters) DeepCopyInto(out *BudgetsBudgetActionParameters) {
	*out = *in
	if in.AccountId != nil {
		in, out := &in.AccountId, &out.AccountId
		*out = new(string)
		**out = **in
	}
	if in.ActionThreshold != nil {
		in, out := &in.ActionThreshold, &out.ActionThreshold
		*out = make([]ActionThresholdParameters, len(*in))
		copy(*out, *in)
	}
	if in.Definition != nil {
		in, out := &in.Definition, &out.Definition
		*out = make([]DefinitionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Subscriber != nil {
		in, out := &in.Subscriber, &out.Subscriber
		*out = make([]SubscriberParameters, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetsBudgetActionParameters.
func (in *BudgetsBudgetActionParameters) DeepCopy() *BudgetsBudgetActionParameters {
	if in == nil {
		return nil
	}
	out := new(BudgetsBudgetActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetsBudgetActionSpec) DeepCopyInto(out *BudgetsBudgetActionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetsBudgetActionSpec.
func (in *BudgetsBudgetActionSpec) DeepCopy() *BudgetsBudgetActionSpec {
	if in == nil {
		return nil
	}
	out := new(BudgetsBudgetActionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetsBudgetActionStatus) DeepCopyInto(out *BudgetsBudgetActionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetsBudgetActionStatus.
func (in *BudgetsBudgetActionStatus) DeepCopy() *BudgetsBudgetActionStatus {
	if in == nil {
		return nil
	}
	out := new(BudgetsBudgetActionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetsBudgetList) DeepCopyInto(out *BudgetsBudgetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BudgetsBudget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetsBudgetList.
func (in *BudgetsBudgetList) DeepCopy() *BudgetsBudgetList {
	if in == nil {
		return nil
	}
	out := new(BudgetsBudgetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BudgetsBudgetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetsBudgetObservation) DeepCopyInto(out *BudgetsBudgetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetsBudgetObservation.
func (in *BudgetsBudgetObservation) DeepCopy() *BudgetsBudgetObservation {
	if in == nil {
		return nil
	}
	out := new(BudgetsBudgetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetsBudgetParameters) DeepCopyInto(out *BudgetsBudgetParameters) {
	*out = *in
	if in.AccountId != nil {
		in, out := &in.AccountId, &out.AccountId
		*out = new(string)
		**out = **in
	}
	if in.CostFilter != nil {
		in, out := &in.CostFilter, &out.CostFilter
		*out = make([]CostFilterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CostFilters != nil {
		in, out := &in.CostFilters, &out.CostFilters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CostTypes != nil {
		in, out := &in.CostTypes, &out.CostTypes
		*out = make([]CostTypesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamePrefix != nil {
		in, out := &in.NamePrefix, &out.NamePrefix
		*out = new(string)
		**out = **in
	}
	if in.Notification != nil {
		in, out := &in.Notification, &out.Notification
		*out = make([]NotificationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TimePeriodEnd != nil {
		in, out := &in.TimePeriodEnd, &out.TimePeriodEnd
		*out = new(string)
		**out = **in
	}
	if in.TimePeriodStart != nil {
		in, out := &in.TimePeriodStart, &out.TimePeriodStart
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetsBudgetParameters.
func (in *BudgetsBudgetParameters) DeepCopy() *BudgetsBudgetParameters {
	if in == nil {
		return nil
	}
	out := new(BudgetsBudgetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetsBudgetSpec) DeepCopyInto(out *BudgetsBudgetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetsBudgetSpec.
func (in *BudgetsBudgetSpec) DeepCopy() *BudgetsBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(BudgetsBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BudgetsBudgetStatus) DeepCopyInto(out *BudgetsBudgetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BudgetsBudgetStatus.
func (in *BudgetsBudgetStatus) DeepCopy() *BudgetsBudgetStatus {
	if in == nil {
		return nil
	}
	out := new(BudgetsBudgetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CostFilterObservation) DeepCopyInto(out *CostFilterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CostFilterObservation.
func (in *CostFilterObservation) DeepCopy() *CostFilterObservation {
	if in == nil {
		return nil
	}
	out := new(CostFilterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CostFilterParameters) DeepCopyInto(out *CostFilterParameters) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CostFilterParameters.
func (in *CostFilterParameters) DeepCopy() *CostFilterParameters {
	if in == nil {
		return nil
	}
	out := new(CostFilterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CostTypesObservation) DeepCopyInto(out *CostTypesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CostTypesObservation.
func (in *CostTypesObservation) DeepCopy() *CostTypesObservation {
	if in == nil {
		return nil
	}
	out := new(CostTypesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CostTypesParameters) DeepCopyInto(out *CostTypesParameters) {
	*out = *in
	if in.IncludeCredit != nil {
		in, out := &in.IncludeCredit, &out.IncludeCredit
		*out = new(bool)
		**out = **in
	}
	if in.IncludeDiscount != nil {
		in, out := &in.IncludeDiscount, &out.IncludeDiscount
		*out = new(bool)
		**out = **in
	}
	if in.IncludeOtherSubscription != nil {
		in, out := &in.IncludeOtherSubscription, &out.IncludeOtherSubscription
		*out = new(bool)
		**out = **in
	}
	if in.IncludeRecurring != nil {
		in, out := &in.IncludeRecurring, &out.IncludeRecurring
		*out = new(bool)
		**out = **in
	}
	if in.IncludeRefund != nil {
		in, out := &in.IncludeRefund, &out.IncludeRefund
		*out = new(bool)
		**out = **in
	}
	if in.IncludeSubscription != nil {
		in, out := &in.IncludeSubscription, &out.IncludeSubscription
		*out = new(bool)
		**out = **in
	}
	if in.IncludeSupport != nil {
		in, out := &in.IncludeSupport, &out.IncludeSupport
		*out = new(bool)
		**out = **in
	}
	if in.IncludeTax != nil {
		in, out := &in.IncludeTax, &out.IncludeTax
		*out = new(bool)
		**out = **in
	}
	if in.IncludeUpfront != nil {
		in, out := &in.IncludeUpfront, &out.IncludeUpfront
		*out = new(bool)
		**out = **in
	}
	if in.UseAmortized != nil {
		in, out := &in.UseAmortized, &out.UseAmortized
		*out = new(bool)
		**out = **in
	}
	if in.UseBlended != nil {
		in, out := &in.UseBlended, &out.UseBlended
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CostTypesParameters.
func (in *CostTypesParameters) DeepCopy() *CostTypesParameters {
	if in == nil {
		return nil
	}
	out := new(CostTypesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionObservation) DeepCopyInto(out *DefinitionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionObservation.
func (in *DefinitionObservation) DeepCopy() *DefinitionObservation {
	if in == nil {
		return nil
	}
	out := new(DefinitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionParameters) DeepCopyInto(out *DefinitionParameters) {
	*out = *in
	if in.IamActionDefinition != nil {
		in, out := &in.IamActionDefinition, &out.IamActionDefinition
		*out = make([]IamActionDefinitionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ScpActionDefinition != nil {
		in, out := &in.ScpActionDefinition, &out.ScpActionDefinition
		*out = make([]ScpActionDefinitionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SsmActionDefinition != nil {
		in, out := &in.SsmActionDefinition, &out.SsmActionDefinition
		*out = make([]SsmActionDefinitionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionParameters.
func (in *DefinitionParameters) DeepCopy() *DefinitionParameters {
	if in == nil {
		return nil
	}
	out := new(DefinitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamActionDefinitionObservation) DeepCopyInto(out *IamActionDefinitionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamActionDefinitionObservation.
func (in *IamActionDefinitionObservation) DeepCopy() *IamActionDefinitionObservation {
	if in == nil {
		return nil
	}
	out := new(IamActionDefinitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamActionDefinitionParameters) DeepCopyInto(out *IamActionDefinitionParameters) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamActionDefinitionParameters.
func (in *IamActionDefinitionParameters) DeepCopy() *IamActionDefinitionParameters {
	if in == nil {
		return nil
	}
	out := new(IamActionDefinitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationObservation) DeepCopyInto(out *NotificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationObservation.
func (in *NotificationObservation) DeepCopy() *NotificationObservation {
	if in == nil {
		return nil
	}
	out := new(NotificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationParameters) DeepCopyInto(out *NotificationParameters) {
	*out = *in
	if in.SubscriberEmailAddresses != nil {
		in, out := &in.SubscriberEmailAddresses, &out.SubscriberEmailAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubscriberSnsTopicArns != nil {
		in, out := &in.SubscriberSnsTopicArns, &out.SubscriberSnsTopicArns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationParameters.
func (in *NotificationParameters) DeepCopy() *NotificationParameters {
	if in == nil {
		return nil
	}
	out := new(NotificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScpActionDefinitionObservation) DeepCopyInto(out *ScpActionDefinitionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScpActionDefinitionObservation.
func (in *ScpActionDefinitionObservation) DeepCopy() *ScpActionDefinitionObservation {
	if in == nil {
		return nil
	}
	out := new(ScpActionDefinitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScpActionDefinitionParameters) DeepCopyInto(out *ScpActionDefinitionParameters) {
	*out = *in
	if in.TargetIds != nil {
		in, out := &in.TargetIds, &out.TargetIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScpActionDefinitionParameters.
func (in *ScpActionDefinitionParameters) DeepCopy() *ScpActionDefinitionParameters {
	if in == nil {
		return nil
	}
	out := new(ScpActionDefinitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsmActionDefinitionObservation) DeepCopyInto(out *SsmActionDefinitionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsmActionDefinitionObservation.
func (in *SsmActionDefinitionObservation) DeepCopy() *SsmActionDefinitionObservation {
	if in == nil {
		return nil
	}
	out := new(SsmActionDefinitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsmActionDefinitionParameters) DeepCopyInto(out *SsmActionDefinitionParameters) {
	*out = *in
	if in.InstanceIds != nil {
		in, out := &in.InstanceIds, &out.InstanceIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsmActionDefinitionParameters.
func (in *SsmActionDefinitionParameters) DeepCopy() *SsmActionDefinitionParameters {
	if in == nil {
		return nil
	}
	out := new(SsmActionDefinitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriberObservation) DeepCopyInto(out *SubscriberObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriberObservation.
func (in *SubscriberObservation) DeepCopy() *SubscriberObservation {
	if in == nil {
		return nil
	}
	out := new(SubscriberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriberParameters) DeepCopyInto(out *SubscriberParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriberParameters.
func (in *SubscriberParameters) DeepCopy() *SubscriberParameters {
	if in == nil {
		return nil
	}
	out := new(SubscriberParameters)
	in.DeepCopyInto(out)
	return out
}
