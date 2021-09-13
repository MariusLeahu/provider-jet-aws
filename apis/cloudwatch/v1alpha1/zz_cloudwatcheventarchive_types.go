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

type CloudwatchEventArchiveObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type CloudwatchEventArchiveParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Optional
	EventPattern *string `json:"eventPattern,omitempty" tf:"event_pattern"`

	// +kubebuilder:validation:Required
	EventSourceArn string `json:"eventSourceArn" tf:"event_source_arn"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RetentionDays *int64 `json:"retentionDays,omitempty" tf:"retention_days"`
}

// CloudwatchEventArchiveSpec defines the desired state of CloudwatchEventArchive
type CloudwatchEventArchiveSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudwatchEventArchiveParameters `json:"forProvider"`
}

// CloudwatchEventArchiveStatus defines the observed state of CloudwatchEventArchive.
type CloudwatchEventArchiveStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudwatchEventArchiveObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventArchive is the Schema for the CloudwatchEventArchives API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CloudwatchEventArchive struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchEventArchiveSpec   `json:"spec"`
	Status            CloudwatchEventArchiveStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventArchiveList contains a list of CloudwatchEventArchives
type CloudwatchEventArchiveList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchEventArchive `json:"items"`
}

// Repository type metadata.
var (
	CloudwatchEventArchiveKind             = "CloudwatchEventArchive"
	CloudwatchEventArchiveGroupKind        = schema.GroupKind{Group: Group, Kind: CloudwatchEventArchiveKind}.String()
	CloudwatchEventArchiveKindAPIVersion   = CloudwatchEventArchiveKind + "." + GroupVersion.String()
	CloudwatchEventArchiveGroupVersionKind = GroupVersion.WithKind(CloudwatchEventArchiveKind)
)

func init() {
	SchemeBuilder.Register(&CloudwatchEventArchive{}, &CloudwatchEventArchiveList{})
}