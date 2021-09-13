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

package integration

import (
	"net/http"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/services/ingest"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationCreateEventsSuccess(t *testing.T) {
	client := getClient(t)
	attributes := make(map[string]interface{})
	attributes1 := make(map[string]interface{})
	attributes1["test1"] = "test1"
	attributes["testkey1"] = attributes1
	timeValue := int64(testutils.RunSuffix * 1000) // Unix millis
	clientURL := client.GetURL("")
	hostURL := clientURL.RequestURI()
	body1 := make(map[string]interface{})
	body1["event"] = "event1"
	body2 := make(map[string]interface{})
	body2["event"] = "event2"

	sourcetypeStr := "sourcetype:eventgen"
	sourceStr := "manual-events"
	event1 := ingest.Event{
		Host:       &hostURL,
		Body:       body1,
		Sourcetype: &sourcetypeStr,
		Source:     &sourceStr,
		Timestamp:  &timeValue,
		Attributes: attributes}
	event2 := ingest.Event{
		Host:       &hostURL,
		Body:       body2,
		Sourcetype: &sourcetypeStr,
		Source:     &sourceStr,
		Timestamp:  &timeValue,
		Attributes: attributes}
	_, err := client.IngestService.PostEvents([]ingest.Event{event1, event2})
	assert.Empty(t, err)
}

func TestIntegrationIngestEventFail(t *testing.T) {
	invalidClient := getInvalidClient(t)
	body1 := make(map[string]interface{})
	body1["event"] = "failed test"
	testIngestEvent := []ingest.Event{{Body: body1}}
	_, err := invalidClient.IngestService.PostEvents(testIngestEvent)

	assert.NotEmpty(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 401, httpErr.HTTPStatusCode)
	errFound := httpErr.Message == "Error validating request" || httpErr.Message == "Invalid or Expired Bearer Token"
	assert.True(t, errFound)
}

func TestIntegrationIngestEventsFailureDetails(t *testing.T) {
	client := getClient(t)
	body1 := make(map[string]interface{})
	body1["event"] = "failed test"
	event1 := ingest.Event{Body: body1}
	event2 := ingest.Event{}
	_, err := client.IngestService.PostEvents([]ingest.Event{event1, event2})

	assert.NotEmpty(t, err)

	httperror, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 400, httperror.HTTPStatusCode)
	assert.Equal(t, "The request isn't valid.", httperror.Message)

	details, ok := httperror.Details.(map[string]interface{})
	require.True(t, ok)
	assert.NotEmpty(t, details["failedEvents"])

	failedEvents, ok := details["failedEvents"].([]interface{})
	require.True(t, ok)
	assert.Equal(t, 1, len(failedEvents))

	failedEvent, ok := failedEvents[0].(map[string]interface{})
	require.True(t, ok)
	assert.Equal(t, float64(1), failedEvent["index"])
	msg, ok := failedEvent["message"].(string)
	require.True(t, ok)
	assert.Equal(t, "event body cannot be null", strings.ToLower(msg))
}

func TestIntegrationIngestEventBadRequest(t *testing.T) {
	client := getClient(t)
	_, err := client.IngestService.PostEvents(nil)
	assert.NotEmpty(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 400, httpErr.HTTPStatusCode)
	assert.Equal(t, "The request isn't valid.", httpErr.Message)
	assert.Equal(t, "INVALID_DATA", httpErr.Code)
	assert.Equal(t, "", httpErr.MoreInfo)
}

func TestIntegrationCreateMetrics(t *testing.T) {
	client := getClient(t)
	value := float64(5.89)
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
	host := "myhost"
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
	_, err := client.IngestService.PostMetrics([]ingest.MetricEvent{metricEvent1})
	assert.Empty(t, err)

	_, err1 := client.IngestService.PostMetrics([]ingest.MetricEvent{metricEvent1, metricEvent1})
	assert.Empty(t, err1)
}

func TestIntegrationUploadFiles(t *testing.T) {
	client := getClient(t)

	dir, err := os.Getwd()
	require.NoError(t, err)

	var resp http.Response
	filename := path.Join(dir, "ingest_integration_test.go")

	err = client.IngestService.UploadFiles(filename, &resp)
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
}

func TestIntegrationUploadFilesNonexistfile(t *testing.T) {
	client := getClient(t)

	err := client.IngestService.UploadFiles("nonexit")
	require.NotNil(t, err)
}
