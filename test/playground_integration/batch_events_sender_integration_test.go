package playgroundintegration

import (
	"fmt"
	"strings"
	"testing"
	"time"
	"sync"
	"math/rand"
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)
var wg sync.WaitGroup

// Should flush when ticker ticked and queue is not full
func TestBatchEventsSenderTickerFlush(t *testing.T) {
	var client = getClient(t)

	event1 := model.Event{Host: "host1", Event: "test1"}
	event2 := model.Event{Host: "host2", Event: "test2"}
	event3 := model.Event{Host: "host3", Event: "test3"}
	done := make(chan bool, 1)

	collector, err := client.NewBatchEventsSender(5, 1000)
	require.Emptyf(t, err, "Error creating NewBatchEventsSender: %s", err)

	collector.Run()
	go blocking(done, 2)
	err = collector.AddEvent(event1)
	assert.Emptyf(t, err, "Error collector.AddEvent(event1): %s", err)
	err = collector.AddEvent(event2)
	assert.Emptyf(t, err, "Error collector.AddEvent(event2): %s", err)
	err = collector.AddEvent(event3)
	assert.Emptyf(t, err, "Error collector.AddEvent(event3): %s", err)
	<-done
	collector.Stop()
	assert.Equal(t, 0, len(collector.EventsQueue))
}

// Should flush when queue is full and ticker has not ticked
func TestBatchEventsSenderQueueFlush(t *testing.T) {
	var client = getClient(t)

	event1 := model.Event{Host: "host1", Event: "test1"}
	event2 := model.Event{Host: "host2", Event: "test2"}
	event3 := model.Event{Host: "host3", Event: "test3"}
	done := make(chan bool, 1)

	collector, err := client.NewBatchEventsSender(5, 1000)
	require.Emptyf(t, err, "Error creating NewBatchEventsSender: %s", err)
	collector.Run()
	go blocking(done, 2)
	err = collector.AddEvent(event1)
	assert.Emptyf(t, err, "Error collector.AddEvent(event1): %s", err)
	err = collector.AddEvent(event2)
	assert.Emptyf(t, err, "Error collector.AddEvent(event2): %s", err)
	err = collector.AddEvent(event3)
	assert.Emptyf(t, err, "Error collector.AddEvent(event3): %s", err)
	collector.Stop()
	<-done
	assert.Equal(t, 0, len(collector.EventsQueue))
}

//Should flush when quit signal is sent
func TestBatchEventsSenderQuitFlush(t *testing.T) {
	var client = getClient(t)

	event1 := model.Event{Host: "host1", Event: "test1"}
	done := make(chan bool, 1)
	collector, err := client.NewBatchEventsSender(5, 1000)
	require.Emptyf(t, err, "Error creating NewBatchEventsSender: %s", err)
	collector.Run()
	go blocking(done, 3)
	err = collector.AddEvent(event1)
	assert.Emptyf(t, err, "Error collector.AddEvent(event1): %s", err)
	collector.Stop()
	assert.Equal(t, 0, len(collector.EventsQueue))
	<-done
}

// This function is purely for blocking purpose so that BatchEventsSender can run for a little while
func blocking(done chan bool, seconds int64) {
	time.Sleep(time.Duration(seconds) * time.Second)
	done <- true
}

func addEventBatch(collector *service.BatchEventsSender, event1 model.Event) error {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		if err := collector.AddEvent(event1); err != nil {
			return err
		}
	}
	return nil
}

// Should return error message when 5 errors are encountered during sending batches
func TestBatchEventsSenderErrorHandle(t *testing.T) {
	var client= getInvalidClient(t)

	event1 := model.Event{Host: "host1", Event: "test10"}

	maxAllowedErr := 5
	collector, _ := client.NewBatchEventsSenderWithMaxAllowedError(2, 2000, maxAllowedErr)
	collector.Run()

	// start 15 threads to send data simultaneously
	wg.Add(15)
	for i := 0; i < 15; i++ {
		go addEventBatch(collector, event1)

		collector, err := client.NewBatchEventsSenderWithMaxAllowedError(2, 1000, 10)
		require.Emptyf(t, err, "Error creating NewBatchEventsSender: %s", err)
		collector.Run()

		for i := 0; i < 10; i++ {
			go addEventBatch(collector, event1)
		}
		wg.Wait()

		s := strings.Split(collector.ErrorMsg, "],")
		fmt.Println(s)

		// it is possible that the stop signal is set by the maxAllowedErr constraint,
		// but while there are some events are pushed to the queue by some threads before we do last flush
		// therefore the last flush that flush all content in queue will add more errors than maxAllowedErr
		assert.True(t, len(s)-1 >= maxAllowedErr)
		//assert.Equal(t, len(s)-1 , maxAllowedErr)

		assert.True(t, strings.Contains(s[0], "[Failed to send all events for batch: [{host1    <nil> test10 map[]}"))
		assert.True(t, strings.Contains(s[0], "\n\tError: Http Error: [401] 401 Unauthorized {\"reason\":\"Error validating request\"}"))
	}
}