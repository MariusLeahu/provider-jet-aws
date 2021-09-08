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

type Ec2LocalGatewayRouteObservation struct {
}

type Ec2LocalGatewayRouteParameters struct {
	DestinationCidrBlock string `json:"destinationCidrBlock" tf:"destination_cidr_block"`

	LocalGatewayRouteTableId string `json:"localGatewayRouteTableId" tf:"local_gateway_route_table_id"`

	LocalGatewayVirtualInterfaceGroupId string `json:"localGatewayVirtualInterfaceGroupId" tf:"local_gateway_virtual_interface_group_id"`
}

// Ec2LocalGatewayRouteSpec defines the desired state of Ec2LocalGatewayRoute
type Ec2LocalGatewayRouteSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Ec2LocalGatewayRouteParameters `json:"forProvider"`
}

// Ec2LocalGatewayRouteStatus defines the observed state of Ec2LocalGatewayRoute.
type Ec2LocalGatewayRouteStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Ec2LocalGatewayRouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2LocalGatewayRoute is the Schema for the Ec2LocalGatewayRoutes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Ec2LocalGatewayRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2LocalGatewayRouteSpec   `json:"spec"`
	Status            Ec2LocalGatewayRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2LocalGatewayRouteList contains a list of Ec2LocalGatewayRoutes
type Ec2LocalGatewayRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2LocalGatewayRoute `json:"items"`
}

// Repository type metadata.
var (
	Ec2LocalGatewayRouteKind             = "Ec2LocalGatewayRoute"
	Ec2LocalGatewayRouteGroupKind        = schema.GroupKind{Group: Group, Kind: Ec2LocalGatewayRouteKind}.String()
	Ec2LocalGatewayRouteKindAPIVersion   = Ec2LocalGatewayRouteKind + "." + GroupVersion.String()
	Ec2LocalGatewayRouteGroupVersionKind = GroupVersion.WithKind(Ec2LocalGatewayRouteKind)
)

func init() {
	SchemeBuilder.Register(&Ec2LocalGatewayRoute{}, &Ec2LocalGatewayRouteList{})
}
