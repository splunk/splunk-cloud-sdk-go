// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package service

import (
	"encoding/json"
	"github.com/splunk/ssc-client-go/model"
)

const ingestServicePrefix = "ingest"
const ingestServiceVersion = "v1"
// this is temporary. All endpoints will be reset to v1
const ingestServiceVersionV2 = "v2"

// IngestService talks to the SSC ingest service
type IngestService service

// CreateEvents post multiple events in one payload
func (h *IngestService) CreateEvents(events []model.Event) error {
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

// CreateMetricEvents implements Ingest metrics endpoint to send multipe metric events
func (h *IngestService) CreateMetricEvents(events []model.MetricEvent) error {
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
