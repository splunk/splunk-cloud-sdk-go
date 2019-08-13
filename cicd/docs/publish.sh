#!/bin/bash

####################################################################################################
# Check for required env vars
####################################################################################################
if [[ "${CI}" != "true" ]] ;  then
    echo "Exiting: $0 can only be run from the CI system."
    exit 1
fi
if [[ -z "${ARTIFACT_USERNAME}" ]] ; then
    echo "ARTIFACT_USERNAME must be set, exiting ..."
    exit 1
fi
if [[ -z "${ARTIFACT_PASSWORD}" ]] ; then
    echo "ARTIFACT_PASSWORD must be set, exiting ..."
    exit 1
fi
if [[ -z "${ARTIFACTORY_NPM_REGISTRY}" ]] ; then
    echo "ARTIFACTORY_NPM_REGISTRY must be set, exiting ..."
    exit 1
fi

####################################################################################################
# Set platform-specific sed extended syntax flag
####################################################################################################
if [[ "$(uname)" == "Darwin" ]] ; then
    # MacOS
    SED_FLG="-E"
else
    # Linux
    SED_FLG="-r"
fi

####################################################################################################
# Get release version from services/client_info.go e.g. 0.9.2
####################################################################################################
NEW_VERSION=$(cat services/client_info.go | sed ${SED_FLG} -n 's/const Version = "([0-9]+\.[0-9]+\.[0-9]+.*)"/\1/p')
if [[ -z "${NEW_VERSION}" ]] ; then
    echo "error setting NEW_VERSION from services/client_info.go, version must be set to match: const Version = \"([0-9]+\.[0-9]+\.[0-9]+.*)\" (e.g. const Version = \"0.8.3\") but format found is:\n\n$(cat services/client_info.go)\n\n..."
    exit 1
fi

echo "Publishing docs for v${NEW_VERSION} ..."

# Run from ci/docs directory
cd "$(dirname "$0")"
####################################################################################################
# Write a package.json file, this is needed by @splunk/cicd-tools even though this isn't js
####################################################################################################
rm package.json
PACKAGE_JSON="{
    \"name\": \"@splunk/splunk-cloud-sdk-go\",
    \"version\": \"${NEW_VERSION}\",
    \"description\": \"Splunk Cloud SDK for Go\"
}"
echo ${PACKAGE_JSON} > package.json
rm -rf build/
####################################################################################################
# Install @splunk/cicd-tools from artifactory
####################################################################################################
npm add --no-save @splunk/cicd-tools --registry "${ARTIFACTORY_NPM_REGISTRY}"
####################################################################################################
# Publish docs to artifactory with @splunk/cicd-tools/cicd-publish-docs
####################################################################################################
npx cicd-publish-docs --force ../../docs/
echo "Docs built and packaged into $(dirname "$0")/build"