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

type ContainerObservation struct {
}

type ContainerParameters struct {

	// +kubebuilder:validation:Optional
	ContainerHostname *string `json:"containerHostname,omitempty" tf:"container_hostname,omitempty"`

	// +kubebuilder:validation:Optional
	Environment map[string]*string `json:"environment,omitempty" tf:"environment,omitempty"`

	// +kubebuilder:validation:Required
	Image *string `json:"image" tf:"image,omitempty"`

	// +kubebuilder:validation:Optional
	ImageConfig []ImageConfigParameters `json:"imageConfig,omitempty" tf:"image_config,omitempty"`

	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// +kubebuilder:validation:Optional
	ModelDataURL *string `json:"modelDataUrl,omitempty" tf:"model_data_url,omitempty"`
}

type ImageConfigObservation struct {
}

type ImageConfigParameters struct {

	// +kubebuilder:validation:Required
	RepositoryAccessMode *string `json:"repositoryAccessMode" tf:"repository_access_mode,omitempty"`
}

type InferenceExecutionConfigObservation struct {
}

type InferenceExecutionConfigParameters struct {

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`
}

type ModelObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ModelParameters struct {

	// +kubebuilder:validation:Optional
	Container []ContainerParameters `json:"container,omitempty" tf:"container,omitempty"`

	// +kubebuilder:validation:Optional
	EnableNetworkIsolation *bool `json:"enableNetworkIsolation,omitempty" tf:"enable_network_isolation,omitempty"`

	// +kubebuilder:validation:Required
	ExecutionRoleArn *string `json:"executionRoleArn" tf:"execution_role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	InferenceExecutionConfig []InferenceExecutionConfigParameters `json:"inferenceExecutionConfig,omitempty" tf:"inference_execution_config,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PrimaryContainer []PrimaryContainerParameters `json:"primaryContainer,omitempty" tf:"primary_container,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VpcConfig []VpcConfigParameters `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type PrimaryContainerImageConfigObservation struct {
}

type PrimaryContainerImageConfigParameters struct {

	// +kubebuilder:validation:Required
	RepositoryAccessMode *string `json:"repositoryAccessMode" tf:"repository_access_mode,omitempty"`
}

type PrimaryContainerObservation struct {
}

type PrimaryContainerParameters struct {

	// +kubebuilder:validation:Optional
	ContainerHostname *string `json:"containerHostname,omitempty" tf:"container_hostname,omitempty"`

	// +kubebuilder:validation:Optional
	Environment map[string]*string `json:"environment,omitempty" tf:"environment,omitempty"`

	// +kubebuilder:validation:Required
	Image *string `json:"image" tf:"image,omitempty"`

	// +kubebuilder:validation:Optional
	ImageConfig []PrimaryContainerImageConfigParameters `json:"imageConfig,omitempty" tf:"image_config,omitempty"`

	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// +kubebuilder:validation:Optional
	ModelDataURL *string `json:"modelDataUrl,omitempty" tf:"model_data_url,omitempty"`
}

type VpcConfigObservation struct {
}

type VpcConfigParameters struct {

	// +kubebuilder:validation:Required
	SecurityGroupIds []*string `json:"securityGroupIds" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Required
	Subnets []*string `json:"subnets" tf:"subnets,omitempty"`
}

// ModelSpec defines the desired state of Model
type ModelSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ModelParameters `json:"forProvider"`
}

// ModelStatus defines the observed state of Model.
type ModelStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ModelObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Model is the Schema for the Models API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Model struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ModelSpec   `json:"spec"`
	Status            ModelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ModelList contains a list of Models
type ModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Model `json:"items"`
}

// Repository type metadata.
var (
	Model_Kind             = "Model"
	Model_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Model_Kind}.String()
	Model_KindAPIVersion   = Model_Kind + "." + CRDGroupVersion.String()
	Model_GroupVersionKind = CRDGroupVersion.WithKind(Model_Kind)
)

func init() {
	SchemeBuilder.Register(&Model{}, &ModelList{})
}