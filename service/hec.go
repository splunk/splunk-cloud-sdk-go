package service

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
)

const hecServicePrefix = "hec2"
const hecServiceVersion = "v1"

// HecService implements hec service
type HecService service

// CreateEvent implements HEC2 event endpoint
func (h *HecService) CreateEvent(event model.HecEvent) error {
	var url = h.client.BuildURL(hecServicePrefix, hecServiceVersion, "events")
	response, err := h.client.Post(url, event)
	return util.ParseError(response, err)
}

// CreateRawEvent implements HEC2 raw endpoint
func (h *HecService) CreateRawEvent(event model.HecEvent) error {
	var url = h.client.BuildURL(hecServicePrefix, hecServiceVersion, "raw")
	url.RawQuery = util.ParseURLParams(event).Encode()
	response, err := h.client.Post(url, event.Event)
	return util.ParseError(response, err)
}
