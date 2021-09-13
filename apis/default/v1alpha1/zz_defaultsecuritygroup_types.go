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

type DefaultSecurityGroupEgressObservation struct {
}

type DefaultSecurityGroupEgressParameters struct {

	// +kubebuilder:validation:Optional
	CidrBlocks []string `json:"cidrBlocks,omitempty" tf:"cidr_blocks"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Required
	FromPort int64 `json:"fromPort" tf:"from_port"`

	// +kubebuilder:validation:Optional
	IPv6CidrBlocks []string `json:"ipv6CidrBlocks,omitempty" tf:"ipv6_cidr_blocks"`

	// +kubebuilder:validation:Optional
	PrefixListIds []string `json:"prefixListIds,omitempty" tf:"prefix_list_ids"`

	// +kubebuilder:validation:Required
	Protocol string `json:"protocol" tf:"protocol"`

	// +kubebuilder:validation:Optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`

	// +kubebuilder:validation:Optional
	Self *bool `json:"self,omitempty" tf:"self"`

	// +kubebuilder:validation:Required
	ToPort int64 `json:"toPort" tf:"to_port"`
}

type DefaultSecurityGroupIngressObservation struct {
}

type DefaultSecurityGroupIngressParameters struct {

	// +kubebuilder:validation:Optional
	CidrBlocks []string `json:"cidrBlocks,omitempty" tf:"cidr_blocks"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Required
	FromPort int64 `json:"fromPort" tf:"from_port"`

	// +kubebuilder:validation:Optional
	IPv6CidrBlocks []string `json:"ipv6CidrBlocks,omitempty" tf:"ipv6_cidr_blocks"`

	// +kubebuilder:validation:Optional
	PrefixListIds []string `json:"prefixListIds,omitempty" tf:"prefix_list_ids"`

	// +kubebuilder:validation:Required
	Protocol string `json:"protocol" tf:"protocol"`

	// +kubebuilder:validation:Optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`

	// +kubebuilder:validation:Optional
	Self *bool `json:"self,omitempty" tf:"self"`

	// +kubebuilder:validation:Required
	ToPort int64 `json:"toPort" tf:"to_port"`
}

type DefaultSecurityGroupObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Description string `json:"description" tf:"description"`

	Name string `json:"name" tf:"name"`

	OwnerID string `json:"ownerId" tf:"owner_id"`
}

type DefaultSecurityGroupParameters struct {

	// +kubebuilder:validation:Optional
	Egress []DefaultSecurityGroupEgressParameters `json:"egress,omitempty" tf:"egress"`

	// +kubebuilder:validation:Optional
	Ingress []DefaultSecurityGroupIngressParameters `json:"ingress,omitempty" tf:"ingress"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RevokeRulesOnDelete *bool `json:"revokeRulesOnDelete,omitempty" tf:"revoke_rules_on_delete"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Optional
	VpcID *string `json:"vpcId,omitempty" tf:"vpc_id"`
}

// DefaultSecurityGroupSpec defines the desired state of DefaultSecurityGroup
type DefaultSecurityGroupSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DefaultSecurityGroupParameters `json:"forProvider"`
}

// DefaultSecurityGroupStatus defines the observed state of DefaultSecurityGroup.
type DefaultSecurityGroupStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DefaultSecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultSecurityGroup is the Schema for the DefaultSecurityGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DefaultSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultSecurityGroupSpec   `json:"spec"`
	Status            DefaultSecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultSecurityGroupList contains a list of DefaultSecurityGroups
type DefaultSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultSecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	DefaultSecurityGroupKind             = "DefaultSecurityGroup"
	DefaultSecurityGroupGroupKind        = schema.GroupKind{Group: Group, Kind: DefaultSecurityGroupKind}.String()
	DefaultSecurityGroupKindAPIVersion   = DefaultSecurityGroupKind + "." + GroupVersion.String()
	DefaultSecurityGroupGroupVersionKind = GroupVersion.WithKind(DefaultSecurityGroupKind)
)

func init() {
	SchemeBuilder.Register(&DefaultSecurityGroup{}, &DefaultSecurityGroupList{})
}