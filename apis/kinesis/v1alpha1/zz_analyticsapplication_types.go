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

type AnalyticsApplicationObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CreateTimestamp *string `json:"createTimestamp,omitempty" tf:"create_timestamp,omitempty"`

	LastUpdateTimestamp *string `json:"lastUpdateTimestamp,omitempty" tf:"last_update_timestamp,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	Version *int64 `json:"version,omitempty" tf:"version,omitempty"`
}

type AnalyticsApplicationParameters struct {

	// +kubebuilder:validation:Optional
	CloudwatchLoggingOptions []CloudwatchLoggingOptionsParameters `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`

	// +kubebuilder:validation:Optional
	Code *string `json:"code,omitempty" tf:"code,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Inputs []InputsParameters `json:"inputs,omitempty" tf:"inputs,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Outputs []OutputsParameters `json:"outputs,omitempty" tf:"outputs,omitempty"`

	// +kubebuilder:validation:Optional
	ReferenceDataSources []ReferenceDataSourcesParameters `json:"referenceDataSources,omitempty" tf:"reference_data_sources,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	StartApplication *bool `json:"startApplication,omitempty" tf:"start_application,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CloudwatchLoggingOptionsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CloudwatchLoggingOptionsParameters struct {

	// +kubebuilder:validation:Required
	LogStreamArn *string `json:"logStreamArn" tf:"log_stream_arn,omitempty"`

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`
}

type CsvObservation struct {
}

type CsvParameters struct {

	// +kubebuilder:validation:Required
	RecordColumnDelimiter *string `json:"recordColumnDelimiter" tf:"record_column_delimiter,omitempty"`

	// +kubebuilder:validation:Required
	RecordRowDelimiter *string `json:"recordRowDelimiter" tf:"record_row_delimiter,omitempty"`
}

type InputsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	StreamNames []*string `json:"streamNames,omitempty" tf:"stream_names,omitempty"`
}

type InputsParameters struct {

	// +kubebuilder:validation:Optional
	KinesisFirehose []KinesisFirehoseParameters `json:"kinesisFirehose,omitempty" tf:"kinesis_firehose,omitempty"`

	// +kubebuilder:validation:Optional
	KinesisStream []KinesisStreamParameters `json:"kinesisStream,omitempty" tf:"kinesis_stream,omitempty"`

	// +kubebuilder:validation:Required
	NamePrefix *string `json:"namePrefix" tf:"name_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	Parallelism []ParallelismParameters `json:"parallelism,omitempty" tf:"parallelism,omitempty"`

	// +kubebuilder:validation:Optional
	ProcessingConfiguration []ProcessingConfigurationParameters `json:"processingConfiguration,omitempty" tf:"processing_configuration,omitempty"`

	// +kubebuilder:validation:Required
	Schema []SchemaParameters `json:"schema" tf:"schema,omitempty"`

	// +kubebuilder:validation:Optional
	StartingPositionConfiguration []StartingPositionConfigurationParameters `json:"startingPositionConfiguration,omitempty" tf:"starting_position_configuration,omitempty"`
}

type JSONObservation struct {
}

type JSONParameters struct {

	// +kubebuilder:validation:Required
	RecordRowPath *string `json:"recordRowPath" tf:"record_row_path,omitempty"`
}

type KinesisFirehoseObservation struct {
}

type KinesisFirehoseParameters struct {

	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`
}

type KinesisStreamObservation struct {
}

type KinesisStreamParameters struct {

	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`
}

type LambdaObservation struct {
}

type LambdaParameters struct {

	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`
}

type MappingParametersCsvObservation struct {
}

type MappingParametersCsvParameters struct {

	// +kubebuilder:validation:Required
	RecordColumnDelimiter *string `json:"recordColumnDelimiter" tf:"record_column_delimiter,omitempty"`

	// +kubebuilder:validation:Required
	RecordRowDelimiter *string `json:"recordRowDelimiter" tf:"record_row_delimiter,omitempty"`
}

type MappingParametersJSONObservation struct {
}

type MappingParametersJSONParameters struct {

	// +kubebuilder:validation:Required
	RecordRowPath *string `json:"recordRowPath" tf:"record_row_path,omitempty"`
}

type MappingParametersObservation struct {
}

type MappingParametersParameters struct {

	// +kubebuilder:validation:Optional
	Csv []CsvParameters `json:"csv,omitempty" tf:"csv,omitempty"`

	// +kubebuilder:validation:Optional
	JSON []JSONParameters `json:"json,omitempty" tf:"json,omitempty"`
}

type OutputsKinesisFirehoseObservation struct {
}

type OutputsKinesisFirehoseParameters struct {

	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`
}

type OutputsKinesisStreamObservation struct {
}

type OutputsKinesisStreamParameters struct {

	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`
}

type OutputsLambdaObservation struct {
}

type OutputsLambdaParameters struct {

	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`
}

type OutputsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OutputsParameters struct {

	// +kubebuilder:validation:Optional
	KinesisFirehose []OutputsKinesisFirehoseParameters `json:"kinesisFirehose,omitempty" tf:"kinesis_firehose,omitempty"`

	// +kubebuilder:validation:Optional
	KinesisStream []OutputsKinesisStreamParameters `json:"kinesisStream,omitempty" tf:"kinesis_stream,omitempty"`

	// +kubebuilder:validation:Optional
	Lambda []OutputsLambdaParameters `json:"lambda,omitempty" tf:"lambda,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Schema []OutputsSchemaParameters `json:"schema" tf:"schema,omitempty"`
}

type OutputsSchemaObservation struct {
}

type OutputsSchemaParameters struct {

	// +kubebuilder:validation:Required
	RecordFormatType *string `json:"recordFormatType" tf:"record_format_type,omitempty"`
}

type ParallelismObservation struct {
}

type ParallelismParameters struct {

	// +kubebuilder:validation:Optional
	Count *int64 `json:"count,omitempty" tf:"count,omitempty"`
}

type ProcessingConfigurationObservation struct {
}

type ProcessingConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Lambda []LambdaParameters `json:"lambda" tf:"lambda,omitempty"`
}

type RecordColumnsObservation struct {
}

type RecordColumnsParameters struct {

	// +kubebuilder:validation:Optional
	Mapping *string `json:"mapping,omitempty" tf:"mapping,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	SQLType *string `json:"sqlType" tf:"sql_type,omitempty"`
}

type RecordFormatMappingParametersObservation struct {
}

type RecordFormatMappingParametersParameters struct {

	// +kubebuilder:validation:Optional
	Csv []MappingParametersCsvParameters `json:"csv,omitempty" tf:"csv,omitempty"`

	// +kubebuilder:validation:Optional
	JSON []MappingParametersJSONParameters `json:"json,omitempty" tf:"json,omitempty"`
}

type RecordFormatObservation struct {
	RecordFormatType *string `json:"recordFormatType,omitempty" tf:"record_format_type,omitempty"`
}

type RecordFormatParameters struct {

	// +kubebuilder:validation:Optional
	MappingParameters []MappingParametersParameters `json:"mappingParameters,omitempty" tf:"mapping_parameters,omitempty"`
}

type ReferenceDataSourcesObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ReferenceDataSourcesParameters struct {

	// +kubebuilder:validation:Required
	S3 []S3Parameters `json:"s3" tf:"s3,omitempty"`

	// +kubebuilder:validation:Required
	Schema []ReferenceDataSourcesSchemaParameters `json:"schema" tf:"schema,omitempty"`

	// +kubebuilder:validation:Required
	TableName *string `json:"tableName" tf:"table_name,omitempty"`
}

type ReferenceDataSourcesSchemaObservation struct {
}

type ReferenceDataSourcesSchemaParameters struct {

	// +kubebuilder:validation:Required
	RecordColumns []SchemaRecordColumnsParameters `json:"recordColumns" tf:"record_columns,omitempty"`

	// +kubebuilder:validation:Optional
	RecordEncoding *string `json:"recordEncoding,omitempty" tf:"record_encoding,omitempty"`

	// +kubebuilder:validation:Required
	RecordFormat []SchemaRecordFormatParameters `json:"recordFormat" tf:"record_format,omitempty"`
}

type S3Observation struct {
}

type S3Parameters struct {

	// +kubebuilder:validation:Required
	BucketArn *string `json:"bucketArn" tf:"bucket_arn,omitempty"`

	// +kubebuilder:validation:Required
	FileKey *string `json:"fileKey" tf:"file_key,omitempty"`

	// +kubebuilder:validation:Required
	RoleArn *string `json:"roleArn" tf:"role_arn,omitempty"`
}

type SchemaObservation struct {
}

type SchemaParameters struct {

	// +kubebuilder:validation:Required
	RecordColumns []RecordColumnsParameters `json:"recordColumns" tf:"record_columns,omitempty"`

	// +kubebuilder:validation:Optional
	RecordEncoding *string `json:"recordEncoding,omitempty" tf:"record_encoding,omitempty"`

	// +kubebuilder:validation:Required
	RecordFormat []RecordFormatParameters `json:"recordFormat" tf:"record_format,omitempty"`
}

type SchemaRecordColumnsObservation struct {
}

type SchemaRecordColumnsParameters struct {

	// +kubebuilder:validation:Optional
	Mapping *string `json:"mapping,omitempty" tf:"mapping,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	SQLType *string `json:"sqlType" tf:"sql_type,omitempty"`
}

type SchemaRecordFormatObservation struct {
	RecordFormatType *string `json:"recordFormatType,omitempty" tf:"record_format_type,omitempty"`
}

type SchemaRecordFormatParameters struct {

	// +kubebuilder:validation:Optional
	MappingParameters []RecordFormatMappingParametersParameters `json:"mappingParameters,omitempty" tf:"mapping_parameters,omitempty"`
}

type StartingPositionConfigurationObservation struct {
}

type StartingPositionConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	StartingPosition *string `json:"startingPosition,omitempty" tf:"starting_position,omitempty"`
}

// AnalyticsApplicationSpec defines the desired state of AnalyticsApplication
type AnalyticsApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AnalyticsApplicationParameters `json:"forProvider"`
}

// AnalyticsApplicationStatus defines the observed state of AnalyticsApplication.
type AnalyticsApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AnalyticsApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AnalyticsApplication is the Schema for the AnalyticsApplications API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type AnalyticsApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AnalyticsApplicationSpec   `json:"spec"`
	Status            AnalyticsApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AnalyticsApplicationList contains a list of AnalyticsApplications
type AnalyticsApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AnalyticsApplication `json:"items"`
}

// Repository type metadata.
var (
	AnalyticsApplication_Kind             = "AnalyticsApplication"
	AnalyticsApplication_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AnalyticsApplication_Kind}.String()
	AnalyticsApplication_KindAPIVersion   = AnalyticsApplication_Kind + "." + CRDGroupVersion.String()
	AnalyticsApplication_GroupVersionKind = CRDGroupVersion.WithKind(AnalyticsApplication_Kind)
)

func init() {
	SchemeBuilder.Register(&AnalyticsApplication{}, &AnalyticsApplicationList{})
}