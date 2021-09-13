#!/bin/bash

echo "==============================================="
echo "Beginning integration tests"
echo "==============================================="
echo "TEST_USERNAME=$TEST_USERNAME"
echo "SPLUNK_CLOUD_HOST_TENANT_SCOPED=$SPLUNK_CLOUD_HOST_TENANT_SCOPED"
echo "TEST_TENANT_SCOPED=$TEST_TENANT_SCOPED"
echo "==============================================="

COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES=$(go list ./... | grep -v test | awk -v ORS=, '{ print $1 }' | sed 's/,$//')

PACKAGE_COVERAGE_PREFIX=./cicd/integration/
FULL_INTEGRATION_TEST_CODECOV_FILE_NAME=integration_test_codecov.out
FULL_INTEGRATION_TEST_CODECOV_PATH=$PACKAGE_COVERAGE_PREFIX$FULL_INTEGRATION_TEST_CODECOV_FILE_NAME

# Required to run just the service tests
if [[ "$allow_failures" == "1" ]]; then
    echo "Running integration tests but not gating on failures..."
else
    echo "Running integration tests and gating on failures..."
fi

set +e
gotestsum --format short-verbose \
          -- -coverpkg $COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES \
             -covermode=count \
             -coverprofile=$FULL_INTEGRATION_TEST_CODECOV_PATH \
             -timeout 20m \
             ./test/integration/...
result=$?

if [[ -z "${CI_PROJECT_DIR}" ]] ; then
    CI_PROJECT_DIR="$(pwd)/cicd/integration"
fi

mkdir -p $CI_PROJECT_DIR/coverage-integration
mv $FULL_INTEGRATION_TEST_CODECOV_PATH $CI_PROJECT_DIR/coverage-integration/coverage.out

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
