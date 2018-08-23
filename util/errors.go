// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTPError is raised when status code is not 2xx
type HTTPError struct {
	HTTPStatusCode int
	Message        string
	Code           string
	MoreInfo       string
	Details        []map[string]string
}

// This allows HTTPError to satisfy the error interface
func (he *HTTPError) Error() string {
	return fmt.Sprintf("Http Error - HTTPStatusCode: [%v], Message: %v",
		he.HTTPStatusCode, he.Message)
}

// ParseHTTPStatusCodeInResponse creates a HTTPError from http status code and message
func ParseHTTPStatusCodeInResponse(response *http.Response) (*http.Response, error) {
	if response != nil && (response.StatusCode < 200 || response.StatusCode >= 400) {
		httpErr := &HTTPError{
			HTTPStatusCode: response.StatusCode,
			Message:        response.Status,
		}
		if response.Body != nil {
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {

				return response, err
			}
			json.Unmarshal(body, &httpErr)
		}
		return response, httpErr
	}
	return response, nil
}
