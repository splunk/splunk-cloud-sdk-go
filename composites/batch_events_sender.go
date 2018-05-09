package composites

import (
	"errors"
	"time"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/service"
)

// HECEventService defines a new interface to avoid cycle import error
type HECEventService interface {
	CreateEvent(event model.HecEvent) error
	CreateEvents(events []model.HecEvent) error
}

// BatchEventsSender sends events in batches or periodically if batch is not full to Splunk HTTP Event Collector endpoint
type BatchEventsSender struct {
	BatchSize   int
	Interval    time.Duration
	EventsChan  chan model.HecEvent
	EventsQueue []model.HecEvent
	QuitChan    chan struct{}
	HecService  service.HecService
}

func (b *BatchEventsSender) Run() {
	ticker := time.NewTicker(b.Interval * time.Millisecond)
	go b.loop(ticker)
}

func (b *BatchEventsSender) loop(ticker *time.Ticker) {
	for {
		select {
		case <-b.QuitChan:
			// this channel is used to make sure the last flush is finished before the loop breaks
			done := make(chan struct{}, 1)
			go b.Flush(b.HecService, b.EventsQueue, done)
			ticker.Stop()
			b.EventsQueue = b.EventsQueue[:0]
			<-done
			return
		case <-ticker.C:
			go b.Flush(b.HecService, b.EventsQueue, nil)
			b.EventsQueue = b.EventsQueue[:0]
		case event := <-b.EventsChan:
			b.EventsQueue = append(b.EventsQueue, event)
			if len(b.EventsQueue) >= b.BatchSize {
				go b.Flush(b.HecService, b.EventsQueue, nil)
				b.EventsQueue = b.EventsQueue[:0]
			}
		}
	}
}

func (b *BatchEventsSender) Start() {
	//
}

func (b *BatchEventsSender) Stop() {
	b.QuitChan <- struct{}{}
}

func (b *BatchEventsSender) AddEvent(event model.HecEvent) {
	b.EventsChan <- event
}

// TODO: Error handling and return results
func (b *BatchEventsSender) Flush(hecService service.HecService, events []model.HecEvent, doneChan chan struct{}) {
	if len(events) > b.BatchSize {
		for i := 0; i < len(events); i += b.BatchSize {
			end := i + b.BatchSize
			if end > len(events) {
				end = len(events)
			}
			hecService.CreateEvents(events[i:end])
		}
		if doneChan != nil {
			doneChan <- struct{}{}
		}
		return
	}
	if len(events) > 0 {
		hecService.CreateEvents(events)
		if doneChan != nil {
			doneChan <- struct{}{}
		}
		return
	}
	if doneChan != nil {
		doneChan <- struct{}{}
	}
	return

}

func (b *BatchEventsSender) NewBatchEventsSender(hecService service.HecService, batchSize int, interval int64) (*BatchEventsSender, error) {
	if batchSize == 0 || interval == 0 {
		return nil, errors.New("batchSize and interval cannot be 0")
	}
	eventsChan := make(chan model.HecEvent, batchSize)
	eventsQueue := make([]model.HecEvent, 0, batchSize)
	quit := make(chan struct{}, 1)

	batchEventsSender := new(BatchEventsSender)
	batchEventsSender.BatchSize = batchSize
	batchEventsSender.Interval = time.Duration(interval)
	batchEventsSender.EventsChan = eventsChan
	batchEventsSender.EventsQueue = eventsQueue
	batchEventsSender.QuitChan = quit
	// TODO: Figure out a way to remove this?
	batchEventsSender.HecService = hecService

	return batchEventsSender, nil

	//return &BatchEventsSender{BatchSize: batchSize, Interval: time.Duration(interval), EventsChan: eventsChan, EventsQueue: eventsQueue, QuitChan: quit, EventService: h}, nil
}
