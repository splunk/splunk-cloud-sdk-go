package service

import (
	"encoding/json"
	"fmt"
	"github.com/splunk/ssc-client-go/lib/model"
	"io/ioutil"
)

// SearchService implements a new service type
type SearchService service

// CreateJob dispatch a search and return the jobID
// POST /search/v1/jobs
func (service *SearchService) CreateJob(job *model.PostJobsRequest) (string,error) {
	jobURL := service.client.BuildSplunkdURL(nil, "search", "v1", "jobs")
	response, err := service.client.Post(jobURL, job)
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
	response, err := service.client.Post(jobURL, job)
	body, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &searchModel)
	return &searchModel, err
}

// GetJob Returns the job resource with the given `jobID`.
// GET /search/v1/jobs/{jobID}
func (service *SearchService) GetJob(jobID string) (string, error) {
	jobURL := service.client.BuildSplunkdURL(nil, "search", "v1", "jobs", jobID)
	response, err := service.client.Get(jobURL)
	//body, err := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", response)
	return "Not Implemented", err
}

// DeleteJob Delete the search job with the given `jobID`, cancelling the search if it is running.
// DELETE /search/v1/jobs/{jobID}
func (service *SearchService) DeleteJob(jobID string) (string, error) {
	jobURL := service.client.BuildSplunkdURL(nil, "search", "v1", "jobs", jobID)
	response, err := service.client.Delete(jobURL)
	//body, err := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", response)
	return "Not Implemented", err
}

// GetResults Returns results for the search job corresponding to "id".
// GET /search/v1/jobs/{jobID}/results
func (service *SearchService) GetResults(jobID string) (*model.SearchEvents, error) {
	var searchModel model.SearchEvents
	jobURL := service.client.BuildSplunkdURL(nil, "search", "v1", "jobs", jobID, "results")
	response, err := service.client.Get(jobURL)
	body, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &searchModel)
	return &searchModel, err
}
