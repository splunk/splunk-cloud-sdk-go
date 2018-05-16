package service

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
)

const searchServicePrefix = "search"

// SearchService talks to the SSC search service
type SearchService service

//CreateJob Dispatches a search and returns the the newly created search job.
func (service *SearchService) CreateJob(job *model.PostJobsRequest) (*model.PostJobResponse, error) {
	var postJobResponse model.PostJobResponse
	jobURL, err := service.client.BuildURL(searchServicePrefix, "jobs")
	if err != nil {
		return nil, err
	}
	response, err := service.client.Post(jobURL, job)
	util.ParseResponse(&postJobResponse, response, err)
	return &postJobResponse, err
}

//CreateSyncJob Dispatches a new search and return results synchronously
func (service *SearchService) CreateSyncJob(job *model.PostJobsRequest) (*model.SearchEvents, error) {
	var searchModel model.SearchEvents
	jobURL, err := service.client.BuildURL(searchServicePrefix, "jobs", "sync")
	if err != nil {
		return nil, err
	}
	response, err := service.client.Post(jobURL, job)
	util.ParseResponse(&searchModel, response, err)
	return &searchModel, err
}

//GetResults Returns the job resource with the given `id`.
func (service *SearchService) GetResults(jobID string) (*model.SearchEvents, error) {
	var searchModel model.SearchEvents
	jobURL, err := service.client.BuildURL(searchServicePrefix, "jobs", jobID, "results")
	if err != nil {
		return nil, err
	}
	response, err := service.client.Get(jobURL)
	util.ParseResponse(&searchModel, response, err)
	return &searchModel, err
}
