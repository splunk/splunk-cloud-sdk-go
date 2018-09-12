// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package service

import (
	"encoding/json"
	"github.com/splunk/splunk-cloud-sdk-go/model"
)

const ingestServicePrefix = "ingest"
const ingestServiceVersion = "v1"
// this is temporary. All endpoints will be reset to v1
const ingestServiceVersionV2 = "v2"

// IngestService talks to the Splunk Cloud ingest service
type IngestService service

// PostEvents post single or multiple events to ingest service
func (h *IngestService) PostEvents(events []model.Event) error {
	url, err := h.client.BuildURL(nil, ingestServicePrefix, ingestServiceVersionV2, "events")
	if err != nil {
		return err
	}
	jsonBytes, err := json.Marshal(events)
	if err != nil {
		return err
	}
	response, err := h.client.Post(RequestParams{URL: url, Body: jsonBytes})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// PostMetrics posts single or multiple metric events to ingest service
func (h *IngestService) PostMetrics(events []model.MetricEvent) error {
	url, err := h.client.BuildURL(nil, ingestServicePrefix, ingestServiceVersion, "metrics")
	if err != nil {
		return err
	}

	jsonBytes, err := json.Marshal(events)
	if err != nil {
		return err
	}

	response, err := h.client.Post(RequestParams{URL: url, Body: jsonBytes})
	if response != nil {
		defer response.Body.Close()
	}

	return err
}
