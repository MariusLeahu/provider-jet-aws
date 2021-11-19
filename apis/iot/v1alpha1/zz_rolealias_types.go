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

type RoleAliasObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type RoleAliasParameters struct {

	// +kubebuilder:validation:Required
	Alias *string `json:"alias" tf:"alias,omitempty"`

	// +kubebuilder:validation:Optional
	CredentialDuration *int64 `json:"credentialDuration,omitempty" tf:"credential_duration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`
}

// RoleAliasSpec defines the desired state of RoleAlias
type RoleAliasSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleAliasParameters `json:"forProvider"`
}

// RoleAliasStatus defines the observed state of RoleAlias.
type RoleAliasStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleAliasObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAlias is the Schema for the RoleAliass API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type RoleAlias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleAliasSpec   `json:"spec"`
	Status            RoleAliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAliasList contains a list of RoleAliass
type RoleAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleAlias `json:"items"`
}

// Repository type metadata.
var (
	RoleAlias_Kind             = "RoleAlias"
	RoleAlias_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleAlias_Kind}.String()
	RoleAlias_KindAPIVersion   = RoleAlias_Kind + "." + CRDGroupVersion.String()
	RoleAlias_GroupVersionKind = CRDGroupVersion.WithKind(RoleAlias_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleAlias{}, &RoleAliasList{})
}