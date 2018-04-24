package service

import (
	"testing"
	"time"

	"github.com/splunk/ssc-client-go/model"
	. "github.com/splunk/ssc-client-go/util"
	"github.com/stretchr/testify/assert"
)

func getSplunkClientNew(tenant string, host string) *Client {
	return NewClient(tenant, TestToken, host, time.Second*5)
}

func getSplunkClient() *Client {
	return NewClient(TestTenantID, TestToken, TestStubbySchme+"://"+TestStubbyHost, time.Second*5)
}

func TestCreateJob(t *testing.T) {

	response, err := getSplunkClient().SearchService.CreateJob(&model.PostJobsRequest{Query: "search index=*"})
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, "SEARCH_ID", response.SearchID)

}

func TestCreateSyncJob(t *testing.T) {

	response, err := getSplunkClient().SearchService.CreateSyncJob(&model.PostJobsRequest{Query: "search index=*"})
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.NotEmpty(t, response.Results)
	assert.NotEmpty(t, response.Results[0].Host)
	assert.NotEmpty(t, response.Fields)
	validateResponse(response, t)
}

func TestGetResults(t *testing.T) {

	response, err := getSplunkClient().SearchService.GetResults("SEARCH_ID")
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.NotEmpty(t, response.Results)
	assert.NotEmpty(t, response.Results[0].Index)
	assert.NotEmpty(t, response.Fields)
	validateResponse(response, t)
}

//Validate response results
func validateResponse(response *model.SearchEvents, t *testing.T) {

	var indexFound bool

	if response.Fields != nil {
		for _, v := range response.Fields {
			for m, n := range v {
				if m == "name" && n == "index" {
					indexFound = true
				}
			}
		}
		if !indexFound {
			t.Errorf("Expected results field element name and corresponding value index not found")
		}
	} else {
		t.Errorf("Expected field elements in results not found")
	}

	if response.Results == nil {
		t.Errorf("Invalid response, missing results in response returned")
	}

}
