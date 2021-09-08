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

type AccountsObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Email string `json:"email" tf:"email"`

	Id string `json:"id" tf:"id"`

	Name string `json:"name" tf:"name"`

	Status string `json:"status" tf:"status"`
}

type AccountsParameters struct {
}

type NonMasterAccountsObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Email string `json:"email" tf:"email"`

	Id string `json:"id" tf:"id"`

	Name string `json:"name" tf:"name"`

	Status string `json:"status" tf:"status"`
}

type NonMasterAccountsParameters struct {
}

type OrganizationsOrganizationObservation struct {
	Accounts []AccountsObservation `json:"accounts" tf:"accounts"`

	Arn string `json:"arn" tf:"arn"`

	MasterAccountArn string `json:"masterAccountArn" tf:"master_account_arn"`

	MasterAccountEmail string `json:"masterAccountEmail" tf:"master_account_email"`

	MasterAccountId string `json:"masterAccountId" tf:"master_account_id"`

	NonMasterAccounts []NonMasterAccountsObservation `json:"nonMasterAccounts" tf:"non_master_accounts"`

	Roots []RootsObservation `json:"roots" tf:"roots"`
}

type OrganizationsOrganizationParameters struct {
	AwsServiceAccessPrincipals []string `json:"awsServiceAccessPrincipals,omitempty" tf:"aws_service_access_principals"`

	EnabledPolicyTypes []string `json:"enabledPolicyTypes,omitempty" tf:"enabled_policy_types"`

	FeatureSet *string `json:"featureSet,omitempty" tf:"feature_set"`
}

type PolicyTypesObservation struct {
	Status string `json:"status" tf:"status"`

	Type string `json:"type" tf:"type"`
}

type PolicyTypesParameters struct {
}

type RootsObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Id string `json:"id" tf:"id"`

	Name string `json:"name" tf:"name"`

	PolicyTypes []PolicyTypesObservation `json:"policyTypes" tf:"policy_types"`
}

type RootsParameters struct {
}

// OrganizationsOrganizationSpec defines the desired state of OrganizationsOrganization
type OrganizationsOrganizationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       OrganizationsOrganizationParameters `json:"forProvider"`
}

// OrganizationsOrganizationStatus defines the observed state of OrganizationsOrganization.
type OrganizationsOrganizationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          OrganizationsOrganizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationsOrganization is the Schema for the OrganizationsOrganizations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type OrganizationsOrganization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationsOrganizationSpec   `json:"spec"`
	Status            OrganizationsOrganizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationsOrganizationList contains a list of OrganizationsOrganizations
type OrganizationsOrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationsOrganization `json:"items"`
}

// Repository type metadata.
var (
	OrganizationsOrganizationKind             = "OrganizationsOrganization"
	OrganizationsOrganizationGroupKind        = schema.GroupKind{Group: Group, Kind: OrganizationsOrganizationKind}.String()
	OrganizationsOrganizationKindAPIVersion   = OrganizationsOrganizationKind + "." + GroupVersion.String()
	OrganizationsOrganizationGroupVersionKind = GroupVersion.WithKind(OrganizationsOrganizationKind)
)

func init() {
	SchemeBuilder.Register(&OrganizationsOrganization{}, &OrganizationsOrganizationList{})
}
