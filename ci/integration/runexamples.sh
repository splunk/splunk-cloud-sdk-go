#!/bin/bash

echo "==============================================="
echo "Running examples"
echo "==============================================="
env | grep SPLUNK_CLOUD_HOST
env | grep URL_PROTOCOL
env | grep TENANT_ID

echo "==============================================="

# Get the BEARER_TOKEN setup
CONFIG_FILE="./okta/.token"
if [ -f $CONFIG_FILE ]; then
    echo "Token found in $CONFIG_FILE"
    export BEARER_TOKEN=$(cat $CONFIG_FILE)
else
    echo "Token was not set to $CONFIG_FILE"
    exit 1
fi

COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES=$(go list ./... | grep -v test | awk -v ORS=, '{ print $1 }' | sed 's/,$//')

if [ "$allow_failures" -eq "1" ]; then
    echo "Running examples but not gating on failures..."
    echo ""
    echo "Running ingestSearch ..."
    go run -v  ./examples/ingestSearch/ingestSearch.go || exit 0
    echo "Running logging ..."
    go run -v  ./examples/logging/logging.go -logfile example.log || exit 0
    echo "example.log output:"
    (ls example.log && cat example.log| sed -e "s/Authorization: Bearer .*/Authorization: Bearer <REDACTED>/g") || exit 0
else
    echo "Running examples and gating on failures..."
    set +e
    echo ""
    echo "Running ingestSearch ..."
    go run -v  ./examples/ingestSearch/ingestSearch.go || exit 1
    echo "Running logging ..."
    go run -v  ./examples/logging/logging.go -logfile example.log || exit 1
    echo "example.log output:"
    (ls example.log && cat example.log| sed -e "s/Authorization: Bearer .*/Authorization: Bearer <REDACTED>/g") || exit 1
fi
