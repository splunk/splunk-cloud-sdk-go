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
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/v2/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/v2/services"
	"github.com/splunk/splunk-cloud-sdk-go/v2/services/catalog"
	"github.com/splunk/splunk-cloud-sdk-go/v2/services/ingest"
	"github.com/splunk/splunk-cloud-sdk-go/v2/services/search"
	testutils "github.com/splunk/splunk-cloud-sdk-go/v2/test/utils"
)

func main() {
	//Get client
	fmt.Println("Get client")
	client := getClient()

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
	indexmodel := catalog.IndexDatasetPost{
		Name:     index,
		Kind:     catalog.IndexDatasetKindIndex,
		Disabled: disabled,
	}

	if index == "main" {
		return index, ""
	}

	result, err := client.CatalogService.CreateDataset(catalog.MakeDatasetPostFromIndexDatasetPost(indexmodel))
	exitOnError(err)

	// it will take some time for the new index to finish the provisioning
	// todo: user dataset endpoint to check the readyness
	time.Sleep(30 * time.Second)
	return index, string(result.IndexDataset().Id)
}

func ingestMetric(client *sdk.Client, index string) string {
	host := fmt.Sprintf("gohost%v", time.Now().Unix())
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

	timestamp := testutils.TimeSec * 1000
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
	exitOnError(err)

	return host
}

func ingestEvent(client *sdk.Client, index string) (string, string) {
	source := fmt.Sprintf("mysource-%v", time.Now().Unix())
	host := fmt.Sprintf("myhost-%v", time.Now().Unix())
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
	exitOnError(err)

	return host, source
}

func performSearch(client *sdk.Client, query string, expected int) {
	start := time.Now()
	timeout := 240 * time.Second
	for {
		if time.Now().Sub(start) > timeout {
			exitOnError(errors.New(fmt.Sprintf("Search failed due to timeout, spend %v", timeout)))
		}

		job, err := client.SearchService.CreateJob(search.SearchJob{Query: query})
		exitOnError(err)

		_, err = client.SearchService.WaitForJob(*job.Sid, 1000*time.Millisecond)
		exitOnError(err)

		query := search.ListResultsQueryParams{}.SetCount(100).SetOffset(0)
		resp, err := client.SearchService.ListResults(*job.Sid, &query)
		exitOnError(err)

		results := (*resp).Results
		fmt.Printf("expect %d, now get %d results\n", expected, len(results))

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
