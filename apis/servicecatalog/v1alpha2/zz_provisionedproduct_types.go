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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ProvisionedProductObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CloudwatchDashboardNames []*string `json:"cloudwatchDashboardNames,omitempty" tf:"cloudwatch_dashboard_names,omitempty"`

	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LastProvisioningRecordID *string `json:"lastProvisioningRecordId,omitempty" tf:"last_provisioning_record_id,omitempty"`

	LastRecordID *string `json:"lastRecordId,omitempty" tf:"last_record_id,omitempty"`

	LastSuccessfulProvisioningRecordID *string `json:"lastSuccessfulProvisioningRecordId,omitempty" tf:"last_successful_provisioning_record_id,omitempty"`

	LaunchRoleArn *string `json:"launchRoleArn,omitempty" tf:"launch_role_arn,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	StatusMessage *string `json:"statusMessage,omitempty" tf:"status_message,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProvisionedProductParameters struct {

	// +kubebuilder:validation:Optional
	AcceptLanguage *string `json:"acceptLanguage,omitempty" tf:"accept_language,omitempty"`

	// +kubebuilder:validation:Optional
	IgnoreErrors *bool `json:"ignoreErrors,omitempty" tf:"ignore_errors,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NotificationArns []*string `json:"notificationArns,omitempty" tf:"notification_arns,omitempty"`

	// +kubebuilder:validation:Optional
	PathID *string `json:"pathId,omitempty" tf:"path_id,omitempty"`

	// +kubebuilder:validation:Optional
	PathName *string `json:"pathName,omitempty" tf:"path_name,omitempty"`

	// +kubebuilder:validation:Optional
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProductName *string `json:"productName,omitempty" tf:"product_name,omitempty"`

	// +kubebuilder:validation:Optional
	ProvisioningArtifactID *string `json:"provisioningArtifactId,omitempty" tf:"provisioning_artifact_id,omitempty"`

	// +kubebuilder:validation:Optional
	ProvisioningArtifactName *string `json:"provisioningArtifactName,omitempty" tf:"provisioning_artifact_name,omitempty"`

	// +kubebuilder:validation:Optional
	ProvisioningParameters []ProvisioningParametersParameters `json:"provisioningParameters,omitempty" tf:"provisioning_parameters,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RetainPhysicalResources *bool `json:"retainPhysicalResources,omitempty" tf:"retain_physical_resources,omitempty"`

	// +kubebuilder:validation:Optional
	StackSetProvisioningPreferences []StackSetProvisioningPreferencesParameters `json:"stackSetProvisioningPreferences,omitempty" tf:"stack_set_provisioning_preferences,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ProvisioningParametersObservation struct {
}

type ProvisioningParametersParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	UsePreviousValue *bool `json:"usePreviousValue,omitempty" tf:"use_previous_value,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type StackSetProvisioningPreferencesObservation struct {
}

type StackSetProvisioningPreferencesParameters struct {

	// +kubebuilder:validation:Optional
	Accounts []*string `json:"accounts,omitempty" tf:"accounts,omitempty"`

	// +kubebuilder:validation:Optional
	FailureToleranceCount *int64 `json:"failureToleranceCount,omitempty" tf:"failure_tolerance_count,omitempty"`

	// +kubebuilder:validation:Optional
	FailureTolerancePercentage *int64 `json:"failureTolerancePercentage,omitempty" tf:"failure_tolerance_percentage,omitempty"`

	// +kubebuilder:validation:Optional
	MaxConcurrencyCount *int64 `json:"maxConcurrencyCount,omitempty" tf:"max_concurrency_count,omitempty"`

	// +kubebuilder:validation:Optional
	MaxConcurrencyPercentage *int64 `json:"maxConcurrencyPercentage,omitempty" tf:"max_concurrency_percentage,omitempty"`

	// +kubebuilder:validation:Optional
	Regions []*string `json:"regions,omitempty" tf:"regions,omitempty"`
}

// ProvisionedProductSpec defines the desired state of ProvisionedProduct
type ProvisionedProductSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProvisionedProductParameters `json:"forProvider"`
}

// ProvisionedProductStatus defines the observed state of ProvisionedProduct.
type ProvisionedProductStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProvisionedProductObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProvisionedProduct is the Schema for the ProvisionedProducts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ProvisionedProduct struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProvisionedProductSpec   `json:"spec"`
	Status            ProvisionedProductStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProvisionedProductList contains a list of ProvisionedProducts
type ProvisionedProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProvisionedProduct `json:"items"`
}

// Repository type metadata.
var (
	ProvisionedProduct_Kind             = "ProvisionedProduct"
	ProvisionedProduct_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProvisionedProduct_Kind}.String()
	ProvisionedProduct_KindAPIVersion   = ProvisionedProduct_Kind + "." + CRDGroupVersion.String()
	ProvisionedProduct_GroupVersionKind = CRDGroupVersion.WithKind(ProvisionedProduct_Kind)
)

func init() {
	SchemeBuilder.Register(&ProvisionedProduct{}, &ProvisionedProductList{})
}
