package model

import (
	"time"
)

// HECEventService defines a new interface to avoid cycle import error
type HECEventService interface {
	CreateEvent(event HecEvent) error
	CreateEvents(events []HecEvent) error
}

// HecEvent contains metadata about the event
type HecEvent struct {
	Host       string            `json:"host,omitempty" key:"host"`
	Index      string            `json:"index,omitempty" key:"index"`
	Sourcetype string            `json:"sourcetype,omitempty" key:"sourcetype"`
	Source     string            `json:"source,omitempty" key:"source"`
	Time       *float64          `json:"time,omitempty" key:"time"`
	Event      interface{}       `json:"event"`
	Fields     map[string]string `json:"fields,omitempty"`
}

// BatchEventsSender sends events in batches or periodically if batch is not full to Splunk HTTP Event Collector endpoint
type BatchEventsSender struct {
	BatchSize    int
	Interval     time.Duration
	EventsChan   chan HecEvent
	EventsQueue  []HecEvent
	QuitChan     chan struct{}
	EventService HECEventService
}

// Run sets up ticker and starts a new goroutine
func (b *BatchEventsSender) Run() {
	ticker := time.NewTicker(b.Interval * time.Millisecond)
	go b.loop(ticker)
}

// loop is a infinity loop which listens to three channels
// ticker.C: ticker channel to be used for timer
// QuitChan: shut down signal channel
// EventsChan: events channel
// the loop will break only if there's signal in QuitChan
// otherwise it'll constantly checking conditions for ticker and events
func (b *BatchEventsSender) loop (ticker *time.Ticker) {
	for {
		select {
		case <- ticker.C:
			b.Flush()
		case <- b.QuitChan:
			ticker.Stop()
			close(b.EventsChan)
			b.Flush()
			return
		case event := <- b.EventsChan:
			if len(b.EventsQueue) >= b.BatchSize {
				b.Flush()
			}
			b.EventsQueue = append(b.EventsQueue, event)
		}
	}
}

// Stop sends a signal to QuitChan and shuts down the goroutine that are created in Run()
func (b *BatchEventsSender) Stop() {
	b.QuitChan <- struct{}{}
}

// AddEvent pushes a single event into EventsChan
func (b *BatchEventsSender) AddEvent(event HecEvent) {
		b.EventsChan <- event
}

// Flush sends off all events currently in EventsQueue and clear EventsQueue afterwards
// If EventsQueue size is bigger than BatchSize, it'll slice the queue into batches and send batches one by one
// TODO: Error handling and return results
func (b *BatchEventsSender) Flush() {
	if len(b.EventsQueue) > b.BatchSize {
		for i := 0; i < len(b.EventsQueue); i += b.BatchSize {
			end := i + b.BatchSize
			if end > len(b.EventsQueue) {
				end = len(b.EventsQueue)
			}
			b.EventService.CreateEvents(b.EventsQueue[i:end])
		}
	} else if len(b.EventsQueue) > 0 {
		b.EventService.CreateEvents(b.EventsQueue)
	}

	// This will keep the memory allocated to underlying array, but clear the value in the queue
	b.EventsQueue = b.EventsQueue[:0]
}