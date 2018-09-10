// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package service

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"encoding/json"
	"github.com/splunk/splunk-cloud-sdk-go/model"
)

//UserErrHandler defines the type of user callback function for batchEventSender
type UserErrHandler func(*BatchEventsSender)

const errMsgSplitter = "[insErrSplit]"

// BatchEventsSender sends events in batches or periodically if batch is not full to Splunk Cloud ingest service endpoints
type BatchEventsSender struct {
	BatchSize      int
	EventsChan     chan model.Event
	EventsQueue    []model.Event
	QuitChan       chan struct{}
	EventService   *IngestService
	IngestTicker   *model.Ticker
	WaitGroup      *sync.WaitGroup
	ErrorChan      chan string
	errorMsg       string
	IsRunning      bool
	mux            sync.Mutex
	callbackFunc   UserErrHandler
	stopMux        sync.Mutex
	chanWaitMillis int
	resetMux       sync.Mutex
}

// SetCallbackHandler allows users to pass their own callback function
func (b *BatchEventsSender) SetCallbackHandler(callback UserErrHandler) {
	b.callbackFunc = callback
}

// Run sets up ticker and starts a new goroutine
func (b *BatchEventsSender) Run() {
	go b.loop()
	b.IsRunning = true
}

// loop is a infinity loop which listens to three channels
// the loop will break only if there's signal in QuitChan
// otherwise it'll constantly checking conditions for ticker and events
func (b *BatchEventsSender) loop() {
	errorMsgCount := 0

	defer close(b.EventsChan)
	for {
		select {
		case err := <-b.ErrorChan:
			errorMsgCount++
			b.errorMsg += err + errMsgSplitter

			if errorMsgCount >= cap(b.ErrorChan) {
				b.Stop()
			}

			if b.callbackFunc != nil {
				go b.callbackFunc(b)
			}

		case <-b.QuitChan:
			b.IsRunning = false
			return

		case <-b.IngestTicker.GetChan():
			b.WaitGroup.Add(1)
			go b.flush(0)

		case event := <-b.EventsChan:
			b.EventsQueue = append(b.EventsQueue, event)
			if len(b.EventsQueue) >= b.BatchSize {
				b.WaitGroup.Add(1)
				go b.flush(1)
			}
		}
	}
}

// Stop sends a signal to QuitChan, wait for all registered goroutines to finish, stop ticker and clear queue
func (b *BatchEventsSender) Stop() {
	b.stopMux.Lock()
	defer b.stopMux.Unlock()

	if b.IsRunning == false && len(b.EventsQueue) == 0 {
		return
	}

	b.IsRunning = false
	// Wait until no element is in channel
	for {
		if len(b.EventsChan) == 0 {
			break
		}
	}

	b.IngestTicker.Stop()

	b.WaitGroup.Add(1)
	// flush one last time before stop
	go b.flush(2)
	b.WaitGroup.Wait()
	b.QuitChan <- struct{}{}
}

// AddEvent pushes a single event into EventsChan
func (b *BatchEventsSender) AddEvent(event model.Event) error {
	if !b.IsRunning {
		return errors.New("Need to start the BatchEventsSender first, call Run() ")
	}

	// Intend to only start ticker when first event is received.
	if len(b.EventsQueue) == 0 && len(b.EventsChan) == 0 && b.IngestTicker.IsRunning() == false {
		b.IngestTicker.Start()
	}

	for len(b.EventsChan) >= cap(b.EventsChan) {
		time.Sleep(time.Duration(b.chanWaitMillis) * time.Millisecond)
	}

	if b.IsRunning {
		b.EventsChan <- event
	}
	return nil
}

// flush sends off all events currently in EventsQueue and resets ticker afterwards
// If EventsQueue size is bigger than BatchSize, it'll slice the queue into batches and send batches one by one
func (b *BatchEventsSender) flush(flushSource int) {
	defer b.WaitGroup.Done()
	defer b.mux.Unlock()

	b.mux.Lock()
	// Reset ticker
	if flushSource == 0 {
		b.IngestTicker.Reset()
	} else if flushSource == 1 && len(b.EventsQueue) < b.BatchSize {
		// it is possible different threads send flush signal while the previous flush already flush everything in queue
		return
	}

	events := append([]model.Event(nil), b.EventsQueue...)
	b.ResetQueue()

	// slice events into batch size to send
	b.sendEventInBatches(events)

}

// sendEventInBatches slices events into batch size to send
func (b *BatchEventsSender) sendEventInBatches(events []model.Event) {
	if len(events) <= 0 {
		return
	}
	for i := 0; i < len(events); {
		end := len(events)
		if i+b.BatchSize < len(events) {
			end = i + b.BatchSize
		}
		batchedEvents := events[i:end]
		err := b.EventService.PostEvents(batchedEvents)
		i = i + b.BatchSize
		if err != nil {
			var errString string
			jsonEvents, parsingErr := json.Marshal(events)
			if parsingErr != nil {
				errString = fmt.Sprintf("failed to send all events:%v\nEventPayload:%v", err, events)
			} else {
				errString = fmt.Sprintf("failed to send all events:%v\nEventPayload:%v", err, string(jsonEvents))
			}
			for len(b.EventsChan) >= cap(b.EventsChan) {
				time.Sleep(time.Duration(b.chanWaitMillis) * time.Millisecond)
			}
			if b.IsRunning {
				b.ErrorChan <- errString
			}
		}
	}
}

// ResetQueue sets b.EventsQueue to empty, but keep memory allocated for underlying array
func (b *BatchEventsSender) ResetQueue() {
	b.EventsQueue = b.EventsQueue[:0]
}

// Restart will reset batch event sender, clear up queue, error msg, timer etc.
func (b *BatchEventsSender) Restart() {
	defer b.resetMux.Unlock()
	b.resetMux.Lock()

	if b.IsRunning {
		b.Stop()
	}

	// reopen channels
	b.EventsChan = make(chan model.Event, b.BatchSize)
	b.ErrorChan = make(chan string, cap(b.ErrorChan))

	b.ResetQueue()
	b.errorMsg = ""
	b.Run()
}

// GetErrors return all the error messages as an array
func (b *BatchEventsSender) GetErrors() []string {
	if b.errorMsg == "" {
		return nil
	}

	errors := strings.Split(b.errorMsg, errMsgSplitter)
	return errors[:len(errors)-1]
}
