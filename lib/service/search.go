package service

import (
	"github.com/splunk/ssc-client-go/lib/model"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

// SearchService implements a new service type
type SearchService service

//
// POST /search/v1/jobs
// NewSearch dispatches a new spl search and returns sid
//
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


//
// POST /search/v1/jobs/sync
// NewSyncSearch (i.e. one-shot) dispatches a new spl search and returns sid
//
func (service *SearchService) CreateSyncJob(spl string) (*model.SearchEvents, error) {
	var searchModel model.SearchEvents
	// TODO(dan): Services are down, using fake data for now.
	//jobURL := service.client.BuildSplunkdURL(nil, "search", "v1", "jobs", "sync")
	//response, err := service.client.Post(jobURL, map[string]string{"query": spl})
	//body, err := ioutil.ReadAll(response.Body)
	err := json.Unmarshal([]byte(model.GetFakeJson1()), &searchModel)
	return &searchModel, err
}


//
// FIXME(dan): GET /search/v1/jobs{jobId}
// Retrieves a job for a jobId
//
func (service *SearchService) GetSearch(jobId string) (string, error) {
	jobURL := service.client.BuildSplunkdURL(nil, "search", "v1", "jobs", jobId)
	response, err := service.client.Get(jobURL)
	//body, err := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", response)
	return "", err
}


//
// FIXME(dan): DELETE /search/v1/jobs{jobId}
// Retrieves a job for a jobId
//
func (service *SearchService) DeleteSearch(jobId string) (string, error) {
	jobURL := service.client.BuildSplunkdURL(nil, "search", "v1", "jobs", jobId)
	response, err := service.client.Delete(jobURL)
	//body, err := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", response)
	return "", err
}