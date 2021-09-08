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

type Ec2TransitGatewayPrefixListReferenceObservation struct {
	PrefixListOwnerId string `json:"prefixListOwnerId" tf:"prefix_list_owner_id"`
}

type Ec2TransitGatewayPrefixListReferenceParameters struct {
	Blackhole *bool `json:"blackhole,omitempty" tf:"blackhole"`

	PrefixListId string `json:"prefixListId" tf:"prefix_list_id"`

	TransitGatewayAttachmentId *string `json:"transitGatewayAttachmentId,omitempty" tf:"transit_gateway_attachment_id"`

	TransitGatewayRouteTableId string `json:"transitGatewayRouteTableId" tf:"transit_gateway_route_table_id"`
}

// Ec2TransitGatewayPrefixListReferenceSpec defines the desired state of Ec2TransitGatewayPrefixListReference
type Ec2TransitGatewayPrefixListReferenceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Ec2TransitGatewayPrefixListReferenceParameters `json:"forProvider"`
}

// Ec2TransitGatewayPrefixListReferenceStatus defines the observed state of Ec2TransitGatewayPrefixListReference.
type Ec2TransitGatewayPrefixListReferenceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Ec2TransitGatewayPrefixListReferenceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayPrefixListReference is the Schema for the Ec2TransitGatewayPrefixListReferences API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Ec2TransitGatewayPrefixListReference struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2TransitGatewayPrefixListReferenceSpec   `json:"spec"`
	Status            Ec2TransitGatewayPrefixListReferenceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayPrefixListReferenceList contains a list of Ec2TransitGatewayPrefixListReferences
type Ec2TransitGatewayPrefixListReferenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TransitGatewayPrefixListReference `json:"items"`
}

// Repository type metadata.
var (
	Ec2TransitGatewayPrefixListReferenceKind             = "Ec2TransitGatewayPrefixListReference"
	Ec2TransitGatewayPrefixListReferenceGroupKind        = schema.GroupKind{Group: Group, Kind: Ec2TransitGatewayPrefixListReferenceKind}.String()
	Ec2TransitGatewayPrefixListReferenceKindAPIVersion   = Ec2TransitGatewayPrefixListReferenceKind + "." + GroupVersion.String()
	Ec2TransitGatewayPrefixListReferenceGroupVersionKind = GroupVersion.WithKind(Ec2TransitGatewayPrefixListReferenceKind)
)

func init() {
	SchemeBuilder.Register(&Ec2TransitGatewayPrefixListReference{}, &Ec2TransitGatewayPrefixListReferenceList{})
}
