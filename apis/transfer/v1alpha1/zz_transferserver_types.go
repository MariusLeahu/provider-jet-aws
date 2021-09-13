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

type EndpointDetailsObservation struct {
}

type EndpointDetailsParameters struct {

	// +kubebuilder:validation:Optional
	AddressAllocationIds []string `json:"addressAllocationIds,omitempty" tf:"address_allocation_ids"`

	// +kubebuilder:validation:Optional
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" tf:"security_group_ids"`

	// +kubebuilder:validation:Optional
	SubnetIds []string `json:"subnetIds,omitempty" tf:"subnet_ids"`

	// +kubebuilder:validation:Optional
	VpcEndpointID *string `json:"vpcEndpointId,omitempty" tf:"vpc_endpoint_id"`

	// +kubebuilder:validation:Optional
	VpcID *string `json:"vpcId,omitempty" tf:"vpc_id"`
}

type TransferServerObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Endpoint string `json:"endpoint" tf:"endpoint"`

	HostKeyFingerprint string `json:"hostKeyFingerprint" tf:"host_key_fingerprint"`
}

type TransferServerParameters struct {

	// +kubebuilder:validation:Optional
	Certificate *string `json:"certificate,omitempty" tf:"certificate"`

	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain"`

	// +kubebuilder:validation:Optional
	EndpointDetails []EndpointDetailsParameters `json:"endpointDetails,omitempty" tf:"endpoint_details"`

	// +kubebuilder:validation:Optional
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type"`

	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy"`

	// +kubebuilder:validation:Optional
	HostKey *string `json:"hostKey,omitempty" tf:"host_key"`

	// +kubebuilder:validation:Optional
	IdentityProviderType *string `json:"identityProviderType,omitempty" tf:"identity_provider_type"`

	// +kubebuilder:validation:Optional
	InvocationRole *string `json:"invocationRole,omitempty" tf:"invocation_role"`

	// +kubebuilder:validation:Optional
	LoggingRole *string `json:"loggingRole,omitempty" tf:"logging_role"`

	// +kubebuilder:validation:Optional
	Protocols []string `json:"protocols,omitempty" tf:"protocols"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityPolicyName *string `json:"securityPolicyName,omitempty" tf:"security_policy_name"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url"`
}

// TransferServerSpec defines the desired state of TransferServer
type TransferServerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       TransferServerParameters `json:"forProvider"`
}

// TransferServerStatus defines the observed state of TransferServer.
type TransferServerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          TransferServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransferServer is the Schema for the TransferServers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type TransferServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransferServerSpec   `json:"spec"`
	Status            TransferServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransferServerList contains a list of TransferServers
type TransferServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransferServer `json:"items"`
}

// Repository type metadata.
var (
	TransferServerKind             = "TransferServer"
	TransferServerGroupKind        = schema.GroupKind{Group: Group, Kind: TransferServerKind}.String()
	TransferServerKindAPIVersion   = TransferServerKind + "." + GroupVersion.String()
	TransferServerGroupVersionKind = GroupVersion.WithKind(TransferServerKind)
)

func init() {
	SchemeBuilder.Register(&TransferServer{}, &TransferServerList{})
}