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

type Apigatewayv2AuthorizerObservation struct {
}

type Apigatewayv2AuthorizerParameters struct {
	ApiId string `json:"apiId" tf:"api_id"`

	AuthorizerCredentialsArn *string `json:"authorizerCredentialsArn,omitempty" tf:"authorizer_credentials_arn"`

	AuthorizerPayloadFormatVersion *string `json:"authorizerPayloadFormatVersion,omitempty" tf:"authorizer_payload_format_version"`

	AuthorizerResultTtlInSeconds *int64 `json:"authorizerResultTtlInSeconds,omitempty" tf:"authorizer_result_ttl_in_seconds"`

	AuthorizerType string `json:"authorizerType" tf:"authorizer_type"`

	AuthorizerUri *string `json:"authorizerUri,omitempty" tf:"authorizer_uri"`

	EnableSimpleResponses *bool `json:"enableSimpleResponses,omitempty" tf:"enable_simple_responses"`

	IdentitySources []string `json:"identitySources,omitempty" tf:"identity_sources"`

	JwtConfiguration []JwtConfigurationParameters `json:"jwtConfiguration,omitempty" tf:"jwt_configuration"`

	Name string `json:"name" tf:"name"`
}

type JwtConfigurationObservation struct {
}

type JwtConfigurationParameters struct {
	Audience []string `json:"audience,omitempty" tf:"audience"`

	Issuer *string `json:"issuer,omitempty" tf:"issuer"`
}

// Apigatewayv2AuthorizerSpec defines the desired state of Apigatewayv2Authorizer
type Apigatewayv2AuthorizerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Apigatewayv2AuthorizerParameters `json:"forProvider"`
}

// Apigatewayv2AuthorizerStatus defines the observed state of Apigatewayv2Authorizer.
type Apigatewayv2AuthorizerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Apigatewayv2AuthorizerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2Authorizer is the Schema for the Apigatewayv2Authorizers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Apigatewayv2Authorizer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Apigatewayv2AuthorizerSpec   `json:"spec"`
	Status            Apigatewayv2AuthorizerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Apigatewayv2AuthorizerList contains a list of Apigatewayv2Authorizers
type Apigatewayv2AuthorizerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Apigatewayv2Authorizer `json:"items"`
}

// Repository type metadata.
var (
	Apigatewayv2AuthorizerKind             = "Apigatewayv2Authorizer"
	Apigatewayv2AuthorizerGroupKind        = schema.GroupKind{Group: Group, Kind: Apigatewayv2AuthorizerKind}.String()
	Apigatewayv2AuthorizerKindAPIVersion   = Apigatewayv2AuthorizerKind + "." + GroupVersion.String()
	Apigatewayv2AuthorizerGroupVersionKind = GroupVersion.WithKind(Apigatewayv2AuthorizerKind)
)

func init() {
	SchemeBuilder.Register(&Apigatewayv2Authorizer{}, &Apigatewayv2AuthorizerList{})
}
