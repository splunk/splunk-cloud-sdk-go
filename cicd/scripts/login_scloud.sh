#!/bin/bash -e

./bin/scloud config set --key env --value $TEST_ENVIRONMENT_1
./bin/scloud config set --key username --value $TEST_USERNAME
./bin/scloud config set --key tenant --value $TEST_SCLOUD_TENANT
./bin/scloud login --pwd $TEST_PASSWORD_1