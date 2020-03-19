#!/bin/bash

echo "==============================================="
echo "Beginning scloud integration tests"
echo "==============================================="
echo "TEST_USERNAME=$TEST_USERNAME"
echo "TEST_SCLOUD_TENANT=$TEST_SCLOUD_TENANT"
echo "SPLUNK_CLOUD_HOST=$SPLUNK_CLOUD_HOST"
echo "==============================================="

COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES=$(go list ./... | grep -v test | awk -v ORS=, '{ print $1 }' | sed 's/,$//')

if [[ "$allow_failures" == "1" ]]; then
    echo "Running integration tests but not gating on failures..."
else
    echo "Running integration tests and gating on failures..."
fi

set +e
gotestsum --format short-verbose \
          -- github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/test \
          -timeout 10m
result=$?

if [[ "$result" -gt "0" ]]; then
    echo "Tests FAILED"
    if [[ "$allow_failures" == "1" ]]; then
        echo "... but not gating, exiting with status 0"
        exit 0
    else
        echo "... gating on failure, exiting with status 1"
        exit 1
    fi
fi