package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTPError is raised when status code is not 2xx
type HTTPError struct {
	Status  int
	Message string
	Body    string
}

// This allows HTTPError to satisfy the error interface
func (he *HTTPError) Error() string {
	return fmt.Sprintf("Http Error: [%v] %v %v",
		he.Status, he.Message, he.Body)
}

// ParseHTTPStatusCodeInResponse creates a HTTPError from http status code and message
func ParseHTTPStatusCodeInResponse(response *http.Response) (*http.Response, error) {
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(response.Body)

		return response, &HTTPError{
			Status:  response.StatusCode,
			Message: response.Status,
			Body:    string(body),
		}
	}

	return response, nil
}
