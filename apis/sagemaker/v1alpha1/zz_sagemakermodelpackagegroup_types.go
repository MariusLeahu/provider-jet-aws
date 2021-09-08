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

type SagemakerModelPackageGroupObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type SagemakerModelPackageGroupParameters struct {
	ModelPackageGroupDescription *string `json:"modelPackageGroupDescription,omitempty" tf:"model_package_group_description"`

	ModelPackageGroupName string `json:"modelPackageGroupName" tf:"model_package_group_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// SagemakerModelPackageGroupSpec defines the desired state of SagemakerModelPackageGroup
type SagemakerModelPackageGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SagemakerModelPackageGroupParameters `json:"forProvider"`
}

// SagemakerModelPackageGroupStatus defines the observed state of SagemakerModelPackageGroup.
type SagemakerModelPackageGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SagemakerModelPackageGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerModelPackageGroup is the Schema for the SagemakerModelPackageGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SagemakerModelPackageGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SagemakerModelPackageGroupSpec   `json:"spec"`
	Status            SagemakerModelPackageGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerModelPackageGroupList contains a list of SagemakerModelPackageGroups
type SagemakerModelPackageGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SagemakerModelPackageGroup `json:"items"`
}

// Repository type metadata.
var (
	SagemakerModelPackageGroupKind             = "SagemakerModelPackageGroup"
	SagemakerModelPackageGroupGroupKind        = schema.GroupKind{Group: Group, Kind: SagemakerModelPackageGroupKind}.String()
	SagemakerModelPackageGroupKindAPIVersion   = SagemakerModelPackageGroupKind + "." + GroupVersion.String()
	SagemakerModelPackageGroupGroupVersionKind = GroupVersion.WithKind(SagemakerModelPackageGroupKind)
)

func init() {
	SchemeBuilder.Register(&SagemakerModelPackageGroup{}, &SagemakerModelPackageGroupList{})
}
