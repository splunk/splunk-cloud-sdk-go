package service

import (
	"fmt"
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuildURL(t *testing.T) {
	var apiURLProtocol = "http"
	var apiHost = "example.com"
	var apiPort = "8882"
	var apiURL = apiURLProtocol + "://" + apiHost + ":" + apiPort
	var tenant = "EXAMPLE_TENANT"
	var token = "EXAMPLE_AUTHENTICATION_TOKEN"
	var timeout = time.Second * 5
	var client = NewClient(tenant, token, apiURL, timeout)

	testURL, err := client.BuildURL("services", "search", "jobs")

	assert.Nil(t, err)
	assert.Equal(t, apiHost, testURL.Hostname())
	assert.Equal(t, apiURLProtocol, testURL.Scheme)
	assert.Equal(t, apiPort, testURL.Port())
	assert.Equal(t, fmt.Sprintf("%s%s", tenant, "/services/search/jobs"), testURL.Path)
	assert.Empty(t, testURL.Fragment)
}

func TestNewClient(t *testing.T) {
	var apiURLProtocol = "http"
	var apiHost = "example.com"
	var apiPort = "8882"
	var apiURL = apiURLProtocol + "://" + apiHost + ":" + apiPort
	var tenant = "EXAMPLE_TENANT"
	var token = "EXAMPLE_AUTHENTICATION_TOKEN"
	var timeout = time.Second * 5
	var client = NewClient(tenant, token, apiURL, timeout)
	var searchService = &SearchService{client: client}

	assert.Equal(t, token, client.token)
	assert.Equal(t, apiURL, fmt.Sprintf("%s://%s", client.URL.Scheme, client.URL.Host))
	assert.Equal(t, timeout, client.httpClient.Timeout)
	assert.Equal(t, searchService, client.SearchService)
}
