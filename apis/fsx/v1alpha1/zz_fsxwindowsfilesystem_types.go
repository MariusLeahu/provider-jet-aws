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

type AuditLogConfigurationObservation struct {
}

type AuditLogConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	AuditLogDestination *string `json:"auditLogDestination,omitempty" tf:"audit_log_destination"`

	// +kubebuilder:validation:Optional
	FileAccessAuditLogLevel *string `json:"fileAccessAuditLogLevel,omitempty" tf:"file_access_audit_log_level"`

	// +kubebuilder:validation:Optional
	FileShareAccessAuditLogLevel *string `json:"fileShareAccessAuditLogLevel,omitempty" tf:"file_share_access_audit_log_level"`
}

type FsxWindowsFileSystemObservation struct {
	Arn string `json:"arn" tf:"arn"`

	DNSName string `json:"dnsName" tf:"dns_name"`

	NetworkInterfaceIds []string `json:"networkInterfaceIds" tf:"network_interface_ids"`

	OwnerID string `json:"ownerId" tf:"owner_id"`

	PreferredFileServerIP string `json:"preferredFileServerIp" tf:"preferred_file_server_ip"`

	RemoteAdministrationEndpoint string `json:"remoteAdministrationEndpoint" tf:"remote_administration_endpoint"`

	VpcID string `json:"vpcId" tf:"vpc_id"`
}

type FsxWindowsFileSystemParameters struct {

	// +kubebuilder:validation:Optional
	ActiveDirectoryID *string `json:"activeDirectoryId,omitempty" tf:"active_directory_id"`

	// +kubebuilder:validation:Optional
	Aliases []string `json:"aliases,omitempty" tf:"aliases"`

	// +kubebuilder:validation:Optional
	AuditLogConfiguration []AuditLogConfigurationParameters `json:"auditLogConfiguration,omitempty" tf:"audit_log_configuration"`

	// +kubebuilder:validation:Optional
	AutomaticBackupRetentionDays *int64 `json:"automaticBackupRetentionDays,omitempty" tf:"automatic_backup_retention_days"`

	// +kubebuilder:validation:Optional
	CopyTagsToBackups *bool `json:"copyTagsToBackups,omitempty" tf:"copy_tags_to_backups"`

	// +kubebuilder:validation:Optional
	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime,omitempty" tf:"daily_automatic_backup_start_time"`

	// +kubebuilder:validation:Optional
	DeploymentType *string `json:"deploymentType,omitempty" tf:"deployment_type"`

	// +kubebuilder:validation:Optional
	KmsKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	// +kubebuilder:validation:Optional
	PreferredSubnetID *string `json:"preferredSubnetId,omitempty" tf:"preferred_subnet_id"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" tf:"security_group_ids"`

	// +kubebuilder:validation:Optional
	SelfManagedActiveDirectory []SelfManagedActiveDirectoryParameters `json:"selfManagedActiveDirectory,omitempty" tf:"self_managed_active_directory"`

	// +kubebuilder:validation:Optional
	SkipFinalBackup *bool `json:"skipFinalBackup,omitempty" tf:"skip_final_backup"`

	// +kubebuilder:validation:Required
	StorageCapacity int64 `json:"storageCapacity" tf:"storage_capacity"`

	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type"`

	// +kubebuilder:validation:Required
	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Required
	ThroughputCapacity int64 `json:"throughputCapacity" tf:"throughput_capacity"`

	// +kubebuilder:validation:Optional
	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime,omitempty" tf:"weekly_maintenance_start_time"`
}

type SelfManagedActiveDirectoryObservation struct {
}

type SelfManagedActiveDirectoryParameters struct {

	// +kubebuilder:validation:Required
	DNSIps []string `json:"dnsIps" tf:"dns_ips"`

	// +kubebuilder:validation:Required
	DomainName string `json:"domainName" tf:"domain_name"`

	// +kubebuilder:validation:Optional
	FileSystemAdministratorsGroup *string `json:"fileSystemAdministratorsGroup,omitempty" tf:"file_system_administrators_group"`

	// +kubebuilder:validation:Optional
	OrganizationalUnitDistinguishedName *string `json:"organizationalUnitDistinguishedName,omitempty" tf:"organizational_unit_distinguished_name"`

	// +kubebuilder:validation:Required
	Password string `json:"password" tf:"password"`

	// +kubebuilder:validation:Required
	Username string `json:"username" tf:"username"`
}

// FsxWindowsFileSystemSpec defines the desired state of FsxWindowsFileSystem
type FsxWindowsFileSystemSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FsxWindowsFileSystemParameters `json:"forProvider"`
}

// FsxWindowsFileSystemStatus defines the observed state of FsxWindowsFileSystem.
type FsxWindowsFileSystemStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FsxWindowsFileSystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FsxWindowsFileSystem is the Schema for the FsxWindowsFileSystems API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type FsxWindowsFileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FsxWindowsFileSystemSpec   `json:"spec"`
	Status            FsxWindowsFileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FsxWindowsFileSystemList contains a list of FsxWindowsFileSystems
type FsxWindowsFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FsxWindowsFileSystem `json:"items"`
}

// Repository type metadata.
var (
	FsxWindowsFileSystemKind             = "FsxWindowsFileSystem"
	FsxWindowsFileSystemGroupKind        = schema.GroupKind{Group: Group, Kind: FsxWindowsFileSystemKind}.String()
	FsxWindowsFileSystemKindAPIVersion   = FsxWindowsFileSystemKind + "." + GroupVersion.String()
	FsxWindowsFileSystemGroupVersionKind = GroupVersion.WithKind(FsxWindowsFileSystemKind)
)

func init() {
	SchemeBuilder.Register(&FsxWindowsFileSystem{}, &FsxWindowsFileSystemList{})
}