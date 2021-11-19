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

type BucketInventoryDestinationObservation struct {
}

type BucketInventoryDestinationParameters struct {

	// +kubebuilder:validation:Required
	Bucket []DestinationBucketParameters `json:"bucket" tf:"bucket,omitempty"`
}

type BucketInventoryFilterObservation struct {
}

type BucketInventoryFilterParameters struct {

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type BucketInventoryObservation struct {
}

type BucketInventoryParameters struct {

	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Required
	Destination []BucketInventoryDestinationParameters `json:"destination" tf:"destination,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Filter []BucketInventoryFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// +kubebuilder:validation:Required
	IncludedObjectVersions *string `json:"includedObjectVersions" tf:"included_object_versions,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	OptionalFields []*string `json:"optionalFields,omitempty" tf:"optional_fields,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	Schedule []ScheduleParameters `json:"schedule" tf:"schedule,omitempty"`
}

type DestinationBucketObservation struct {
}

type DestinationBucketParameters struct {

	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// +kubebuilder:validation:Required
	BucketArn *string `json:"bucketArn" tf:"bucket_arn,omitempty"`

	// +kubebuilder:validation:Optional
	Encryption []EncryptionParameters `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// +kubebuilder:validation:Required
	Format *string `json:"format" tf:"format,omitempty"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type EncryptionObservation struct {
}

type EncryptionParameters struct {

	// +kubebuilder:validation:Optional
	SseKms []SseKmsParameters `json:"sseKms,omitempty" tf:"sse_kms,omitempty"`

	// +kubebuilder:validation:Optional
	SseS3 []SseS3Parameters `json:"sseS3,omitempty" tf:"sse_s3,omitempty"`
}

type ScheduleObservation struct {
}

type ScheduleParameters struct {

	// +kubebuilder:validation:Required
	Frequency *string `json:"frequency" tf:"frequency,omitempty"`
}

type SseKmsObservation struct {
}

type SseKmsParameters struct {

	// +kubebuilder:validation:Required
	KeyID *string `json:"keyId" tf:"key_id,omitempty"`
}

type SseS3Observation struct {
}

type SseS3Parameters struct {
}

// BucketInventorySpec defines the desired state of BucketInventory
type BucketInventorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketInventoryParameters `json:"forProvider"`
}

// BucketInventoryStatus defines the observed state of BucketInventory.
type BucketInventoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketInventoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketInventory is the Schema for the BucketInventorys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type BucketInventory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketInventorySpec   `json:"spec"`
	Status            BucketInventoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketInventoryList contains a list of BucketInventorys
type BucketInventoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketInventory `json:"items"`
}

// Repository type metadata.
var (
	BucketInventory_Kind             = "BucketInventory"
	BucketInventory_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketInventory_Kind}.String()
	BucketInventory_KindAPIVersion   = BucketInventory_Kind + "." + CRDGroupVersion.String()
	BucketInventory_GroupVersionKind = CRDGroupVersion.WithKind(BucketInventory_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketInventory{}, &BucketInventoryList{})
}