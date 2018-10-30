// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/service"
	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/services/identity"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// This is the latest/correct client initialization to use
func getSdkClient(t *testing.T) *sdk.Client {
	client, err := sdk.NewClient(&services.Config{
		Token:   testutils.TestAuthenticationToken,
		Host:    testutils.TestSplunkCloudHost,
		Tenant:  testutils.TestTenant,
		Timeout: testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "error calling sdk.NewClient(): %s", err)
	return client
}

// TestSDKClientTokenInit tests initializing a service-wide Splunk Cloud client and validating the token provided
func TestSDKClientInit(t *testing.T) {
	client, err := sdk.NewClient(&services.Config{
		Token:  testutils.TestAuthenticationToken,
		Host:   testutils.TestSplunkCloudHost,
		Tenant: "system",
	})
	require.Emptyf(t, err, "error calling sdk.NewClient(): %s", err)
	info, err := client.IdentityService.Validate()
	assert.Emptyf(t, err, "error calling client.IdentityService.Validate(): %s", err)
	assert.NotNil(t, info)
}

// This is the legacy client initialization
// Deprecated: please use sdk.NewClient()
func getClient(t *testing.T) *service.Client {
	client, err := service.NewClient(&service.Config{
		Token:   testutils.TestAuthenticationToken,
		Host:    testutils.TestSplunkCloudHost,
		Tenant:  testutils.TestTenant,
		Timeout: testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "error calling service.NewClient(): %s", err)
	return client
}

// This is the legacy client initialization
// Deprecated: please use sdk.NewClient()
func getInvalidTenantClient(t *testing.T) *service.Client {
	client, err := service.NewClient(&service.Config{
		Token:   testutils.TestAuthenticationToken,
		Host:    testutils.TestSplunkCloudHost,
		Tenant:  testutils.TestInvalidTestTenant,
		Timeout: testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "error calling service.NewClient(): %s", err)
	return client
}

// This is the legacy client initialization
// Deprecated: please use sdk.NewClient()
func getInvalidClient(t *testing.T) *service.Client {
	client, err := service.NewClient(&service.Config{
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

func (rh *noOpHandler) HandleResponse(client *services.BaseClient, request *service.Request, response *http.Response) (*http.Response, error) {
	rh.N++
	return response, nil
}

const rHandlerErrMsg = "my custom response handler error"

type rHandlerErr struct {
	N int
}

func (rh *rHandlerErr) HandleResponse(client *services.BaseClient, request *service.Request, response *http.Response) (*http.Response, error) {
	rh.N++
	return nil, fmt.Errorf(rHandlerErrMsg)
}

// This is the legacy client initialization
// Deprecated: please use sdk.NewClient()
func TestClientMultipleResponseHandlers(t *testing.T) {
	var handler1 = &noOpHandler{}
	var handler2 = &rHandlerErr{}
	var handler3 = &noOpHandler{}
	var handlers = []service.ResponseHandler{handler1, handler2, handler3}
	client, err := service.NewClient(&service.Config{
		Token:            testutils.TestAuthenticationToken,
		Host:             testutils.TestSplunkCloudHost,
		Tenant:           testutils.TestInvalidTestTenant,
		Timeout:          testutils.TestTimeOut,
		ResponseHandlers: handlers,
	})
	require.Nil(t, err, "Error calling service.NewClient(): %s", err)
	_, err = client.SearchService.ListJobs()
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
	if !ok {
		return
	}
	LoggerOutput = append(LoggerOutput, text)
}

func TestRoundTripperWithSdkClient(t *testing.T) {
	client, err := sdk.NewClient(&service.Config{
		Token:        testutils.TestAuthenticationToken,
		Host:         testutils.TestSplunkCloudHost,
		Tenant:       testutils.TestTenant,
		RoundTripper: util.CreateRoundTripperWithLogger(&MyLogger{}),
	})
	require.Nil(t, err, "Error calling service.NewClient(): %s", err)

	webhookActionName := "testaction"
	webhookAction := model.NewWebhookAction(webhookActionName, webhookURL, webhookMsg)
	action, err := client.ActionService.CreateAction(*webhookAction)
	defer client.ActionService.DeleteAction(webhookActionName)
	require.Nil(t, err)
	require.NotEmpty(t, action)
	assert.Equal(t, 2, len(LoggerOutput))

	// verify log the request method and url
	assert.Contains(t, LoggerOutput[0], "POST")
	assert.Contains(t, LoggerOutput[0], "Host:")
	// verify log the request body
	assert.Contains(t, LoggerOutput[0], "\"name\":\"testaction\"")
	assert.Contains(t, LoggerOutput[0], "\"webhookUrl\":\"https://webhook.site/test\"")

	// verify log the response
	assert.Contains(t, LoggerOutput[1], "\"name\":\"testaction\"")
	assert.Contains(t, LoggerOutput[1], "\"webhookUrl\":\"https://webhook.site/test\"")
}

func TestRoundTripperWithIdentityClient(t *testing.T) {
	identityClient, _ := identity.NewService(&services.Config{
		Token:        testutils.TestAuthenticationToken,
		Host:         testutils.TestSplunkCloudHost,
		Tenant:       "system",
		RoundTripper: util.CreateRoundTripperWithLogger(&MyLogger{}),
	})

	LoggerOutput = LoggerOutput[:0]
	identityClient.Validate()
	assert.Equal(t, 2, len(LoggerOutput))
	assert.Contains(t, LoggerOutput[0], "GET /system/identity/v1/validate HTTP/1.1")
}
