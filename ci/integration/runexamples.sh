#!/bin/bash

source ./ci/integration/okta.sh

echo "==============================================="
echo "Beginning integration tests"
echo "==============================================="
env | grep TEST_SSC_HOST
env | grep TEST_URL_PROTOCOL
env | grep TEST_TENANT_ID

echo "==============================================="

# Get the BEARER_TOKEN setup
CONFIG_FILE="./okta/.token"
if [ -f $CONFIG_FILE ]; then
    echo "Token found in $CONFIG_FILE"
    TEST_BEARER_TOKEN=$(cat $CONFIG_FILE)
else
    echo "Token was not set to $CONFIG_FILE"
    exit 1
fi

COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES=$(go list ./... | grep -v test | awk -v ORS=, '{ print $1 }' | sed 's/,$//')

PACKAGE_COVERAGE_PREFIX=./ci/integration/
FULL_INTEGRATION_TEST_CODECOV_FILE_NAME=integration_test_codecov.out
FULL_INTEGRATION_TEST_CODECOV_PATH=$PACKAGE_COVERAGE_PREFIX$FULL_INTEGRATION_TEST_CODECOV_FILE_NAME

if [ "$allow_failures" -eq "1" ]; then
    echo "Running examples but not gating on failures..."
    set +e
    go run -v  ./examples/ingestSearch.go
else
    echo "Running examples and gating on failures..."
    go run -v ./examples/ingestSearch.go \
        || exit 1
fi
