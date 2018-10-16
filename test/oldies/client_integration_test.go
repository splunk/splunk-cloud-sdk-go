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

	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/service"
	"github.com/splunk/splunk-cloud-sdk-go/services"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// This is the latest/correct client initialization to use
func getSdkClient(t *testing.T) *sdk.Client {
	client, err := sdk.NewClient(&services.Config{
		Token:   testutils.TestAuthenticationToken,
		Scheme:  testutils.TestURLProtocol,
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
		Scheme:  testutils.TestURLProtocol,
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
		Scheme:  testutils.TestURLProtocol,
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
		Scheme:  testutils.TestURLProtocol,
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
		Scheme:           testutils.TestURLProtocol,
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