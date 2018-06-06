package service

import (
	"bytes"
	"encoding/json"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
)

const hecServicePrefix = "ingest"

// HecService talks to the SSC hec service
type HecService service

// CreateEvent implements HEC2 event endpoint
func (h *HecService) CreateEvent(event model.HecEvent) error {
	url, err := h.client.BuildURL(nil, hecServicePrefix, "v1", "events")
	if err != nil {
		return err
	}
	response, err := h.client.Post(url, event)
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// CreateEvents post multiple events in one payload
func (h *HecService) CreateEvents(events []model.HecEvent) error {
	url, err := h.client.BuildURL(nil, hecServicePrefix, "v1", "events")
	if err != nil {
		return err
	}
	hecEvents, err := h.buildMultiEventsPayload(events)
	if err != nil {
		return err
	}
	response, err := h.client.Post(url, hecEvents)
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// CreateRawEvent implements HEC2 raw endpoint
func (h *HecService) CreateRawEvent(event model.HecEvent) error {
	url, err := h.client.BuildURL(nil, hecServicePrefix, "v1", "raw")
	if err != nil {
		return err
	}
	if param := util.ParseURLParams(event).Encode(); len(param) > 0 {
		url.RawQuery = param
	}
	response, err := h.client.Post(url, event.Event)
	if response != nil {
		defer response.Body.Close()
	}
	return err
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
