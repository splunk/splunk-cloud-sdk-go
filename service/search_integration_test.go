// +build !integration

package service

import (
	"os"
	"testing"
	"time"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
	"github.com/stretchr/testify/assert"
)

var hostID = os.Getenv("SSC_HOST")
var token = os.Getenv("BEARER_TOKEN")
var tenantID = os.Getenv("TENANT_ID")

const DefaultSearchQuery = "search index=_internal | head 5"

var (
	PostJobsRequest                        = &model.PostJobsRequest{Query: DefaultSearchQuery}
	PostJobsRequestBadRequest              = &model.PostJobsRequest{}
	PostJobsRequestBadQuery                = &model.PostJobsRequest{Query: "index=_internal | head 5"}
	PostJobsRequestTimeout                 = &model.PostJobsRequest{Query: DefaultSearchQuery, Timeout: 5}
	PostJobsRequestTTL                     = &model.PostJobsRequest{Query: DefaultSearchQuery, TTL: 5}
	PostJobsRequestLimit                   = &model.PostJobsRequest{Query: DefaultSearchQuery, Limit: 10}
	PostJobsRequestDisableAutoFinalization = &model.PostJobsRequest{Query: DefaultSearchQuery, Limit: 0}
	PostJobsRequestMultiArgs               = &model.PostJobsRequest{Query: DefaultSearchQuery, Timeout: 5, TTL: 10, Limit: 10}
)

func getSplunkClientForPlaygroundTests() *Client {
	return NewClient(tenantID, token, hostID, time.Second*5)
}

// TestIntegrationEnvironment for token and tenant
func TestIntegrationEnvironment(t *testing.T) {
	assert.NotEmpty(t, token)
	assert.NotEmpty(t, tenantID)
}

// TestIntegrationNewSearchJob asynchronously
func TestIntegrationNewSearchJob(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateJob(PostJobsRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, response.SearchID)
}

// TestIntegrationNewSearchJobBadRequest asynchronously
func TestIntegrationNewSearchJobBadRequest(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateJob(PostJobsRequestBadRequest)

	// HTTP 400 Error Code
	expectedError := &util.HTTPError{Status: 400, Message: "400 Bad Request"}

	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	assert.Empty(t, response.SearchID)
}

// TestIntegrationNewSearchJobBadQuery asynchronously
func TestIntegrationNewSearchJobBadQuery(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateJob(PostJobsRequestBadQuery)
	assert.Nil(t, err)
	assert.NotEmpty(t, response.SearchID)
}

// TestIntegrationNewSearchJobDuplicates
func TestIntegrationNewSearchJobDuplicates(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateJob(PostJobsRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, response.SearchID)

	response, err = client.SearchService.CreateJob(PostJobsRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, response.SearchID)
}

// TestIntegrationNewSearchJobTimeout with timeout at 5 sec
func TestIntegrationNewSearchJobTimeout(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateJob(PostJobsRequestTimeout)
	assert.Nil(t, err)
	assert.NotEmpty(t, response.SearchID)
}

// TestIntegrationNewSearchJobTTL with TTL at 5 sec
func TestIntegrationNewSearchJobTTL(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateJob(PostJobsRequestTTL)
	assert.Nil(t, err)
	assert.NotEmpty(t, response.SearchID)
}

// TestIntegrationNewSearchJobLimit with Limit at 10
func TestIntegrationNewSearchJobLimit(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateJob(PostJobsRequestLimit)
	assert.Nil(t, err)
	assert.NotEmpty(t, response.SearchID)
}

// TestIntegrationNewSearchJobDisableAutoFinalization with Limit at 0, disable automatic finalization
func TestIntegrationNewSearchJobDisableAutoFinalization(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateJob(PostJobsRequestDisableAutoFinalization)
	assert.Nil(t, err)
	assert.NotEmpty(t, response.SearchID)
}

// TestIntegrationNewSearchJobMultiArgs with multiple args
func TestIntegrationNewSearchJobMultiArgs(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateJob(PostJobsRequestMultiArgs)
	assert.Nil(t, err)
	assert.NotEmpty(t, response.SearchID)
}

// TestIntegrationNewSearchJobSync
func TestIntegrationNewSearchJobSync(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateSyncJob(PostJobsRequest)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	ValidateResponses(response, t)
}

// TestIntegrationNewSearchJobBadRequest asynchronously
func TestIntegrationNewSearchJobSyncBadRequest(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateSyncJob(PostJobsRequestBadRequest)

	// HTTP 400 Error Code
	expectedError := &util.HTTPError{Status: 400, Message: "400 Bad Request"}
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)

	// expected Search Event
	expectedResult := &model.SearchEvents{Preview: false, InitOffset: 0, Messages: []interface{}(nil),
		Results: []*model.Result(nil), Fields: []map[string]interface{}(nil), Highlighted: map[string]interface{}(nil)}
	assert.NotNil(t, response)
	assert.EqualValues(t, expectedResult, response)
}

// TestIntegrationNewSearchJobBadQuery asynchronously
func TestIntegrationNewSearchJobSyncBadQuery(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateJob(PostJobsRequestBadQuery)
	assert.Nil(t, err)
	assert.NotEmpty(t, response.SearchID)
}

// TestIntegrationNewSearchJobSyncDuplicates
func TestIntegrationNewSearchJobSyncDuplicates(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateSyncJob(PostJobsRequest)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	ValidateResponses(response, t)

	response, err = client.SearchService.CreateSyncJob(PostJobsRequest)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	ValidateResponses(response, t)
}

// TestIntegrationNewSearchJobSyncTimeout with timeout at 5 sec
func TestIntegrationNewSearchJobSyncTimeout(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateSyncJob(PostJobsRequestTimeout)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	ValidateResponses(response, t)
}

// TestIntegrationNewSearchJobSyncTTL with TTL at 5 sec
func TestIntegrationNewSearchJobSyncTTL(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateSyncJob(PostJobsRequestTTL)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	ValidateResponses(response, t)
}

// TestIntegrationNewSearchJobSyncLimit with Limit at 10
func TestIntegrationNewSearchJobSyncLimit(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateSyncJob(PostJobsRequestLimit)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	ValidateResponses(response, t)
}

// TestIntegrationNewSearchJobSyncDisableAutoFinalization with Limit at 0, disable automatic finalization
func TestIntegrationNewSearchJobSyncDisableAutoFinalization(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateSyncJob(PostJobsRequestDisableAutoFinalization)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	ValidateResponses(response, t)
}

// TestIntegrationNewSearchJobSyncMultiArgs with multiple args
func TestIntegrationNewSearchJobSyncMultiArgs(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateSyncJob(PostJobsRequestMultiArgs)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	ValidateResponses(response, t)
}

// TestIntegrationGetJobResults
func TestIntegrationGetJobResults(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	response, err := client.SearchService.CreateJob(PostJobsRequest)
	assert.Nil(t, err)
	assert.NotNil(t, response)

	time.Sleep(3000 * time.Millisecond)

	resp, err := client.SearchService.GetResults(response.SearchID)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	ValidateResponses(resp, t)
}

// TestIntegrationGetJobResultsBadSearchID
func TestIntegrationGetJobResultsBadSearchID(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	assert.NotNil(t, client)

	// HTTP Code 500 Error
	expectedError := &util.HTTPError{Status: 500, Message: "500 Internal Server Error"}

	resp, err := client.SearchService.GetResults("NON_EXISTING_SEARCH_ID")
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)

	// empty SearchEvent
	expectedSearchEvent := &model.SearchEvents{Preview: false, InitOffset: 0, Messages: []interface{}(nil),
		Results: []*model.Result(nil), Fields: []map[string]interface{}(nil), Highlighted: map[string]interface{}(nil)}

	assert.NotNil(t, resp)
	assert.EqualValues(t, expectedSearchEvent, resp)
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
