// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package stubbyintegration

import (
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateJob(t *testing.T) {
	client := getClient(t)
	response, err := client.SearchService.CreateJob(&model.CreateJobRequest{Query: "search index=*"})
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, "SEARCH_ID", response)

}

func TestGetJobs(t *testing.T) {
	response, err := getClient(t).SearchService.ListJobs()
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, "SEARCH_ID", response[0].ID)
}

func TestGetJob(t *testing.T) {
	response, err := getClient(t).SearchService.GetJob("SEARCH_ID")
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, "SEARCH_ID", response.ID)
}

func TestGetResults(t *testing.T) {
	response, err := getClient(t).SearchService.GetResults("SEARCH_ID", 0, 0)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.NotEmpty(t, response.(*model.SearchResults).Results)
	assert.NotEmpty(t, response.(*model.SearchResults).Results[0])
	assert.NotEmpty(t, response.(*model.SearchResults).Fields)
	validateResponse(response.(*model.SearchResults), t)
}

// Validate response results
func validateResponse(response *model.SearchResults, t *testing.T) {

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
