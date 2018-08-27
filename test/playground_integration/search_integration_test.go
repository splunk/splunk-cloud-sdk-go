// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package playgroundintegration

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"strconv"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
)

const DefaultSearchQuery = "| from index:main | head 5"

var timeout uint = 5

var (
	PostJobsRequest                        = &model.PostJobsRequest{Search: DefaultSearchQuery}
	PostJobsRequestBadRequest              = &model.PostJobsRequest{Search: "hahdkfdksf=main | dfsdfdshead 5"}
	PostJobsRequestTimeout                 = &model.PostJobsRequest{Search: DefaultSearchQuery, Timeout: &timeout}
	PostJobsRequestTTL                     = &model.PostJobsRequest{Search: DefaultSearchQuery, TTL: 5}
	PostJobsRequestLimit                   = &model.PostJobsRequest{Search: DefaultSearchQuery, Limit: 10}
	PostJobsRequestDisableAutoFinalization = &model.PostJobsRequest{Search: DefaultSearchQuery, Limit: 0}
	PostJobsRequestMultiArgs               = &model.PostJobsRequest{Search: DefaultSearchQuery, Timeout: &timeout, TTL: 10, Limit: 10}
	PostJobsRequestLowThresholds           = &model.PostJobsRequest{Search: DefaultSearchQuery, Timeout: &timeout, TTL: 1}
	PostJobsRequestModule                  = &model.PostJobsRequest{Search: DefaultSearchQuery, Module: ""} // Empty string until catalog is updated
)

func TestGetJobsDefaultParams(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.GetJobs(nil)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestGetJobsCustomParams(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.GetJobs(&model.JobsRequest{Count: 1, Offset: 0})
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, len(response), 1)
}

func TestGetJob(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	sid, err := client.SearchService.CreateJob(PostJobsRequest)
	assert.Emptyf(t, err, "Error creating job: %s", err)
	response, err := client.SearchService.GetJob(sid)
	assert.Nil(t, err)
	err = client.SearchService.WaitForJob(sid, 1000*time.Millisecond)
	assert.Emptyf(t, err, "Error waiting for job: %s", err)
	assert.NotEmpty(t, response)
}

func TestGetJobWithModule(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	sid, err := client.SearchService.CreateJob(PostJobsRequestModule)
	assert.Emptyf(t, err, "Error creating job: %s", err)
	response, err := client.SearchService.GetJob(sid)
	assert.Nil(t, err)
	err = client.SearchService.WaitForJob(sid, 1000*time.Millisecond)
	assert.Emptyf(t, err, "Error waiting for job: %s", err)
	assert.NotEmpty(t, response)
}

func TestGetJobResults(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	sid, err := client.SearchService.CreateJob(PostJobsRequest)
	assert.Emptyf(t, err, "Error creating job: %s", err)
	err = client.SearchService.WaitForJob(sid, 1000*time.Millisecond)
	assert.Emptyf(t, err, "Error waiting for job: %s", err)
	response, err := client.SearchService.GetJobResults(sid, &model.FetchResultsRequest{Count: 5})
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, 5, len(response.Results))
}

func TestGetJobEvents(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	sid, err := client.SearchService.CreateJob(PostJobsRequest)
	assert.Emptyf(t, err, "Error creating job: %s", err)
	err = client.SearchService.WaitForJob(sid, 1000*time.Millisecond)
	assert.Emptyf(t, err, "Error waiting for job: %s", err)
	response, err := client.SearchService.GetJobEvents(sid, &model.FetchEventsRequest{Count: 5})
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, 5, len(response.Results))
}

// TestIntegrationNewSearchJob asynchronously
func TestIntegrationNewSearchJob(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
}

//
// TestIntegrationNewSearchJobBadRequest asynchronously
func TestIntegrationNewSearchJobBadRequest(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestBadRequest)
	// HTTP 400 Error Code
	expectedError := &util.HTTPError{HTTPStatusCode:400, Message:"{\"type\":\"ERROR_SPL_PARSE\",\"reason\":\"no viable alternative at input '|searchhahdkfdksf=main|dfsdfdshead'\",\"rule\":\"search\",\"line\":1,\"position\":27,\"token\":\"dfsdfdshead\",\"ok\":false}", Code:"1019"}

	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	assert.Empty(t, response)
}

// TestIntegrationNewSearchJobDuplicates
func TestIntegrationNewSearchJobDuplicates(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	response, err = client.SearchService.CreateJob(PostJobsRequest)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
}

//
// TestIntegrationNewSearchJobTimeout with timeout at 5 sec
func TestIntegrationNewSearchJobTimeout(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestTimeout)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
}

// TestIntegrationNewSearchJobTTL with TTL at 5 sec
func TestIntegrationNewSearchJobTTL(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestTTL)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
}

// TestIntegrationNewSearchJobLimit with Limit at 10
func TestIntegrationNewSearchJobLimit(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestLimit)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
}

// TestIntegrationNewSearchJobDisableAutoFinalization with Limit at 0, disable automatic finalization
func TestIntegrationNewSearchJobDisableAutoFinalization(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestDisableAutoFinalization)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
}

// TestIntegrationNewSearchJobMultiArgs with multiple args
func TestIntegrationNewSearchJobMultiArgs(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestMultiArgs)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
}

// TestIntegrationGetJobResults
func TestIntegrationGetJobResults(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	sid, err := client.SearchService.CreateJob(PostJobsRequest)
	assert.Nil(t, err)
	assert.NotNil(t, sid)
	err = client.SearchService.WaitForJob(sid, 1000*time.Millisecond)
	assert.Emptyf(t, err, "Error waiting for job: %s", err)
	resp, err := client.SearchService.GetJobResults(sid, &model.FetchResultsRequest{Count: 30})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	validateResponses(resp, t)
}

// TestIntegrationGetJobResultsTTL
func TestIntegrationGetJobResultsTTL(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestTTL)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	err = client.SearchService.WaitForJob(response, 1000*time.Millisecond)
	assert.Emptyf(t, err, "Error waiting for job: %s", err)
	resp, err := client.SearchService.GetJobResults(response, &model.FetchResultsRequest{Count: 30})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	validateResponses(resp, t)
}

// TestIntegrationGetJobResultsLimit
func TestIntegrationGetJobResultsLimit(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestLimit)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	err = client.SearchService.WaitForJob(response, 1000*time.Millisecond)
	assert.Emptyf(t, err, "Error waiting for job: %s", err)
	resp, err := client.SearchService.GetJobResults(response, &model.FetchResultsRequest{Count: 30})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	validateResponses(resp, t)
}

// TestIntegrationGetJobResultsDisableAutoFinalization
func TestIntegrationGetJobResultsDisableAutoFinalization(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestDisableAutoFinalization)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	err = client.SearchService.WaitForJob(response, 1000*time.Millisecond)
	assert.Emptyf(t, err, "Error waiting for job: %s", err)
	resp, err := client.SearchService.GetJobResults(response, &model.FetchResultsRequest{Count: 30})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	validateResponses(resp, t)
}

// TestIntegrationGetJobResultsMultipleArgs
func TestIntegrationGetJobResultsMultipleArgs(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestMultiArgs)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	err = client.SearchService.WaitForJob(response, 1000*time.Millisecond)
	assert.Emptyf(t, err, "Error waiting for job: %s", err)
	resp, err := client.SearchService.GetJobResults(response, &model.FetchResultsRequest{Count: 30})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	validateResponses(resp, t)
}

// TestIntegrationGetJobResultsLowThresholds
func TestIntegrationGetJobResultsLowThresholds(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestLowThresholds)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	err = client.SearchService.WaitForJob(response, 1000*time.Millisecond)
	assert.Emptyf(t, err, "Error waiting for job: %s", err)
	resp, err := client.SearchService.GetJobResults(response, &model.FetchResultsRequest{Count: 30})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	validateResponses(resp, t)
}

// TestIntegrationGetJobResultsBadSearchID
func TestIntegrationGetJobResultsBadSearchID(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	// HTTP Code 500 Error
	expectedError := &util.HTTPError{HTTPStatusCode: 404, Message: "404 Not Found", Code:"404"}

	resp, err := client.SearchService.GetJobResults("NON_EXISTING_SEARCH_ID", &model.FetchResultsRequest{Count: 30})
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	// empty SearchResults
	var expectedSearchEvent *model.SearchResults
	assert.Nil(t, resp)
	assert.EqualValues(t, expectedSearchEvent, resp)
}

func TestQueryEvents(t *testing.T) {
	client := getClient(t)
	search, err := client.SearchService.SubmitSearch(PostJobsRequest)
	assert.Emptyf(t, err, "Error submitting search: %s", err)
	pages, err := search.QueryEvents(2, 0, &model.FetchEventsRequest{Count: 5})
	assert.Emptyf(t, err, "Error querying events: %s", err)
	defer pages.Close()
	for pages.Next() {
		values, err := pages.Value()
		assert.Emptyf(t, err, "Error calling pages.Value(): %s", err)
		assert.NotNil(t, values)
	}
	err = pages.Err()
	assert.Nil(t, err)
}

func TestQueryResults(t *testing.T) {
	client := getClient(t)
	search, err := client.SearchService.SubmitSearch(PostJobsRequest)
	assert.Emptyf(t, err, "Error submitting search: %s", err)
	pages, err := search.QueryResults(3, 0, &model.FetchResultsRequest{Count: 5})
	assert.Emptyf(t, err, "Error querying events: %s", err)
	defer pages.Close()
	for pages.Next() {
		values, err := pages.Value()
		assert.Emptyf(t, err, "Error calling pages.Value(): %s", err)
		assert.NotNil(t, values)
	}
	err = pages.Err()
	assert.Nil(t, err)
}

// retry
func retry(attempts int, sleep time.Duration, callback func() (interface{}, error)) error {
	var err error
	for i := 1; i <= attempts; i++ {
		fmt.Println("Retry Attempts: " + strconv.Itoa(i))
		_, err = callback()
		if err != nil {
			time.Sleep(sleep)
		} else {
			// stop retrying
			return nil
		}
	}
	return err
}

// validateResponse
func validateResponses(response *model.SearchResults, t *testing.T) {
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
