// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package playgroundintegration

import (
	"fmt"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"strings"
	"time"
)

const DefaultSearchQuery = "| from index:main | head 5"

var timeout uint = 5

var (
	PostJobsRequest             = &model.CreateJobRequest{Query: DefaultSearchQuery}
	PostJobsBadRequest          = &model.CreateJobRequest{Query: "hahdkfdksf=main | dfsdfdshead 5"}
	PostJobsRequestModule       = &model.CreateJobRequest{Query: DefaultSearchQuery, Module: ""} // Empty string until catalog is updated
	PostJobsRequestWithEarliest = &model.CreateJobRequest{Query: DefaultSearchQuery, Earliest: "12h@h"}
)

func TestListJobs(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.ListJobs()
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestGetJob(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	job, err := client.SearchService.CreateJob(PostJobsRequest)
	assert.Emptyf(t, err, "Error creating job: %s", err)
	response, err := client.SearchService.GetJob(job.Id)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.NotNil(t, response.Messages)
	assert.Equal(t, job.Id, response.Id)
	assert.NotEmpty(t, response.Status)
	assert.Equal(t, DefaultSearchQuery, response.Query)
}

func TestCreateJobWithTimerange(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestWithEarliest)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, PostJobsRequest.Query, response.Query)
	assert.Equal(t, model.Running, response.Status)
	assert.Equal(t, "12h@h", response.Parameters.Earliest)
}

func TestCreateJobWithModule(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	job, err := client.SearchService.CreateJob(PostJobsRequestModule)
	assert.Emptyf(t, err, "Error creating job: %s", err)
	response, err := client.SearchService.GetJob(job.Id)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.NotNil(t, response.Messages)
	assert.Equal(t, job.Id, response.Id)
	assert.NotEmpty(t, response.Status)
	assert.Equal(t, PostJobsRequestModule, response.Query)
}

func TestUpdateJob(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	job, err := client.SearchService.CreateJob(PostJobsRequest)
	assert.Emptyf(t, err, "Error creating job: %s", err)
	patchResponse, err := client.SearchService.UpdateJob(job.Id, model.CancelJob)
	assert.Nil(t, err)
	assert.Equal(t, "INFO", patchResponse.Messages[0].Type)
	assert.Equal(t, "Search job cancelled.", patchResponse.Messages[0].Text)
	assert.NotEmpty(t, patchResponse)
	_, err = client.SearchService.GetJob(job.Id)
	fmt.Println(job.Id)
	assert.Equal(t, 404, err.(*util.HTTPError).HTTPStatusCode)
	assert.Equal(t, "404 Not Found", err.(*util.HTTPError).Message)
	assert.NotNil(t, err)
}

func TestGetJobResultsNextLink(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	job, err := client.SearchService.CreateJob(PostJobsRequest)
	assert.Emptyf(t, err, "Error creating job: %s", err)
	response, err := client.SearchService.GetResults(job.Id, 0, 0)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.NotEmpty(t, response.(*model.ResultsNotReadyResponse).NextLink)
}

func TestGetJobResults(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	job, err := client.SearchService.CreateJob(PostJobsRequest)
	assert.Emptyf(t, err, "Error creating job: %s", err)
	state, err := client.SearchService.WaitForJob(job.Id, 1000*time.Millisecond)
	assert.Emptyf(t, err, "Error waiting for job: %s", err)
	assert.Equal(t, model.Done, state)
	response, err := client.SearchService.GetResults(job.Id, 5, 0)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, 5, response.(*model.SearchResults).Results)
}

// TestIntegrationNewSearchJobBadRequest asynchronously
func TestIntegrationNewSearchJobBadRequest(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsBadRequest)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.(*util.HTTPError).Message, "ERROR_SPL_PARSE"))
	assert.Equal(t, "400", err.(*util.HTTPError).HTTPStatusCode)
	assert.Empty(t, response)
}

// TestIntegrationGetJobResultsBadSearchID
func TestIntegrationGetJobResultsBadSearchID(t *testing.T) {
	client := getClient(t)
	assert.NotNil(t, client)
	expectedError := &util.HTTPError{HTTPStatusCode: 404, HTTPStatus: "404 Not Found", Code: "404"}

	resp, err := client.SearchService.GetResults("NON_EXISTING_SEARCH_ID", 0, 0)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
	// empty SearchResults
	var expectedSearchEvent *model.SearchResults
	assert.Nil(t, resp)
	assert.EqualValues(t, expectedSearchEvent, resp)
}
