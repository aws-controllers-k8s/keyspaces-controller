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

"""Utilities for working with Type resources"""

import datetime
import time
import typing
import logging

import boto3
import pytest

from acktest.aws.identity import get_region

DEFAULT_WAIT_UNTIL_TIMEOUT_SECONDS = 60*2
DEFAULT_WAIT_UNTIL_INTERVAL_SECONDS = 15
DEFAULT_WAIT_UNTIL_DELETED_TIMEOUT_SECONDS = 60*2
DEFAULT_WAIT_UNTIL_DELETED_INTERVAL_SECONDS = 15

TypeMatchFunc = typing.NewType(
    'TypeMatchFunc',
    typing.Callable[[dict], bool],
)

class StatusMatcher:
    def __init__(self, status):
        self.match_on = status

    def __call__(self, record: dict) -> bool:
        return ('TypeStatus' in record
                and record['TypeStatus'] == self.match_on)


def status_matches(status: str) -> TypeMatchFunc:
    return StatusMatcher(status)

def wait_until(
        type_name: str,
        keyspace_name: str,
        match_fn: TypeMatchFunc,
        timeout_seconds: int = DEFAULT_WAIT_UNTIL_TIMEOUT_SECONDS,
        interval_seconds: int = DEFAULT_WAIT_UNTIL_INTERVAL_SECONDS,
    ) -> None:
    """Waits until a Type with a supplied name is returned from the Keyspaces
    API and the matching functor returns True.

    Usage:
        from e2e.type_ import wait_until, status_matches

        wait_until(
            type_name,
            keyspace_name,
            status_matches("ACTIVE"),
        )

    Raises:
        pytest.fail upon timeout
    """
    now = datetime.datetime.now()
    timeout = now + datetime.timedelta(seconds=timeout_seconds)

    while not match_fn(get(type_name, keyspace_name)):
        if datetime.datetime.now() >= timeout:
            pytest.fail("failed to match type before timeout")
        time.sleep(interval_seconds)

def wait_until_deleted(
        type_name: str,
        keyspace_name: str,
        timeout_seconds: int = DEFAULT_WAIT_UNTIL_DELETED_TIMEOUT_SECONDS,
        interval_seconds: int = DEFAULT_WAIT_UNTIL_DELETED_INTERVAL_SECONDS,
    ) -> None:
    """Waits until Type with a supplied type_name is no longer returned from
    the Keyspaces API.

    Usage:
        from e2e.type_ import wait_until_deleted

        wait_until_deleted(
          type_name,
          keyspace_name
        )

    Raises:
        pytest.fail upon timeout or if the Type goes to any other status
        other than 'deleting'
    """
    now = datetime.datetime.now()
    timeout = now + datetime.timedelta(seconds=timeout_seconds)

    while True:
        if datetime.datetime.now() >= timeout:
            pytest.fail(
                "Timed out waiting for Type to be "
                "deleted in Keyspaces API"
            )
        time.sleep(interval_seconds)

        latest = get(type_name,keyspace_name)
        if latest is None:
            break

        if latest['status'] != "DELETING":
            pytest.fail(
                "Status is not 'deleting' for Type that was "
                "deleted. Status is " + latest['status']
            )

def get(type_name,keyspace_name):
    """Returns a dict containing the Type record from the keyspaces API.

    If no such Type exists, returns None.
    """
    c = boto3.client('keyspaces', region_name=get_region())
    try:
        resp = c.get_type(keyspaceName=keyspace_name, typeName=type_name)
        return resp
    except c.exceptions.ResourceNotFoundException:
        logging.info("Type %s not found", type_name)
        return None
    except c.exceptions.ValidationException:
        logging.info(
          "Couldn't verify %s exists. Here's why: %s",
          type_name,
          c.exceptions
        )
        return None