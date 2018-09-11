// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package stubbyintegration

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/service"
	"github.com/splunk/splunk-cloud-sdk-go/testutils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

func getInvalidTokenClient(t *testing.T) *service.Client {
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

func TestNewRequest(t *testing.T) {
	client := getClient(t)
	body := []byte(`{"test":"This is a test body"}`)
	expectedAuth := []string{"Bearer " + testutils.TestAuthenticationToken}
	requestBody := bytes.NewBuffer(body)
	tests := []struct {
		method string
		url    string
		body   io.Reader
	}{
		{http.MethodGet, testutils.TestSplunkCloudHost, nil},
		{http.MethodPost, testutils.TestSplunkCloudHost, requestBody},
		{http.MethodPut, testutils.TestSplunkCloudHost, requestBody},
		{http.MethodPatch, testutils.TestSplunkCloudHost, requestBody},
		{http.MethodDelete, testutils.TestSplunkCloudHost, nil},
	}
	for _, test := range tests {
		req, err := client.NewRequest(test.method, test.url, test.body, nil)
		if err != nil {
			t.Fatalf("client.NewRequest returns unexpected error: %v", err)
		}
		if got, want := req.Method, test.method; got != want {
			t.Errorf("NewRequest http method is %v, want %v", got, want)
		}
		if got, want := req.URL.String(), test.url; got != want {
			t.Errorf("NewRequest url is %v, want %v", got, want)
		}
		if got, want := req.Header["Authorization"], expectedAuth; !reflect.DeepEqual(got, want) {
			t.Errorf("NewRequest authorization is %v, want %v", got, want)
		}
		if test.method == http.MethodGet || test.method == http.MethodDelete {
			t.Skipf("Skip NewRequest body test for %v and %v method", http.MethodGet, http.MethodDelete)
		} else {
			gotBody, _ := ioutil.ReadAll(req.Body)
			if bytes.Compare(gotBody, body) != -1 {
				t.Errorf("NewRequest url is %v, want %v", gotBody, body)
			}
		}
	}
}

func TestNewRequestBearerAuthHeader(t *testing.T) {
	client := getClient(t)
	req, err := client.NewRequest(http.MethodGet, testutils.TestSplunkCloudHost, nil, nil)
	if err != nil {
		t.Errorf("NewRequest returns unexpected error %v", err)
	}
	expectedAuth := []string{"Bearer " + testutils.TestAuthenticationToken}
	if got, want := req.Header["Authorization"], expectedAuth; !reflect.DeepEqual(got, want) {
		t.Errorf("NewRequest authorization is %v, want %v", got, want)
	}
}

func TestNewRequestError(t *testing.T) {
	client := getClient(t)
	_, err := client.NewRequest("#~/", testutils.TestSplunkCloudHost, nil, nil)
	if err == nil {
		t.Errorf("NewRequest expected to return error, got %v", err)
	}
}

func TestNewStubbyRequest(t *testing.T) {
	client := getClient(t)
	resp, err := client.DoRequest(service.RequestParams{Method: http.MethodGet, URL: url.URL{Scheme: testutils.TestURLProtocol, Host: testutils.TestSplunkCloudHost, Path: "/error"}})
	defer resp.Body.Close()

	assert.NotNil(t, err)

	assert.Equal(t, 500, resp.StatusCode)

	assert.Equal(t, err.(*util.HTTPError).Code, "1234")
	assert.Equal(t, err.(*util.HTTPError).MoreInfo, "/url/test")
	assert.Equal(t, err.(*util.HTTPError).Message, "error response")
	assert.Equal(t, err.(*util.HTTPError).Details[0]["code"], "123")
	assert.Equal(t, err.(*util.HTTPError).Details[0]["field"], "username")
	assert.Equal(t, err.(*util.HTTPError).Details[0]["message"], "Username must be at least 8 characters")
}

func TestNewBatchEventsSenderState(t *testing.T) {
	var client = getClient(t)
	collector, err := client.NewBatchEventsSender(5, 1000)
	assert.Nil(t, err)

	// Initial queue should
	assert.Equal(t, 0, len(collector.EventsQueue))
	assert.Equal(t, 5, cap(collector.EventsQueue))
	assert.Equal(t, 0, len(collector.EventsChan))
	assert.Equal(t, 5, cap(collector.EventsChan))
	assert.Equal(t, 0, len(collector.QuitChan))
	assert.Equal(t, 1, cap(collector.QuitChan))
	assert.Equal(t, 5, collector.BatchSize)
}
