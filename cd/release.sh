#!/bin/sh
echo "Input the version to release omitting the leading 'v' (e.g. 0.7.1) followed by [ENTER]:"
read NEW_VERSION

echo "Preparing v$NEW_VERSION for release ..."
echo "Running \"git checkout develop\" ..."
git checkout develop
echo "Running \"git fetch --all && git pull --all\" ..."
git fetch --all && git pull --all
echo "Checking out a release/v$NEW_VERSION branch ..."
BRANCH_NAME=release/v$NEW_VERSION
git checkout -b $BRANCH_NAME
echo "Regenerating the services/*/interface.go files ..."
make generate_interface
git add services/*/interface.go
echo "Updating Version in services/client_info.go ..."
if [ "$(uname)" == "Darwin" ]; then
    # MacOS
    sed -E -i '' -e "s/[0-9]+\.[0-9]+\.[0-9]+/$NEW_VERSION/g" services/client_info.go
else
    # Linux
    sed -r -i '' -e "s/[0-9]+\.[0-9]+\.[0-9]+/$NEW_VERSION/g" services/client_info.go
fi
git add services/client_info.go
echo "Updating docs and generating cicd-publish artifact ..."
make docs_publish
git add docs/
echo "Showing changes with \"git status\" ..."
git status
echo "Review the git status above and if your changes look good press y followed by [ENTER] to commit and push your release branch:"
read PUSH_TO_GIT
if [ "$PUSH_TO_GIT" = "y" ]
then
    echo "Creating commit for client_info.go and docs/ changes ..."
    git commit -m "Release v$NEW_VERSION"
    echo "Pushing branch $BRANCH_NAME ..."
    git push origin $BRANCH_NAME
    echo
    echo " Remaining steps: "
    echo "   1. Create a PR for $BRANCH_NAME -> master"
    echo ""
    echo "   2. Merge the PR - ! USE A MERGE COMMIT NOT A MERGE-SQUASH !"
    echo ""
    echo "   3. Create the release in Github with the option of creating a new tag = v$NEW_VERSION"
    echo ""
    echo "   4. Pull the master branch and run \"make docs_publish\" ... then deliver the contents of ci/docs/build/*v$NEW_VERSION.tgz to the portals team"
    echo ""
    echo "   5. Create/merge a PR from master -> develop"
    echo ""
    echo ""
else
    echo "No changes pushed, branch $BRANCH_NAME only created locally ..."
fi