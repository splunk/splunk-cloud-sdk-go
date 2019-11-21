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

// This example demonstrates how to setup the Splunk Cloud SDK for Go to return mocked
// responses rather than sending requests to the Splunk Cloud Platform itself.
//
// Invoke this example using ```$ go run -v ./examples/mock/mock.go```
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/services/identity"
	"github.com/splunk/splunk-cloud-sdk-go/services/search"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

func main() {
	client, err := sdk.NewClient(&services.Config{
		Tenant:       "foo",
		Token:        "faked.token",
		RoundTripper: mockTransport{},
	})
	exitOnError(err)
	var resp http.Response
	fmt.Println("Sending request to mocked IdentityService.ValidateToken() endpoint ...")
	info, err := client.IdentityService.ValidateToken(nil, &resp)
	// MOCKING   [/foo/identity/v2beta1/validate]
	exitOnError(err)
	fmt.Printf("%-10s[%d] %+v\n\n", "RESPONSE", resp.StatusCode, info)
	// RESPONSE  [200] &{ClientId: Name:someone@domain.com Principal:<nil> Tenant:<nil>}
	fmt.Println("Sending request to unmocked SearchService.CreateJob() endpoint ...")
	job := search.SearchJob{Query: "| from index:main | head 5"}
	_, err = client.SearchService.CreateJob(job, &resp)
	// MOCKING   [/foo/search/v2beta1/jobs]
	// Expecting an err this time
	if httpErr, ok := err.(*util.HTTPError); ok {
		fmt.Printf("%-10s[%d] %+v\n\n", "RESPONSE", resp.StatusCode, httpErr)
		// RESPONSE  [404] {"HTTPStatusCode":404,"HTTPStatus":"Not Found","message":"endpoint not found","code":"route_not_found"}
		fmt.Printf("Success.")
		os.Exit(0)
	}
	exitOnError(fmt.Errorf("expected to receive a 404 error but none was returned"))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// mockTransport defines a RoundTripper to be added to our services.Config
// which intercepts requests from the client and returns its own mocked
// responses
type mockTransport struct{}

// RoundTrip is the method needed to implement the http.RoundTripper interface.
// This is where mocks for various endpoints are implemented.
func (m mockTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	resp := &http.Response{}
	fmt.Printf("%-10s[%s]\n", "MOCKING", request.URL.EscapedPath())
	switch request.URL.EscapedPath() {
	// Define our mocked responses here, for request URLs that we haven't
	// mocked, return a 404 (http.StatusNotFound) response
	case "/foo/identity/v2beta1/validate":
		info := identity.ValidateInfo{
			Name: "someone@domain.com",
		}
		b, _ := json.Marshal(&info)
		// Return a 200 with our ValidateInfo model contents defined above
		resp.StatusCode = http.StatusOK
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		return resp, nil
	default:
		resp.StatusCode = http.StatusNotFound
		resp.Status = "Not Found"
		resp.Body = ioutil.NopCloser(bytes.NewBufferString(`{"code":"route_not_found", "message":"endpoint not found"}`))
		return resp, nil
	}
}
