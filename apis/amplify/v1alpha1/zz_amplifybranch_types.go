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

type AmplifyBranchObservation struct {
	Arn string `json:"arn" tf:"arn"`

	AssociatedResources []string `json:"associatedResources" tf:"associated_resources"`

	CustomDomains []string `json:"customDomains" tf:"custom_domains"`

	DestinationBranch string `json:"destinationBranch" tf:"destination_branch"`

	SourceBranch string `json:"sourceBranch" tf:"source_branch"`
}

type AmplifyBranchParameters struct {
	AppId string `json:"appId" tf:"app_id"`

	BackendEnvironmentArn *string `json:"backendEnvironmentArn,omitempty" tf:"backend_environment_arn"`

	BasicAuthCredentials *string `json:"basicAuthCredentials,omitempty" tf:"basic_auth_credentials"`

	BranchName string `json:"branchName" tf:"branch_name"`

	Description *string `json:"description,omitempty" tf:"description"`

	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`

	EnableAutoBuild *bool `json:"enableAutoBuild,omitempty" tf:"enable_auto_build"`

	EnableBasicAuth *bool `json:"enableBasicAuth,omitempty" tf:"enable_basic_auth"`

	EnableNotification *bool `json:"enableNotification,omitempty" tf:"enable_notification"`

	EnablePerformanceMode *bool `json:"enablePerformanceMode,omitempty" tf:"enable_performance_mode"`

	EnablePullRequestPreview *bool `json:"enablePullRequestPreview,omitempty" tf:"enable_pull_request_preview"`

	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty" tf:"environment_variables"`

	Framework *string `json:"framework,omitempty" tf:"framework"`

	PullRequestEnvironmentName *string `json:"pullRequestEnvironmentName,omitempty" tf:"pull_request_environment_name"`

	Stage *string `json:"stage,omitempty" tf:"stage"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	Ttl *string `json:"ttl,omitempty" tf:"ttl"`
}

// AmplifyBranchSpec defines the desired state of AmplifyBranch
type AmplifyBranchSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AmplifyBranchParameters `json:"forProvider"`
}

// AmplifyBranchStatus defines the observed state of AmplifyBranch.
type AmplifyBranchStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AmplifyBranchObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AmplifyBranch is the Schema for the AmplifyBranchs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AmplifyBranch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AmplifyBranchSpec   `json:"spec"`
	Status            AmplifyBranchStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AmplifyBranchList contains a list of AmplifyBranchs
type AmplifyBranchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AmplifyBranch `json:"items"`
}

// Repository type metadata.
var (
	AmplifyBranchKind             = "AmplifyBranch"
	AmplifyBranchGroupKind        = schema.GroupKind{Group: Group, Kind: AmplifyBranchKind}.String()
	AmplifyBranchKindAPIVersion   = AmplifyBranchKind + "." + GroupVersion.String()
	AmplifyBranchGroupVersionKind = GroupVersion.WithKind(AmplifyBranchKind)
)

func init() {
	SchemeBuilder.Register(&AmplifyBranch{}, &AmplifyBranchList{})
}
