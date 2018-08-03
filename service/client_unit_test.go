// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBuildURL(t *testing.T) {
	var apiURLProtocol = "http"
	var apiHost = "example.com"
	var apiPort = "8882"
	var apiURL = apiURLProtocol + "://" + apiHost + ":" + apiPort
	var tenant = "EXAMPLE_TENANT"
	var token = "EXAMPLE_AUTHENTICATION_TOKEN"
	var timeout = time.Second * 5
	var client, _ = NewClient(&Config{token, apiURL, tenant, timeout})

	testURL, err := client.BuildURL(nil, "services", "search", "jobs")

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
	var client, err = NewClient(&Config{token, apiURL, tenant, timeout})
	var searchService = &SearchService{client: client}
	var catalogService = &CatalogService{client: client}
	var identityService = &IdentityService{client: client}
	var ingestService = &IngestService{client: client}
	var kvStoreService = &KVStoreService{client: client}
	assert.Nil(t, err)

	clientURL, err := client.GetURL()
	assert.Nil(t, err)
	assert.Equal(t, token, client.config.Token)
	assert.Equal(t, apiURL, fmt.Sprintf("%s://%s", clientURL.Scheme, clientURL.Host))
	assert.Equal(t, timeout, client.httpClient.Timeout)
	assert.Equal(t, searchService, client.SearchService)
	assert.Equal(t, catalogService, client.CatalogService)
	assert.Equal(t, identityService, client.IdentityService)
	assert.Equal(t, ingestService, client.IngestService)
	assert.Equal(t, kvStoreService, client.KVStoreService)
}
