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

type VpcLinkObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VpcLinkParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	SecurityGroupIds []*string `json:"securityGroupIds" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Required
	SubnetIds []*string `json:"subnetIds" tf:"subnet_ids,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VpcLinkSpec defines the desired state of VpcLink
type VpcLinkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VpcLinkParameters `json:"forProvider"`
}

// VpcLinkStatus defines the observed state of VpcLink.
type VpcLinkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VpcLinkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VpcLink is the Schema for the VpcLinks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type VpcLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcLinkSpec   `json:"spec"`
	Status            VpcLinkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcLinkList contains a list of VpcLinks
type VpcLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpcLink `json:"items"`
}

// Repository type metadata.
var (
	VpcLink_Kind             = "VpcLink"
	VpcLink_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VpcLink_Kind}.String()
	VpcLink_KindAPIVersion   = VpcLink_Kind + "." + CRDGroupVersion.String()
	VpcLink_GroupVersionKind = CRDGroupVersion.WithKind(VpcLink_Kind)
)

func init() {
	SchemeBuilder.Register(&VpcLink{}, &VpcLinkList{})
}