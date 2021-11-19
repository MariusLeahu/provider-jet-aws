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

type V2HsmObservation struct {
	HsmEniID *string `json:"hsmEniId,omitempty" tf:"hsm_eni_id,omitempty"`

	HsmID *string `json:"hsmId,omitempty" tf:"hsm_id,omitempty"`

	HsmState *string `json:"hsmState,omitempty" tf:"hsm_state,omitempty"`
}

type V2HsmParameters struct {

	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

// V2HsmSpec defines the desired state of V2Hsm
type V2HsmSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     V2HsmParameters `json:"forProvider"`
}

// V2HsmStatus defines the observed state of V2Hsm.
type V2HsmStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        V2HsmObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// V2Hsm is the Schema for the V2Hsms API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type V2Hsm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              V2HsmSpec   `json:"spec"`
	Status            V2HsmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// V2HsmList contains a list of V2Hsms
type V2HsmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []V2Hsm `json:"items"`
}

// Repository type metadata.
var (
	V2Hsm_Kind             = "V2Hsm"
	V2Hsm_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: V2Hsm_Kind}.String()
	V2Hsm_KindAPIVersion   = V2Hsm_Kind + "." + CRDGroupVersion.String()
	V2Hsm_GroupVersionKind = CRDGroupVersion.WithKind(V2Hsm_Kind)
)

func init() {
	SchemeBuilder.Register(&V2Hsm{}, &V2HsmList{})
}