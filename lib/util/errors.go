package util

import (
	"fmt"
	"net/http"
)

// HTTPError is raised when status code is not 2xx
type HTTPError struct {
	Status  int
	Message string
}

// This allows HTTPError to satisfy the error interface
func (he *HTTPError) Error() string {
	return fmt.Sprintf("Http Error: [%v] %v",
		he.Status, he.Message)
}

// ParseHTTPStatusCodeInResponse creates a HTTPError from http status code and message
func ParseHTTPStatusCodeInResponse(response *http.Response) (*http.Response, error) {
	if response.StatusCode <= 200 || response.StatusCode >= 300 {
		return response, &HTTPError{Status: response.StatusCode, Message: response.Status}
	}
	return response, nil
}
