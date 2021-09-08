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

type AbortIncompleteMultipartUploadObservation struct {
}

type AbortIncompleteMultipartUploadParameters struct {
	DaysAfterInitiation int64 `json:"daysAfterInitiation" tf:"days_after_initiation"`
}

type ExpirationObservation struct {
}

type ExpirationParameters struct {
	Date *string `json:"date,omitempty" tf:"date"`

	Days *int64 `json:"days,omitempty" tf:"days"`

	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker"`
}

type FilterObservation struct {
}

type FilterParameters struct {
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type RuleObservation struct {
}

type RuleParameters struct {
	AbortIncompleteMultipartUpload []AbortIncompleteMultipartUploadParameters `json:"abortIncompleteMultipartUpload,omitempty" tf:"abort_incomplete_multipart_upload"`

	Expiration []ExpirationParameters `json:"expiration,omitempty" tf:"expiration"`

	Filter []FilterParameters `json:"filter,omitempty" tf:"filter"`

	Id string `json:"id" tf:"id"`

	Status *string `json:"status,omitempty" tf:"status"`
}

type S3ControlBucketLifecycleConfigurationObservation struct {
}

type S3ControlBucketLifecycleConfigurationParameters struct {
	Bucket string `json:"bucket" tf:"bucket"`

	Rule []RuleParameters `json:"rule" tf:"rule"`
}

// S3ControlBucketLifecycleConfigurationSpec defines the desired state of S3ControlBucketLifecycleConfiguration
type S3ControlBucketLifecycleConfigurationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       S3ControlBucketLifecycleConfigurationParameters `json:"forProvider"`
}

// S3ControlBucketLifecycleConfigurationStatus defines the observed state of S3ControlBucketLifecycleConfiguration.
type S3ControlBucketLifecycleConfigurationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          S3ControlBucketLifecycleConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// S3ControlBucketLifecycleConfiguration is the Schema for the S3ControlBucketLifecycleConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type S3ControlBucketLifecycleConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3ControlBucketLifecycleConfigurationSpec   `json:"spec"`
	Status            S3ControlBucketLifecycleConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3ControlBucketLifecycleConfigurationList contains a list of S3ControlBucketLifecycleConfigurations
type S3ControlBucketLifecycleConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3ControlBucketLifecycleConfiguration `json:"items"`
}

// Repository type metadata.
var (
	S3ControlBucketLifecycleConfigurationKind             = "S3ControlBucketLifecycleConfiguration"
	S3ControlBucketLifecycleConfigurationGroupKind        = schema.GroupKind{Group: Group, Kind: S3ControlBucketLifecycleConfigurationKind}.String()
	S3ControlBucketLifecycleConfigurationKindAPIVersion   = S3ControlBucketLifecycleConfigurationKind + "." + GroupVersion.String()
	S3ControlBucketLifecycleConfigurationGroupVersionKind = GroupVersion.WithKind(S3ControlBucketLifecycleConfigurationKind)
)

func init() {
	SchemeBuilder.Register(&S3ControlBucketLifecycleConfiguration{}, &S3ControlBucketLifecycleConfigurationList{})
}
