# splunk-cloud-sdk-go
A Go client for Splunk Cloud services

# Terms of Service (TOS)
[Splunk Cloud Terms of Service](https://www.splunk.com/en_us/legal/terms/splunk-cloud-pre-release-terms-of-service.html)

# Getting started
---
Install Go 1.11 (or later)
* [Install Go and Setup your Go environment](https://golang.org/doc/install)

Install recommended Go tools
  * `go get golang.org/x/lint/golint`
  * `go get -u golang.org/x/tools/cmd/goimports`

Below are steps to initialize your project using Go Modules for dependency support, for more info see: https://github.com/golang/go/wiki/Modules

Initialize your project which can be outside $GOPATH/src
(NOTE: if the project is within your $GOPATH you must set `GO111MODULE=on` in your environment variables before continuing):
```bash
$ mkdir myproject && cd myproject
$ go mod init github.com/example/myproject
```

Create a `myproject/main.go` file containing:
```go
package main

import (
	"fmt"
	"os"

	"github.com/splunk/splunk-cloud-sdk-go/service"
)

func main() {
	checkForTenantToken()
	// Initialize the client
	client, err := service.NewClient(&service.Config{
		Token: os.Getenv("BEARER_TOKEN"),
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

Setup your project's Go module dependencies to point to the local copy of ./splunk-sdk-go:
```bash
$ go mod edit -replace=github.com/splunk/splunk-cloud-sdk-go=./splunk-cloud-sdk-go
```

Set your tenant and token, tokens can be retrieved from https://sdc.splunkbeta.com/settings:
```bash
$ export BEARER_TOKEN=<INSERT_TOKEN>
$ export TENANT=<INSERT_TENANT>
```

Finally, build and run your project:
```bash
$ go build
$ ./myproject
info: &{Name:me@example.com Tenants:[]}
```

## Documentation
For general documentation about the Splunk Cloud SDK for Go, see:
- https://sdc.splunkbeta.com/docs/sdks/gosdk

For the API reference for the Splunk Cloud SDK for Go, see:
- https://sdc.splunkbeta.com/reference/sdk/splunk-cloud-sdk-go

The API reference contains detailed information about all classes and functions, with clearly-defined parameters and return types.
