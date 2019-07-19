#!/bin/bash -e

TARGET_ROOT_DIR=bin/cross-compiled
ARCHIVE_DIR=${TARGET_ROOT_DIR}/archive

if [ "${CI}" != "true" ] ; then
    echo "Exiting: $0 can only be run from the CI system."
    exit 1
fi

if [ -z "${ARTIFACTORY_TOKEN}" ] ; then
    echo "Exiting: $0 no \$ARTIFACTORY_TOKEN set"
    exit 1
fi

for pkg in ${ARCHIVE_DIR[@]}/*;
do
    echo "Uploading ${pkg} to artifactory.";
    curl -H 'Content-Type:text/plain' -H "X-JFrog-Art-Api: ${ARTIFACTORY_TOKEN}" -X PUT "https://repo.splunk.com/artifactory/Solutions/scloud/" -T ${pkg}
done