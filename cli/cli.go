package main

import (
	"github.com/splunk/ssc-client-go/lib/service"
	"flag"
	"time"
	"fmt"
)


const (
	ClientId = "REDACTED"
	ClientSecret = "REDACTED"
	BaseURL = "api.splunknovadev-playground.com"
)



var (
	// Initialize parameters needed for communicating with Splunkd
	user     = flag.String("user", "admin", "Splunk username, defaults to admin")
	password = flag.String("password", "changed", "Splunk password, defaults to changeme")
	host     = flag.String("host", "localhost:8089", "Splunkd host, defaults to localhost:8089")
	search   = flag.String("search", "", "The search spl, defaults to an empty string")
	timeout  = flag.Duration("timeout", time.Second*5, "The timeout used for the client, defaults to 5 seconds")
)

func main() {
	flag.Parse()

	//var err error
	splunkdClient := service.NewSplunkdClient("", [2]string{ClientId, ClientSecret}, BaseURL, service.NewSplunkdHTTPClient(*timeout, true))

	url := splunkdClient.BuildSplunkdURL(nil, "v1", "events")

	body := make(map[string]string)
	body["log"] = "This is my first Nova event"
	body["source"] = "fake cli"
	body["entity"] = "test_api"

	resp, err := splunkdClient.Post(url, body)

	if err != nil {
		fmt.Println(resp)
	} else {
		fmt.Println(err)
	}




}