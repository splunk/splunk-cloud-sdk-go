package stubbyintegration

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"fmt"
	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/testutils"
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

	client, err := service.NewClient(testutils.TestTenantID, testutils.TestAuthenticationToken, url, testutils.TestTimeOut)
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
		req, err := client.NewRequest(test.method, test.url, test.body)
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
	req, err := client.NewRequest(http.MethodGet, testutils.TestSSCHost, nil)
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
	_, err := client.NewRequest("#~/", testutils.TestSSCHost, nil)
	if err == nil {
		t.Errorf("NewRequest expected to return error, got %v", err)
	}
}

func TestNewStubbyRequest(t *testing.T) {
	client := getClient(t)
	resp, _ := client.DoRequest(http.MethodGet, url.URL{Scheme: testutils.TestURLProtocol, Host: testutils.TestSSCHost, Path: "/error"}, nil)
	if resp.StatusCode != 500 {
		t.Fatalf("client.DoRequest to /error endpoint expected Response Code: %d, Received: %d", 500, resp.StatusCode)
	}
	defer resp.Body.Close()
	b := new(bytes.Buffer)
	b.ReadFrom(resp.Body)
	content := new(map[string]string)
	if err := json.NewDecoder(b).Decode(content); err != nil {
		t.Fatalf("client.DoRequest error unmarshalling response, err: %v", err)
	}
	if (*content)["message"] != "Something exploded" {
		t.Fatalf("client.DoRequest error/ expecting response {\"message\":\"Something exploded\"} Received: %+v", content)
	}
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
