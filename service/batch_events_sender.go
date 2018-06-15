package service

import (
	"fmt"
	"github.com/splunk/ssc-client-go/model"
	"sync"
)

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
}

// Run sets up ticker and starts a new goroutine
func (b *BatchEventsSender) Run() {
	go b.loop()
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

			if errorMsgCount == cap(b.ErrorChan)-1 {
				b.Stop()
			}

		case <-b.QuitChan:
			events := append([]model.HecEvent(nil), b.EventsQueue...)
			b.WaitGroup.Add(1)
			// Flush one last time before exit
			go b.Flush(events)
			return
		case <-b.HecTicker.GetChan():
			events := append([]model.HecEvent(nil), b.EventsQueue...)
			b.WaitGroup.Add(1)
			go b.Flush(events)
			b.ResetQueue()
		case event := <-b.EventsChan:
			b.EventsQueue = append(b.EventsQueue, event)
			if len(b.EventsQueue) == b.BatchSize {
				events := append([]model.HecEvent(nil), b.EventsQueue...)
				b.WaitGroup.Add(1)
				go b.Flush(events)
				b.ResetQueue()
			}
		}
	}
}

// Stop sends a signal to QuitChan, wait for all registered goroutines to finish, stop ticker and clear queue
func (b *BatchEventsSender) Stop() {
	// Wait until no element is in channel
	for {
		if len(b.EventsChan) == 0 {
			break
		}
	}
	b.QuitChan <- struct{}{}
	b.WaitGroup.Wait()
	b.HecTicker.Stop()
	b.ResetQueue()
}

// AddEvent pushes a single event into EventsChan
func (b *BatchEventsSender) AddEvent(event model.HecEvent) {
	// Intend to only start ticker when first event is received.
	if len(b.EventsQueue) == 0 && len(b.EventsChan) == 0 && b.HecTicker.IsRunning() == false {
		b.HecTicker.Start()
	}
	b.EventsChan <- event
}

// Flush sends off all events currently in EventsQueue and resets ticker afterwards
// If EventsQueue size is bigger than BatchSize, it'll slice the queue into batches and send batches one by one
// TODO: Error handling and return results
func (b *BatchEventsSender) Flush(events []model.HecEvent) error {
	defer b.WaitGroup.Done()
	// Reset ticker
	b.HecTicker.Reset()
	if len(events) > 0 {
		err := b.EventService.CreateEvents(events)
		if err != nil {
			str := fmt.Sprintf("Failed to send all events for batch: %v\n\tError: %v", events, err)
			b.ErrorChan <- str
		}
	}

	return nil
}

// ResetQueue sets b.EventsQueue to empty, but keep memory allocated for underlying array
func (b *BatchEventsSender) ResetQueue() {
	b.EventsQueue = b.EventsQueue[:0]
}
