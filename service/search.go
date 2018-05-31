package service

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
	"fmt"
	"io/ioutil"
	"strconv"
	"errors"
)

const searchServicePrefix = "search-service"
const searchServiceVersion = "v1"

// SearchService talks to the SSC search service
type SearchService service

// GetJobsHandler gets details of all current searches.
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
		return jobs, err
	}
	err = util.ParseResponse(&jobs, response, err)
	return jobs, err
}

//CreateJob Dispatches a search and returns the the newly created search job.
func (service *SearchService) CreateJob(job *model.NewSearchConfig) (string, error) {
	// var postJobResponse model.PostJobResponse
	jobURL, err := service.client.BuildURL(nil, searchServicePrefix, searchServiceVersion, "jobs")
	if err != nil {
		return "", err
	}
	response, err := service.client.Post(jobURL, job)
	if response == nil {
		return "", fmt.Errorf("nil response provided")
	} else {
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
		return "", errors.New("Unable to parse jobID")
	}
	return jobID, err
}

func (service *SearchService) GetJob(jobID string) (*model.JobResponse, error) {
	var jobsResponse model.JobResponse
	jobURL, err := service.client.BuildURL(nil, searchServicePrefix, searchServiceVersion, "jobs", jobID)
	response, err := service.client.Get(jobURL)
	if err != nil {
		return nil, err
	}
	err = util.ParseResponse(&jobsResponse, response, err)
	return &jobsResponse, err
}

func (service *SearchService) PostJobControl(jobId string, action *model.JobControlAction) (*model.JobControlReplyMsg, error) {
	var msg model.JobControlReplyMsg
	jobURL, err := service.client.BuildURL(nil, searchServicePrefix, searchServiceVersion, "jobs", jobId, "control")
	if err != nil {
		return nil, err
	}
	response, err := service.client.Post(jobURL, action)
	err = util.ParseResponse(&msg, response, err)
	return &msg, err
}

//GetJobResults Returns the job results with the given `id`.
func (service *SearchService) GetJobResults(jobID string, params *model.FetchResultsRequest) (*model.SearchResults, error) {
	var results model.SearchResults
	jobURL, err := service.client.BuildURL(util.ParseURLParams(*params), searchServicePrefix, searchServiceVersion, "jobs", jobID, "results")
	if err != nil {
		return nil, err
	}
	response, err := service.client.Get(jobURL)
	err = util.ParseResponse(&results, response, err)
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
	err = util.ParseResponse(&results, response, err)
	return &results, err
}
