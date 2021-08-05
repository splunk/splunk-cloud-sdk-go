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
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"syscall"
	"testing"

	"github.com/splunk/go-dependencies/services"
	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const retrievedToken = "MY.RETRIEVED.TOKEN"

type testTokenRetriever struct{}

func (tr *testTokenRetriever) GetTokenContext() (*idp.Context, error) {
	return &idp.Context{AccessToken: retrievedToken}, nil
}

type testErrTokenRetriever struct{}

func (tr *testErrTokenRetriever) GetTokenContext() (*idp.Context, error) {
	return nil, fmt.Errorf("no luck with that token")
}

func TestAuthnResponseHandlerGoodToken(t *testing.T) {
	rh := AuthnResponseHandler{
		TokenRetriever: &testTokenRetriever{},
	}
	ctx, err := rh.TokenRetriever.GetTokenContext()
	assert.NoError(t, err)
	assert.NotNil(t, ctx)
	assert.Equal(t, ctx.AccessToken, retrievedToken)
}

func TestAuthnResponseHandlerErrorToken(t *testing.T) {
	rh := AuthnResponseHandler{
		TokenRetriever: &testErrTokenRetriever{},
	}
	ctx, err := rh.TokenRetriever.GetTokenContext()
	assert.NotNil(t, err)
	assert.Nil(t, ctx)
}

// respHandler is a ResponseHandler but not a ResponseOrErrorHandler
type respHandler struct {
	NResp int
}

func (rh *respHandler) HandleResponse(client *BaseClient, request *Request, response *http.Response) (*http.Response, error) {
	rh.NResp++
	// Try a new request if >= 400
	if response.StatusCode >= 400 {
		return client.Do(request)
	}
	return response, nil
}

type test429RT struct {
	N int
}

// RoundTrip for this test429RT provides 429 for three attempts, then success
func (rt *test429RT) RoundTrip(req *http.Request) (*http.Response, error) {
	rt.N++
	b := ioutil.NopCloser(bytes.NewReader([]byte("")))
	switch rt.N {
	case 1, 2, 3:
		return &http.Response{Status: "429 Too Many Requests", StatusCode: 429, Body: b}, nil
	}
	return &http.Response{Status: "200 OK", StatusCode: 200, Body: b}, nil
}

func TestClientResponseHandler(t *testing.T) {
	var handler1 = &respHandler{}
	// Verify that this is not a ResponseOrErrorHandler
	verify := func(hr ResponseHandler) {
		_, ok := hr.(ResponseOrErrorHandler)
		require.False(t, ok, "Handler provided was should not be a ResponseOrErrorHandler")
	}
	verify(handler1)
	handlers := []ResponseHandler{handler1}
	rt := &test429RT{}
	client, err := NewClient(&Config{
		Token:            "testtoken",
		ResponseHandlers: handlers,
		RoundTripper:     rt,
		RetryConfig: RetryStrategyConfig{
			ConfigurableRetryConfig: &ConfigurableRetryConfig{
				RetryNum: 4,
				Interval: 10, // 10 ms so tests execute quickly
			},
		},
	})
	require.Nil(t, err, "Error calling NewClient(): %s", err)
	resp, err := client.Get(services.RequestParams{})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 4, handler1.NResp, "first handler's HandleResponse() should have been called three times for the 429 responses and once for 200")
}

// errAndRespHandler implements ResponseOrErrorHandler and retries any requests that err
// while returning the first response
type errAndRespHandler struct {
	NErr  int
	NResp int
}

func (rh *errAndRespHandler) HandleRequestError(client *BaseClient, request *Request, err error) (*http.Response, error) {
	rh.NErr++
	// Try a new request
	return client.Get(services.RequestParams{})
}

func (rh *errAndRespHandler) HandleResponse(client *BaseClient, request *Request, response *http.Response) (*http.Response, error) {
	rh.NResp++
	return response, nil
}

type testRT struct {
	N int
}

// RoundTrip for this testRT provides a connection timeout for first two attempts,
// then 429, then success
func (rt *testRT) RoundTrip(req *http.Request) (*http.Response, error) {
	rt.N++
	b := ioutil.NopCloser(bytes.NewReader([]byte("")))
	switch rt.N {
	case 1, 2:
		return nil, syscall.ECONNRESET
	case 3:
		return &http.Response{Status: "429 Too Many Requests", StatusCode: 429, Body: b}, nil
	}
	return &http.Response{Status: "200 OK", StatusCode: 200, Body: b}, nil
}

func TestClientErrorAndResponseHandler(t *testing.T) {
	var handler1 = &errAndRespHandler{}
	// Verify that this is in fact a ResponseOrErrorHandler
	verify := func(hr ResponseHandler) {
		_, ok := hr.(ResponseOrErrorHandler)
		require.True(t, ok, "Handler provided was not ResponseOrErrorHandler")
	}
	verify(handler1)
	handlers := []ResponseHandler{handler1}
	rt := &testRT{}
	client, err := NewClient(&Config{
		Token:            "testtoken",
		ResponseHandlers: handlers,
		RoundTripper:     rt,
	})
	require.Nil(t, err, "Error calling NewClient(): %s", err)
	resp, err := client.Get(services.RequestParams{})
	require.NotNil(t, err)
	assert.Nil(t, resp)
	he, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, "429 Too Many Requests", he.HTTPStatus)
	assert.Equal(t, 429, he.HTTPStatusCode)
	assert.Equal(t, 2, handler1.NErr, "first handler's HandleRequestError() should have been called TWICE for connection resets")
	assert.Equal(t, 1, handler1.NResp, "first handler's HandleResponse() should have been called ONCE for the 429 response")
}

func TestClientConnectionResetRetry(t *testing.T) {
	rt := &testRT{}
	client, err := NewClient(&Config{
		Token:         "testtoken",
		RetryRequests: true,
		RetryConfig: RetryStrategyConfig{
			ConfigurableRetryConfig: &ConfigurableRetryConfig{
				RetryNum: 6,
				Interval: 10, // 10 ms so tests execute quickly
			},
		},
		RoundTripper: rt,
	})
	require.Nil(t, err, "Error calling NewClient(): %s", err)
	require.Equal(t, 1, len(client.responseHandlers))
	// Verify that this is in fact a ResponseOrErrorHandler
	verify := func(hr ResponseHandler) {
		_, ok := hr.(ResponseOrErrorHandler)
		require.True(t, ok, "Handler provided was not ResponseOrErrorHandler")
	}
	verify(client.responseHandlers[0])
	resp, err := client.Get(services.RequestParams{})
	assert.Nil(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "200 OK", resp.Status)
	assert.Equal(t, 200, resp.StatusCode)
	// RoundTripper should be called a total of four times: twice for connection reset errors
	// being retried, once for a 429 response being retried, then one successful 200 OK response
	assert.Equal(t, 4, rt.N, "RoundTripper should have been called 4 times")
}

func TestClientCustomRetryFunc(t *testing.T) {
	rt := &testRT{}
	retryFn := func(request *Request, reqErr error, response *http.Response, maxRetries uint) bool {
		// Provide custom retry logic to override the maxRetries specified in the config and stop after 2 retries (3 total requests)
		return request.NumAttempts < 3
	}
	client, err := NewClient(&Config{
		Token:         "testtoken",
		RetryRequests: true,
		RetryConfig: RetryStrategyConfig{
			ConfigurableRetryConfig: &ConfigurableRetryConfig{
				RetryNum:      6,  // This should be ignored by our custom ShouldRetryFn above
				Interval:      10, // 10 ms so tests execute quickly
				ShouldRetryFn: retryFn,
			},
		},
		RoundTripper: rt,
	})
	require.Nil(t, err, "Error calling NewClient(): %s", err)
	require.Equal(t, 1, len(client.responseHandlers))
	// Verify that this is in fact a ResponseOrErrorHandler
	verify := func(hr ResponseHandler) {
		_, ok := hr.(ResponseOrErrorHandler)
		require.True(t, ok, "Handler provided was not ResponseOrErrorHandler")
	}
	verify(client.responseHandlers[0])
	resp, err := client.Get(services.RequestParams{})
	require.NotNil(t, err)
	assert.Nil(t, resp)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	// The second retry should result in 429
	assert.Equal(t, "429 Too Many Requests", httpErr.HTTPStatus)
	assert.Equal(t, 429, httpErr.HTTPStatusCode)
	// RoundTripper should be called a total of three times due to custom retry func
	assert.Equal(t, 3, rt.N, "RoundTripper should have been called 3 times")
}
