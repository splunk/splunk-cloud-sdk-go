// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package services

import (
	"net/http"

	"time"

	"github.com/splunk/splunk-cloud-sdk-go/idp"
)

const (
	// DefaultMaxAuthnAttempts defines the maximum number of retries that will be performed for a request encountering an authentication issue
	DefaultMaxAuthnAttempts = 1
)

const (
	maxRetryCount = 6
)

// ResponseHandler defines the interface for implementing custom response
// handling logic
type ResponseHandler interface {
	HandleResponse(client *BaseClient, request *Request, response *http.Response) (*http.Response, error)
}

// AuthnResponseHandler handles logic for updating the client access token in response to 401 errors
type AuthnResponseHandler struct {
	TokenRetriever idp.TokenRetriever
}

// ConfigurableRetryResponseHandler handles logic for retrying requests with user configurable settings for Retry number and interval
type ConfigurableRetryResponseHandler struct {
	ConfigurableRetryConfig ConfigurableRetryConfig
}

// DefaultRetryResponseHandler handles logic for retrying requests with default settings for Retry number and interval
type DefaultRetryResponseHandler struct {
	DefaultRetryConfig DefaultRetryConfig
}

// HandleResponse will retry a request once after re-authenticating if a 401 response code is encountered
func (rh AuthnResponseHandler) HandleResponse(client *BaseClient, request *Request, response *http.Response) (*http.Response, error) {
	if response.StatusCode != 401 || rh.TokenRetriever == nil || request.GetNumErrorsByResponseCode(401) > DefaultMaxAuthnAttempts {
		return response, nil
	}
	ctx, err := rh.TokenRetriever.GetTokenContext()
	if err != nil {
		return response, err
	}
	// Replace the access token in the request's Authorization: Bearer header
	request.UpdateToken(ctx.AccessToken)
	// Re-initialize body (otherwise body is empty)
	body, err := request.GetBody()
	request.Body = body
	// Update the client such that future requests will use the new access token and retain context information
	client.UpdateTokenContext(ctx)
	// Retry the request with the updated token
	return client.Do(request)
}

// HandleResponse will retry a request once a 429 is encountered using a Default exponential BackOff Retry Strategy
func (defRh DefaultRetryResponseHandler) HandleResponse(client *BaseClient, request *Request, response *http.Response) (*http.Response, error) {
	return handleRequestResponse(client, request, response, maxRetryCount, 500)
}

// HandleResponse will retry a request once a 429 is encountered using a Configurable exponential BackOff Retry Strategy
func (configRh ConfigurableRetryResponseHandler) HandleResponse(client *BaseClient, request *Request, response *http.Response) (*http.Response, error) {
	return handleRequestResponse(client, request, response, configRh.ConfigurableRetryConfig.RetryNum, configRh.ConfigurableRetryConfig.Interval)
}

//handleRequestResponse - helper function to handle the retry to a 429 response
func handleRequestResponse(client *BaseClient, request *Request, response *http.Response, retryCount uint, interval int) (*http.Response, error) {
	if response.StatusCode != 429 {
		return response, nil
	}
	if request.NumAttempts > retryCount {
		return response, nil
	}
	//implement exponential back off by increasing the waiting time between retries after each retry failure.
	backoffMillis := time.Duration((1<<request.NumAttempts)*interval) * time.Millisecond
	time.Sleep(backoffMillis)

	// reinitialize body, otherwise it will be empty
	body, err := request.GetBody()
	if err != nil {
		return nil, err
	}
	request.Body = body
	return client.Do(request)
}
