// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package ingest

import (
	"errors"
	"sync"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/util"
)

//UserErrHandler defines the type of user callback function for batchEventSender
type UserErrHandler func(*BatchEventsSender)

//ingestError defines the type of the event payload sent and ingest error incurred
type ingestError struct {
	Error  error
	Events []Event
}

//Default Payload Size in unit Bytes
const (
	payLoadSize = 1040000 // ~1MiB 1048576 bytes
)

// BatchEventsSender sends events in batches or periodically if batch is not full to Splunk Cloud ingest service endpoints
type BatchEventsSender struct {
	PayLoadBytes   int
	BatchSize      int
	EventsChan     chan Event
	EventsQueue    []Event
	QuitChan       chan struct{}
	EventService   *Service
	IngestTicker   *util.Ticker
	WaitGroup      *sync.WaitGroup
	ErrorChan      chan struct{}
	IsRunning      bool
	mux            sync.Mutex
	callbackFunc   UserErrHandler
	stopMux        sync.Mutex
	chanWaitMillis int
	resetMux       sync.Mutex
	Errors         []ingestError
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
	batchPayLoadSize := 0

	defer close(b.EventsChan)
	for {
		select {
		case <-b.ErrorChan:
			errorMsgCount++

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

			for i := 0; i < len(b.EventsQueue); i++ {
				batchPayLoadSize += len(b.EventsQueue[i].Body.(string))
			}
			if len(b.EventsQueue) >= b.BatchSize || batchPayLoadSize >= b.PayLoadBytes {
				b.WaitGroup.Add(1)
				go b.flush(1)
				batchPayLoadSize = 0
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
func (b *BatchEventsSender) AddEvent(event Event) error {
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

	events := append([]Event(nil), b.EventsQueue...)
	b.ResetQueue()

	// slice events into batch size to send
	b.sendEventInBatches(events)

}

//sendEventInBatches will slice Event Queue into batches.
//Add events from event queue into a batch until either the batch events counts size is reached or the payload size limit is hit
//Once the batch is flushed, another batch is initialized with the remaining elements from events queue until either of the two limits are reached
func (b *BatchEventsSender) sendEventInBatches(events []Event) {
	if len(events) <= 0 {
		return
	}
	end := 0
	for i := 0; i < len(events); {
		batchedEvents := events[i : i+1]
		batchPayLoadSize := len(events[i].Body.(string))

		end = i + 1
		//Increment batch until Payload Size limit is reached or batch events count is hit
		for batchPayLoadSize <= b.PayLoadBytes && len(batchedEvents) < b.BatchSize && end < len(events) {
			batchPayLoadSize += len(events[end].Body.(string))
			if batchPayLoadSize <= b.PayLoadBytes {
				end = end + 1
				batchedEvents = events[i:end]
			}
		}
		i = end
		err := b.EventService.PostEvents(batchedEvents)

		if err != nil {
			b.Errors = append(b.Errors, ingestError{Error: err, Events: events})

			for len(b.EventsChan) >= cap(b.EventsChan) {
				time.Sleep(time.Duration(b.chanWaitMillis) * time.Millisecond)
			}
			if b.IsRunning {
				b.ErrorChan <- struct{}{}
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
	b.EventsChan = make(chan Event, b.BatchSize)
	b.ErrorChan = make(chan struct{}, cap(b.ErrorChan))

	b.Errors = nil
	b.ResetQueue()
	b.Run()
}
