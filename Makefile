# Copyright 2018 Splunk

.DEFAULT_GOAL := noop

GO_SOURCES := $(shell find . -name '*.go')
GO_NON_VENDOR_PACKAGES := $(shell go list ./... | grep -v /vendor/)

GIT_COMMIT_TAG := $(shell git rev-parse --verify HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

noop:
	@echo "No make target specified."

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
		printf "Encrypted deploy/secret.env  to deploy/env.encrypted\n"; \
	fi;

decrypt:
	@if [ -f deploy/env.encrypted ]; then \
		jet decrypt deploy/env.encrypted deploy/secret.env && \
		printf "Decrypted deploy/env.encrypted to deploy/secret.env\n"; \
	fi;

install_dep:
	go get -u github.com/golang/dep/cmd/dep

dependencies:
	dep ensure -vendor-only
	go get -u golang.org/x/tools/cmd/goimports

dependencies_update:
	dep ensure -update