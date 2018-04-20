package service

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
)

const searchServicePrefix = "search"
const searchServiceVersion = "v1"

// SearchService implements a new service type
type SearchService service

//CreateJob Dispatches a search and returns the the newly created search job.
func (service *SearchService) CreateJob(job *model.PostJobsRequest) (*model.PostJobResponse, error) {
	var postJobResponse model.PostJobResponse
	jobURL := service.client.BuildURLWithTenantID(searchServicePrefix, searchServiceVersion, "jobs")
	response, err := service.client.Post(jobURL, job)
	util.ParseResponse(&postJobResponse, response, err)
	return &postJobResponse, err
}

//CreateSyncJob Dispatches a new search and return results synchronously
func (service *SearchService) CreateSyncJob(job *model.PostJobsRequest) (*model.SearchEvents, error) {
	var searchModel model.SearchEvents
	jobURL := service.client.BuildURLWithTenantID(searchServicePrefix, searchServiceVersion, "jobs", "sync")
	response, err := service.client.Post(jobURL, job)
	if err != nil {
		return nil, err
	}
	util.ParseResponse(&searchModel, response, err)
	return &searchModel, err
}

//GetResults Returns the job resource with the given `id`.
func (service *SearchService) GetResults(jobID string) (*model.SearchEvents, error) {
	var searchModel model.SearchEvents
	jobURL := service.client.BuildURLWithTenantID(searchServicePrefix, searchServiceVersion, "jobs", jobID, "results")
	response, err := service.client.Get(jobURL)
	if err != nil {
		return nil, err
	}
	util.ParseResponse(&searchModel, response, err)
	return &searchModel, err
}
