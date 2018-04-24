package service

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
)

const hecServicePrefix = "hec2"

// HecService implements hec service
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
