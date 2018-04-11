package service

import (
	"encoding/json"
	"github.com/splunk/ssc-client-go/lib/model"
	"github.com/splunk/ssc-client-go/lib/util"
	"io/ioutil"
)

// SearchService implements a new service type
type SearchService service

//CreateJob Dispatches a search and returns the the newly created search job.
func (service *SearchService) CreateJob(job *model.PostJobsRequest) (string, error) {
	jobURL := service.client.BuildURL(nil, "search", "v1", "jobs")
	response, err := service.client.Post(jobURL, job)
	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return "", err
	}
	data := make(map[string]string)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}
	return string(body), err
}

//CreateSyncJob Dispatches a new search and return results synchronously
func (service *SearchService) CreateSyncJob(job *model.PostJobsRequest) (*model.SearchEvents, error) {
	var searchModel model.SearchEvents
	jobURL := service.client.BuildURL(nil, "search", "v1", "jobs", "sync")
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
	jobURL := service.client.BuildURL(nil, "search", "v1", "jobs", jobID, "results")
	response, err := service.client.Get(jobURL)
	if err != nil {
		return nil, err
	}
	util.ParseResponse(&searchModel, response, err)
	return &searchModel, err
}
