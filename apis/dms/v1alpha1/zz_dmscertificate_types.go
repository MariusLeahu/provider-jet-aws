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

type DmsCertificateObservation struct {
	CertificateArn string `json:"certificateArn" tf:"certificate_arn"`
}

type DmsCertificateParameters struct {
	CertificateId string `json:"certificateId" tf:"certificate_id"`

	CertificatePem *string `json:"certificatePem,omitempty" tf:"certificate_pem"`

	CertificateWallet *string `json:"certificateWallet,omitempty" tf:"certificate_wallet"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// DmsCertificateSpec defines the desired state of DmsCertificate
type DmsCertificateSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DmsCertificateParameters `json:"forProvider"`
}

// DmsCertificateStatus defines the observed state of DmsCertificate.
type DmsCertificateStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DmsCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DmsCertificate is the Schema for the DmsCertificates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DmsCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DmsCertificateSpec   `json:"spec"`
	Status            DmsCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DmsCertificateList contains a list of DmsCertificates
type DmsCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DmsCertificate `json:"items"`
}

// Repository type metadata.
var (
	DmsCertificateKind             = "DmsCertificate"
	DmsCertificateGroupKind        = schema.GroupKind{Group: Group, Kind: DmsCertificateKind}.String()
	DmsCertificateKindAPIVersion   = DmsCertificateKind + "." + GroupVersion.String()
	DmsCertificateGroupVersionKind = GroupVersion.WithKind(DmsCertificateKind)
)

func init() {
	SchemeBuilder.Register(&DmsCertificate{}, &DmsCertificateList{})
}
