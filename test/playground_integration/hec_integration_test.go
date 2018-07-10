package playgroundintegration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
)

func TestIntegrationCreateEventSuccess(t *testing.T) {
	timeValue := float64(1523637597)
	client := getClient(t)
	testIngestEvent := model.Event{
		Host:       client.URL.RequestURI(),
		Index:      "main",
		Event:      "test",
		Sourcetype: "sourcetype:eventgen",
		Source:     "manual-events",
		Time:       &timeValue,
		Fields:     map[string]string{"testKey": "testValue"}}

	err := client.IngestService.CreateEvent(testIngestEvent)
	assert.Empty(t, err)
}

// TODO: Deal with later
func TestIntegrationIngestEventFail(t *testing.T) {
	invalidClient := getInvalidClient(t)
	testIngestEvent := model.Event{Event: "failed test"}
	err := invalidClient.IngestService.CreateEvent(testIngestEvent)

	assert.NotEmpty(t, err)
	assert.Equal(t, 401, err.(*util.HTTPError).Status)
	assert.Equal(t, "401 Unauthorized", err.(*util.HTTPError).Message)
}

func TestIntegrationCreateRawEventSuccess(t *testing.T) {
	client := getClient(t)
	testEvent := model.Event{Event: "test"}

	err := client.IngestService.CreateRawEvent(testEvent)
	assert.Empty(t, err)
}

func TestIntegrationCreateEvents(t *testing.T) {
	client := getClient(t)
	event1 := model.Event{Host: "host1", Event: "test1"}
	event2 := model.Event{Host: "host2", Event: "test2"}
	err := client.IngestService.CreateEvents([]model.Event{event1, event2})
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

	err1 := client.IngestService.CreateMetricEvents([]model.MetricEvent{metricEvent1, metricEvent1})
	assert.Empty(t, err1)
}
