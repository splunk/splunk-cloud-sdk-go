#!/bin/bash -e

# Run from ./cicd/docs/../.. directory (root of repo)
cd "$(dirname "$0")/../.."

GO111MODULE=off go get github.com/robertkrimen/godocdown/godocdown
GO_NON_TEST_NON_EXAMPLE_PACKAGES=$(go list ./... | grep -v scloud | grep -v test | grep -v examples)

rm -rf docs/
mkdir -p docs/pkg

DOC_INDEX="

#  github.com/splunk/splunk-cloud-sdk-go

## Packages

"

for PACKAGE in $GO_NON_TEST_NON_EXAMPLE_PACKAGES
do
    DOC_FILE=$(echo $PACKAGE | sed "s/[\.|\/|-]/_/g").md
    godocdown $PACKAGE > docs/pkg/$DOC_FILE
    echo "Wrote docs for $PACKAGE to docs/pkg/$DOC_FILE"
    DOC_INDEX+="* [$PACKAGE](pkg/$DOC_FILE)
"
done

echo "$DOC_INDEX" > docs/README.md
echo "Wrote docs index to docs/README.md"