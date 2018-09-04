// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package main

import (
	"errors"
	"fmt"
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/testutils"
	"os"
	"time"
)

func main() {
	//Get client
	fmt.Println("Get client")
	client := getClient()

	//Create index
	fmt.Println("Create index")
	index, id := createIndex(client)
	if index != "main" {
		defer client.CatalogService.DeleteDataset(id)
	}

	//Ingest data
	fmt.Println("Ingest data")
	host, source := ingestData(client, index)

	//Ingest metrics data
	fmt.Println("Ingest metrics data")
	metricHost := ingestMetric(client, index)

	//Do search and verify results
	fmt.Println("Search event data")
	query := fmt.Sprintf("|from  index:%v where host=\"%v\" and source=\"%v\"", index, host, source)
	fmt.Println(query)
	search(client, query, 5)

	//Search metrics data and verify
	fmt.Println("Search metric data")
	query = fmt.Sprintf("| from metric:metrics group by host SELECT sum(CPU) as cpu,host |search host=\"%v\" AND cpu > 0", metricHost)
	fmt.Println(query)
	search(client, query, 1)
}

func checkIfQuit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getClient() *service.Client {
	var url = testutils.TestURLProtocol + "://" + testutils.TestSSCHost

	client, err := service.NewClient(&service.Config{
		Token:    testutils.TestAuthenticationToken,
		URL:      url,
		TenantID: testutils.TestTenantID,
		Timeout:  testutils.TestTimeOut})

	checkIfQuit(err)

	return client
}

func createIndex(client *service.Client) (string, string) {
	//index := fmt.Sprintf("goexample%v", float64(time.Now().Second()))
	index := "main"
	indexinfo := model.DatasetInfo{
		Owner:    "splunk",
		Name:     index,
		Kind:     "index",
		Disabled: false,
	}

	if index == "main" {
		return index, ""
	}

	result, err := client.CatalogService.CreateDataset(indexinfo)
	checkIfQuit(err)

	// it will take some time for the new index to finish the provisioning
	time.Sleep(30 * time.Second)
	return index, result.ID
}

func ingestMetric(client *service.Client, index string) string {
	host := fmt.Sprintf("gohost%v", time.Now().Unix())

	metrics := []model.Metric{
		{Name: "CPU", Value: 100,
			Dimensions: map[string]string{"Server": "redhat"}, Unit: "percentage"},

		{Name: "Memory", Value: 20.27,
			Dimensions: map[string]string{"Region": "us-east-5"}, Type: "g"},

		{Name: "Disk", Value: 15.444,
			Unit: "GB"},
	}

	metricEvent1 := model.MetricEvent{
		Body:       metrics,
		Timestamp:  time.Now().Unix() * 1000,
		Nanos:      1,
		Source:     "mysource",
		Sourcetype: "newsourcetype",
		Host:       host,
		ID:         "metric0001",
		Attributes: model.MetricAttribute{
			DefaultDimensions: map[string]string{"defaultDimensions": ""},
			DefaultType:       "g",
			DefaultUnit:       "MB",
		},
	}

	// Use the Ingest Service raw endpoint to send data
	err := client.IngestService.CreateMetricEvents([]model.MetricEvent{metricEvent1, metricEvent1, metricEvent1})
	checkIfQuit(err)

	return host
}

func ingestData(client *service.Client, index string) (string, string) {
	source := fmt.Sprintf("mysource-%v", float64(time.Now().Second()))
	host := fmt.Sprintf("myhost-%v", float64(time.Now().Second()))

	event1 := model.Event{
		Host:   host,
		Source: source,
		Event:  fmt.Sprintf("device_id=aa1 haha0 my new event %v,%v", host, source),
		Index:  index}

	event2 := model.Event{
		Host:   host,
		Source: source,
		Event:  fmt.Sprintf("04-24-2018 12:32:23.252 -0700 INFO  device_id=[www]401:sdfsf haha1 %v,%v", host, source),
		Fields: map[string]string{"fieldkey1": "fieldval1", "fieldkey2": "fieldkey2"},
		Index:  index}

	event3 := model.Event{
		Host:       host,
		Source:     source,
		Sourcetype: "splunkd",
		Event:      fmt.Sprintf("04-24-2018 12:32:23.258 -0700 INFO device_id:aa2 device_id=[code]error3: haha2 \"9765f1bebdb4\" %v,%v", host, source),
		Fields:     map[string]string{"fieldkey1": "fieldval1", "fieldkey2": "fieldkey2"},
		Index:      index}

	// Use the Ingest Service raw endpoint to send data
	err := client.IngestService.CreateRawEvent(event1)
	checkIfQuit(err)

	// Use the Ingest Service endpoint to send one event
	err = client.IngestService.CreateEvent(event1)
	checkIfQuit(err)

	// Use the Ingest endpoint to send multiple events
	err = client.IngestService.CreateEvents([]model.Event{event1, event2, event3})
	checkIfQuit(err)

	return host, source
}

func search(client *service.Client, query string, expected int) {
	start := time.Now()
	timeout := 60 * time.Second
	for {
		if time.Now().Sub(start) > timeout {
			checkIfQuit(errors.New("Search failed due to timeout "))
		}

		sid, err := client.SearchService.CreateJob(&model.PostJobsRequest{Search: query})
		checkIfQuit(err)

		err = client.SearchService.WaitForJob(sid, 1000*time.Millisecond)
		checkIfQuit(err)

		resp, err := client.SearchService.GetJobResults(sid, &model.FetchResultsRequest{Count: 100})
		fmt.Println(resp.Results)
		checkIfQuit(err)

		if len(resp.Results) == expected {
			fmt.Println("Search succeed")

			return
		}

		if len(resp.Results) < expected {
			fmt.Println("Not found all yet, keep searching")
			time.Sleep(20 * time.Second)
		} else {
			checkIfQuit(errors.New("Search failed: Get more results than expected"))
		}
	}
}
