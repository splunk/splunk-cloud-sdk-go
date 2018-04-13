package service

import (
	"testing"
	"github.com/splunk/ssc-client-go/lib/model"
	"github.com/stretchr/testify/assert"
	"time"
)

func getSplunkClient() *Client {

	//return NewClient([2]string{"admin", "changeme"},
	//	"http://localhost:8882", time.Second*5, true)
	return NewClient([2]string{"admin", "changeme"},
		"http://ssc-sdk-shared-stubby:8882", time.Second*5, true)
}


func TestCreateJob(t *testing.T) {

	response, err := getSplunkClient().SearchService.CreateJob(&model.PostJobsRequest{Query: "search index=*"})
	assert.Empty(t, err)
	assert.NotEmpty(t, response)

}

func TestCreateSyncJob(t *testing.T) {

	response, err := getSplunkClient().SearchService.CreateSyncJob(&model.PostJobsRequest{Query: "search index=*"})
	assert.Empty(t, err)
	assert.NotEmpty(t, response)
	assert.NotEmpty(t, response.Results)
	assert.NotEmpty(t, response.Results[0].Host)
	assert.NotEmpty(t, response.Fields)
}

func TestGetResults(t *testing.T) {

	response, err := getSplunkClient().SearchService.GetResults("SEARCH_ID")
	assert.Empty(t, err)
	assert.NotEmpty(t, response)
	assert.NotEmpty(t, response.Results)
	assert.NotEmpty(t, response.Results[0].Index)
	assert.NotEmpty(t, response.Fields)
}