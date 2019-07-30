#!/bin/bash

echo "==============================================="
echo "Beginning scloud integration tests"
echo "==============================================="
echo "TEST_USERNAME=$TEST_USERNAME"
echo "TEST_SCLOUD_TENANT=$TEST_SCLOUD_TENANT"
echo "==============================================="

COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES=$(go list ./... | grep -v test | awk -v ORS=, '{ print $1 }' | sed 's/,$//')

PACKAGE_COVERAGE_PREFIX=./cicd/integration/
FULL_INTEGRATION_TEST_CODECOV_FILE_NAME=integration_test_scloud_codecov.out
FULL_INTEGRATION_TEST_CODECOV_PATH=$PACKAGE_COVERAGE_PREFIX$FULL_INTEGRATION_TEST_CODECOV_FILE_NAME

if [[ "$allow_failures" == "1" ]]; then
    echo "Running integration tests but not gating on failures..."
else
    echo "Running integration tests and gating on failures..."
fi

set +e
go test -v -coverpkg $COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES \
            -covermode=count \
            -coverprofile=$FULL_INTEGRATION_TEST_CODECOV_PATH \
            github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cli \
            -timeout 10m \
            -run '^(TestScloudBinaryWithCoverage)$'
result=$?

if [[ -z "${CI_PROJECT_DIR}" ]] ; then
    CI_PROJECT_DIR="$(pwd)/cicd/integration"
fi

mkdir -p $CI_PROJECT_DIR/coverage-integration-scloud
mv $FULL_INTEGRATION_TEST_CODECOV_PATH $CI_PROJECT_DIR/coverage-integration-scloud/coverage.out

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