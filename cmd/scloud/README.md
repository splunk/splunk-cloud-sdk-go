# Splunk Cloud CLI

Splunk Cloud CLI, `scloud`, is a command-line tool for developers using Splunk Developer Cloud to make API calls in the Splunk Cloud Platform. The Splunk Cloud CLI uses the Splunk Cloud SDK for Go for many API calls and is a great way to explore the available functionality.

## Terms of Service

[Splunk Cloud Terms of Service](https://www.splunk.com/en_us/legal/terms/splunk-cloud-pre-release-terms-of-service.html)

## Download and install the Splunk Cloud CLI tool

1. Download the Splunk Cloud CLI package for your operating system from the **/releases** directory in this repository.

2. Extract the package file to a home directory on your computer. 

    For example: 

    - On *nix, extract the .tar.gz file to a directory such as `/usr/local/bin`. 

    - On Windows, unzip the package to a directory such as `C:\scloud`. Consider adding this location to the PATH system variable.

3. Accept the Terms of Service by logging in to [Splunk Investigate](https://si.scp.splunk.com/).

## Run Splunk Cloud CLI

To access the Splunk Cloud CLI, you need: 
* A user account with Splunk Developer Cloud
* A shell prompt, command prompt, or PowerShell session

Then, use `scloud <command>` at the command line to perform almost any operation in SDC.

Here are some commands to get started: 

| To...                            | Enter:                               |
|----------------------------------|--------------------------------------|
| Get help for all commands        | `scloud help`                        |
| Get help for a specific command  | `scloud <commandname> help`          |
| Log in                           | `scloud -u <username> login`         |
| Save your user name in settings  | `scloud set username <username>`     |
| Save a tenant name in settings   | `scloud set tenant <tenantname>`     |
| Display saved settings           | `scloud get-settings`                |
| List your tenants                | `scloud identity list-tenants`       |

## Documentation
For more about using `scloud`, see [Splunk Cloud CLI](https://sdc.splunkbeta.com/docs/overview/sdctools/tools_scloud) on the Splunk Developer Cloud Portal.

## Contact
If you have questions, reach out to us on [Slack](https://splunkdevplatform.slack.com) in the **#sdc** channel or email us at _sdcbeta@splunk.com_.
