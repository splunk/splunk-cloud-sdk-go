// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package service

import (
	"fmt"
	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"net/url"
	"time"
)

const searchServicePrefix = "search"
const searchServiceVersion = "v1beta1"

// SearchService talks to the Splunk Cloud search service
type SearchService service

// ListJobs gets the matching list of search jobs
func (service *SearchService) ListJobs() ([]model.SearchJob, error) {
	var searchJobs []model.SearchJob
	jobsURL, err := service.client.BuildURL(nil, searchServicePrefix, searchServiceVersion, "jobs")
	if err != nil {
		return searchJobs, err
	}
	response, err := service.client.Get(RequestParams{URL: jobsURL})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	err = util.ParseResponse(&searchJobs, response)
	return searchJobs, err
}

// CreateJob creates a new search job
func (service *SearchService) CreateJob(job *model.CreateJobRequest) (*model.SearchJob, error) {
	var postJobResponse model.SearchJob
	jobURL, err := service.client.BuildURL(nil, searchServicePrefix, searchServiceVersion, "jobs")
	if err != nil {
		return &postJobResponse, err
	}
	response, err := service.client.Post(RequestParams{URL: jobURL, Body: job})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return &postJobResponse, err
	}
	err = util.ParseResponse(&postJobResponse, response)
	return &postJobResponse, err
}

// GetJob retrieves information about the specified search.
func (service *SearchService) GetJob(jobID string) (*model.SearchJob, error) {
	var searchJob model.SearchJob
	jobURL, err := service.client.BuildURL(nil, searchServicePrefix, searchServiceVersion, "jobs", jobID)
	response, err := service.client.Get(RequestParams{URL: jobURL})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	err = util.ParseResponse(&searchJob, response)
	return &searchJob, err
}

// UpdateJob updates an existing job with actions and TTL
func (service *SearchService) UpdateJob(jobID string, action model.JobStatus) (*model.PatchJobResponse, error) {
	var patchResponse model.PatchJobResponse
	jobURL, err := service.client.BuildURL(nil, searchServicePrefix, searchServiceVersion, "jobs", jobID)
	if err != nil {
		return nil, err
	}
	requestBody := struct {
		Action model.JobStatus `json:"action"`
	}{action}
	response, err := service.client.Patch(RequestParams{URL: jobURL, Body: requestBody})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	err = util.ParseResponse(&patchResponse, response)
	return &patchResponse, err
}

// GetResults Returns the job results with the given `id`. count=0 returns default number of results from search
func (service *SearchService) GetResults(jobID string, count, offset int) (interface{}, error) {
	query := url.Values{}
	if count > 0 {
		query.Set("count", fmt.Sprintf("%d", count))
	}
	query.Set("offset", fmt.Sprintf("%d", offset))
	jobURL, err := service.client.BuildURL(query, searchServicePrefix, searchServiceVersion, "jobs", jobID, "results")
	if err != nil {
		return nil, err
	}
	response, err := service.client.Get(RequestParams{URL: jobURL})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	// Create a temporary struct to check if nextLink field exists in payload
	var tempNextLinkModel struct {
		NextLink *string `json:"nextLink,omitempty"`
	}
	err = util.ParseResponse(&tempNextLinkModel, response)
	if err != nil {
		return nil, err
	}
	// NextLink exists
	if tempNextLinkModel.NextLink != nil {
		var emptyResponse model.ResultsNotReadyResponse
		err = util.ParseResponse(&emptyResponse, response)
		return &emptyResponse, err
	}
	// NextLink does not exist
	var results model.SearchResults
	err = util.ParseResponse(&results, response)
	return &results, err
}

// WaitForJob polls the job until it's completed or errors out
func (service *SearchService) WaitForJob(jobID string, pollInterval time.Duration) (interface{}, error) {
	for {
		job, err := service.GetJob(jobID)
		if err != nil {
			return nil, err
		}
		// wait for terminal state
		switch job.Status {
		case model.Done, model.Failed:
			return job.Status, nil
		default:
			time.Sleep(pollInterval)
		}
	}
}
