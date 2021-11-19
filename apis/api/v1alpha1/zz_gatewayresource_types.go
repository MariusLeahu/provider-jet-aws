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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GatewayResourceObservation struct {
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type GatewayResourceParameters struct {

	// +kubebuilder:validation:Required
	ParentID *string `json:"parentId" tf:"parent_id,omitempty"`

	// +kubebuilder:validation:Required
	PathPart *string `json:"pathPart" tf:"path_part,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	RestAPIID *string `json:"restApiId" tf:"rest_api_id,omitempty"`
}

// GatewayResourceSpec defines the desired state of GatewayResource
type GatewayResourceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayResourceParameters `json:"forProvider"`
}

// GatewayResourceStatus defines the observed state of GatewayResource.
type GatewayResourceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayResourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayResource is the Schema for the GatewayResources API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type GatewayResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewayResourceSpec   `json:"spec"`
	Status            GatewayResourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayResourceList contains a list of GatewayResources
type GatewayResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayResource `json:"items"`
}

// Repository type metadata.
var (
	GatewayResource_Kind             = "GatewayResource"
	GatewayResource_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GatewayResource_Kind}.String()
	GatewayResource_KindAPIVersion   = GatewayResource_Kind + "." + CRDGroupVersion.String()
	GatewayResource_GroupVersionKind = CRDGroupVersion.WithKind(GatewayResource_Kind)
)

func init() {
	SchemeBuilder.Register(&GatewayResource{}, &GatewayResourceList{})
}