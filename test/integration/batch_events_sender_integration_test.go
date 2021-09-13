/*
 * Copyright 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package integration

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/services/ingest"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var wg sync.WaitGroup

// Should flush when ticker ticked and queue is not full
func TestBatchEventsSenderTickerFlush(t *testing.T) {
	var client = getClient(t)

	host := "host1"
	host1 := "host2"
	host2 := "host3"

	body := make(map[string]interface{})
	body["event"] = "test1"
	body1 := make(map[string]interface{})
	body1["event"] = "test2"
	body2 := make(map[string]interface{})
	body2["event"] = "test3"
	event1 := ingest.Event{Host: &host, Body: body}
	event2 := ingest.Event{Host: &host1, Body: body1}
	event3 := ingest.Event{Host: &host2, Body: body2}
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

	host := "host1"
	host1 := "host2"
	host2 := "host3"

	body := make(map[string]interface{})
	body["event"] = "test1"
	body1 := make(map[string]interface{})
	body1["event"] = "test2"
	body2 := make(map[string]interface{})
	body2["event"] = "test3"

	event1 := ingest.Event{Host: &host, Body: body}
	event2 := ingest.Event{Host: &host1, Body: body1}
	event3 := ingest.Event{Host: &host2, Body: body2}
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

// Should flush when queue is full, payLoadSize limitation hit before batch size
// First flush 2 events because we hit the payloadsize, then the next 3 due to the same reason and then flush the final one
func TestBatchEventsSenderPayloadSizeQueueFlush(t *testing.T) {
	var client = getClient(t)

	host1 := "host1"
	host2 := "host2"
	host3 := "host3"
	host4 := "host4"
	host5 := "host5"
	host6 := "host6"

	body1 := make(map[string]interface{})
	body1["event"] = "first host"
	body2 := make(map[string]interface{})
	body2["event"] = "this is the second host"
	body3 := make(map[string]interface{})
	body3["event"] = "third host"
	body4 := make(map[string]interface{})
	body4["event"] = "fourth host"
	body5 := make(map[string]interface{})
	body5["event"] = "fifth host"
	body6 := make(map[string]interface{})
	body6["event"] = "this is the sixth host"

	event1 := ingest.Event{Host: &host1, Body: body1}
	event2 := ingest.Event{Host: &host2, Body: body2}
	event3 := ingest.Event{Host: &host3, Body: body3}
	event4 := ingest.Event{Host: &host4, Body: body4}
	event5 := ingest.Event{Host: &host5, Body: body5}
	event6 := ingest.Event{Host: &host6, Body: body6}

	done := make(chan bool, 1)

	collector, err := client.IngestService.NewBatchEventsSender(5, 1000, 30)
	require.Emptyf(t, err, "Error creating NewBatchEventsSender: %s", err)
	collector.Run()
	go blocking(done, 10)
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

	<-done
	collector.Stop()
	assert.Equal(t, 0, len(collector.EventsQueue))
}

// Should flush when queue is full, only payloadSize is hit, batchsize has no impact
func TestBatchEventsSenderBatchSizeNoImpact(t *testing.T) {
	var client = getClient(t)

	host := "host1"
	host1 := "host2"
	host2 := "host3"
	host3 := "host4"
	host4 := "host5"

	body := make(map[string]interface{})
	body["event"] = "thisisalongstring"
	body1 := make(map[string]interface{})
	body1["event"] = "for event"
	body2 := make(map[string]interface{})
	body2["event"] = "thisisalongstring"
	body3 := make(map[string]interface{})
	body3["event"] = "host4"
	body4 := make(map[string]interface{})
	body4["event"] = "host5"

	event1 := ingest.Event{Host: &host, Body: body}
	event2 := ingest.Event{Host: &host1, Body: body1}
	event3 := ingest.Event{Host: &host2, Body: body2}
	event4 := ingest.Event{Host: &host3, Body: body3}
	event5 := ingest.Event{Host: &host4, Body: body4}
	done := make(chan bool, 1)

	collector, err := client.IngestService.NewBatchEventsSender(100, 10000, 20)
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

	//this sleep should not exceed the interval, otherwise interval will force the flush
	time.Sleep(5 * time.Millisecond)
	collector.Stop()
	<-done
	assert.Equal(t, 0, len(collector.EventsQueue))
}

// Should flush when queue is full, batchSize is hit before payLoadSize
func TestBatchEventsSenderPayloadSizeNoImpact(t *testing.T) {
	t.Skip("Skip TestBatchEventsSenderPayloadSizeNoImpact until len(collector.EventsQueue) consistenly = 0 (flaky)")
	var client = getClient(t)

	host := "host1"
	host1 := "host2"
	host2 := "host3"
	host3 := "host4"
	host4 := "host5"

	body := make(map[string]interface{})
	body["event"] = "host1host1"
	body1 := make(map[string]interface{})
	body1["event"] = "host2host1"
	body2 := make(map[string]interface{})
	body2["event"] = "host3host1"
	body3 := make(map[string]interface{})
	body3["event"] = "host3host1"
	body4 := make(map[string]interface{})
	body4["event"] = "host5host1"

	event1 := ingest.Event{Host: &host, Body: body}
	event2 := ingest.Event{Host: &host1, Body: body1}
	event3 := ingest.Event{Host: &host2, Body: body2}
	event4 := ingest.Event{Host: &host3, Body: body3}
	event5 := ingest.Event{Host: &host4, Body: body4}
	done := make(chan bool, 1)

	collector, err := client.IngestService.NewBatchEventsSender(3, 1000, 2000)
	require.Emptyf(t, err, "Error creating NewBatchEventsSender: %s", err)
	collector.Run()
	go blocking(done, 5)
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

	host := "host1"
	host1 := "host2"
	host2 := "host3"
	host3 := "host4"
	host4 := "host5"

	body := make(map[string]interface{})
	body["event"] = "test1"
	body1 := make(map[string]interface{})
	body1["event"] = "test2"
	body2 := make(map[string]interface{})
	body2["event"] = "this is a long host body so it can hit payload Size"

	event1 := ingest.Event{Body: body, Host: &host}
	event2 := ingest.Event{Body: body1, Host: &host1}
	event3 := ingest.Event{Body: true, Host: &host2}
	event4 := ingest.Event{Body: 123, Host: &host3}
	event5 := ingest.Event{Body: "string-only", Host: &host4}
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
	time.Sleep(1 * time.Millisecond)
	assert.Emptyf(t, err, "Error collector.AddEvent(event4): %s", err)
	err = collector.AddEvent(event5)
	time.Sleep(1 * time.Millisecond)
	assert.Emptyf(t, err, "Error collector.AddEvent(event5): %s", err)

	collector.Stop()
	<-done
	time.Sleep(3 * time.Millisecond)
	assert.Equal(t, 0, len(collector.EventsQueue))
}

// Should flush when quit signal is sent
func TestBatchEventsSenderQuitFlush(t *testing.T) {
	var client = getClient(t)

	host := "host1"
	body := make(map[string]interface{})
	body["event"] = "test1"

	event1 := ingest.Event{Host: &host, Body: body}
	done := make(chan bool, 1)
	collector, err := client.IngestService.NewBatchEventsSender(5, 1000, 0)
	require.Emptyf(t, err, "Error creating NewBatchEventsSender: %s", err)

	collector.Run()
	go blocking(done, 3)
	err = collector.AddEvent(event1)
	assert.Emptyf(t, err, "Error collector.AddEvent(event1): %s", err)

	collector.Stop()
	time.Sleep(3 * time.Millisecond)
	assert.Equal(t, 0, len(collector.EventsQueue))
	<-done
}

// This function is purely for blocking purpose so that BatchEventsSender can run for a little while
func blocking(done chan bool, seconds int64) {
	time.Sleep(time.Duration(seconds) * time.Second)
	done <- true
}

func addEventBatch(collector *ingest.BatchEventsSender, event1 ingest.Event) {
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

	host := "host1"
	body := make(map[string]interface{})
	body["event"] = "test1"

	event1 := ingest.Event{Host: &host, Body: body}

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

	assert.Equal(t, &host, collector.Errors[0].Events[0].Host)
	collector.Stop()
}

func TestBatchEventsSenderErrorHandleWithCallBack(t *testing.T) {
	var client = getInvalidClient(t)

	host := "host1"
	body := make(map[string]interface{})
	body["event"] = "test10"

	event1 := ingest.Event{Host: &host, Body: body}

	maxAllowedErr := 5

	collector, err := client.IngestService.NewBatchEventsSenderWithMaxAllowedError(2, 2000, 0, maxAllowedErr)
	require.Emptyf(t, err, "Error creating NewBatchEventsSender: %s", err)

	callbackPrint := ""
	callback := func(b *ingest.BatchEventsSender) {
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

	host := "host1"
	body := make(map[string]interface{})
	body["event"] = "test10"

	event1 := ingest.Event{Host: &host, Body: body}

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
