package stubbyintegration

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateJob(t *testing.T) {
	client := getClient(t)
	response, err := client.SearchService.CreateJob(&model.PostJobsRequest{Search: "search index=*"})
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, "SEARCH_ID", response)

}

func TestGetJobs(t *testing.T) {
	response, err := getClient(t).SearchService.GetJobs(nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, "SEARCH_ID", response[0].Sid)
}

func TestGetJob(t *testing.T) {
	response, err := getClient(t).SearchService.GetJob("SEARCH_ID")
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, "SEARCH_ID", response.Sid)
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

func TestGetEvents(t *testing.T) {
	response, err := getClient(t).SearchService.GetJobEvents("SEARCH_ID", nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.NotEmpty(t, response.Results)
	assert.NotEmpty(t, response.Results[0].Index)
	assert.NotEmpty(t, response.Fields)
	validateResponse(response, t)
}

func TestJobActionControl(t *testing.T) {
	response, err := getClient(t).SearchService.PostJobControl("SEARCH_ID", &model.JobControlAction{Action:"pause"})
	assert.Nil(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, 1, len(response.Msg))
	assert.Equal(t, "INFO", response.Msg[0].TypeKey)
	assert.Equal(t, "Search job paused.", response.Msg[0].Text)
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
