// +build !integration

package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
)

func TestIntegrationCreateEventSuccess(t *testing.T) {
	timeValue := float64(1523637597)
	client := getSplunkClientForPlaygroundTests()
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

func TestIntegrationHecEventFail(t *testing.T) {
	client := NewClient(tenantID, "wrongToken", hostID, util.TestTimeOut)
	testHecEvent := model.HecEvent{Event: "failed test"}
	err := client.HecService.CreateEvent(testHecEvent)

	assert.NotEmpty(t, err)
	assert.Equal(t, 401, err.(*util.HTTPError).Status)
	assert.Equal(t, "401 Unauthorized", err.(*util.HTTPError).Message)
}

func TestIntegrationCreateRawEventSuccess(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	testHecEvent := model.HecEvent{Event: "test"}

	err := client.HecService.CreateRawEvent(testHecEvent)
	assert.Empty(t, err)
}

func TestIntegrationCreateEvents(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	event1 := model.HecEvent{Host: "host1", Event: "test1"}
	event2 := model.HecEvent{Host: "host2", Event: "test2"}
	err := client.HecService.CreateEvents([]model.HecEvent{event1, event2})
	assert.Empty(t, err)
}

func TestIntegrationBuildMultiEventsPayload(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	event1 := model.HecEvent{Host: "host1", Event: "test1"}
	event2 := model.HecEvent{Host: "host2", Event: "test2"}
	event3WithEmptyFields := model.HecEvent{Host: "", Event: "test3"}
	payload1, err := client.HecService.buildMultiEventsPayload([]model.HecEvent{event1, event2})
	assert.Nil(t, err)
	assert.Equal(t, `{"host":"host1","event":"test1"}{"host":"host2","event":"test2"}`, string(payload1[:]))
	payload2, err := client.HecService.buildMultiEventsPayload([]model.HecEvent{event1, event3WithEmptyFields})
	assert.Nil(t, err)
	assert.Equal(t, `{"host":"host1","event":"test1"}{"event":"test3"}`, string(payload2[:]))
}
