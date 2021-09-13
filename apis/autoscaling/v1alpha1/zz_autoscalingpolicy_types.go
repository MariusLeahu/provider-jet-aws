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

type AutoscalingPolicyObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type AutoscalingPolicyParameters struct {

	// +kubebuilder:validation:Optional
	AdjustmentType *string `json:"adjustmentType,omitempty" tf:"adjustment_type"`

	// +kubebuilder:validation:Required
	AutoscalingGroupName string `json:"autoscalingGroupName" tf:"autoscaling_group_name"`

	// +kubebuilder:validation:Optional
	Cooldown *int64 `json:"cooldown,omitempty" tf:"cooldown"`

	// +kubebuilder:validation:Optional
	EstimatedInstanceWarmup *int64 `json:"estimatedInstanceWarmup,omitempty" tf:"estimated_instance_warmup"`

	// +kubebuilder:validation:Optional
	MetricAggregationType *string `json:"metricAggregationType,omitempty" tf:"metric_aggregation_type"`

	// +kubebuilder:validation:Optional
	MinAdjustmentMagnitude *int64 `json:"minAdjustmentMagnitude,omitempty" tf:"min_adjustment_magnitude"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	PolicyType *string `json:"policyType,omitempty" tf:"policy_type"`

	// +kubebuilder:validation:Optional
	PredictiveScalingConfiguration []PredictiveScalingConfigurationParameters `json:"predictiveScalingConfiguration,omitempty" tf:"predictive_scaling_configuration"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ScalingAdjustment *int64 `json:"scalingAdjustment,omitempty" tf:"scaling_adjustment"`

	// +kubebuilder:validation:Optional
	StepAdjustment []StepAdjustmentParameters `json:"stepAdjustment,omitempty" tf:"step_adjustment"`

	// +kubebuilder:validation:Optional
	TargetTrackingConfiguration []TargetTrackingConfigurationParameters `json:"targetTrackingConfiguration,omitempty" tf:"target_tracking_configuration"`
}

type CustomizedMetricSpecificationObservation struct {
}

type CustomizedMetricSpecificationParameters struct {

	// +kubebuilder:validation:Optional
	MetricDimension []MetricDimensionParameters `json:"metricDimension,omitempty" tf:"metric_dimension"`

	// +kubebuilder:validation:Required
	MetricName string `json:"metricName" tf:"metric_name"`

	// +kubebuilder:validation:Required
	Namespace string `json:"namespace" tf:"namespace"`

	// +kubebuilder:validation:Required
	Statistic string `json:"statistic" tf:"statistic"`

	// +kubebuilder:validation:Optional
	Unit *string `json:"unit,omitempty" tf:"unit"`
}

type MetricDimensionObservation struct {
}

type MetricDimensionParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	Value string `json:"value" tf:"value"`
}

type MetricSpecificationObservation struct {
}

type MetricSpecificationParameters struct {

	// +kubebuilder:validation:Optional
	PredefinedLoadMetricSpecification []PredefinedLoadMetricSpecificationParameters `json:"predefinedLoadMetricSpecification,omitempty" tf:"predefined_load_metric_specification"`

	// +kubebuilder:validation:Optional
	PredefinedMetricPairSpecification []PredefinedMetricPairSpecificationParameters `json:"predefinedMetricPairSpecification,omitempty" tf:"predefined_metric_pair_specification"`

	// +kubebuilder:validation:Optional
	PredefinedScalingMetricSpecification []PredefinedScalingMetricSpecificationParameters `json:"predefinedScalingMetricSpecification,omitempty" tf:"predefined_scaling_metric_specification"`

	// +kubebuilder:validation:Required
	TargetValue int64 `json:"targetValue" tf:"target_value"`
}

type PredefinedLoadMetricSpecificationObservation struct {
}

type PredefinedLoadMetricSpecificationParameters struct {

	// +kubebuilder:validation:Required
	PredefinedMetricType string `json:"predefinedMetricType" tf:"predefined_metric_type"`

	// +kubebuilder:validation:Required
	ResourceLabel string `json:"resourceLabel" tf:"resource_label"`
}

type PredefinedMetricPairSpecificationObservation struct {
}

type PredefinedMetricPairSpecificationParameters struct {

	// +kubebuilder:validation:Required
	PredefinedMetricType string `json:"predefinedMetricType" tf:"predefined_metric_type"`

	// +kubebuilder:validation:Required
	ResourceLabel string `json:"resourceLabel" tf:"resource_label"`
}

type PredefinedMetricSpecificationObservation struct {
}

type PredefinedMetricSpecificationParameters struct {

	// +kubebuilder:validation:Required
	PredefinedMetricType string `json:"predefinedMetricType" tf:"predefined_metric_type"`

	// +kubebuilder:validation:Optional
	ResourceLabel *string `json:"resourceLabel,omitempty" tf:"resource_label"`
}

type PredefinedScalingMetricSpecificationObservation struct {
}

type PredefinedScalingMetricSpecificationParameters struct {

	// +kubebuilder:validation:Required
	PredefinedMetricType string `json:"predefinedMetricType" tf:"predefined_metric_type"`

	// +kubebuilder:validation:Required
	ResourceLabel string `json:"resourceLabel" tf:"resource_label"`
}

type PredictiveScalingConfigurationObservation struct {
}

type PredictiveScalingConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	MaxCapacityBreachBehavior *string `json:"maxCapacityBreachBehavior,omitempty" tf:"max_capacity_breach_behavior"`

	// +kubebuilder:validation:Optional
	MaxCapacityBuffer *string `json:"maxCapacityBuffer,omitempty" tf:"max_capacity_buffer"`

	// +kubebuilder:validation:Required
	MetricSpecification []MetricSpecificationParameters `json:"metricSpecification" tf:"metric_specification"`

	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode"`

	// +kubebuilder:validation:Optional
	SchedulingBufferTime *string `json:"schedulingBufferTime,omitempty" tf:"scheduling_buffer_time"`
}

type StepAdjustmentObservation struct {
}

type StepAdjustmentParameters struct {

	// +kubebuilder:validation:Optional
	MetricIntervalLowerBound *string `json:"metricIntervalLowerBound,omitempty" tf:"metric_interval_lower_bound"`

	// +kubebuilder:validation:Optional
	MetricIntervalUpperBound *string `json:"metricIntervalUpperBound,omitempty" tf:"metric_interval_upper_bound"`

	// +kubebuilder:validation:Required
	ScalingAdjustment int64 `json:"scalingAdjustment" tf:"scaling_adjustment"`
}

type TargetTrackingConfigurationObservation struct {
}

type TargetTrackingConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	CustomizedMetricSpecification []CustomizedMetricSpecificationParameters `json:"customizedMetricSpecification,omitempty" tf:"customized_metric_specification"`

	// +kubebuilder:validation:Optional
	DisableScaleIn *bool `json:"disableScaleIn,omitempty" tf:"disable_scale_in"`

	// +kubebuilder:validation:Optional
	PredefinedMetricSpecification []PredefinedMetricSpecificationParameters `json:"predefinedMetricSpecification,omitempty" tf:"predefined_metric_specification"`

	// +kubebuilder:validation:Required
	TargetValue float64 `json:"targetValue" tf:"target_value"`
}

// AutoscalingPolicySpec defines the desired state of AutoscalingPolicy
type AutoscalingPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AutoscalingPolicyParameters `json:"forProvider"`
}

// AutoscalingPolicyStatus defines the observed state of AutoscalingPolicy.
type AutoscalingPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AutoscalingPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingPolicy is the Schema for the AutoscalingPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AutoscalingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingPolicySpec   `json:"spec"`
	Status            AutoscalingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingPolicyList contains a list of AutoscalingPolicys
type AutoscalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoscalingPolicy `json:"items"`
}

// Repository type metadata.
var (
	AutoscalingPolicyKind             = "AutoscalingPolicy"
	AutoscalingPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: AutoscalingPolicyKind}.String()
	AutoscalingPolicyKindAPIVersion   = AutoscalingPolicyKind + "." + GroupVersion.String()
	AutoscalingPolicyGroupVersionKind = GroupVersion.WithKind(AutoscalingPolicyKind)
)

func init() {
	SchemeBuilder.Register(&AutoscalingPolicy{}, &AutoscalingPolicyList{})
}