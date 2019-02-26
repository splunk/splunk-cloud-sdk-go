#!/bin/bash

echo "==============================================="
echo "Beginning unit tests"
echo "==============================================="

GO_NON_TEST_PACKAGES=$(go list ./... | grep -v test)

PACKAGE_COVERAGE_PREFIX=./ci/unit_tests/
PACKAGE_COVERAGE_SUFFIX=_unit_test_code_cov.out
FULL_UNIT_TEST_CODECOV_FILE_NAME=unit_test_codecov.out
FULL_UNIT_TEST_CODECOV_PATH=$PACKAGE_COVERAGE_PREFIX$FULL_UNIT_TEST_CODECOV_FILE_NAME

echo "------------------------------------------------------------------------------"
echo "Unit tests will be output to $FULL_UNIT_TEST_CODECOV_PATH"
echo "------------------------------------------------------------------------------"

for PACKAGE in $GO_NON_TEST_PACKAGES
do
    echo "-------------------------------------------------------------------"
    echo "Beginning unit tests for $PACKAGE"
    echo "-------------------------------------------------------------------"
    SANITIZED_PACKAGE_NAME=$(echo $PACKAGE | sed "s/[\.|\/|-]/_/g")
    COVERAGE_PACKAGE_OUTPUT_FILE=$PACKAGE_COVERAGE_PREFIX$SANITIZED_PACKAGE_NAME$PACKAGE_COVERAGE_SUFFIX

    # echo $SANITIZED_PACKAGE_NAME
    # echo $COVERAGE_PACKAGE_OUTPUT_FILE

    go test -v -covermode=count -coverprofile=$COVERAGE_PACKAGE_OUTPUT_FILE $PACKAGE

    RESULT=$?
    if [ $RESULT -ne 0 ]
        then
            echo "There was an error testing the $PACKAGE package's unit tests."
            exit 1
    fi

    if [ -f $COVERAGE_PACKAGE_OUTPUT_FILE ]; then
        cat $COVERAGE_PACKAGE_OUTPUT_FILE >> $FULL_UNIT_TEST_CODECOV_PATH
    else
        echo "No unit test results were found for $PACKAGE"
    fi
done

# Upload code cov report
echo "TODO: CODE COVERAGE IS NOT CURRENTLY SUPPORTED! CODECOV REPORT WILL NOT BE UPLOADED."