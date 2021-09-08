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

type IpSetDescriptorsObservation struct {
}

type IpSetDescriptorsParameters struct {
	Type string `json:"type" tf:"type"`

	Value string `json:"value" tf:"value"`
}

type WafIpsetObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type WafIpsetParameters struct {
	IpSetDescriptors []IpSetDescriptorsParameters `json:"ipSetDescriptors,omitempty" tf:"ip_set_descriptors"`

	Name string `json:"name" tf:"name"`
}

// WafIpsetSpec defines the desired state of WafIpset
type WafIpsetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WafIpsetParameters `json:"forProvider"`
}

// WafIpsetStatus defines the observed state of WafIpset.
type WafIpsetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WafIpsetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WafIpset is the Schema for the WafIpsets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type WafIpset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafIpsetSpec   `json:"spec"`
	Status            WafIpsetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafIpsetList contains a list of WafIpsets
type WafIpsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafIpset `json:"items"`
}

// Repository type metadata.
var (
	WafIpsetKind             = "WafIpset"
	WafIpsetGroupKind        = schema.GroupKind{Group: Group, Kind: WafIpsetKind}.String()
	WafIpsetKindAPIVersion   = WafIpsetKind + "." + GroupVersion.String()
	WafIpsetGroupVersionKind = GroupVersion.WithKind(WafIpsetKind)
)

func init() {
	SchemeBuilder.Register(&WafIpset{}, &WafIpsetList{})
}
