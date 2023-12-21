// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TableSpec defines the desired state of Table.
type TableSpec struct {

	// Specifies the read/write throughput capacity mode for the table. The options
	// are:
	//
	//   - throughputMode:PAY_PER_REQUEST and
	//
	//   - throughputMode:PROVISIONED - Provisioned capacity mode requires readCapacityUnits
	//     and writeCapacityUnits as input.
	//
	// The default is throughput_mode:PAY_PER_REQUEST.
	//
	// For more information, see Read/write capacity modes (https://docs.aws.amazon.com/keyspaces/latest/devguide/ReadWriteCapacityMode.html)
	// in the Amazon Keyspaces Developer Guide.
	CapacitySpecification *CapacitySpecification `json:"capacitySpecification,omitempty"`
	// Enables client-side timestamps for the table. By default, the setting is
	// disabled. You can enable client-side timestamps with the following option:
	//
	//   - status: "enabled"
	//
	// Once client-side timestamps are enabled for a table, this setting cannot
	// be disabled.
	ClientSideTimestamps *ClientSideTimestamps `json:"clientSideTimestamps,omitempty"`
	// This parameter allows to enter a description of the table.
	Comment *Comment `json:"comment,omitempty"`
	// The default Time to Live setting in seconds for the table.
	//
	// For more information, see Setting the default TTL value for a table (https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL-how-it-works.html#ttl-howitworks_default_ttl)
	// in the Amazon Keyspaces Developer Guide.
	DefaultTimeToLive *int64 `json:"defaultTimeToLive,omitempty"`
	// Specifies how the encryption key for encryption at rest is managed for the
	// table. You can choose one of the following KMS key (KMS key):
	//
	//   - type:AWS_OWNED_KMS_KEY - This key is owned by Amazon Keyspaces.
	//
	//   - type:CUSTOMER_MANAGED_KMS_KEY - This key is stored in your account and
	//     is created, owned, and managed by you. This option requires the kms_key_identifier
	//     of the KMS key in Amazon Resource Name (ARN) format as input.
	//
	// The default is type:AWS_OWNED_KMS_KEY.
	//
	// For more information, see Encryption at rest (https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html)
	// in the Amazon Keyspaces Developer Guide.
	EncryptionSpecification *EncryptionSpecification `json:"encryptionSpecification,omitempty"`
	// The name of the keyspace that the table is going to be created in.
	// +kubebuilder:validation:Required
	KeyspaceName *string `json:"keyspaceName"`
	// Specifies if pointInTimeRecovery is enabled or disabled for the table. The
	// options are:
	//
	//   - status=ENABLED
	//
	//   - status=DISABLED
	//
	// If it's not specified, the default is status=DISABLED.
	//
	// For more information, see Point-in-time recovery (https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery.html)
	// in the Amazon Keyspaces Developer Guide.
	PointInTimeRecovery *PointInTimeRecovery `json:"pointInTimeRecovery,omitempty"`
	// The schemaDefinition consists of the following parameters.
	//
	// For each column to be created:
	//
	//   - name - The name of the column.
	//
	//   - type - An Amazon Keyspaces data type. For more information, see Data
	//     types (https://docs.aws.amazon.com/keyspaces/latest/devguide/cql.elements.html#cql.data-types)
	//     in the Amazon Keyspaces Developer Guide.
	//
	// The primary key of the table consists of the following columns:
	//
	//   - partitionKeys - The partition key can be a single column, or it can
	//     be a compound value composed of two or more columns. The partition key
	//     portion of the primary key is required and determines how Amazon Keyspaces
	//     stores your data.
	//
	//   - name - The name of each partition key column.
	//
	//   - clusteringKeys - The optional clustering column portion of your primary
	//     key determines how the data is clustered and sorted within each partition.
	//
	//   - name - The name of the clustering column.
	//
	//   - orderBy - Sets the ascendant (ASC) or descendant (DESC) order modifier.
	//     To define a column as static use staticColumns - Static columns store
	//     values that are shared by all rows in the same partition:
	//
	//   - name - The name of the column.
	//
	//   - type - An Amazon Keyspaces data type.
	//
	// +kubebuilder:validation:Required
	SchemaDefinition *SchemaDefinition `json:"schemaDefinition"`
	// The name of the table.
	// +kubebuilder:validation:Required
	TableName *string `json:"tableName"`
	// A list of key-value pair tags to be attached to the resource.
	//
	// For more information, see Adding tags and labels to Amazon Keyspaces resources
	// (https://docs.aws.amazon.com/keyspaces/latest/devguide/tagging-keyspaces.html)
	// in the Amazon Keyspaces Developer Guide.
	Tags []*Tag `json:"tags,omitempty"`
	// Enables Time to Live custom settings for the table. The options are:
	//
	//   - status:enabled
	//
	//   - status:disabled
	//
	// The default is status:disabled. After ttl is enabled, you can't disable it
	// for the table.
	//
	// For more information, see Expiring data by using Amazon Keyspaces Time to
	// Live (TTL) (https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL.html)
	// in the Amazon Keyspaces Developer Guide.
	Ttl *TimeToLive `json:"ttl,omitempty"`
}

// TableStatus defines the observed state of Table
type TableStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The unique identifier of the table in the format of an Amazon Resource Name
	// (ARN).
	// +kubebuilder:validation:Optional
	ResourceARN *string `json:"resourceARN,omitempty"`
}

// Table is the Schema for the Tables API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Table struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TableSpec   `json:"spec,omitempty"`
	Status            TableStatus `json:"status,omitempty"`
}

// TableList contains a list of Table
// +kubebuilder:object:root=true
type TableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Table `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Table{}, &TableList{})
}