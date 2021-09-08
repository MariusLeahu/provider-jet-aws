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

type JupyterServerAppSettingsDefaultResourceSpecObservation struct {
}

type JupyterServerAppSettingsDefaultResourceSpecParameters struct {
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`

	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn"`
}

type KernelGatewayAppSettingsCustomImageObservation struct {
}

type KernelGatewayAppSettingsCustomImageParameters struct {
	AppImageConfigName string `json:"appImageConfigName" tf:"app_image_config_name"`

	ImageName string `json:"imageName" tf:"image_name"`

	ImageVersionNumber *int64 `json:"imageVersionNumber,omitempty" tf:"image_version_number"`
}

type SagemakerUserProfileObservation struct {
	Arn string `json:"arn" tf:"arn"`

	HomeEfsFileSystemUid string `json:"homeEfsFileSystemUid" tf:"home_efs_file_system_uid"`
}

type SagemakerUserProfileParameters struct {
	DomainId string `json:"domainId" tf:"domain_id"`

	SingleSignOnUserIdentifier *string `json:"singleSignOnUserIdentifier,omitempty" tf:"single_sign_on_user_identifier"`

	SingleSignOnUserValue *string `json:"singleSignOnUserValue,omitempty" tf:"single_sign_on_user_value"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	UserProfileName string `json:"userProfileName" tf:"user_profile_name"`

	UserSettings []UserSettingsParameters `json:"userSettings,omitempty" tf:"user_settings"`
}

type UserSettingsJupyterServerAppSettingsObservation struct {
}

type UserSettingsJupyterServerAppSettingsParameters struct {
	DefaultResourceSpec []JupyterServerAppSettingsDefaultResourceSpecParameters `json:"defaultResourceSpec" tf:"default_resource_spec"`
}

type UserSettingsKernelGatewayAppSettingsDefaultResourceSpecObservation struct {
}

type UserSettingsKernelGatewayAppSettingsDefaultResourceSpecParameters struct {
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`

	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn"`
}

type UserSettingsKernelGatewayAppSettingsObservation struct {
}

type UserSettingsKernelGatewayAppSettingsParameters struct {
	CustomImage []KernelGatewayAppSettingsCustomImageParameters `json:"customImage,omitempty" tf:"custom_image"`

	DefaultResourceSpec []UserSettingsKernelGatewayAppSettingsDefaultResourceSpecParameters `json:"defaultResourceSpec" tf:"default_resource_spec"`
}

type UserSettingsObservation struct {
}

type UserSettingsParameters struct {
	ExecutionRole string `json:"executionRole" tf:"execution_role"`

	JupyterServerAppSettings []UserSettingsJupyterServerAppSettingsParameters `json:"jupyterServerAppSettings,omitempty" tf:"jupyter_server_app_settings"`

	KernelGatewayAppSettings []UserSettingsKernelGatewayAppSettingsParameters `json:"kernelGatewayAppSettings,omitempty" tf:"kernel_gateway_app_settings"`

	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`

	SharingSettings []UserSettingsSharingSettingsParameters `json:"sharingSettings,omitempty" tf:"sharing_settings"`

	TensorBoardAppSettings []UserSettingsTensorBoardAppSettingsParameters `json:"tensorBoardAppSettings,omitempty" tf:"tensor_board_app_settings"`
}

type UserSettingsSharingSettingsObservation struct {
}

type UserSettingsSharingSettingsParameters struct {
	NotebookOutputOption *string `json:"notebookOutputOption,omitempty" tf:"notebook_output_option"`

	S3KmsKeyId *string `json:"s3KmsKeyId,omitempty" tf:"s3_kms_key_id"`

	S3OutputPath *string `json:"s3OutputPath,omitempty" tf:"s3_output_path"`
}

type UserSettingsTensorBoardAppSettingsDefaultResourceSpecObservation struct {
}

type UserSettingsTensorBoardAppSettingsDefaultResourceSpecParameters struct {
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`

	SagemakerImageArn *string `json:"sagemakerImageArn,omitempty" tf:"sagemaker_image_arn"`
}

type UserSettingsTensorBoardAppSettingsObservation struct {
}

type UserSettingsTensorBoardAppSettingsParameters struct {
	DefaultResourceSpec []UserSettingsTensorBoardAppSettingsDefaultResourceSpecParameters `json:"defaultResourceSpec" tf:"default_resource_spec"`
}

// SagemakerUserProfileSpec defines the desired state of SagemakerUserProfile
type SagemakerUserProfileSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SagemakerUserProfileParameters `json:"forProvider"`
}

// SagemakerUserProfileStatus defines the observed state of SagemakerUserProfile.
type SagemakerUserProfileStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SagemakerUserProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerUserProfile is the Schema for the SagemakerUserProfiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SagemakerUserProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SagemakerUserProfileSpec   `json:"spec"`
	Status            SagemakerUserProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerUserProfileList contains a list of SagemakerUserProfiles
type SagemakerUserProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SagemakerUserProfile `json:"items"`
}

// Repository type metadata.
var (
	SagemakerUserProfileKind             = "SagemakerUserProfile"
	SagemakerUserProfileGroupKind        = schema.GroupKind{Group: Group, Kind: SagemakerUserProfileKind}.String()
	SagemakerUserProfileKindAPIVersion   = SagemakerUserProfileKind + "." + GroupVersion.String()
	SagemakerUserProfileGroupVersionKind = GroupVersion.WithKind(SagemakerUserProfileKind)
)

func init() {
	SchemeBuilder.Register(&SagemakerUserProfile{}, &SagemakerUserProfileList{})
}
