package service

import (
	"bytes"
	"encoding/json"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
	"time"
	"errors"
)

const hecServicePrefix = "hec2"

// HecService talks to the SSC hec service
type HecService service

// CreateEvent implements HEC2 event endpoint
func (h *HecService) CreateEvent(event model.HecEvent) error {
	url, err := h.client.BuildURL(hecServicePrefix, "events")
	if err != nil {
		return err
	}
	response, err := h.client.Post(url, event)
	return util.ParseError(response, err)
}

// CreateEvents post multiple events in one payload
func (h *HecService) CreateEvents(events []model.HecEvent) error {
	url, err := h.client.BuildURL(hecServicePrefix, "events")
	if err != nil {
		return err
	}
	hecEvents, err := h.buildMultiEventsPayload(events)
	response, err := h.client.Post(url, hecEvents)
	return util.ParseError(response, err)
}

// CreateRawEvent implements HEC2 raw endpoint
func (h *HecService) CreateRawEvent(event model.HecEvent) error {
	url, err := h.client.BuildURL(hecServicePrefix, "raw")
	if err != nil {
		return err
	}
	if param := util.ParseURLParams(event).Encode(); len(param) > 0 {
		url.RawQuery = param
	}
	response, err := h.client.Post(url, event.Event)
	return util.ParseError(response, err)
}

func (h *HecService) buildMultiEventsPayload(events []model.HecEvent) ([]byte, error) {
	var eventBuffer bytes.Buffer
	for _, event := range events {
		jsonBytes, err := json.Marshal(event)
		if err != nil {
			return nil, err
		}
		eventBuffer.Write(jsonBytes)
	}
	return eventBuffer.Bytes(), nil
}

// NewBatchEventsSender creates a new batch events sender
func (h *HecService) NewBatchEventsSender(batchSize int, interval int64) (*model.BatchEventsSender, error) {
	if batchSize == 0 || interval == 0 {
		return nil, errors.New("batchSize and interval cannot be 0")
	}
	eventsChan := make(chan model.HecEvent, batchSize)
	eventsQueue := make([]model.HecEvent, 0, batchSize)
	quit := make(chan struct{}, 1)
	return &model.BatchEventsSender{BatchSize:batchSize, Interval:time.Duration(interval), EventsChan:eventsChan, EventsQueue:eventsQueue, QuitChan:quit, EventService: h}, nil
}
