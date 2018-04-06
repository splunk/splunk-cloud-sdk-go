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
	Host         = "api.splunknovadev-playground.com"
	Scheme       = "https"
	Timeout      = time.Second * 5
)

var splunkClient = service.NewClient(
	"", [2]string{ClientID, ClientSecret}, Host, Scheme, service.NewHTTPClient(Timeout, true))

func printSearchModel(searchModel *model.SearchEvents) {
	fmt.Println("Preview: ", searchModel.Preview)
	fmt.Println("InitOffset: ", searchModel.InitOffset)
	fmt.Println("Messages: ", searchModel.Messages)
	fmt.Println("Results: ")
	for _, result := range searchModel.Results {
		fmt.Println("\tBkt", result.Bkt)
		fmt.Println("\tCd", result.Cd)
		fmt.Println("\tIndexTime", result.IndexTime)
		fmt.Println("\tRaw", result.Raw)
		fmt.Println("\tSerial", result.Serial)
		fmt.Println("\tSi", result.Si)
		fmt.Println("\tSourceType1", result.SourceType1)
		fmt.Println("\tTime", result.Time)
		fmt.Println("\ttEntity", result.Entity)
		fmt.Println("\tHost", result.Host)
		fmt.Println("\tIndex", result.Index)
		fmt.Println("\tLineCount", result.LineCount)
		fmt.Println("\tLog", result.Log)
		fmt.Println("\tPunct", result.Punct)
		fmt.Println("\tSource", result.Source)
		fmt.Println("\tSourceType2", result.SourceType2)
		fmt.Println("\tSplunkServer", result.SplunkServer)
	}
	fmt.Println("Fields: ", searchModel.Fields)
	fmt.Println("Highlighted: ", searchModel.Highlighted)
}

func createJob() (string, string) {
	jobID, json, _ := splunkClient.SearchService.CreateJob(&model.PostJobsRequest{Query:"search index=*"})
	return jobID, json
}

func createSyncJob() *model.SearchEvents {
	searchModel, _ := splunkClient.SearchService.CreateSyncJob(&model.PostJobsRequest{Query:"search index=*"})
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
	jobID, json := createJob()
	fmt.Println(jobID)
	fmt.Println(json)

	///////////////////////////////////
	// 1b) retrieve a job results by id
	///////////////////////////////////
	//searchModel1 := getResults("1b9c6b21-3277-4dc6-b80a-76894678986f")
	//printSearchModel(searchModel1)

	////////////////////////////////////
	// 2) create a new synchronous search
	////////////////////////////////////
	//searchModel2 := createSyncJob()
	//printSearchModel(searchModel2)

	// TODO(dan): delete, get results for job id

}
