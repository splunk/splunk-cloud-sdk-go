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
	HTTPStatus     string
	Message        string              `json:"message"`
	Code           string              `json:"code"`
	MoreInfo       string              `json:"moreInfo"`
	Details        []map[string]string `json:"details"`
}

// This allows HTTPError to satisfy the error interface
func (he *HTTPError) Error() string {
	if jsonErrMsg, err := json.Marshal(he); err != nil {
		return fmt.Sprintf("Unknown Http Error - HTTPStatusCode: [%v], HTTPStatus: %v, Error Message: %v, Error Code: %v",
			he.HTTPStatusCode, he.HTTPStatus, he.Message, he.Code)
	} else {
		return string(jsonErrMsg)
	}
}

// ParseHTTPStatusCodeInResponse creates a HTTPError from http status code and message
func ParseHTTPStatusCodeInResponse(response *http.Response) (*http.Response, error) {
	if response != nil && (response.StatusCode < 200 || response.StatusCode >= 400) {
		httpErr := &HTTPError{
			HTTPStatusCode: response.StatusCode,
			HTTPStatus:     response.Status,
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
