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
# Convert the release candidate tag (e.g. v0.5.11rc2) -> NEW_RELEASE (e.g. 0.5.11)
####################################################################################################
NEW_VERSION=$(echo "${CI_COMMIT_REF_NAME}" | sed "${SED_FLG}" -n 's/^v([0-9]+\.[0-9]+\.[0-9]+.*)rc.?/\1/p')
if [[ -z "${NEW_VERSION}" ]] ; then
    echo "error setting NEW_VERSION, the release candidate tag must match this pattern: \"^v([0-9]+\.[0-9]+\.[0-9]+.*)rc.?\" (e.g. v0.5.11rc1) but tag is: ${CI_COMMIT_REF_NAME} ..."
    exit 1
fi

####################################################################################################
# Checkout latest version of Gitlab code to prepare for release
####################################################################################################
OWNER=devplat-pr-32
PROJECT=splunk-cloud-sdk-go
BRANCH_NAME=develop
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
# Re-generate the latest interface.go files for each service
####################################################################################################
echo "Preparing v${NEW_VERSION} for release ..."
####################################################################################################
# Update services/client_info.go using the pre-release tag e.g. 0.9.1 for v0.9.1rc2
####################################################################################################
echo "Updating Version in services/client_info.go using sed \"${SED_FLG}\" -i '' -e \"s/\\"[0-9]+\.[0-9]+\.[0-9]+.*\\"/\\"${NEW_VERSION}\\"/g\" services/client_info.go ..."
sed "${SED_FLG}" -i -e "s/\"[0-9]+\.[0-9]+\.[0-9]+.*\"/\"${NEW_VERSION}\"/g" services/client_info.go
git add services/client_info.go
####################################################################################################
# Re-generate the latest docs/
####################################################################################################
echo "Updating docs ..."
make build
make docs_md
####################################################################################################
# Append the message from the pre-release tag to the top of the CHANGELOG.md under ## Version x.y.z
####################################################################################################
echo "Adding tag message to changelog ..."
set +x
TAG_MESSAGE=$(git cat-file -p "${CI_COMMIT_REF_NAME}" | tail -n +6)
set -x
printf "Adding release notes to top of CHANGELOG.md:\n\n${TAG_MESSAGE}\n\n"
set +x
CL_HEADER=$(head -n 1 CHANGELOG.md)
CL_CONTENTS=$(tail -n +2 CHANGELOG.md)
rm CHANGELOG.md
printf "${CL_HEADER}\n\n## Version ${NEW_VERSION}\n${TAG_MESSAGE}\n${CL_CONTENTS}" > CHANGELOG.md
set -x
git add CHANGELOG.md
####################################################################################################
# Adding/pushing changed files to develop (client_info.go, CHANGELOG.md, */interface.go, and docs/)
####################################################################################################
echo "Showing changes with \"git status\" ..."
git status
echo "Creating commit for client_info.go, CHANGELOG.md, */interface.go, and docs/ changes ..."
git commit -m "Prepare v${NEW_VERSION} for release"
echo "Pushing branch ${BRANCH_NAME} ..."
git push origin "${BRANCH_NAME}"
