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

type Apigatewayv2ModelObservation struct {
}

type Apigatewayv2ModelParameters struct {

	// +kubebuilder:validation:Required
	APIID string `json:"apiId" tf:"api_id"`

	// +kubebuilder:validation:Required
	ContentType string `json:"contentType" tf:"content_type"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	Schema string `json:"schema" tf:"schema"`
}

// Apigatewayv2ModelSpec defines the desired state of Apigatewayv2Model
type Apigatewayv2ModelSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Apigatewayv2ModelParameters `json:"forProvider"`
}

// Apigatewayv2ModelStatus defines the observed state of Apigatewayv2Model.
type Apigatewayv2ModelStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Apigatewayv2ModelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2Model is the Schema for the Apigatewayv2Models API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Apigatewayv2Model struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Apigatewayv2ModelSpec   `json:"spec"`
	Status            Apigatewayv2ModelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2ModelList contains a list of Apigatewayv2Models
type Apigatewayv2ModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2Model `json:"items"`
}

// Repository type metadata.
var (
	Apigatewayv2ModelKind             = "Apigatewayv2Model"
	Apigatewayv2ModelGroupKind        = schema.GroupKind{Group: Group, Kind: Apigatewayv2ModelKind}.String()
	Apigatewayv2ModelKindAPIVersion   = Apigatewayv2ModelKind + "." + GroupVersion.String()
	Apigatewayv2ModelGroupVersionKind = GroupVersion.WithKind(Apigatewayv2ModelKind)
)

func init() {
	SchemeBuilder.Register(&Apigatewayv2Model{}, &Apigatewayv2ModelList{})
}