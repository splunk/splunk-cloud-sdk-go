package playgroundintegration

import (
	"fmt"
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/service"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

// Should flush when ticker ticked and queue is not full
func TestBatchEventsSenderTickerFlush(t *testing.T) {
	var client = getClient(t)

	event1 := model.HecEvent{Host: "host1", Event: "test1"}
	event2 := model.HecEvent{Host: "host2", Event: "test2"}
	event3 := model.HecEvent{Host: "host3", Event: "test3"}
	done := make(chan bool, 1)

	collector, _ := client.NewBatchEventsSender(5, 1000)

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
func TestBatchEventsSenderQueueFlush(t *testing.T) {
	var client = getClient(t)

	event1 := model.HecEvent{Host: "host1", Event: "test1"}
	event2 := model.HecEvent{Host: "host2", Event: "test2"}
	event3 := model.HecEvent{Host: "host3", Event: "test3"}
	done := make(chan bool, 1)

	collector, _ := client.NewBatchEventsSender(5, 1000)
	collector.Run()
	go blocking(done, 2)
	collector.AddEvent(event1)
	collector.AddEvent(event2)
	collector.AddEvent(event3)
	collector.Stop()
	<-done
	assert.Equal(t, 0, len(collector.EventsQueue))
}

//Should flush when quit signal is sent
func TestBatchEventsSenderQuitFlush(t *testing.T) {
	var client = getClient(t)

	event1 := model.HecEvent{Host: "host1", Event: "test1"}
	done := make(chan bool, 1)
	collector, _ := client.NewBatchEventsSender(5, 1000)
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

func addEventBatch(collector *service.BatchEventsSender, event1 model.HecEvent) {
	for i := 0; i < 5; i++ {
		collector.AddEvent(event1)
	}
}

// Should return error message when 5 errors are encountered during sending batches
func TestBatchEventsSenderErrorHandle(t *testing.T) {
	var client = getInvalidClient(t)

	event1 := model.HecEvent{Host: "host1", Event: "test10"}
	done := make(chan bool, 1)

	collector, _ := client.NewBatchEventsSenderWithMaxAllowedError(2, 1000, 10)
	collector.Run()
	go blocking(done, 15)
	for i := 0; i < 10; i++ {
		go addEventBatch(collector, event1)
	}

	<-done

	s := strings.Split(collector.ErrorMsg, "],")
	fmt.Println(s)
	assert.Equal(t, 10, len(s)-1)
}

//// TestBatchEventsSenderFlush should send events right away
//func TestBatchEventsSenderFlush(t *testing.T) {
//	var client = getClient(t)
//
//	event1 := model.HecEvent{Host: "host1", Event: "test1"}
//	event2 := model.HecEvent{Host: "host2", Event: "test2"}
//	event3 := model.HecEvent{Host: "host3", Event: "test3"}
//
//	collector, _ := client.NewBatchEventsSender(5, 1000)
//	collector.Run()
//	collector.AddEvent(event1)
//	collector.AddEvent(event2)
//	collector.AddEvent(event3)
//	collector.Flush()
//	collector.Stop()
//
//	assert.Equal(t, 0, len(collector.EventsQueue))
//	assert.Empty(t, collector.ErrorMsg)
//}