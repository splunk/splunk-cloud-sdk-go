#!/bin/bash

go get github.com/robertkrimen/godocdown/godocdown
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
    godocdown $PACKAGE > docs/pkg/$DOC_FILE
    # Escape < and > to &lt; and &gt; per portal requirements
    if [ "$(uname)" == "Darwin" ]; then
        # MacOS
        sed -E -i '' -e "s/\</\&lt\;/g" -e "s/\>/\&gt\;/g" docs/pkg/$DOC_FILE
    else
        # Linux
        sed -r -i '' -e "s/\</\&lt\;/g" -e "s/\>/\&gt\;/g" docs/pkg/$DOC_FILE
    fi
    echo "Wrote docs for $PACKAGE to docs/pkg/$DOC_FILE"
    DOC_INDEX+="* [$PACKAGE](pkg/$DOC_FILE)
"
done

echo "$DOC_INDEX" > docs/README.md
echo "Wrote docs index to docs/README.md"