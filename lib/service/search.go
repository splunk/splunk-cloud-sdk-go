package service

import (
	"encoding/json"
	"github.com/splunk/ssc-client-go/lib/model"
	"github.com/splunk/ssc-client-go/lib/util"
	"io/ioutil"
)

// SearchService implements a new service type
type SearchService service

// CreateJob dispatch a search and return the jobID
// POST /search/v1/jobs
func (service *SearchService) CreateJob(job *model.PostJobsRequest) (string, error) {
	jobURL := service.client.BuildSplunkdURL(nil, "search", "v1", "jobs")
	response, err := service.client.Post(jobURL, job, JSON)
	body, err := ioutil.ReadAll(response.Body)

	//
	// simple parsing for now, data binding later
	//
	data := make(map[string]string)
	json.Unmarshal(body, &data)
	return string(body), err
}

// CreateSyncJob Dispatch a search and return the newly created search job synchronously
// POST /search/v1/jobs/sync
func (service *SearchService) CreateSyncJob(job *model.PostJobsRequest) (*model.SearchEvents, error) {
	var searchModel model.SearchEvents
	jobURL := service.client.BuildSplunkdURL(nil, "search", "v1", "jobs", "sync")
	response, err := service.client.Post(jobURL, job, JSON)
	util.ParseResponse(&searchModel, response, err)
	return &searchModel, err
}

// GetResults Returns results for the search job corresponding to "id".
// GET /search/v1/jobs/{jobID}/results
func (service *SearchService) GetResults(jobID string) (*model.SearchEvents, error) {
	var searchModel model.SearchEvents
	jobURL := service.client.BuildSplunkdURL(nil, "search", "v1", "jobs", jobID, "results")
	response, err := service.client.Get(jobURL, JSON)
	util.ParseResponse(&searchModel, response, err)
	return &searchModel, err
}
