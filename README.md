# splunk-cloud-sdk-go
A Go client for Splunk Cloud services

| Branch | Codeship | Coverage |
|:------:|:--------:|:--------:|
| develop | [![Codeship Status for splunk/splunk-cloud-sdk-go](https://app.codeship.com/projects/d0ec9ea0-15c2-0136-e7ad-1a0f3e5cdd95/status?branch=develop)](https://app.codeship.com/projects/283638) | [![codecov](https://codecov.io/gh/splunk/splunk-cloud-sdk-go/branch/develop/graph/badge.svg?token=o4BjP93wQt)](https://codecov.io/gh/splunk/splunk-cloud-sdk-go/branch/develop) |
| master | [![Codeship Status for splunk/splunk-cloud-sdk-go](https://app.codeship.com/projects/d0ec9ea0-15c2-0136-e7ad-1a0f3e5cdd95/status?branch=master)](https://app.codeship.com/projects/283638) | [![codecov](https://codecov.io/gh/splunk/splunk-cloud-sdk-go/branch/master/graph/badge.svg?token=o4BjP93wQt)](https://codecov.io/gh/splunk/splunk-cloud-sdk-go/branch/master) |


# Getting started
---
### macOS
* [Install Brew](https://brew.sh/)
* [Install Docker for Mac](https://docs.docker.com/docker-for-mac/install/)
* [Install Go and Setup your Go environment](https://golang.org/doc/install) and checkout this repository into `$GOPATH/src/github.com/splunk/splunk-cloud-sdk-go`
* Recommended Go tools:
  * `go get -u github.com/golang/dep/cmd/dep`
  * `go get golang.org/x/lint/golint`
  * `go get -u golang.org/x/tools/cmd/goimports`
* Clone/unzip our splunk/splunk-cloud-sdk-go repo into your project's vendor/github.com/splunk/splunk-cloud-sdk-go directory
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
	// Validate access to the platform
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