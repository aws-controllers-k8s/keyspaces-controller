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
	"github.com/aws/aws-sdk-go/aws"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = ackv1alpha1.AWSAccountID("")
)

// Amazon Keyspaces has two read/write capacity modes for processing reads and
// writes on your tables:
//
//   - On-demand (default)
//
//   - Provisioned
//
// The read/write capacity mode that you choose controls how you are charged
// for read and write throughput and how table throughput capacity is managed.
//
// For more information, see Read/write capacity modes (https://docs.aws.amazon.com/keyspaces/latest/devguide/ReadWriteCapacityMode.html)
// in the Amazon Keyspaces Developer Guide.
type CapacitySpecification struct {
	ReadCapacityUnits  *int64  `json:"readCapacityUnits,omitempty"`
	ThroughputMode     *string `json:"throughputMode,omitempty"`
	WriteCapacityUnits *int64  `json:"writeCapacityUnits,omitempty"`
}

// The read/write throughput capacity mode for a table. The options are:
//
//   - throughputMode:PAY_PER_REQUEST and
//
//   - throughputMode:PROVISIONED.
//
// For more information, see Read/write capacity modes (https://docs.aws.amazon.com/keyspaces/latest/devguide/ReadWriteCapacityMode.html)
// in the Amazon Keyspaces Developer Guide.
type CapacitySpecificationSummary struct {
	LastUpdateToPayPerRequestTimestamp *metav1.Time `json:"lastUpdateToPayPerRequestTimestamp,omitempty"`
	ReadCapacityUnits                  *int64       `json:"readCapacityUnits,omitempty"`
	ThroughputMode                     *string      `json:"throughputMode,omitempty"`
	WriteCapacityUnits                 *int64       `json:"writeCapacityUnits,omitempty"`
}

// The client-side timestamp setting of the table.
//
// For more information, see How it works: Amazon Keyspaces client-side timestamps
// (https://docs.aws.amazon.com/keyspaces/latest/devguide/client-side-timestamps-how-it-works.html)
// in the Amazon Keyspaces Developer Guide.
type ClientSideTimestamps struct {
	Status *string `json:"status,omitempty"`
}

// The optional clustering column portion of your primary key determines how
// the data is clustered and sorted within each partition.
type ClusteringKey struct {
	Name    *string `json:"name,omitempty"`
	OrderBy *string `json:"orderBy,omitempty"`
}

// The names and data types of regular columns.
type ColumnDefinition struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type_,omitempty"`
}

// An optional comment that describes the table.
type Comment struct {
	Message *string `json:"message,omitempty"`
}

// Amazon Keyspaces encrypts and decrypts the table data at rest transparently
// and integrates with Key Management Service for storing and managing the encryption
// key. You can choose one of the following KMS keys (KMS keys):
//
//   - Amazon Web Services owned key - This is the default encryption type.
//     The key is owned by Amazon Keyspaces (no additional charge).
//
//   - Customer managed key - This key is stored in your account and is created,
//     owned, and managed by you. You have full control over the customer managed
//     key (KMS charges apply).
//
// For more information about encryption at rest in Amazon Keyspaces, see Encryption
// at rest (https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html)
// in the Amazon Keyspaces Developer Guide.
//
// For more information about KMS, see KMS management service concepts (https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html)
// in the Key Management Service Developer Guide.
type EncryptionSpecification struct {
	KMSKeyIdentifier *string `json:"kmsKeyIdentifier,omitempty"`
	Type             *string `json:"type_,omitempty"`
}

// Represents the properties of a keyspace.
type KeyspaceSummary struct {
	KeyspaceName        *string   `json:"keyspaceName,omitempty"`
	ReplicationRegions  []*string `json:"replicationRegions,omitempty"`
	ReplicationStrategy *string   `json:"replicationStrategy,omitempty"`
	ResourceARN         *string   `json:"resourceARN,omitempty"`
}

// The partition key portion of the primary key is required and determines how
// Amazon Keyspaces stores the data. The partition key can be a single column,
// or it can be a compound value composed of two or more columns.
type PartitionKey struct {
	Name *string `json:"name,omitempty"`
}

// Point-in-time recovery (PITR) helps protect your Amazon Keyspaces tables
// from accidental write or delete operations by providing you continuous backups
// of your table data.
//
// For more information, see Point-in-time recovery (https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery.html)
// in the Amazon Keyspaces Developer Guide.
type PointInTimeRecovery struct {
	Status *string `json:"status,omitempty"`
}

// The point-in-time recovery status of the specified table.
type PointInTimeRecoverySummary struct {
	EarliestRestorableTimestamp *metav1.Time `json:"earliestRestorableTimestamp,omitempty"`
	Status                      *string      `json:"status,omitempty"`
}

// The replication specification of the keyspace includes:
//
//   - regionList - up to six Amazon Web Services Regions where the keyspace
//     is replicated in.
//
//   - replicationStrategy - the required value is SINGLE_REGION or MULTI_REGION.
type ReplicationSpecification struct {
	RegionList          []*string `json:"regionList,omitempty"`
	ReplicationStrategy *string   `json:"replicationStrategy,omitempty"`
}

// Describes the schema of the table.
type SchemaDefinition struct {
	AllColumns     []*ColumnDefinition `json:"allColumns,omitempty"`
	ClusteringKeys []*ClusteringKey    `json:"clusteringKeys,omitempty"`
	PartitionKeys  []*PartitionKey     `json:"partitionKeys,omitempty"`
	StaticColumns  []*StaticColumn     `json:"staticColumns,omitempty"`
}

// The static columns of the table. Static columns store values that are shared
// by all rows in the same partition.
type StaticColumn struct {
	Name *string `json:"name,omitempty"`
}

// Returns the name of the specified table, the keyspace it is stored in, and
// the unique identifier in the format of an Amazon Resource Name (ARN).
type TableSummary struct {
	KeyspaceName *string `json:"keyspaceName,omitempty"`
	ResourceARN  *string `json:"resourceARN,omitempty"`
	TableName    *string `json:"tableName,omitempty"`
}

// Describes a tag. A tag is a key-value pair. You can add up to 50 tags to
// a single Amazon Keyspaces resource.
//
// Amazon Web Services-assigned tag names and values are automatically assigned
// the aws: prefix, which the user cannot assign. Amazon Web Services-assigned
// tag names do not count towards the tag limit of 50. User-assigned tag names
// have the prefix user: in the Cost Allocation Report. You cannot backdate
// the application of a tag.
//
// For more information, see Adding tags and labels to Amazon Keyspaces resources
// (https://docs.aws.amazon.com/keyspaces/latest/devguide/tagging-keyspaces.html)
// in the Amazon Keyspaces Developer Guide.
type Tag struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Enable custom Time to Live (TTL) settings for rows and columns without setting
// a TTL default for the specified table.
//
// For more information, see Enabling TTL on tables (https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL-how-it-works.html#ttl-howitworks_enabling)
// in the Amazon Keyspaces Developer Guide.
type TimeToLive struct {
	Status *string `json:"status,omitempty"`
}
