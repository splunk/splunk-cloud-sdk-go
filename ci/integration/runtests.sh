#!/bin/bash

FULL_PATH=$(cd "$(dirname "$0")"; pwd)

source $FULL_PATH/okta.sh
if [ "$allow_failures" -eq "1" ]; then
    echo "Running integration tests but not gating on failures..."
    set +e
    go test -v -covermode=count -coverprofile="codecov.integration.out" -run ^TestIntegration* ./...
else
    echo "Running integration tests and gating on failures..."
    go test -v -covermode=count -coverprofile="codecov.integration.out" -run ^TestIntegration* ./... || exit 1
fi
$FULL_PATH/../codecov -f "codecov.integration.out" -F integration