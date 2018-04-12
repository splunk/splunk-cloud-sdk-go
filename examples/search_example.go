package main

import (
	"fmt"
	"github.com/splunk/ssc-client-go/lib/model"
	"github.com/splunk/ssc-client-go/lib/service"
	"time"
)

// Canned configs for running this example
const (
	ClientID     = "4zRqusbLAq754mX5WCDfoiQFzFJFWWkO"
	ClientSecret = "ff9odDwxiZqSVEQzcBeOU-_ALDLKksXlELySNdjkbPxRH7rV9gybNhhbgbucteGe"
	URL          = "https://api.splunknovadev-playground.com"
	Timeout      = time.Second * 5
)

var splunkClient = service.NewClient([2]string{ClientID, ClientSecret}, URL, Timeout, true)

func printSearchModel(searchModel *model.SearchEvents) {
	fmt.Println("Preview: ", searchModel.Preview)
	fmt.Println("InitOffset: ", searchModel.InitOffset)
	fmt.Println("Messages: ", searchModel.Messages)
	fmt.Println("Results: ")
	for _, result := range searchModel.Results {
		fmt.Printf("%+v\n", result)
	}
	fmt.Println("Fields: ", searchModel.Fields)
	fmt.Println("Highlighted: ", searchModel.Highlighted)
}

func createJob() string {
	json, _ := splunkClient.SearchService.CreateJob(&model.PostJobsRequest{Query: "search index=*"})
	return json
}

func createSyncJob() *model.SearchEvents {
	searchModel, _ := splunkClient.SearchService.CreateSyncJob(&model.PostJobsRequest{Query: "search index=*"})
	return searchModel
}

func getResults(searchID string) *model.SearchEvents {
	searchModel, _ := splunkClient.SearchService.GetResults(searchID)
	return searchModel
}

func main() {
	///////////////////////////////
	// 1a) create a new search job
	///////////////////////////////
	json := createJob()
	fmt.Println(json)

	///////////////////////////////////
	// 1b) retrieve a job results by id
	///////////////////////////////////
	searchModel1 := getResults("1b9c6b21-3277-4dc6-b80a-76894678986f")
	printSearchModel(searchModel1)

	////////////////////////////////////
	// 2) create a new synchronous search
	////////////////////////////////////åå
	searchModel2 := createSyncJob()
	printSearchModel(searchModel2)
}
