package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"reflect"
	"testing"

	. "github.com/splunk/ssc-client-go/util"
	"github.com/stretchr/testify/assert"
)

func getClient() *Client {
	return NewClient(TestTenantID, TestToken, TestHost, TestTimeOut)
}

func TestBuildURLNoURLPath(t *testing.T) {
	client := getClient()
	url := client.BuildURL("")

	if got, want := url.Hostname(), "localhost"; got != want {
		t.Errorf("hostname invalid, got %s, want %s", got, want)
	}
	if got, want := url.Scheme, "https"; got != want {
		t.Errorf("scheme invalid, got %s, want %s", got, want)
	}
	if got, want := url.Port(), "8089"; got != want {
		t.Errorf("port invalid, got %s, want %s", got, want)
	}
	if got, want := url.Path, ""; got != want {
		t.Errorf("path invalid, got %s, want %s", got, want)
	}
	if got, want := url.Fragment, ""; got != want {
		t.Errorf("fragment invalid, got %s, want %s", got, want)
	}
	if url.User != nil {
		t.Errorf("user invalid, got %s, want %v", url.User, nil)
	}
}

func TestBuildURLNoHost(t *testing.T) {
	client := getClient()
	url := client.BuildURL("services",
		"search", "jobs")

	if got, want := url.Hostname(), "localhost"; got != want {
		t.Errorf("hostname invalid, got %s, want %s", got, want)
	}
	if got, want := url.Scheme, "https"; got != want {
		t.Errorf("scheme invalid, got %s, want %s", got, want)
	}
	if got, want := url.Port(), "8089"; got != want {
		t.Errorf("port invalid, got %s, want %s", got, want)
	}
	if got, want := url.Path, "services/search/jobs"; got != want {
		t.Errorf("path invalid, got %s, want %s", got, want)
	}
	if got, want := url.Fragment, ""; got != want {
		t.Errorf("fragment invalid, got %s, want %s", got, want)
	}
	if url.User != nil {
		t.Errorf("user invalid, got %s, want %v", url.User, nil)
	}
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
		{MethodGet, TestHost, nil},
		{MethodPost, TestHost, requestBody},
		{MethodPut, TestHost, requestBody},
		{MethodPatch, TestHost, requestBody},
		{MethodDelete, TestHost, nil},
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
		if test.method == MethodGet || test.method == MethodDelete {
			t.Skipf("Skip NewRequest body test for %v and %v method", MethodGet, MethodDelete)
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
	req, err := client.NewRequest(MethodGet, TestHost, nil)
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
	resp, _ := client.DoRequest(MethodGet, url.URL{Scheme: TestStubbySchme, Host: TestStubbyHost, Path: "/error"}, nil)
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
