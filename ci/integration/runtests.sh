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

# Required to run just the service tests
if [ "$allow_failures" -eq "1" ]; then
    echo "Running integration tests but not gating on failures..."
    set +e
    go test -v -cover -covermode=count -coverprofile="codecov.integration.out" ./test/playground_integration/...
    exit 0
else
    echo "Running integration tests and gating on failures..."
    go test -v -cover -covermode=count -coverprofile="codecov.integration.out" ./test/playground_integration/... || exit 1
fi

# Upload coverage information
./ci/codecov -f "codecov.integration.out" -F integration

