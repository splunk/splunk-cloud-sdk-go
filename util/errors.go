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
	Message        string      `json:"message,omitempty"`
	Code           string      `json:"code,omitempty"`
	MoreInfo       string      `json:"moreInfo,omitempty"`
	Details        interface{} `json:"details,omitempty"`
}

// This allows HTTPError to satisfy the error interface
func (he *HTTPError) Error() string {
	jsonErrMsg, err := json.Marshal(he)
	if err != nil {
		return fmt.Sprintf("Http Error - HTTPStatusCode: [%v], HTTPStatus: %v, Error Message: %v, Error Code: %v",
			he.HTTPStatusCode, he.HTTPStatus, he.Message, he.Code)
	}
	return string(jsonErrMsg)
}

// ParseHTTPStatusCodeInResponse returns http response and HTTPError struct based on response status code
func ParseHTTPStatusCodeInResponse(response *http.Response) (*http.Response, error) {
	if response != nil && (response.StatusCode < 200 || response.StatusCode >= 400) {
		httpErr := HTTPError{
			HTTPStatusCode: response.StatusCode,
			HTTPStatus:     response.Status,
		}
		if response.Body != nil {
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return response, &httpErr
			}
			err = json.Unmarshal(body, &httpErr)
			if err != nil {
				return nil, &httpErr
			}
		}
		return response, &httpErr
	}
	return response, nil
}
