// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package model

import (
	"github.com/splunk/splunk-cloud-sdk-go/services/ingest"
)

// Event in the model package is DEPRECATED - please use service.ingest.Event
type Event = ingest.Event

// MetricEvent in the model package is DEPRECATED - please use service.ingest.MetricEvent
type MetricEvent = ingest.MetricEvent

// Metric in the model package is DEPRECATED - please use service.ingest.Metric
type Metric = ingest.Metric

// MetricAttribute in the model package is DEPRECATED - please use service.ingest.MetricAttribute
type MetricAttribute = ingest.MetricAttribute
