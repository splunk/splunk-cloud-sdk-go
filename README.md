# splunk-cloud-sdk-go

\#TODO(dan): test

A Go client for Splunk Cloud services

# Terms of Service (TOS)
[Splunk Cloud Terms of Service](https://www.splunk.com/en_us/legal/terms/splunk-cloud-pre-release-terms-of-service.html)

# Getting started
---

## Install Go and Go tools

1. Install Go 1.11 or later from the [Getting Started](https://golang.org/doc/install) page on the Go Programmming Language website.
   
2. Install recommended tools for Go by running the following commands:
    
    ```bash
    $ go get golang.org/x/lint/golint
    $ go get -u golang.org/x/tools/cmd/goimports
    ```


## Initialize your project

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
        info, err := client.IdentityService.Validate()
        exitOnErr(err)
        fmt.Printf("info: %+v", info)
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
    
4. Clone the Splunk Cloud SDK for Go repository by navigating to your project directory and running the following command:
    
    ```bash
    $ git clone https://github.com/splunk/splunk-cloud-sdk-go
    ```

5. Set up your project's Go module dependencies to point to the cloned SDK repository by running the following command:
    
    ```bash
    $ go mod edit -replace=github.com/splunk/splunk-cloud-sdk-go=./splunk-cloud-sdk-go
    ```

6. Set your tenant and token by running the following commands, but replace the values for `<mytoken>` and `<mytenant>`. You can retrieve these values from https://sdc.splunkbeta.com/settings. 

    ```bash
    $ export BEARER_TOKEN=<mytoken>
    $ export TENANT=<mytenant>
    ```

7. Build and run your project by running the following commands, where `<myproject>` is the name of your project, and `<me@example.com>` is your user name:
    
    ```bash
    $ go build
    $ ./<myproject>
    info: &{Name:<me@example.com> Tenants:[]}
    ```

## Documentation
For general documentation about the Splunk Cloud SDK for Go, see:
- https://sdc.splunkbeta.com/docs/sdks/gosdk

For the API reference for the Splunk Cloud SDK for Go, see:
- https://sdc.splunkbeta.com/reference/sdk/splunk-cloud-sdk-go

The API reference contains detailed information about all classes and functions, with clearly-defined parameters and return types.
