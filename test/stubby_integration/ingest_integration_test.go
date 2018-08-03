/*
 * Copyright © 2018 Splunk Inc.
 * SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
 * without a valid written license from Splunk Inc. is PROHIBITED.
 *
 */

package stubbyintegration

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/splunk/ssc-client-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateEventSuccess(t *testing.T) {
	timeValue := float64(1523637597)

	client := getClient(t)
	clientURL, err := client.GetURL()
	assert.Empty(t, err)

	err = client.IngestService.CreateEvent(
		model.Event{Host: clientURL.RequestURI(), Index: "main", Event: "test", Sourcetype: "sourcetype:eventgen", Source: "manual-events", Time: &timeValue, Fields: map[string]string{"testKey": "testValue"}})
	assert.Empty(t, err)
}

func TestCreateRawEventSuccess(t *testing.T) {
	err := getClient(t).IngestService.CreateRawEvent(
		model.Event{Event: "test"})
	assert.Empty(t, err)
}

func TestIngestEventFail(t *testing.T) {
	client, _ := service.NewClient(&service.Config{Token: "wrongToken", URL: testutils.TestURLProtocol + "://" + testutils.TestSSCHost, TenantID: testutils.TestTenantID, Timeout: time.Second * 5})
	err := client.IngestService.CreateEvent(model.Event{Event: "failed test"})
	assert.NotEmpty(t, err)
	assert.Equal(t, 401, err.(*util.HTTPError).Status)
	assert.Equal(t, "401 Unauthorized", err.(*util.HTTPError).Message)
}

func TestCreateEvents(t *testing.T) {
	event1 := model.Event{Host: "host1", Event: "test1"}
	event2 := model.Event{Host: "host2", Event: "test2"}
	err := getClient(t).IngestService.CreateEvents([]model.Event{event1, event2})
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
	err := client.IngestService.CreateMetricEvent(metricEvent1)
	assert.Empty(t, err)
}
