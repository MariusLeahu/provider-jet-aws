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

type AttributeObservation struct {
}

type AttributeParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

type DynamodbTableObservation struct {
	Arn string `json:"arn" tf:"arn"`

	StreamArn string `json:"streamArn" tf:"stream_arn"`

	StreamLabel string `json:"streamLabel" tf:"stream_label"`
}

type DynamodbTableParameters struct {

	// +kubebuilder:validation:Required
	Attribute []AttributeParameters `json:"attribute" tf:"attribute"`

	// +kubebuilder:validation:Optional
	BillingMode *string `json:"billingMode,omitempty" tf:"billing_mode"`

	// +kubebuilder:validation:Optional
	GlobalSecondaryIndex []GlobalSecondaryIndexParameters `json:"globalSecondaryIndex,omitempty" tf:"global_secondary_index"`

	// +kubebuilder:validation:Required
	HashKey string `json:"hashKey" tf:"hash_key"`

	// +kubebuilder:validation:Optional
	LocalSecondaryIndex []LocalSecondaryIndexParameters `json:"localSecondaryIndex,omitempty" tf:"local_secondary_index"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	PointInTimeRecovery []PointInTimeRecoveryParameters `json:"pointInTimeRecovery,omitempty" tf:"point_in_time_recovery"`

	// +kubebuilder:validation:Optional
	RangeKey *string `json:"rangeKey,omitempty" tf:"range_key"`

	// +kubebuilder:validation:Optional
	ReadCapacity *int64 `json:"readCapacity,omitempty" tf:"read_capacity"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Replica []DynamodbTableReplicaParameters `json:"replica,omitempty" tf:"replica"`

	// +kubebuilder:validation:Optional
	ServerSideEncryption []ServerSideEncryptionParameters `json:"serverSideEncryption,omitempty" tf:"server_side_encryption"`

	// +kubebuilder:validation:Optional
	StreamEnabled *bool `json:"streamEnabled,omitempty" tf:"stream_enabled"`

	// +kubebuilder:validation:Optional
	StreamViewType *string `json:"streamViewType,omitempty" tf:"stream_view_type"`

	// +kubebuilder:validation:Optional
	TTL []TTLParameters `json:"ttl,omitempty" tf:"ttl"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Optional
	WriteCapacity *int64 `json:"writeCapacity,omitempty" tf:"write_capacity"`
}

type DynamodbTableReplicaObservation struct {
}

type DynamodbTableReplicaParameters struct {

	// +kubebuilder:validation:Optional
	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`

	// +kubebuilder:validation:Required
	RegionName string `json:"regionName" tf:"region_name"`
}

type GlobalSecondaryIndexObservation struct {
}

type GlobalSecondaryIndexParameters struct {

	// +kubebuilder:validation:Required
	HashKey string `json:"hashKey" tf:"hash_key"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	NonKeyAttributes []string `json:"nonKeyAttributes,omitempty" tf:"non_key_attributes"`

	// +kubebuilder:validation:Required
	ProjectionType string `json:"projectionType" tf:"projection_type"`

	// +kubebuilder:validation:Optional
	RangeKey *string `json:"rangeKey,omitempty" tf:"range_key"`

	// +kubebuilder:validation:Optional
	ReadCapacity *int64 `json:"readCapacity,omitempty" tf:"read_capacity"`

	// +kubebuilder:validation:Optional
	WriteCapacity *int64 `json:"writeCapacity,omitempty" tf:"write_capacity"`
}

type LocalSecondaryIndexObservation struct {
}

type LocalSecondaryIndexParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	NonKeyAttributes []string `json:"nonKeyAttributes,omitempty" tf:"non_key_attributes"`

	// +kubebuilder:validation:Required
	ProjectionType string `json:"projectionType" tf:"projection_type"`

	// +kubebuilder:validation:Required
	RangeKey string `json:"rangeKey" tf:"range_key"`
}

type PointInTimeRecoveryObservation struct {
}

type PointInTimeRecoveryParameters struct {

	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled" tf:"enabled"`
}

type ServerSideEncryptionObservation struct {
}

type ServerSideEncryptionParameters struct {

	// +kubebuilder:validation:Required
	Enabled bool `json:"enabled" tf:"enabled"`

	// +kubebuilder:validation:Optional
	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`
}

type TTLObservation struct {
}

type TTLParameters struct {

	// +kubebuilder:validation:Required
	AttributeName string `json:"attributeName" tf:"attribute_name"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	// +kubebuilder:validation:Optional
	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`
}

// DynamodbTableSpec defines the desired state of DynamodbTable
type DynamodbTableSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DynamodbTableParameters `json:"forProvider"`
}

// DynamodbTableStatus defines the observed state of DynamodbTable.
type DynamodbTableStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DynamodbTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DynamodbTable is the Schema for the DynamodbTables API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DynamodbTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DynamodbTableSpec   `json:"spec"`
	Status            DynamodbTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DynamodbTableList contains a list of DynamodbTables
type DynamodbTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DynamodbTable `json:"items"`
}

// Repository type metadata.
var (
	DynamodbTableKind             = "DynamodbTable"
	DynamodbTableGroupKind        = schema.GroupKind{Group: Group, Kind: DynamodbTableKind}.String()
	DynamodbTableKindAPIVersion   = DynamodbTableKind + "." + GroupVersion.String()
	DynamodbTableGroupVersionKind = GroupVersion.WithKind(DynamodbTableKind)
)

func init() {
	SchemeBuilder.Register(&DynamodbTable{}, &DynamodbTableList{})
}