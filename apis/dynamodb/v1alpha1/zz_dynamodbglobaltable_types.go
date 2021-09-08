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

type DynamodbGlobalTableObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type DynamodbGlobalTableParameters struct {
	Name string `json:"name" tf:"name"`

	Replica []ReplicaParameters `json:"replica" tf:"replica"`
}

type ReplicaObservation struct {
}

type ReplicaParameters struct {
	RegionName string `json:"regionName" tf:"region_name"`
}

// DynamodbGlobalTableSpec defines the desired state of DynamodbGlobalTable
type DynamodbGlobalTableSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DynamodbGlobalTableParameters `json:"forProvider"`
}

// DynamodbGlobalTableStatus defines the observed state of DynamodbGlobalTable.
type DynamodbGlobalTableStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DynamodbGlobalTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DynamodbGlobalTable is the Schema for the DynamodbGlobalTables API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DynamodbGlobalTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DynamodbGlobalTableSpec   `json:"spec"`
	Status            DynamodbGlobalTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DynamodbGlobalTableList contains a list of DynamodbGlobalTables
type DynamodbGlobalTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DynamodbGlobalTable `json:"items"`
}

// Repository type metadata.
var (
	DynamodbGlobalTableKind             = "DynamodbGlobalTable"
	DynamodbGlobalTableGroupKind        = schema.GroupKind{Group: Group, Kind: DynamodbGlobalTableKind}.String()
	DynamodbGlobalTableKindAPIVersion   = DynamodbGlobalTableKind + "." + GroupVersion.String()
	DynamodbGlobalTableGroupVersionKind = GroupVersion.WithKind(DynamodbGlobalTableKind)
)

func init() {
	SchemeBuilder.Register(&DynamodbGlobalTable{}, &DynamodbGlobalTableList{})
}
