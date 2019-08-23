# Splunk Cloud SDK for Go
[![Go Report Card](https://goreportcard.com/badge/github.com/splunk/splunk-cloud-sdk-go)](https://goreportcard.com/report/github.com/splunk/splunk-cloud-sdk-go) 
[![GoDoc](https://godoc.org/github.com/splunk/splunk-cloud-sdk-go?status.svg)](https://godoc.org/github.com/splunk/splunk-cloud-sdk-go)

The Splunk Cloud software development kit (SDK) for Go contains library code and examples to enable you to build apps using the Splunk Cloud services with the Go programming language.

To use Splunk Cloud SDKs, you must be included in the Splunk Investigate Beta Program.
Sign up here: https://si.scp.splunk.com/.

# Splunk Cloud CLI

Splunk Cloud CLI, `scloud`, is a command-line tool for developers using Splunk Developer Cloud to make API calls in the Splunk Cloud Platform.

For more information about Splunk Cloud CLI [see the README](cmd/scloud/README.md).

## Terms of Service (TOS)
[Splunk Cloud Terms of Service](https://www.splunk.com/en_us/legal/terms/splunk-cloud-pre-release-terms-of-service.html)


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
        // Validate access to the platform and tenant
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

    -  Retrieve your access token from the [Splunk Developer Cloud Portal](https://sdc.splunkbeta.com/settings).

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
For general documentation, see the [Splunk Developer Cloud Portal](https://sdc.splunkbeta.com/).

For reference documentation, see the [Splunk Cloud SDK for Go API Reference](https://sdc.splunkbeta.com/reference/sdk/splunk-cloud-sdk-go).

## Contributing

Do not directly edit any source file with `_generated` in the name because that file was generated from service specifications.

## Contact
If you have questions, reach out to us on [Slack](https://splunkdevplatform.slack.com) in the **#sdc** channel or email us at _sdcbeta@splunk.com_.
