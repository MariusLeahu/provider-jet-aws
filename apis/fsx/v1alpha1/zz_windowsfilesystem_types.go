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

type AuditLogConfigurationObservation struct {
}

type AuditLogConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	AuditLogDestination *string `json:"auditLogDestination,omitempty" tf:"audit_log_destination,omitempty"`

	// +kubebuilder:validation:Optional
	FileAccessAuditLogLevel *string `json:"fileAccessAuditLogLevel,omitempty" tf:"file_access_audit_log_level,omitempty"`

	// +kubebuilder:validation:Optional
	FileShareAccessAuditLogLevel *string `json:"fileShareAccessAuditLogLevel,omitempty" tf:"file_share_access_audit_log_level,omitempty"`
}

type SelfManagedActiveDirectoryObservation struct {
}

type SelfManagedActiveDirectoryParameters struct {

	// +kubebuilder:validation:Required
	DNSIps []*string `json:"dnsIps" tf:"dns_ips,omitempty"`

	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// +kubebuilder:validation:Optional
	FileSystemAdministratorsGroup *string `json:"fileSystemAdministratorsGroup,omitempty" tf:"file_system_administrators_group,omitempty"`

	// +kubebuilder:validation:Optional
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName,omitempty" tf:"organizational_unit_distinguished_name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Username *string `json:"username" tf:"username,omitempty"`
}

type WindowsFileSystemObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	NetworkInterfaceIds []*string `json:"networkInterfaceIds,omitempty" tf:"network_interface_ids,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	PreferredFileServerIP *string `json:"preferredFileServerIp,omitempty" tf:"preferred_file_server_ip,omitempty"`

	RemoteAdministrationEndpoint *string `json:"remoteAdministrationEndpoint,omitempty" tf:"remote_administration_endpoint,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	VpcID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type WindowsFileSystemParameters struct {

	// +kubebuilder:validation:Optional
	ActiveDirectoryID *string `json:"activeDirectoryId,omitempty" tf:"active_directory_id,omitempty"`

	// +kubebuilder:validation:Optional
	Aliases []*string `json:"aliases,omitempty" tf:"aliases,omitempty"`

	// +kubebuilder:validation:Optional
	AuditLogConfiguration []AuditLogConfigurationParameters `json:"auditLogConfiguration,omitempty" tf:"audit_log_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	AutomaticBackupRetentionDays *int64 `json:"automaticBackupRetentionDays,omitempty" tf:"automatic_backup_retention_days,omitempty"`

	// +kubebuilder:validation:Optional
	CopyTagsToBackups *bool `json:"copyTagsToBackups,omitempty" tf:"copy_tags_to_backups,omitempty"`

	// +kubebuilder:validation:Optional
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime,omitempty" tf:"daily_automatic_backup_start_time,omitempty"`

	// +kubebuilder:validation:Optional
	DeploymentType *string `json:"deploymentType,omitempty" tf:"deployment_type,omitempty"`

	// +kubebuilder:validation:Optional
	KmsKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	PreferredSubnetID *string `json:"preferredSubnetId,omitempty" tf:"preferred_subnet_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	SelfManagedActiveDirectory []SelfManagedActiveDirectoryParameters `json:"selfManagedActiveDirectory,omitempty" tf:"self_managed_active_directory,omitempty"`

	// +kubebuilder:validation:Optional
	SkipFinalBackup *bool `json:"skipFinalBackup,omitempty" tf:"skip_final_backup,omitempty"`

	// +kubebuilder:validation:Required
	StorageCapacity *int64 `json:"storageCapacity" tf:"storage_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// +kubebuilder:validation:Required
	SubnetIds []*string `json:"subnetIds" tf:"subnet_ids,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	ThroughputCapacity *int64 `json:"throughputCapacity" tf:"throughput_capacity,omitempty"`

	// +kubebuilder:validation:Optional
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime,omitempty" tf:"weekly_maintenance_start_time,omitempty"`
}

// WindowsFileSystemSpec defines the desired state of WindowsFileSystem
type WindowsFileSystemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WindowsFileSystemParameters `json:"forProvider"`
}

// WindowsFileSystemStatus defines the observed state of WindowsFileSystem.
type WindowsFileSystemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WindowsFileSystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WindowsFileSystem is the Schema for the WindowsFileSystems API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type WindowsFileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WindowsFileSystemSpec   `json:"spec"`
	Status            WindowsFileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WindowsFileSystemList contains a list of WindowsFileSystems
type WindowsFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WindowsFileSystem `json:"items"`
}

// Repository type metadata.
var (
	WindowsFileSystem_Kind             = "WindowsFileSystem"
	WindowsFileSystem_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WindowsFileSystem_Kind}.String()
	WindowsFileSystem_KindAPIVersion   = WindowsFileSystem_Kind + "." + CRDGroupVersion.String()
	WindowsFileSystem_GroupVersionKind = CRDGroupVersion.WithKind(WindowsFileSystem_Kind)
)

func init() {
	SchemeBuilder.Register(&WindowsFileSystem{}, &WindowsFileSystemList{})
}