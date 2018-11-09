// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationCreateEventsSuccess(t *testing.T) {
	client := getClient(t)
	attributes := make(map[string]interface{})
	attributes["testKey"] = "testValue"
	timeValue := int64(testutils.TimeSec * 1000) // Unix millis
	clientURL := client.GetURL()
	event1 := model.Event{
		Host:       clientURL.RequestURI(),
		Body:       "event1",
		Sourcetype: "sourcetype:eventgen",
		Source:     "manual-events",
		Timestamp:  timeValue,
		Attributes: attributes}
	event2 := model.Event{
		Host:       clientURL.RequestURI(),
		Body:       "event2",
		Sourcetype: "sourcetype:eventgen",
		Source:     "manual-events",
		Timestamp:  timeValue,
		Attributes: attributes}
	err := client.IngestService.PostEvents([]model.Event{event1, event2})
	assert.Empty(t, err)
}

func TestIntegrationIngestEventFail(t *testing.T) {
	invalidClient := getInvalidClient(t)
	testIngestEvent := []model.Event{{Body: "failed test"}}
	err := invalidClient.IngestService.PostEvents(testIngestEvent)

	assert.NotEmpty(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 401, httpErr.HTTPStatusCode)
	assert.Equal(t, "Error validating request", httpErr.Message)
}

func TestIntegrationIngestEventsFailureDetails(t *testing.T) {
	client := getClient(t)
	event1 := model.Event{Body: "some event"}
	event2 := model.Event{}
	err := client.IngestService.PostEvents([]model.Event{event1, event2})

	assert.NotEmpty(t, err)

	httperror, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 400, httperror.HTTPStatusCode)
	assert.Equal(t, "Invalid data format", httperror.Message)

	details, ok := httperror.Details.(map[string]interface{})
	require.True(t, ok)
	assert.NotEmpty(t, details["failedEvents"])

	failedEvents, ok := details["failedEvents"].([]interface{})
	require.True(t, ok)
	assert.Equal(t, 1, len(failedEvents))

	failedEvent, ok := failedEvents[0].(map[string]interface{})
	require.True(t, ok)
	assert.Equal(t, float64(1), failedEvent["index"])
	assert.Equal(t, "Event body cannot be empty", failedEvent["message"])
}

func TestIntegrationIngestEventBadRequest(t *testing.T) {
	client := getClient(t)
	err := client.IngestService.PostEvents(nil)
	assert.NotEmpty(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 400, httpErr.HTTPStatusCode)
	assert.Equal(t, "Invalid data format", httpErr.Message)
	assert.Equal(t, "INVALID_DATA", httpErr.Code)
	assert.Equal(t, "", httpErr.MoreInfo)
}

func TestIntegrationCreateMetrics(t *testing.T) {
	client := getClient(t)

	metrics := []model.Metric{
		{Name: "CPU", Value: 5.89,
			Dimensions: map[string]string{"Server": "redhat"}, Unit: "percentage"},

		{Name: "Memory", Value: 20.27,
			Dimensions: map[string]string{"Region": "us-east-5"}, Type: "g"},

		{Name: "Disk", Value: 15.444,
			Unit: "GB"},
	}

	metricEvent1 := model.MetricEvent{
		Body:       metrics,
		Timestamp:  testutils.TimeSec * 1000,
		Nanos:      1,
		Source:     "mysource",
		Sourcetype: "mysourcetype",
		Host:       "myhost",
		ID:         "metric0001",
		Attributes: model.MetricAttribute{
			DefaultDimensions: map[string]string{"defaultDimensions": ""},
			DefaultType:       "g",
			DefaultUnit:       "MB",
		},
	}
	err := client.IngestService.PostMetrics([]model.MetricEvent{metricEvent1})
	assert.Empty(t, err)

	err1 := client.IngestService.PostMetrics([]model.MetricEvent{metricEvent1, metricEvent1})
	assert.Empty(t, err1)
}
