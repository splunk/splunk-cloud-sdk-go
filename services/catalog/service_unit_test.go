// Copyright © 2019 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package catalog

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	lookupDS = `{
		"owner": "me@example.com",
		"created": "2018-11-29 08:25:19.000987",
		"modified": "2018-11-29 08:25:19.000987",
		"version": 1,
		"id": "99999aaaabbbbcccc",
		"module": "mymodule",
		"name": "mylookup",
		"kind": "lookup",
		"createdby": "me@example.com",
		"modifiedby": "me@example.com",
		"internalname": "mymodule.mylookup",
		"resourcename": "mymodule.mylookup",
		"filter": "kind==\"lookup\"",
		"caseSensitiveMatch": true,
		"externalName": "test_externalName",
		"externalKind": "kvcollection"
	  }`
	kvDS = `{
		"owner": "me@example.com",
		"created": "2018-11-28 11:54:21.000634",
		"modified": "2018-11-28 11:54:21.000634",
		"version": 1,
		"id": "99999aaaabbbbcccc",
		"module": "mymodule",
		"name": "mycollection",
		"kind": "kvcollection",
		"createdby": "me@example.com",
		"modifiedby": "me@example.com",
		"internalname": "mymodule.mycollection",
		"resourcename": "mymodule.mycollection"
	  }`
	indexDS = `{
		"owner": "me@example.com",
		"created": "2018-11-28 11:54:13.000889",
		"modified": "2018-11-28 11:54:14.000000",
		"version": 1,
		"id": "99999aaaabbbbcccc",
		"module": "mymodule",
		"name": "myindex",
		"kind": "index",
		"createdby": "me@example.com",
		"modifiedby": "me@example.com",
		"internalname": "mymodule_____myindex",
		"resourcename": "mymodule.myindex",
		"disabled": true,
		"frozenTimePeriodInSecs": 999
	  }`
	metricDS = `{
		"owner": "test1@splunk.com",
		"created": "2018-11-28 11:39:16.000228",
		"modified": "2018-11-28 11:39:16.000286",
		"version": 1,
		"id": "99999aaaabbbbcccc",
		"module": "mymodule",
		"name": "mymetric",
		"kind": "metric",
		"createdby": "me@example.com",
		"modifiedby": "me@example.com",
		"internalname": "mymodule_____mymetric",
		"resourcename": "mymodule.mymetric",
		"disabled": true,
		"frozenTimePeriodInSecs": 7
	  }`
	importDS = `{
		"owner": "me@example.com",
		"created": "2018-11-29 11:27:40.000519",
		"modified": "2018-11-29 11:27:40.000519",
		"version": 1,
		"id": "99999aaaabbbbcccc",
		"module": "mymodule",
		"name": "myimport",
		"kind": "import",
		"createdby": "me@example.com",
		"modifiedby": "me@example.com",
		"resourcename": "mymodule.myimport",
		"originalDatasetId": "1234aaabbbccc",
		"sourceName": "myimport",
		"sourceModule": ""
	  }`
	viewDS = `{
		"owner": "me@example.com",
		"created": "2018-11-30 08:13:36.000727",
		"modified": "2018-11-30 08:13:36.000727",
		"version": 1,
		"id": "99999aaaabbbbcccc",
		"module": "",
		"name": "myview",
		"kind": "view",
		"createdby": "me@example.com",
		"modifiedby": "me@example.com",
		"internalname": "myview",
		"resourcename": "myview",
		"search": "search index=main | head limit=10 | stats count()"
	  }`
)

//Actual error message for this test is not readable, will follow up to see what we can do in Codegen UnmarshalJson implementation in models
//func TestParseResponseNotMap(t *testing.T) {
//	var nMap *Dataset
//	httpResp := &http.Response{
//		Body: ioutil.NopCloser(bytes.NewReader([]byte(notMap))),
//	}
//	err := util.ParseResponse(&nMap, httpResp)
//	require.NotNil(t, err)
//	assert.Contains(t, err.Error(), "json: cannot unmarshal string into Go value of type catalog.discriminator")
//}

// Test known dataset kinds:
func TestParseResponseLookup(t *testing.T) {
	var ds *Dataset
	httpResp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(lookupDS))),
	}
	err := util.ParseResponse(&ds, httpResp)
	require.NoError(t, err)
	assert.Equal(t, LookupDatasetKindLookup, ds.LookupDataset().Kind)

	require.NotNil(t, ds.LookupDataset().Id)
	assert.Equal(t, "mylookup", ds.LookupDataset().Name)
	assert.Equal(t, "kind==\"lookup\"", *ds.LookupDataset().Filter)
}
func TestParseResponseKvCollection(t *testing.T) {
	var kv *Dataset
	httpResp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(kvDS))),
	}
	err := util.ParseResponse(&kv, httpResp)
	require.NoError(t, err)
	assert.Equal(t, KvCollectionDatasetKindKvcollection, kv.KvCollectionDataset().Kind)

	require.NotNil(t, kv.KvCollectionDataset().Id)
	assert.Equal(t, "mycollection", kv.KvCollectionDataset().Name)
}
func TestParseResponseIndex(t *testing.T) {
	var indx *Dataset
	httpResp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(indexDS))),
	}
	err := util.ParseResponse(&indx, httpResp)
	require.NoError(t, err)
	assert.Equal(t, IndexDatasetKindIndex, indx.IndexDataset().Kind)

	require.NotNil(t, indx.IndexDataset().Id)
	assert.Equal(t, "myindex", indx.IndexDataset().Name)
	assert.Equal(t, int32(999), *indx.IndexDataset().FrozenTimePeriodInSecs)
}

func TestParseResponseMetric(t *testing.T) {
	var metric *Dataset
	httpResp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(metricDS))),
	}
	err := util.ParseResponse(&metric, httpResp)
	require.NoError(t, err)
	assert.Equal(t, MetricDatasetKindMetric, metric.MetricDataset().Kind)

	require.NotNil(t, metric.MetricDataset().Id)
	assert.Equal(t, "mymetric", metric.MetricDataset().Name)
	assert.Equal(t, int32(7), *metric.MetricDataset().FrozenTimePeriodInSecs)
}

func TestParseResponseImport(t *testing.T) {
	var imp *Dataset
	httpResp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(importDS))),
	}
	err := util.ParseResponse(&imp, httpResp)
	require.NoError(t, err)
	assert.Equal(t, ImportDatasetKindModelImport, imp.ImportDataset().Kind)

	require.NotNil(t, imp.ImportDataset().Id)
	assert.Equal(t, "myimport", imp.ImportDataset().Name)
	assert.Equal(t, "myimport", imp.ImportDataset().SourceName)
}

func TestParseResponseView(t *testing.T) {
	var view *Dataset
	httpResp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(viewDS))),
	}
	err := util.ParseResponse(&view, httpResp)
	require.NoError(t, err)
	assert.Equal(t, ViewDatasetKindView, view.ViewDataset().Kind)

	require.NotNil(t, view.ViewDataset().Id)
	assert.Equal(t, "myview", view.ViewDataset().Name)
	assert.Equal(t, "search index=main | head limit=10 | stats count()", view.ViewDataset().Search)
}
