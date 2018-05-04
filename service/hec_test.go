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

func TestHecServiceNewBatchEventsSenderSuccess(t *testing.T) {
	collector, err := getSplunkClient().HecService.NewBatchEventsSender(5, 1000)
	assert.Nil(t, err)
	// Initial queue should
	assert.Equal(t, 0, len(collector.EventsQueue))
	assert.Equal(t, 5, cap(collector.EventsQueue))
	assert.Equal(t, 0, len(collector.EventsChan))
	assert.Equal(t, 5, cap(collector.EventsChan))
	assert.Equal(t, 0, len(collector.QuitChan))
	assert.Equal(t, 1, cap(collector.QuitChan))
	assert.Equal(t, 5, collector.BatchSize)
	assert.Equal(t, time.Duration(1000), collector.Interval)

}

func TestHecServiceNewBatchEventsCollectorFail(t *testing.T) {
	_, err := getSplunkClient().HecService.NewBatchEventsSender(0, 0)
	assert.EqualError(t, err, "batchSize and interval cannot be 0")
}

// Should flush when ticker ticked and queue is not full
func TestHecServiceNewBatchEventsCollectorTickerFlush(t *testing.T) {
	event1 := model.HecEvent{Host: "host1", Event: "test1"}
	event2 := model.HecEvent{Host: "host2", Event: "test2"}
	event3 := model.HecEvent{Host: "host3", Event: "test3"}
	done := make(chan bool, 1)
	collector, _ := getSplunkClient().HecService.NewBatchEventsSender(5, 1000)
	collector.Run()
	go blocking(done, 2)
	collector.AddEvent(event1)
	collector.AddEvent(event2)
	collector.AddEvent(event3)
	<- done
	collector.Stop()
	assert.Equal(t, 0, len(collector.EventsQueue))
}

// Should flush when queue is full and ticker has not ticked
func TestHecServiceNewBatchEventsCollectorQueueFlush(t *testing.T) {
	event1 := model.HecEvent{Host: "host1", Event: "test1"}
	event2 := model.HecEvent{Host: "host2", Event: "test2"}
	event3 := model.HecEvent{Host: "host3", Event: "test3"}
	done := make(chan bool, 1)
	collector, _ := getSplunkClient().HecService.NewBatchEventsSender(2, 100000)
	collector.Run()
	go blocking(done, 2)
	collector.AddEvent(event1)
	collector.AddEvent(event2)
	collector.AddEvent(event3)
	collector.Stop()
	<- done
	assert.Equal(t, 0, len(collector.EventsQueue))
}

// Should flush when quit signal is sent
func TestHecServiceNewBatchEventsCollectorQuitFlush(t *testing.T) {
	event1 := model.HecEvent{Host: "host1", Event: "test1"}
	done := make(chan bool, 1)
	collector, _ := getSplunkClient().HecService.NewBatchEventsSender(5, 10000)
	collector.Run()
	go blocking(done, 3)
	collector.AddEvent(event1)
	collector.Stop()
	<- done
	assert.Equal(t, 0, len(collector.EventsQueue))
}

// This function is purely for blocking purpose so that BatchEventsSender can run for a little while
func blocking(done chan bool, seconds int64) {
	time.Sleep(time.Duration(seconds)*time.Second)
	done <- true
}
