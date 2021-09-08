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

type AppversionLifecycleObservation struct {
}

type AppversionLifecycleParameters struct {
	DeleteSourceFromS3 *bool `json:"deleteSourceFromS3,omitempty" tf:"delete_source_from_s3"`

	MaxAgeInDays *int64 `json:"maxAgeInDays,omitempty" tf:"max_age_in_days"`

	MaxCount *int64 `json:"maxCount,omitempty" tf:"max_count"`

	ServiceRole string `json:"serviceRole" tf:"service_role"`
}

type ElasticBeanstalkApplicationObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type ElasticBeanstalkApplicationParameters struct {
	AppversionLifecycle []AppversionLifecycleParameters `json:"appversionLifecycle,omitempty" tf:"appversion_lifecycle"`

	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// ElasticBeanstalkApplicationSpec defines the desired state of ElasticBeanstalkApplication
type ElasticBeanstalkApplicationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ElasticBeanstalkApplicationParameters `json:"forProvider"`
}

// ElasticBeanstalkApplicationStatus defines the observed state of ElasticBeanstalkApplication.
type ElasticBeanstalkApplicationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ElasticBeanstalkApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticBeanstalkApplication is the Schema for the ElasticBeanstalkApplications API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ElasticBeanstalkApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticBeanstalkApplicationSpec   `json:"spec"`
	Status            ElasticBeanstalkApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticBeanstalkApplicationList contains a list of ElasticBeanstalkApplications
type ElasticBeanstalkApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticBeanstalkApplication `json:"items"`
}

// Repository type metadata.
var (
	ElasticBeanstalkApplicationKind             = "ElasticBeanstalkApplication"
	ElasticBeanstalkApplicationGroupKind        = schema.GroupKind{Group: Group, Kind: ElasticBeanstalkApplicationKind}.String()
	ElasticBeanstalkApplicationKindAPIVersion   = ElasticBeanstalkApplicationKind + "." + GroupVersion.String()
	ElasticBeanstalkApplicationGroupVersionKind = GroupVersion.WithKind(ElasticBeanstalkApplicationKind)
)

func init() {
	SchemeBuilder.Register(&ElasticBeanstalkApplication{}, &ElasticBeanstalkApplicationList{})
}
