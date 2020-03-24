# Splunk Cloud Services SDK for Go
[![Go Report Card](https://goreportcard.com/badge/github.com/splunk/splunk-cloud-sdk-go)](https://goreportcard.com/report/github.com/splunk/splunk-cloud-sdk-go) 
[![GoDoc](https://godoc.org/github.com/splunk/splunk-cloud-sdk-go?status.svg)](https://godoc.org/github.com/splunk/splunk-cloud-sdk-go)

The Splunk Cloud Services software development kit (SDK) for Go contains library code and examples to enable you to build apps using the Splunk Cloud Services with the Go programming language.

    **Note:** This SDK is not used for Splunk Enterprise or Splunk Cloud development. For information about developing apps and add-ons for those products, see the [Splunk Developer Portal for Splunk Enterprise](https://dev.splunk.com/enterprise/).

# Splunk Cloud Services CLI

Splunk Cloud Services CLI, `scloud`, is a command-line tool for developers to make API calls to Splunk Cloud Services.

For more information about Splunk Cloud Services CLI, see the [Splunk Developer Portal](https://dev.splunk.com/scs/docs/overview/tools/tools_scloud).

## Terms of Service

Log in to [Splunk Investigate](https://si.scp.splunk.com/) and accept the Terms of Service when prompted.

## Get started


### Install Go and Go tools

1. Install Go 1.11 or later from the [Getting Started](https://golang.org/doc/install) page on the Go Programmming Language website.

2. Install recommended tools for Go by running the following commands:

    ```bash
    $ go get golang.org/x/lint/golint
    $ go get -u golang.org/x/tools/cmd/goimports
    ```


### Initialize your project

Initialize your project using Go modules for dependency support. Your project can be located outside of the **$GOPATH/src** directory. For more about modules, see [Go 1.11 Modules](https://github.com/golang/go/wiki/Modules) on the GitHub website.

1. If your project is within your **$GOPATH**, set `GO111MODULE=on` in your environment variables.

2. Initialize your project by running the following commands, but replace the `<github.com/example/myproject>` path with your Git host, organization, user name, and project name as appropriate:

    ```bash
    $ mkdir myproject && cd myproject
    $ go mod init <github.com/example/myproject>
    ```

3. Create a **main.go** file within your project directory containing the following code:

    ```go
    package main

    import (
        "fmt"
        "os"

        "github.com/splunk/splunk-cloud-sdk-go/sdk"
        "github.com/splunk/splunk-cloud-sdk-go/services"
        "github.com/splunk/splunk-cloud-sdk-go/services/identity"
    )

    func main() {
        checkForTenantToken()
        // Initialize the client
        client, err := sdk.NewClient(&services.Config{
            Token:  os.Getenv("BEARER_TOKEN"),
            Tenant: os.Getenv("TENANT"),
        })
        exitOnErr(err)
        // Validate access to Splunk Cloud Services and tenant
        query := identity.ValidateTokenQueryParams{Include: []string{"principal", "tenant"}}
        info, err := client.IdentityService.ValidateToken(&query)
        exitOnErr(err)
        fmt.Println("name: " + info.Name)
        fmt.Println("principal name: " + info.Principal.Name)
        fmt.Println("tenant name: " + info.Tenant.Name)
    }

    func exitOnErr(err error) {
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    }

    func checkForTenantToken() {
        if os.Getenv("BEARER_TOKEN") == "" {
            exitOnErr(fmt.Errorf("$BEARER_TOKEN must be set"))
        }
        if os.Getenv("TENANT") == "" {
            exitOnErr(fmt.Errorf("$TENANT must be set"))
        }
    }
    ```

4. Set your access token and tenant.

    -  Retrieve your access token from the [Splunk Developer Portal](https://dev.splunk.com/scs/).
       -  Log in at the top of the page. 
       -  Open the [Reference](https://dev.splunk.com/scs/reference).
       -  Navigate to any REST API. For example, go to [Provisioner GET /system/provisioner/v1beta1/tenants](https://api.scp.splunk.com/system/provisioner/v1beta1/tenants).
       -  Click **Console**, then next to Authorization Token click **Show**, and then click **Copy to Clipboard** to copy your token.
    -  List your tenants using the following REST command, replacing `<accessToken>` with your access token:

        ```curl
        curl -X GET "https://api.scp.splunk.com/system/identity/v2beta1/tenants" \
        -H "Authorization: Bearer <accessToken>"
        ```

    -  Run the following command, replacing `<accessToken>` and `<tenant>` with your values:

        ```bash
        $ export BEARER_TOKEN=<accessToken>
        $ export TENANT=<tenant>
        ```

5. Build and run your project by running the following commands, where `<project>` is the name of your project, and `<me@example.com>` is your user name:

    ```bash
    $ go build
    $ ./<project>
    name: <me@example.com>
    principal name: <me@example.com>
    tenant name: <mytenant>
    ```

## Documentation
For general documentation, see the [Splunk Developer Portal](https://dev.splunk.com/scs/).

For reference documentation, see the [Splunk Cloud Services SDK for Go API Reference](https://dev.splunk.com/scs/reference/sdk/splunk-cloud-sdk-go).

## Contributing

Do not directly edit any source file with `_generated` in the name because that file was generated from service specifications.

## Contact
If you have questions, reach out to us on [Slack](https://splunkdevplatform.slack.com) in the **#sdc** channel or email us at _devinfo@splunk.com_.
