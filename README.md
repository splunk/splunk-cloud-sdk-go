# ssc-client-go
A Go client for Self Service Cloud services

| Branch | Codeship | Coverage |
|:------:|:--------:|:--------:|
| develop | [![Codeship Status for splunk/ssc-client-go](https://app.codeship.com/projects/d0ec9ea0-15c2-0136-e7ad-1a0f3e5cdd95/status?branch=develop)](https://app.codeship.com/projects/283638) | [![codecov](https://codecov.io/gh/splunk/ssc-client-go/branch/develop/graph/badge.svg?token=o4BjP93wQt)](https://codecov.io/gh/splunk/ssc-client-go/branch/develop) |
| master | [![Codeship Status for splunk/ssc-client-go](https://app.codeship.com/projects/d0ec9ea0-15c2-0136-e7ad-1a0f3e5cdd95/status?branch=master)](https://app.codeship.com/projects/283638) | [![codecov](https://codecov.io/gh/splunk/ssc-client-go/branch/master/graph/badge.svg?token=o4BjP93wQt)](https://codecov.io/gh/splunk/ssc-client-go/branch/master) |


# Development
---
### Getting setup with macOS
* [Install Brew](https://brew.sh/)
* [Install Docker for Mac](https://docs.docker.com/docker-for-mac/install/)
* [Setup your Go environment](https://golang.org/doc/install) and checkout this repository into `$GOPATH/src/github.com/splunk/ssc-client-go`
* Install local development tools using `make install_local`
### Optional
* For encypting or decrypting secrets you will need to [download the AES key from Codeship](https://app.codeship.com/projects/283638/configure) and move/rename it to `$GOPATH/src/github.com/splunk/ssc-client-go/codeship.aes`
### Running a stubby server locally for testing
* Download the shared stubby Docker image from ECR using `jet load ssc-sdk-shared-stubby`
* Run the stubby server locally with `make stubby_local`
* Test by visiting http://localhost:8882/error in your browser. You should see `{"message":"Something exploded"}`
* NOTE: At the moment the stubby container does not fail gracefully, you may need to stop and rm all containers in another shell terminal using `docker stop $(docker ps -a -q)` and `docker rm $(docker ps -a -q)`
### Making updates to imported libraries
* This repo uses `dep` for dependency management
* If you have imported an outside library as part of your changes and are seeing errors when running `make vet` such as `cannot find package` you will need to update the dependencies using `make dependencies_update` which may regenerate the `Gopkg.*` files. Changes to the `Gopkg.*` should be commited along with your imports in code.