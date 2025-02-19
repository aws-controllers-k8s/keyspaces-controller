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

// KeyspaceSpec defines the desired state of Keyspace.
type KeyspaceSpec struct {

	// The name of the keyspace to be created.

	// +kubebuilder:validation:Required

	KeyspaceName *string `json:"keyspaceName"`
	// The replication specification of the keyspace includes:
	//
	//    * replicationStrategy - the required value is SINGLE_REGION or MULTI_REGION.
	//
	//    * regionList - if the replicationStrategy is MULTI_REGION, the regionList
	//    requires the current Region and at least one additional Amazon Web Services
	//    Region where the keyspace is going to be replicated in. The maximum number
	//    of supported replication Regions including the current Region is six.

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable once set"

	ReplicationSpecification *ReplicationSpecification `json:"replicationSpecification,omitempty"`
	// A list of key-value pair tags to be attached to the keyspace.
	//
	// For more information, see Adding tags and labels to Amazon Keyspaces resources
	// (https://docs.aws.amazon.com/keyspaces/latest/devguide/tagging-keyspaces.html)
	// in the Amazon Keyspaces Developer Guide.

	Tags []*Tag `json:"tags,omitempty"`
}

// KeyspaceStatus defines the observed state of Keyspace
type KeyspaceStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The unique identifier of the keyspace in the format of an Amazon Resource
	// Name (ARN).
	// +kubebuilder:validation:Optional
	ResourceARN *string `json:"resourceARN,omitempty"`
}

// Keyspace is the Schema for the Keyspaces API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Keyspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyspaceSpec   `json:"spec,omitempty"`
	Status            KeyspaceStatus `json:"status,omitempty"`
}

// KeyspaceList contains a list of Keyspace
// +kubebuilder:object:root=true
type KeyspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Keyspace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Keyspace{}, &KeyspaceList{})
}
