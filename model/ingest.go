// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package model

import (
	"github.com/splunk/splunk-cloud-sdk-go/services/ingest"
)

// Event is Deprecated: please use services/ingest.Event
type Event = ingest.Event

// MetricEvent is Deprecated: please use services/ingest.MetricEvent
type MetricEvent = ingest.MetricEvent

// Metric is Deprecated: please use services/ingest.Metric
type Metric = ingest.Metric

// MetricAttribute is Deprecated: please use services/ingest.MetricAttribute
type MetricAttribute = ingest.MetricAttribute
