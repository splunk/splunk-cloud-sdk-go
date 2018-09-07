#!/bin/bash

echo "==============================================="
echo "Running examples"
echo "==============================================="
env | grep TEST_SPLUNK_CLOUD_HOST
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


if [ "$allow_failures" -eq "1" ]; then
    echo "Running examples but not gating on failures..."
    set +e
    go run -v  ./examples/ingestSearch.go
else
    echo "Running examples and gating on failures..."
    go run -v ./examples/ingestSearch.go \
        || exit 1
fi
