// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package util

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// Logger define logger interface for roundtripper
type Logger interface {
	Debug(string)
}

// SdkTransport is to define a transport RoundTripper with user-defined logger
type SdkTransport struct {
	transport http.RoundTripper
	logger    Logger
}

// RoundTrip implements the RoundTripper interface
func (st *SdkTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		return nil, err
	}

	st.logger.Debug(fmt.Sprintf("===Request:\n%s\n", string(requestDump)))

	response, err := st.transport.RoundTrip(request)

	responseDump, err := httputil.DumpResponse(response, true)
	if err != nil {
		return response, err
	}

	st.logger.Debug(fmt.Sprintf("===Response:\n%s\n", string(responseDump)))

	return response, err
}

// CreateRoundTripperWithLogger Creates a RoundTripper with user defined logger
func CreateRoundTripperWithLogger(logger Logger) *SdkTransport {
	return &SdkTransport{transport: &http.Transport{}, logger: logger}
}
