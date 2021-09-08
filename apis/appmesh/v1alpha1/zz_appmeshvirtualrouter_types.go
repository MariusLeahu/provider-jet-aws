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

type AppmeshVirtualRouterObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreatedDate string `json:"createdDate" tf:"created_date"`

	LastUpdatedDate string `json:"lastUpdatedDate" tf:"last_updated_date"`

	ResourceOwner string `json:"resourceOwner" tf:"resource_owner"`
}

type AppmeshVirtualRouterParameters struct {
	MeshName string `json:"meshName" tf:"mesh_name"`

	MeshOwner *string `json:"meshOwner,omitempty" tf:"mesh_owner"`

	Name string `json:"name" tf:"name"`

	Spec []AppmeshVirtualRouterSpecParameters `json:"spec" tf:"spec"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type AppmeshVirtualRouterSpecListenerObservation struct {
}

type AppmeshVirtualRouterSpecListenerParameters struct {
	PortMapping []SpecListenerPortMappingParameters `json:"portMapping" tf:"port_mapping"`
}

type AppmeshVirtualRouterSpecObservation struct {
}

type AppmeshVirtualRouterSpecParameters struct {
	Listener []AppmeshVirtualRouterSpecListenerParameters `json:"listener" tf:"listener"`
}

type SpecListenerPortMappingObservation struct {
}

type SpecListenerPortMappingParameters struct {
	Port int64 `json:"port" tf:"port"`

	Protocol string `json:"protocol" tf:"protocol"`
}

// AppmeshVirtualRouterSpec defines the desired state of AppmeshVirtualRouter
type AppmeshVirtualRouterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AppmeshVirtualRouterParameters `json:"forProvider"`
}

// AppmeshVirtualRouterStatus defines the observed state of AppmeshVirtualRouter.
type AppmeshVirtualRouterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AppmeshVirtualRouterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppmeshVirtualRouter is the Schema for the AppmeshVirtualRouters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AppmeshVirtualRouter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppmeshVirtualRouterSpec   `json:"spec"`
	Status            AppmeshVirtualRouterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppmeshVirtualRouterList contains a list of AppmeshVirtualRouters
type AppmeshVirtualRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppmeshVirtualRouter `json:"items"`
}

// Repository type metadata.
var (
	AppmeshVirtualRouterKind             = "AppmeshVirtualRouter"
	AppmeshVirtualRouterGroupKind        = schema.GroupKind{Group: Group, Kind: AppmeshVirtualRouterKind}.String()
	AppmeshVirtualRouterKindAPIVersion   = AppmeshVirtualRouterKind + "." + GroupVersion.String()
	AppmeshVirtualRouterGroupVersionKind = GroupVersion.WithKind(AppmeshVirtualRouterKind)
)

func init() {
	SchemeBuilder.Register(&AppmeshVirtualRouter{}, &AppmeshVirtualRouterList{})
}
