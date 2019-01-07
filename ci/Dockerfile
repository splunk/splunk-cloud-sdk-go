# Starting from the latest Golang image
FROM golang:1.11-alpine

# INSTALL any further tools you need here so they are cached in the docker build
ENV DEP_VERSION=0.4.1
ENV LINT_VERSION=1.12.5
RUN apk update && \
    apk upgrade && \
    apk add --no-cache bash \
                       # Needed for gcc in order to support Go
                       build-base \
                       curl \
                       git \
                       jq \
                       make \
                       openssh

RUN curl --fail \
         --silent \
         --show-error \
         --location \
         --output /usr/local/bin/dep \
         --url https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 \
         -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v${LINT_VERSION} \
    && chmod +x /usr/local/bin/dep

# Set the WORKDIR to the project path in your GOPATH
WORKDIR /go/src/github.com/splunk/splunk-cloud-sdk-go

# Copy the content of your repository into the image
COPY . ./

# Install dependencies through go get, unless you vendored them in your repository before
# Vendoring can be done with an external tool like godep or glide
# Go versions after 1.5.1 include support for a vendor directory
RUN make dependencies

VOLUME /go/src/github.com/splunk/splunk-cloud-sdk-go
