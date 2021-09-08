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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type CloudformationStackSetInstanceObservation struct {
	StackId string `json:"stackId" tf:"stack_id"`
}

type CloudformationStackSetInstanceParameters struct {
	AccountId *string `json:"accountId,omitempty" tf:"account_id"`

	ParameterOverrides map[string]string `json:"parameterOverrides,omitempty" tf:"parameter_overrides"`

	Region *string `json:"region,omitempty" tf:"region"`

	RetainStack *bool `json:"retainStack,omitempty" tf:"retain_stack"`

	StackSetName string `json:"stackSetName" tf:"stack_set_name"`
}

// CloudformationStackSetInstanceSpec defines the desired state of CloudformationStackSetInstance
type CloudformationStackSetInstanceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudformationStackSetInstanceParameters `json:"forProvider"`
}

// CloudformationStackSetInstanceStatus defines the observed state of CloudformationStackSetInstance.
type CloudformationStackSetInstanceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudformationStackSetInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudformationStackSetInstance is the Schema for the CloudformationStackSetInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CloudformationStackSetInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudformationStackSetInstanceSpec   `json:"spec"`
	Status            CloudformationStackSetInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudformationStackSetInstanceList contains a list of CloudformationStackSetInstances
type CloudformationStackSetInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudformationStackSetInstance `json:"items"`
}

// Repository type metadata.
var (
	CloudformationStackSetInstanceKind             = "CloudformationStackSetInstance"
	CloudformationStackSetInstanceGroupKind        = schema.GroupKind{Group: Group, Kind: CloudformationStackSetInstanceKind}.String()
	CloudformationStackSetInstanceKindAPIVersion   = CloudformationStackSetInstanceKind + "." + GroupVersion.String()
	CloudformationStackSetInstanceGroupVersionKind = GroupVersion.WithKind(CloudformationStackSetInstanceKind)
)

func init() {
	SchemeBuilder.Register(&CloudformationStackSetInstance{}, &CloudformationStackSetInstanceList{})
}
