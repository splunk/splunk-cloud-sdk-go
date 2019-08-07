/*
* Splunk Ingest Service
*
 */

package ingest

import (
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

//
import (
	"errors"
)

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

/*
	UploadFiles - Upload a CSV or text file that contains events.
	Parameters:
		upfile
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) UploadFiles(upfile string, resp ...*http.Response) error {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ingest/v1beta2/files`, nil)

	if err != nil {
		return err
	}

	response, err := s.Client.Post(services.RequestParams{URL: u, Body: map[string]string{"fieldname": "upfile", "upfile": upfile}, IsFormData: true})

	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}

	if err != nil {
		return err
	}

	return nil
}

/*
	UploadFiles - Upload stream of io.Reader.
	Parameters:
		stream
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) UploadStream(stream io.Reader, resp ...*http.Response) error {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ingest/v1beta2/files`, nil)

	if err != nil {
		return err
	}

	response, err := s.Client.Post(services.RequestParams{URL: u, Body: map[string]interface{}{"fieldname": "upfile", "upfile": stream}, IsFormData: true})

	// populate input *http.Response if provided
	if response != nil && len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}

	if err != nil {
		return err
	}

	return nil
}
