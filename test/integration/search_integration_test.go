/*
 * Copyright 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package integration

import (
	"sync"
	"testing"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/services/search"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const DefaultSearchQuery = "| from index:main | head 5"

var (
	PostJobsRequest             = search.SearchJob{Query: DefaultSearchQuery}
	PostJobsBadRequest          = search.SearchJob{Query: "hahdkfdksf=main | dfsdfdshead 5"}
	ModuleName                  = ""
	PostJobsRequestModule       = search.SearchJob{Query: DefaultSearchQuery, Module: &ModuleName} // Empty string until catalog is updated
	earliest                    = "-12h@h"
	QueryParams                 = search.QueryParameters{Earliest: &earliest}
	PostJobsRequestWithEarliest = search.SearchJob{Query: DefaultSearchQuery, QueryParameters: &QueryParams}
)

func createSearchJob(client *sdk.Client, postJobsRequest search.SearchJob, t *testing.T) *search.SearchJob {
	job, err := client.SearchService.CreateJob(postJobsRequest)
	require.Emptyf(t, err, "Error creating job: %s", err)
	state, err := client.SearchService.WaitForJob(*job.Sid, 1000*time.Millisecond)
	require.Emptyf(t, err, "Error waiting for job: %s", err)
	assert.Equal(t, search.SearchStatusDone, state)
	return job
}

func TestListJobs(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)

	query := search.ListJobsQueryParams{}.SetCount(0).SetStatus(search.SearchStatusRunning)
	response, err := client.SearchService.ListJobs(&query)
	require.NoError(t, err)
	assert.NotNil(t, response)
}

func TestGetJob(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	job, err := client.SearchService.CreateJob(PostJobsRequest)
	require.Emptyf(t, err, "Error creating job: %s", err)
	response, err := client.SearchService.GetJob(*job.Sid)
	assert.NoError(t, err)
	require.NotEmpty(t, response)
	assert.Equal(t, job.Sid, response.Sid)
	assert.NotEmpty(t, response.Status)
	assert.Equal(t, DefaultSearchQuery, response.Query)
}

func TestCreateJobWithTimerange(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsRequestWithEarliest)
	assert.NoError(t, err)
	require.NotEmpty(t, response)
	assert.Equal(t, PostJobsRequest.Query, response.Query)
	assert.Equal(t, search.SearchStatusRunning, *response.Status)
	assert.Equal(t, "-12h@h", *response.QueryParameters.Earliest)
}

func TestCreateJobWithModule(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	job, err := client.SearchService.CreateJob(PostJobsRequestModule)
	require.Emptyf(t, err, "Error creating job: %s", err)
	response, err := client.SearchService.GetJob(*job.Sid)
	assert.NoError(t, err)
	require.NotEmpty(t, response)
	//assert.Equal(t, *job.Sid, response.Sid)
	assert.NotEmpty(t, response.Status)
	assert.Equal(t, PostJobsRequestModule.Query, response.Query)
}

func TestUpdateJobToBeCanceled(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	job, err := client.SearchService.CreateJob(PostJobsRequest)
	require.Emptyf(t, err, "Error creating job: %s", err)
	assert.Equal(t, *(job.Status), search.SearchStatusRunning)
	require.Emptyf(t, err, "Error creating job: %s", err)
	_, err = client.SearchService.UpdateJob(*job.Sid, search.UpdateJob{Status: search.UpdateJobStatusCanceled})
	assert.NoError(t, err)
	job, err = client.SearchService.GetJob(*job.Sid)
	// status should be canceled ??? but now we always get status failed from search service
	//assert.Equal(t,*(job.Status),search.SEARCH_JOB_STATUS_CANCELED)
	assert.Equal(t, *(job.Status), search.SearchStatusFailed)
}

func TestUpdateJobToBeFinalized(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	job, err := client.SearchService.CreateJob(PostJobsRequest)
	require.NoError(t, err)
	_, err = client.SearchService.UpdateJob(*job.Sid, search.UpdateJob{Status: search.UpdateJobStatusFinalized})
	assert.NoError(t, err)
}

func TestGetJobResultsNextLink(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	job, err := client.SearchService.CreateJob(PostJobsRequest)
	require.NoError(t, err)
	query := search.ListResultsQueryParams{}.SetCount(0).SetOffset(0)
	response, err := client.SearchService.ListResults(*job.Sid, &query)
	require.NoError(t, err)
	assert.NotEmpty(t, response)
	assert.NotEmpty(t, response.NextLink)
}

func TestGetJobResults(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	job := createSearchJob(client, PostJobsRequest, t)

	query := search.ListResultsQueryParams{}.SetCount(5).SetOffset(0)
	response, err := client.SearchService.ListResults(*job.Sid, &query)

	assert.NoError(t, err)
	require.NotEmpty(t, response)
	assert.Equal(t, 5, len((*response).Results))
}

//TestIntegrationNewSearchJobBadRequest asynchronously
func TestIntegrationNewSearchJobBadRequest(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)
	response, err := client.SearchService.CreateJob(PostJobsBadRequest)
	require.NotNil(t, err)
	assert.Empty(t, response)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 400, httpErr.HTTPStatusCode)
	assert.Equal(t, "400 Bad Request", httpErr.HTTPStatus)
}

//TestIntegrationGetJobResultsBadSearchID
func TestIntegrationGetJobResultsBadSearchID(t *testing.T) {
	// id format: [\da-z]{32}_\d{19}_[\da-z]{5}
	nonexistentSearchID := "nonexistent000000000000000000000_0000000000000000000_000id"

	client := getClient(t)
	require.NotNil(t, client)

	query := search.ListResultsQueryParams{}.SetCount(0).SetOffset(0)

	resp, err := client.SearchService.ListResults(nonexistentSearchID, &query)
	require.Error(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
	assert.Equal(t, "404 Not Found", httpErr.HTTPStatus)
	assert.Equal(t, "Failed to get job status: unknown sid", httpErr.Message)
	assert.Nil(t, resp)
}

func TestListEventsSummary(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)

	boolvar := true
	postJobsRequest := search.SearchJob{Query: DefaultSearchQuery, CollectEventSummary: &boolvar}

	job := createSearchJob(client, postJobsRequest, t)

	query := search.ListEventsSummaryQueryParams{}.SetCount(3).SetOffset(0).SetField("host")
	response, err := client.SearchService.ListEventsSummary(*job.Sid, &query)

	assert.NoError(t, err)
	require.NotEmpty(t, response)
	assert.Equal(t, 3, len((*response).Results))
}

func TestListFieldsSummary(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)

	boolvar := true
	postJobsRequest := search.SearchJob{Query: DefaultSearchQuery, CollectFieldSummary: &boolvar}

	job := createSearchJob(client, postJobsRequest, t)

	query := search.ListFieldsSummaryQueryParams{}.SetEarliest("-1s")
	response, err := client.SearchService.ListFieldsSummary(*job.Sid, &query)

	assert.NoError(t, err)
	require.NotEmpty(t, response)
	assert.True(t, *(*response).EventCount == 5)
	assert.NotEmpty(t, (*response).Fields)
	assert.True(t, len((*response).Fields) > 0)
}

func TestListTimeBuckets(t *testing.T) {
	client := getClient(t)
	require.NotNil(t, client)

	boolvar := true
	postJobsRequest := search.SearchJob{Query: DefaultSearchQuery, CollectTimeBuckets: &boolvar}
	job := createSearchJob(client, postJobsRequest, t)

	response, err := client.SearchService.ListTimeBuckets(*job.Sid)

	assert.NoError(t, err)
	require.NotEmpty(t, response)
	assert.True(t, *(*response).EventCount == 5)
}

//TestCreateJobConfigurableBackOffRetry and validate that all the job requests are created successfully after retries
func TestCreateJobConfigurableBackOffRetry(t *testing.T) {
	config := &services.Config{
		Token:         testutils.TestAuthenticationToken,
		Host:          testutils.TestSplunkCloudHost,
		Tenant:        testutils.TestTenant,
		Timeout:       testutils.LongTestTimeout,
		RetryRequests: true,
		RetryConfig:   services.RetryStrategyConfig{ConfigurableRetryConfig: &services.ConfigurableRetryConfig{RetryNum: 5, Interval: 800}},
	}
	client, err := services.NewClient(config)
	assert.NoError(t, err)
	searchService := search.NewService(client)

	concurrentSearches := 20

	var wg sync.WaitGroup
	wg.Add(concurrentSearches)

	jobIDs := make(chan string, concurrentSearches)
	errors := make(chan error, concurrentSearches)

	for i := 0; i < concurrentSearches; i++ {
		go func(service *search.Service) {
			defer wg.Done()
			job, err := service.CreateJob(PostJobsRequest)

			if err != nil {
				errors <- err
			} else {
				jobIDs <- *job.Sid
			}
		}(searchService)
	}

	// block on all jobs being created
	wg.Wait()
	close(jobIDs)
	close(errors)

	for err := range errors {
		assert.NoError(t, err)
	}

	cnt := 0
	for id := range jobIDs {
		assert.NotEmpty(t, id)
		cnt++
	}
	assert.Equal(t, concurrentSearches, cnt)
}

//TestCreateJobDefaultBackOffRetry and validate that all the job requests are created successfully after retries
func TestCreateJobDefaultBackOffRetry(t *testing.T) {
	config := &services.Config{
		Token:         testutils.TestAuthenticationToken,
		Host:          testutils.TestSplunkCloudHost,
		Tenant:        testutils.TestTenant,
		Timeout:       testutils.LongTestTimeout,
		RetryRequests: true,
		RetryConfig:   services.RetryStrategyConfig{DefaultRetryConfig: &services.DefaultRetryConfig{}},
	}

	client, err := services.NewClient(config)
	assert.NoError(t, err)
	searchService := search.NewService(client)

	concurrentSearches := 15

	var wg sync.WaitGroup
	wg.Add(concurrentSearches)

	jobIDs := make(chan string, concurrentSearches)
	errors := make(chan error, concurrentSearches)

	for i := 0; i < concurrentSearches; i++ {
		go func(service *search.Service) {
			defer wg.Done()
			job, err := service.CreateJob(PostJobsRequest)

			if err != nil {
				errors <- err
			} else {
				jobIDs <- *job.Sid
			}
		}(searchService)
	}

	// block on all jobs being created
	wg.Wait()
	close(jobIDs)
	close(errors)

	for err := range errors {
		assert.NoError(t, err)
	}

	cnt := 0
	for id := range jobIDs {
		assert.NotEmpty(t, id)
		cnt++
	}
	assert.Equal(t, concurrentSearches, cnt)
}
