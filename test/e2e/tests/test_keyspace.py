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

"""Integration tests for the Keyspaces API.
"""

import pytest
import time
import logging
import time
import boto3

from acktest.resources import random_suffix_name
from acktest.aws.identity import get_region
from acktest.k8s import resource as k8s
from acktest import tags
from e2e import (
    service_marker, CRD_GROUP, CRD_VERSION,
    load_keyspaces_resource, 
)
from e2e.replacement_values import REPLACEMENT_VALUES

RESOURCE_PLURAL = "keyspaces"

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

@pytest.fixture(scope="module")
def keyspace_basic():
    resource_name = random_suffix_name("keyspace", 16)
    (ref, cr) = create_keyspace(resource_name, "keyspace_basic")

    yield ref, cr
    try:
        _, deleted = k8s.delete_custom_resource(ref, wait_periods=3, period_length=10)
        assert deleted
    except:
        pass

@service_marker
@pytest.mark.canary
class TestKeyspace:
    def keyspace_exists(self, keyspace_name: str) -> bool:
      """Returns a dict containing the Role record from the keyspaces API.

      If no such Keyspace exists, returns None.
      """
      c = boto3.client('keyspaces', region_name=get_region())
      try:
          resp = c.get_keyspace(keyspaceName=keyspace_name)
          return resp is not None
      except c.exceptions.ResourceNotFoundException:
          logging.info("Keyspace %s not found", keyspace_name)
          return None
      except c.exceptions.ValidationException:
          logging.info(
            "Couldn't verify %s exists. Here's why: %s",
            keyspace_name,
            c.exceptions
          )
          return None

    def test_create_delete(self, keyspace_basic):
        (ref, res) = keyspace_basic

        keyspace_name = res["spec"]["keyspaceName"]

        # Check Keyspace exists
        assert self.keyspace_exists(keyspace_name)

        # Delete k8s resource
        _, deleted = k8s.delete_custom_resource(ref, wait_periods=15, period_length=15)
        assert deleted is True

        time.sleep(DELETE_WAIT_AFTER_SECONDS)

        # Check keyspace doesn't exist
        assert not self.keyspace_exists(keyspace_name)
    