// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package playgroundintegration

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/service"
	"github.com/splunk/splunk-cloud-sdk-go/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getClient(t *testing.T) *service.Client {
	var url = testutils.TestURLProtocol + "://" + testutils.TestSplunkCloudHost

	//fmt.Printf("=================================================================")
	//fmt.Printf("CREATING A CLIENT WITH THESE SETTINGS")
	//fmt.Printf("=================================================================")
	//fmt.Printf("Authentication Token: " + testutils.TestAuthenticationToken + "\n")
	//fmt.Printf("Splunk Cloud Host API: " + testutils.TestSplunkCloudHost + "\n")
	//fmt.Printf("Tenant ID: " + testutils.TestTenantID + "\n")
	//fmt.Printf("URL Protocol: " + testutils.TestURLProtocol + "\n")
	//fmt.Printf("Fully Qualified URL: " + url + "\n")
	client, err := service.NewClient(&service.Config{Token: testutils.TestAuthenticationToken, URL: url, TenantID: testutils.TestTenantID, Timeout: testutils.TestTimeOut})
	require.Emptyf(t, err, "Error calling service.NewClient(): %s", err)
	return client
}

func getInvalidClient(t *testing.T) *service.Client {
	var url = testutils.TestURLProtocol + "://" + testutils.TestSplunkCloudHost

	client, err := service.NewClient(&service.Config{Token: testutils.TestInvalidAuthenticationToken, URL: url, TenantID: testutils.TestTenantID, Timeout: testutils.TestTimeOut})
	require.Emptyf(t, err, "Error calling service.NewClient(): %s", err)
	return client
}

type noOpHandler struct {
	N int
}

func (rh *noOpHandler) HandleResponse(client *service.Client, request *service.Request, response *http.Response) (*http.Response, error) {
	rh.N++
	return response, nil
}

const rHandlerErrMsg = "my custom response handler error"

type rHandlerErr struct {
	N int
}

func (rh *rHandlerErr) HandleResponse(client *service.Client, request *service.Request, response *http.Response) (*http.Response, error) {
	rh.N++
	return nil, fmt.Errorf(rHandlerErrMsg)
}

func TestClientMultipleResponseHandlers(t *testing.T) {
	var url = testutils.TestURLProtocol + "://" + testutils.TestSplunkCloudHost
	var handler1 = &noOpHandler{}
	var handler2 = &rHandlerErr{}
	var handler3 = &noOpHandler{}
	var handlers = []service.ResponseHandler{handler1, handler2, handler3}
	client, err := service.NewClient(&service.Config{Token: testutils.TestAuthenticationToken, URL: url, TenantID: testutils.TestInvalidTestTenantID, Timeout: testutils.TestTimeOut, ResponseHandlers: handlers})
	require.Nil(t, err, "Error calling service.NewClient(): %s", err)
	_, err = client.SearchService.GetJobs(nil)
	assert.True(t, strings.Contains(err.Error(), rHandlerErrMsg), "error should match custom error from response handler")
	assert.Equal(t, handler1.N, 1, "first handler should have been called")
	assert.Equal(t, handler2.N, 1, "second (error) handler should have been called")
	assert.Equal(t, handler3.N, 0, "third handler should not have been called")
}
