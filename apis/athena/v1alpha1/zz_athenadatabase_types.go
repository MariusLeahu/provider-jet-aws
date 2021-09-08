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

type AthenaDatabaseObservation struct {
}

type AthenaDatabaseParameters struct {
	Bucket string `json:"bucket" tf:"bucket"`

	EncryptionConfiguration []EncryptionConfigurationParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration"`

	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy"`

	Name string `json:"name" tf:"name"`
}

type EncryptionConfigurationObservation struct {
}

type EncryptionConfigurationParameters struct {
	EncryptionOption string `json:"encryptionOption" tf:"encryption_option"`

	KmsKey *string `json:"kmsKey,omitempty" tf:"kms_key"`
}

// AthenaDatabaseSpec defines the desired state of AthenaDatabase
type AthenaDatabaseSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AthenaDatabaseParameters `json:"forProvider"`
}

// AthenaDatabaseStatus defines the observed state of AthenaDatabase.
type AthenaDatabaseStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AthenaDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AthenaDatabase is the Schema for the AthenaDatabases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AthenaDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AthenaDatabaseSpec   `json:"spec"`
	Status            AthenaDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AthenaDatabaseList contains a list of AthenaDatabases
type AthenaDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AthenaDatabase `json:"items"`
}

// Repository type metadata.
var (
	AthenaDatabaseKind             = "AthenaDatabase"
	AthenaDatabaseGroupKind        = schema.GroupKind{Group: Group, Kind: AthenaDatabaseKind}.String()
	AthenaDatabaseKindAPIVersion   = AthenaDatabaseKind + "." + GroupVersion.String()
	AthenaDatabaseGroupVersionKind = GroupVersion.WithKind(AthenaDatabaseKind)
)

func init() {
	SchemeBuilder.Register(&AthenaDatabase{}, &AthenaDatabaseList{})
}
