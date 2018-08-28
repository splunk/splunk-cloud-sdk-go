#!/bin/bash
# Run from ci/docs directory
cd "$(dirname "$0")"
npm add -D @splunk/cicd-tools --registry https://repo.splunk.com/artifactory/api/npm/npm
yarn cicd-publish-docs ../../docs/
