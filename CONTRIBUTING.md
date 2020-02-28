# Contribution Guidelines


## Issues and bug reports

If you see unexpected behavior with this project, please [create an issue on GitHub](/issues) with the following information:

-  A title and a clear description of the issue.
-  The project version (for example "0.1.4").
-  The framework version (for example "Go 1.11").

If possible, include the following to help us reproduce the issue: 
-  A code sample that demonstrates the issue.
-  Any unit test cases that show how the expected behavior is not occurring.
-  An executable test case. 

If you have a question about Splunk, see [Splunk Answers](https://answers.splunk.com).

## Development

Configure your development environment as described in the project [README](/blob/master/README.md).

## Submit a pull request

1. Fill out the [Splunk Contribution Agreement](https://www.splunk.com/goto/contributions).
2. Create a new branch. For example:

    ```
    git checkout -b my-branch develop
    ```

3. Make code changes in your branch with tests. 
4. Commit your changes.
5. Push your branch to GitHub.

    ```
    git push origin my-branch
    ```

6. In GitHub, create a pull request that targets the **develop** branch. CI tests are run automatically.
7. After the pull request is merged, delete your branch.
8. Pull changes from the **develop** branch.

    ```
    git checkout develop
    git pull develop
    ```

## Contact us

If you have questions, reach out to us on [Slack](https://splunkdevplatform.slack.com) in the **#sdc** channel or email us at _sdcbeta@splunk.com_.