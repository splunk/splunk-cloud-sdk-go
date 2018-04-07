package tests

import (
	"github.com/splunk/ssc-client-go/lib/service"
	"time"
	"testing"
	"github.com/splunk/ssc-client-go/lib/model"
	"fmt"
)


const (
	ClientID     = "admin"
	ClientSecret = "changeme"
	BaseURL      = "localhost:32769"
)

func TestNewSearchJobWithStubby(t *testing.T) {
	client := service.NewClient(
		"", [2]string{ClientID, ClientSecret}, BaseURL, "http", service.NewHTTPClient(time.Second*5, true))

	response, json,_ := client.SearchService.CreateJob(&model.PostJobsRequest{Query:"search index=*"})
	fmt.Println("json: ", json)  // TODO(dan): add a test?

	if response != "SEARCH_ID" {
		t.Errorf("Expected SEARCHID not found in response")
	}
}

func TestGetJobResultsWithStubby(t *testing.T) {
	client := service.NewClient(
		"", [2]string{ClientID, ClientSecret}, BaseURL, "http", service.NewHTTPClient(time.Second*5, true))

	client.SearchService.CreateJob(&model.PostJobsRequest{Query:"search index=*"})
	response, _ := client.SearchService.GetResults("SEARCH_ID")

	ValidateResponse(response, t)
}

func TestNewSearchJobSyncWithStubby(t *testing.T) {
	client := service.NewClient(
		"", [2]string{ClientID, ClientSecret}, BaseURL, "http", service.NewHTTPClient(time.Second*5, true))

	response, _ := client.SearchService.CreateSyncJob(&model.PostJobsRequest{Query:"search index=*"})

	ValidateResponse(response, t)
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
