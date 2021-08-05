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

package services

import (
	"errors"
	"net/http"
	"syscall"

	"time"

	"github.com/splunk/splunk-cloud-sdk-go/idp"
)

const (
	// DefaultMaxAuthnAttempts defines the maximum number of retries that will be performed for a request encountering an authentication issue
	DefaultMaxAuthnAttempts = 1
)

// ShouldRetry returns whether or not the request should be retried based on the original
// request, request error received, response received, and max number of retries specified in the config
type ShouldRetry func(request *Request, reqErr error, response *http.Response, maxRetries uint) bool

//ConfigurableRetryConfig that will accept a user configurable RetryNumber and Interval between retries
type ConfigurableRetryConfig struct {
	// RetryNum defines the number of attempts to retry, if attempts > RetryNum then no more retries are attempted
	RetryNum uint
	// Interval defines the interval in milliseconds for exponential backoff where backoff = interval * 2^(retry num)
	Interval int
	// ShouldRetryFn defines a custom function to determine whether or not to retry the request
	ShouldRetryFn ShouldRetry
}

//DefaultRetryConfig that will use a default RetryNumber and a default Interval between retries
type DefaultRetryConfig struct {
}

// These are the values used for the DefaultRetryConfig
const (
	defaultMaxRetryCount  = 6
	defaultIntervalMillis = 500
)

func defaultShouldRetryFn(request *Request, reqErr error, response *http.Response, maxRetries uint) bool {
	// Regardless of other conditions if we're over the retry count, we should not retry
	if request != nil && request.NumAttempts > maxRetries {
		return false
	}
	// If response is 429, retry
	if response != nil && response.StatusCode == 429 {
		return true
	}
	// If connection reset error encountered, retry
	if reqErr != nil && errors.Is(reqErr, syscall.ECONNRESET) {
		return true
	}
	return false
}

//RetryStrategyConfig to be specified while creating a NewClient
type RetryStrategyConfig struct {
	DefaultRetryConfig      *DefaultRetryConfig
	ConfigurableRetryConfig *ConfigurableRetryConfig
}

// ResponseHandler defines the interface for implementing custom response
// handling logic, request errors are not handled - implement ResponseOrErrorHandler for
// handling of request errors
type ResponseHandler interface {
	HandleResponse(client *BaseClient, request *Request, response *http.Response) (*http.Response, error)
}

// ResponseOrErrorHandler defines the interface for implementing custom response
// handling logic including when response is returned for HandleResponse and
// when request returns with an error for HandleError
type ResponseOrErrorHandler interface {
	HandleRequestError(client *BaseClient, request *Request, err error) (*http.Response, error)
	HandleResponse(client *BaseClient, request *Request, response *http.Response) (*http.Response, error)
}

// AuthnResponseHandler handles logic for updating the client access token in response to 401 errors
type AuthnResponseHandler struct {
	TokenRetriever idp.TokenRetriever
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
	if err != nil {
		return nil, err
	}
	request.Body = body
	// Update the client such that future requests will use the new access token and retain context information
	client.UpdateTokenContext(ctx)
	// Retry the request with the updated token
	return client.Do(request)
}

// DefaultRetryResponseHandler handles logic for retrying requests with default settings for Retry number and interval
type DefaultRetryResponseHandler struct {
	DefaultRetryConfig DefaultRetryConfig
}

// HandleResponse will retry a request once a 429 is encountered using a Default exponential BackOff Retry Strategy
func (defRh DefaultRetryResponseHandler) HandleResponse(client *BaseClient, request *Request, response *http.Response) (*http.Response, error) {
	return handleRequestResponse(client, request, nil, response, defaultMaxRetryCount, defaultIntervalMillis, defaultShouldRetryFn)
}

// HandleRequestError will retry a request if a connection reset is encountered using a
// Default exponential BackOff Retry Strategy
func (defRh DefaultRetryResponseHandler) HandleRequestError(client *BaseClient, request *Request, err error) (*http.Response, error) {
	return handleRequestResponse(client, request, err, nil, defaultMaxRetryCount, defaultIntervalMillis, defaultShouldRetryFn)
}

// ConfigurableRetryResponseHandler handles logic for retrying requests with user configurable
// settings for Retry number and interval
type ConfigurableRetryResponseHandler struct {
	ConfigurableRetryConfig ConfigurableRetryConfig
}

// HandleResponse will retry a request if a 429 is encountered using a configurable exponential
// BackOff Retry Strategy
func (configRh ConfigurableRetryResponseHandler) HandleResponse(client *BaseClient, request *Request, response *http.Response) (*http.Response, error) {
	retryFn := configRh.ConfigurableRetryConfig.ShouldRetryFn
	if retryFn == nil {
		retryFn = defaultShouldRetryFn
	}
	return handleRequestResponse(client, request, nil, response, configRh.ConfigurableRetryConfig.RetryNum, configRh.ConfigurableRetryConfig.Interval, retryFn)
}

// HandleRequestError will retry a request once a connection reset is encountered using
// a Configurable exponential BackOff Retry Strategy
func (configRh ConfigurableRetryResponseHandler) HandleRequestError(client *BaseClient, request *Request, err error) (*http.Response, error) {
	retryFn := configRh.ConfigurableRetryConfig.ShouldRetryFn
	if retryFn == nil {
		retryFn = defaultShouldRetryFn
	}
	return handleRequestResponse(client, request, err, nil, configRh.ConfigurableRetryConfig.RetryNum, configRh.ConfigurableRetryConfig.Interval, retryFn)
}

//handleRequestResponse - helper function to handle the retry to a 429 response
func handleRequestResponse(client *BaseClient, request *Request, reqErr error, response *http.Response, maxRetries uint, interval int, shouldRetry ShouldRetry) (*http.Response, error) {
	if request == nil {
		return response, reqErr // can't retry the request without it
	}
	// Is retry warranted? If not, return
	if !shouldRetry(request, reqErr, response, maxRetries) {
		return response, reqErr
	}
	// implement exponential back off by increasing the waiting time between retries after each retry failure.
	backoffMillis := time.Duration((1<<request.NumAttempts)*interval) * time.Millisecond
	time.Sleep(backoffMillis)

	// reinitialize body, otherwise it will be empty
	if request.Body != nil {
		body, err := request.GetBody()
		if err != nil {
			return nil, reqErr
		}
		request.Body = body
	}
	return client.Do(request)
}
