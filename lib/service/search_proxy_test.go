package service

import (
	"encoding/json"
	"github.com/splunk/ssc-client-go/lib/model"
	"testing"
	"time"
)

const (
	ClientID     = "admin"
	ClientSecret = "changeme"
	Host   		 = "http://ssc-sdk-shared-stubby:8882"
)

func TestNewSearchJobWithStubby(t *testing.T) {
	client := NewClient([2]string{ClientID, ClientSecret}, Host, time.Second*5, true)

	response, _ := client.SearchService.CreateJob(&model.PostJobsRequest{Query: "search index=*"})

	data := make(map[string]string)
	err := json.Unmarshal([]byte(response), &data)

	if err != nil && data["searchId"] == "SEARCH_ID" {
		t.Errorf("Expected SEARCHID not found in response")
	}
}

func TestGetJobResultsWithStubby(t *testing.T) {
	client := NewClient([2]string{ClientID, ClientSecret}, Host, time.Second*5, true)

	response, err := client.SearchService.GetResults("SEARCH_ID")

	if err == nil {
		ValidateResponse(response, t)
	} else {
		t.Errorf("Encountered error in Get Results %v", err)
	}
}

func TestNewSearchJobSyncWithStubby(t *testing.T) {
	client := NewClient([2]string{ClientID, ClientSecret}, Host , time.Second*5, true)

	response, err := client.SearchService.CreateSyncJob(&model.PostJobsRequest{Query: "search index=*"})

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
