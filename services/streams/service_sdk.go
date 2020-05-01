package streams

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/splunk/splunk-cloud-sdk-go/services"
)

/*
	UploadFiles - Upload a CSV or text file that contains events.
	Parameters:
		filename
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) UploadFiles(filename string, resp ...*http.Response) error {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/files`, nil)

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
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/files`, nil)

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

	response, err := s.Client.Post(services.RequestParams{URL: u, Body: form, Headers: map[string]string{"Content-Type": "multipart/form-data"}})

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
