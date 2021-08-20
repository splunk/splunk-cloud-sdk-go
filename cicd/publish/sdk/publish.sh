#!/bin/bash -e

####################################################################################################
# Check for required env vars
####################################################################################################
if [[ "${CI}" != "true" ]] ;  then
    echo "Exiting: $0 can only be run from the CI system."
    exit 1
fi
if [[ -z "${GITLAB_HOST}" ]] ; then
    echo "GITLAB_HOST must be set, exiting ..."
    exit 1
fi
if [[ -z "${GITLAB_TOKEN}" ]] ; then
    echo "GITLAB_TOKEN must be set, exiting ..."
    exit 1
fi
if [[ -z "${GITHUB_TOKEN}" ]] ; then
    echo "GITHUB_TOKEN must be set, exiting ..."
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

####################################################################################################
# Get release notes for version x.y.z from CHANGELOG.md
####################################################################################################
echo "Retrieving CHANGELOG.md entry for ${NEW_VERSION} ..."
set +x
RELEASE_NOTES=$(sed -n "/Version ${NEW_VERSION}/,/Version/p" CHANGELOG.md | sed '$ d')
# Sanitize this to escape double quotes and newlines such that it is valid json
RELEASE_NOTES=$(printf "$RELEASE_NOTES" | tr -d '\r' | sed -e 's/\"/\\"/g' | sed -e ':a' -e 'N' -e '$!ba' -e 's/\n/\\n/g')
set -x
echo "Changelog entry found: ${RELEASE_NOTES}"

####################################################################################################
# Checkout master branch in Gitlab
####################################################################################################
OWNER=devplat-pr-32
PROJECT=splunk-cloud-sdk-go
BRANCH_NAME=master
set +x
git remote set-url origin "https://oauth2:${GITLAB_TOKEN}@${GITLAB_HOST}/${OWNER}/${PROJECT}.git"
set -x
git config user.email "srv-dev-platform@splunk.com"
git config user.name "srv-dev-platform"
echo "Running \"git checkout ${BRANCH_NAME}\" ..."
git checkout "${BRANCH_NAME}"
echo "Running \"git fetch --all && git pull --all\" ..."
git fetch --all && git pull --all
####################################################################################################
# Tag master branch in Gitlab
####################################################################################################
echo "Running git tag/push for v${NEW_VERSION} on ${BRANCH_NAME} ..."
git tag -a -s "v${NEW_VERSION}" -m "${RELEASE_NOTES}"
git push origin "v${NEW_VERSION}"
SPLUNK_SHA=$(git rev-parse HEAD)

####################################################################################################
# Create a release for splunk Github (commit should be pushed automatically)
####################################################################################################
OWNER=splunk
PROJECT=splunk-cloud-sdk-go
echo "Creating a release on ${OWNER}/${PROJECT} ..."
set +x
RELEASE_RESPONSE=$(curl -s -H "Authorization: token $GITHUB_TOKEN" "https://api.github.com/repos/${OWNER}/${PROJECT}/releases" -d "{\"tag_name\": \"v${NEW_VERSION}\", \"target_commitish\": \"${SPLUNK_SHA}\", \"name\": \"Release v${NEW_VERSION}\", \"body\": \"${RELEASE_NOTES}\"}")
RELEASE_EXIT_CODE=$?
set -x
echo "Release response: ${RELEASE_RESPONSE}"
