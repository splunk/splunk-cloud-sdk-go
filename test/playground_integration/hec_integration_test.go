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
	testHecEvent := model.HecEvent{
		Host:       client.URL.RequestURI(),
		Index:      "main",
		Event:      "test",
		Sourcetype: "sourcetype:eventgen",
		Source:     "manual-events",
		Time:       &timeValue,
		Fields:     map[string]string{"testKey": "testValue"}}

	err := client.HecService.CreateEvent(testHecEvent)
	assert.Empty(t, err)
}

// TODO: Deal with later
func TestIntegrationHecEventFail(t *testing.T) {
	invalidClient := getInvalidClient(t)
	testHecEvent := model.HecEvent{Event: "failed test"}
	err := invalidClient.HecService.CreateEvent(testHecEvent)

	assert.NotEmpty(t, err)
	assert.Equal(t, 401, err.(*util.HTTPError).Status)
	assert.Equal(t, "401 Unauthorized", err.(*util.HTTPError).Message)
}

func TestIntegrationCreateRawEventSuccess(t *testing.T) {
	client := getClient(t)
	testHecEvent := model.HecEvent{Event: "test"}

	err := client.HecService.CreateRawEvent(testHecEvent)
	assert.Empty(t, err)
}

func TestIntegrationCreateEvents(t *testing.T) {
	client := getClient(t)
	event1 := model.HecEvent{Host: "host1", Event: "test1"}
	event2 := model.HecEvent{Host: "host2", Event: "test2"}
	err := client.HecService.CreateEvents([]model.HecEvent{event1, event2})
	assert.Empty(t, err)
}
