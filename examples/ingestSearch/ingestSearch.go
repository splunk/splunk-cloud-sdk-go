/*
 * Copyright 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/util"

	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/services/ingest"
	"github.com/splunk/splunk-cloud-sdk-go/services/search"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
)

func main() {
	//Get client
	fmt.Println("Get client")
	client := getClient()

	index := "main"

	// An example showing UploadFiles functionality of ingest service
	demoIngestServiceUploadFile(client)

	//Ingest data
	host, source := demoIngestServiceIngestEvents(client)

	//Ingest metrics data
	metricHost := demoIngestServiceIngestMetric(client)

	//Do search and verify results
	fmt.Println("Search event data")
	query := fmt.Sprintf("|from  index:%v where host=\"%v\" and source=\"%v\"", index, host, source)
	fmt.Println(query)
	demoSearchServiceSearchResults(client, query, 3, false)

	//Search metrics data and verify
	fmt.Println("Search metric data")
	query = fmt.Sprintf("| from metrics group by host SELECT sum(CPU) as cpu,host |search host=\"%v\" AND cpu > 0", metricHost)
	fmt.Println(query)
	demoSearchServiceSearchResults(client, query, 1, false)
}

func demoIngestServiceUploadFile(client *sdk.Client) {
	dir, err := os.Getwd()

	var resp http.Response
	filename := path.Join(dir, "examples/ingestSearch/ingestSearch.go")
	err = client.IngestService.UploadFiles(filename, &resp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("INFO: File upload using ingest service successfully completed")
}

// Based on 'shouldFailOnAnyError' flag value, it decides whether 429 or 500 errors that indicate that the service is overloaded should fail the example or not
// Setting this to true will exit the program run if any error is encountered
// Setting this to false will exit the program run on all errors except 429 or 500 errors
func handleError(err error, shouldFailOnAnyError bool) {
	if shouldFailOnAnyError == false {
		httpErr, _ := err.(*util.HTTPError)
		fmt.Println("http Error: ", httpErr)
		if httpErr.HTTPStatusCode == 429 || httpErr.HTTPStatusCode == 500 {
			fmt.Printf("INFO: Skipping example - Service is overloaded. Error message received is: %s, "+
				"Error status Code is: %d, Error Status is: %s, Error code is: %s", httpErr.Message, httpErr.HTTPStatusCode, httpErr.HTTPStatus, httpErr.Code)
		}
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getClient() *sdk.Client {
	client, err := sdk.NewClient(&services.Config{
		Token:  testutils.TestAuthenticationToken,
		Host:   testutils.TestSplunkCloudHost,
		Tenant: testutils.TestTenant,
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return client
}

func demoIngestServiceIngestMetric(client *sdk.Client) string {
	fmt.Println("Ingest metrics data")
	host := fmt.Sprintf("gohost%v", time.Now().Unix())
	fmt.Println("Ingest metric Host: ", host)
	value := float64(100)
	unit := "percentage"
	value1 := float64(20.27)
	unit1 := "GB"
	value2 := float64(15.444)
	typestr := "g"

	metrics := []ingest.Metric{
		{Name: "CPU", Value: &value,
			Dimensions: map[string]string{"Server": "redhat"}, Unit: &unit},

		{Name: "Memory", Value: &value1,
			Dimensions: map[string]string{"Region": "us-east-5"}, Type: &typestr},

		{Name: "Disk", Value: &value2,
			Unit: &unit1},
	}

	timestamp := testutils.RunSuffix * 1000
	nanos := int32(1)
	source := "mysource"
	sourcetype := "mysourcetype"
	id := "metric0001"
	defaulttype := "g"
	defaultunit := "MB"

	metricEvent1 := ingest.MetricEvent{
		Body:       metrics,
		Timestamp:  &timestamp,
		Nanos:      &nanos,
		Source:     &source,
		Sourcetype: &sourcetype,
		Host:       &host,
		Id:         &id,
		Attributes: &ingest.MetricAttribute{
			DefaultDimensions: map[string]string{"defaultDimensions": ""},
			DefaultType:       &defaulttype,
			DefaultUnit:       &defaultunit,
		},
	}

	// Use the Ingest Service send metrics
	_, err := client.IngestService.PostMetrics([]ingest.MetricEvent{metricEvent1, metricEvent1, metricEvent1})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return host
}

func demoIngestServiceIngestEvents(client *sdk.Client) (string, string) {
	fmt.Println("Ingest event data")
	source := fmt.Sprintf("mysource-%v", time.Now().Unix())
	fmt.Println("Ingest events Source: ", source)
	host := fmt.Sprintf("myhost-%v", time.Now().Unix())
	fmt.Println("Ingest events Host: ", host)
	body := make(map[string]interface{})
	body["event"] = fmt.Sprintf("device_id=aa1 haha0 my new event %v,%v", host, source)

	event1 := ingest.Event{
		Host:   &host,
		Source: &source,
		Body:   body,
	}

	body["event"] = fmt.Sprintf("04-24-2018 12:32:23.252 -0700 INFO  device_id=[www]401:sdfsf haha1 %v,%v", host, source)
	attributes := make(map[string]interface{})
	attributes1 := make(map[string]interface{})
	attributes1["index"] = "index"
	attributes["attr"] = attributes1

	event2 := ingest.Event{
		Host:       &host,
		Source:     &source,
		Body:       body,
		Attributes: attributes,
	}

	sourcetype := "splunkd"
	body["event"] = fmt.Sprintf("04-24-2018 12:32:23.258 -0700 INFO device_id:aa2 device_id=[code]error3: haha2 \"9765f1bebdb4\" %v,%v", host, source)

	event3 := ingest.Event{
		Host:       &host,
		Source:     &source,
		Sourcetype: &sourcetype,
		Body:       body,
		Attributes: attributes,
	}

	// Use the Ingest endpoint to send multiple events
	_, err := client.IngestService.PostEvents([]ingest.Event{event1, event2, event3})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return host, source
}

// Query
func demoSearchServiceSearchResults(client *sdk.Client, query string, expected int, shouldFailOnAnyError bool) {
	start := time.Now()
	timeout := 160 * time.Second
	for {
		if time.Now().Sub(start) > timeout {
			fmt.Printf("INFO: Unable to fetch the search results, Search exceeded the timeout period, Wait longer before fetching the data, spend %v \n", timeout)
			break
		}

		fmt.Println("INFO: Creating a new job")
		job, err := client.SearchService.CreateJob(search.SearchJob{Query: query})
		if err != nil {
			handleError(err, shouldFailOnAnyError)
			break
		}

		fmt.Println("INFO: Waiting for job until it completes")
		_, err = client.SearchService.WaitForJob(*job.Sid, 1000*time.Millisecond)
		if err != nil {
			handleError(err, shouldFailOnAnyError)
			break
		}

		fmt.Printf("INFO: Fetching results for the job with job ID: %s \n", *job.Sid)
		query := search.ListResultsQueryParams{}.SetCount(100).SetOffset(0)
		resp, err := client.SearchService.ListResults(*job.Sid, &query)
		if err != nil {
			handleError(err, shouldFailOnAnyError)
			break
		}

		results := (*resp).Results
		fmt.Printf("expect %d, now get %d results\n", expected, len(results))

		if len(results) >= expected {
			fmt.Println("INFO: Search succeed")
			return
		}

		if len(results) < expected {
			fmt.Println("Not found all yet, keep searching, data takes a while to appear")
			time.Sleep(40 * time.Second)
		}
	}
}
