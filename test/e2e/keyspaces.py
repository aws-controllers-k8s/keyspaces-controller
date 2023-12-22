# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
#	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Utilities for working with Keyspaces resources"""

import datetime
import time
import typing
import logging

import boto3
import pytest

from acktest.aws.identity import get_region

DEFAULT_WAIT_UNTIL_TIMEOUT_SECONDS = 60
DEFAULT_WAIT_UNTIL_INTERVAL_SECONDS = 5

KeyspaceMatchFunc = typing.NewType(
    'KeyspaceMatchFunc',
    typing.Callable[[dict], bool],
)

class StatusMatcher:
    def __init__(self, status):
        self.match_on = status

    def __call__(self, record: dict) -> bool:
        return ('KeyspaceStatus' in record
                and record['KeyspaceStatus'] == self.match_on)


def status_matches(status: str) -> KeyspaceMatchFunc:
    return StatusMatcher(status)

def wait_until(
        keyspace_name: str,
        match_fn: KeyspaceMatchFunc,
        timeout_seconds: int = DEFAULT_WAIT_UNTIL_TIMEOUT_SECONDS,
        interval_seconds: int = DEFAULT_WAIT_UNTIL_INTERVAL_SECONDS,
    ) -> None:
    """Waits until a Keyspace with a supplied name is returned from the Keyspaces
    API and the matching functor returns True.

    Usage:
        from e2e.keyspaces import wait_until, status_matches

        wait_until(
            keyspace_name,
            status_matches("ACTIVE"),
        )

    Raises:
        pytest.fail upon timeout
    """
    now = datetime.datetime.now()
    timeout = now + datetime.timedelta(seconds=timeout_seconds)

    while not match_fn(get(keyspace_name)):
        if datetime.datetime.now() >= timeout:
            pytest.fail("failed to match keyspace before timeout")
        time.sleep(interval_seconds)


def get(keyspace_name):
    """Returns a dict containing the Role record from the keyspaces API.

    If no such Keyspace exists, returns None.
    """
    c = boto3.client('keyspaces', region_name=get_region())
    try:
        resp = c.get_keyspace(keyspaceName=keyspace_name)
        return resp['keyspaceName']
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