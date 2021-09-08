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

type AlbListenerCertificateObservation struct {
}

type AlbListenerCertificateParameters struct {
	CertificateArn string `json:"certificateArn" tf:"certificate_arn"`

	ListenerArn string `json:"listenerArn" tf:"listener_arn"`
}

// AlbListenerCertificateSpec defines the desired state of AlbListenerCertificate
type AlbListenerCertificateSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AlbListenerCertificateParameters `json:"forProvider"`
}

// AlbListenerCertificateStatus defines the observed state of AlbListenerCertificate.
type AlbListenerCertificateStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AlbListenerCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AlbListenerCertificate is the Schema for the AlbListenerCertificates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AlbListenerCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlbListenerCertificateSpec   `json:"spec"`
	Status            AlbListenerCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlbListenerCertificateList contains a list of AlbListenerCertificates
type AlbListenerCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlbListenerCertificate `json:"items"`
}

// Repository type metadata.
var (
	AlbListenerCertificateKind             = "AlbListenerCertificate"
	AlbListenerCertificateGroupKind        = schema.GroupKind{Group: Group, Kind: AlbListenerCertificateKind}.String()
	AlbListenerCertificateKindAPIVersion   = AlbListenerCertificateKind + "." + GroupVersion.String()
	AlbListenerCertificateGroupVersionKind = GroupVersion.WithKind(AlbListenerCertificateKind)
)

func init() {
	SchemeBuilder.Register(&AlbListenerCertificate{}, &AlbListenerCertificateList{})
}
