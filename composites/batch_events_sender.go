package composites

import (
	"errors"
	"time"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/service"
)

// BatchEventsSender sends events in batches or periodically if batch is not full to Splunk HTTP Event Collector endpoint
type BatchEventsSender struct {
	BatchSize   int
	Interval    time.Duration
	EventsChan  chan model.HecEvent
	EventsQueue []model.HecEvent
	QuitChan    chan struct{}
	HecService  *service.HecService
}

// Run sets up ticker and starts a new goroutine
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

// Stop sends a signal to QuitChan and shuts down the goroutine that are created in Run()
func (b *BatchEventsSender) Stop() {
	b.QuitChan <- struct{}{}
}

// AddEvent pushes a single event into EventsChan
func (b *BatchEventsSender) AddEvent(event model.HecEvent) {
	b.EventsChan <- event
}

// Flush sends off all events currently in EventsQueue and clear EventsQueue afterwards
// If EventsQueue size is bigger than BatchSize, it'll slice the queue into batches and send batches one by one
// TODO: Error handling and return results
func (b *BatchEventsSender) Flush(hecService *service.HecService, events []model.HecEvent, doneChan chan struct{}) {
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

// NewBatchEventsSender used to initialize dependencies and set values
func NewBatchEventsSender(hecService *service.HecService, batchSize int, interval int64) (*BatchEventsSender, error) {
	// Rather than return a super general error for both it will block on batchSize first
	if batchSize == 0 {
		return nil, errors.New("batchSize cannot be 0")
	}
	if interval == 0 {
		return nil, errors.New("interval cannot be 0")
	}

	eventsChan := make(chan model.HecEvent, batchSize)
	eventsQueue := make([]model.HecEvent, 0, batchSize)
	quit := make(chan struct{}, 1)

	batchEventsSender := &BatchEventsSender{
		BatchSize:   batchSize,
		EventsChan:  eventsChan,
		EventsQueue: eventsQueue,
		HecService:  hecService,
		Interval:    time.Duration(interval),
		QuitChan:    quit,
	}

	return batchEventsSender, nil
}
