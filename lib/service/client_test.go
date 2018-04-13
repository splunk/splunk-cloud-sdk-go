package service

import (
	"bytes"
	"io"
	"io/ioutil"
	"reflect"
	"testing"
	"time"
	"net/url"
	"encoding/json"
)

const (
	testUser                     = "admin"
	testPassword                 = "changeme"
	testHost                     = "localhost:8089"
	testStubbyHost               = "ssc-sdk-shared-stubby:8882"
	testScheme                   = "https"
	baseURL                      = "https://localhost:8089"
	testURL                      = "https://test:8089/test"
	testTimeOut    time.Duration = time.Second*10
)

func getClient() *Client {
	return NewClient([2]string{testUser, testPassword}, baseURL, testTimeOut, true)
}


func TestBuildURLNoURLPath(t *testing.T) {
	client := getClient()
	url := client.BuildURL("")

	if got, want := url.Hostname(), "localhost"; got != want {
		t.Errorf("hostname invalid, got %s, want %s", got, want)
	}
	if got, want := url.Scheme, testScheme; got != want {
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
	if got, want := url.Scheme, testScheme; got != want {
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
	var defaultAuth = [2]string{"admin", "changeme"}
	client := getClient()
	searchService := &SearchService{client: client}
	if got, want := client.Auth, defaultAuth; got != want {
		t.Errorf("NewClient Auth is %v, want %v", got, want)
	}
	var u *url.URL
	u,_ = url.Parse(client.URL)

	if got, want := u.Host, testHost; got != want {
		t.Errorf("NewClient Host is %v, want %v", got, want)
	}
	if got, want := client.httpClient.Timeout, testTimeOut; got != want {
		t.Errorf("NewClient httpClient is %v, want %v", got, want)
	}
	if got, want := client.SearchService, searchService; *got != *want {
		t.Errorf("NewClient SearchService is %v, want %v", got, want)
	}
}

func TestNewRequest(t *testing.T) {
	client := getClient()
	body := []byte(`{"test":"This is a test body"}`)
	expectedBasicAuth := []string{"Basic YWRtaW46Y2hhbmdlbWU="}
	requestBody := bytes.NewBuffer(body)
	tests := []struct {
		method string
		url    string
		body   io.Reader
	}{
		{MethodGet, testURL, nil},
		{MethodPost, testURL, requestBody},
		{MethodPut, testURL, requestBody},
		{MethodPatch, testURL, requestBody},
		{MethodDelete, testURL, nil},
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
		if got, want := req.Header["Authorization"], expectedBasicAuth; !reflect.DeepEqual(got, want) {
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

func TestNewRequesthBasicAuthHeader(t *testing.T) {
	client := getClient()
	req, err := client.NewRequest(MethodGet, testURL, nil)
	if err != nil {
		t.Errorf("NewRequest returns unexpected error %v", err)
	}
	expectedBasicAuth := []string{"Basic YWRtaW46Y2hhbmdlbWU="}
	if got, want := req.Header["Authorization"], expectedBasicAuth; !reflect.DeepEqual(got, want) {
		t.Errorf("NewRequest authorization is %v, want %v", got, want)
	}
}

func TestNewRequestError(t *testing.T) {
	client := getClient()
	_, err := client.NewRequest("#~/", testURL, nil)
	if err == nil {
		t.Errorf("NewRequest expected to return error, got %v", err)
	}
}

func TestNewStubbyRequest(t *testing.T) {
	client := getClient()
	resp, _ := client.DoRequest(MethodGet, url.URL{Scheme: "http", Host: testStubbyHost, Path: "/error"}, nil)
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
