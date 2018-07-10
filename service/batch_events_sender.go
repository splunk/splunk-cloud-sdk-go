package service

import (
	"errors"
	"fmt"
	"github.com/splunk/ssc-client-go/model"
	"sync"
)

type UserErrHandlerFunc func(*BatchEventsSender)

// BatchEventsSender sends events in batches or periodically if batch is not full to Splunk HTTP Event Collector endpoint
type BatchEventsSender struct {
	BatchSize    int
	EventsChan   chan model.HecEvent
	EventsQueue  []model.HecEvent
	QuitChan     chan struct{}
	EventService *HecService
	HecTicker    *model.Ticker
	WaitGroup    *sync.WaitGroup
	ErrorChan    chan string
	ErrorMsg     string
	IsRunning    bool
	mux          sync.Mutex
	callbackFunc UserErrHandlerFunc
}

// Run sets up ticker and starts a new goroutine
func (b *BatchEventsSender) SetCallbackFunc(callback UserErrHandlerFunc) {
	b.callbackFunc = callback
}

// Run sets up ticker and starts a new goroutine
func (b *BatchEventsSender) Run() {
	go b.loop()
	b.IsRunning = true
}

// loop is a infinity loop which listens to three channels
// ticker.C: ticker channel to be used for timer
// QuitChan: shut down signal channel
// EventsChan: events channel
// the loop will break only if there's signal in QuitChan
// otherwise it'll constantly checking conditions for ticker and events
func (b *BatchEventsSender) loop() {
	errorMsgCount := 0

	defer close(b.EventsChan)
	for {
		select {
		case err := <-b.ErrorChan:
			errorMsgCount++
			b.ErrorMsg += "[" + err + "],"
			fmt.Println("got an error : " + err)

			if b.callbackFunc != nil {
				b.callbackFunc(b)
			}

			if errorMsgCount == cap(b.ErrorChan) {
				b.Stop()
			}

		case <-b.QuitChan:
			return

		case <-b.HecTicker.GetChan():
			b.WaitGroup.Add(1)
			go b.flush(true)

		case event := <-b.EventsChan:
			b.EventsQueue = append(b.EventsQueue, event)
			if len(b.EventsQueue) == b.BatchSize {
				b.WaitGroup.Add(1)
				go b.flush(false)
			}
		}
	}
}

// Stop sends a signal to QuitChan, wait for all registered goroutines to finish, stop ticker and clear queue
func (b *BatchEventsSender) Stop() {
	b.IsRunning = false
	// Wait until no element is in channel
	for {
		if len(b.EventsChan) == 0 {
			break
		}
	}
	b.HecTicker.Stop()

	b.WaitGroup.Add(1)
	// flush one last time before stop
	go b.flush(true)
	b.WaitGroup.Wait()

	b.QuitChan <- struct{}{}
}

// AddEvent pushes a single event into EventsChan
func (b *BatchEventsSender) AddEvent(event model.HecEvent) error {
	if !b.IsRunning {
		return errors.New("Need to start the BatchEventsSender first, call Run() ")
	}

	// Intend to only start ticker when first event is received.
	if len(b.EventsQueue) == 0 && len(b.EventsChan) == 0 && b.HecTicker.IsRunning() == false {
		b.HecTicker.Start()
	}
	b.EventsChan <- event
	return nil
}

// flush sends off all events currently in EventsQueue and resets ticker afterwards
// If EventsQueue size is bigger than BatchSize, it'll slice the queue into batches and send batches one by one
func (b *BatchEventsSender) flush(fromTicker bool) error {
	defer b.WaitGroup.Done()

	b.mux.Lock()
	// Reset ticker
	if fromTicker {
		b.HecTicker.Reset()
	} else if len(b.EventsQueue) < b.BatchSize {
		// it is possible different threads send flush signal while the previous flush already flush everything in queue
		b.mux.Unlock()
		return nil
	}

	events := append([]model.HecEvent(nil), b.EventsQueue...)
	b.ResetQueue()

	// slice events into batch size to send
	if len(events) > 0 {
		for i := 0; i < len(events); {
			end := len(events)
			if i+b.BatchSize < len(events) {
				end = i + b.BatchSize
			}

			err := b.EventService.CreateEvents(events[i:end])
			i = i + b.BatchSize
			if err != nil {
				str := fmt.Sprintf("Failed to send all events for batch: %v\n\tError: %v", events, err)
				b.ErrorChan <- str
			}
		}
	}

	b.mux.Unlock()
	return nil
}

// ResetQueue sets b.EventsQueue to empty, but keep memory allocated for underlying array
func (b *BatchEventsSender) ResetQueue() {
	b.EventsQueue = b.EventsQueue[:0]
}