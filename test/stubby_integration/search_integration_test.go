package stubbyintegration

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateJob(t *testing.T) {

	response, err := getClient(t).SearchService.CreateJob(&model.PostJobsRequest{Search: "search index=*"})
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, "SEARCH_ID", response)

}

func TestGetResults(t *testing.T) {

	response, err := getClient(t).SearchService.GetJobResults("SEARCH_ID", nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.NotEmpty(t, response.Results)
	assert.NotEmpty(t, response.Results[0].Index)
	assert.NotEmpty(t, response.Fields)
	validateResponse(response, t)
}

//Validate response results
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
