package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	. "github.com/splunk/ssc-client-go/util"
	"github.com/stretchr/testify/assert"
)

func getClient() *Client {
	return NewClient(TestTenantID, TestToken, TestHost, TestTimeOut)
}

func TestBuildURL(t *testing.T) {
	client := getClient()
	url, err := client.BuildURL("services", "search", "jobs")

	assert.Nil(t, err)
	assert.Equal(t, "localhost", url.Hostname())
	assert.Equal(t, "https", url.Scheme)
	assert.Equal(t, "8089", url.Port())
	assert.Equal(t, fmt.Sprintf("%s%s%s", "api/", TestTenantID, "/services/search/jobs"), url.Path)
	assert.Empty(t, url.Fragment)
}

func TestNewClient(t *testing.T) {
	client := getClient()
	searchService := &SearchService{client: client}
	assert.Equal(t, TestToken, client.token)
	assert.Equal(t, TestHost, fmt.Sprintf("%s://%s", client.URL.Scheme, client.URL.Host))
	assert.Equal(t, TestTimeOut, client.httpClient.Timeout)
	assert.Equal(t, searchService, client.SearchService)
}

func TestNewRequest(t *testing.T) {
	client := getClient()
	body := []byte(`{"test":"This is a test body"}`)
	expectedAuth := []string{"Bearer " + TestToken}
	requestBody := bytes.NewBuffer(body)
	tests := []struct {
		method string
		url    string
		body   io.Reader
	}{
		{http.MethodGet, TestHost, nil},
		{http.MethodPost, TestHost, requestBody},
		{http.MethodPut, TestHost, requestBody},
		{http.MethodPatch, TestHost, requestBody},
		{http.MethodDelete, TestHost, nil},
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
	client := getClient()
	req, err := client.NewRequest(http.MethodGet, TestHost, nil)
	if err != nil {
		t.Errorf("NewRequest returns unexpected error %v", err)
	}
	expectedAuth := []string{"Bearer " + TestToken}
	if got, want := req.Header["Authorization"], expectedAuth; !reflect.DeepEqual(got, want) {
		t.Errorf("NewRequest authorization is %v, want %v", got, want)
	}
}

func TestNewRequestError(t *testing.T) {
	client := getClient()
	_, err := client.NewRequest("#~/", TestHost, nil)
	if err == nil {
		t.Errorf("NewRequest expected to return error, got %v", err)
	}
}

func TestNewStubbyRequest(t *testing.T) {
	client := getClient()
	resp, _ := client.DoRequest(http.MethodGet, url.URL{Scheme: TestStubbySchme, Host: TestStubbyHost, Path: "/error"}, nil)
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
	var client = getSplunkClient()
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
