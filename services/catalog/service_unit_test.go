// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package catalog

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	notMap   = `foo=bar`
	noKind   = `{"foo":"bar"}`
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
	jobDS = `{
		"owner": "me@example.com",
		"created": "2018-11-28 09:25:42.000262",
		"modified": "2018-11-28 09:25:43.000377",
		"version": 1,
		"id": "99999aaaabbbbcccc",
		"module": "",
		"name": "sid_1234567890_123",
		"kind": "job",
		"createdby": "me@example.com",
		"modifiedby": "me@example.com",
		"internalname": "sid_1234567890_123",
		"resourcename": "sid_1234567890_123",
		"maxTime": 3600,
		"timeOfSearch": "1543440342",
		"deleteTime": "2018-11-29T21:25:42.258412+00:00",
		"query": "| from mylookup",
		"timeFormat": "%FT%T.%Q%:z",
		"extractAllFields": false,
		"percentComplete": 100,
		"parameters": {
		  "earliest": "-24h@h",
		  "latest": "now"
		},
		"spl": "| inputlookup mylookup",
		"resultsAvailable": 2,
		"sid": "1543440342.234",
		"status": "done"
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
	randoDS = `{
		"kind": "narwhal",
		"createdby": "me@example.com",
		"modifiedby": "me@example.com",
		"tusks": 1
	  }`
)

// Test error cases:
func TestParseRawDatasetNil(t *testing.T) {
	_, err := ParseRawDataset(nil)
	require.NotNil(t, err)
	assert.Contains(t, err.Error(), "catalog: response was not of type map[string]interface{}")
}

func TestParseRawDatasetNotMap(t *testing.T) {
	_, err := ParseRawDataset(notMap)
	require.NotNil(t, err)
	assert.Contains(t, err.Error(), "catalog: response was not of type map[string]interface{}")
}

func TestParseRawDatasetMarshalFail(t *testing.T) {
	var noKindMap interface{}
	err := json.Unmarshal([]byte(noKind), &noKindMap)
	require.Nil(t, err)
	_, err = ParseRawDataset(noKindMap)
	require.NotNil(t, err)
	assert.Contains(t, err.Error(), "catalog: dataset response did not contain key 'kind' with string value in contents")
}

// Test unknown dataset kind:
func TestParseRawDatasetOther(t *testing.T) {
	var dsMap interface{}
	err := json.Unmarshal([]byte(randoDS), &dsMap)
	require.Nil(t, err)
	ds, err := ParseRawDataset(dsMap)
	require.Nil(t, err)
	assert.Equal(t, "narwhal", ds.GetKind())
	rando, ok := ds.(DatasetBase)
	require.True(t, ok)
	assert.Equal(t, "", rando.Name) // this is not in the payload, so set to zero value
	assert.Equal(t, "me@example.com", rando.CreatedBy)
}

// Test known dataset kinds:

func TestParseRawDatasetLookup(t *testing.T) {
	var dsMap interface{}
	err := json.Unmarshal([]byte(lookupDS), &dsMap)
	require.Nil(t, err)
	ds, err := ParseRawDataset(dsMap)
	require.Nil(t, err)
	assert.Equal(t, string(Lookup), ds.GetKind())
	lookup, ok := ds.(LookupDataset)
	require.True(t, ok)
	assert.Equal(t, "mylookup", lookup.Name)
	assert.Equal(t, "kind==\"lookup\"", lookup.Filter)
}
func TestParseRawDatasetKvCollection(t *testing.T) {
	var dsMap interface{}
	err := json.Unmarshal([]byte(kvDS), &dsMap)
	require.Nil(t, err)
	ds, err := ParseRawDataset(dsMap)
	require.Nil(t, err)
	assert.Equal(t, string(KvCollection), ds.GetKind())
	kvCollection, ok := ds.(KVCollectionDataset)
	require.True(t, ok)
	assert.Equal(t, "mycollection", kvCollection.Name)
}
func TestParseRawDatasetIndex(t *testing.T) {
	var dsMap interface{}
	err := json.Unmarshal([]byte(indexDS), &dsMap)
	require.Nil(t, err)
	ds, err := ParseRawDataset(dsMap)
	require.Nil(t, err)
	assert.Equal(t, string(Index), ds.GetKind())
	index, ok := ds.(IndexDataset)
	require.True(t, ok)
	assert.Equal(t, "myindex", index.Name)
	assert.Equal(t, 999, *index.FrozenTimePeriodInSecs)
}

func TestParseRawDatasetMetric(t *testing.T) {
	var dsMap interface{}
	err := json.Unmarshal([]byte(metricDS), &dsMap)
	require.Nil(t, err)
	ds, err := ParseRawDataset(dsMap)
	require.Nil(t, err)
	assert.Equal(t, string(Metric), ds.GetKind())
	metric, ok := ds.(MetricDataset)
	require.True(t, ok)
	assert.Equal(t, "mymetric", metric.Name)
	assert.Equal(t, 7, *metric.FrozenTimePeriodInSecs)
}

func TestParseRawDatasetImport(t *testing.T) {
	var dsMap interface{}
	err := json.Unmarshal([]byte(importDS), &dsMap)
	require.Nil(t, err)
	ds, err := ParseRawDataset(dsMap)
	require.Nil(t, err)
	assert.Equal(t, string(Import), ds.GetKind())
	imp, ok := ds.(ImportDataset)
	require.True(t, ok)
	assert.Equal(t, "myimport", imp.Name)
	assert.Equal(t, "1234aaabbbccc", *imp.OriginalDatasetID)
}

func TestParseRawDatasetJob(t *testing.T) {
	var dsMap interface{}
	err := json.Unmarshal([]byte(jobDS), &dsMap)
	require.Nil(t, err)
	ds, err := ParseRawDataset(dsMap)
	require.Nil(t, err)
	assert.Equal(t, string(Job), ds.GetKind())
	job, ok := ds.(JobDataset)
	require.True(t, ok)
	assert.Equal(t, "sid_1234567890_123", job.Name)
	assert.Equal(t, "done", *job.Status)
}

func TestParseRawDatasetView(t *testing.T) {
	var dsMap interface{}
	err := json.Unmarshal([]byte(viewDS), &dsMap)
	require.Nil(t, err)
	ds, err := ParseRawDataset(dsMap)
	require.Nil(t, err)
	assert.Equal(t, string(View), ds.GetKind())
	view, ok := ds.(ViewDataset)
	require.True(t, ok)
	assert.Equal(t, "myview", view.Name)
	assert.Equal(t, "search index=main | head limit=10 | stats count()", *view.Search)
}
