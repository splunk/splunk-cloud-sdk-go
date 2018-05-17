#!/bin/bash

FULL_PATH_OF_DIRECTORY_CONTAINING_THIS_SCRIPT=$(cd "$(dirname "$0")"; pwd)

source $FULL_PATH_OF_DIRECTORY_CONTAINING_THIS_SCRIPT/okta.sh
if [ "$allow_failures" -eq "1" ]; then
    echo "Running integration tests but not gating on failures..."
    set +e
    go test -v -covermode=count -coverprofile="codecov.integration.out" -run ^TestIntegration* ./...
else
    echo "Running integration tests and gating on failures..."
    go test -v -covermode=count -coverprofile="codecov.integration.out" -run ^TestIntegration* ./... || exit 1
fi
$FULL_PATH_OF_DIRECTORY_CONTAINING_THIS_SCRIPT/../codecov -f "codecov.integration.out" -F integration