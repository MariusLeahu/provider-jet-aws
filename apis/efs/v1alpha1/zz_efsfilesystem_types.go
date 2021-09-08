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

type EfsFileSystemObservation struct {
	Arn string `json:"arn" tf:"arn"`

	AvailabilityZoneId string `json:"availabilityZoneId" tf:"availability_zone_id"`

	DnsName string `json:"dnsName" tf:"dns_name"`

	NumberOfMountTargets int64 `json:"numberOfMountTargets" tf:"number_of_mount_targets"`

	OwnerId string `json:"ownerId" tf:"owner_id"`

	SizeInBytes []SizeInBytesObservation `json:"sizeInBytes" tf:"size_in_bytes"`
}

type EfsFileSystemParameters struct {
	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty" tf:"availability_zone_name"`

	CreationToken *string `json:"creationToken,omitempty" tf:"creation_token"`

	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted"`

	KmsKeyId *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	LifecyclePolicy []LifecyclePolicyParameters `json:"lifecyclePolicy,omitempty" tf:"lifecycle_policy"`

	PerformanceMode *string `json:"performanceMode,omitempty" tf:"performance_mode"`

	ProvisionedThroughputInMibps *float64 `json:"provisionedThroughputInMibps,omitempty" tf:"provisioned_throughput_in_mibps"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	ThroughputMode *string `json:"throughputMode,omitempty" tf:"throughput_mode"`
}

type LifecyclePolicyObservation struct {
}

type LifecyclePolicyParameters struct {
	TransitionToIa string `json:"transitionToIa" tf:"transition_to_ia"`
}

type SizeInBytesObservation struct {
	Value int64 `json:"value" tf:"value"`

	ValueInIa int64 `json:"valueInIa" tf:"value_in_ia"`

	ValueInStandard int64 `json:"valueInStandard" tf:"value_in_standard"`
}

type SizeInBytesParameters struct {
}

// EfsFileSystemSpec defines the desired state of EfsFileSystem
type EfsFileSystemSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EfsFileSystemParameters `json:"forProvider"`
}

// EfsFileSystemStatus defines the observed state of EfsFileSystem.
type EfsFileSystemStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EfsFileSystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EfsFileSystem is the Schema for the EfsFileSystems API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type EfsFileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EfsFileSystemSpec   `json:"spec"`
	Status            EfsFileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EfsFileSystemList contains a list of EfsFileSystems
type EfsFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EfsFileSystem `json:"items"`
}

// Repository type metadata.
var (
	EfsFileSystemKind             = "EfsFileSystem"
	EfsFileSystemGroupKind        = schema.GroupKind{Group: Group, Kind: EfsFileSystemKind}.String()
	EfsFileSystemKindAPIVersion   = EfsFileSystemKind + "." + GroupVersion.String()
	EfsFileSystemGroupVersionKind = GroupVersion.WithKind(EfsFileSystemKind)
)

func init() {
	SchemeBuilder.Register(&EfsFileSystem{}, &EfsFileSystemList{})
}
