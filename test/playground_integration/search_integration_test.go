package playgroundintegration

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/service"
)

const DefaultSearchQuery = "search index=_internal | head 5"

// var (
// 	PostJobsRequest                        = &model.PostJobsRequest{Query: DefaultSearchQuery}
// 	PostJobsRequestBadRequest              = &model.PostJobsRequest{}
// 	PostJobsRequestBadQuery                = &model.PostJobsRequest{Query: "index=_internal | head 5"}
// 	PostJobsRequestTimeout                 = &model.PostJobsRequest{Query: DefaultSearchQuery, Timeout: 5}
// 	PostJobsRequestTTL                     = &model.PostJobsRequest{Query: DefaultSearchQuery, TTL: 5}
// 	PostJobsRequestLimit                   = &model.PostJobsRequest{Query: DefaultSearchQuery, Limit: 10}
// 	PostJobsRequestDisableAutoFinalization = &model.PostJobsRequest{Query: DefaultSearchQuery, Limit: 0}
// 	PostJobsRequestMultiArgs               = &model.PostJobsRequest{Query: DefaultSearchQuery, Timeout: 5, TTL: 10, Limit: 10}
// 	PostJobsRequestLowThresholds           = &model.PostJobsRequest{Query: DefaultSearchQuery, Timeout: 1, TTL: 1}
// )

func TestGetJobs(t *testing.T) {
	client := service.NewClient("123","","http://0.0.0.0:8443",time.Second * 5)
	response, err := client.SearchService.GetJobs(nil)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestGetJob(t *testing.T) {
	client := service.NewClient("123","","http://0.0.0.0:8443",time.Second * 5)
	sid, _ := client.SearchService.CreateJob(&model.PostJobsRequest{Search: "search index=_internal | head 5"})
	response, err := client.SearchService.GetJob(sid)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
}

func TestPostJobAction(t *testing.T) {
	client := service.NewClient("123","","http://0.0.0.0:8443",time.Second * 5)
	sid, _ := client.SearchService.CreateJob(&model.PostJobsRequest{Search: "search index=_internal | head 5"})
	msg, err := client.SearchService.PostJobControl(sid, &model.JobControlAction{Action:"pause"})
	assert.Nil(t, err)
	assert.NotEmpty(t, msg)
}

func TestGetJobResults(t *testing.T) {
	client := service.NewClient("123","","http://0.0.0.0:8443",time.Second * 5)
	sid, _ := client.SearchService.CreateJob(&model.PostJobsRequest{Search: "search index=_internal | head 5"})
	time.Sleep(time.Second * 5)
	response, err := client.SearchService.GetJobResults(sid, &model.FetchResultsRequest{Count:5, OutputMode:"json"})
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
}

func TestGetJobEvents(t *testing.T) {
	client := service.NewClient("123","","http://0.0.0.0:8443",time.Second * 5)
	sid, _ := client.SearchService.CreateJob(&model.PostJobsRequest{Search: "search index=_internal | head 5"})
	time.Sleep(time.Second * 5)
	response, err := client.SearchService.GetJobResults(sid, &model.FetchResultsRequest{Count:5, OutputMode:"json"})
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
}

// TestIntegrationNewSearchJob asynchronously
func TestIntegrationNewSearchJob(t *testing.T) {
	// client := getClient()
	client := service.NewClient("123","","http://0.0.0.0:8443",time.Second * 5)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(&model.PostJobsRequest{Search: "search index=_internal | head 5"})
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
}
//
// // TestIntegrationNewSearchJobBadRequest asynchronously
// func TestIntegrationNewSearchJobBadRequest(t *testing.T) {
// 	client := getClient()
// 	assert.NotNil(t, client)
//
// 	response, err := client.SearchService.CreateJob(PostJobsRequestBadRequest)
//
// 	// HTTP 400 Error Code
// 	expectedError := &util.HTTPError{Status: 400, Message: "400 Bad Request"}
//
// 	assert.NotNil(t, err)
// 	assert.Equal(t, expectedError, err)
// 	assert.Empty(t, response.SearchID)
// }
//
// // TestIntegrationNewSearchJobBadQuery asynchronously
// func TestIntegrationNewSearchJobBadQuery(t *testing.T) {
// 	client := getClient()
// 	assert.NotNil(t, client)
//
// 	response, err := client.SearchService.CreateJob(PostJobsRequestBadQuery)
// 	assert.Nil(t, err)
// 	assert.NotEmpty(t, response.SearchID)
// }
//
// // TestIntegrationNewSearchJobDuplicates
// func TestIntegrationNewSearchJobDuplicates(t *testing.T) {
// 	client := getClient()
// 	assert.NotNil(t, client)
//
// 	response, err := client.SearchService.CreateJob(PostJobsRequest)
// 	assert.Nil(t, err)
// 	assert.NotEmpty(t, response.SearchID)
//
// 	response, err = client.SearchService.CreateJob(PostJobsRequest)
// 	assert.Nil(t, err)
// 	assert.NotEmpty(t, response.SearchID)
// }
//
// // TestIntegrationNewSearchJobTimeout with timeout at 5 sec
// func TestIntegrationNewSearchJobTimeout(t *testing.T) {
// 	client := getClient()
// 	assert.NotNil(t, client)
//
// 	response, err := client.SearchService.CreateJob(PostJobsRequestTimeout)
// 	assert.Nil(t, err)
// 	assert.NotEmpty(t, response.SearchID)
// }
//
// // TestIntegrationNewSearchJobTTL with TTL at 5 sec
// func TestIntegrationNewSearchJobTTL(t *testing.T) {
// 	client := getClient()
// 	assert.NotNil(t, client)
//
// 	response, err := client.SearchService.CreateJob(PostJobsRequestTTL)
// 	assert.Nil(t, err)
// 	assert.NotEmpty(t, response.SearchID)
// }
//
// // TestIntegrationNewSearchJobLimit with Limit at 10
// func TestIntegrationNewSearchJobLimit(t *testing.T) {
// 	client := getClient()
// 	assert.NotNil(t, client)
//
// 	response, err := client.SearchService.CreateJob(PostJobsRequestLimit)
// 	assert.Nil(t, err)
// 	assert.NotEmpty(t, response.SearchID)
// }
//
// // TestIntegrationNewSearchJobDisableAutoFinalization with Limit at 0, disable automatic finalization
// func TestIntegrationNewSearchJobDisableAutoFinalization(t *testing.T) {
// 	client := getClient()
// 	assert.NotNil(t, client)
//
// 	response, err := client.SearchService.CreateJob(PostJobsRequestDisableAutoFinalization)
// 	assert.Nil(t, err)
// 	assert.NotEmpty(t, response.SearchID)
// }
//
// // TestIntegrationNewSearchJobMultiArgs with multiple args
// func TestIntegrationNewSearchJobMultiArgs(t *testing.T) {
// 	client := getClient()
// 	assert.NotNil(t, client)
//
// 	response, err := client.SearchService.CreateJob(PostJobsRequestMultiArgs)
// 	assert.Nil(t, err)
// 	assert.NotEmpty(t, response.SearchID)
// }
//
// // TestIntegrationNewSearchJobBadQuery asynchronously
// func TestIntegrationNewSearchJobSyncBadQuery(t *testing.T) {
// 	client := getClient()
// 	assert.NotNil(t, client)
//
// 	response, err := client.SearchService.CreateJob(PostJobsRequestBadQuery)
// 	assert.Nil(t, err)
// 	assert.NotEmpty(t, response.SearchID)
// }
// // TestIntegrationGetJobResults
// func TestIntegrationGetJobResults(t *testing.T) {
// 	client := getClient()
// 	assert.NotNil(t, client)
//
// 	response, err := client.SearchService.CreateJob(PostJobsRequest)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, response)
// 	validateGetResults(client, response, t)
// }
//
// // TestIntegrationGetJobResultsTimeout
// func TestIntegrationGetJobResultsTimeout(t *testing.T) {
// 	client := getClient()
// 	assert.NotNil(t, client)
//
// 	response, err := client.SearchService.CreateJob(PostJobsRequestTimeout)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, response)
// 	validateGetResults(client, response, t)
// }
//
// // TestIntegrationGetJobResultsTTL
// func TestIntegrationGetJobResultsTTL(t *testing.T) {
// 	client := getClient()
// 	assert.NotNil(t, client)
//
// 	response, err := client.SearchService.CreateJob(PostJobsRequestTTL)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, response)
// 	validateGetResults(client, response, t)
// }
//
// // TestIntegrationGetJobResultsLimit
// func TestIntegrationGetJobResultsLimit(t *testing.T) {
// 	client := getClient()
// 	assert.NotNil(t, client)
//
// 	response, err := client.SearchService.CreateJob(PostJobsRequestLimit)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, response)
// 	validateGetResults(client, response, t)
// }
//
// // TestIntegrationGetJobResultsDisableAutoFinalization
// func TestIntegrationGetJobResultsDisableAutoFinalization(t *testing.T) {
// 	client := getClient()
// 	assert.NotNil(t, client)
//
// 	response, err := client.SearchService.CreateJob(PostJobsRequestDisableAutoFinalization)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, response)
// 	validateGetResults(client, response, t)
// }
//
// // TestIntegrationGetJobResultsMultipleArgs
// func TestIntegrationGetJobResultsMultipleArgs(t *testing.T) {
// 	client := getClient()
// 	assert.NotNil(t, client)
//
// 	response, err := client.SearchService.CreateJob(PostJobsRequestMultiArgs)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, response)
// 	validateGetResults(client, response, t)
// }
//
// // TestIntegrationGetJobResultsLowThresholds
// func TestIntegrationGetJobResultsLowThresholds(t *testing.T) {
// 	client := getClient()
// 	assert.NotNil(t, client)
//
// 	response, err := client.SearchService.CreateJob(PostJobsRequestLowThresholds)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, response)
//
// 	time.Sleep(10000 * time.Millisecond)
//
// 	resp, err := client.SearchService.GetResults(response.SearchID)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, resp)
// 	validateResponses(resp, t)
// }

// TestIntegrationGetJobResultsBadSearchID
// func TestIntegrationGetJobResultsBadSearchID(t *testing.T) {
// 	client := getClient()
// 	assert.NotNil(t, client)
//
// 	// HTTP Code 500 Error
// 	expectedError := &util.HTTPError{Status: 500, Message: "500 Internal Server Error"}
//
// 	resp, err := client.SearchService.GetResults("NON_EXISTING_SEARCH_ID")
// 	assert.NotNil(t, err)
// 	assert.Equal(t, expectedError, err)
//
// 	// empty SearchEvent
// 	expectedSearchEvent := &model.SearchEvents{Preview: false, InitOffset: 0, Messages: []interface{}(nil),
// 		Results: []*model.Result(nil), Fields: []map[string]interface{}(nil), Highlighted: map[string]interface{}(nil)}
//
// 	assert.NotNil(t, resp)
// 	assert.EqualValues(t, expectedSearchEvent, resp)
// }

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
// func validateGetResults(client *service.Client, response *model.PostJobResponse, t *testing.T) {
// 	var resp *model.SearchEvents
// 	var err error
//
// 	retryError := retry(3, 3000*time.Millisecond, func() (interface{}, error) {
// 		resp, err = client.SearchService.GetResults(response.SearchID)
// 		return resp, err
// 	})
// 	assert.Nil(t, retryError)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, resp)
// 	validateResponses(resp, t)
// }
//
// // validateResponse
// func validateResponses(response *model.SearchEvents, t *testing.T) {
// 	indexFound := false
// 	if response.Fields != nil {
// 		for _, v := range response.Fields {
// 			for m, n := range v {
// 				if m == "name" && n == "index" {
// 					indexFound = true
// 				}
// 			}
// 		}
// 		if !indexFound {
// 			t.Errorf("Expected results field element name and corresponding value index not found")
// 		}
// 	} else {
// 		t.Errorf("Expected field elements in results not found")
// 	}
//
// 	if response.Results == nil {
// 		t.Errorf("Invalid response, missing results in response returned")
// 	}
// }
