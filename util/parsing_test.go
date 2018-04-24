package util

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
)

func TestParseResponseError(t *testing.T) {
	testData := []byte("testbool=true&testfloat32=0.999&testfloat64=0.555&testint=1")
	type TestModel struct {
		TestID string `json:"TestID"`
	}

	httpResp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(testData)),
	}
	var testModel TestModel
	err := ParseResponse(&testModel, httpResp, nil)
	if err == nil {
		t.Errorf("ParseResponse expected to raise an error, got %v", err)
	}
}

func TestParseResponse(t *testing.T) {
	testData := []byte(`{"TestID":"1"}`)
	type TestModel struct {
		TestID string `json:"TestID"`
	}
	httpResp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(testData)),
	}
	var testModel TestModel
	err := ParseResponse(&testModel, httpResp, nil)
	if err != nil {
		t.Errorf("ParseResponse expected to return an error, got %v", err)
	}
}

func TestParseErrorNoError(t *testing.T) {
	testData := []byte(`{"TestID":"1"}`)
	httpResp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(testData)),
	}
	if err := ParseError(httpResp, nil); err != nil {
		t.Errorf("ParseError expected to not return an error, got %v", err)
	}
}

func TestParseErrorReturnError(t *testing.T) {
	testData := []byte(`{"TestID":"1"}`)
	httpResp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(testData)),
	}
	err := errors.New("TestParseErrorReturnError should return this error")
	if err = ParseError(httpResp, err); err == nil {
		t.Errorf("ParseError expected to return an error, got %v", err)
	}
}

func TestParseUrlParams(t *testing.T) {
	params := model.HecEvent{Host: "http://ssc-sdk-shared-stubby:8882", Event: "test", Source: "manual-events", Sourcetype: "sourcetype:eventgen"}
	values := ParseURLParams(params)
	assert.Equal(t, "http://ssc-sdk-shared-stubby:8882", values.Get("host"))
	assert.Equal(t, "manual-events", values.Get("source"))
	assert.Equal(t, "sourcetype:eventgen", values.Get("sourcetype"))
	assert.Empty(t, values["event"])
}

type Enum string
type Inner struct {
	ID   int    `key:"id"`
	Name string `key:"name"`
}

type Outter struct {
	Enum  Enum `key:"enum"`
	Inner Inner
	Other string   `key:"other"`
	Arr   []string `key:"arr"`
}

func TestParseUrlParamsConstRecursiveAndArray(t *testing.T) {
	o := Outter{Enum: "typedef", Inner: Inner{ID: 33, Name: "anonymous"}, Other: "stuff", Arr: []string{"ele1"}}
	values := ParseURLParams(o)
	assert.Equal(t, "typedef", values.Get("enum"))
	assert.Equal(t, "stuff", values.Get("other"))
	assert.Equal(t, "33", values.Get("id"))
	assert.Equal(t, "anonymous", values.Get("name"))
	assert.Equal(t, "ele1", values.Get("arr"))
}
