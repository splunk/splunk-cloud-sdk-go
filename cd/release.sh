#!/bin/sh
echo "Input the version to release omitting the leading 'v' (e.g. 0.7.1) followed by [ENTER]:"
read NEW_VERSION

echo "Preparing v$NEW_VERSION for release ..."
echo "Running `git checkout develop` ..."
git checkout develop
echo "Running `git fetch --all && git pull --all` ..."
git fetch --all && git pull --all
echo "Checking out a release/v$NEW_VERSION branch ..."
BRANCH_NAME=release/v$NEW_VERSION
git checkout -b $BRANCH_NAME
echo "Updating Version in service/client_info.go ..."
sed -i '' -e "s/[0-9].[0-9].[0-9]/$NEW_VERSION/g" service/client_info.go
git add service/client_info.go
echo "Updating docs and generating cicd-publish artifact ..."
make docs_publish
git add docs/
echo "Showing changes with `git status` ..."
git status
echo "Review the git status above and if your changes look good press y followed by [ENTER] to commit and push your release branch and release tag:"
read PUSH_TO_GIT
if [ "$PUSH_TO_GIT" -eq "y" ]
then
    echo "Creating commit for client_info.go and docs/ changes ..."
    git commit -m "Release v$NEW_VERSION"
    echo "Creating tag: v$NEW_VERSION ..."
    git tag -a v$NEW_VERSION -m "Release v$NEW_VERSION"
    echo "Pushing branch $BRANCH_NAME ..."
    git push origin $BRANCH_NAME
    echo "Pushing tag v$NEW_VERSION ..."
    git push origin v$NEW_VERSION
    echo "PRs should be created for $BRANCH_NAME -> master AND $BRANCH_NAME -> develop ..."
    echo "Finally, the docs package in ci/docs/build/ should be delivered to the portals team ..."
else
    echo "No changes pushed, branch $BRANCH_NAME only created locally ..."
fi