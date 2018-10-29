# Copyright © 2018 Splunk Inc.
# SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
# without a valid written license from Splunk Inc. is PROHIBITED.
#

.DEFAULT_GOAL := noop

GO_NON_VENDOR_PACKAGES := $(shell go list ./... | grep -v /vendor/)
GO_NON_TEST_NON_VENDOR_PACKAGES := $(shell go list ./... | grep -v /vendor/ | grep -v test)

GIT_COMMIT_TAG := $(shell git rev-parse --verify HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

LOCAL_SPLUNK_CLOUD_HOST := localhost:8882
LOCAL_BEARER_TOKEN := AUTH_TOKEN
LOCAL_TENANT_ID := TENANT

DOCKER_STUBBY_SPLUNK_CLOUD_HOST := splunk-cloud-sdk-shared-stubby:8882
DOCKER_STUBBY_BEARER_TOKEN := AUTH_TOKEN
DOCKER_STUBBY_TENANT_ID := TENANT

noop:
	@echo "No make target specified."

clean: build

lint:
	go get golang.org/x/lint/golint && golint --set_exit_status $(GO_NON_VENDOR_PACKAGES)

vet:
	go vet $(GO_NON_VENDOR_PACKAGES)

build:
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

docs: docs_md

docs_md: .FORCE
	./ci/docs/docs_md.sh

docs_publish: docs_md
	./ci/docs/publish.sh

release: .FORCE
	./cd/release.sh

install_local:
	printf "Installing dep to manage Go dependencies ..." && \
	make install_dep
	printf "Installing Codeship jet for local build acceptance testing, if there are any issues installing see: https://documentation.codeship.com/pro/jet-cli/installation/ ..." && \
	brew cask install codeship/taps/jet

docker_stubby:
	jet load splunk-cloud-sdk-go-with-stubby
	docker run -p 8889:8889 -p 8882:8882 -p 7443:7443 splunk/splunk-cloud-sdk-shared-stubby

install_dep:
	go get -u github.com/golang/dep/cmd/dep

install_test_dep:
	go get -u github.com/stretchr/testify

dependencies:
	dep ensure -vendor-only
	go get -u golang.org/x/tools/cmd/goimports

dependencies_update:
	dep ensure -update

debug_local_environment_variables:
	@echo "Local Testing Environment Variables"
	@echo "LOCAL_TEST_SPLUNK_CLOUD_HOST: $(LOCAL_TEST_SPLUNK_CLOUD_HOST)"
	@echo "LOCAL_TEST_BEARER_TOKEN: $(LOCAL_TEST_BEARER_TOKEN)"
	@echo "LOCAL_TEST_TENANT_ID: $(LOCAL_TEST_TENANT_ID)"
	@echo

debug_docker_environment_variables:
	@echo "Docker Testing Environment Variables"
	@echo "DOCKER_STUBBY_SPLUNK_CLOUD_HOST: $(DOCKER_STUBBY_SPLUNK_CLOUD_HOST)"
	@echo "DOCKER_STUBBY_BEARER_TOKEN: $(DOCKER_STUBBY_BEARER_TOKEN)"
	@echo "DOCKER_STUBBY_TENANT_ID: $(DOCKER_STUBBY_TENANT_ID)"
	@echo

run_unit_tests:
	make install_test_dep
	sh ./ci/unit_tests/run_unit_tests.sh

run_local_stubby_tests: debug_local_environment_variables
	make install_test_dep
	SPLUNK_CLOUD_HOST=$(LOCAL_SPLUNK_CLOUD_HOST) \
	BEARER_TOKEN=$(LOCAL_BEARER_TOKEN) \
	TENANT_ID=$(LOCAL_TENANT_ID) \
	sh ./ci/functional/runtests.sh

run_docker_stubby_tests: debug_docker_environment_variables
	make install_test_dep
	SPLUNK_CLOUD_HOST=$(DOCKER_STUBBY_SPLUNK_CLOUD_HOST) \
	BEARER_TOKEN=$(DOCKER_STUBBY_BEARER_TOKEN) \
	TENANT_ID=$(DOCKER_STUBBY_TENANT_ID) \
	sh ./ci/functional/runtests.sh

.FORCE: