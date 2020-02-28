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

package util

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp/syntax"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseResponseParsingError(t *testing.T) {
	testData := []byte("testbool=true&testfloat32=0.999&testfloat64=0.555&testint=1")
	type TestModel struct {
		TestID string `json:"TestID"`
	}
	httpResp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(testData)),
	}
	var testModel TestModel
	err := ParseResponse(&testModel, httpResp)
	if err == nil {
		t.Errorf("ParseResponse expected to raise an error, got %v", err)
	}
}

func TestParseResponseSuccess(t *testing.T) {
	testData := []byte(`{"TestID":"1"}`)
	type TestModel struct {
		TestID string `json:"TestID"`
	}
	httpResp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(testData)),
	}
	var testModel TestModel
	err := ParseResponse(&testModel, httpResp)
	if err != nil {
		t.Errorf("ParseResponse expected to not return an error, got %v", err)
	}
}

func TestParseResponseNilResponseError(t *testing.T) {
	type TestModel struct {
		TestID string `json:"TestID"`
	}
	parsingError := errors.New("nil response provided")
	var testModel TestModel
	err := ParseResponse(&testModel, nil)
	assert.NotNil(t, err)
	assert.Equal(t, parsingError, err)
}

func TestParseEmptyResponse(t *testing.T) {
	testData := []byte("")
	type TestModel struct {
		TestID string `json:"TestID"`
	}
	httpResp := &http.Response{
		StatusCode: 204,
		Body:       ioutil.NopCloser(bytes.NewReader(testData)),
	}
	var testModel TestModel
	err := ParseResponse(&testModel, httpResp)
	assert.Nil(t, err)
}

func TestParseUrlParams(t *testing.T) {
	type Event struct {
		Host       string      `key:"host"`
		Event      interface{} `json:"event"`
		Source     string      `key:"source"`
		Sourcetype string      `key:"sourcetype"`
	}
	params := Event{Host: "http://splunk-cloud-sdk-shared-stubby:8882", Event: "test", Source: "manual-events", Sourcetype: "sourcetype:eventgen"}
	values := ParseURLParams(params)
	assert.Equal(t, "http://splunk-cloud-sdk-shared-stubby:8882", values.Get("host"))
	assert.Equal(t, "manual-events", values.Get("source"))
	assert.Equal(t, "sourcetype:eventgen", values.Get("sourcetype"))
	assert.Empty(t, values["event"])
}

type Enum string

const (
	Do Enum = "do"
	Re Enum = "re"
	Mi Enum = "mi"
)

type Struct struct {
	ID   int    `key:"id"`
	Name string `key:"name"`
}

// URLParams with defaults: exploded objects and unexploded arrays
type URLParams struct {
	Arr        []string               `key:"arr"`
	ArrEnum    []Enum                 `key:"arrenums"`
	Struct     Struct                 `key:"strct"`
	Map        map[string]interface{} `key:"map"`
	MapStrings map[string]string
	Enum       Enum     `key:"enum"`
	Str        string   `key:"str"`
	Int        int64    `key:"integer"`
	Float      float32  `key:"flt"`
	OptStr     *string  `key:"ostr"`
	OptInt     *int     `key:"ointeger"`
	OptFloat   *float32 `key:"oflt"`
}

type URLParamsUnexploded struct {
	Arr        []string               `key:"arr" explode:"false"`
	ArrEnum    []Enum                 `key:"arrenums" explode:"false"`
	Struct     Struct                 `key:"strct" explode:"false"`
	Map        map[string]interface{} `key:"map" explode:"false"`
	MapStrings map[string]string      `explode:"false"`
}

type URLParamsExploded struct {
	Arr        []string               `key:"arr" explode:"true"`
	ArrEnum    []Enum                 `key:"arrenums" explode:"true"`
	Struct     Struct                 `key:"strct" explode:"true"`
	Map        map[string]interface{} `key:"map" explode:"true"`
	MapStrings map[string]string      `explode:"true"`
}

func TestParseUrlParamsConstRecursiveAndArray(t *testing.T) {
	oint := 45
	o := URLParams{
		Arr:        []string{"ele1", "ele2", "ele3"},
		ArrEnum:    []Enum{Mi, Do, Do},
		Struct:     Struct{ID: 33, Name: "anonymous"},
		Map:        map[string]interface{}{"foo": "bar", "five": 5},
		MapStrings: map[string]string{"a": "aa", "b": "bb", "c": "cc"},
		Enum:       Re,
		Str:        "stuff",
		Int:        55,
		Float:      66.666,
		OptInt:     &oint,
	}
	values := ParseURLParams(o)
	assert.Equal(t, "ele1,ele2,ele3", values.Get("arr"))
	assert.Equal(t, "mi,do,do", values.Get("arrenums"))
	assert.Equal(t, "33", values.Get("id"))
	assert.Equal(t, "anonymous", values.Get("name"))
	assert.Equal(t, "aa", values.Get("a"))
	assert.Equal(t, "bb", values.Get("b"))
	assert.Equal(t, "cc", values.Get("c"))
	assert.Equal(t, "5", values.Get("five"))
	assert.Equal(t, "bar", values.Get("foo"))
	assert.Equal(t, "re", values.Get("enum"))
	assert.Equal(t, "stuff", values.Get("str"))
	assert.Equal(t, "55", values.Get("integer"))
	assert.Equal(t, "66.666", values.Get("flt"))
	assert.Equal(t, "", values.Get("ostr"))
	assert.Equal(t, "45", values.Get("ointeger"))
	assert.Equal(t, "", values.Get("oflt"))
	// Assert the encoded params replacing %2C with comma for readability
	enc := strings.Replace(values.Encode(), "%2C", ",", -1)
	assert.Equal(t, "a=aa&arr=ele1,ele2,ele3&arrenums=mi,do,do&b=bb&c=cc&enum=re&five=5&flt=66.666&foo=bar&id=33&integer=55&name=anonymous&ointeger=45&str=stuff", enc)
}

func TestParseURLParamsUnexploded(t *testing.T) {
	o := URLParamsUnexploded{
		Arr:        []string{"ele1", "ele2", "ele3"},
		ArrEnum:    []Enum{Mi, Do, Do},
		Struct:     Struct{ID: 33, Name: "anonymous"},
		Map:        map[string]interface{}{"foo": "bar", "five": 5, "list": []string{"x", "y", "z"}},
		MapStrings: map[string]string{"a": "aa", "b": "bb", "c": "cc"},
	}
	values := ParseURLParams(o)
	assert.Equal(t, "ele1,ele2,ele3", values.Get("arr"))
	assert.Equal(t, "mi,do,do", values.Get("arrenums"))
	assert.Equal(t, "id,33,name,anonymous", values.Get("strct"))
	// For maps the key ordering can change so simple assert that kv pairs are contained within the string
	m := values.Get("map")
	assert.Equal(t, len("foo,bar,five,5,list,x,list,y,list,z"), len(m))
	assert.Contains(t, m, "foo,bar")
	assert.Contains(t, m, "five,5")
	assert.Contains(t, m, "list,x")
	assert.Contains(t, m, "list,y")
	assert.Contains(t, m, "list,z")
	ms := values.Get("MapStrings")
	assert.Equal(t, len("a,aa,b,bb,c,cc"), len(ms))
	assert.Contains(t, ms, "a,aa")
	assert.Contains(t, ms, "b,bb")
	assert.Contains(t, ms, "c,cc")

	// Assert the encoded params replacing %2C with comma for readability
	enc := strings.Replace(values.Encode(), "%2C", ",", -1)
	assert.Equal(t, fmt.Sprintf("MapStrings=%s&arr=ele1,ele2,ele3&arrenums=mi,do,do&map=%s&strct=id,33,name,anonymous", ms, m), enc)
}

func TestParseURLParamsExploded(t *testing.T) {
	o := URLParamsExploded{
		Arr:        []string{"ele1", "ele2", "ele3"},
		ArrEnum:    []Enum{Mi, Do, Do},
		Struct:     Struct{ID: 33, Name: "anonymous"},
		Map:        map[string]interface{}{"foo": "bar", "five": 5, "list": []string{"x", "y", "z"}},
		MapStrings: map[string]string{"a": "aa", "b": "bb", "c": "cc"},
	}
	values := ParseURLParams(o)
	assert.Len(t, values["arr"], 3)
	assert.Equal(t, "ele1", values["arr"][0])
	assert.Equal(t, "ele2", values["arr"][1])
	assert.Equal(t, "ele3", values["arr"][2])
	assert.Len(t, values["arrenums"], 3)
	assert.Equal(t, "mi", values["arrenums"][0])
	assert.Equal(t, "do", values["arrenums"][1])
	assert.Equal(t, "do", values["arrenums"][2])
	assert.Equal(t, "33", values.Get("id"))
	assert.Equal(t, "anonymous", values.Get("name"))
	assert.Equal(t, "aa", values.Get("a"))
	assert.Equal(t, "bb", values.Get("b"))
	assert.Equal(t, "cc", values.Get("c"))
	assert.Equal(t, "5", values.Get("five"))
	assert.Equal(t, "bar", values.Get("foo"))
	assert.Equal(t, []string{"x", "y", "z"}, values["list"])
	// Assert the encoded params replacing %2C with comma for readability
	enc := strings.Replace(values.Encode(), "%2C", ",", -1)
	assert.Equal(t, "a=aa&arr=ele1&arr=ele2&arr=ele3&arrenums=mi&arrenums=do&arrenums=do&b=bb&c=cc&five=5&foo=bar&id=33&list=x&list=y&list=z&name=anonymous", enc)
}

func TestParseTemplatedPath(t *testing.T) {
	template := "/{tenant}/myservice/v1beta2/all-things/thing_id_2939535/{thing_name}/status/{status1}"
	tenant := "tenant230_t35"
	name := "my-thing-name"
	status := "FAILED"
	path := fmt.Sprintf("/%s/myservice/v1beta2/all-things/thing_id_2939535/%s/status/%s", tenant, name, status)
	params, err := ParseTemplatedPath(template, path)
	require.Nil(t, err)
	require.NotEmpty(t, params)
	assert.Equal(t, tenant, params["tenant"])
	assert.Equal(t, name, params["thing_name"])
	assert.Equal(t, status, params["status1"])
}

func TestParseTemplatedPathBadCaptureGroup(t *testing.T) {
	// Dashes are not valid for capture group names
	template := "/{tenant}/myservice/v1beta2/all-things/thing_id_2939535/{thing-name}/status/{status-1}"
	tenant := "tenant230_t35"
	name := "my-thing-name"
	status := "FAILED"
	path := fmt.Sprintf("/%s/myservice/v1beta2/all-things/thing_id_2939535/%s/status/%s", tenant, name, status)
	_, err := ParseTemplatedPath(template, path)
	require.Contains(t, err.Error(), syntax.ErrInvalidNamedCapture)
}

func TestParseTemplatedPathMissingMatch(t *testing.T) {
	template := "/{tenant}/myservice/v1beta2/all-things/thing_id_2939535/{thing_name}/status/{status1}"
	tenant := "tenant230_t35"
	name := "my-thing-name"
	// no status
	path := fmt.Sprintf("/%s/myservice/v1beta2/all-things/thing_id_2939535/%s/status", tenant, name)
	params, err := ParseTemplatedPath(template, path)
	require.Nil(t, err)
	// no matches are returned if path doesn't match regex
	require.Empty(t, params)
}
