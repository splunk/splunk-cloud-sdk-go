package service

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
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
	if err != nil {
		return jobs, err
	}
	err = util.ParseResponse(&jobs, response, err)
	return jobs, err
}

//CreateJob Dispatches a search and returns the the newly created search job.
func (service *SearchService) CreateJob(job *model.PostJobsRequest) (*model.PostJobResponse, error) {
	var postJobResponse model.PostJobResponse
	jobURL, err := service.client.BuildURL(nil, searchServicePrefix, searchServiceVersion, "jobs")
	if err != nil {
		return nil, err
	}
	response, err := service.client.Post(jobURL, job)
	err = util.ParseResponse(&postJobResponse, response, err)
	return &postJobResponse, err
}

//GetResults Returns the job resource with the given `id`.
func (service *SearchService) GetResults(jobID string) (*model.SearchEvents, error) {
	var searchModel model.SearchEvents
	jobURL, err := service.client.BuildURL(nil, searchServicePrefix, searchServiceVersion, "jobs", jobID, "results")
	if err != nil {
		return nil, err
	}
	response, err := service.client.Get(jobURL)
	err = util.ParseResponse(&searchModel, response, err)
	return &searchModel, err
}
