package main

import (
	"github.com/splunk/ssc-client-go/lib/service"
	"flag"
	"time"
	"fmt"
	//"bytes"
)


const (
	ClientId = "4zRqusbLAq754mX5WCDfoiQFzFJFWWkO"
	ClientSecret = "ff9odDwxiZqSVEQzcBeOU-_ALDLKksXlELySNdjkbPxRH7rV9gybNhhbgbucteGe"
	BaseURL = "api.splunknovadev-playground.com"
)

var (
	// Initialize parameters needed for communicating with Splunkd
	user     = flag.String("user", "admin", "Splunk username, defaults to admin")
	password = flag.String("password", "changed", "Splunk password, defaults to changeme")
	host     = flag.String("host", "localhost:8089", "Splunkd host, defaults to localhost:8089")
	search   = flag.String("search", "search index=*", "The search spl, defaults to an empty string")
	timeout  = flag.Duration("timeout", time.Second*5, "The timeout used for the client, defaults to 5 seconds")
)

func main() {
	flag.Parse()

	//var err error
	splunkdClient := service.NewSplunkdClient(
		"", [2]string{ClientId, ClientSecret}, BaseURL, service.NewSplunkdHTTPClient(*timeout, true))

	//url := splunkdClient.BuildSplunkdURL(nil,  "v1", "events")

	//body := make(map[string]string)
	//body["log"] = "This is my first Nova event"
	//body["source"] = "fake cli"
	//body["entity"] = "test_api"
	//
	//resp, err := splunkdClient.Post(url, body)
	//
	//if err != nil {
	//	fmt.Println(resp)
	//
	//	defer resp.Body.Close()
	//	b := new(bytes.Buffer)
	//	b.ReadFrom(resp.Body)
	//	fmt.Println(b)
	//
	//} else {
	//	fmt.Println(err)
	//}

	jobModel, _ := splunkdClient.SearchService.NewSearch(*search)

	fmt.Sprintf("sid: %s ", jobModel.Sid)


}