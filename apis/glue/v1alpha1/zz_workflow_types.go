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

type WorkflowObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type WorkflowParameters struct {

	// +kubebuilder:validation:Optional
	DefaultRunProperties map[string]string `json:"defaultRunProperties,omitempty" tf:"default_run_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	MaxConcurrentRuns *int64 `json:"maxConcurrentRuns,omitempty" tf:"max_concurrent_runs,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WorkflowSpec defines the desired state of Workflow
type WorkflowSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkflowParameters `json:"forProvider"`
}

// WorkflowStatus defines the observed state of Workflow.
type WorkflowStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkflowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Workflow is the Schema for the Workflows API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Workflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkflowSpec   `json:"spec"`
	Status            WorkflowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkflowList contains a list of Workflows
type WorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workflow `json:"items"`
}

// Repository type metadata.
var (
	Workflow_Kind             = "Workflow"
	Workflow_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Workflow_Kind}.String()
	Workflow_KindAPIVersion   = Workflow_Kind + "." + CRDGroupVersion.String()
	Workflow_GroupVersionKind = CRDGroupVersion.WithKind(Workflow_Kind)
)

func init() {
	SchemeBuilder.Register(&Workflow{}, &WorkflowList{})
}