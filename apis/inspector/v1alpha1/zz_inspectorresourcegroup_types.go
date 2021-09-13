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

type InspectorResourceGroupObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type InspectorResourceGroupParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	Tags map[string]string `json:"tags" tf:"tags"`
}

// InspectorResourceGroupSpec defines the desired state of InspectorResourceGroup
type InspectorResourceGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       InspectorResourceGroupParameters `json:"forProvider"`
}

// InspectorResourceGroupStatus defines the observed state of InspectorResourceGroup.
type InspectorResourceGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          InspectorResourceGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InspectorResourceGroup is the Schema for the InspectorResourceGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type InspectorResourceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InspectorResourceGroupSpec   `json:"spec"`
	Status            InspectorResourceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InspectorResourceGroupList contains a list of InspectorResourceGroups
type InspectorResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InspectorResourceGroup `json:"items"`
}

// Repository type metadata.
var (
	InspectorResourceGroupKind             = "InspectorResourceGroup"
	InspectorResourceGroupGroupKind        = schema.GroupKind{Group: Group, Kind: InspectorResourceGroupKind}.String()
	InspectorResourceGroupKindAPIVersion   = InspectorResourceGroupKind + "." + GroupVersion.String()
	InspectorResourceGroupGroupVersionKind = GroupVersion.WithKind(InspectorResourceGroupKind)
)

func init() {
	SchemeBuilder.Register(&InspectorResourceGroup{}, &InspectorResourceGroupList{})
}