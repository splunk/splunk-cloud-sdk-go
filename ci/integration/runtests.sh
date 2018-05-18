#!/bin/bash

FULL_PATH_OF_DIRECTORY_CONTAINING_THIS_SCRIPT=$(cd "$(dirname "$0")"; pwd)

# Get the BEARER_TOKEN setup
CONFIG_FILE="./okta/.token"
if [ -f $CONFIG_FILE ]; then
    echo "Token found in $CONFIG_FILE"
    BEARER_TOKEN=$(cat $CONFIG_FILE)
else
    echo "Token was not set to $CONFIG_FILE"
    exit 1
fi

cd service # Required to run just the service tests
if [ "$allow_failures" -eq "1" ]; then
    echo "Running integration tests but not gating on failures..."
    set +e
    go test -v -covermode=count -coverprofile="codecov.integration.out" -run ^TestIntegration* ./...
else
    echo "Running integration tests and gating on failures..."
    go test -v -covermode=count -coverprofile="codecov.integration.out" -run ^TestIntegration* ./... || exit 1
fi

# Upload coverage information
$FULL_PATH_OF_DIRECTORY_CONTAINING_THIS_SCRIPT/../codecov -f "codecov.integration.out" -F integration