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

// Represents the properties of a keyspace.
type KeyspaceSummary struct {
	KeyspaceName        *string   `json:"keyspaceName,omitempty"`
	ReplicationRegions  []*string `json:"replicationRegions,omitempty"`
	ReplicationStrategy *string   `json:"replicationStrategy,omitempty"`
	ResourceARN         *string   `json:"resourceARN,omitempty"`
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

// Returns the name of the specified table, the keyspace it is stored in, and
// the unique identifier in the format of an Amazon Resource Name (ARN).
type TableSummary struct {
	KeyspaceName *string `json:"keyspaceName,omitempty"`
	ResourceARN  *string `json:"resourceARN,omitempty"`
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
