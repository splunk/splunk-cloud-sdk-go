package model

import (
	"time"
	"fmt"
)

type batchEventService interface {
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

type BatchEventCollector struct {
	BatchSize       int
	Interval        time.Duration
	EventsChan      chan HecEvent
	QuitChan        chan struct{}
	HecEventService batchEventService
}

func (b *BatchEventCollector) Start() {
	ticker := time.NewTicker(b.Interval * time.Millisecond)
	go b.loop(ticker)
}

func (b *BatchEventCollector) loop (ticker *time.Ticker) {
	defer close(b.EventsChan)
	eventsQueue := make([]HecEvent, 0, b.BatchSize)
	for {
		select {
		case <- ticker.C:
			fmt.Println("flush from ticker")
			b.Flush(eventsQueue)
		case <- b.QuitChan:
			fmt.Println("done")
			ticker.Stop()
			b.Flush(eventsQueue)
			return
		case event, ok := <- b.EventsChan:
			if ok {
				eventsQueue = append(eventsQueue, event)
				fmt.Println(event.Host, "added")
			}
			if len(eventsQueue) >= b.BatchSize {
				b.Flush(eventsQueue)
				eventsQueue = nil
			}
		}
	}
}

func (b *BatchEventCollector) Stop() {
	b.QuitChan <- struct{}{}
}

func (b *BatchEventCollector) AddEvent(event HecEvent) {
		b.EventsChan <- event
}

func (b *BatchEventCollector) Flush(events []HecEvent) error {
	//if len(events) > b.BatchSize {
	//	for i := 0; i < len(events); i += b.BatchSize {
	//		end := i + b.BatchSize
	//		if end > len(events) {
	//			end = len(events)
	//		}
	//		err := b.HecEventService.CreateEvents(events[i:end])
	//		if err != nil {
	//			return err
	//		}
	//	}
	//} else {
	//	return b.HecEventService.CreateEvents(events)
	//}
	//return nil
	fmt.Println("flush")
	return nil
}