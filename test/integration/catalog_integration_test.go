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

package integration

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services/catalog"
	"github.com/splunk/splunk-cloud-sdk-go/services/search"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test Dataset variables
var (
	// Base:
	dsNameTemplate = fmt.Sprintf("gointegds%s_%d", "%s", testutils.RunSuffix)
	newname        = fmt.Sprintf("newmx%d", testutils.RunSuffix)
	newmod         = fmt.Sprintf("newmod%d", testutils.RunSuffix)
	newowner       = "test1@splunk.com"
	// Lookup:
	caseMatch    = true
	externalName = "test_externalName"
	filter       = `kind=="lookup"`
	// Metric/Index:
	disabled               = false
	frozenTimePeriodInSecs = int32(60)
	newftime               = int32(999)
	// View:
	searchString = "search index=main|stats count()"
)

// Test Rule variables
var (
	ruleNameTemplate = fmt.Sprintf("gointegrule%s_%d", "%s", testutils.RunSuffix)
	ruleModule       = "gointeg"
	ruleMatch        = "sourcetype::integration_test_match"
)

func makeDSName(ctx string) string {
	return fmt.Sprintf(dsNameTemplate, ctx)
}

func makeRuleName(ctx string) string {
	return fmt.Sprintf(ruleNameTemplate, ctx)
}

func cleanupDataset(t *testing.T, id string) {
	client := getSdkClient(t)
	err := client.CatalogService.DeleteDatasetById(id)
	assert.Emptyf(t, err, "Error deleting dataset: %s", err)
}

func cleanupRule(t *testing.T, id string) {
	client := getSdkClient(t)
	err := client.CatalogService.DeleteRule(id)
	assert.Emptyf(t, err, "Error deleting rule: %s", err)
}

func cleanupRuleAction(t *testing.T, ruleID, actionID string) {
	client := getSdkClient(t)

	action, _ := client.CatalogService.GetActionByIdForRuleById(ruleID, actionID)
	if action != nil {
		err := client.CatalogService.DeleteActionByIdForRuleById(ruleID, actionID)
		assert.Emptyf(t, err, "Error deleting rule action: %s", err)
	}
}

// createLookupDataset - Helper function for creating a valid Lookup in Catalog
func createLookupDataset(t *testing.T, name string) (*catalog.Dataset, error) {
	externalKind := catalog.LookupDatasetExternalKindKvcollection
	createLookup := catalog.LookupDatasetPost{
		Name:               name,
		Kind:               catalog.LookupDatasetKindLookup,
		Module:             &testutils.TestModule,
		CaseSensitiveMatch: &caseMatch,
		ExternalKind:       externalKind,
		ExternalName:       externalName,
		Filter:             &filter,
	}
	return getSdkClient(t).CatalogService.CreateDataset(catalog.MakeDatasetPostFromLookupDatasetPost(createLookup))
}

// createKVCollectionDataset - Helper function for creating a valid KVCollection in Catalog
func createKVCollectionDataset(t *testing.T, name string) (*catalog.Dataset, error) {
	createKVCollection := catalog.KvCollectionDatasetPost{
		Name:   name,
		Kind:   catalog.KvCollectionDatasetKindKvcollection,
		Module: &testutils.TestModule,
	}
	return getSdkClient(t).CatalogService.CreateDataset(catalog.MakeDatasetPostFromKvCollectionDatasetPost(createKVCollection))
}

// createMetricDataset - Helper function for creating a valid Metric in Catalog
func createMetricDataset(t *testing.T, name string) (*catalog.Dataset, error) {
	createMetric := catalog.MetricDatasetPost{
		Name:                   name,
		Kind:                   catalog.MetricDatasetKindMetric,
		Module:                 &testutils.TestModule,
		Disabled:               disabled,
		FrozenTimePeriodInSecs: &frozenTimePeriodInSecs,
	}
	return getSdkClient(t).CatalogService.CreateDataset(catalog.MakeDatasetPostFromMetricDatasetPost(createMetric))
}

// createIndexDataset - Helper function for creating a valid Index in Catalog
func createIndexDataset(t *testing.T, name string) (*catalog.Dataset, error) {
	return createIndexDatasetWithModule(t, name, testutils.TestModule)
}

// createIndexDataset - Helper function for creating a valid Index in Catalog
func createIndexDatasetWithModule(t *testing.T, name string, module string) (*catalog.Dataset, error) {
	catalogIndex := catalog.IndexDatasetKindIndex
	createIndex := catalog.IndexDatasetPost{
		Name:                   name,
		Kind:                   catalogIndex,
		Module:                 &module,
		Disabled:               disabled,
		FrozenTimePeriodInSecs: &frozenTimePeriodInSecs,
	}
	return getSdkClient(t).CatalogService.CreateDataset(catalog.MakeDatasetPostFromIndexDatasetPost(createIndex))
}

// createImportDatasetByID - Helper function for creating a valid Import in Catalog
//SourceID is no longer supported

func createImportDatasetByID(t *testing.T, name, importID string) (*catalog.Dataset, error) {
	createImport := catalog.ImportDatasetByIdPost{
		Name:   name,
		Kind:   catalog.ImportDatasetKindModelImport,
		Module: &testutils.TestModule,
		Id:     &importID,
	}

	req := catalog.MakeImportDatasetPostFromImportDatasetByIdPost(createImport)
	return getSdkClient(t).CatalogService.CreateDataset(catalog.MakeDatasetPostFromImportDatasetPost(req))
}

// createImportDatasetByName - Helper function for creating a valid Import in Catalog
func createImportDatasetByName(t *testing.T, name, importName, importModule string) (*catalog.Dataset, error) {
	createImport := catalog.ImportDatasetByNamePost{
		Name:         name,
		Kind:         catalog.ImportDatasetKindModelImport,
		Module:       &testutils.TestModule,
		SourceName:   importName,
		SourceModule: importModule,
	}

	req := catalog.MakeImportDatasetPostFromImportDatasetByNamePost(createImport)

	return getSdkClient(t).CatalogService.CreateDataset(catalog.MakeDatasetPostFromImportDatasetPost(req))
}

// createViewDataset - Helper function for creating a valid View in Catalog
func createViewDataset(t *testing.T, name string) (*catalog.Dataset, error) {
	createView := catalog.ViewDatasetPost{
		Name:   name,
		Kind:   catalog.ViewDatasetKindView,
		Module: &testutils.TestModule,
		Search: searchString,
	}
	return getSdkClient(t).CatalogService.CreateDataset(catalog.MakeDatasetPostFromViewDatasetPost(createView))
}

// createViewDataset - Helper function for creating Fields
func createDatasetField(datasetID string, client *sdk.Client, t *testing.T) *catalog.Field {
	return createDatasetFieldName(datasetID, "integ_test_field", client, t)
}

// createViewDataset - Helper function for creating Fields
func createDatasetFieldName(datasetID string, fieldName string, client *sdk.Client, t *testing.T) *catalog.Field {
	dataType := catalog.FieldDataTypeString
	fieldType := catalog.FieldTypeDimension
	prevalenceType := catalog.FieldPrevalenceAll

	testField := catalog.FieldPost{Name: fieldName, Datatype: &dataType, Fieldtype: &fieldType, Prevalence: &prevalenceType}
	resultField, err := client.CatalogService.CreateFieldForDataset(datasetID, testField)
	require.NoError(t, err)
	require.NotEmpty(t, resultField)
	return resultField
}

// assertDatasetKind - Helper to assert that the kind for the Dataset matches model associated with that kind
func assertDatasetKind(t *testing.T, dataset catalog.DatasetGet) {
	if dataset.IsIndexDataset() {
		assert.NotEmpty(t, dataset.IndexDataset().Id)
	} else if dataset.IsViewDataset() {
		assert.NotEmpty(t, dataset.ViewDataset().Id)
	} else if dataset.IsLookupDataset() {
		assert.NotEmpty(t, dataset.LookupDataset().Id)
	} else if dataset.IsIndexDataset() {
		assert.NotEmpty(t, dataset.ImportDataset().Id)
	} else if dataset.IsJobDatasetGet() {
		assert.NotEmpty(t, dataset.JobDatasetGet().Id)
	} else if dataset.IsMetricDataset() {
		assert.NotEmpty(t, dataset.MetricDataset().Id)
	} else if dataset.IsImportDataset() {
		assert.NotEmpty(t, dataset.ImportDataset().Id)
	} else if dataset.IsKvCollectionDataset() {
		assert.NotEmpty(t, dataset.KvCollectionDataset().Id)
	} else if dataset.IsRawInterface() { //handle unknown kinds
		rawDataset := dataset.RawInterface()
		m, _ := rawDataset.(map[string]interface{})
		require.NotNil(t, m["kind"])

	} else {
		// If catalog dataset does not either a known or an unknown kind
		fmt.Println(dataset)
		assert.Fail(t, "Invalid catalog dataset", "Invalid catalog dataset")
	}
}

// Test CreateIndexDataset
func TestCreateIndexDataset(t *testing.T) {
	indexds, err := createIndexDataset(t, makeDSName("crix"))
	require.NoError(t, err)
	defer cleanupDataset(t, indexds.IndexDataset().Id)
	require.NotNil(t, indexds)
	require.Equal(t, catalog.IndexDatasetKindIndex, indexds.IndexDataset().Kind)
}

//
// Test CreateImportDataset
func TestCreateImportDataset(t *testing.T) {
	indexds, err := createIndexDataset(t, makeDSName("crix2"))
	require.NoError(t, err)
	defer cleanupDataset(t, indexds.IndexDataset().Id)
	importds, err := createImportDatasetByName(t, makeDSName("crim"), indexds.IndexDataset().Name, testutils.TestModule)
	require.NoError(t, err)
	defer cleanupDataset(t, importds.ImportDataset().Id)
	require.NotNil(t, importds)
	require.Equal(t, catalog.ImportDatasetKindModelImport, importds.ImportDataset().Kind)
}

// Test CreateDatasetImport
func TestCreateDatasetImport(t *testing.T) {
	client := getSdkClient(t)
	ds1, err := createIndexDataset(t, makeDSName("crix1"))
	require.NoError(t, err)
	require.NotNil(t, ds1.IndexDataset())
	indexds1 := ds1.IndexDataset()
	defer cleanupDataset(t, indexds1.Id)
	ds2, err := createIndexDatasetWithModule(t, makeDSName("crix2"), testutils.TestModule2)
	require.NoError(t, err)
	require.NotNil(t, ds2.IndexDataset())
	indexds2 := ds2.IndexDataset()
	defer cleanupDataset(t, indexds2.Id)

	impby := catalog.DatasetImportedBy{
		Name:   &indexds1.Name,
		Module: indexds2.Module,
	}
	ids, err := client.CatalogService.CreateDatasetImport(indexds1.Id, impby)
	require.NoError(t, err)
	assert.NotNil(t, ids)
}

//
// Test CreateKVCollectionDataset
func TestKVCollectionDataset(t *testing.T) {
	kvds, err := createKVCollectionDataset(t, makeDSName("crkv"))
	require.NoError(t, err)
	defer cleanupDataset(t, kvds.KvCollectionDataset().Id)
	require.NotNil(t, kvds)
	require.Equal(t, catalog.KvCollectionDatasetKindKvcollection, kvds.KvCollectionDataset().Kind)
}

// Test CreateLookupDataset
func TestLookupDataset(t *testing.T) {
	lookupds, err := createLookupDataset(t, makeDSName("crlk"))
	require.NoError(t, err)
	defer cleanupDataset(t, lookupds.LookupDataset().Id)
	require.NotNil(t, lookupds)
	require.Equal(t, catalog.LookupDatasetKindLookup, lookupds.LookupDataset().Kind)
}

// Test CreateMetricDataset
func TestMetricDataset(t *testing.T) {
	metricds, err := createMetricDataset(t, makeDSName("crmx"))
	require.NoError(t, err)
	defer cleanupDataset(t, metricds.MetricDataset().Id)
	require.NotNil(t, metricds)
	require.Equal(t, catalog.MetricDatasetKindMetric, metricds.MetricDataset().Kind)
}

// Test CreateViewDataset
func TestViewDataset(t *testing.T) {
	viewds, err := createViewDataset(t, makeDSName("crvw"))
	require.NoError(t, err)
	defer cleanupDataset(t, viewds.ViewDataset().Id)
	require.NotNil(t, viewds)
	require.Equal(t, catalog.ViewDatasetKindView, viewds.ViewDataset().Kind)
}

// Test CreateDataset for 409 DatasetInfo already present error
func TestCreateDatasetDataAlreadyPresentError(t *testing.T) {
	// create dataset
	ds, err := createLookupDataset(t, makeDSName("409"))
	require.NoError(t, err)
	defer cleanupDataset(t, ds.LookupDataset().Id)
	_, err = createLookupDataset(t, makeDSName("409"))
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	require.True(t, ok)
	assert.Equal(t, 409, httpErr.HTTPStatusCode)
}

// Test CreateDataset for 401 Unauthorized operation error
func TestCreateDatasetUnauthorizedOperationError(t *testing.T) {
	name := makeDSName("401")
	createView := catalog.ViewDatasetPost{
		Name:   name,
		Kind:   catalog.ViewDatasetKindView,
		Module: &testutils.TestModule,
		Search: searchString,
	}

	ds, err := getInvalidClient(t).CatalogService.CreateDataset(catalog.MakeDatasetPostFromViewDatasetPost(createView))
	if ds != nil {
		defer cleanupDataset(t, ds.ViewDataset().Id)
	}
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 401, httpErr.HTTPStatusCode)
}

// Test ListDatasets
func TestListAllDatasets(t *testing.T) {
	ds, err := createLookupDataset(t, makeDSName("getall"))
	require.NoError(t, err)
	defer cleanupDataset(t, ds.LookupDataset().Id)
	req := &catalog.ListDatasetsQueryParams{Filter: "kind!=\"reserved\""}
	datasets, err := getSdkClient(t).CatalogService.ListDatasets(req)
	require.NoError(t, err)
	assert.NotZero(t, len(datasets))

	// We should be able assert that the kinds are known and relate to their associated model
	for i := 0; i < len(datasets); i++ {
		assertDatasetKind(t, datasets[i])
	}
}

// Test TestListDatasetsComplexFilter
func TestListDatasetsComplexFilter(t *testing.T) {
	ds, err := createLookupDataset(t, makeDSName("cxfil"))
	require.NoError(t, err)
	defer cleanupDataset(t, ds.LookupDataset().Id)

	req := &catalog.ListDatasetsQueryParams{Filter: "kind==\"kvcollection\" AND name==\"test_externalName\""}
	datasets, err := getClient(t).CatalogService.ListDatasets(req)
	assert.Emptyf(t, err, "Error retrieving the datasets: %s", err)
	assert.NotNil(t, len(datasets))
}

// Test TestListDatasetsCount
func TestListDatasetsCount(t *testing.T) {
	// Create three datasets
	ds1, err := createLookupDataset(t, makeDSName("cnt1"))
	require.NoError(t, err)
	defer cleanupDataset(t, ds1.LookupDataset().Id)
	ds2, err := createLookupDataset(t, makeDSName("cnt2"))
	require.NoError(t, err)
	defer cleanupDataset(t, ds2.LookupDataset().Id)
	ds3, err := createLookupDataset(t, makeDSName("cnt3"))
	require.NoError(t, err)
	defer cleanupDataset(t, ds3.LookupDataset().Id)

	// There should be at least three
	query := catalog.ListDatasetsQueryParams{}.SetCount(3)
	datasets, err := getSdkClient(t).CatalogService.ListDatasets(&query)
	require.NoError(t, err)
	require.Equal(t, 3, len(datasets))
}

//Test TestListDatasetsOrderBy
func TestListDatasetsOrderBy(t *testing.T) {
	ds, err := createLookupDataset(t, makeDSName("orby"))
	require.NoError(t, err)
	defer cleanupDataset(t, ds.LookupDataset().Id)

	datasets, err := getSdkClient(t).CatalogService.ListDatasets(&catalog.ListDatasetsQueryParams{Orderby: []string{"id Descending"}})
	assert.NoError(t, err)
	assert.NotZero(t, len(datasets))
}

// Test TestListDatasetsAll with filter, count, and orderby
func TestListDatasetsAll(t *testing.T) {
	ds, err := createViewDataset(t, makeDSName("fco"))
	require.NoError(t, err)
	defer cleanupDataset(t, ds.ViewDataset().Id)

	query := catalog.ListDatasetsQueryParams{}.SetFilter(`kind=="view"`).SetFilter(`id=="` + ds.ViewDataset().Id + "\"").SetOrderby([]string{"id Descending"}).SetCount(1)
	datasets, err := getSdkClient(t).CatalogService.ListDatasets(&query)
	assert.NoError(t, err)
	assert.NotZero(t, len(datasets))
}

// Test GetDataset by id and resource name
func TestGetDatasetByID(t *testing.T) {
	ds, err := createLookupDataset(t, makeDSName("cnt"))
	require.NoError(t, err)
	require.NotNil(t, ds.LookupDataset())
	defer cleanupDataset(t, ds.LookupDataset().Id)

	ds1, err := getSdkClient(t).CatalogService.GetDatasetById(ds.LookupDataset().Id, nil)
	require.NoError(t, err)
	require.NotNil(t, ds1.LookupDataset())
	assert.Equal(t, ds.LookupDataset().Name, ds1.LookupDataset().Name)
	ds2, err := getSdkClient(t).CatalogService.GetDataset(ds.LookupDataset().Module+"."+ds.LookupDataset().Name, nil)
	require.NoError(t, err)
	require.NotNil(t, ds2.LookupDataset())
	assert.Equal(t, ds.LookupDataset().Name, ds2.LookupDataset().Name)
}

// Test GetDataset for 404 DatasetInfo not found error
func TestGetDatasetByIDDatasetNotFoundError(t *testing.T) {
	_, err := getSdkClient(t).CatalogService.GetDataset("idonotexist", nil)
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
}

// Test UpdateIndexDataset
func TestUpdateIndexDataset(t *testing.T) {
	client := getSdkClient(t)

	indexds, err := createIndexDataset(t, makeDSName("uix"))
	require.NoError(t, err)
	defer cleanupDataset(t, indexds.IndexDataset().Id)
	require.NotNil(t, indexds.IndexDataset())
	notdisabled := !disabled
	uidx := catalog.IndexDatasetPatch{
		Disabled:               &notdisabled,
		FrozenTimePeriodInSecs: &newftime,
	}
	newindexds, err := client.CatalogService.UpdateDataset(
		indexds.IndexDataset().Module+"."+indexds.IndexDataset().Name,
		catalog.MakeDatasetPatchFromIndexDatasetPatch(uidx))
	require.NoError(t, err)
	assert.Equal(t, !disabled, newindexds.IndexDataset().Disabled)
	assert.Equal(t, newftime, *newindexds.IndexDataset().FrozenTimePeriodInSecs)

	uidx1 := catalog.IndexDatasetPatch{
		Disabled:               &disabled,
		FrozenTimePeriodInSecs: &frozenTimePeriodInSecs,
	}
	newindexds1, err := client.CatalogService.UpdateDatasetById(
		indexds.IndexDataset().Id,
		catalog.MakeDatasetPatchFromIndexDatasetPatch(uidx1))
	require.NoError(t, err)
	assert.Equal(t, disabled, newindexds1.IndexDataset().Disabled)
	assert.Equal(t, frozenTimePeriodInSecs, *newindexds1.IndexDataset().FrozenTimePeriodInSecs)
}

// Test UpdateMetricDataset
func TestUpdateMetricDataset(t *testing.T) {
	client := getSdkClient(t)

	metricds, err := createMetricDataset(t, makeDSName("umx"))
	require.NoError(t, err)
	defer cleanupDataset(t, metricds.MetricDataset().Id)
	require.NotNil(t, metricds)
	notdisabled := !disabled
	// Update the metrics dataset
	// N.B.: Name and module not allowed to update
	umx := catalog.MetricDatasetPatch{
		Owner:                  &newowner,
		Disabled:               &notdisabled,
		FrozenTimePeriodInSecs: &newftime,
	}
	newmetricsds, err := client.CatalogService.UpdateDataset(metricds.MetricDataset().Id, catalog.MakeDatasetPatchFromMetricDatasetPatch(umx))
	require.NoError(t, err)
	assert.Equal(t, newowner, newmetricsds.MetricDataset().Owner)
	assert.Equal(t, !disabled, newmetricsds.MetricDataset().Disabled)
	assert.True(t, newftime == *newmetricsds.MetricDataset().FrozenTimePeriodInSecs)
}

//
// ToDo: Needs investigation to see if this works now
// Test UpdateImportDataset
// func TestUpdateImportDataset(t *testing.T) {
// 	client := getSdkClient(t)
//
//	 newmetricsds, err := createMetricDataset(t, makeDSName("umx"))
//	 require.NoError(t, err)
//	 defer cleanupDataset(t, newmetricsds.Id)
//	 require.NotNil(t, newmetricsds)
//
//	 newindexds, err := createIndexDataset(t, makeDSName("indx"))
//	 require.NoError(t, err)
//	 defer cleanupDataset(t, newindexds.Id)
//	 require.NotNil(t, newindexds)
//
//
// 	importds, err := createImportDatasetByName(t, makeDSName("uim"), newmetricsds.Name, newmetricsds.Module)
// 	require.NoError(t, err)
// 	defer cleanupDataset(t, importds.Id)
// 	require.NotNil(t, importds)
// 	uim := &catalog.ImportDatasetPatch{
// 		Name:   &newindexds.Name,
// 		Module: &newindexds.Module,
// 	}
// 	newimportds, err := client.CatalogService.UpdateImportDataset(uim, importds.Id)
// 	require.NoError(t, err)
// 	assert.Equal(t, newindexds.Name, newimportds.SourceName)
// 	assert.Equal(t, newindexds.Module, newimportds.SourceModule)
// }

// Test UpdateJobDataset
func TestUpdateJobDataset(t *testing.T) {
	client := getSdkClient(t)

	// Create a search job to ensure at least one job exists
	searchjobReq := search.SearchJob{Query: "| from index:main | head 1"}
	searchjob, err := client.SearchService.CreateJob(searchjobReq)
	require.NoError(t, err)
	query := catalog.ListDatasetsQueryParams{}.SetFilter(fmt.Sprintf(`sid=="%s"`, *searchjob.Sid)).SetCount(1)
	datasets, err := getSdkClient(t).CatalogService.ListDatasets(&query)
	require.NoError(t, err)
	require.NotZero(t, len(datasets))
	jobds := datasets[0].JobDatasetGet()

	require.NotNil(t, jobds.Id)
	newstatus := string(search.SearchStatusCanceled)

	// This job should not be canceled since it was just created
	require.NotEqual(t, newstatus, *jobds.Status)
}

// Test UpdateLookupDataset
func TestUpdateLookupDataset(t *testing.T) {
	client := getSdkClient(t)

	lookupds, err := createLookupDataset(t, makeDSName("ulk"))
	require.NoError(t, err)
	defer cleanupDataset(t, lookupds.LookupDataset().Id)
	require.NotNil(t, lookupds)
	notcasematch := !caseMatch
	newxname := "newxname"
	filter := `kind=="lookup"`
	owner := "test1@splunk.com"
	ulk := catalog.LookupDatasetPatch{
		Owner:              &owner,
		CaseSensitiveMatch: &notcasematch,
		ExternalName:       &newxname,
		Filter:             &filter,
	}
	newlookupds, err := client.CatalogService.UpdateDataset(lookupds.LookupDataset().Id, catalog.MakeDatasetPatchFromLookupDatasetPatch(ulk))
	require.NoError(t, err)
	assert.NotEqual(t, "cantchangethis", newlookupds.LookupDataset().Id)
	assert.NotEqual(t, "cantchangethat", newlookupds.LookupDataset().Kind)
	assert.Equal(t, "test1@splunk.com", newlookupds.LookupDataset().Owner)
	assert.Equal(t, !caseMatch, *newlookupds.LookupDataset().CaseSensitiveMatch)
	assert.Equal(t, newxname, newlookupds.LookupDataset().ExternalName)
	assert.Equal(t, `kind=="lookup"`, *newlookupds.LookupDataset().Filter)
}

// Test UpdateViewDataset
func TestUpdateViewDataset(t *testing.T) {
	client := getSdkClient(t)

	viewds, err := createViewDataset(t, makeDSName("uvw"))
	require.NoError(t, err)
	defer cleanupDataset(t, viewds.ViewDataset().Id)
	require.NotNil(t, viewds)
	newname = fmt.Sprintf("newvw%d", testutils.RunSuffix)
	uvw := catalog.ViewDatasetPatch{
		Name:   &newname,
		Module: &newmod,
		Owner:  &newowner,
	}
	newviewds, err := client.CatalogService.UpdateDataset(viewds.ViewDataset().Id, catalog.MakeDatasetPatchFromViewDatasetPatch(uvw))
	require.NoError(t, err)
	assert.Equal(t, newname, newviewds.ViewDataset().Name)
	assert.Equal(t, newmod, newviewds.ViewDataset().Module)
	assert.Equal(t, newowner, newviewds.ViewDataset().Owner)
}

// Test UpdateDataset for 404 Datasetnot found error
func TestUpdateExistingDatasetDataNotFoundError(t *testing.T) {
	uvw := catalog.ViewDatasetPatch{
		Search: &searchString,
	}
	_, err := getSdkClient(t).CatalogService.UpdateDataset("idonotexist", catalog.MakeDatasetPatchFromViewDatasetPatch(uvw))
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
}

// Test DeleteDataset
func TestDeleteDataset(t *testing.T) {
	client := getSdkClient(t)

	ds, err := createViewDataset(t, makeDSName("delv"))
	require.NoError(t, err)

	err = client.CatalogService.DeleteDataset(ds.ViewDataset().Id)
	require.NoError(t, err)

	_, err = client.CatalogService.GetDataset(ds.ViewDataset().Id, nil)
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
}

// Test DeleteDataset for 404 DatasetInfo not found error
func TestDeleteDatasetDataNotFoundError(t *testing.T) {
	err := getSdkClient(t).CatalogService.DeleteDataset("idonotexist")
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
}

// Test CreateRules
func TestCreateRules(t *testing.T) {
	client := getSdkClient(t)

	// create rule
	ruleName := makeRuleName("crmat")
	rule, err := client.CatalogService.CreateRule(catalog.RulePost{Name: ruleName, Module: &ruleModule, Match: ruleMatch})
	require.NoError(t, err)
	defer cleanupRule(t, rule.Id)
	assert.Equal(t, ruleName, rule.Name)
	assert.Equal(t, ruleMatch, rule.Match)
}

// Test UpdateRule
func TestUpdateRule(t *testing.T) {
	client := getSdkClient(t)

	// create rule
	ruleName := makeRuleName("crmatu")
	rule, err := client.CatalogService.CreateRule(catalog.RulePost{Name: ruleName, Module: &ruleModule, Match: ruleMatch})
	require.NoError(t, err)
	defer cleanupRule(t, rule.Id)

	mat := `sourcetype::new_sourcetype`
	urule, err := client.CatalogService.UpdateRule(rule.Module+"."+rule.Name, catalog.RulePatch{Match: &mat})
	assert.Equal(t, ruleName, urule.Name)
	assert.Equal(t, mat, urule.Match)
}

// Test UpdateRuleById
func TestUpdateRuleById(t *testing.T) {
	client := getSdkClient(t)

	// create rule
	ruleName := makeRuleName("crmatu")
	rule, err := client.CatalogService.CreateRule(catalog.RulePost{Name: ruleName, Module: &ruleModule, Match: ruleMatch})
	require.NoError(t, err)
	defer cleanupRule(t, rule.Id)

	mat := `sourcetype::new_sourcetype`
	urule, err := client.CatalogService.UpdateRuleById(rule.Id, catalog.RulePatch{Match: &mat})
	assert.Equal(t, ruleName, urule.Name)
	assert.Equal(t, mat, urule.Match)
}

// Test CreateRule for 409 Rule already present error
func TestCreateRuleDataAlreadyPresent(t *testing.T) {
	client := getSdkClient(t)

	// create rule
	ruleName := makeRuleName("409")
	rule, err := client.CatalogService.CreateRule(catalog.RulePost{Name: ruleName, Module: &ruleModule, Match: ruleMatch})
	require.NoError(t, err)
	defer cleanupRule(t, rule.Id)
	assert.Equal(t, ruleName, rule.Name)
	assert.Equal(t, ruleMatch, rule.Match)

	_, err = client.CatalogService.CreateRule(catalog.RulePost{Id: &rule.Id, Name: ruleName, Module: &ruleModule, Match: ruleMatch})
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 409, httpErr.HTTPStatusCode)
}

// Test CreateRule for 401 Unauthorized operation error
func TestCreateRuleUnauthorizedOperationError(t *testing.T) {
	// create rule
	ruleName := makeRuleName("401")
	rule, err := getInvalidClient(t).CatalogService.CreateRule(catalog.RulePost{Name: ruleName, Module: &ruleModule, Match: ruleMatch})
	if rule != nil {
		defer cleanupRule(t, rule.Id)
	}
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 401, httpErr.HTTPStatusCode)
}

// Test GetRules
func TestGetAllRules(t *testing.T) {
	client := getSdkClient(t)

	// create rule
	ruleName := makeRuleName("getall")
	rule, err := client.CatalogService.CreateRule(catalog.RulePost{Name: ruleName, Module: &ruleModule, Match: ruleMatch})
	require.NoError(t, err)
	defer cleanupRule(t, rule.Id)

	rules, err := client.CatalogService.ListRules(nil)
	require.NoError(t, err)
	assert.NotZero(t, len(rules))
}

// Test GetRule By ID
func TestGetRuleByID(t *testing.T) {
	client := getSdkClient(t)

	// create rule
	ruleName := makeRuleName("getid")
	rule, err := client.CatalogService.CreateRule(catalog.RulePost{Name: ruleName, Module: &ruleModule, Match: ruleMatch})
	require.NoError(t, err)
	require.NotEmpty(t, rule.Id)

	ruleByID, err := client.CatalogService.GetRule(rule.Id)
	assert.NoError(t, err)
	assert.NotNil(t, ruleByID)
}

// Test GetRules for 404 Rule not found error
func TestGetRuleByIDRuleNotFoundError(t *testing.T) {
	_, err := getSdkClient(t).CatalogService.GetRule("idonotexist")
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
}

// Test DeleteRule by ID
func TestDeleteRuleById(t *testing.T) {
	client := getSdkClient(t)

	// create rule
	ruleName := makeRuleName("delid")
	rule, err := client.CatalogService.CreateRule(catalog.RulePost{Name: ruleName, Module: &ruleModule, Match: ruleMatch})
	require.NoError(t, err)
	require.NotEmpty(t, rule.Id)

	err = client.CatalogService.DeleteRuleById(rule.Id)
	assert.NoError(t, err)

	var resp http.Response
	r, err := client.CatalogService.GetRuleById(rule.Id, &resp)
	assert.Nil(t, r)
	assert.NotNil(t, err)
	assert.Equal(t, 404, resp.StatusCode)
}

// Test DeleteRule by ID
func TestDeleteRule(t *testing.T) {
	client := getSdkClient(t)

	// create rule
	ruleName := makeRuleName("delrl")
	rule, err := client.CatalogService.CreateRule(catalog.RulePost{Name: ruleName, Module: &ruleModule, Match: ruleMatch})
	require.NoError(t, err)
	require.NotEmpty(t, rule.Id)

	err = client.CatalogService.DeleteRule(ruleModule + "." + ruleName)
	assert.NoError(t, err)

	var resp http.Response
	r, err := client.CatalogService.GetRuleById(rule.Id, &resp)
	assert.Nil(t, r)
	assert.NotNil(t, err)
	assert.Equal(t, 404, resp.StatusCode)
}

// Test DeleteRule for 404 Rule not found error
func TestDeleteRulebyIDRuleNotFoundError(t *testing.T) {
	err := getSdkClient(t).CatalogService.DeleteRule("idonotexist")
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
}

// Test TestDatasetFieldGetList
func TestDatasetFieldGetList(t *testing.T) {
	client := getSdkClient(t)

	// Create dataset
	dsname := makeDSName("getfs")
	fieldName1 := "integ_test_field1"
	fieldName2 := "integ_test_field2"
	dataType := catalog.FieldDataTypeString
	fieldType := catalog.FieldTypeDimension
	prevalenceType := catalog.FieldPrevalenceAll

	d, err := createLookupDataset(t, dsname)
	require.NoError(t, err)
	require.NotNil(t, d.LookupDataset())
	ds := d.LookupDataset()
	defer cleanupDataset(t, ds.Id)

	// create new fields in the dataset
	// Initial property values for creating a new field using a POST request.

	testField1 := catalog.FieldPost{Name: fieldName1, Datatype: &dataType, Fieldtype: &fieldType, Prevalence: &prevalenceType}
	testField2 := catalog.FieldPost{Name: fieldName2, Datatype: &dataType, Fieldtype: &fieldType, Prevalence: &prevalenceType}
	f1, err := client.CatalogService.CreateFieldForDataset(ds.Id, testField1)
	f2, err := client.CatalogService.CreateFieldForDataset(ds.Id, testField2)
	require.NoError(t, err)
	// Get
	results, err := client.CatalogService.GetFieldByIdForDataset(ds.Id, f1.Id)
	require.NoError(t, err)
	require.NotNil(t, results)
	require.NotNil(t, f1)
	require.NotNil(t, f2)
	assert.Equal(t, testField1.Name, f1.Name)
	assert.Equal(t, testField2.Name, f2.Name)
	res1, err := client.CatalogService.GetFieldByIdForDataset(ds.Module+"."+ds.Name, f1.Id)
	require.NoError(t, err)
	assert.NotNil(t, res1)
	assert.Equal(t, testField1.Name, res1.Name)
	res2, err := client.CatalogService.GetFieldByIdForDataset(ds.Id, f2.Id)
	require.NoError(t, err)
	assert.NotNil(t, res2)
	assert.Equal(t, testField2.Name, res2.Name)
	fdid1, err := client.CatalogService.GetFieldByIdForDatasetById(ds.Id, f1.Id)
	require.NoError(t, err)
	assert.NotNil(t, fdid1)
	assert.Equal(t, testField1.Name, fdid1.Name)
	fdid2, err := client.CatalogService.GetFieldByIdForDatasetById(ds.Id, f2.Id)
	require.NoError(t, err)
	assert.NotNil(t, fdid2)
	assert.Equal(t, testField2.Name, fdid2.Name)
	fr1, err := client.CatalogService.GetFieldById(f1.Id)
	require.NoError(t, err)
	assert.NotNil(t, fr1)
	assert.Equal(t, testField1.Name, fr1.Name)
	fr2, err := client.CatalogService.GetFieldById(f2.Id)
	require.NoError(t, err)
	assert.NotNil(t, fr2)
	assert.Equal(t, testField2.Name, fr2.Name)
	// List
	fs, err := client.CatalogService.ListFields(nil)
	require.NoError(t, err)
	assert.True(t, len(fs) > 0)
	fs, err = client.CatalogService.ListFieldsForDataset(ds.Module+"."+ds.Name, nil)
	require.NoError(t, err)
	assert.True(t, len(fs) > 0)
	fs, err = client.CatalogService.ListFieldsForDatasetById(ds.Id, nil)
	require.NoError(t, err)
	assert.True(t, len(fs) > 0)
}

// Test CreateDatasetField
func TestCreateDatasetField(t *testing.T) {
	client := getSdkClient(t)

	// Create dataset
	dsname := makeDSName("crf")
	fieldName := "integ_test_field"
	dataType := catalog.FieldDataTypeString
	fieldType := catalog.FieldTypeDimension
	prevalenceType := catalog.FieldPrevalenceAll

	ds, err := createLookupDataset(t, dsname)
	require.NoError(t, err)
	defer cleanupDataset(t, ds.LookupDataset().Id)

	// Create a new dataset field
	testField := catalog.FieldPost{Name: fieldName, Datatype: &dataType, Fieldtype: &fieldType, Prevalence: &prevalenceType}
	resultField, err := client.CatalogService.CreateFieldForDataset(ds.LookupDataset().Id, testField)
	require.NoError(t, err)
	require.NotEmpty(t, resultField)
	assert.Equal(t, "integ_test_field", resultField.Name)
	// TODO: catalog.String, Dimension, and All do not match "S", "D", "A" - why is this?
	assert.Equal(t, catalog.FieldDataTypeString, resultField.Datatype)
	assert.Equal(t, catalog.FieldTypeDimension, resultField.Fieldtype)
	assert.Equal(t, catalog.FieldPrevalenceAll, resultField.Prevalence)
	assert.Emptyf(t, err, "Error creating dataset field: %s", err)

	// Validate the creation of a new dataset field
	resultField, err = client.CatalogService.GetFieldByIdForDataset(ds.LookupDataset().Id, resultField.Id)
	require.NoError(t, err)
	assert.NotEmpty(t, resultField)
}

// Test PatchDatasetField
func TestPatchDatasetField(t *testing.T) {
	client := getSdkClient(t)

	// Create dataset
	dsname := makeDSName("patf")
	ds, err := createLookupDataset(t, dsname)
	require.NoError(t, err)
	defer cleanupDataset(t, ds.LookupDataset().Id)

	// Create a new dataset field
	resultField := createDatasetField(ds.LookupDataset().Id, client, t)

	// Update the existing dataset field
	dataType := catalog.FieldDataTypeString
	resultField, err = client.CatalogService.UpdateFieldByIdForDataset(ds.LookupDataset().Id, resultField.Id, catalog.FieldPatch{Datatype: &dataType})
	require.NoError(t, err)
	require.NotNil(t, resultField)
	assert.Equal(t, "integ_test_field", resultField.Name)
	assert.Equal(t, catalog.FieldDataTypeString, resultField.Datatype)
	assert.Equal(t, catalog.FieldTypeDimension, resultField.Fieldtype)
	assert.Equal(t, catalog.FieldPrevalenceAll, resultField.Prevalence)

	dataType = catalog.FieldDataTypeNumber
	resultField, err = client.CatalogService.UpdateFieldByIdForDatasetById(ds.LookupDataset().Id, resultField.Id, catalog.FieldPatch{Datatype: &dataType})
	require.NoError(t, err)
	require.NotNil(t, resultField)
	assert.Equal(t, "integ_test_field", resultField.Name)
	assert.Equal(t, catalog.FieldDataTypeNumber, resultField.Datatype)
	assert.Equal(t, catalog.FieldTypeDimension, resultField.Fieldtype)
	assert.Equal(t, catalog.FieldPrevalenceAll, resultField.Prevalence)
}

// Test DeleteDatasetField
func TestDeleteDatasetField(t *testing.T) {
	client := getSdkClient(t)

	// Create dataset
	dsname := makeDSName("delf")
	ds, err := createLookupDataset(t, dsname)
	require.NoError(t, err)
	defer cleanupDataset(t, ds.LookupDataset().Id)

	// Create a new dataset field
	resultField := createDatasetField(ds.LookupDataset().Id, client, t)

	// Delete dataset field
	err = client.CatalogService.DeleteFieldByIdForDataset(ds.LookupDataset().Id, resultField.Id)
	require.NoError(t, err)

	// Validate the deletion of the dataset field
	_, err = client.CatalogService.GetFieldByIdForDataset(ds.LookupDataset().Id, resultField.Id)
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
}

// Test CreateDatasetField for 401 error
func TestCreateDatasetFieldUnauthorizedOperationError(t *testing.T) {
	invalidClient := getInvalidClient(t)

	// Create dataset
	dsname := makeDSName("f401")

	fieldName := "integ_test_field1"
	dataType := catalog.FieldDataTypeString
	fieldType := catalog.FieldTypeDimension
	prevalenceType := catalog.FieldPrevalenceAll

	ds, err := createLookupDataset(t, dsname)
	require.NoError(t, err)
	defer cleanupDataset(t, ds.LookupDataset().Id)

	// Create a new dataset field
	testField := catalog.FieldPost{Name: fieldName, Datatype: &dataType, Fieldtype: &fieldType, Prevalence: &prevalenceType}
	_, err = invalidClient.CatalogService.CreateFieldForDatasetById(ds.LookupDataset().Id, testField)
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 401, httpErr.HTTPStatusCode)
}

// Test CreateDatasetField for 409 error
func TestCreateDatasetFieldDataAlreadyPresent(t *testing.T) {

	client := getSdkClient(t)

	// Create dataset
	dsname := makeDSName("f409")
	fieldName := "integ_test_field"
	dataType := catalog.FieldDataTypeString
	fieldType := catalog.FieldTypeDimension
	prevalenceType := catalog.FieldPrevalenceAll
	ds, err := createLookupDataset(t, dsname)
	require.NoError(t, err)
	defer cleanupDataset(t, ds.LookupDataset().Id)

	// Create a new dataset field
	createDatasetField(ds.LookupDataset().Id, client, t)

	// Post an already created dataset field
	duplicateTestField := catalog.FieldPost{Name: fieldName, Datatype: &dataType, Fieldtype: &fieldType, Prevalence: &prevalenceType}
	fmt.Println(&duplicateTestField.Name)
	_, err = client.CatalogService.CreateFieldForDatasetById(ds.LookupDataset().Id, duplicateTestField)
	fmt.Println(err)
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 409, httpErr.HTTPStatusCode)
}

// Test CreateDatasetField for 400 error
func TestCreateDatasetFieldInvalidDataFormat(t *testing.T) {
	t.Skip("SCP-24776")

	client := getSdkClient(t)

	// Create dataset
	dsname := makeDSName("f400")
	ds, err := createLookupDataset(t, dsname)
	require.NoError(t, err)
	defer cleanupDataset(t, ds.LookupDataset().Id)

	// Create a new dataset field but with no data
	testField := catalog.FieldPost{}
	_, err = client.CatalogService.CreateFieldForDatasetById(ds.LookupDataset().Id, testField)
	require.Error(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 400, httpErr.HTTPStatusCode)
}

// Test PatchDatasetField for 404 error
func TestPatchDatasetFieldDataNotFound(t *testing.T) {
	client := getSdkClient(t)

	// Create dataset
	dsname := makeDSName("fp404")
	datatype := catalog.FieldDataTypeString
	ds, err := createLookupDataset(t, dsname)
	require.NoError(t, err)
	defer cleanupDataset(t, ds.LookupDataset().Id)

	// Update non-existent dataset field
	_, err = client.CatalogService.UpdateFieldByIdForDataset("idonotexist", ds.LookupDataset().Id, catalog.FieldPatch{Datatype: &datatype})
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
}

// Test DeleteDatasetField for 404 error
func TestDeleteDatasetFieldDataNotFound(t *testing.T) {
	client := getSdkClient(t)

	// Create dataset
	dsname := makeDSName("fd404")
	ds, err := createLookupDataset(t, dsname)
	require.NoError(t, err)
	defer cleanupDataset(t, ds.LookupDataset().Id)

	// Delete dataset field
	err = client.CatalogService.DeleteFieldByIdForDatasetById(ds.LookupDataset().Id, "idonotexist")
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
}

// Test rule actions endpoints
func TestRuleActions(t *testing.T) {
	client := getSdkClient(t)

	// Create dataset
	dsname := makeDSName("acts")
	ds, err := createLookupDataset(t, dsname)
	require.NoError(t, err)
	defer cleanupDataset(t, ds.LookupDataset().Id)

	// create new field in the dataset
	field := createDatasetField(ds.LookupDataset().Id, client, t)

	// Create rule and rule action
	ruleName := makeRuleName("acts")
	rule, err := client.CatalogService.CreateRule(catalog.RulePost{Name: ruleName, Module: &ruleModule, Match: ruleMatch})
	require.NoError(t, err)
	defer cleanupRule(t, rule.Id)

	// Create rule action - alias
	alias := "myfieldalias"
	aap := catalog.AliasActionPost{Field: field.Name, Alias: alias, Kind: catalog.AliasActionKindAlias}
	action1, err := client.CatalogService.CreateActionForRule(rule.Module+"."+rule.Name, catalog.MakeActionPostFromAliasActionPost(aap))
	require.NoError(t, err)
	defer cleanupRuleAction(t, rule.Id, action1.AliasAction().Id)

	//update rule action by id
	tmpstr := "newaliasi"
	updateact1, err := client.CatalogService.UpdateActionByIdForRuleById(rule.Id, action1.AliasAction().Id,
		catalog.MakeActionPatchFromAliasActionPatch(catalog.AliasActionPatch{Alias: &tmpstr}))
	require.NotNil(t, updateact1)
	assert.Equal(t, tmpstr, updateact1.AliasAction().Alias)

	//update rule action by resourcename
	tmpstr2 := "newaliasii"
	updateact1a, err := client.CatalogService.UpdateActionByIdForRule(rule.Module+"."+rule.Name, action1.AliasAction().Id,
		catalog.MakeActionPatchFromAliasActionPatch(catalog.AliasActionPatch{Alias: &tmpstr2}))
	require.NotNil(t, updateact1a)

	assert.Equal(t, tmpstr2, updateact1a.AliasAction().Alias)

	// Create rule action - autokv
	mode := "auto"
	actionPost := catalog.AutoKvActionPost{Mode: mode, Kind: catalog.AutoKvActionKindAutokv}
	action2, err := client.CatalogService.CreateActionForRule(rule.Module+"."+rule.Name, catalog.MakeActionPostFromAutoKvActionPost(actionPost))
	require.NoError(t, err)
	defer cleanupRuleAction(t, rule.Id, action2.AutoKvAction().Id)

	//update rule action
	tmpstr = "json"
	updateact2, err := client.CatalogService.UpdateActionByIdForRuleById(rule.Id, action2.AutoKvAction().Id,
		catalog.MakeActionPatchFromAutoKvActionPatch(catalog.AutoKvActionPatch{Mode: &tmpstr}))
	require.NotNil(t, updateact2)
	assert.Equal(t, tmpstr, updateact2.AutoKvAction().Mode)

	// Create rule action - eval
	expr := "\"some expression\""
	eap := catalog.EvalActionPost{Field: field.Name, Expression: expr, Kind: catalog.EvalActionKindEval}
	action3, err := client.CatalogService.CreateActionForRule(rule.Module+"."+rule.Name, catalog.MakeActionPostFromEvalActionPost(eap))
	require.NoError(t, err)
	defer cleanupRuleAction(t, rule.Id, action3.EvalAction().Id)

	//update rule eval action
	tmpstr = "newField"
	updateact3, err := client.CatalogService.UpdateActionByIdForRuleById(rule.Id, action3.EvalAction().Id,
		catalog.MakeActionPatchFromEvalActionPatch(catalog.EvalActionPatch{Field: &tmpstr}))
	require.NotNil(t, updateact3)
	assert.Equal(t, tmpstr, updateact3.EvalAction().Field)

	// Create rule action - lookup
	expr2 := "myexpression2"
	lap := catalog.LookupActionPost{Expression: expr2, Kind: catalog.LookupActionKindLookup}
	action4, err := client.CatalogService.CreateActionForRule(rule.Module+"."+rule.Name, catalog.MakeActionPostFromLookupActionPost(lap))
	require.NoError(t, err)
	defer cleanupRuleAction(t, rule.Id, action4.LookupAction().Id)

	//update rule action
	tmpstr = "newexpr"
	updateact4, err := client.CatalogService.UpdateActionByIdForRuleById(rule.Id, action4.LookupAction().Id,
		catalog.MakeActionPatchFromLookupActionPatch(catalog.LookupActionPatch{Expression: &tmpstr}))
	require.NotNil(t, updateact4)
	assert.Equal(t, tmpstr, updateact4.LookupAction().Expression)

	// Create rule action - regex
	limit := int32(5)
	pattern := `field=myfield "From: (?<from>.*) To: (?<to>.*)"`
	rap := catalog.RegexActionPost{Field: field.Name, Pattern: pattern, Limit: &limit, Kind: catalog.RegexActionKindRegex}
	action5, err := client.CatalogService.CreateActionForRule(rule.Module+"."+rule.Name, catalog.MakeActionPostFromRegexActionPost(rap))
	require.NoError(t, err)
	assert.True(t, *action5.RegexAction().Limit == 5)
	defer cleanupRuleAction(t, rule.Id, action5.RegexAction().Id)

	rap = catalog.RegexActionPost{Field: field.Name, Pattern: pattern, Limit: nil, Kind: catalog.RegexActionKindRegex}
	action6, err := client.CatalogService.CreateActionForRule(rule.Module+"."+rule.Name, catalog.MakeActionPostFromRegexActionPost(rap))
	require.NoError(t, err)
	assert.Equal(t, (*int32)(nil), action6.RegexAction().Limit)
	defer cleanupRuleAction(t, rule.Id, action6.RegexAction().Id)

	//update rule action
	tmpstr = `field=myotherfield "To: (?<to>.*) From: (?<from>.*)"`
	fieldnew := "something"
	updateact5, err := client.CatalogService.UpdateActionByIdForRuleById(rule.Id, action5.RegexAction().Id,
		catalog.MakeActionPatchFromRegexActionPatch(catalog.RegexActionPatch{Pattern: &tmpstr, Field: &fieldnew}))
	require.NotNil(t, updateact5)
	assert.Equal(t, tmpstr, updateact5.RegexAction().Pattern)
	assert.Equal(t, fieldnew, updateact5.RegexAction().Field)

	//Get rule actions
	actions, err := client.CatalogService.ListActionsForRuleById(rule.Id, nil)
	require.NotNil(t, actions)
	assert.Equal(t, 6, len(actions))

	// List actions for rule
	actions1, err := client.CatalogService.ListActionsForRule(rule.Module+"."+rule.Name, nil)
	require.NotNil(t, actions1)
	assert.Equal(t, 6, len(actions1))

	action7, err := client.CatalogService.GetActionByIdForRuleById(rule.Id, action1.AliasAction().Id)
	require.NotNil(t, action7)
	require.NotNil(t, action7.AliasAction())

	action8, err := client.CatalogService.GetActionByIdForRule(rule.Module+"."+rule.Name, action1.AliasAction().Id)
	require.NotNil(t, action8)
	require.NotNil(t, action8.AliasAction())
	assert.Equal(t, action7.AliasAction().Id, action8.AliasAction().Id)
	assert.Equal(t, action7.AliasAction().Created, action8.AliasAction().Created)
	assert.Equal(t, action7.AliasAction().Modifiedby, action8.AliasAction().Modifiedby)

	// Delete action
	err = client.CatalogService.DeleteActionByIdForRule(rule.Module+"."+rule.Name, action6.RegexAction().Id)
	require.NoError(t, err)
}

func TestCreateActionForRuleById(t *testing.T) {
	client := getSdkClient(t)

	// Create dataset
	dsname := makeDSName("actrid")
	ds, err := createLookupDataset(t, dsname)
	require.NoError(t, err)
	defer cleanupDataset(t, ds.LookupDataset().Id)

	// create new field in the dataset
	field := createDatasetField(ds.LookupDataset().Id, client, t)

	// Create rule and rule action
	ruleName := makeRuleName("actid")
	rule, err := client.CatalogService.CreateRule(catalog.RulePost{Name: ruleName, Module: &ruleModule, Match: ruleMatch})
	require.NoError(t, err)
	defer cleanupRule(t, rule.Id)

	// Create rule action - alias
	alias := "myfieldalias2"
	aap := catalog.AliasActionPost{Field: field.Name, Alias: alias, Kind: catalog.AliasActionKindAlias}
	action1, err := client.CatalogService.CreateActionForRuleById(rule.Id, catalog.MakeActionPostFromAliasActionPost(aap))
	require.NoError(t, err)
	defer cleanupRuleAction(t, rule.Id, action1.AliasAction().Id)
}

// Test list modules
func TestGetModules(t *testing.T) {
	client := getSdkClient(t)

	// test using NO filter
	modules, err := client.CatalogService.ListModules(nil)
	require.NoError(t, err)
	assert.True(t, len(modules) > 0)

	// test using filter
	filter := make(url.Values)
	filter.Add("filter", `module==""`)
	modules, err = client.CatalogService.ListModules(&catalog.ListModulesQueryParams{Filter: `module==""`})
	require.NoError(t, err)
	assert.Equal(t, 1, len(modules))
	assert.Equal(t, "", *modules[0].Name)
}

// Test dashboard
func TestCRUDDashboard(t *testing.T) {
	client := getSdkClient(t)
	dashboardName := makeDSName("dashboard")
	module := "allmembers"

	//create
	db, err := client.CatalogService.CreateDashboard(catalog.DashboardPost{
		Name:       dashboardName,
		Module:     module,
		Definition: "{\"title\":\"this is my test dashboard\"}"})

	defer client.CatalogService.DeleteDashboardById(db.Id)
	require.NoError(t, err)
	assert.Equal(t, dashboardName, db.Name)
	assert.Equal(t, module, db.Module)

	//List
	db, err = client.CatalogService.GetDashboardById(db.Id)
	require.NoError(t, err)
	assert.Equal(t, dashboardName, db.Name)

	db, err = client.CatalogService.GetDashboardByResourceName(module + "." + dashboardName)
	require.NoError(t, err)
	assert.Equal(t, dashboardName, db.Name)

	//ListAll
	count := int32(10)
	dbs, err := client.CatalogService.ListDashboards(&catalog.ListDashboardsQueryParams{Count: &count, Filter: `module=="allmembers"`, Orderby: []string{"id Descending"}})
	require.NoError(t, err)
	assert.NotZero(t, len(dbs))

	//Update
	name_new := dashboardName + "_updated"
	dashboard, err := client.CatalogService.UpdateDashboardById(db.Id, catalog.DashboardPatch{Name: &name_new})
	require.NoError(t, err)
	assert.Equal(t, name_new, dashboard.Name)

	db, err = client.CatalogService.GetDashboardByResourceName(module + "." + name_new)
	assert.Equal(t, name_new, db.Name)

	//Delete
	err = client.CatalogService.DeleteDashboardById(db.Id)
	require.NoError(t, err)

	// Create
	name1 := makeDSName("dashboard")
	db, err = client.CatalogService.CreateDashboard(catalog.DashboardPost{
		Name:       name1,
		Module:     module,
		Definition: "{\"title\":\"this is my other test dashboard\"}"})

	defer client.CatalogService.DeleteDashboardById(db.Id)

	// Update by name
	name_new1 := name1 + "_updated"
	err = client.CatalogService.UpdateDashboardByResourceName(module+"."+name1, catalog.DashboardPatch{Name: &name_new1})
	require.NoError(t, err)
	db, err = client.CatalogService.GetDashboardByResourceName(module + "." + name_new1)
	assert.Equal(t, name_new, db.Name)

	// Delete by name
	err = client.CatalogService.DeleteDashboardByResourceName(module + "." + name_new1)
	require.NoError(t, err)
}

func createWorkflow(client *sdk.Client) (*catalog.Workflow, error) {
	name := makeDSName("workflow")

	return client.CatalogService.CreateWorkflow(catalog.WorkflowPost{
		Name:  &name,
		Tasks: []catalog.TaskPost{}})
}

func createWorkflowBuild(client *sdk.Client, wf *catalog.Workflow) (*catalog.WorkflowBuild, error) {

	return client.CatalogService.CreateWorkflowBuild(
		wf.Id,
		catalog.WorkflowBuildPost{
			Outputdata:  []string{},
			Inputdata:   []string{},
			Timeoutsecs: 10})
}

// Test workflow
func TestCRUDWorkflow(t *testing.T) {
	client := getSdkClient(t)

	wf, err := createWorkflow(client)
	require.NoError(t, err)
	defer client.CatalogService.DeleteWorkflowById(wf.Id)

	//List
	wfg, err := client.CatalogService.GetWorkflowById(wf.Id)
	require.NoError(t, err)
	assert.Equal(t, wf.Id, wfg.Id)

	//ListAll
	dbs, err := client.CatalogService.ListWorkflows(nil)
	require.NoError(t, err)
	assert.True(t, len(dbs) > 0)

	//Update
	name_new := "newname"
	err = client.CatalogService.UpdateWorkflowById(wf.Id, catalog.WorkflowPatch{Name: &name_new})
	require.NoError(t, err)
	wf, err = client.CatalogService.GetWorkflowById(wf.Id)
	assert.Equal(t, name_new, wf.Name)

	//Delete
	err = client.CatalogService.DeleteWorkflowById(wf.Id)
	require.NoError(t, err)
}

// Test workflowBuild
func TestCRUDWorkflowBuild(t *testing.T) {
	client := getSdkClient(t)

	wf, err := createWorkflow(client)
	require.NoError(t, err)
	defer client.CatalogService.DeleteWorkflowById(wf.Id)

	wfb, err := createWorkflowBuild(client, wf)
	require.NoError(t, err)
	defer client.CatalogService.DeleteWorkflowBuildById(wf.Id, wfb.Id)

	//List
	wfg, err := client.CatalogService.GetWorkflowBuildById(wf.Id, wfb.Id)
	require.NoError(t, err)
	assert.Equal(t, wfb.Id, wfg.Id)

	//ListAll
	dbs, err := client.CatalogService.ListWorkflowBuilds(wf.Id, nil)
	require.NoError(t, err)
	assert.True(t, len(dbs) > 0)

	//Update
	name_new := "newname"
	err = client.CatalogService.UpdateWorkflowBuildById(wf.Id, wfb.Id, catalog.WorkflowBuildPatch{Name: &name_new})
	require.NoError(t, err)
	wfb, err = client.CatalogService.GetWorkflowBuildById(wf.Id, wfb.Id)
	assert.Equal(t, name_new, *wfb.Name)

	//Delete
	err = client.CatalogService.DeleteWorkflowBuildById(wf.Id, wfb.Id)
	require.NoError(t, err)
}

// Test workflowRun
func TestCRUDWorkflowRun(t *testing.T) {
	client := getSdkClient(t)

	wf, err := createWorkflow(client)
	require.NoError(t, err)
	defer client.CatalogService.DeleteWorkflowById(wf.Id)

	wfb, err := createWorkflowBuild(client, wf)
	require.NoError(t, err)
	defer client.CatalogService.DeleteWorkflowBuildById(wf.Id, wfb.Id)

	wfr, err := client.CatalogService.CreateWorkflowRun(
		wf.Id,
		wfb.Id,
		catalog.WorkflowRunPost{
			Outputdata:  []string{},
			Inputdata:   []string{},
			Timeoutsecs: 10})

	require.NoError(t, err)
	defer client.CatalogService.DeleteWorkflowRunById(wf.Id, wfb.Id, wfr.Id)

	//List
	wfg, err := client.CatalogService.GetWorkflowRunById(wf.Id, wfb.Id, wfr.Id)
	require.NoError(t, err)
	assert.Equal(t, wfr.Id, wfg.Id)

	//ListAll
	dbs, err := client.CatalogService.ListWorkflowRuns(wf.Id, wfb.Id, nil)
	require.NoError(t, err)
	assert.True(t, len(dbs) > 0)

	//Update
	name_new := "newname"
	err = client.CatalogService.UpdateWorkflowRunById(wf.Id, wfb.Id, wfr.Id, catalog.WorkflowRunPatch{Name: &name_new})
	require.NoError(t, err)
	wfr, err = client.CatalogService.GetWorkflowRunById(wf.Id, wfb.Id, wfr.Id)
	assert.Equal(t, name_new, *wfr.Name)

	//Delete
	err = client.CatalogService.DeleteWorkflowRunById(wf.Id, wfb.Id, wfr.Id)
	require.NoError(t, err)
}

// TODO: tim - CreateRelationship is failing with 500 error - catalog is working the issue
// Test relationships
// func TestCRUDRelationships(t *testing.T) {
// 	client, _, _ := getSdkClientWithLoggers(t)

// 	// Create datasets
// 	dsname1 := makeDSName("rs1")
// 	ds1, err := createIndexDataset(t, dsname1)
// 	require.NoError(t, err)
// 	require.NotNil(t, ds1.GetIndexDataset())
// 	inds1 := ds1.GetIndexDataset()
// 	defer cleanupDataset(t, inds1.Id)
// 	dsname2 := makeDSName("rs2")
// 	ds2, err := createIndexDataset(t, dsname2)
// 	require.NoError(t, err)
// 	require.NotNil(t, ds2.GetIndexDataset())
// 	inds2 := ds2.GetIndexDataset()
// 	defer cleanupDataset(t, inds2.Id)

// 	// Create relationship fields
// 	// f1 := catalog.RelationshipFieldPost{
// 	// 	Kind:     catalog.RelationshipFieldKindExact,
// 	// 	Sourceid: inds1.Id,
// 	// 	Targetid: inds2.Id,
// 	// }
// 	relname := fmt.Sprintf("gorel%d", testutils.RunSuffix)
// 	// src := inds1.Module + "." + inds1.Name
// 	// tgt := inds2.Module + "." + inds2.Name
// 	r := catalog.RelationshipPost{
// 		Kind:     catalog.RelationshipKindDependency,
// 		Name:     relname,
// 		Sourceid: inds1.Id,
// 		Targetid: inds2.Id,
// 		Fields:   []catalog.RelationshipFieldPost{},
// 	}
// 	var resp http.Response
// 	rel, err := client.CatalogService.CreateRelationship(r, &resp)
// 	// TODO: 500 err
// 	require.NoError(t, err)
// 	require.NotNil(t, rel)
// 	// CreateRelationship
// 	// GetRelationshipById
// 	// UpdateRelationshipById
// 	// ListRelationships
// 	// DeleteRelationshipById
// }

// Test annotations
// Test annotations
func TestCRUDAnnotations(t *testing.T) {
	client := getSdkClient(t)
	ds, err := createLookupDataset(t, makeDSName("annx"))
	require.NoError(t, err)
	require.NotNil(t, ds.LookupDataset())
	lds := ds.LookupDataset()

	// There is exactly one that can be used at the moment
	const DefaultAnnotationTypeId = "00000000000000000000008b"
	// Create

	m := map[string]string{"annotationtypeid": DefaultAnnotationTypeId}
	an1, err := client.CatalogService.CreateAnnotationForDatasetById(lds.Id, m)

	var annotationId string
	for k, p := range *an1 {
		if k == "annotationtypeid" {
			annotationId = p.(string)
		}
	}

	require.NoError(t, err)
	require.NotNil(t, annotationId)

	defer client.CatalogService.DeleteAnnotationOfDatasetById(lds.Id, annotationId)
	require.Equal(t, DefaultAnnotationTypeId, annotationId)
	// Create
	an2, err := client.CatalogService.CreateAnnotationForDatasetByResourceName(lds.Module+"."+lds.Name, m)
	require.NoError(t, err)
	for k, p := range *an2 {
		if k == "annotationtypeid" {
			annotationId = p.(string)
		}
	}
	require.NotNil(t, annotationId)
	defer client.CatalogService.DeleteAnnotationOfDatasetById(lds.Id, annotationId)
	require.Equal(t, DefaultAnnotationTypeId, annotationId)

	// List
	ans1, err := client.CatalogService.ListAnnotationsForDatasetById(lds.Id, nil)
	require.NoError(t, err)
	require.True(t, len(ans1) > 0)

	ans2, err := client.CatalogService.ListAnnotationsForDatasetByResourceName(lds.Module+"."+lds.Name, nil)
	require.NoError(t, err)
	require.True(t, len(ans2) > 0)

	// Delete
	//Currently failing with a 404 - Annotation not found error, under investigation
	//err = client.CatalogService.DeleteAnnotationOfDatasetById(lds.Id, annotationId)
	//require.NoError(t, err)
	//err = client.CatalogService.DeleteAnnotationOfDatasetByResourceName(lds.Module+"."+lds.Name, annotationId)
	//require.NoError(t, err)

}
