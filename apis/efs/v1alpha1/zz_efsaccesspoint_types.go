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

type CreationInfoObservation struct {
}

type CreationInfoParameters struct {

	// +kubebuilder:validation:Required
	OwnerGID int64 `json:"ownerGid" tf:"owner_gid"`

	// +kubebuilder:validation:Required
	OwnerUID int64 `json:"ownerUid" tf:"owner_uid"`

	// +kubebuilder:validation:Required
	Permissions string `json:"permissions" tf:"permissions"`
}

type EfsAccessPointObservation struct {
	Arn string `json:"arn" tf:"arn"`

	FileSystemArn string `json:"fileSystemArn" tf:"file_system_arn"`

	OwnerID string `json:"ownerId" tf:"owner_id"`
}

type EfsAccessPointParameters struct {

	// +kubebuilder:validation:Required
	FileSystemID string `json:"fileSystemId" tf:"file_system_id"`

	// +kubebuilder:validation:Optional
	PosixUser []PosixUserParameters `json:"posixUser,omitempty" tf:"posix_user"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RootDirectory []RootDirectoryParameters `json:"rootDirectory,omitempty" tf:"root_directory"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type PosixUserObservation struct {
}

type PosixUserParameters struct {

	// +kubebuilder:validation:Required
	GID int64 `json:"gid" tf:"gid"`

	// +kubebuilder:validation:Optional
	SecondaryGids []int64 `json:"secondaryGids,omitempty" tf:"secondary_gids"`

	// +kubebuilder:validation:Required
	UID int64 `json:"uid" tf:"uid"`
}

type RootDirectoryObservation struct {
}

type RootDirectoryParameters struct {

	// +kubebuilder:validation:Optional
	CreationInfo []CreationInfoParameters `json:"creationInfo,omitempty" tf:"creation_info"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path"`
}

// EfsAccessPointSpec defines the desired state of EfsAccessPoint
type EfsAccessPointSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EfsAccessPointParameters `json:"forProvider"`
}

// EfsAccessPointStatus defines the observed state of EfsAccessPoint.
type EfsAccessPointStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EfsAccessPointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EfsAccessPoint is the Schema for the EfsAccessPoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type EfsAccessPoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EfsAccessPointSpec   `json:"spec"`
	Status            EfsAccessPointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EfsAccessPointList contains a list of EfsAccessPoints
type EfsAccessPointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EfsAccessPoint `json:"items"`
}

// Repository type metadata.
var (
	EfsAccessPointKind             = "EfsAccessPoint"
	EfsAccessPointGroupKind        = schema.GroupKind{Group: Group, Kind: EfsAccessPointKind}.String()
	EfsAccessPointKindAPIVersion   = EfsAccessPointKind + "." + GroupVersion.String()
	EfsAccessPointGroupVersionKind = GroupVersion.WithKind(EfsAccessPointKind)
)

func init() {
	SchemeBuilder.Register(&EfsAccessPoint{}, &EfsAccessPointList{})
}