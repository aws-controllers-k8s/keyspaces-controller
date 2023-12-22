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

import pytest
import time
import logging
from typing import Dict, Tuple

from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s
from acktest import tags
from e2e import (
    service_marker, CRD_GROUP, CRD_VERSION,
    load_keyspaces_resource, 
)
from e2e.replacement_values import REPLACEMENT_VALUES
from e2e import condition
from e2e import table

RESOURCE_PLURAL = "tables"

CREATE_WAIT_AFTER_SECONDS = 45
DELETE_WAIT_AFTER_SECONDS = 15
MODIFY_WAIT_AFTER_SECONDS = 30

def create_keyspace(name: str, resource_template):
    replacements = REPLACEMENT_VALUES.copy()
    replacements["KEYSPACE_NAME"] = name

    # load resource
    resource_data = load_keyspaces_resource(
        resource_template,
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

def create_table(name: str, keyspace_name: str, resource_template):
    replacements = REPLACEMENT_VALUES.copy()
    replacements["TABLE_NAME"] = name
    replacements["KEYSPACE_NAME"] = keyspace_name

    # load resource
    resource_data = load_keyspaces_resource(
        resource_template,
        additional_replacements=replacements,
    )

    table_reference = k8s.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, "tables",
        name, namespace="default",
    )

    # Create table
    k8s.create_custom_resource(table_reference, resource_data)
    time.sleep(CREATE_WAIT_AFTER_SECONDS)
    table_resource = k8s.wait_resource_consumed_by_controller(table_reference)

    assert table_resource is not None
    assert k8s.get_resource_exists(table_reference)

    return table_reference, table_resource

@pytest.fixture(scope="module")
def table_basic():
    resource_name = random_suffix_name("table", 32, "test")
    keyspace_resource_name = random_suffix_name("keyspace", 32, "test")

    create_keyspace(keyspace_resource_name, "keyspace_basic")

    (ref, cr) = create_table(resource_name, keyspace_resource_name, "table_basic")

    yield ref, cr
    try:
        _, deleted_table = k8s.delete_custom_resource(ref, wait_periods=3, period_length=10)
        assert deleted_table
        time.sleep(DELETE_WAIT_AFTER_SECONDS)
    except:
        pass
    
    table.wait_until_deleted(keyspace_resource_name, resource_name)

@service_marker
@pytest.mark.canary
class TestTable:
    def table_exists(self, keyspace_name: str, table_name: str) -> bool:
        return table.get(keyspace_name, table_name) is not None

    def test_create_delete_table(self, table_basic):
        (ref, res) = table_basic

        keyspace_name = res["spec"]["keyspaceName"]
        table_name = res["spec"]["tableName"]
        condition.assert_synced(ref)

        # Check Table exists
        assert self.table_exists(keyspace_name, table_name)
