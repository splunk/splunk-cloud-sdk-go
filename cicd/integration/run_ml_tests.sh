#!/bin/bash

echo "==============================================="
echo "Beginning ML integration tests"
echo "==============================================="
echo "TEST_USERNAME=$TEST_USERNAME"
echo "SPLUNK_CLOUD_HOST=$SPLUNK_CLOUD_HOST"
echo "TENANT_ID=$TENANT_ID"
echo "==============================================="

# Get the BEARER_TOKEN setup
CONFIG_FILE="./.token"
if [[ -f $CONFIG_FILE ]]; then
    echo "Token found in $CONFIG_FILE"
    export BEARER_TOKEN=$(cat $CONFIG_FILE)
else
    echo "Token was not set to $CONFIG_FILE"
    exit 1
fi

COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES=$(go list ./... | grep -v test | awk -v ORS=, '{ print $1 }' | sed 's/,$//')

PACKAGE_COVERAGE_PREFIX=./cicd/integration/
FULL_INTEGRATION_TEST_CODECOV_FILE_NAME=integration_test_ml_codecov.out
FULL_INTEGRATION_TEST_CODECOV_PATH=$PACKAGE_COVERAGE_PREFIX$FULL_INTEGRATION_TEST_CODECOV_FILE_NAME

# Required to run just the service tests
if [[ "$allow_failures" == "1" ]]; then
    echo "Running ML integration tests but not gating on failures..."
else
    echo "Running ML integration tests and gating on failures..."
fi

set +e
go test -v -coverpkg $COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES \
            -covermode=count \
            -coverprofile=$FULL_INTEGRATION_TEST_CODECOV_PATH \
            ./test/not_gated/...
result=$?

if [[ -z "${CI_PROJECT_DIR}" ]] ; then
    CI_PROJECT_DIR="$(pwd)/cicd/integration"
fi

mkdir -p $CI_PROJECT_DIR/coverage-integration-ml
mv $FULL_INTEGRATION_TEST_CODECOV_PATH $CI_PROJECT_DIR/coverage-integration-ml/coverage.out

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