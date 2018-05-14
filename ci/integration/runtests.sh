#!/bin/bash

# source ./ci/integration/okta.sh

CONFIG_FILE="./okta/.token"
if [ -f $CONFIG_FILE ]; then
    echo "Token found in $CONFIG_FILE"
    BEARER_TOKEN=$(cat $CONFIG_FILE)
else
    echo "Token was not set to $CONFIG_FILE"
    exit 1
fi

cd service # TODO: this shouldn't be necessary...
if [ "$allow_failures" -eq "1" ]; then
    echo "Running integration tests but not gating on failures..."
    set +e
    go test -v -run ^TestIntegration*
    exit 0
else
    echo "Running integration tests and gating on failures..."
    go test -v -run ^TestIntegration* || exit 1
fi