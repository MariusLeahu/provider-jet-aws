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

type GuarddutyMemberObservation struct {
	RelationshipStatus string `json:"relationshipStatus" tf:"relationship_status"`
}

type GuarddutyMemberParameters struct {

	// +kubebuilder:validation:Required
	AccountID string `json:"accountId" tf:"account_id"`

	// +kubebuilder:validation:Required
	DetectorID string `json:"detectorId" tf:"detector_id"`

	// +kubebuilder:validation:Optional
	DisableEmailNotification *bool `json:"disableEmailNotification,omitempty" tf:"disable_email_notification"`

	// +kubebuilder:validation:Required
	Email string `json:"email" tf:"email"`

	// +kubebuilder:validation:Optional
	InvitationMessage *string `json:"invitationMessage,omitempty" tf:"invitation_message"`

	// +kubebuilder:validation:Optional
	Invite *bool `json:"invite,omitempty" tf:"invite"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`
}

// GuarddutyMemberSpec defines the desired state of GuarddutyMember
type GuarddutyMemberSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GuarddutyMemberParameters `json:"forProvider"`
}

// GuarddutyMemberStatus defines the observed state of GuarddutyMember.
type GuarddutyMemberStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GuarddutyMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GuarddutyMember is the Schema for the GuarddutyMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type GuarddutyMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GuarddutyMemberSpec   `json:"spec"`
	Status            GuarddutyMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GuarddutyMemberList contains a list of GuarddutyMembers
type GuarddutyMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GuarddutyMember `json:"items"`
}

// Repository type metadata.
var (
	GuarddutyMemberKind             = "GuarddutyMember"
	GuarddutyMemberGroupKind        = schema.GroupKind{Group: Group, Kind: GuarddutyMemberKind}.String()
	GuarddutyMemberKindAPIVersion   = GuarddutyMemberKind + "." + GroupVersion.String()
	GuarddutyMemberGroupVersionKind = GroupVersion.WithKind(GuarddutyMemberKind)
)

func init() {
	SchemeBuilder.Register(&GuarddutyMember{}, &GuarddutyMemberList{})
}