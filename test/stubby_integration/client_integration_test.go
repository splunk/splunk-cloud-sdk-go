// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package stubbyintegration

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/splunk/ssc-client-go/util"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func getClient(t *testing.T) *service.Client {
	var url = testutils.TestURLProtocol + "://" + testutils.TestSSCHost

	//fmt.Printf("=================================================================")
	//fmt.Printf("CREATING A CLIENT WITH THESE SETTINGS")
	//fmt.Printf("=================================================================")
	//fmt.Printf("Authentication Token: " + testutils.TestAuthenticationToken + "\n")
	//fmt.Printf("SSC Host API: " + testutils.TestSSCHost + "\n")
	//fmt.Printf("Tenant ID: " + testutils.TestTenantID + "\n")
	//fmt.Printf("URL Protocol: " + testutils.TestURLProtocol + "\n")
	//fmt.Printf("Fully Qualified URL: " + url + "\n")

	client, err := service.NewClient(&service.Config{Token: testutils.TestAuthenticationToken, URL: url, TenantID: testutils.TestTenantID, Timeout: testutils.TestTimeOut})
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}

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
		{http.MethodGet, testutils.TestSSCHost, nil},
		{http.MethodPost, testutils.TestSSCHost, requestBody},
		{http.MethodPut, testutils.TestSSCHost, requestBody},
		{http.MethodPatch, testutils.TestSSCHost, requestBody},
		{http.MethodDelete, testutils.TestSSCHost, nil},
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
	req, err := client.NewRequest(http.MethodGet, testutils.TestSSCHost, nil, nil)
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
	_, err := client.NewRequest("#~/", testutils.TestSSCHost, nil, nil)
	if err == nil {
		t.Errorf("NewRequest expected to return error, got %v", err)
	}
}

func TestNewStubbyRequest(t *testing.T) {
	client := getClient(t)
	resp, err := client.DoRequest(service.RequestParams{Method: http.MethodGet, URL: url.URL{Scheme: testutils.TestURLProtocol, Host: testutils.TestSSCHost, Path: "/error"}})
	defer resp.Body.Close()

	assert.NotNil(t, err)

	assert.Equal(t, 500, resp.StatusCode)
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