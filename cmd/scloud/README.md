# Splunk Cloud Services CLI

Splunk Cloud Services CLI, `scloud`, is a command-line tool for developers to make API calls to Splunk Cloud Services. The Splunk Cloud Services CLI uses the Splunk Cloud Services SDK for Go for many API calls and is a great way to explore the available functionality.

## Terms of Service

[Splunk Cloud Services Terms of Service](https://auth.scp.splunk.com/tos)

## Install the Splunk Cloud Services CLI tool

### Download from GitHub

1. Download the Splunk Cloud Services CLI package for your operating system from the **/releases** directory in this repository.

2. Extract the package file to a home directory on your computer. 

    For example: 

    - On *nix, extract the .tar.gz file to a directory such as `/usr/local/bin`. 

    - On Windows, unzip the package to a directory such as `C:\scloud`. Consider adding this location to the PATH system variable.

3. Accept the Terms of Service by logging in to [Splunk Investigate](https://si.scp.splunk.com/).


### Install from Homebrew (macOS, Linux)

If you have Homebrew installed, install the Splunk Cloud Services CLI by doing the following: 

1. Register the Splunk Homebrew Tap. At the command line, enter: 

   ```
   brew tap splunk/tap
   ```

2. Install the Splunk Cloud Services CLI package: 

   ```
   brew install scloud
   ```

3. Accept the Terms of Service by logging in to [Splunk Investigate](https://si.scp.splunk.com/).


## Run Splunk Cloud Services CLI

To access the Splunk Cloud Services CLI, you need: 
* A user account with Splunk Cloud Services
* A shell prompt, command prompt, or PowerShell session

Then, use `scloud <command>` at the command line to perform almost any operation in Splunk Cloud Services.

Here are some commands to get started: 

| To...                            | Enter:                               |
|----------------------------------|--------------------------------------|
| Get help for all commands        | `scloud help`                        |
| Get help for a specific command  | `scloud <commandname> help`          |
| Log in                           | `scloud -u <username> login`         |
| Save your user name in settings  | `scloud set username <username>`     |
| Save a tenant name in settings   | `scloud set tenant <tenantname>`     |
| Display saved settings           | `scloud get-settings`                |
| List your tenants                | `scloud provisioner list-tenants`    |

## Documentation
For more about using `scloud`, see [Splunk Cloud Services CLI](https://developer.splunk.com/scs/docs/overview/tools/tools_scloud) on the Splunk Developer Portal.

## Contact
If you have questions, reach out to us on [Slack](https://splunkdevplatform.slack.com) in the **#sdc** channel or email us at _devinfo@splunk.com_.
