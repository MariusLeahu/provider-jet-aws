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

type CacheNodesObservation struct {
	Address string `json:"address" tf:"address"`

	AvailabilityZone string `json:"availabilityZone" tf:"availability_zone"`

	Id string `json:"id" tf:"id"`

	Port int64 `json:"port" tf:"port"`
}

type CacheNodesParameters struct {
}

type ElasticacheClusterObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CacheNodes []CacheNodesObservation `json:"cacheNodes" tf:"cache_nodes"`

	ClusterAddress string `json:"clusterAddress" tf:"cluster_address"`

	ConfigurationEndpoint string `json:"configurationEndpoint" tf:"configuration_endpoint"`

	EngineVersionActual string `json:"engineVersionActual" tf:"engine_version_actual"`
}

type ElasticacheClusterParameters struct {
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately"`

	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`

	AzMode *string `json:"azMode,omitempty" tf:"az_mode"`

	ClusterId string `json:"clusterId" tf:"cluster_id"`

	Engine *string `json:"engine,omitempty" tf:"engine"`

	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version"`

	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier"`

	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window"`

	NodeType *string `json:"nodeType,omitempty" tf:"node_type"`

	NotificationTopicArn *string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn"`

	NumCacheNodes *int64 `json:"numCacheNodes,omitempty" tf:"num_cache_nodes"`

	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name"`

	Port *int64 `json:"port,omitempty" tf:"port"`

	PreferredAvailabilityZones []string `json:"preferredAvailabilityZones,omitempty" tf:"preferred_availability_zones"`

	ReplicationGroupId *string `json:"replicationGroupId,omitempty" tf:"replication_group_id"`

	SecurityGroupIds []string `json:"securityGroupIds,omitempty" tf:"security_group_ids"`

	SecurityGroupNames []string `json:"securityGroupNames,omitempty" tf:"security_group_names"`

	SnapshotArns []string `json:"snapshotArns,omitempty" tf:"snapshot_arns"`

	SnapshotName *string `json:"snapshotName,omitempty" tf:"snapshot_name"`

	SnapshotRetentionLimit *int64 `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit"`

	SnapshotWindow *string `json:"snapshotWindow,omitempty" tf:"snapshot_window"`

	SubnetGroupName *string `json:"subnetGroupName,omitempty" tf:"subnet_group_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// ElasticacheClusterSpec defines the desired state of ElasticacheCluster
type ElasticacheClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ElasticacheClusterParameters `json:"forProvider"`
}

// ElasticacheClusterStatus defines the observed state of ElasticacheCluster.
type ElasticacheClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ElasticacheClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticacheCluster is the Schema for the ElasticacheClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ElasticacheCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticacheClusterSpec   `json:"spec"`
	Status            ElasticacheClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticacheClusterList contains a list of ElasticacheClusters
type ElasticacheClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticacheCluster `json:"items"`
}

// Repository type metadata.
var (
	ElasticacheClusterKind             = "ElasticacheCluster"
	ElasticacheClusterGroupKind        = schema.GroupKind{Group: Group, Kind: ElasticacheClusterKind}.String()
	ElasticacheClusterKindAPIVersion   = ElasticacheClusterKind + "." + GroupVersion.String()
	ElasticacheClusterGroupVersionKind = GroupVersion.WithKind(ElasticacheClusterKind)
)

func init() {
	SchemeBuilder.Register(&ElasticacheCluster{}, &ElasticacheClusterList{})
}
