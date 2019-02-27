#!/bin/bash

echo "==============================================="
echo "Beginning functional tests"
echo "==============================================="
echo "SPLUNK_CLOUD_HOST=$SPLUNK_CLOUD_HOST"
echo "TENANT_ID=$TENANT_ID"
echo "==============================================="

COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES=$(go list ./... | grep -v test | awk -v ORS=, '{ print $1 }' | sed 's/,$//')

PACKAGE_COVERAGE_PREFIX=./ci/functional/
FULL_FUNCTIONAL_TEST_CODECOV_FILE_NAME=functional_test_codecov.out
FULL_FUNCTIONAL_TEST_CODECOV_PATH=$PACKAGE_COVERAGE_PREFIX$FULL_FUNCTIONAL_TEST_CODECOV_FILE_NAME

# Required to run just the service tests
if [ "$allow_failures" = "1" ]; then
    echo "Running functional tests but not gating on failures..."
    set +e
    go test -v -coverpkg $COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES \
               -covermode=count \
               -coverprofile=$FULL_FUNCTIONAL_TEST_CODECOV_PATH \
               ./test/stubby_integration/...
else
    echo "Running functional tests and gating on failures..."
    go test -v -coverpkg $COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES \
               -covermode=count \
               -coverprofile=$FULL_FUNCTIONAL_TEST_CODECOV_PATH \
               ./test/stubby_integration/... \
        || exit 1
fi

# Upload code cov report
echo "TODO: CODE COVERAGE IS NOT CURRENTLY SUPPORTED! CODECOV REPORT WILL NOT BE UPLOADED."