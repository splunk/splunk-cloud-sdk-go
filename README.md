# splunk-cloud-sdk-go
A Go client for Splunk Cloud services

# Terms of Service (TOS)
[Splunk Cloud Terms of Service](https://www.splunk.com/en_us/legal/terms/splunk-cloud-pre-release-terms-of-service.html)

# Getting started
---
### macOS
* [Install Brew](https://brew.sh/)
* [Install Docker for Mac](https://docs.docker.com/docker-for-mac/install/)
* [Install Go and Setup your Go environment](https://golang.org/doc/install)
* Recommended Go tools:
  * `go get -u github.com/golang/dep/cmd/dep`
  * `go get golang.org/x/lint/golint`
  * `go get -u golang.org/x/tools/cmd/goimports`
* Clone/unzip our splunk-cloud-sdk-go repo into your project's vendor/github.com/splunk/splunk-cloud-sdk-go directory
* Initialize a new client:

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

## Documentation
For general documentation about the Splunk Cloud SDK for Go, see:
- https://sdc.splunkbeta.com/docs/sdks/gosdk

For the API reference for the Splunk Cloud SDK for Go, see:
- https://sdc.splunkbeta.com/reference/sdk/splunk-cloud-sdk-go

The API reference contains detailed information about all classes and functions, with clearly-defined parameters and return types.
