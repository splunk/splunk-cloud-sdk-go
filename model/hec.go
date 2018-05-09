package model

import (
	"fmt"
	"sync"
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
	EventsChan   chan HecEvent
	EventsQueue  []HecEvent
	QuitChan     chan struct{}
	EventService HECEventService
	HecTicker    *Ticker
	wg           sync.WaitGroup
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
	for {
		select {
		case <-b.QuitChan:
			fmt.Println("stop")
			events := append([]HecEvent(nil), b.EventsQueue...)
			fmt.Println("queue size from stop:", len(events))
			b.wg.Add(1)
			go b.Flush(events)
			return
		case <-b.HecTicker.ticker.C:
			fmt.Println("ticker ticked")
			events := append([]HecEvent(nil), b.EventsQueue...)
			b.wg.Add(1)
			go b.Flush(events)
			b.ResetQueue()
		case event := <-b.EventsChan:
			b.EventsQueue = append(b.EventsQueue, event)
			fmt.Println("queue size from eventschan:", len(b.EventsQueue))
			if len(b.EventsQueue) == b.BatchSize {
				events := append([]HecEvent(nil), b.EventsQueue...)
				b.wg.Add(1)
				go b.Flush(events)
				b.ResetQueue()
			}
		}
	}
}

// Stop sends a signal to QuitChan and shuts down the goroutine that are created in Run()
func (b *BatchEventsSender) Stop() {
	b.QuitChan <- struct{}{}
	b.wg.Wait()
	b.HecTicker.Stop()
	b.ResetQueue()
}

// AddEvent pushes a single event into EventsChan
func (b *BatchEventsSender) AddEvent(event HecEvent) {
	b.EventsChan <- event
}

// Flush sends off all events currently in EventsQueue and clear EventsQueue afterwards
// If EventsQueue size is bigger than BatchSize, it'll slice the queue into batches and send batches one by one
// TODO: Error handling and return results
func (b *BatchEventsSender) Flush(events []HecEvent) {
	defer b.wg.Done()
	fmt.Println("events size", len(events))
	if len(events) > 0 {
		b.EventService.CreateEvents(events)
		fmt.Println("flushed")
	}
	// Reset ticker
	b.HecTicker.Reset()
}

func (b *BatchEventsSender) ResetQueue() {
	b.EventsQueue = b.EventsQueue[:0]
}
