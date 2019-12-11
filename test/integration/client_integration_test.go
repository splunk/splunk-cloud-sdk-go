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

package integration

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/services/search"

	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/services/identity"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// This is the latest/correct client initialization to use
func getSdkClient(t *testing.T) *sdk.Client {
	client, err := MakeSdkClient(nil)
	require.Emptyf(t, err, "error calling sdk.NewClient(): %s", err)
	return client
}

func getSdkClientWithLoggers(t *testing.T) (client *sdk.Client, infoLogger *log.Logger, errLogger *log.Logger) {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	client, err := MakeSdkClient(util.NewSdkTransport(infoLogger))
	require.Emptyf(t, err, "error calling sdk.NewClient(): %s", err)
	return
}

// Get an client without the testing interface
func MakeSdkClient(rt http.RoundTripper) (*sdk.Client, error) {
	return sdk.NewClient(&services.Config{
		Token:        testutils.TestAuthenticationToken,
		Host:         testutils.TestSplunkCloudHost,
		Tenant:       testutils.TestTenant,
		Timeout:      testutils.TestTimeOut,
		RoundTripper: rt,
	})
}

// TestSDKClientTokenInit tests initializing a service-wide Splunk Cloud client and validating the token provided
func TestSDKClientInit(t *testing.T) {
	client, err := sdk.NewClient(&services.Config{
		Token:  testutils.TestAuthenticationToken,
		Host:   testutils.TestSplunkCloudHost,
		Tenant: testutils.TestTenant,
	})
	require.Emptyf(t, err, "error calling sdk.NewClient(): %s", err)
	input := identity.ValidateTokenQueryParams{Include: []identity.ValidateTokenincludeEnum{"principal", "tenant"}}
	info, err := client.IdentityService.ValidateToken(&input)
	assert.Emptyf(t, err, "error calling client.IdentityService.Validate(): %s", err)
	assert.NotNil(t, info)
}

func getClient(t *testing.T) *sdk.Client {
	client, err := sdk.NewClient(&services.Config{
		Token:   testutils.TestAuthenticationToken,
		Host:    testutils.TestSplunkCloudHost,
		Tenant:  testutils.TestTenant,
		Timeout: testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "error calling service.NewClient(): %s", err)
	return client
}

// Get a client with an invalid (expired) token
func getInvalidClient(t *testing.T) *sdk.Client {
	client, err := sdk.NewClient(&services.Config{
		Token:   testutils.ExpiredAuthenticationToken,
		Host:    testutils.TestSplunkCloudHost,
		Tenant:  testutils.TestTenant,
		Timeout: testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "error calling service.NewClient(): %s", err)
	return client
}

type noOpHandler struct {
	N int
}

func (rh *noOpHandler) HandleResponse(client *services.BaseClient, request *services.Request, response *http.Response) (*http.Response, error) {
	rh.N++
	return response, nil
}

const rHandlerErrMsg = "my custom response handler error"

type rHandlerErr struct {
	N int
}

func (rh *rHandlerErr) HandleResponse(client *services.BaseClient, request *services.Request, response *http.Response) (*http.Response, error) {
	rh.N++
	return nil, fmt.Errorf(rHandlerErrMsg)
}

// This is the legacy client initialization
// Deprecated: please use sdk.NewClient()
func TestClientMultipleResponseHandlers(t *testing.T) {
	var handler1 = &noOpHandler{}
	var handler2 = &rHandlerErr{}
	var handler3 = &noOpHandler{}
	var handlers = []services.ResponseHandler{handler1, handler2, handler3}
	client, err := sdk.NewClient(&services.Config{
		Token:            testutils.TestAuthenticationToken,
		Host:             testutils.TestSplunkCloudHost,
		Tenant:           testutils.TestInvalidTestTenant,
		Timeout:          testutils.TestTimeOut,
		ResponseHandlers: handlers,
	})
	require.Nil(t, err, "Error calling service.NewClient(): %s", err)
	query := search.ListJobsQueryParams{}.SetCount(0).SetStatus(search.SearchStatusDone)
	_, err = client.SearchService.ListJobs(&query)
	assert.True(t, strings.Contains(err.Error(), rHandlerErrMsg), "error should match custom error from response handler")
	assert.Equal(t, handler1.N, 1, "first handler should have been called")
	assert.Equal(t, handler2.N, 1, "second (error) handler should have been called")
	assert.Equal(t, handler3.N, 0, "third handler should not have been called")
}

// example to show how to create/pass RoundTripper
var LoggerOutput []string

type MyLogger struct {
}

func (ml *MyLogger) Print(v ...interface{}) {
	text, ok := v[0].(string)
	r := regexp.MustCompile(`Authorization:.*`)
	t := r.ReplaceAllString(text, "Authorization: Bearer XXX")
	if !ok {
		return
	}
	LoggerOutput = append(LoggerOutput, t)
}

func TestRoundTripperWithSdkClient(t *testing.T) {
	client, err := sdk.NewClient(&services.Config{
		Token:        testutils.TestAuthenticationToken,
		Host:         testutils.TestSplunkCloudHost,
		Tenant:       testutils.TestTenant,
		RoundTripper: util.NewSdkTransport(&MyLogger{}),
	})
	require.Nil(t, err, "Error calling service.NewClient(): %s", err)

	webhookAction := genWebhookAction()
	action, err := client.ActionService.CreateAction(webhookAction)
	require.Nil(t, err)
	defer client.ActionService.DeleteAction((*webhookAction.WebhookAction()).Name)
	require.NotEmpty(t, action)
	assert.Equal(t, 4, len(LoggerOutput))

	// verify the logged request method, url & body
	assert.Contains(t, LoggerOutput[1], "POST")
	assert.Contains(t, LoggerOutput[1], "Host:")
	assert.Contains(t, LoggerOutput[1], testutils.TestSplunkCloudHost)
	// verify log the request body
	assert.Contains(t, LoggerOutput[1], fmt.Sprintf("\"name\":\"%s\"", (*webhookAction.WebhookAction()).Name))
	assert.Contains(t, LoggerOutput[1], "\"webhookUrl\":\"https://webhook.site/test\"")

	// verify the logged request is a guid
	assert.Contains(t, LoggerOutput[1], "POST")
	assert.Regexp(t, "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}", LoggerOutput[2])
}

func TestRoundTripperWithIdentityClient(t *testing.T) {
	identityClient, err := identity.NewService(&services.Config{
		Token:        testutils.TestAuthenticationToken,
		Host:         testutils.TestSplunkCloudHost,
		Tenant:       testutils.TestTenant,
		RoundTripper: util.NewSdkTransport(&MyLogger{}),
	})
	require.Nil(t, err, "Error calling service.NewClient(): %s", err)

	LoggerOutput = LoggerOutput[:0]
	input := identity.ValidateTokenQueryParams{Include: []identity.ValidateTokenincludeEnum{"principal", "tenant"}}
	_, err = identityClient.ValidateToken(&input)
	assert.Equal(t, 4, len(LoggerOutput))
	assert.Contains(t, LoggerOutput[1], fmt.Sprintf("GET /%s/identity/v2beta1/validate?include=principal%stenant HTTP/1.1", testutils.TestTenant, "%2C"))
}

func TestRoundTripperWithInvalidClient(t *testing.T) {
	identityClient, err := identity.NewService(&services.Config{
		Token:        testutils.TestAuthenticationToken,
		Host:         "invalid.host",
		Tenant:       testutils.TestTenant,
		RoundTripper: util.NewSdkTransport(&MyLogger{}),
	})
	require.Nil(t, err, "Error calling service.NewClient(): %s", err)

	LoggerOutput = LoggerOutput[:0]
	input := identity.ValidateTokenQueryParams{Include: []identity.ValidateTokenincludeEnum{"principal", "tenant"}}
	_, err = identityClient.ValidateToken(&input)
	assert.NotNil(t, err)
	assert.Equal(t, 3, len(LoggerOutput))
	assert.Contains(t, LoggerOutput[2], "Request error")
	assert.Contains(t, LoggerOutput[2], "no such host")
}
