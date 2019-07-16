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
	"fmt"
	"net/http"
	"net/http/httputil"
)

// Logger compatible with standard "log" library
type Logger interface {
	Print(v ...interface{})
}

// SdkTransport is to define a transport RoundTripper with user-defined logger
type SdkTransport struct {
	transport http.RoundTripper
	logger    Logger
}

// 80 chars wide
const divider = "================================================================================"

// RoundTrip implements the RoundTripper interface
func (st *SdkTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	st.logger.Print("\n")
	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		st.logger.Print(fmt.Sprintf("%s\nRequest error:\n%s\n%s\n", divider, divider, err.Error()))
		return nil, err
	}

	st.logger.Print(fmt.Sprintf("%s\nRequest:\n%s\n", divider, string(requestDump)))

	response, err := st.transport.RoundTrip(request)
	if response != nil {
		st.logger.Print(fmt.Sprintf("%s\nResponse from: %s %s [%d]\nRequest ID: %s\n", divider, request.Method, request.URL, response.StatusCode, response.Header.Get("X-Request-ID")))
	}
	if err != nil {
		st.logger.Print(fmt.Sprintf("%s\nRequest error:\n%s\n", divider, err.Error()))
		return response, err
	}

	responseDump, err := httputil.DumpResponse(response, true)
	if err != nil {
		st.logger.Print(fmt.Sprintf("%s\nResponse error:\n%s\n", divider, err.Error()))
		return response, err
	}

	st.logger.Print(fmt.Sprintf("%s\nResponse:\n%s\n%s\n", divider, string(responseDump), divider))

	return response, err
}

// NewSdkTransport Creates a RoundTripper with user defined logger
func NewSdkTransport(logger Logger) *SdkTransport {
	return NewCustomSdkTransport(logger, &http.Transport{})
}

// NewCustomSdkTransport Creates a RoundTripper with user defined logger and http.Transport
func NewCustomSdkTransport(logger Logger, transport *http.Transport) *SdkTransport {
	return &SdkTransport{logger: logger, transport: transport}
}
