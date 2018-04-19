package service

import (
	"testing"

	"github.com/splunk/ssc-client-go/model"
)

func TestNewSearchJobWithStubby(t *testing.T) {

	response, err := getSplunkClient().SearchService.CreateJob(&model.PostJobsRequest{Query: "search index=*"})
	if err != nil && response.SearchID != "SEARCH_ID" {
		t.Errorf("Expected SEARCHID not found in response")
	}
}

func TestGetJobResultsWithStubby(t *testing.T) {

	response, err := getSplunkClient().SearchService.GetResults("SEARCH_ID")

	if err == nil {
		ValidateResponse(response, t)
	} else {
		t.Errorf("Encountered error in Get Results %v", err)
	}
}

func TestNewSearchJobSyncWithStubby(t *testing.T) {

	response, err := getSplunkClient().SearchService.CreateSyncJob(&model.PostJobsRequest{Query: "search index=*"})

	if err == nil {
		ValidateResponse(response, t)
	} else {
		t.Errorf("Encountered error in Create Sync Job %v", err)
	}
}

//Validate response results
func ValidateResponse(response *model.SearchEvents, t *testing.T) {

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
