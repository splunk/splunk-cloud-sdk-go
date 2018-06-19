package playgroundintegration

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/util"
	"strconv"
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
	sid, _ := client.SearchService.CreateJob(PostJobsRequest)
	response, err := client.SearchService.GetJob(sid)
	assert.Nil(t, err)
	client.SearchService.WaitForJob(sid, 1000*time.Millisecond)
	assert.NotEmpty(t, response)
}

func TestPostJobAction(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	sid, _ := client.SearchService.CreateJob(PostJobsRequest)
	msg, err := client.SearchService.PostJobControl(sid, &model.JobControlAction{Action: model.PAUSE})
	assert.Nil(t, err)
	assert.NotEmpty(t, msg)
}

func TestGetJobResults(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	sid, _ := client.SearchService.CreateJob(PostJobsRequest)
	client.SearchService.WaitForJob(sid, 1000*time.Millisecond)
	response, err := client.SearchService.GetJobResults(sid, &model.FetchResultsRequest{Count: 5, OutputMode: "json"})
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, 5, len(response.Results))
}

func TestGetJobEvents(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	sid, _ := client.SearchService.CreateJob(PostJobsRequest)
	client.SearchService.WaitForJob(sid, 1000*time.Millisecond)
	response, err := client.SearchService.GetJobResults(sid, &model.FetchResultsRequest{Count: 5, OutputMode: "json"})
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
	expectedError := &util.HTTPError{Status: 400, Message: "400 Bad Request"}
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
	validateGetResults(client, sid, t)
}

// TestIntegrationGetJobResultsTTL
func TestIntegrationGetJobResultsTTL(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestTTL)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	validateGetResults(client, response, t)
}

// TestIntegrationGetJobResultsLimit
func TestIntegrationGetJobResultsLimit(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestLimit)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	validateGetResults(client, response, t)
}

// TestIntegrationGetJobResultsDisableAutoFinalization
func TestIntegrationGetJobResultsDisableAutoFinalization(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestDisableAutoFinalization)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	validateGetResults(client, response, t)
}

// TestIntegrationGetJobResultsMultipleArgs
func TestIntegrationGetJobResultsMultipleArgs(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestMultiArgs)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	validateGetResults(client, response, t)
}

// TestIntegrationGetJobResultsLowThresholds
func TestIntegrationGetJobResultsLowThresholds(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestLowThresholds)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	client.SearchService.WaitForJob(response, 1000*time.Millisecond)
	resp, err := client.SearchService.GetJobResults(response, &model.FetchResultsRequest{OutputMode: "json", Count: 30})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	validateResponses(resp, t)
}

// TestIntegrationGetJobResultsBadSearchID
func TestIntegrationGetJobResultsBadSearchID(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	// HTTP Code 500 Error
	expectedError := &util.HTTPError{Status: 404, Message: "404 Not Found"}

	resp, err := client.SearchService.GetJobResults("NON_EXISTING_SEARCH_ID", &model.FetchResultsRequest{OutputMode: "json", Count: 30})
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	// empty SearchResults
	var expectedSearchEvent *model.SearchResults
	assert.Nil(t, resp)
	assert.EqualValues(t, expectedSearchEvent, resp)
}

func TestQueryEvents(t *testing.T) {
	client, _ := service.NewClient(
		"test123",
		"eyJraWQiOiI3cXV0SjFleUR6V2lOeGlTbktsakZHRWhmU0VzWFlMQWt0NUVNbzJaNFk4IiwiYWxnIjoiUlMyNTYifQ.eyJ2ZXIiOjEsImp0aSI6IkFULms1eUVMNzJzTGxSYkR0c1dXVmtkUmxIay1YemU2cVIxMEVBOXVvSS1sbDQiLCJpc3MiOiJodHRwczovL3NwbHVuay1jaWFtLm9rdGEuY29tL29hdXRoMi9kZWZhdWx0IiwiYXVkIjoiYXBpOi8vZGVmYXVsdCIsImlhdCI6MTUyOTM1Nzg0NSwiZXhwIjoxNTI5NDAxMDQ1LCJjaWQiOiIwb2FwYmcyem1MYW1wV2daNDJwNiIsInVpZCI6IjAwdXpsMHdlZFdxM2tvWEFDMnA2Iiwic2NwIjpbInByb2ZpbGUiLCJlbWFpbCIsIm9wZW5pZCJdLCJzdWIiOiJ4Y2hlbmdAc3BsdW5rLmNvbSJ9.TMkxl81pCILkeNaLJ31lOITyXSuAnhKDnse7sOAQ8qOvPTywgSuRfrhO6casTVagS_WF421TfGEHTGk1Mdar8ZCddueDUY1JJMQuuocfM600uHE4tKPFC-gUqQgg32RlWPJZL1RQyMBIWd6u92rCL5PFJb2KkX7kFBaePsdW_xqlJYTMuh688znL858y9MlQ8gyjzx3hVAjHPCoBhx11kRo1W7BkZ73bkX_S_k3To93_6E_rFz9arACMl9sr64Yb5maPe8PgXVXhyLED90qWUrKo5bj3DnV-gXPLUe4J-s4uECCDo9fQc76Rd5XfSd85jtRwCQmR0Jg2IifbpggpfQ",
		"https://gateway.splunknovadev-playground.com",
		5*time.Second,
	)
	search, _ := client.SearchService.SubmitSearch(PostJobsRequest)
	pages, _ := search.QueryEvents(2, 0, &model.FetchEventsRequest{Count: 5})
	defer pages.Close()
	for pages.Next() {
		values, _ := pages.Value()
		fmt.Printf("%#v", values)
		assert.NotNil(t, values)
		assert.Equal(t, 2, len(values.Results))
	}
	err := pages.Err()
	assert.Nil(t, err)
}

func TestQueryResults(t *testing.T) {
	client := getClient(t)
	search, _ := client.SearchService.SubmitSearch(PostJobsRequest)
	pages, _ := search.QueryResults(2, 0, &model.FetchResultsRequest{Count: 5})
	defer pages.Close()
	for pages.Next() {
		values, _ := pages.Value()
		assert.NotNil(t, values)
		assert.Equal(t, 2, len(values.Results))
	}
	err := pages.Err()
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

// validateGetResults tests the GetResults calls, tries 3x before giving up
func validateGetResults(client *service.Client, sid string, t *testing.T) {
	var resp *model.SearchResults
	var err error

	retryError := retry(3, 3000*time.Millisecond, func() (interface{}, error) {
		resp, err = client.SearchService.GetJobResults(sid, &model.FetchResultsRequest{OutputMode: "json", Count: 30})
		return resp, err
	})
	assert.Nil(t, retryError)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	validateResponses(resp, t)
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
