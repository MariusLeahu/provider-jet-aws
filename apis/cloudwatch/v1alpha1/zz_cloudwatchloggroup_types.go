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

type CloudwatchLogGroupObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type CloudwatchLogGroupParameters struct {

	// +kubebuilder:validation:Optional
	KmsKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// +kubebuilder:validation:Optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RetentionInDays *int64 `json:"retentionInDays,omitempty" tf:"retention_in_days"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// CloudwatchLogGroupSpec defines the desired state of CloudwatchLogGroup
type CloudwatchLogGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudwatchLogGroupParameters `json:"forProvider"`
}

// CloudwatchLogGroupStatus defines the observed state of CloudwatchLogGroup.
type CloudwatchLogGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudwatchLogGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchLogGroup is the Schema for the CloudwatchLogGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CloudwatchLogGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchLogGroupSpec   `json:"spec"`
	Status            CloudwatchLogGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchLogGroupList contains a list of CloudwatchLogGroups
type CloudwatchLogGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchLogGroup `json:"items"`
}

// Repository type metadata.
var (
	CloudwatchLogGroupKind             = "CloudwatchLogGroup"
	CloudwatchLogGroupGroupKind        = schema.GroupKind{Group: Group, Kind: CloudwatchLogGroupKind}.String()
	CloudwatchLogGroupKindAPIVersion   = CloudwatchLogGroupKind + "." + GroupVersion.String()
	CloudwatchLogGroupGroupVersionKind = GroupVersion.WithKind(CloudwatchLogGroupKind)
)

func init() {
	SchemeBuilder.Register(&CloudwatchLogGroup{}, &CloudwatchLogGroupList{})
}