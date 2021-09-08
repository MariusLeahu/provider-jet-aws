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

type RamPrincipalAssociationObservation struct {
}

type RamPrincipalAssociationParameters struct {
	Principal string `json:"principal" tf:"principal"`

	ResourceShareArn string `json:"resourceShareArn" tf:"resource_share_arn"`
}

// RamPrincipalAssociationSpec defines the desired state of RamPrincipalAssociation
type RamPrincipalAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       RamPrincipalAssociationParameters `json:"forProvider"`
}

// RamPrincipalAssociationStatus defines the observed state of RamPrincipalAssociation.
type RamPrincipalAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          RamPrincipalAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RamPrincipalAssociation is the Schema for the RamPrincipalAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type RamPrincipalAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RamPrincipalAssociationSpec   `json:"spec"`
	Status            RamPrincipalAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RamPrincipalAssociationList contains a list of RamPrincipalAssociations
type RamPrincipalAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RamPrincipalAssociation `json:"items"`
}

// Repository type metadata.
var (
	RamPrincipalAssociationKind             = "RamPrincipalAssociation"
	RamPrincipalAssociationGroupKind        = schema.GroupKind{Group: Group, Kind: RamPrincipalAssociationKind}.String()
	RamPrincipalAssociationKindAPIVersion   = RamPrincipalAssociationKind + "." + GroupVersion.String()
	RamPrincipalAssociationGroupVersionKind = GroupVersion.WithKind(RamPrincipalAssociationKind)
)

func init() {
	SchemeBuilder.Register(&RamPrincipalAssociation{}, &RamPrincipalAssociationList{})
}
