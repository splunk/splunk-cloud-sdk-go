/*
* Splunk Ingest Service
*
 */

package ingest

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
	"time"

	"errors"

	gdepservices "github.com/splunk/go-dependencies/services"
	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

/*
	NewBatchEventsSenderWithMaxAllowedError initializes a BatchEventsSender to collect events and send them as a single
        batched request when a maximum event batch size, time interval, or maximum payload size is reached. It also
        validates the user input for BatchEventSender.
	Parameters:
		batchSize: maximum number of events to reach before sending the batch, default maximum is 500
		interval: milliseconds to wait before sending the batch if other conditions have not been met
		dataSize: bytes that the overall payload should not exceed before sending, default maximum is 1040000 ~1MiB
		maxErrorsAllowed: number of errors after which the BatchEventsSender will stop
*/
func (s *Service) NewBatchEventsSenderWithMaxAllowedError(batchSize int, interval int64, dataSize int, maxErrorsAllowed int) (*BatchEventsSender, error) {
	// Rather than return a super general error for both it will block on batchSize first
	if batchSize == 0 {
		return nil, errors.New("batchSize cannot be 0")
	}
	if batchSize > eventCount {
		batchSize = eventCount
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

/*
	NewBatchEventsSender initializes a BatchEventsSender to collect events and send them as a single batched
        request when a maximum event batch size, time interval, or maximum payload size is reached.
	Parameters:
		batchSize: maximum number of events to reach before sending the batch, default maximum is 500
		interval: milliseconds to wait before sending the batch if other conditions have not been met
		payLoadSize: bytes that the overall payload should not exceed before sending, default maximum is 1040000 ~1MiB
*/
func (s *Service) NewBatchEventsSender(batchSize int, interval int64, payLoadSize int) (*BatchEventsSender, error) {
	return s.NewBatchEventsSenderWithMaxAllowedError(batchSize, interval, payLoadSize, 1)
}

/*
	UploadFiles - Upload a CSV or text file that contains events.
	Parameters:
		filename
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) UploadFiles(filename string, resp ...*http.Response) error {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ingest/v1beta2/files`, nil)

	if err != nil {
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var response *http.Response
	if len(resp) > 0 && resp[0] != nil {
		response = resp[0]
	}
	return s.uploadFileStream(u, file, filepath.Base(filename), response)
}

/*
	UploadFilesStream - Upload stream of io.Reader.
	Parameters:
		stream
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) UploadFilesStream(stream io.Reader, resp ...*http.Response) error {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ingest/v1beta2/files`, nil)

	if err != nil {
		return err
	}

	var response *http.Response
	if len(resp) > 0 && resp[0] != nil {
		response = resp[0]
	}

	return s.uploadFileStream(u, stream, "stream", response)
}

func (s *Service) uploadFileStream(u url.URL, stream io.Reader, filename string, resp *http.Response) error {

	form := services.FormData{Filename: filename, Stream: stream, Key: "upfile"}

	response, err := s.Client.Post(gdepservices.RequestParams{URL: u, Body: form, Headers: map[string]string{"Content-Type": "multipart/form-data"}})

	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if resp != nil {
			*resp = *response
		}
	}

	if err != nil {
		return err
	}

	return nil
}
