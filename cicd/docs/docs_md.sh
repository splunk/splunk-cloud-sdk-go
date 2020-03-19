#!/bin/bash -e

# Run from ./cicd/docs/../.. directory (root of repo)
cd "$(dirname "$0")/../.."

GO111MODULE=off go get github.com/robertkrimen/godocdown/godocdown
GO_NON_TEST_NON_EXAMPLE_PACKAGES=$(go list ./... | grep -v test | grep -v examples)

rm -rf docs/
mkdir -p docs/pkg

DOC_INDEX="

#  github.com/splunk/splunk-cloud-sdk-go

## Packages

"

for PACKAGE in $GO_NON_TEST_NON_EXAMPLE_PACKAGES
do
    DOC_FILE=$(echo $PACKAGE | sed "s/[\.|\/|-]/_/g").md
    PKG_LOC=$(echo $PACKAGE | sed "s/github\.com\/splunk\/splunk-cloud-sdk-go/./g")
    echo "running $GOPATH/bin/godocdown $PKG_LOC > docs/pkg/$DOC_FILE ..."
    # escape . and / characters for replacing in the next step
    PACKAGE_ESC=$(echo $PACKAGE | sed 's|\.|\\.|g')
    $GOPATH/bin/godocdown $PKG_LOC | sed "s|import \"\.\"|import \"$PACKAGE_ESC\"|g" > docs/pkg/$DOC_FILE
    echo "Wrote docs for $PACKAGE to docs/pkg/$DOC_FILE"
    echo ""
    DOC_INDEX+="* [$PACKAGE](pkg/$DOC_FILE)
"
done

echo "$DOC_INDEX" > docs/README.md
echo "Wrote docs index to docs/README.md"