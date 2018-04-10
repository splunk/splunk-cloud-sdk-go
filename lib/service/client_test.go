package service

import (
	"bytes"
	"io"
	"io/ioutil"
	"math"
	"reflect"
	"testing"
	"time"
)

const (
	testUser       = "admin"
	testPassword   = "changeme"
	testSessionKey = "123"
	testHost       = "localhost:8089"
	testStubbyHost = "ssc-sdk-shared-stubby:8882"
	testScheme     = "https"
	testURL        = "https://test:8089/test"
	testTimeOut = time.Second*10
)

func getClient() *Client {
	return NewClient("", [2]string{testUser, testPassword}, testHost, testScheme, testTimeOut, true)
}


func TestBuildSplunkdURLNoURLPath(t *testing.T) {
	client := getClient()
	url := client.BuildSplunkdURL(nil, "")

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

func TestBuildSplunkdURLNoHost(t *testing.T) {
	client := getClient()
	url := client.BuildSplunkdURL(nil, "services",
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
	if got, want := client.SessionKey, ""; got != want {
		t.Errorf("NewClient SessionKey is %v, want %v", got, want)
	}
	if got, want := client.Auth, defaultAuth; got != want {
		t.Errorf("NewClient Auth is %v, want %v", got, want)
	}
	if got, want := client.Host, testHost; got != want {
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

func TestNewRequestSessionKey(t *testing.T) {
	client := getClient()
	client.SessionKey = testSessionKey
	req, err := client.NewRequest(MethodGet, testURL, nil)
	if err != nil {
		t.Errorf("NewRequest returns unexpected error %v", err)
	}
	expectedBasicAuth := []string{"Splunk " + client.SessionKey}
	if got, want := req.Header["Authorization"], expectedBasicAuth; !reflect.DeepEqual(got, want) {
		t.Errorf("NewRequest authorization is %v, want %v", got, want)
	}
}

func TestNewRequestError(t *testing.T) {
	client := getClient()
	client.SessionKey = testSessionKey
	_, err := client.NewRequest("#~/", testURL, nil)
	if err == nil {
		t.Errorf("NewRequest expected to return error, got %v", err)
	}
}

func TestEncodeRequestBodyNil(t *testing.T) {
	client := getClient()
	response, err := client.EncodeRequestBody(nil)
	if len(response) > 0 {
		t.Errorf("EncodeRequestBody expected to return nil, got %v", response)
	}
	if err != nil {
		t.Errorf("EncodeRequestBody expected to not return error, got %v", err)
	}
}

func TestEncodeRequestBodyString(t *testing.T) {
	client := getClient()
	got, err := client.EncodeRequestBody(`{"test":"This is a test body"}`)
	// expect := []byte(`{"test":"This is a test body"}`)
	if value := reflect.ValueOf(got); value.Kind() != reflect.Slice {
		t.Errorf("EncodeRequestBody expected to return []byte, got %v", got)
	}
	if err != nil {
		t.Errorf("EncodeRequestBody expected to not return error, got %v", err)
	}
}

func TestTestEncodeRequestBodyMap(t *testing.T) {
	client := getClient()
	testData := map[string]string{
		"testKey": "testValue",
	}
	got, err := client.EncodeRequestBody(testData)
	if value := reflect.ValueOf(got); value.Kind() != reflect.Slice {
		t.Errorf("EncodeRequestBody expected to return []byte, got %v", got)
	}
	if err != nil {
		t.Errorf("EncodeRequestBody expected to not return error, got %v", err)
	}
}

func TestTestEncodeRequestBodyStruct(t *testing.T) {
	client := getClient()
	type TestModel struct {
		testID    string
		testValue string
	}
	testData := TestModel{
		testID:    "123",
		testValue: "test",
	}
	got, err := client.EncodeRequestBody(testData)
	if value := reflect.ValueOf(got); value.Kind() != reflect.Slice {
		t.Errorf("EncodeRequestBody expected to return []byte, got %v", got)
	}
	if err != nil {
		t.Errorf("EncodeRequestBody expected to not return error, got %v", err)
	}
}

func TestTestEncodeRequestBodyInvalid(t *testing.T) {
	client := getClient()
	_, err := client.EncodeRequestBody(123)
	if err == nil {
		t.Errorf("EncodeRequestBody expected to raise an error, got %v", err)
	}
}

func TestEncodeObjectError(t *testing.T) {
	client := getClient()
	_, err := client.EncodeObject(math.Inf(1))
	if err == nil {
		t.Errorf("EncodeObject expected to raise an error, got %v", err)
	}
}

func TestEncodeObjectTypeConversion(t *testing.T) {
	client := getClient()
	intVal := 1
	var float32Val float32 = 0.999
	testData := map[string]interface{}{
		"testBool":    true,
		"testInt":     intVal,
		"testFloat32": float32Val,
		"testFloat64": 0.555,
	}
	want := "testbool=true&testfloat32=0.999&testfloat64=0.555&testint=1"
	got, err := client.EncodeObject(testData)
	gotString := string(got[:])
	if gotString != want {
		t.Errorf("EncodeObject expected to return %v, got %v", want, gotString)
	}
	if err != nil {
		t.Errorf("EncodeObject expected to not return error, got %v", err)
	}
}

//func TestNewStubbyRequest(t *testing.T) {
//	client := getClient()
//	resp, _ := client.DoRequest(MethodGet, url.URL{Scheme: "http", Host: testStubbyHost, Path: "/error"}, nil, URLEncoded)
//	if resp.StatusCode != 500 {
//		t.Fatalf("client.DoRequest to /error endpoint expected Response Code: %d, Received: %d", 500, resp.StatusCode)
//	}
//	defer resp.Body.Close()
//	b := new(bytes.Buffer)
//	b.ReadFrom(resp.Body)
//	content := new(map[string]string)
//	if err := json.NewDecoder(b).Decode(content); err != nil {
//		t.Fatalf("client.DoRequest error unmarshalling response, err: %v", err)
//	}
//	if (*content)["message"] != "Something exploded" {
//		t.Fatalf("client.DoRequest error/ expecting response {\"message\":\"Something exploded\"} Received: %+v", content)
//	}
//}
