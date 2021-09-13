#!/bin/bash

echo "==============================================="
echo "Running examples"
echo "==============================================="
echo "SPLUNK_CLOUD_HOST_TENANT_SCOPED=$SPLUNK_CLOUD_HOST_TENANT_SCOPED"
echo "TEST_TENANT_SCOPED=$TEST_TENANT_SCOPED"
echo "==============================================="

COMMA_SEPARATED_FULLY_QUALIFIED_PACKAGES=$(go list ./... | grep -v test | awk -v ORS=, '{ print $1 }' | sed 's/,$//')

if [[ "$allow_failures" == "1" ]]; then
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
    echo "Running mock ..."
    go run -v ./examples/mock/mock.go || exit 1
fi
