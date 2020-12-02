#!/bin/bash -e

if [[ "$(uname)" == "Darwin" ]] ; then
    # MacOS
    SED_FLG="-E"
else
    # Linux
    SED_FLG="-r"
fi

echo "Publishing changes to homebrew"

####################################################################################################
# Obtain release tag, scloud version, SHA and URL
####################################################################################################
# Release Tag Example: v1.0.0
RELEASE_TAG=v$(cat services/client_info.go | sed ${SED_FLG} -n 's/const Version = "([0-9]+\.[0-9]+\.[0-9]+.*)"/\1/p')
if [ -n "${OVERRIDE_RELEASE_TAG}" ] ; then
   echo "\$OVERRIDE_RELEASE_TAG was set so uploading cross-compiled artifacts to ${OVERRIDE_RELEASE_TAG} rather than the default for this tag (${RELEASE_TAG}) ..."
   RELEASE_TAG="${OVERRIDE_RELEASE_TAG}"
fi
echo "release tag is: $RELEASE_TAG"

# Scloud Version Example: 1.0.0
SCLOUD_VERSION=$(cat cmd/scloud/cmd/scloud/version/client_info.go | sed ${SED_FLG} -n 's/const ScloudVersion = "([0-9]+\.[0-9]+\.[0-9]+.*)"/\1/p')
if [[ -z "${SCLOUD_VERSION}" ]] ; then
    echo "error setting SCLOUD_VERSION from cmd/scloud/cmd/scloud/version/client_info.go, version must be set to match: const ScloudVersion = \"([0-9]+\.[0-9]+\.[0-9]+.*)\" (e.g. const ScloudVersion = \"0.8.3\") but format found is:\n\n$(cat cmd/scloud/cmd/scloud/version/client_info.go)\n\n..."
    exit 1
fi
echo "scloud version is: $SCLOUD_VERSION"

# MAC URL
MAC_URL='\"https:\/\/github.com\/splunk\/splunk-cloud-sdk-go\/releases\/download\/'${RELEASE_TAG}'\/scloud_v'${SCLOUD_VERSION}'_darwin_amd64.tar.gz\"'
echo "MAC_URL is: $MAC_URL"

# LINUX URL
LINUX_URL='\"https:\/\/github.com\/splunk\/splunk-cloud-sdk-go\/releases\/download\/'${RELEASE_TAG}'\/scloud_v'${SCLOUD_VERSION}'_linux_amd64.tar.gz\"'
echo "LINUX_URL is: $LINUX_URL"

# Ensure resources are uploaded before running SHA. Otherwise wait 5 minutes and check
i=0
while [ $i -le 1 ]
do
echo "Checking for resources..."
status=$(curl --head --silent https://github.com/splunk/splunk-cloud-sdk-go/releases/download/${RELEASE_TAG}/scloud_v${SCLOUD_VERSION}_darwin_amd64.tar.gz | head -n 1)
if [ "$i" -lt 1 ] && echo "$status" | grep -q 404
then
  echo "Resources hasn't been uploaded"
  echo "Sleeping for 5 minutes"
  sleep 5m
  i= $(( i++ ))
elif echo "$status" | grep -q 404
then
  echo "Cannot find resources"
  exit 1
else
  echo "Resources uploaded"
  i=$[$i+2]
fi
done

# Download Resource
"$(wget https://github.com/splunk/splunk-cloud-sdk-go/releases/download/${RELEASE_TAG}/scloud_v${SCLOUD_VERSION}_darwin_amd64.tar.gz)"
"$(wget https://github.com/splunk/splunk-cloud-sdk-go/releases/download/${RELEASE_TAG}/scloud_v${SCLOUD_VERSION}_linux_amd64.tar.gz)"

MAC_SHA="$(sha256sum -b scloud_v${SCLOUD_VERSION}_darwin_amd64.tar.gz)"
LINUX_SHA="$(sha256sum -b scloud_v${SCLOUD_VERSION}_linux_amd64.tar.gz)"

# Binary is in 0th position
MAC_SHA="$(echo $MAC_SHA | head -n1 | sed -e 's/\s.*$//')"
LINUX_SHA="$(echo $LINUX_SHA | head -n1 | sed -e 's/\s.*$//')"

echo "MAC_SHA is: $MAC_SHA"
echo "LINUX_SHA is: $LINUX_SHA"

if [[ "$MAC_SHA" == "$LINUX_SHA" ]]
then
    echo "Invalid SHA"
    exit 1
fi

####################################################################################################
# Clone homebrew-tap repo
####################################################################################################
echo "cloning homebrew-tap repo..."
BRANCH_NAME=master
git clone "https://${GITHUB_TOKEN}@github.com/splunk/homebrew-tap.git"
cd homebrew-tap

git remote set-url origin "https://srv-dev-platform:${GITHUB_TOKEN}@github.com/splunk/homebrew-tap.git"
git config user.email "srv-dev-platform@splunk.com"
git config user.name "srv-dev-platform"
git checkout "${BRANCH_NAME}"

####################################################################################################
# Print current scloud.rb file
####################################################################################################
echo "printing scloud.rb BEFORE update..."
cat scloud.rb
echo "Done printing..."

####################################################################################################
# Update scloud.rb
####################################################################################################
echo "updating scloud.rb..."
sed -ie  ${SED_FLG} '1,/version/ s/version.*/version \"'$SCLOUD_VERSION'\"/g' scloud.rb
grep "${SCLOUD_VERSION}" -q scloud.rb && echo "version updated successfully" || { echo "version updated failed" ; exit 1; }

sed -ie  ${SED_FLG} '10,/sha256/ s/.*sha256.*/    sha256 \"'$MAC_SHA'\"/g' scloud.rb
grep "${MAC_SHA}" -q scloud.rb && echo "MAC_SHA updated successfully" || { echo "MAC_SHA updated failed" ; exit 1; }

sed -ie  ${SED_FLG} '14,/sha256/ s/.*sha256.*/    sha256 \"'$LINUX_SHA'\"/g' scloud.rb
grep "${LINUX_SHA}" -q scloud.rb && echo "LINUX_SHA updated successfully" || { echo "LINUX_SHA updated failed" ; exit 1; }

sed -ie  ${SED_FLG} 's/.*darwin_amd64.*/    url '$MAC_URL'/g; s/.*linux_amd64.*/    url '$LINUX_URL'/g' scloud.rb
grep "${MAC_URL}" -q scloud.rb && echo "MAC_URL updated successfully" || { echo "MAC_URL updated failed" ; exit 1; }
grep "${LINUX_URL}" -q scloud.rb && echo "LINUX_URL updated successfully" || { echo "LINUX_URL updated failed" ; exit 1; }


####################################################################################################
# Print current scloud.rb file
####################################################################################################
echo "printing scloud.rb AFTER update..."
cat scloud.rb
echo "Done printing..."

####################################################################################################
# Commit and Push changes to homebrew-tap repo
####################################################################################################
echo "git status"
git status

diffs=$(git diff -- scloud.rb)
if [[ -z "${diffs}" ]] ; then
  echo "Failed: no changes were made to scloud.rb"
  exit 1
fi

git add scloud.rb
git commit -m "update scloud.rb for new release"
echo "git push origin ${BRANCH_NAME}"

if git push origin ${BRANCH_NAME}
then
  echo "Successfully published changes to homebrew"
else
  echo "Failed to push changes to homebrew"
  exit 1
fi
