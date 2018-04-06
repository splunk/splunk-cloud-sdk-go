package service

import (
	"encoding/json"
	"fmt"
	"github.com/splunk/ssc-client-go/lib/model"
	"io/ioutil"
)

// SearchService implements a new service type
type SearchService service

// CreateJob dispatches a new spl search and returns sid
// POST /search/v1/jobs
func (service *SearchService) CreateJob(spl string) (string, error) {
	jobURL := service.client.BuildSplunkdURL(nil, "search", "v1", "jobs")
	response, err := service.client.Post(jobURL, map[string]string{"query": spl})
	body, err := ioutil.ReadAll(response.Body)

	//
	// simple parsing for now, data binding later
	//
	data := make(map[string]string)
	json.Unmarshal(body, &data)
	return data["searchId"], err
}

// CreateSyncJob (i.e. one-shot) dispatches a new spl search and returns sid
// POST /search/v1/jobs/sync
func (service *SearchService) CreateSyncJob(spl string) (*model.SearchEvents, error) {
	var searchModel model.SearchEvents
	jobURL := service.client.BuildSplunkdURL(nil, "search", "v1", "jobs", "sync")
	response, err := service.client.Post(jobURL, map[string]string{"query": spl})
	body, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &searchModel)
	return &searchModel, err
}

// GetJob retrieves a job for a jobID
// GET /search/v1/jobs/{jobID}
func (service *SearchService) GetJob(jobID string) (string, error) {
	jobURL := service.client.BuildSplunkdURL(nil, "search", "v1", "jobs", jobID)
	response, err := service.client.Get(jobURL)
	//body, err := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", response)
	return "Not Implemented", err
}

// DeleteJob a search job by jobID
// DELETE /search/v1/jobs/{jobID}
func (service *SearchService) DeleteJob(jobID string) (string, error) {
	jobURL := service.client.BuildSplunkdURL(nil, "search", "v1", "jobs", jobID)
	response, err := service.client.Delete(jobURL)
	//body, err := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", response)
	return "Not Implemented", err
}

// GetResults returns a SearchEvent for a jobID
// GET /search/v1/jobs/{jobID}/results
func (service *SearchService) GetResults(jobID string) (*model.SearchEvents, error) {
	var searchModel model.SearchEvents
	jobURL := service.client.BuildSplunkdURL(nil, "search", "v1", "jobs", jobID, "results")
	response, err := service.client.Get(jobURL)
	body, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &searchModel)
	return &searchModel, err
}
