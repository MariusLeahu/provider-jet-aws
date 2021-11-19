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

type BeanstalkApplicationVersionObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type BeanstalkApplicationVersionParameters struct {

	// +kubebuilder:validation:Required
	Application *string `json:"application" tf:"application,omitempty"`

	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// BeanstalkApplicationVersionSpec defines the desired state of BeanstalkApplicationVersion
type BeanstalkApplicationVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BeanstalkApplicationVersionParameters `json:"forProvider"`
}

// BeanstalkApplicationVersionStatus defines the observed state of BeanstalkApplicationVersion.
type BeanstalkApplicationVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BeanstalkApplicationVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BeanstalkApplicationVersion is the Schema for the BeanstalkApplicationVersions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type BeanstalkApplicationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BeanstalkApplicationVersionSpec   `json:"spec"`
	Status            BeanstalkApplicationVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BeanstalkApplicationVersionList contains a list of BeanstalkApplicationVersions
type BeanstalkApplicationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BeanstalkApplicationVersion `json:"items"`
}

// Repository type metadata.
var (
	BeanstalkApplicationVersion_Kind             = "BeanstalkApplicationVersion"
	BeanstalkApplicationVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BeanstalkApplicationVersion_Kind}.String()
	BeanstalkApplicationVersion_KindAPIVersion   = BeanstalkApplicationVersion_Kind + "." + CRDGroupVersion.String()
	BeanstalkApplicationVersion_GroupVersionKind = CRDGroupVersion.WithKind(BeanstalkApplicationVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&BeanstalkApplicationVersion{}, &BeanstalkApplicationVersionList{})
}