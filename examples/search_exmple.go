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
	splunkClient.SearchService.GetSearch(searchId)
}


func main() {

	// create a new search job
	//searchId := createJob()
	//fmt.Println(searchId)

	////////////////////////////////////
	// create a new synchronous search
	////////////////////////////////////
	searchModel := createSyncJob()
	fmt.Println("Preview: ", searchModel.Preview)
	fmt.Println("InitOffset: ", searchModel.InitOffset)
	fmt.Println("Messages: ", searchModel.Messages)

	//fmt.Println("Results: ", searchModel.Results)
	fmt.Print("Results: ")
	for _, result := range searchModel.Results {
		fmt.Println("Bkt", result.Bkt)
		fmt.Println("Cd", result.Cd)
		fmt.Println("IndexTime", result.IndexTime)
		fmt.Println("Raw", result.Raw)
		fmt.Println("Serial", result.Serial)
		fmt.Println("Si", result.Si)
		fmt.Println("SourceType1", result.SourceType1)
		fmt.Println("Time", result.Time)
		fmt.Println("Entity", result.Entity)
		fmt.Println("Host", result.Host)
		fmt.Println("Index", result.Index)
		fmt.Println("LineCount", result.LineCount)
		fmt.Println("Log", result.Log)
		fmt.Println("Punct", result.Punct)
		fmt.Println("Source", result.Source)
		fmt.Println("SourceType2", result.SourceType2)
		fmt.Println("SplunkServer", result.SplunkServer)
		}
	fmt.Println("Fields: ", searchModel.Fields)
	fmt.Println("Highlighted: ", searchModel.Highlighted)


	// retrieve a job by id
	//getSearch(searchId)


}
