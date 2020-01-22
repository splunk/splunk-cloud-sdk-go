# Copyright 2018 Splunk Inc.
# SPLUNK CONFIDENTIAL - Use or disclosure of this material in whole or in part
# without a valid written license from Splunk Inc. is PROHIBITED.

.DEFAULT_GOAL:=build

NON_VENDOR_GO_FILES:=$(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -name "version.go")

SCLOUD_SRC_PATH:=github.com/splunk/splunk-cloud-sdk-go/cmd/scloud
SCLOUD_GEN_SRC_PATH:=github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/cmd
CONFIG_VER_FILE:=cmd/scloud/cli/static/config.ver
SCLOUD_CONFIG_VERSION:=$(shell cat $(CONFIG_VER_FILE))

setup: prereqs download_config statik version

lint: linttest
	# vendor/ needed for golangci-lint to work at the moment
	GO111MODULE=on go mod vendor
	golangci-lint run ./... --skip-dirs test --skip-files ".*_generated.go"  --skip-files "interface.go" --enable golint --disable megacheck
	rm -rf vendor/

linttest: statik version
	# vendor/ needed for golangci-lint to work at the moment
	GO111MODULE=on go mod vendor
	golangci-lint run test/... --disable-all
	rm -rf vendor/

build: statik version
	GO111MODULE=on go build -v ./...
	make build_scloud
	make build_scloud_generated

build_scloud: statik version
	@echo "Building scloud .."
	GO111MODULE=on go build -v -o bin/scloud $(SCLOUD_SRC_PATH)/cli

build_scloud_generated: statik version
	@echo "Building scloud generated.."
	GO111MODULE=on go build -v -o bin/scloud_gen $(SCLOUD_GEN_SRC_PATH)/scloud/.

build_cross_compile:
	SCLOUD_SRC_PATH=$(SCLOUD_SRC_PATH) ./cicd/scripts/build_cross_compile.sh

format:
	GO111MODULE=on gofmt -s -w $(NON_VENDOR_GO_FILES)
	GO111MODULE=on goimports -w $(NON_VENDOR_GO_FILES)

format_check:
	echo "Checking gofmt / goimports, if this fails you need to re-run 'make format' ..."
	test -z $(shell GO111MODULE=on gofmt -l $(NON_VENDOR_GO_FILES))
	test -z $(shell GO111MODULE=on goimports -l $(NON_VENDOR_GO_FILES))

vet: statik version
	GO111MODULE=on go vet ./...

login: build_scloud
	./cicd/scripts/login.sh

token:
	./cicd/scripts/token.sh

clean: download_config
	@rm -rf bin/
	build

generate_interface:
	@GO111MODULE=off && go get github.com/vburenin/ifacemaker && cd services && GO111MODULE=on go generate

download_config:
	@echo "Downloading current config ($(SCLOUD_CONFIG_VERSION)) from artifactory ..."
ifndef SKIP_DOWNLOAD_CONFIG
	@test -n "$(ARTIFACTORY_SCLOUD_LOC)" || (echo "ARTIFACTORY_SCLOUD_LOC must be set ..." && exit 1)
	@curl -s "$(ARTIFACTORY_SCLOUD_LOC)/config/$(SCLOUD_CONFIG_VERSION)/default.yaml" -o "./cmd/scloud/cli/static/default.yaml.$(SCLOUD_CONFIG_VERSION)"
	@test -f "./cmd/scloud/cli/static/default.yaml.$(SCLOUD_CONFIG_VERSION)" || (echo "default.yaml not downloaded ..." && exit 1)
	@cat "./cmd/scloud/cli/static/default.yaml.$(SCLOUD_CONFIG_VERSION)" | grep 'environments:' || (echo "default.yaml contents not correct ..." && rm -f "./cmd/scloud/cli/static/default.yaml.$(SCLOUD_CONFIG_VERSION)" && exit 1)
	@rm -f ./cmd/scloud/cli/static/default.yaml
	@mv "./cmd/scloud/cli/static/default.yaml.$(SCLOUD_CONFIG_VERSION)" ./cmd/scloud/cli/static/default.yaml
endif

upload_config:
	@echo "Uploading current config from ./cmd/scloud/cli/static/default.yaml to artifactory ..."
	@test -n "$(ARTIFACTORY_SCLOUD_LOC)" || (echo "ARTIFACTORY_SCLOUD_LOC must be set ..." && exit 1)
	@test -n "$(ARTIFACTORY_TOKEN)" || (echo "ARTIFACTORY_TOKEN must be set ..." && exit 1)
	@cat "./cmd/scloud/cli/static/default.yaml" | grep 'environments:' || (echo "default.yaml contents not correct ..." && exit 1)
	@curl -s -H 'Content-Type:text/plain' -H "X-JFrog-Art-Api: $(ARTIFACTORY_TOKEN)" -X PUT "$(ARTIFACTORY_SCLOUD_LOC)/config/$(DATETIME)/default.yaml" -T ./cmd/scloud/cli/static/default.yaml | jq -e .downloadUri || (echo "Error uploading config ..." && exit 1)
	@echo "Upload successful, updating contents of $(CONFIG_VER_FILE) to contain $(DATETIME) ..."
	@echo "$(DATETIME)" > $(CONFIG_VER_FILE)

prereqs:
	echo "Installing goimports .."
	GO111MODULE=off go get golang.org/x/tools/cmd/goimports
	echo "Installing statik .."
	GO111MODULE=off go get github.com/rakyll/statik
	echo "Installing golangci-lint .."
	GO111MODULE=off go get github.com/golangci/golangci-lint/cmd/golangci-lint
	@echo "Installing gotestsum ..."
	@go get gotest.tools/gotestsum

statik:
	@echo "Generate static assets .."
	@statik -src=$(CURDIR)/cmd/scloud/cli/static -dest $(CURDIR)/cmd/scloud/cli

version:
	@echo "Generate version.go .."
	cd $(CURDIR)/cmd/scloud/cli && go generate

# This is a generic target that should invoke all levels of tests, i.e. unit tests, integration tests.
test: test_unit test_integration

test_unit: build
	GO111MODULE=on sh ./cicd/unit_tests/run_unit_tests.sh

test_integration: build
	GO111MODULE=on sh ./cicd/integration/runtests.sh

test_integration_scloud: login build
	export PYTHONPATH=$(PYTHONPATH):.
	SCLOUD_TEST_DIR=$(shell pwd)/test/scloud GO111MODULE=on sh ./cicd/integration/run_scloud_tests.sh

test_integration_examples: build
	GO111MODULE=on sh ./cicd/integration/runexamples.sh

test_ml_integration_tests: build
	GO111MODULE=on sh ./cicd/integration/run_ml_tests.sh

prerelease: .FORCE
	./cicd/prerelease/prerelease.sh

docs: docs_md

docs_md: .FORCE
	./cicd/docs/docs_md.sh

docs_publish: docs_md
	./cicd/docs/publish.sh

publish:
	./cicd/publish/sdk/publish.sh
	make docs_publish

publish_scloud:
	SIGN_PACKAGES=true SCLOUD_SRC_PATH=$(SCLOUD_SRC_PATH) ./cicd/scripts/build_cross_compile.sh
	./cicd/publish/scloud/publish_github.sh

.FORCE:
