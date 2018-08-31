// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package playgroundintegration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
)

func TestIntegrationCreateEventsSuccess(t *testing.T) {
	client := getClient(t)
	attributes := make(map[string]interface{})
	attributes["testKey"] = "testValue"
	timeValue := int64(time.Now().Unix() * 1000) // Unix millis
	clientURL, err := client.GetURL()
	event1 := model.Event{
		Host:       clientURL.RequestURI(),
		Body:       "event1",
		Sourcetype: "sourcetype:eventgen",
		Source:     "manual-events",
		Timestamp:   timeValue,
		Attributes:  attributes}
	event2 := model.Event{
		Host:       clientURL.RequestURI(),
		Body:       "event2",
		Sourcetype: "sourcetype:eventgen",
		Source:     "manual-events",
		Timestamp:   timeValue,
		Attributes:  attributes}
	err = client.IngestService.CreateEvents([]model.Event{event1, event2})
	assert.Empty(t, err)
}

func TestIntegrationIngestEventFail(t *testing.T) {
	invalidClient := getInvalidClient(t)
	testIngestEvent := []model.Event{{Body: "failed test"}}
	err := invalidClient.IngestService.CreateEvents(testIngestEvent)

	assert.NotEmpty(t, err)
	assert.Equal(t, 401, err.(*util.HTTPError).HTTPStatusCode)
	assert.Equal(t, "401 Unauthorized", err.(*util.HTTPError).Message)
}

func TestIntegrationIngestEventBadRequest(t *testing.T) {
	client := getClient(t)
	err := client.IngestService.CreateEvents(nil)

	assert.NotEmpty(t, err)
	assert.Equal(t, 400, err.(*util.HTTPError).HTTPStatusCode)
	assert.Equal(t, "Invalid data format", err.(*util.HTTPError).Message)
	assert.Equal(t, "INVALID_DATA", err.(*util.HTTPError).Code)
	assert.Equal(t, "", err.(*util.HTTPError).MoreInfo)
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
		Timestamp:  time.Now().Unix() * 1000,
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
	err := client.IngestService.CreateMetricEvents([]model.MetricEvent{metricEvent1})
	assert.Empty(t, err)

	err1 := client.IngestService.CreateMetricEvents([]model.MetricEvent{metricEvent1, metricEvent1})
	assert.Empty(t, err1)
}
