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
	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/util"
)

func TestIntegrationCreateEventSuccess(t *testing.T) {
	timeValue := int64(time.Now().Unix() * 1000) // Unix millis
	// client := getClient(t)
	client, _ := service.NewClient(&service.Config{Token:"eyJraWQiOiJIR1RMbXJGUWNsSGVTRUZzdTIzQ1k4cTZ3S3pjR2JwUGtvT014R2hQVVBVIiwiYWxnIjoiUlMyNTYifQ.eyJ2ZXIiOjEsImp0aSI6IkFULlJYdkI0dlNXWEQwMEtObVRtOG5JT2VOWXVibl9EOGhVbEFzSEgxMmh3STQiLCJpc3MiOiJodHRwczovL3NwbHVuay1jaWFtLm9rdGEuY29tL29hdXRoMi9kZWZhdWx0IiwiYXVkIjoiYXBpOi8vZGVmYXVsdCIsImlhdCI6MTUzNTY2OTcyMiwiZXhwIjoxNTM1NzEyOTIyLCJjaWQiOiIwb2FwYmcyem1MYW1wV2daNDJwNiIsInVpZCI6IjAwdXpsMHdlZFdxM2tvWEFDMnA2Iiwic2NwIjpbInByb2ZpbGUiLCJlbWFpbCIsIm9wZW5pZCJdLCJzdWIiOiJ4Y2hlbmdAc3BsdW5rLmNvbSJ9.TLvTMjp1s6ADz_-rKCUpskMz9g2WiKVSiRYosi9hRk-nSPHY9Ce3jzVoqdyCD54RUufW3DpeBZwF9A4LyrOW11wxOuUQyvlGXFbtJjbA0uNNUzhONsXzkOrStzwCdEw_4bwwKRdCd52wawyaTQJsKJ9NWurEy2lDDf0r8OzH9fdO6pU2aWxxjCDt2Z6ZE9vlc7NhyPs-_MUcBvdPQxb8Nl377JwoRNUNa-ysd35KMUPoj7oX77vqG-XSFQ8kV5rcAor3ume2sH-gUrwkTcYodFm87oFQ4eJyPLBV3T6Eft6fLQkbQIFxkMohwTb10H5aLp0mcWCtJvD1x-KaOI1Mtw" , URL: "https://api.playground.splunkbeta.com/xcheng-workshop", TenantID: "xcheng-workshop3", Timeout: time.Second * 5})
	clientURL, err := client.GetURL()
	assert.Empty(t, err)
	attributes := make(map[string]interface{})
	attributes["testKey"] = "testValue"
	testIngestEvent := model.Event{
		Host:       clientURL.RequestURI(),
		Body:      "test",
		Sourcetype: "sourcetype:eventgen",
		Source:     "manual-events",
		Timestamp:   timeValue,
		Attributes:  attributes}

	err = client.IngestService.CreateEvent(testIngestEvent)
	assert.Empty(t, err)
}

// TODO: Deal with later
func TestIntegrationIngestEventFail(t *testing.T) {
	invalidClient := getInvalidClient(t)
	testIngestEvent := model.Event{Body: "failed test"}
	err := invalidClient.IngestService.CreateEvent(testIngestEvent)

	assert.NotEmpty(t, err)
	assert.Equal(t, 401, err.(*util.HTTPError).HTTPStatusCode)
	assert.Equal(t, "401 Unauthorized", err.(*util.HTTPError).Message)
}

func TestIntegrationCreateEvents(t *testing.T) {
	client := getClient(t)
	event1 := model.Event{Host: "host1", Body: "test1"}
	event2 := model.Event{Host: "host2", Body: "test2"}
	err := client.IngestService.CreateEvents([]model.Event{event1, event2})
	assert.Empty(t, err)
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
	err := client.IngestService.CreateMetricEvent(metricEvent1)
	assert.Empty(t, err)

	err1 := client.IngestService.CreateMetricEvents([]model.MetricEvent{metricEvent1, metricEvent1})
	assert.Empty(t, err1)
}
