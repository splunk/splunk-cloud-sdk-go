// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package service

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
	"time"
	"net/url"
)

const searchServicePrefix = "search"
const searchServiceVersion = "v2"

// SearchService talks to the SSC search service
type SearchService service

// GetJobs gets the matching list of search jobs
func (service *SearchService) ListJobs() (*[]model.SearchJob, error) {
	var searchJobs []model.SearchJob
	jobsURL, err := service.client.BuildURL(nil, searchServicePrefix, searchServiceVersion, "jobs")
	if err != nil {
		return &searchJobs, err
	}
	response, err := service.client.Get(RequestParams{URL: jobsURL})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	err = util.ParseResponse(&searchJobs, response)
	return &searchJobs, err
}

// CreateJob creates a new search job
func (service *SearchService) CreateJob(job *model.SearchJobBase) (*model.SearchJob, error) {
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
func (service *SearchService) GetJob(jobId string) (*model.SearchJob, error) {
	var searchJob model.SearchJob
	jobURL, err := service.client.BuildURL(nil, searchServicePrefix, searchServiceVersion, "jobs", jobId)
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
func (service *SearchService) UpdateJob(jobId string, patch *model.PatchJobRequest) (*model.PatchJobResponse, error){
	var patchResponse model.PatchJobResponse
	jobURL, err := service.client.BuildURL(nil, searchServicePrefix, searchServiceVersion, "jobs", jobId)
	if err != nil {
		return nil, err
	}
	response, err := service.client.Patch(RequestParams{URL: jobURL, Body: patch})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	err = util.ParseResponse(&patchResponse, response)
	return &patchResponse, err
}

// GetJobResults Returns the job results with the given `id`.
func (service *SearchService) GetJobResults(jobID string, params *model.JobResultsParams) (interface{}, error) {
	var query url.Values
	if params != nil {
		query = util.ParseURLParams(*params)
	}
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
	// check if there's any search result available
	jobStatus, err := service.GetJob(jobID)
	if err != nil {
		return nil, err
	}

	// return EmptyResultsResponse model if no results is available
	if jobStatus.ResultsAvailable == 0 {
		var emptyResponse model.EmptyResultsResponse
		err = util.ParseResponse(&emptyResponse, response)
		return &emptyResponse, err
	} else {
		// return available results
		var jobResponse model.SearchResults
		err = util.ParseResponse(&jobResponse, response)
		return &jobResponse, err
	}
}

// WaitForJob polls the job until it's completed or errors out
func (service *SearchService) WaitForJob(sid string, pollInterval time.Duration) (interface{}, error) {
	for {
		job, err := service.GetJob(sid)
		if err != nil {
			return nil,err
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
