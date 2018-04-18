package main

import (
	"fmt"
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/service"
	"time"
)

// Canned configs for running this example
const (
	Token   = "4zRqusbLAq754mX5WCDfoiQFzFJFWWkO"
	Host    = "https://api.splunknovadev-playground.com"
	Timeout = time.Second * 5
)

var splunkClient = service.NewClient(Token, Host, Timeout, true)

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

func createJob() *model.PostJobResponse {
	postJobResponse, _ := splunkClient.SearchService.CreateJob(&model.PostJobsRequest{Query: "search index=*"})
	return postJobResponse
}

func createSyncJob() *model.SearchEvents {
	searchEvent, _ := splunkClient.SearchService.CreateSyncJob(&model.PostJobsRequest{Query: "search index=*"})
	return searchEvent
}

func getResults(searchID string) *model.SearchEvents {
	searchEvent, _ := splunkClient.SearchService.GetResults(searchID)
	return searchEvent
}

func main() {
	///////////////////////////////
	// 1a) create a new search job
	///////////////////////////////
	postJobResponse := createJob()
	fmt.Printf("%+v\n", &postJobResponse)

	///////////////////////////////////
	// 1b) retrieve a job results by id
	///////////////////////////////////
	searchModel1 := getResults("1b9c6b21-3277-4dc6-b80a-76894678986f")
	printSearchModel(searchModel1)

	/////////////////////////////////////
	// 2) create a new synchronous search
	/////////////////////////////////////
	searchModel2 := createSyncJob()
	printSearchModel(searchModel2)
}
