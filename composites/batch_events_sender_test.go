package composites

import (
	"testing"
	"time"

	"github.com/splunk/ssc-client-go/model"
	//. "github.com/splunk/ssc-client-go/util"
	"github.com/stretchr/testify/assert"
)

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
	<-done
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
	<-done
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
	assert.Equal(t, 0, len(collector.EventsQueue))
	<-done
}

// This function is purely for blocking purpose so that BatchEventsSender can run for a little while
func blocking(done chan bool, seconds int64) {
	time.Sleep(time.Duration(seconds) * time.Second)
	done <- true
}
