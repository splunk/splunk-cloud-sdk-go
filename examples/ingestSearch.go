// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package main

import (
	"errors"
	"fmt"
	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/service"
	"github.com/splunk/splunk-cloud-sdk-go/testutils"
	"os"
	"time"
)

func main() {
	//Get client
	fmt.Println("Get client")
	client := getClient()

	//todo will need to wait pipeline to get index data to non-main index
	////Create index
	//fmt.Println("Create index")
	//index, id := createIndex(client)
	//if index != "main" {
	//	defer client.CatalogService.DeleteDataset(id)
	//}
	index:="main"

	//Ingest data
	fmt.Println("Ingest event data")
	host, source := ingestEvent(client, index)

	//Ingest metrics data
	fmt.Println("Ingest metrics data")
	metricHost := ingestMetric(client, index)

	//Do search and verify results
	fmt.Println("Search event data")
	query := fmt.Sprintf("|from  index:%v where host=\"%v\" and source=\"%v\"", index, host, source)
	fmt.Println(query)
	search(client, query, 3)

	//Search metrics data and verify
	fmt.Println("Search metric data")
	query = fmt.Sprintf("| from metric:metrics group by host SELECT sum(CPU) as cpu,host |search host=\"%v\" AND cpu > 0", metricHost)
	fmt.Println(query)
	search(client, query, 1)
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getClient() *service.Client {
	client, err := service.NewClient(&service.Config{
		Token:  testutils.TestAuthenticationToken,
		Scheme: testutils.TestURLProtocol,
		Host:   testutils.TestSplunkCloudHost,
		Tenant: testutils.TestTenantID,
	})

	exitOnError(err)

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
	exitOnError(err)

	// it will take some time for the new index to finish the provisioning
	// todo: user dataset endpoint to check the readyness
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

	// Use the Ingest Service send metrics
	err := client.IngestService.PostMetrics([]model.MetricEvent{metricEvent1, metricEvent1, metricEvent1})
	exitOnError(err)

	return host
}

func ingestEvent(client *service.Client, index string) (string, string) {
	source := fmt.Sprintf("mysource-%v", float64(time.Now().Second()))
	host := fmt.Sprintf("myhost-%v", float64(time.Now().Second()))

	event1 := model.Event{
		Host:   host,
		Source: source,
		Body:  fmt.Sprintf("device_id=aa1 haha0 my new event %v,%v", host, source),
		}

	event2 := model.Event{
		Host:   host,
		Source: source,
		Body:  fmt.Sprintf("04-24-2018 12:32:23.252 -0700 INFO  device_id=[www]401:sdfsf haha1 %v,%v", host, source),
		Attributes: map[string]interface{}{"fieldkey1": "fieldval1", "fieldkey2": "fieldkey2"},
		}

	event3 := model.Event{
		Host:       host,
		Source:     source,
		Sourcetype: "splunkd",
		Body:      fmt.Sprintf("04-24-2018 12:32:23.258 -0700 INFO device_id:aa2 device_id=[code]error3: haha2 \"9765f1bebdb4\" %v,%v", host, source),
		Attributes:     map[string]interface{}{"fieldkey1": "fieldval1", "fieldkey2": "fieldkey2"},
		}

	// Use the Ingest endpoint to send multiple events
	err := client.IngestService.PostEvents([]model.Event{event1,event2,event3})
	exitOnError(err)

	return host, source
}

func search(client *service.Client, query string, expected int) {
	start := time.Now()
	timeout := 60 * time.Second
	for {
		if time.Now().Sub(start) > timeout {
			exitOnError(errors.New("Search failed due to timeout "))
		}

		sid, err := client.SearchService.CreateJob(&model.PostJobsRequest{Search: query})
		exitOnError(err)

		err = client.SearchService.WaitForJob(sid, 1000*time.Millisecond)
		exitOnError(err)

		resp, err := client.SearchService.GetJobResults(sid, &model.FetchResultsRequest{Count: 100})
		fmt.Println(resp.Results)
		exitOnError(err)

		if len(resp.Results) == expected {
			fmt.Println("Search succeed")

			return
		}

		if len(resp.Results) < expected {
			fmt.Println("Not found all yet, keep searching")
			time.Sleep(20 * time.Second)
		} else {
			exitOnError(errors.New("Search failed: Get more results than expected"))
		}
	}
}
