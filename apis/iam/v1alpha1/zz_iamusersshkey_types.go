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

type IamUserSshKeyObservation struct {
	Fingerprint string `json:"fingerprint" tf:"fingerprint"`

	SshPublicKeyId string `json:"sshPublicKeyId" tf:"ssh_public_key_id"`
}

type IamUserSshKeyParameters struct {
	Encoding string `json:"encoding" tf:"encoding"`

	PublicKey string `json:"publicKey" tf:"public_key"`

	Status *string `json:"status,omitempty" tf:"status"`

	Username string `json:"username" tf:"username"`
}

// IamUserSshKeySpec defines the desired state of IamUserSshKey
type IamUserSshKeySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IamUserSshKeyParameters `json:"forProvider"`
}

// IamUserSshKeyStatus defines the observed state of IamUserSshKey.
type IamUserSshKeyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IamUserSshKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IamUserSshKey is the Schema for the IamUserSshKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type IamUserSshKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamUserSshKeySpec   `json:"spec"`
	Status            IamUserSshKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamUserSshKeyList contains a list of IamUserSshKeys
type IamUserSshKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamUserSshKey `json:"items"`
}

// Repository type metadata.
var (
	IamUserSshKeyKind             = "IamUserSshKey"
	IamUserSshKeyGroupKind        = schema.GroupKind{Group: Group, Kind: IamUserSshKeyKind}.String()
	IamUserSshKeyKindAPIVersion   = IamUserSshKeyKind + "." + GroupVersion.String()
	IamUserSshKeyGroupVersionKind = GroupVersion.WithKind(IamUserSshKeyKind)
)

func init() {
	SchemeBuilder.Register(&IamUserSshKey{}, &IamUserSshKeyList{})
}
