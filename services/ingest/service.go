// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package ingest

import (
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

const servicePrefix = "ingest"
const serviceVersion = "v1beta2"
const serviceCluster = "api"

// Service talks to the Splunk Cloud ingest service
type Service services.BaseService

// NewService creates a new ingest service client from the given Config
func NewService(config *services.Config) (*Service, error) {
	baseClient, err := services.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Service{Client: baseClient}, nil
}

// PostEvents post single or multiple events to ingest service
func (s *Service) PostEvents(events []Event) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "events")
	if err != nil {
		return err
	}
	jsonBytes, err := json.Marshal(events)
	if err != nil {
		return err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: jsonBytes})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// PostMetrics posts single or multiple metric events to ingest service
func (s *Service) PostMetrics(events []MetricEvent) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "metrics")
	if err != nil {
		return err
	}

	jsonBytes, err := json.Marshal(events)
	if err != nil {
		return err
	}

	response, err := s.Client.Post(services.RequestParams{URL: url, Body: jsonBytes})
	if response != nil {
		defer response.Body.Close()
	}

	return err
}

// NewBatchEventsSenderWithMaxAllowedError used to initialize dependencies and set values, the maxErrorsAllowed is the max number of errors allowed before the eventsender quit
func (s *Service) NewBatchEventsSenderWithMaxAllowedError(batchSize int, interval int64, dataSize int, maxErrorsAllowed int) (*BatchEventsSender, error) {
	// Rather than return a super general error for both it will block on batchSize first
	if batchSize == 0 {
		return nil, errors.New("batchSize cannot be 0")
	}
	if interval == 0 {
		return nil, errors.New("interval cannot be 0")
	}
	if dataSize == 0 {
		dataSize = payLoadSize
	}

	if maxErrorsAllowed < 0 {
		maxErrorsAllowed = 1
	}

	eventsChan := make(chan Event, batchSize)
	eventsQueue := make([]Event, 0, batchSize)
	quit := make(chan struct{}, 1)
	ticker := util.NewTicker(time.Duration(interval) * time.Millisecond)
	var wg sync.WaitGroup
	errorChan := make(chan struct{}, maxErrorsAllowed)

	batchEventsSender := &BatchEventsSender{
		BatchSize:      batchSize,
		PayLoadBytes:   dataSize,
		EventsChan:     eventsChan,
		EventsQueue:    eventsQueue,
		EventService:   s,
		QuitChan:       quit,
		IngestTicker:   ticker,
		WaitGroup:      &wg,
		ErrorChan:      errorChan,
		IsRunning:      false,
		chanWaitMillis: 300,
		callbackFunc:   nil,
	}

	return batchEventsSender, nil
}

// NewBatchEventsSender used to initialize dependencies and set values
func (s *Service) NewBatchEventsSender(batchSize int, interval int64, payLoadSize int) (*BatchEventsSender, error) {
	return s.NewBatchEventsSenderWithMaxAllowedError(batchSize, interval, payLoadSize, 1)
}
