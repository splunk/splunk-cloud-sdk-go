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

COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES=$(go list ./... | grep -v /vendor/ | grep -v test | awk -v ORS=, '{ print $1 }' | sed 's/,$//')

PACKAGE_COVERAGE_PREFIX=./ci/integration/
FULL_INTEGRATION_TEST_CODECOV_FILE_NAME=integration_test_codecov.out
FULL_INTEGRATION_TEST_CODECOV_PATH=$PACKAGE_COVERAGE_PREFIX$FULL_INTEGRATION_TEST_CODECOV_FILE_NAME

# Required to run just the service tests
if [ "$allow_failures" -eq "1" ]; then
    echo "Running integration tests but not gating on failures..."
    set +e
    go test -v -coverpkg $COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES \
               -covermode=count \
               -coverprofile=$FULL_INTEGRATION_TEST_CODECOV_PATH \
               ./test/playground_integration/...
else
    echo "Running integration tests and gating on failures..."
    go test -v -coverpkg $COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES \
               -covermode=count \
               -coverprofile=$FULL_INTEGRATION_TEST_CODECOV_PATH \
               ./test/playground_integration/...
fi

echo "==============================================="
echo "PERFORMING COMMAND"
echo "./ci/codecov -f $FULL_INTEGRATION_TEST_CODECOV_PATH -F integration -t $CODECOV_TOKEN"
echo "==============================================="

if [[ -z "$CODECOV_TOKEN" ]];
then
    echo "THE CODE COVERAGE TOKEN IS NOT SET! CODECOV REPORT WILL NOT BE UPLOADED."
else
    # Upload coverage information
    ./ci/codecov -f $FULL_INTEGRATION_TEST_CODECOV_PATH -F integration -t $CODECOV_TOKEN
fi
