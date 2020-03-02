#!/bin/bash -e

./bin/scloud_gen config set --key env --value $TEST_ENVIRONMENT_1
./bin/scloud_gen config set --key username --value $TEST_USERNAME
./bin/scloud_gen config set --key tenant --value $TEST_SCLOUD_TENANT
./bin/scloud_gen login --pwd $TEST_PASSWORD_1