// Code generated by gen_interface.go. DO NOT EDIT.

package ingest

// Servicer ...
type Servicer interface {
	// PostEvents post single or multiple events to ingest service
	PostEvents(events []Event) error
	// PostMetrics posts single or multiple metric events to ingest service
	PostMetrics(events []MetricEvent) error
	// NewBatchEventsSenderWithMaxAllowedError used to initialize dependencies and set values, the maxErrorsAllowed is the max number of errors allowed before the eventsender quit
	NewBatchEventsSenderWithMaxAllowedError(batchSize int, interval int64, dataSize int, maxErrorsAllowed int) (*BatchEventsSender, error)
	// NewBatchEventsSender used to initialize dependencies and set values
	NewBatchEventsSender(batchSize int, interval int64, payLoadSize int) (*BatchEventsSender, error)
}
