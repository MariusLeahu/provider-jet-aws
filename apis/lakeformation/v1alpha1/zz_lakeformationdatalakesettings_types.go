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

type CreateDatabaseDefaultPermissionsObservation struct {
}

type CreateDatabaseDefaultPermissionsParameters struct {

	// +kubebuilder:validation:Optional
	Permissions []string `json:"permissions,omitempty" tf:"permissions"`

	// +kubebuilder:validation:Optional
	Principal *string `json:"principal,omitempty" tf:"principal"`
}

type CreateTableDefaultPermissionsObservation struct {
}

type CreateTableDefaultPermissionsParameters struct {

	// +kubebuilder:validation:Optional
	Permissions []string `json:"permissions,omitempty" tf:"permissions"`

	// +kubebuilder:validation:Optional
	Principal *string `json:"principal,omitempty" tf:"principal"`
}

type LakeformationDataLakeSettingsObservation struct {
}

type LakeformationDataLakeSettingsParameters struct {

	// +kubebuilder:validation:Optional
	Admins []string `json:"admins,omitempty" tf:"admins"`

	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id"`

	// +kubebuilder:validation:Optional
	CreateDatabaseDefaultPermissions []CreateDatabaseDefaultPermissionsParameters `json:"createDatabaseDefaultPermissions,omitempty" tf:"create_database_default_permissions"`

	// +kubebuilder:validation:Optional
	CreateTableDefaultPermissions []CreateTableDefaultPermissionsParameters `json:"createTableDefaultPermissions,omitempty" tf:"create_table_default_permissions"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	TrustedResourceOwners []string `json:"trustedResourceOwners,omitempty" tf:"trusted_resource_owners"`
}

// LakeformationDataLakeSettingsSpec defines the desired state of LakeformationDataLakeSettings
type LakeformationDataLakeSettingsSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LakeformationDataLakeSettingsParameters `json:"forProvider"`
}

// LakeformationDataLakeSettingsStatus defines the observed state of LakeformationDataLakeSettings.
type LakeformationDataLakeSettingsStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LakeformationDataLakeSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LakeformationDataLakeSettings is the Schema for the LakeformationDataLakeSettingss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type LakeformationDataLakeSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LakeformationDataLakeSettingsSpec   `json:"spec"`
	Status            LakeformationDataLakeSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LakeformationDataLakeSettingsList contains a list of LakeformationDataLakeSettingss
type LakeformationDataLakeSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LakeformationDataLakeSettings `json:"items"`
}

// Repository type metadata.
var (
	LakeformationDataLakeSettingsKind             = "LakeformationDataLakeSettings"
	LakeformationDataLakeSettingsGroupKind        = schema.GroupKind{Group: Group, Kind: LakeformationDataLakeSettingsKind}.String()
	LakeformationDataLakeSettingsKindAPIVersion   = LakeformationDataLakeSettingsKind + "." + GroupVersion.String()
	LakeformationDataLakeSettingsGroupVersionKind = GroupVersion.WithKind(LakeformationDataLakeSettingsKind)
)

func init() {
	SchemeBuilder.Register(&LakeformationDataLakeSettings{}, &LakeformationDataLakeSettingsList{})
}