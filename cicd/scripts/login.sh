#!/bin/bash -e

./bin/scloud set env staging
./bin/scloud set username $TEST_USERNAME
./bin/scloud set tenant $TEST_SCLOUD_TENANT
./bin/scloud -no-prompt -p $TEST_PASSWORD login