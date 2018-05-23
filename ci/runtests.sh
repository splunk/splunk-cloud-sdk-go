#!/bin/bash

go test -v -covermode=count -coverprofile="codecov.out"  .test/stubby_integration/... || exit 1

ci/codecov  -f "codecov.out"  -F test