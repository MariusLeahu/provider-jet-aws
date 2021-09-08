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

type CloudwatchEventPermissionObservation struct {
}

type CloudwatchEventPermissionParameters struct {
	Action *string `json:"action,omitempty" tf:"action"`

	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition"`

	EventBusName *string `json:"eventBusName,omitempty" tf:"event_bus_name"`

	Principal string `json:"principal" tf:"principal"`

	StatementId string `json:"statementId" tf:"statement_id"`
}

type ConditionObservation struct {
}

type ConditionParameters struct {
	Key string `json:"key" tf:"key"`

	Type string `json:"type" tf:"type"`

	Value string `json:"value" tf:"value"`
}

// CloudwatchEventPermissionSpec defines the desired state of CloudwatchEventPermission
type CloudwatchEventPermissionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudwatchEventPermissionParameters `json:"forProvider"`
}

// CloudwatchEventPermissionStatus defines the observed state of CloudwatchEventPermission.
type CloudwatchEventPermissionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudwatchEventPermissionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventPermission is the Schema for the CloudwatchEventPermissions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CloudwatchEventPermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchEventPermissionSpec   `json:"spec"`
	Status            CloudwatchEventPermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventPermissionList contains a list of CloudwatchEventPermissions
type CloudwatchEventPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchEventPermission `json:"items"`
}

// Repository type metadata.
var (
	CloudwatchEventPermissionKind             = "CloudwatchEventPermission"
	CloudwatchEventPermissionGroupKind        = schema.GroupKind{Group: Group, Kind: CloudwatchEventPermissionKind}.String()
	CloudwatchEventPermissionKindAPIVersion   = CloudwatchEventPermissionKind + "." + GroupVersion.String()
	CloudwatchEventPermissionGroupVersionKind = GroupVersion.WithKind(CloudwatchEventPermissionKind)
)

func init() {
	SchemeBuilder.Register(&CloudwatchEventPermission{}, &CloudwatchEventPermissionList{})
}
