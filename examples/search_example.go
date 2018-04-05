package main

import (
	"github.com/splunk/ssc-client-go/lib/service"
	"github.com/splunk/ssc-client-go/lib/model"
	"time"
	"fmt"
)

const (
	ClientId     = "4zRqusbLAq754mX5WCDfoiQFzFJFWWkO"
	ClientSecret = "ff9odDwxiZqSVEQzcBeOU-_ALDLKksXlELySNdjkbPxRH7rV9gybNhhbgbucteGe"
	BaseURL      = "api.splunknovadev-playground.com"
	Timeout 	 = time.Second * 5
)

var splunkClient = service.NewSplunkdClient(
"", [2]string{ClientId, ClientSecret}, BaseURL, service.NewSplunkdHTTPClient(Timeout, true))


func createJob() (string) {
	searchId, _ := splunkClient.SearchService.CreateJob("search index=*")
	return searchId
}

func createSyncJob() (*model.SearchEvents) {
	searchModel, _ := splunkClient.SearchService.CreateSyncJob("search index=*")
	return searchModel
}

func getJob(searchId string) {
	splunkClient.SearchService.GetJob(searchId)
}


func main() {
	///////////////////////////////
	// 1) create a new search job
	///////////////////////////////
	//searchId := createJob()
	//fmt.Println(searchId)

	///////////////////////////////
	// 2) retrieve a job by id
	///////////////////////////////
	//getJob(searchId)

	////////////////////////////////////
	// 3) create a new synchronous search
	////////////////////////////////////
	searchModel := createSyncJob()
	fmt.Println("Preview: ", searchModel.Preview)
	fmt.Println("InitOffset: ", searchModel.InitOffset)
	fmt.Println("Messages: ", searchModel.Messages)

	//fmt.Println("Results: ", searchModel.Results)
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


	// TODO(dan): delete, get results for job id


}
