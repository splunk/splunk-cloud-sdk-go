// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package stubbyintegration

import (
	"testing"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/service"
	"github.com/splunk/splunk-cloud-sdk-go/testutils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
)

func TestIngestEventFail(t *testing.T) {
	client, _ := service.NewClient(&service.Config{Token: "wrongToken", URL: testutils.TestURLProtocol + "://" + testutils.TestSplunkCloudHost, TenantID: testutils.TestTenantID, Timeout: time.Second * 5})
	err := client.IngestService.PostEvents([]model.Event{{Body: "failed test"}})
	assert.NotEmpty(t, err)
	assert.Equal(t, 401, err.(*util.HTTPError).HTTPStatusCode)
	assert.Equal(t, "401 Unauthorized", err.(*util.HTTPError).Message)
}

func TestCreateEvents(t *testing.T) {
	attributes := make(map[string]interface{})
	attributes["testKey"] = "testValue"
	timeValue := int64(time.Now().Unix() * 1000) // Unix millis
	event1 := model.Event{
		Host:       "stubby",
		Body:       "event1",
		Sourcetype: "sourcetype:eventgen",
		Source:     "manual-events",
		Timestamp:  timeValue,
		Attributes: attributes}
	event2 := model.Event{
		Host:       "stubby",
		Body:       "event2",
		Sourcetype: "sourcetype:eventgen",
		Source:     "manual-events",
		Timestamp:  timeValue,
		Attributes: attributes}
	err := getClient(t).IngestService.PostEvents([]model.Event{event1, event2})
	assert.Empty(t, err)
}

func TestIntegrationCreateMetrics(t *testing.T) {
	client := getClient(t)

	metrics := []model.Metric{
		{Name: "CPU", Value: 55.89,
			Dimensions: map[string]string{"Server": "redhat"}, Unit: "percentage"},

		{Name: "Memory", Value: 20.27,
			Dimensions: map[string]string{"Region": "us-east-5"}, Type: "g"},

		{Name: "Disk", Value: 15.444,
			Unit: "GB"},
	}

	metricEvent1 := model.MetricEvent{
		Body:       metrics,
		Timestamp:  1529020697,
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
}
