// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/services/catalog"
	"github.com/splunk/splunk-cloud-sdk-go/services/ingest"
	"github.com/splunk/splunk-cloud-sdk-go/services/search"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
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
	index := "main"

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
	performSearch(client, query, 3)

	//Search metrics data and verify
	fmt.Println("Search metric data")
	query = fmt.Sprintf("| from metrics group by host SELECT sum(CPU) as cpu,host |search host=\"%v\" AND cpu > 0", metricHost)
	fmt.Println(query)
	performSearch(client, query, 1)
}

func exitOnError(err error) {
	if err != nil {
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

	exitOnError(err)

	return client
}

func createIndex(client *sdk.Client) (string, string) {
	//index := fmt.Sprintf("goexample%v", float64(time.Now().Second()))
	index := "main"
	disabled := false
	indexmodel := &catalog.IndexDataset{
		Name:     index,
		Kind:     catalog.Index,
		Disabled: &disabled,
	}

	if index == "main" {
		return index, ""
	}

	result, err := client.CatalogService.CreateIndexDataset(indexmodel)
	exitOnError(err)

	// it will take some time for the new index to finish the provisioning
	// todo: user dataset endpoint to check the readyness
	time.Sleep(30 * time.Second)
	return index, string(result.ID)
}

func ingestMetric(client *sdk.Client, index string) string {
	host := fmt.Sprintf("gohost%v", time.Now().Unix())

	metrics := []ingest.Metric{
		{Name: "CPU", Value: 100,
			Dimensions: map[string]string{"Server": "redhat"}, Unit: "percentage"},

		{Name: "Memory", Value: 20.27,
			Dimensions: map[string]string{"Region": "us-east-5"}, Type: "g"},

		{Name: "Disk", Value: 15.444,
			Unit: "GB"},
	}

	metricEvent1 := ingest.MetricEvent{
		Body:       metrics,
		Timestamp:  time.Now().Unix() * 1000,
		Nanos:      1,
		Source:     "mysource",
		Sourcetype: "newsourcetype",
		Host:       host,
		ID:         "metric0001",
		Attributes: ingest.MetricAttribute{
			DefaultDimensions: map[string]string{"defaultDimensions": ""},
			DefaultType:       "g",
			DefaultUnit:       "MB",
		},
	}

	// Use the Ingest Service send metrics
	err := client.IngestService.PostMetrics([]ingest.MetricEvent{metricEvent1, metricEvent1, metricEvent1})
	exitOnError(err)

	return host
}

func ingestEvent(client *sdk.Client, index string) (string, string) {
	source := fmt.Sprintf("mysource-%v", time.Now().Unix())
	host := fmt.Sprintf("myhost-%v", time.Now().Unix())

	event1 := ingest.Event{
		Host:   host,
		Source: source,
		Body:   fmt.Sprintf("device_id=aa1 haha0 my new event %v,%v", host, source),
	}

	event2 := ingest.Event{
		Host:       host,
		Source:     source,
		Body:       fmt.Sprintf("04-24-2018 12:32:23.252 -0700 INFO  device_id=[www]401:sdfsf haha1 %v,%v", host, source),
		Attributes: map[string]interface{}{"index": index},
	}

	event3 := ingest.Event{
		Host:       host,
		Source:     source,
		Sourcetype: "splunkd",
		Body:       fmt.Sprintf("04-24-2018 12:32:23.258 -0700 INFO device_id:aa2 device_id=[code]error3: haha2 \"9765f1bebdb4\" %v,%v", host, source),
		Attributes: map[string]interface{}{"index": index},
	}

	// Use the Ingest endpoint to send multiple events
	err := client.IngestService.PostEvents([]ingest.Event{event1, event2, event3})
	exitOnError(err)

	return host, source
}

func performSearch(client *sdk.Client, query string, expected int) {
	start := time.Now()
	timeout := 60 * time.Second
	for {
		if time.Now().Sub(start) > timeout {
			exitOnError(errors.New("Search failed due to timeout "))
		}

		job, err := client.SearchService.CreateJob(&search.CreateJobRequest{Query: query})
		exitOnError(err)

		_, err = client.SearchService.WaitForJob(job.ID, 1000*time.Millisecond)
		exitOnError(err)

		resp, err := client.SearchService.GetResults(job.ID, 100, 0)
		results := resp.(*search.Results).Results
		fmt.Println(results)
		exitOnError(err)

		if len(results) >= expected {
			fmt.Println("Search succeed")

			return
		}

		// TODO: Duplicates occurring when ingesting new data. Known issue (SSC-4179). Should follow up with ingest team.
		if len(results) < expected {
			fmt.Println("Not found all yet, keep searching")
			time.Sleep(20 * time.Second)
		} /*else {
			exitOnError(errors.New("Search failed: Get more results than expected"))
		}*/
	}
}
