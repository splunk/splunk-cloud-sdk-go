# Copyright 2018 Splunk

.DEFAULT_GOAL := noop

GO_SOURCES := $(shell find . -name '*.go')
GO_NON_VENDOR_PACKAGES := $(shell go list ./... | grep -v /vendor/)

GIT_COMMIT_TAG := $(shell git rev-parse --verify HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

noop:
	@echo "No make target specified."

clean:
	docker rmi -f 137462835382.dkr.ecr.us-west-1.amazonaws.com/ssc-sdk-shared-stubby

ssc-client-go: $(GO_SOURCES)
	go build -v ./...

lint:
	go get golang.org/x/lint/golint && golint --set_exit_status $(GO_NON_VENDOR_PACKAGES)

vet:
	go vet $(GO_NON_VENDOR_PACKAGES)

build:
	make ssc-client-go

encrypt:
	@if [ -f deploy/secret.env ]; then \
		jet encrypt deploy/secret.env deploy/env.encrypted && \
		printf "Encrypted deploy/secret.env to deploy/env.encrypted\n"; \
	fi;
	@if [ -f deploy/shared/secret.env ]; then \
		jet encrypt deploy/shared/secret.env deploy/shared/env.encrypted && \
		printf "Encrypted deploy/shared/secret.env to deploy/shared/env.encrypted\n"; \
	fi;
	@if [ -f deploy/integration/secret.env ]; then \
		jet encrypt deploy/integration/secret.env deploy/integration/env.encrypted && \
		printf "Encrypted deploy/integration/secret.env to deploy/integration/env.encrypted\n"; \
	fi;

decrypt:
	@if [ -f deploy/env.encrypted ]; then \
		jet decrypt deploy/env.encrypted deploy/secret.env && \
		printf "Decrypted deploy/env.encrypted to deploy/secret.env\n"; \
	fi;
	@if [ -f deploy/shared/env.encrypted ]; then \
		jet decrypt deploy/shared/env.encrypted deploy/shared/secret.env && \
		printf "Decrypted deploy/shared/env.encrypted to deploy/shared/env.encrypted\n"; \
	fi;
	@if [ -f deploy/integration/env.encrypted ]; then \
		jet decrypt deploy/integration/env.encrypted deploy/integration/secret.env && \
		printf "Decrypted deploy/integration/env.encrypted to deploy/integration/env.encrypted\n"; \
	fi;

install_local:
	printf "Installing dep to manage Go dependencies ..." && \
	make install_dep
	printf "Installing Codeship jet for local build acceptance testing, if there are any issues installing see: https://documentation.codeship.com/pro/jet-cli/installation/ ..." && \
	brew cask install codeship/taps/jet

stubby_local:
	jet load ssc-client-go-with-stubby
	docker run -p 8889:8889 -p 8882:8882 -p 7443:7443 137462835382.dkr.ecr.us-west-1.amazonaws.com/ssc-sdk-shared-stubby

install_dep:
	go get -u github.com/golang/dep/cmd/dep

dependencies:
	dep ensure -vendor-only
	go get -u golang.org/x/tools/cmd/goimports

dependencies_update:
	dep ensure -update