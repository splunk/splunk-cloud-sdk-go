// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

const servicePrefix = "search"
const serviceVersion = "v1beta1"
const serviceCluster = "api"

// Service talks to the Splunk Cloud search service
type Service services.BaseService

// NewService creates a new search service client from the given Config
func NewService(config *services.Config) (*Service, error) {
	baseClient, err := services.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Service{Client: baseClient}, nil
}

// JobsQuery represents Query Parameters that can be provided for ListJobs endpoint
type JobsQuery struct {
	//The supported statuses are running, done and failed
	Status string `key:"status"`
}

// ListJobs gets the matching list of search jobs
func (s *Service) ListJobs() ([]Job, error) {
	return s.ListJobsByQueryParameters(JobsQuery{})
}

// ListJobsByQueryParameters gets the matching list of search jobs filtered by query parameters specified
func (s *Service) ListJobsByQueryParameters(query JobsQuery) ([]Job, error) {
	var searchJobs []Job
	values := util.ParseURLParams(query)
	jobsURL, err := s.Client.BuildURL(values, serviceCluster, servicePrefix, serviceVersion, "jobs")
	if err != nil {
		return searchJobs, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: jobsURL})
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
func (s *Service) CreateJob(job *CreateJobRequest) (*Job, error) {
	var postJobResponse Job
	jobURL, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "jobs")
	if err != nil {
		return &postJobResponse, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: jobURL, Body: job})
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
func (s *Service) GetJob(jobID string) (*Job, error) {
	var searchJob Job
	jobURL, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "jobs", jobID)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: jobURL})
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
func (s *Service) UpdateJob(jobID string, jobStatus JobAction) (*PatchJobResponse, error) {
	var patchResponse PatchJobResponse
	jobURL, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "jobs", jobID)
	if err != nil {
		return nil, err
	}
	requestBody := struct {
		Status JobAction `json:"status"`
	}{jobStatus}
	response, err := s.Client.Patch(services.RequestParams{URL: jobURL, Body: requestBody})
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
func (s *Service) GetResults(jobID string, count, offset int) (interface{}, error) {
	query := url.Values{}
	if count > 0 {
		query.Set("count", fmt.Sprintf("%d", count))
	}
	query.Set("offset", fmt.Sprintf("%d", offset))
	jobURL, err := s.Client.BuildURL(query, serviceCluster, servicePrefix, serviceVersion, "jobs", jobID, "results")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: jobURL})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	// assign response.Body to a variable so that we can reuse response.Body later
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	// Create a temporary struct to check if nextLink field exists in payload
	var tempNextLinkModel struct {
		NextLink *string `json:"nextLink"`
		Wait     string  `json:"wait"`
	}
	err = json.Unmarshal(bodyBytes, &tempNextLinkModel)
	if err != nil {
		return nil, err
	}
	// NextLink exists
	if tempNextLinkModel.NextLink != nil {
		var emptyResponse ResultsNotReadyResponse
		// reset the response body to the original state
		response.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		err = util.ParseResponse(&emptyResponse, response)
		return &emptyResponse, err
	}
	// NextLink does not exist
	var results Results
	// reset the response body to the original state
	response.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	err = util.ParseResponse(&results, response)
	return &results, err
}

// WaitForJob polls the job until it's completed or errors out
func (s *Service) WaitForJob(jobID string, pollInterval time.Duration) (interface{}, error) {
	for {
		job, err := s.GetJob(jobID)
		if err != nil {
			return nil, err
		}
		// wait for terminal state
		switch job.Status {
		case Done, Failed:
			return job.Status, nil
		default:
			time.Sleep(pollInterval)
		}
	}
}
