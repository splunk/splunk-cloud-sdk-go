// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/service"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var wg sync.WaitGroup

// Should flush when ticker ticked and queue is not full
func TestBatchEventsSenderTickerFlush(t *testing.T) {
	var client = getClient(t)

	event1 := model.Event{Host: "host1", Body: "test1"}
	event2 := model.Event{Host: "host2", Body: "test2"}
	event3 := model.Event{Host: "host3", Body: "test3"}
	done := make(chan bool, 1)

	collector, err := client.IngestService.NewBatchEventsSender(5, 1000, 0)
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

	event1 := model.Event{Host: "host1", Body: "test1"}
	event2 := model.Event{Host: "host2", Body: "test2"}
	event3 := model.Event{Host: "host3", Body: "test3"}
	done := make(chan bool, 1)

	collector, err := client.IngestService.NewBatchEventsSender(5, 1000, 0)
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

// Should flush when queue is full , payLoadSize limitation hit before batch size
func TestBatchEventsSenderPayloadSizeQueueFlush(t *testing.T) {
	var client = getClient(t)

	event1 := model.Event{Host: "host1", Body: "first host"}
	event2 := model.Event{Host: "host2", Body: "this is the second host"}
	event3 := model.Event{Host: "host3", Body: "third host"}
	event4 := model.Event{Host: "host4", Body: "fourth host"}
	event5 := model.Event{Host: "host5", Body: "fifth host"}
	event6 := model.Event{Host: "host6", Body: "this is the sixth host but not a very long one"}

	done := make(chan bool, 1)

	collector, err := client.IngestService.NewBatchEventsSender(5, 1000, 30)
	require.Emptyf(t, err, "Error creating NewBatchEventsSender: %s", err)
	collector.Run()
	go blocking(done, 3)
	err = collector.AddEvent(event1)
	assert.Emptyf(t, err, "Error collector.AddEvent(event1): %s", err)
	err = collector.AddEvent(event2)
	assert.Emptyf(t, err, "Error collector.AddEvent(event2): %s", err)
	err = collector.AddEvent(event3)
	assert.Emptyf(t, err, "Error collector.AddEvent(event3): %s", err)
	err = collector.AddEvent(event4)
	assert.Emptyf(t, err, "Error collector.AddEvent(event4): %s", err)
	err = collector.AddEvent(event5)
	assert.Emptyf(t, err, "Error collector.AddEvent(event5): %s", err)
	err = collector.AddEvent(event6)
	assert.Emptyf(t, err, "Error collector.AddEvent(event6): %s", err)

	collector.Stop()
	<-done
	assert.Equal(t, 0, len(collector.EventsQueue))
}

// Should flush when queue is full, only payloadSize is hit, batchsize has no impact
func TestBatchEventsSenderBatchSizeNoImpact(t *testing.T) {
	var client = getClient(t)

	event1 := model.Event{Host: "host1", Body: "thisisalongstring"}
	event2 := model.Event{Host: "host2", Body: "for event"}
	event3 := model.Event{Host: "host3", Body: "thisisalongstring"}
	event4 := model.Event{Host: "host4", Body: "host4"}
	event5 := model.Event{Host: "host5", Body: "host5"}
	done := make(chan bool, 1)

	collector, err := client.IngestService.NewBatchEventsSender(10, 1000, 20)
	require.Emptyf(t, err, "Error creating NewBatchEventsSender: %s", err)
	collector.Run()
	go blocking(done, 3)
	err = collector.AddEvent(event1)
	assert.Emptyf(t, err, "Error collector.AddEvent(event1): %s", err)
	err = collector.AddEvent(event2)
	assert.Emptyf(t, err, "Error collector.AddEvent(event2): %s", err)
	err = collector.AddEvent(event3)
	assert.Emptyf(t, err, "Error collector.AddEvent(event3): %s", err)
	err = collector.AddEvent(event4)
	assert.Emptyf(t, err, "Error collector.AddEvent(event4): %s", err)
	err = collector.AddEvent(event5)
	assert.Emptyf(t, err, "Error collector.AddEvent(event5): %s", err)

	collector.Stop()
	<-done
	assert.Equal(t, 0, len(collector.EventsQueue))
}

// Should flush when queue is full, batchSize is hit before payLoadSize
func TestBatchEventsSenderPayloadSizeNoImpact(t *testing.T) {
	var client = getClient(t)

	event1 := model.Event{Host: "host1", Body: "host1host1"}
	event2 := model.Event{Host: "host2", Body: "host2host1"}
	event3 := model.Event{Host: "host3", Body: "host3host1"}
	event4 := model.Event{Host: "host4", Body: "host4host1"}
	event5 := model.Event{Host: "host5", Body: "host5host1"}
	done := make(chan bool, 1)

	collector, err := client.IngestService.NewBatchEventsSender(3, 1000, 1000)
	require.Emptyf(t, err, "Error creating NewBatchEventsSender: %s", err)
	collector.Run()
	go blocking(done, 3)
	err = collector.AddEvent(event1)
	assert.Emptyf(t, err, "Error collector.AddEvent(event1): %s", err)
	err = collector.AddEvent(event2)
	assert.Emptyf(t, err, "Error collector.AddEvent(event2): %s", err)
	err = collector.AddEvent(event3)
	assert.Emptyf(t, err, "Error collector.AddEvent(event3): %s", err)
	err = collector.AddEvent(event4)
	assert.Emptyf(t, err, "Error collector.AddEvent(event4): %s", err)
	err = collector.AddEvent(event5)
	assert.Emptyf(t, err, "Error collector.AddEvent(event5): %s", err)

	collector.Stop()
	<-done
	assert.Equal(t, 0, len(collector.EventsQueue))
}

// Should flush when Queue is full, both batchSize and payLoadSize are hit while processing the Events Queue
func TestBatchEventsSenderFlushBothBatchSizePayloadSize(t *testing.T) {
	var client = getClient(t)

	event1 := model.Event{Body: "test1", Host: "host1"}
	event2 := model.Event{Body: "test2", Host: "host2"}
	event3 := model.Event{Body: "this is a long host body so it can hit payload Size", Host: "host3"}
	event4 := model.Event{Body: "test3test4", Host: "host4"}
	event5 := model.Event{Body: "test5test6", Host: "test5"}
	done := make(chan bool, 1)

	collector, err := client.IngestService.NewBatchEventsSender(2, 1000, 20)
	require.Emptyf(t, err, "Error creating NewBatchEventsSender: %s", err)
	collector.Run()
	go blocking(done, 3)
	err = collector.AddEvent(event1)
	assert.Emptyf(t, err, "Error collector.AddEvent(event1): %s", err)
	err = collector.AddEvent(event2)
	assert.Emptyf(t, err, "Error collector.AddEvent(event2): %s", err)
	err = collector.AddEvent(event3)
	assert.Emptyf(t, err, "Error collector.AddEvent(event3): %s", err)
	err = collector.AddEvent(event4)
	assert.Emptyf(t, err, "Error collector.AddEvent(event4): %s", err)
	err = collector.AddEvent(event5)
	assert.Emptyf(t, err, "Error collector.AddEvent(event5): %s", err)

	collector.Stop()
	<-done
	assert.Equal(t, 0, len(collector.EventsQueue))
}

// Should flush when quit signal is sent
func TestBatchEventsSenderQuitFlush(t *testing.T) {
	var client = getClient(t)

	event1 := model.Event{Host: "host1", Body: "test1"}
	done := make(chan bool, 1)
	collector, err := client.IngestService.NewBatchEventsSender(5, 1000, 0)
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

func addEventBatch(collector *service.BatchEventsSender, event1 model.Event) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		if err := collector.AddEvent(event1); err != nil {
			fmt.Println(err)
			return
		}
	}
}

// Should return error message when 5 errors are encountered during sending batches
func TestBatchEventsSenderErrorHandle(t *testing.T) {
	var client = getInvalidClient(t)

	event1 := model.Event{Host: "host1", Body: "test10"}

	maxAllowedErr := 4

	collector, err := client.IngestService.NewBatchEventsSenderWithMaxAllowedError(2, 2000, 0, maxAllowedErr)
	require.Emptyf(t, err, "Error creating NewBatchEventsSender: %s", err)

	collector.Run()

	// start 15 threads to send data simultaneously
	wg.Add(8)
	for i := 0; i < 8; i++ {
		go addEventBatch(collector, event1)
	}
	wg.Wait()

	// it is possible that the stop signal is set by the maxAllowedErr constraint,
	// but while there are some events are pushed to the queue by some threads before we do last flush
	// therefore the last flush that flush all content in queue will add more errors than maxAllowedErr
	assert.True(t, len(collector.Errors) >= maxAllowedErr)

	httpError, ok := collector.Errors[0].Error.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, httpError.HTTPStatusCode, 401)
	assert.Equal(t, httpError.HTTPStatus, "401 Unauthorized")
	errFound := httpError.Message == "Error validating request" || httpError.Message == "Invalid or Expired Bearer Token"
	assert.True(t, errFound)

	assert.Equal(t, "host1", collector.Errors[0].Events[0].Host)
	collector.Stop()
}

func TestBatchEventsSenderErrorHandleWithCallBack(t *testing.T) {
	var client = getInvalidClient(t)

	event1 := model.Event{Host: "host1", Body: "test10"}

	maxAllowedErr := 5

	collector, err := client.IngestService.NewBatchEventsSenderWithMaxAllowedError(2, 2000, 0, maxAllowedErr)
	require.Emptyf(t, err, "Error creating NewBatchEventsSender: %s", err)

	callbackPrint := ""
	callback := func(b *service.BatchEventsSender) {
		assert.True(t, len(b.Errors) > 0)
		callbackPrint = "call from callback function"
	}

	collector.SetCallbackHandler(callback)

	assert.Equal(t, "", callbackPrint)

	// this should call the callback func when err happens during sending the batch and update the value of callbackPrint
	collector.Run()
	wg.Add(1)
	go addEventBatch(collector, event1)
	wg.Wait()

	// this wait is to make the sure the callback func finish its execution
	done := make(chan bool, 1)
	go blocking(done, 2)
	<-done

	assert.Equal(t, "call from callback function", callbackPrint)
	collector.Stop()
}

func TestBatchEventsSenderRestart(t *testing.T) {
	var client = getInvalidClient(t)
	event1 := model.Event{Host: "host1", Body: "test10"}

	maxAllowedErr := 4

	collector, err := client.IngestService.NewBatchEventsSenderWithMaxAllowedError(2, 2000, 0, maxAllowedErr)
	require.Emptyf(t, err, "Error creating NewBatchEventsSender: %s", err)

	// Initial run of the batchSender
	collector.Run()
	// start 15 threads to send data simultaneously
	wg.Add(8)
	for i := 0; i < 8; i++ {
		go addEventBatch(collector, event1)
	}
	wg.Wait()

	// batchSender should have stopped due to maxError hit
	assert.False(t, collector.IsRunning)
	assert.True(t, len(collector.Errors) >= 4)

	// restart the batchSender and resend events, everything should work just like the initial run
	collector.Restart()
	assert.Nil(t, collector.Errors)
	assert.True(t, collector.IsRunning)

	// start 15 threads to send data simultaneously
	wg.Add(8)
	for i := 0; i < 8; i++ {
		go addEventBatch(collector, event1)
	}
	wg.Wait()
	assert.True(t, len(collector.Errors) >= 4)
}
