// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"testing"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/service"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const DefaultSearchQuery = "| from index:main | head 5"

var (
	PostJobsRequest             = &model.CreateJobRequest{Query: DefaultSearchQuery}
	PostJobsBadRequest          = &model.CreateJobRequest{Query: "hahdkfdksf=main | dfsdfdshead 5"}
	PostJobsRequestModule       = &model.CreateJobRequest{Query: DefaultSearchQuery, Module: ""} // Empty string until catalog is updated
	QueryParams                 = &model.QueryParameters{Earliest: "-12h@h"}
	PostJobsRequestWithEarliest = &model.CreateJobRequest{Query: DefaultSearchQuery, QueryParameters: QueryParams}
)

func TestListJobs(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	response, err := client.SearchService.ListJobs()
	require.Nil(t, err)
	assert.NotNil(t, response)
}

func TestListJobsByStatusRunning(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	response, err := client.SearchService.ListJobsByQueryParameters(service.JobsQuery{Status: "running"})
	require.Nil(t, err)
	assert.NotNil(t, response)
}

func TestListJobsByMultipleStatuses(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	response, err := client.SearchService.ListJobsByQueryParameters(service.JobsQuery{Status: "running, done"})
	require.Nil(t, err)
	assert.NotNil(t, response)
}

func TestGetJob(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	job, err := client.SearchService.CreateJob(PostJobsRequest)
	require.Emptyf(t, err, "Error creating job: %s", err)
	response, err := client.SearchService.GetJob(job.ID)
	assert.Nil(t, err)
	require.NotEmpty(t, response)
	assert.NotNil(t, response.Messages)
	assert.Equal(t, job.ID, response.ID)
	assert.NotEmpty(t, response.Status)
	assert.Equal(t, DefaultSearchQuery, response.Query)
}

func TestCreateJobWithTimerange(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestWithEarliest)
	assert.Nil(t, err)
	require.NotEmpty(t, response)
	assert.Equal(t, PostJobsRequest.Query, response.Query)
	assert.Equal(t, model.Running, response.Status)
	assert.Equal(t, "-12h@h", response.QueryParameters.Earliest)
}

func TestCreateJobWithModule(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	job, err := client.SearchService.CreateJob(PostJobsRequestModule)
	require.Emptyf(t, err, "Error creating job: %s", err)
	response, err := client.SearchService.GetJob(job.ID)
	assert.Nil(t, err)
	require.NotEmpty(t, response)
	assert.NotNil(t, response.Messages)
	assert.Equal(t, job.ID, response.ID)
	assert.NotEmpty(t, response.Status)
	assert.Equal(t, PostJobsRequestModule.Query, response.Query)
}

func TestUpdateJobToBeCanceled(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	job, err := client.SearchService.CreateJob(PostJobsRequest)
	require.Emptyf(t, err, "Error creating job: %s", err)
	patchResponse, err := client.SearchService.UpdateJob(job.ID, model.JobCanceled)
	assert.Nil(t, err)
	require.NotEmpty(t, patchResponse)
	assert.Equal(t, "INFO", patchResponse.Messages[0].Type)
	assert.Equal(t, "Search job cancelled.", patchResponse.Messages[0].Text)
	_, err = client.SearchService.GetJob(job.ID)
	assert.Equal(t, 404, err.(*util.HTTPError).HTTPStatusCode)
	assert.Equal(t, "404 Not Found", err.(*util.HTTPError).HTTPStatus)
	assert.Equal(t, "Failed to get search job status by sid", err.(*util.HTTPError).Message)
	assert.NotNil(t, err)
}

func TestUpdateJobToBeFinalized(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	job, err := client.SearchService.CreateJob(PostJobsRequest)
	require.Emptyf(t, err, "Error creating job: %s", err)
	patchResponse, err := client.SearchService.UpdateJob(job.ID, model.JobFinalized)
	assert.Nil(t, err)
	require.NotEmpty(t, patchResponse)
	assert.Equal(t, "INFO", patchResponse.Messages[0].Type)
	assert.Equal(t, "Search job finalized.", patchResponse.Messages[0].Text)
}

func TestGetJobResultsNextLink(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	job, err := client.SearchService.CreateJob(PostJobsRequest)
	require.Emptyf(t, err, "Error creating job: %s", err)
	response, err := client.SearchService.GetResults(job.ID, 0, 0)
	require.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.NotEmpty(t, response.(*model.ResultsNotReadyResponse).NextLink)
}

func TestGetJobResults(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	job, err := client.SearchService.CreateJob(PostJobsRequest)
	require.Emptyf(t, err, "Error creating job: %s", err)
	state, err := client.SearchService.WaitForJob(job.ID, 1000*time.Millisecond)
	require.Emptyf(t, err, "Error waiting for job: %s", err)
	assert.Equal(t, model.Done, state)
	response, err := client.SearchService.GetResults(job.ID, 5, 0)
	assert.Nil(t, err)
	require.NotEmpty(t, response)
	assert.Equal(t, 5, len(response.(*model.SearchResults).Results))
}

// TestIntegrationNewSearchJobBadRequest asynchronously
func TestIntegrationNewSearchJobBadRequest(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsBadRequest)
	require.NotNil(t, err)
	assert.Empty(t, response)
	assert.Equal(t, 400, err.(*util.HTTPError).HTTPStatusCode)
	assert.Equal(t, "400 Bad Request", err.(*util.HTTPError).HTTPStatus)
}

// TestIntegrationGetJobResultsBadSearchID
func TestIntegrationGetJobResultsBadSearchID(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	resp, err := client.SearchService.GetResults("NON_EXISTING_SEARCH_ID", 0, 0)
	require.NotNil(t, err)
	assert.Equal(t, 404, err.(*util.HTTPError).HTTPStatusCode)
	assert.Equal(t, "404 Not Found", err.(*util.HTTPError).HTTPStatus)
	assert.Equal(t, "Failed to list search results.", err.(*util.HTTPError).Message)
	assert.Nil(t, resp)
}
