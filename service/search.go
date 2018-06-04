package service

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
	"io/ioutil"
	"strconv"
	"errors"
)

const searchServicePrefix = "search"
const searchServiceVersion = "v1"

// SearchService talks to the SSC search service
type SearchService service

// GetJobs gets details of all current searches.
func (service *SearchService) GetJobs(params *model.JobsRequest) ([]model.SearchJob, error) {
	var jobs []model.SearchJob
	if params == nil {
		params = model.NewDefaultPaginationParams()
	}
	jobsURL, err := service.client.BuildURL(util.ParseURLParams(*params), searchServicePrefix, searchServiceVersion, "jobs")
	if err != nil {
		return jobs, err
	}
	response, err := service.client.Get(jobsURL)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	err = util.ParseResponse(&jobs, response)
	return jobs, err
}

// CreateJob dispatches a search and returns sid.
func (service *SearchService) CreateJob(job *model.PostJobsRequest) (string, error) {
	// var postJobResponse model.PostJobResponse
	jobURL, err := service.client.BuildURL(nil, searchServicePrefix, searchServiceVersion, "jobs")
	if err != nil {
		return "", err
	}
	response, err := service.client.Post(jobURL, job)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	jobID, err := strconv.Unquote(string(body))
	if err != nil {
		return "", errors.New("unable to parse jobID")
	}
	return jobID, err
}

// GetJob retrieves information about the specified search.
func (service *SearchService) GetJob(jobID string) (*model.SearchJobContent, error) {
	var jobsResponse model.SearchJobContent
	jobURL, err := service.client.BuildURL(nil, searchServicePrefix, searchServiceVersion, "jobs", jobID)
	response, err := service.client.Get(jobURL)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	err = util.ParseResponse(&jobsResponse, response)
	return &jobsResponse, err
}

// PostJobControl runs a job control command for the specified search.
func (service *SearchService) PostJobControl(jobID string, action *model.JobControlAction) (*model.JobControlReplyMsg, error) {
	var msg model.JobControlReplyMsg
	jobURL, err := service.client.BuildURL(nil, searchServicePrefix, searchServiceVersion, "jobs", jobID, "control")
	if err != nil {
		return nil, err
	}
	response, err := service.client.Post(jobURL, action)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	err = util.ParseResponse(&msg, response)
	return &msg, err
}

// GetJobResults Returns the job results with the given `id`.
func (service *SearchService) GetJobResults(jobID string, params *model.FetchResultsRequest) (*model.SearchResults, error) {
	var results model.SearchResults
	jobURL, err := service.client.BuildURL(util.ParseURLParams(*params), searchServicePrefix, searchServiceVersion, "jobs", jobID, "results")
	if err != nil {
		return nil, err
	}
	response, err := service.client.Get(jobURL)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	err = util.ParseResponse(&results, response)
	return &results, err
}

//GetJobEvents Returns the job events with the given `id`.
func (service *SearchService) GetJobEvents(jobID string, params *model.FetchEventsRequest) (*model.SearchResults, error) {
	var results model.SearchResults
	jobURL, err := service.client.BuildURL(util.ParseURLParams(*params), searchServicePrefix, searchServiceVersion, "jobs", jobID, "events")
	if err != nil {
		return nil, err
	}
	response, err := service.client.Get(jobURL)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	err = util.ParseResponse(&results, response)
	return &results, err
}
