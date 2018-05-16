# Starting from the latest Golang image
FROM golang:1.10-alpine

# INSTALL any further tools you need here so they are cached in the docker build
ENV DEP_VERSION=0.4.1
RUN apk update && apk upgrade && \
    apk add --no-cache bash && \
    apk add --no-cache openssh && \
    apk add --no-cache git && \
    apk add --no-cache make && \
    apk add --no-cache curl && \
    apk add --no-cache jq && \
    curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 && chmod +x /usr/local/bin/dep

# Set the WORKDIR to the project path in your GOPATH
WORKDIR /go/src/github.com/splunk/ssc-client-go

# Copy the content of your repository into the image
COPY . ./

# Install dependencies through go get, unless you vendored them in your repository before
# Vendoring can be done with an external tool like godep or glide
# Go versions after 1.5.1 include support for a vendor directory
RUN make dependencies