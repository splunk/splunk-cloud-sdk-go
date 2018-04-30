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
	@if [ -f ci/secret.env ]; then \
		jet encrypt ci/secret.env ci/env.encrypted && \
		printf "Encrypted ci/secret.env to ci/env.encrypted\n"; \
	fi;
	@if [ -f ci/shared/secret.env ]; then \
		jet encrypt ci/shared/secret.env ci/shared/env.encrypted && \
		printf "Encrypted ci/shared/secret.env to ci/shared/env.encrypted\n"; \
	fi;
	@if [ -f ci/integration/secret.env ]; then \
		jet encrypt ci/integration/secret.env ci/integration/env.encrypted && \
		printf "Encrypted ci/integration/secret.env to ci/integration/env.encrypted\n"; \
	fi;

decrypt:
	@if [ -f ci/env.encrypted ]; then \
		jet decrypt ci/env.encrypted ci/secret.env && \
		printf "Decrypted ci/env.encrypted to ci/secret.env\n"; \
	fi;
	@if [ -f ci/shared/env.encrypted ]; then \
		jet decrypt ci/shared/env.encrypted ci/shared/secret.env && \
		printf "Decrypted ci/shared/env.encrypted to ci/shared/env.encrypted\n"; \
	fi;
	@if [ -f ci/integration/env.encrypted ]; then \
		jet decrypt ci/integration/env.encrypted ci/integration/secret.env && \
		printf "Decrypted ci/integration/env.encrypted to ci/integration/env.encrypted\n"; \
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