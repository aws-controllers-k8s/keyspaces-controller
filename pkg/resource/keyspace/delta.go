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

package keyspace

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.KeyspaceName, b.ko.Spec.KeyspaceName) {
		delta.Add("Spec.KeyspaceName", a.ko.Spec.KeyspaceName, b.ko.Spec.KeyspaceName)
	} else if a.ko.Spec.KeyspaceName != nil && b.ko.Spec.KeyspaceName != nil {
		if *a.ko.Spec.KeyspaceName != *b.ko.Spec.KeyspaceName {
			delta.Add("Spec.KeyspaceName", a.ko.Spec.KeyspaceName, b.ko.Spec.KeyspaceName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ReplicationSpecification, b.ko.Spec.ReplicationSpecification) {
		delta.Add("Spec.ReplicationSpecification", a.ko.Spec.ReplicationSpecification, b.ko.Spec.ReplicationSpecification)
	} else if a.ko.Spec.ReplicationSpecification != nil && b.ko.Spec.ReplicationSpecification != nil {
		if len(a.ko.Spec.ReplicationSpecification.RegionList) != len(b.ko.Spec.ReplicationSpecification.RegionList) {
			delta.Add("Spec.ReplicationSpecification.RegionList", a.ko.Spec.ReplicationSpecification.RegionList, b.ko.Spec.ReplicationSpecification.RegionList)
		} else if len(a.ko.Spec.ReplicationSpecification.RegionList) > 0 {
			if !ackcompare.SliceStringPEqual(a.ko.Spec.ReplicationSpecification.RegionList, b.ko.Spec.ReplicationSpecification.RegionList) {
				delta.Add("Spec.ReplicationSpecification.RegionList", a.ko.Spec.ReplicationSpecification.RegionList, b.ko.Spec.ReplicationSpecification.RegionList)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ReplicationSpecification.ReplicationStrategy, b.ko.Spec.ReplicationSpecification.ReplicationStrategy) {
			delta.Add("Spec.ReplicationSpecification.ReplicationStrategy", a.ko.Spec.ReplicationSpecification.ReplicationStrategy, b.ko.Spec.ReplicationSpecification.ReplicationStrategy)
		} else if a.ko.Spec.ReplicationSpecification.ReplicationStrategy != nil && b.ko.Spec.ReplicationSpecification.ReplicationStrategy != nil {
			if *a.ko.Spec.ReplicationSpecification.ReplicationStrategy != *b.ko.Spec.ReplicationSpecification.ReplicationStrategy {
				delta.Add("Spec.ReplicationSpecification.ReplicationStrategy", a.ko.Spec.ReplicationSpecification.ReplicationStrategy, b.ko.Spec.ReplicationSpecification.ReplicationStrategy)
			}
		}
	}
	if !ackcompare.MapStringStringEqual(ToACKTags(a.ko.Spec.Tags), ToACKTags(b.ko.Spec.Tags)) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}

	return delta
}
