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

type AutomationParametersObservation struct {
}

type AutomationParametersParameterObservation struct {
}

type AutomationParametersParameterParameters struct {
	Name string `json:"name" tf:"name"`

	Values []string `json:"values" tf:"values"`
}

type AutomationParametersParameters struct {
	DocumentVersion *string `json:"documentVersion,omitempty" tf:"document_version"`

	Parameter []AutomationParametersParameterParameters `json:"parameter,omitempty" tf:"parameter"`
}

type CloudwatchConfigObservation struct {
}

type CloudwatchConfigParameters struct {
	CloudwatchLogGroupName *string `json:"cloudwatchLogGroupName,omitempty" tf:"cloudwatch_log_group_name"`

	CloudwatchOutputEnabled *bool `json:"cloudwatchOutputEnabled,omitempty" tf:"cloudwatch_output_enabled"`
}

type LambdaParametersObservation struct {
}

type LambdaParametersParameters struct {
	ClientContext *string `json:"clientContext,omitempty" tf:"client_context"`

	Payload *string `json:"payload,omitempty" tf:"payload"`

	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier"`
}

type NotificationConfigObservation struct {
}

type NotificationConfigParameters struct {
	NotificationArn *string `json:"notificationArn,omitempty" tf:"notification_arn"`

	NotificationEvents []string `json:"notificationEvents,omitempty" tf:"notification_events"`

	NotificationType *string `json:"notificationType,omitempty" tf:"notification_type"`
}

type RunCommandParametersObservation struct {
}

type RunCommandParametersParameterObservation struct {
}

type RunCommandParametersParameterParameters struct {
	Name string `json:"name" tf:"name"`

	Values []string `json:"values" tf:"values"`
}

type RunCommandParametersParameters struct {
	CloudwatchConfig []CloudwatchConfigParameters `json:"cloudwatchConfig,omitempty" tf:"cloudwatch_config"`

	Comment *string `json:"comment,omitempty" tf:"comment"`

	DocumentHash *string `json:"documentHash,omitempty" tf:"document_hash"`

	DocumentHashType *string `json:"documentHashType,omitempty" tf:"document_hash_type"`

	DocumentVersion *string `json:"documentVersion,omitempty" tf:"document_version"`

	NotificationConfig []NotificationConfigParameters `json:"notificationConfig,omitempty" tf:"notification_config"`

	OutputS3Bucket *string `json:"outputS3Bucket,omitempty" tf:"output_s3_bucket"`

	OutputS3KeyPrefix *string `json:"outputS3KeyPrefix,omitempty" tf:"output_s3_key_prefix"`

	Parameter []RunCommandParametersParameterParameters `json:"parameter,omitempty" tf:"parameter"`

	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn"`

	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty" tf:"timeout_seconds"`
}

type SsmMaintenanceWindowTaskObservation struct {
}

type SsmMaintenanceWindowTaskParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	MaxConcurrency string `json:"maxConcurrency" tf:"max_concurrency"`

	MaxErrors string `json:"maxErrors" tf:"max_errors"`

	Name *string `json:"name,omitempty" tf:"name"`

	Priority *int64 `json:"priority,omitempty" tf:"priority"`

	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn"`

	Targets []SsmMaintenanceWindowTaskTargetsParameters `json:"targets,omitempty" tf:"targets"`

	TaskArn string `json:"taskArn" tf:"task_arn"`

	TaskInvocationParameters []TaskInvocationParametersParameters `json:"taskInvocationParameters,omitempty" tf:"task_invocation_parameters"`

	TaskType string `json:"taskType" tf:"task_type"`

	WindowId string `json:"windowId" tf:"window_id"`
}

type SsmMaintenanceWindowTaskTargetsObservation struct {
}

type SsmMaintenanceWindowTaskTargetsParameters struct {
	Key string `json:"key" tf:"key"`

	Values []string `json:"values" tf:"values"`
}

type StepFunctionsParametersObservation struct {
}

type StepFunctionsParametersParameters struct {
	Input *string `json:"input,omitempty" tf:"input"`

	Name *string `json:"name,omitempty" tf:"name"`
}

type TaskInvocationParametersObservation struct {
}

type TaskInvocationParametersParameters struct {
	AutomationParameters []AutomationParametersParameters `json:"automationParameters,omitempty" tf:"automation_parameters"`

	LambdaParameters []LambdaParametersParameters `json:"lambdaParameters,omitempty" tf:"lambda_parameters"`

	RunCommandParameters []RunCommandParametersParameters `json:"runCommandParameters,omitempty" tf:"run_command_parameters"`

	StepFunctionsParameters []StepFunctionsParametersParameters `json:"stepFunctionsParameters,omitempty" tf:"step_functions_parameters"`
}

// SsmMaintenanceWindowTaskSpec defines the desired state of SsmMaintenanceWindowTask
type SsmMaintenanceWindowTaskSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SsmMaintenanceWindowTaskParameters `json:"forProvider"`
}

// SsmMaintenanceWindowTaskStatus defines the observed state of SsmMaintenanceWindowTask.
type SsmMaintenanceWindowTaskStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SsmMaintenanceWindowTaskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SsmMaintenanceWindowTask is the Schema for the SsmMaintenanceWindowTasks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SsmMaintenanceWindowTask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmMaintenanceWindowTaskSpec   `json:"spec"`
	Status            SsmMaintenanceWindowTaskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsmMaintenanceWindowTaskList contains a list of SsmMaintenanceWindowTasks
type SsmMaintenanceWindowTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsmMaintenanceWindowTask `json:"items"`
}

// Repository type metadata.
var (
	SsmMaintenanceWindowTaskKind             = "SsmMaintenanceWindowTask"
	SsmMaintenanceWindowTaskGroupKind        = schema.GroupKind{Group: Group, Kind: SsmMaintenanceWindowTaskKind}.String()
	SsmMaintenanceWindowTaskKindAPIVersion   = SsmMaintenanceWindowTaskKind + "." + GroupVersion.String()
	SsmMaintenanceWindowTaskGroupVersionKind = GroupVersion.WithKind(SsmMaintenanceWindowTaskKind)
)

func init() {
	SchemeBuilder.Register(&SsmMaintenanceWindowTask{}, &SsmMaintenanceWindowTaskList{})
}
