# Copyright © 2018 Splunk Inc.
# SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
# without a valid written license from Splunk Inc. is PROHIBITED.
#

.DEFAULT_GOAL := noop

GO_NON_VENDOR_PACKAGES := $(shell go list ./... | grep -v /vendor/)
GO_NON_TEST_NON_VENDOR_PACKAGES := $(shell go list ./... | grep -v /vendor/ | grep -v test)

GIT_COMMIT_TAG := $(shell git rev-parse --verify HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
GIT_VERSION_TAG := $(shell git describe origin/master --tags --match="v*" | sed 's/v//g')

LOCAL_TEST_URL_PROTOCOL := http
LOCAL_TEST_SSC_HOST := localhost:8882
LOCAL_TEST_BEARER_TOKEN := TEST_AUTH_TOKEN
LOCAL_TEST_TENANT_ID := TEST_TENANT

DOCKER_STUBBY_TEST_URL_PROTOCOL := http
DOCKER_STUBBY_TEST_SSC_HOST := ssc-sdk-shared-stubby:8882
DOCKER_STUBBY_TEST_BEARER_TOKEN := TEST_AUTH_TOKEN
DOCKER_STUBBY_TEST_TENANT_ID := TEST_TENANT

noop:
	@echo "No make target specified."

clean:
	docker rmi -f 137462835382.dkr.ecr.us-west-1.amazonaws.com/ssc-sdk-shared-stubby

lint:
	go get golang.org/x/lint/golint && golint --set_exit_status $(GO_NON_VENDOR_PACKAGES)

vet:
	go vet $(GO_NON_VENDOR_PACKAGES)

build:
	sed -i '' -e 's/[0-9].[0-9].[0-9]/$(GIT_VERSION_TAG)/g' service/client_info.go
	go build $(GO_NON_TEST_NON_VENDOR_PACKAGES)

encrypt:
	@if [ -f ci/secret.env ]; then \
		jet encrypt ci/secret.env ci/env.encrypted && \
		printf "Encrypted ci/secret.env to ci/env.encrypted\n"; \
	fi;
	@if [ -f ci/shared/secret.env ]; then \
		jet encrypt ci/shared/secret.env ci/shared/env.encrypted && \
		printf "Encrypted ci/shared/secret.env to ci/shared/env.encrypted\n"; \
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

docs_md:
	./ci/docs/docs_md.sh

docs_publish: docs_md
	./ci/docs/publish.sh

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

debug_local_environment_variables:
	@echo "Local Testing Environment Variables"
	@echo "LOCAL_TEST_URL_PROTOCOL: $(LOCAL_TEST_URL_PROTOCOL)"
	@echo "LOCAL_TEST_SSC_HOST: $(LOCAL_TEST_SSC_HOST)"
	@echo "LOCAL_TEST_BEARER_TOKEN: $(LOCAL_TEST_BEARER_TOKEN)"
	@echo "LOCAL_TEST_TENANT_ID: $(LOCAL_TEST_TENANT_ID)"
	@echo

debug_docker_environment_variables:
	@echo "Docker Testing Environment Variables"
	@echo "DOCKER_STUBBY_TEST_URL_PROTOCOL: $(DOCKER_STUBBY_TEST_URL_PROTOCOL)"
	@echo "DOCKER_STUBBY_TEST_SSC_HOST: $(DOCKER_STUBBY_TEST_SSC_HOST)"
	@echo "DOCKER_STUBBY_TEST_BEARER_TOKEN: $(DOCKER_STUBBY_TEST_BEARER_TOKEN)"
	@echo "DOCKER_STUBBY_TEST_TENANT_ID: $(DOCKER_STUBBY_TEST_TENANT_ID)"
	@echo

run_unit_tests:
	sh ./ci/unit_tests/run_unit_tests.sh

run_local_stubby_tests: debug_local_environment_variables
	TEST_URL_PROTOCOL=$(LOCAL_TEST_URL_PROTOCOL) \
	TEST_SSC_HOST=$(LOCAL_TEST_SSC_HOST) \
	TEST_BEARER_TOKEN=$(LOCAL_TEST_BEARER_TOKEN) \
	TEST_TENANT_ID=$(LOCAL_TEST_TENANT_ID) \
	go test -v ./test/stubby_integration/...

run_docker_stubby_tests: debug_docker_environment_variables
	TEST_URL_PROTOCOL=$(DOCKER_STUBBY_TEST_URL_PROTOCOL) \
	TEST_SSC_HOST=$(DOCKER_STUBBY_TEST_SSC_HOST) \
	TEST_BEARER_TOKEN=$(DOCKER_STUBBY_TEST_BEARER_TOKEN) \
	TEST_TENANT_ID=$(DOCKER_STUBBY_TEST_TENANT_ID) \
	go test -v ./test/stubby_integration/...
