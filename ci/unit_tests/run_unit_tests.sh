#!/bin/bash

echo "==============================================="
echo "Beginning unit tests"
echo "==============================================="

GO_NON_TEST_NON_VENDOR_PACKAGES=$(go list ./... | grep -v /vendor/ | grep -v test)
#GO_NON_TEST_NON_VENDOR_PACKAGE_NAMES=$(go list -f "{{.Name}}" ./... | grep -v /vendor/ | grep -v test)
#echo $GO_NON_TEST_NON_VENDOR_PACKAGE_NAMES

PACKAGE_COVERAGE_PREFIX=./ci/unit_tests/
PACKAGE_COVERAGE_SUFFIX=_unit_test_code_cov.out
FULL_UNIT_TEST_CODECOV_FILE_NAME=unit_test_codecov.out
FULL_UNIT_TEST_CODECOV_PATH=$PACKAGE_COVERAGE_PREFIX$FULL_UNIT_TEST_CODECOV_FILE_NAME

echo "------------------------------------------------------------------------------"
echo "Unit tests will be output to $FULL_UNIT_TEST_CODECOV_PATH"
echo "------------------------------------------------------------------------------"

for PACKAGE in $GO_NON_TEST_NON_VENDOR_PACKAGES
do
    echo "-------------------------------------------------------------------"
    echo "Beginning unit tests for $PACKAGE"
    echo "-------------------------------------------------------------------"
    SANITIZED_PACKAGE_NAME=$(echo $PACKAGE | sed "s/[\.|\/|-]/_/g")
    COVERAGE_PACKAGE_OUTPUT_FILE=$PACKAGE_COVERAGE_PREFIX$SANITIZED_PACKAGE_NAME$PACKAGE_COVERAGE_SUFFIX

    # echo $SANITIZED_PACKAGE_NAME
    # echo $COVERAGE_PACKAGE_OUTPUT_FILE

    go test -v -covermode=count -coverprofile=$COVERAGE_PACKAGE_OUTPUT_FILE $PACKAGE

    if [ -f $COVERAGE_PACKAGE_OUTPUT_FILE ]; then
        cat $COVERAGE_PACKAGE_OUTPUT_FILE >> $FULL_UNIT_TEST_CODECOV_PATH
    else
        echo "No unit test results were found for $PACKAGE"
    fi
done


echo "==============================================="
echo "PERFORMING COMMAND"
echo "./ci/codecov -f "codecov.integration.out" -F unit -t $CODECOV_TOKEN"
echo "==============================================="

if [[ -z "$CODECOV_TOKEN" ]];
then
    echo "THE CODE COVERAGE TOKEN IS NOT SET! CODECOV REPORT WILL NOT BE UPLOADED."
else
    # Upload coverage information
    ./ci/codecov -f $FULL_UNIT_TEST_CODECOV_PATH -F unit -t $CODECOV_TOKEN
fi
