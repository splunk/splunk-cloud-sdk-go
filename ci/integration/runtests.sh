#!/bin/bash

echo "==============================================="
echo "Beginning integration tests"
echo "==============================================="
env | grep TEST_SSC_HOST
env | grep TEST_URL_PROTOCOL
env | grep TEST_TENANT_ID

# Get the BEARER_TOKEN setup
TOKEN_FILE="./.config/okta/token"
if [ -f $TOKEN_FILE ]; then
    echo "Token found in $TOKEN_FILE"
    TEST_BEARER_TOKEN=$(cat $TOKEN_FILE)
else
    echo "Token was not set to $TOKEN_FILE"
    exit 1
fi

TENANT_FILE="./.config/ssc/tenant_id"
if [ -f $TENANT_FILE ]; then
    echo "Tenant found in $TENANT_FILE"
    TEST_TENANT_ID=$(cat $TENANT_FILE)
else
    echo "Tenant was not set to $TENANT_FILE"
    exit 1
fi

echo "==============================================="

if [ "$allow_failures" -eq "1" ]; then
    echo "Running integration tests but not gating on failures..."
    set +e
    go test -v -covermode=count -coverprofile="codecov.integration.out" ./test/playground_integration/...
    exit 0
else
    echo "Running integration tests and gating on failures..."
    go test -v -covermode=count -coverprofile="codecov.integration.out" ./test/playground_integration/... || exit 1
fi

# Upload coverage information
../codecov -f "codecov.integration.out" -F integration

