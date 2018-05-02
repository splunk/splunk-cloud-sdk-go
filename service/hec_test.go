package service

import (
	"testing"
	"time"

	"github.com/splunk/ssc-client-go/model"
	. "github.com/splunk/ssc-client-go/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateEventSuccess(t *testing.T) {
	timeValue := float64(1523637597)
	err := getSplunkClient().HecService.CreateEvent(
		model.HecEvent{Host: "http://ssc-sdk-shared-stubby:8882", Index: "main", Event: "test", Sourcetype: "sourcetype:eventgen", Source: "manual-events", Time: &timeValue, Fields: map[string]string{"testKey": "testValue"}})
	assert.Empty(t, err)
}

func TestCreateRawEventSuccess(t *testing.T) {
	err := getSplunkClient().HecService.CreateRawEvent(
		model.HecEvent{Event: "test"})
	assert.Empty(t, err)
}

func TestHecEventFail(t *testing.T) {
	client := NewClient(TestTenantID, "wrongToken", TestStubbySchme+"://"+TestStubbyHost, time.Second*5)
	err := client.HecService.CreateEvent(
		model.HecEvent{Event: "failed test"})
	assert.NotEmpty(t, err)
	assert.Equal(t, 401, err.(*HTTPError).Status)
	assert.Equal(t, "401 Unauthorized", err.(*HTTPError).Message)
}

func TestCreateEvents(t *testing.T) {
	event1 := model.HecEvent{Host: "host1", Event: "test1"}
	event2 := model.HecEvent{Host: "host2", Event: "test2"}
	err := getSplunkClient().HecService.CreateEvents([]model.HecEvent{event1, event2})
	assert.Empty(t, err)
}

func TestBuildMultiEventsPayload(t *testing.T) {
	event1 := model.HecEvent{Host: "host1", Event: "test1"}
	event2 := model.HecEvent{Host: "host2", Event: "test2"}
	event3WithEmptyFields := model.HecEvent{Host: "", Event: "test3"}
	payload1, err := getSplunkClient().HecService.buildMultiEventsPayload([]model.HecEvent{event1, event2})
	assert.Nil(t, err)
	assert.Equal(t, `{"host":"host1","event":"test1"}{"host":"host2","event":"test2"}`, string(payload1[:]))
	payload2, err := getSplunkClient().HecService.buildMultiEventsPayload([]model.HecEvent{event1, event3WithEmptyFields})
	assert.Nil(t, err)
	assert.Equal(t, `{"host":"host1","event":"test1"}{"event":"test3"}`, string(payload2[:]))
}

func TestHecService_NewBatchEventsCollector(t *testing.T) {
	event1 := model.HecEvent{Host: "host1", Event: "test1"}
	event2 := model.HecEvent{Host: "host2", Event: "test2"}
	event3 := model.HecEvent{Host: "host3", Event: "test2"}
	done := make(chan bool, 1)
	collector := getSplunkClient().HecService.NewBatchEventsCollector(2, 1000)
	collector.Start()
	go blocking(done)
	collector.AddEvent(event1)
	collector.AddEvent(event2)
	collector.AddEvent(event3)
	<- done
	collector.Stop()
}

func blocking(done chan bool) {
	time.Sleep(time.Duration(5)*time.Second)
	done <- true
}
