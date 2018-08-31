#!/bin/bash

go get github.com/robertkrimen/godocdown
GO_NON_TEST_PACKAGES=$(go list ./... | grep -v test)

rm -rf docs/
mkdir -p docs/pkg

DOC_INDEX="

#  github.com/splunk/ssc-client-go

## Packages

"

for PACKAGE in $GO_NON_TEST_PACKAGES
do
    DOC_FILE=$(echo $PACKAGE | sed "s/[\.|\/|-]/_/g").md
    godocdown $PACKAGE > docs/pkg/$DOC_FILE
    echo "Wrote docs for $PACKAGE to docs/pkg/$DOC_FILE"
    DOC_INDEX+="* [$PACKAGE](pkg/$DOC_FILE)
"
done

echo "$DOC_INDEX" > docs/README.md
echo "Wrote docs index to docs/README.md"