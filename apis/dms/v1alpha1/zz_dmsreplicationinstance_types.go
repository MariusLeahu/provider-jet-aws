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

type DmsReplicationInstanceObservation struct {
	ReplicationInstanceArn string `json:"replicationInstanceArn" tf:"replication_instance_arn"`

	ReplicationInstancePrivateIps []string `json:"replicationInstancePrivateIps" tf:"replication_instance_private_ips"`

	ReplicationInstancePublicIps []string `json:"replicationInstancePublicIps" tf:"replication_instance_public_ips"`
}

type DmsReplicationInstanceParameters struct {
	AllocatedStorage *int64 `json:"allocatedStorage,omitempty" tf:"allocated_storage"`

	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade"`

	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately"`

	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade"`

	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`

	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version"`

	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`

	MultiAz *bool `json:"multiAz,omitempty" tf:"multi_az"`

	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window"`

	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible"`

	ReplicationInstanceClass string `json:"replicationInstanceClass" tf:"replication_instance_class"`

	ReplicationInstanceId string `json:"replicationInstanceId" tf:"replication_instance_id"`

	ReplicationSubnetGroupId *string `json:"replicationSubnetGroupId,omitempty" tf:"replication_subnet_group_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids"`
}

// DmsReplicationInstanceSpec defines the desired state of DmsReplicationInstance
type DmsReplicationInstanceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DmsReplicationInstanceParameters `json:"forProvider"`
}

// DmsReplicationInstanceStatus defines the observed state of DmsReplicationInstance.
type DmsReplicationInstanceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DmsReplicationInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DmsReplicationInstance is the Schema for the DmsReplicationInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DmsReplicationInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DmsReplicationInstanceSpec   `json:"spec"`
	Status            DmsReplicationInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DmsReplicationInstanceList contains a list of DmsReplicationInstances
type DmsReplicationInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DmsReplicationInstance `json:"items"`
}

// Repository type metadata.
var (
	DmsReplicationInstanceKind             = "DmsReplicationInstance"
	DmsReplicationInstanceGroupKind        = schema.GroupKind{Group: Group, Kind: DmsReplicationInstanceKind}.String()
	DmsReplicationInstanceKindAPIVersion   = DmsReplicationInstanceKind + "." + GroupVersion.String()
	DmsReplicationInstanceGroupVersionKind = GroupVersion.WithKind(DmsReplicationInstanceKind)
)

func init() {
	SchemeBuilder.Register(&DmsReplicationInstance{}, &DmsReplicationInstanceList{})
}
