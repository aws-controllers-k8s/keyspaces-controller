# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
# 	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Integration tests for the Keyspaces Table API.
"""

import time
import logging
import pytest
from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from e2e import (
    service_marker, CRD_GROUP, CRD_VERSION,
    load_keyspaces_resource, type_, condition
)
from e2e.replacement_values import REPLACEMENT_VALUES

RESOURCE_PLURAL = "types"

CREATE_WAIT_AFTER_SECONDS = 45
DELETE_WAIT_AFTER_SECONDS = 15

def create_keyspace_for_type(name: str):
    replacements = REPLACEMENT_VALUES.copy()
    replacements["KEYSPACE_NAME"] = name

    # load resource
    resource_data = load_keyspaces_resource(
        "keyspace_basic",
        additional_replacements=replacements,
    )
    logging.debug(resource_data)

    keyspace_reference = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, "keyspaces",
        name, namespace="default",
    )

    # Create keyspace
    k8s.create_custom_resource(keyspace_reference, resource_data)
    time.sleep(CREATE_WAIT_AFTER_SECONDS)
    keyspace_resource = k8s.wait_resource_consumed_by_controller(keyspace_reference)

    assert keyspace_reference is not None
    assert k8s.get_resource_exists(keyspace_reference)
    return keyspace_reference, keyspace_resource

def create_type(name: str, keyspace_name: str, resource_template):
    replacements = REPLACEMENT_VALUES.copy()
    replacements["TYPE_NAME"] = name
    replacements["KEYSPACE_NAME"] = keyspace_name

    # load resource
    resource_data = load_keyspaces_resource(
        resource_template,
        additional_replacements=replacements,
    )

    type_reference = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, RESOURCE_PLURAL,
        name, namespace="default",
    )

    # Create type
    k8s.create_custom_resource(type_reference, resource_data)
    time.sleep(CREATE_WAIT_AFTER_SECONDS)
    type_resource = k8s.wait_resource_consumed_by_controller(type_reference)

    assert type_resource is not None
    assert k8s.get_resource_exists(type_reference)
    return type_reference, type_resource

def create_table_with_udt(name: str, keyspace_name: str, type_name: str, resource_template):
    replacements = REPLACEMENT_VALUES.copy()
    replacements["TABLE_NAME"] = name
    replacements["KEYSPACE_NAME"] = keyspace_name
    replacements["TYPE_NAME"] = type_name

    # load resource
    resource_data = load_keyspaces_resource(
        resource_template,
        additional_replacements=replacements,
    )

    table_reference = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, "tables",
        name, namespace="default",
    )

    # Create type
    k8s.create_custom_resource(table_reference, resource_data)
    time.sleep(CREATE_WAIT_AFTER_SECONDS)
    table_resource = k8s.wait_resource_consumed_by_controller(table_reference)

    assert table_resource is not None
    assert k8s.get_resource_exists(table_reference)
    return table_reference, table_resource

@pytest.fixture(scope="module")
def keyspace_basic():
    keyspace_name = random_suffix_name("udt", 32, "test")
    (ref, cr) = create_keyspace_for_type(keyspace_name)
    yield ref, cr
    deleted_keyspace = k8s.delete_custom_resource(ref, wait_periods=3, period_length=10)
    assert deleted_keyspace

@pytest.fixture(scope="module")
def type_basic(keyspace_basic):
    keyspace_ref, _ = keyspace_basic
    type_name = random_suffix_name("udt", 32, "test")
    (ref, cr) = create_type(type_name,keyspace_ref.name, "type_basic")
    yield ref, cr
    deleted_type = k8s.delete_custom_resource(ref, wait_periods=3, period_length=10)
    assert deleted_type

@pytest.fixture(scope="module")
def table_with_udt(keyspace_basic,type_basic):
    table_name = random_suffix_name("udttable", 32, "test")
    keyspace_ref, _ = keyspace_basic
    type_ref, _ = type_basic
    (ref, cr) = create_table_with_udt(table_name, keyspace_ref.name, type_ref.name, "table_with_udt")
    yield ref, cr
    deleted_table = k8s.delete_custom_resource(ref, wait_periods=3, period_length=10)
    assert deleted_table


@service_marker
@pytest.mark.canary
class TestType:
    def type_exists(self, type_name: str, keyspace_name: str) -> bool:
        return type_.get(type_name, keyspace_name) is not None

    def test_create_delete_type(self, keyspace_basic, type_basic):
        (keyspace_basic_ref, _) = keyspace_basic
        (ref, res) = type_basic
        type_name = res["spec"]["typeName"]
        keyspace_name = res["spec"]["keyspaceName"]
        condition.assert_synced(ref)
        condition.assert_synced(keyspace_basic_ref)

        # Check Type exists
        assert self.type_exists(type_name, keyspace_name)

    def test_create_delete_table_with_udt(self, keyspace_basic, type_basic, table_with_udt):
        (keyspace_basic_ref, _) = keyspace_basic
        (ref, res) = type_basic
        (ref_table, res_table) = table_with_udt
        type_name = res["spec"]["typeName"]
        keyspace_name = res["spec"]["keyspaceName"]
        table_name = res_table["spec"]["tableName"]
        condition.assert_synced(keyspace_basic_ref)
        condition.assert_synced(ref)        
        condition.assert_synced(ref_table)

        # Check Type exists
        assert self.type_exists(type_name, keyspace_name)
