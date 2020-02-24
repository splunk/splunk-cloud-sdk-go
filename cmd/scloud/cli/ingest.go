/*
 * Copyright 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package main

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strings"
	"time"

	"github.com/golang/glog"

	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/argx"
	"github.com/splunk/splunk-cloud-sdk-go/services/ingest"
)

const (
	IngestServiceVersion = "v1beta2"
)

var createIngestService = func() *ingest.Service {
	return apiClient().IngestService
}

type IngestCommand struct {
	ingestService *ingest.Service
}

func newIngestCommand() *IngestCommand {
	return &IngestCommand{
		ingestService: createIngestService(),
	}
}

func (cmd *PostEventsCommand) client() ingest.Servicer {
	if cmd._client == nil {
		cmd._client = createIngestService()
	}
	return cmd._client
}

func (cmd *IngestCommand) Dispatch(argv []string) (result interface{}, err error) {
	arg, argv := head(argv)
	switch arg {
	case "":
		eusage("too few arguments")
	case "get-spec-json":
		result, err = cmd.getSpecJSON(argv)
	case "get-spec-yaml":
		result, err = cmd.getSpecYaml(argv)
	case "help":
		err = help("ingest.txt")
	case "post-events":
		result, err = newPostEventsCommand().postEvents(argv)
	case "post-metrics":
		result, err = newPostMetricsCommand().postMetrics(argv)
	case "upload-file":
		result = cmd.uploadFile(argv)
	default:
		fatal("unknown command: '%s'", arg)
	}
	return
}

func (cmd *IngestCommand) getSpecJSON(args []string) (interface{}, error) {
	checkEmpty(args)
	return GetSpecJSON("api", IngestServiceVersion, "ingest", cmd.ingestService.Client)
}

func (cmd *IngestCommand) getSpecYaml(args []string) (interface{}, error) {
	checkEmpty(args)
	return GetSpecYaml("api", IngestServiceVersion, "ingest", cmd.ingestService.Client)
}

func (cmd *IngestCommand) uploadFile(args []string) error {
	fileName := head1(args)

	if _, err := os.Stat(fileName); err != nil {
		fatal("Error with file : %s, please check if the filename is valid.", err.Error())
	}

	return newIngestCommand().ingestService.UploadFiles(fileName)
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

type postEventsArgs struct {
	Format     string `arg:"format"`
	Host       string `arg:"host"`
	Source     string `arg:"source"`
	Sourcetype string `arg:"sourcetype"`
}

type PostEventsCommand struct {
	args    postEventsArgs
	_client ingest.Servicer
}

func newPostEventsCommand() *PostEventsCommand {
	result := &PostEventsCommand{}
	result.args.Format = "raw" // default
	return result
}

func (cmd *PostEventsCommand) newEvent(body interface{}) ingest.Event {
	args := &cmd.args

	return ingest.Event{
		Host:       &args.Host,
		Source:     &args.Source,
		Sourcetype: &args.Sourcetype,
		Body:       body}
}

// Answers if the given HTTP request error indicates a request we should retry.
func shouldRetry(err error) bool {
	msg := err.Error()
	if strings.Contains(msg, "Client.Timeout exceeded while awaiting headers") {
		return true
	}
	return false
}

// Post the given batch of event objects, with retry on timeout.
func (cmd *PostEventsCommand) doPostEvents(events []ingest.Event) (err error) {
	client := cmd.client()
	for i := 0; ; i++ {
		_, err = client.PostEvents(events)
		if err == nil {
			return
		}
		glog.Error(err.Error())
		if i >= 2 {
			return
		}
		if !shouldRetry(err) {
			return
		}
	}
}

// Post the given batch of inputs.
func (cmd *PostEventsCommand) postBatch(batch []string) error {
	var err error
	switch cmd.args.Format {
	case "event":
		err = cmd.postBatchEvent(batch)
	case "json":
		err = cmd.postBatchJSON(batch)
	case "raw":
		err = cmd.postBatchRaw(batch)
	default:
		fatal("bad format: '%s'", cmd.args.Format)
	}
	return err
}

// Post the given batch of inputs, interpreted as event objects.
func (cmd *PostEventsCommand) postBatchEvent(batch []string) error {
	args := &cmd.args
	events := make([]ingest.Event, len(batch))
	for i, item := range batch {
		event := &events[i]
		if err := json.Unmarshal([]byte(item), event); err != nil {
			return err
		}
		if args.Host != "" {
			event.Host = &args.Host
		}
		if args.Source != "" {
			event.Source = &args.Source
		}
		if args.Sourcetype != "" {
			event.Sourcetype = &args.Sourcetype
		}
	}
	return cmd.doPostEvents(events)
}

// Post the given batch of inputs, interpreted as raw event data.
func (cmd *PostEventsCommand) postBatchRaw(batch []string) error {
	events := make([]ingest.Event, len(batch))
	for i, item := range batch {
		events[i] = cmd.newEvent(item)
	}
	return cmd.doPostEvents(events)
}

// Post the given batch of inputs, interpreted as JSON event data.
func (cmd *PostEventsCommand) postBatchJSON(batch []string) error {
	events := make([]ingest.Event, len(batch))
	for i, item := range batch {
		body, err := loads(item)
		if err != nil {
			return err
		}
		events[i] = cmd.newEvent(body)
	}
	return cmd.doPostEvents(events)
}

const (
	alpha = 0.02
	beta  = 1.0 - alpha
)

// Dispatch the post-events command.
func (cmd *PostEventsCommand) postEvents(argv []string) (interface{}, error) {
	argv, err := argx.Parse(argv, &cmd.args)
	if err != nil {
		return nil, err
	}
	checkEmpty(argv)

	total := 0
	start := time.Now()
	format := cmd.args.Format

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
		if err = cmd.postBatch(batch); err != nil {
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
			format, count, total, deltaBatch, rps, eps)
	}
	return nil, nil
}

type postMetricsArgs struct {
	Format     string            // `arg:"format"`
	Host       string            `arg:"host"`
	Source     string            `arg:"source"`
	Sourcetype string            `arg:"sourcetype"`
	Dimensions map[string]string `arg:"dimensions"`
	Type       string            `arg:"type"`
	Unit       string            `arg:"unit"`
}

type PostMetricsCommand struct {
	args    postMetricsArgs
	_client ingest.Servicer
}

func newPostMetricsCommand() *PostMetricsCommand {
	return &PostMetricsCommand{}
}

func (cmd *PostMetricsCommand) client() ingest.Servicer {
	if cmd._client == nil {
		cmd._client = createIngestService()
	}
	return cmd._client
}

func (cmd *PostMetricsCommand) newMetricEvent(body []ingest.Metric) *ingest.MetricEvent {
	args := &cmd.args
	result := &ingest.MetricEvent{
		Host:       &args.Host,
		Source:     &args.Source,
		Sourcetype: &args.Sourcetype,
		Body:       body}
	attrs := result.Attributes
	attrs.DefaultDimensions = args.Dimensions
	attrs.DefaultType = &args.Type
	attrs.DefaultUnit = &args.Unit
	return result
}

func (cmd *PostMetricsCommand) postBatch(batch []string) error {
	var err error
	switch cmd.args.Format {
	case "": // default
		_, err = cmd.postBatchJSON(batch)
	default:
		fatal("bad format: '%s'", cmd.args.Format)
	}
	return err
}

// Loads a list of metrics from the given string.
func loadMetrics(value string) ([]ingest.Metric, error) {
	var result []ingest.Metric
	r := strings.NewReader(value)
	if err := json.NewDecoder(r).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (cmd *PostMetricsCommand) postBatchJSON(batch []string) (ingest.HttpResponse, error) {
	events := make([]ingest.MetricEvent, len(batch))
	for i, item := range batch {
		body, err := loadMetrics(item)
		if err != nil {
			return ingest.HttpResponse{}, err
		}
		events[i] = *cmd.newMetricEvent(body)
	}
	glog.Infof("postMetrics format=json count=%d", len(events))
	httpResponse, err := cmd.client().PostMetrics(events)
	return *httpResponse, err
}

func (cmd *PostMetricsCommand) postMetrics(argv []string) (interface{}, error) {
	argv, err := argx.Parse(argv, &cmd.args)
	if err != nil {
		return nil, err
	}
	checkEmpty(argv)
	r := bufio.NewReader(os.Stdin)
	for {
		batch, err := readBatch(r)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if err = cmd.postBatch(batch); err != nil {
			return nil, err
		}
	}
	return nil, nil
}
