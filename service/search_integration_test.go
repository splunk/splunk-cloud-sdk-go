// +build !integration

package service

import (
	"os"
	"testing"
	"time"

	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
)

var hostID = os.Getenv("SSC_HOST")
var token = os.Getenv("BEARER_TOKEN")
var tenantID = os.Getenv("TENANT_ID")

func getSplunkClientForPlaygroundTests() *Client {
	return NewClient(tenantID, token, hostID, time.Second*5)
}

func TestIntegrationEnvironment(t *testing.T) {
	assert.NotEmpty(t, token)
	assert.NotEmpty(t, tenantID)
}

func TestIntegrationNewSearchJob(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateJob(&model.PostJobsRequest{Query: "search index=_internal | head 5"})
	assert.Nil(t, err)
	assert.NotEmpty(t, response.SearchID)
}

func TestIntegrationNewSearchJobSync(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateSyncJob(&model.PostJobsRequest{Query: "search index=_internal | head 5"})
	assert.Nil(t, err)
	assert.NotNil(t, response)
	ValidateResponses(response, t)
}

func TestIntegrationGetJobResults(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateJob(&model.PostJobsRequest{Query: "search index=_internal | head 5"})
	assert.Nil(t, err)
	assert.NotNil(t, response)

	time.Sleep(3000 * time.Millisecond)

	resp, err := client.SearchService.GetResults(response.SearchID)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	ValidateResponses(resp, t)
}

// Validate response results
func ValidateResponses(response *model.SearchEvents, t *testing.T) {
	indexFound := false
	if response.Fields != nil {
		for _, v := range response.Fields {
			for m, n := range v {
				if m == "name" && n == "index" {
					indexFound = true
				}
			}
		}
		if !indexFound {
			t.Errorf("Expected results field element name and corresponding value index not found")
		}
	} else {
		t.Errorf("Expected field elements in results not found")
	}

	if response.Results == nil {
		t.Errorf("Invalid response, missing results in response returned")
	}
}
