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

type PredicatesObservation struct {
}

type PredicatesParameters struct {
	DataId string `json:"dataId" tf:"data_id"`

	Negated bool `json:"negated" tf:"negated"`

	Type string `json:"type" tf:"type"`
}

type WafRateBasedRuleObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type WafRateBasedRuleParameters struct {
	MetricName string `json:"metricName" tf:"metric_name"`

	Name string `json:"name" tf:"name"`

	Predicates []PredicatesParameters `json:"predicates,omitempty" tf:"predicates"`

	RateKey string `json:"rateKey" tf:"rate_key"`

	RateLimit int64 `json:"rateLimit" tf:"rate_limit"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// WafRateBasedRuleSpec defines the desired state of WafRateBasedRule
type WafRateBasedRuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WafRateBasedRuleParameters `json:"forProvider"`
}

// WafRateBasedRuleStatus defines the observed state of WafRateBasedRule.
type WafRateBasedRuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WafRateBasedRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WafRateBasedRule is the Schema for the WafRateBasedRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type WafRateBasedRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafRateBasedRuleSpec   `json:"spec"`
	Status            WafRateBasedRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafRateBasedRuleList contains a list of WafRateBasedRules
type WafRateBasedRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafRateBasedRule `json:"items"`
}

// Repository type metadata.
var (
	WafRateBasedRuleKind             = "WafRateBasedRule"
	WafRateBasedRuleGroupKind        = schema.GroupKind{Group: Group, Kind: WafRateBasedRuleKind}.String()
	WafRateBasedRuleKindAPIVersion   = WafRateBasedRuleKind + "." + GroupVersion.String()
	WafRateBasedRuleGroupVersionKind = GroupVersion.WithKind(WafRateBasedRuleKind)
)

func init() {
	SchemeBuilder.Register(&WafRateBasedRule{}, &WafRateBasedRuleList{})
}
