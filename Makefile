# Copyright 2018 Splunk Inc.
# SPLUNK CONFIDENTIAL - Use or disclosure of this material in whole or in part
# without a valid written license from Splunk Inc. is PROHIBITED.

.DEFAULT_GOAL := noop
GOLANGCI_VER:=1.12.3

GO_NON_TEST_NON_VENDOR_PACKAGES := $(shell go list ./... | grep -v /vendor/ | grep -v test)

setup: prereqs dependencies

lint:
	golangci-lint run ./... --skip-dirs test --enable golint --disable megacheck

linttest:
	golangci-lint run test/... --disable-all

build:
	go build $(GO_NON_TEST_NON_VENDOR_PACKAGES)

clean:
	build

format:
	gofmt -s -w .
	goimports -w .

format_check:
	echo "Checking gofmt / goimports, if this fails you need to re-run 'make format' ..."
	test -z $(shell gofmt -l .)
	test -z $(shell goimports -l .)

vet: install_test_dep
	go vet ./...

install_dep:
	go get -u github.com/golang/dep/cmd/dep

install_test_dep:
	go get -u github.com/stretchr/testify

dependencies:
	@echo "Ensuring dependencies .."
	@rm -rf vendor/
	@rm -rf .vendor-new/
	@go get github.com/vburenin/ifacemaker
	@dep ensure -vendor-only

dependencies_update:
	@echo "Updating dep files (Gopkg.toml. Gopkg.lock) .."
	@rm -rf vendor/
	@rm -rf .vendor-new/
	@dep ensure -update
	@echo "Rebuilding mod files (go.mod, go.sum) .."
	@rm go.mod go.sum
	GO111MODULE=on go mod init
	GO111MODULE=on go mod tidy
	@make dependencies

generate_interface:
	cd services && go generate

prereqs:
	@echo "Installing dep .."
	@DEP_RELEASE_TAG=0.5.0 curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	@echo "Installing goimports ..."
	@go get golang.org/x/tools/cmd/goimports
	@echo "Installing golangci-lint .."
	@curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sudo sh -s  -- -b $(shell go env GOPATH)/bin v$(GOLANGCI_VER)

# This is a generic target that should invoke all levels of tests, i.e. unit tests, integration tests.
test: test_unit test_integration

test_unit:
	make install_test_dep
	sh ./ci/unit_tests/run_unit_tests.sh

test_integration:
	make install_test_dep
	sh ./ci/integration/runtests.sh

test_integration_examples:
	sh ./ci/integration/runexamples.sh

test_ml_integration_tests:
	make install_test_dep
	sh ./ci/integration/run_ml_tests.sh

release: .FORCE
	./cd/release.sh

docs: docs_md

docs_md: .FORCE
	./ci/docs/docs_md.sh

docs_publish: docs_md
	./ci/docs/publish.sh

publish:
	./cd/publish.sh
	make docs_publish

.FORCE: