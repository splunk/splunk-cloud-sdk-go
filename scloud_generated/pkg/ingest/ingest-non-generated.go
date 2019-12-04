package ingest

import (
	"bufio"
	"encoding/json"
	"time"

	"github.com/golang/glog"

	"io"
	"os"
	"strings"

	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/auth"
	model "github.com/splunk/splunk-cloud-sdk-go/services/ingest"
)

const (
	alpha = 0.02
	beta  = 1.0 - alpha
)

// PostEvents Sends events.
func PostEventsOverride(args []model.Event, format string) (*model.HttpResponse, error) {
	var resp *model.HttpResponse

	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	total := 0
	start := time.Now()

	defer func() {
		delta := time.Since(start)
		eps := float64(total) / delta.Seconds()
		glog.Infof("Posted %d events (%v), %.2f events/sec", total, delta, eps)
	}()

	rps := 0.0 // requests/second
	eps := 0.0 // events/second

	r := bufio.NewReader(os.Stdin)

	for {
		batch, err := readBatch(r)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		startBatch := time.Now()
		events, err := makeEventBatch(batch, args[0], format)
		if err != nil {
			return nil, err
		}

		resp, err = client.IngestService.PostEvents(events)

		if err != nil {
			return nil, err
		}

		deltaBatch := time.Since(startBatch)

		count := len(batch)
		total += count

		secs := deltaBatch.Seconds()
		rpsBatch := 1.0 / secs
		epsBatch := float64(count) / secs

		if rps == 0.0 {
			rps = rpsBatch // initialize
			eps = epsBatch
		} else {
			rps = (rpsBatch * alpha) + (rps * beta)
			eps = (epsBatch * alpha) + (eps * beta)
		}

		glog.Infof("postEvents format=%s batch=%d total=%d rtt=%v rps=%.3f eps=%.3f",
			"raw", count, total, deltaBatch, rps, eps)
	}

	return resp, nil
}

// PostEvents Sends events.
func PostMetricsOverride(args []model.MetricEvent) (*model.HttpResponse, error) {

	var resp *model.HttpResponse

	total := 0
	start := time.Now()

	defer func() {
		delta := time.Since(start)
		eps := float64(total) / delta.Seconds()
		glog.Infof("Posted %d events (%v), %.2f events/sec", total, delta, eps)
	}()

	rps := 0.0 // requests/second
	eps := 0.0 // events/second

	r := bufio.NewReader(os.Stdin)

	for {
		batch, err := readBatch(r)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		startBatch := time.Now()

		resp, err = postBatchJSON(batch, args[0])

		if err != nil {
			return nil, err
		}

		deltaBatch := time.Since(startBatch)

		count := len(batch)
		total += count

		secs := deltaBatch.Seconds()
		rpsBatch := 1.0 / secs
		epsBatch := float64(count) / secs

		if rps == 0.0 {
			rps = rpsBatch // initialize
			eps = epsBatch
		} else {
			rps = (rpsBatch * alpha) + (rps * beta)
			eps = (epsBatch * alpha) + (eps * beta)
		}

		glog.Infof("postEvents format=%s batch=%d total=%d rtt=%v rps=%.3f eps=%.3f",
			"raw", count, total, deltaBatch, rps, eps)
	}

	return resp, nil
}

// Post the given batch of inputs.
func makeEventBatch(batch []string, args model.Event, format string) ([]model.Event, error) {
	var err error
	events := make([]model.Event, len(batch))

	switch format {
	case "event":
		events, err = postBatchEvent(batch, args)
	case "json":
		events, err = postBatchEventJSON(batch, args)
	case "raw":
		events, err = postBatchRaw(batch, args)
	default:
		events, err = postBatchRaw(batch, args)
	}
	return events, err
}

// Post the given batch of inputs, interpreted as event objects.
func postBatchEvent(batch []string, args model.Event) ([]model.Event, error) {
	events := make([]model.Event, len(batch))
	for i, item := range batch {
		event := &events[i]
		if err := json.Unmarshal([]byte(item), event); err != nil {
			return nil, err
		}
		var empty = ""
		if args.Host != &empty {
			event.Host = args.Host
		}
		if args.Source != &empty {
			event.Source = args.Source
		}
		if args.Sourcetype != &empty {
			event.Sourcetype = args.Sourcetype
		}
		if args.Attributes != nil {
			event.Attributes = args.Attributes
		}
		if args.Id != &empty {
			event.Id = args.Id
		}
		if args.Timestamp != nil {
			event.Timestamp = args.Timestamp
		}
	}
	return events, nil
}

// Post the given batch of inputs, interpreted as raw event data.
func postBatchRaw(batch []string, args model.Event) ([]model.Event, error) {
	events := make([]model.Event, len(batch))
	for i, item := range batch {
		events[i] = model.Event{
			Host:       args.Host,
			Source:     args.Source,
			Sourcetype: args.Sourcetype,
			Attributes: args.Attributes,
			Timestamp:  args.Timestamp,
			Id:         args.Id,
			Nanos:      args.Nanos,
			Body:       item}
	}
	return events, nil
}

// Post the given batch of inputs, interpreted as JSON event data.
func postBatchEventJSON(batch []string, args model.Event) ([]model.Event, error) {
	events := make([]model.Event, len(batch))
	for i, item := range batch {
		body, err := loads(item)
		if err != nil {
			return nil, err
		}
		events[i] = model.Event{
			Host:       args.Host,
			Source:     args.Source,
			Sourcetype: args.Sourcetype,
			Attributes: args.Attributes,
			Timestamp:  args.Timestamp,
			Id:         args.Id,
			Nanos:      args.Nanos,
			Body:       body}
	}
	return events, nil
}

func newMetricEvent(body []model.Metric, args model.MetricEvent) *model.MetricEvent {
	result := &model.MetricEvent{
		Host:       args.Host,
		Source:     args.Source,
		Sourcetype: args.Sourcetype,
		Body:       body,
		Attributes: args.Attributes,
	}
	return result
}

// Loads a list of metrics from the given string.
func loadMetrics(value string) ([]model.Metric, error) {
	var result []model.Metric
	r := strings.NewReader(value)
	if err := json.NewDecoder(r).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func postBatchJSON(batch []string, args model.MetricEvent) (*model.HttpResponse, error) {
	client, err := auth.GetClient()
	events := make([]model.MetricEvent, len(batch))
	for i, item := range batch {
		body, err := loadMetrics(item)
		if err != nil {
			return &model.HttpResponse{}, err
		}
		events[i] = *newMetricEvent(body, args)
	}
	glog.Infof("postMetrics format=json count=%d", len(events))
	httpResponse, err := client.IngestService.PostMetrics(events)
	return httpResponse, err
}

// Read a batch of event or metrics data, not to exceed size threshold.
// Note, this is an attempt to ensure we fit under the 1M kinesis limit, but
// this is not robust, because in the pathalogical case where event body is
// ~= envelope size, we can blow the kinesis limit. Doing this correctly will
// require a refactoring of the SDK APIs to give us more direct control of
// the request buffer.
func readBatch(r *bufio.Reader) ([]string, error) {
	size := 0
	batch := make([]string, 0, 1024)
	const maxBatchSize = 500 * 1024
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		line = strings.TrimRight(line, " \n")
		if len(line) == 0 {
			continue
		}
		batch = append(batch, line)
		size += len(line)
		if size > maxBatchSize {
			break
		}
	}
	if len(batch) == 0 {
		return nil, io.EOF
	}
	return batch, nil
}

func UploadFilesOverride(arg string) (*model.HttpResponse, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	fileName := arg

	if _, err := os.Stat(fileName); err != nil {
		return nil, err
	}

	err = client.IngestService.UploadFiles(fileName)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Loads a json "object" from the given string.
func loads(value string) (interface{}, error) {
	var result interface{}
	r := strings.NewReader(value)
	if err := json.NewDecoder(r).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
