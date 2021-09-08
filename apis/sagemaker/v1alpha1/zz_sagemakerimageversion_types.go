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

type SagemakerImageVersionObservation struct {
	Arn string `json:"arn" tf:"arn"`

	ContainerImage string `json:"containerImage" tf:"container_image"`

	ImageArn string `json:"imageArn" tf:"image_arn"`

	Version int64 `json:"version" tf:"version"`
}

type SagemakerImageVersionParameters struct {
	BaseImage string `json:"baseImage" tf:"base_image"`

	ImageName string `json:"imageName" tf:"image_name"`
}

// SagemakerImageVersionSpec defines the desired state of SagemakerImageVersion
type SagemakerImageVersionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SagemakerImageVersionParameters `json:"forProvider"`
}

// SagemakerImageVersionStatus defines the observed state of SagemakerImageVersion.
type SagemakerImageVersionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SagemakerImageVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerImageVersion is the Schema for the SagemakerImageVersions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SagemakerImageVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SagemakerImageVersionSpec   `json:"spec"`
	Status            SagemakerImageVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerImageVersionList contains a list of SagemakerImageVersions
type SagemakerImageVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SagemakerImageVersion `json:"items"`
}

// Repository type metadata.
var (
	SagemakerImageVersionKind             = "SagemakerImageVersion"
	SagemakerImageVersionGroupKind        = schema.GroupKind{Group: Group, Kind: SagemakerImageVersionKind}.String()
	SagemakerImageVersionKindAPIVersion   = SagemakerImageVersionKind + "." + GroupVersion.String()
	SagemakerImageVersionGroupVersionKind = GroupVersion.WithKind(SagemakerImageVersionKind)
)

func init() {
	SchemeBuilder.Register(&SagemakerImageVersion{}, &SagemakerImageVersionList{})
}
