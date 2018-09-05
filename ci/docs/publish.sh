#!/bin/bash
# Run from ci/docs directory
cd "$(dirname "$0")"
rm package.json
PACKAGE_JSON="{
    \"name\": \"@splunk/splunk-cloud-sdk-go\",
    \"version\": \"$(git describe --abbrev=0 --tags | sed 's/v//')\",
    \"description\": \"Splunk Cloud SDK for Go\"
}"
echo $PACKAGE_JSON > package.json
rm -rf build/
npm add -D @splunk/cicd-tools --registry https://repo.splunk.com/artifactory/api/npm/npm
yarn cicd-publish-docs ../../docs/
echo "Docs built and packaged into $(dirname "$0")/build"