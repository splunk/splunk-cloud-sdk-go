// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"net/url"
	"testing"

	"fmt"

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
	dsNameTemplate = fmt.Sprintf("gointegds%s_%d", "%s", testutils.TimeSec)
	newname        = fmt.Sprintf("newmx%d", testutils.TimeSec)
	newmod         = fmt.Sprintf("newmod%d", testutils.TimeSec)
	newowner       = "test1@splunk.com"
	// Lookup:
	caseMatch    = true
	externalName = "test_externalName"
	filter       = `kind=="lookup"`
	// Metric/Index:
	disabled               = false
	frozenTimePeriodInSecs = 60
	newftime               = 999
	// View:
	searchString = "search index=main|stats count()"
)

// Test Rule variables
var (
	ruleNameTemplate = fmt.Sprintf("gointegrule%s_%d", "%s", testutils.TimeSec)
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
	err := client.CatalogService.DeleteDataset(id)
	assert.Emptyf(t, err, "Error deleting dataset: %s", err)
}

func cleanupRule(t *testing.T, id string) {
	client := getSdkClient(t)
	err := client.CatalogService.DeleteRule(id)
	assert.Emptyf(t, err, "Error deleting rule: %s", err)
}

func cleanupRuleAction(t *testing.T, ruleID, actionID string) {
	client := getSdkClient(t)
	err := client.CatalogService.DeleteRuleAction(ruleID, actionID)
	assert.Emptyf(t, err, "Error deleting rule action: %s", err)
}

// createLookupDataset - Helper function for creating a valid Lookup in Catalog
func createLookupDataset(t *testing.T, name string) (*catalog.LookupDataset, error) {
	externalKind := string(catalog.KvCollection)
	createLookup := &catalog.LookupDataset{
		Name:               name,
		Kind:               string(catalog.Lookup),
		Module:             testutils.TestModule,
		CaseSensitiveMatch: &caseMatch,
		ExternalKind:       &externalKind,
		ExternalName:       &externalName,
		Filter:             &filter,
	}
	return getSdkClient(t).CatalogService.CreateLookupDataset(createLookup)
}

// createKVCollectionDataset - Helper function for creating a valid KVCollection in Catalog
func createKVCollectionDataset(t *testing.T, name string) (*catalog.KVCollectionDataset, error) {
	createKVCollection := &catalog.CreateKVCollectionDataset{
		CreateDatasetBase: catalog.NewCreateDatasetBaseByName(name, catalog.KvCollection, testutils.TestModule),
	}
	return getSdkClient(t).CatalogService.CreateKVCollectionDataset(createKVCollection)
}

// createMetricDataset - Helper function for creating a valid Metric in Catalog
func createMetricDataset(t *testing.T, name string) (*catalog.MetricDataset, error) {
	createMetric := &catalog.CreateMetricDataset{
		CreateDatasetBase: catalog.NewCreateDatasetBaseByName(name, catalog.Metric, testutils.TestModule),
		MetricProperties:  catalog.NewMetricProperties(disabled, frozenTimePeriodInSecs),
	}
	return getSdkClient(t).CatalogService.CreateMetricDataset(createMetric)
}

// createIndexDataset - Helper function for creating a valid Index in Catalog
func createIndexDataset(t *testing.T, name string) (*catalog.IndexDataset, error) {
	createIndex := &catalog.CreateIndexDataset{
		CreateDatasetBase: catalog.NewCreateDatasetBaseByName(name, catalog.Index, testutils.TestModule),
		IndexProperties:   catalog.NewIndexProperties(disabled, frozenTimePeriodInSecs),
	}
	return getSdkClient(t).CatalogService.CreateIndexDataset(createIndex)
}

// createImportDatasetByID - Helper function for creating a valid Import in Catalog
func createImportDatasetByID(t *testing.T, name, importID string) (*catalog.ImportDataset, error) {
	createImport := &catalog.CreateImportDataset{
		CreateDatasetBase: catalog.NewCreateDatasetBaseByName(name, catalog.Import, testutils.TestModule),
		ImportProperties:  catalog.NewImportPropertiesByID(importID),
	}
	return getSdkClient(t).CatalogService.CreateImportDataset(createImport)
}

// createImportDatasetByName - Helper function for creating a valid Import in Catalog
func createImportDatasetByName(t *testing.T, name, importName, importModule string) (*catalog.ImportDataset, error) {
	createImport := &catalog.CreateImportDataset{
		CreateDatasetBase: catalog.NewCreateDatasetBaseByName(name, catalog.Import, testutils.TestModule),
		ImportProperties:  catalog.NewImportPropertiesByName(importModule, importName),
	}
	return getSdkClient(t).CatalogService.CreateImportDataset(createImport)
}

// createViewDataset - Helper function for creating a valid View in Catalog
func createViewDataset(t *testing.T, name string) (*catalog.ViewDataset, error) {
	createView := &catalog.CreateViewDataset{
		CreateDatasetBase: catalog.NewCreateDatasetBaseByName(name, catalog.View, testutils.TestModule),
		ViewProperties:    catalog.NewViewProperties(searchString),
	}
	return getSdkClient(t).CatalogService.CreateViewDataset(createView)
}

// createViewDataset - Helper function for creating Fields
func createDatasetField(datasetID string, client *sdk.Client, t *testing.T) *catalog.Field {
	testField := catalog.Field{Name: "integ_test_field", DataType: "S", FieldType: "D", Prevalence: "A"}
	resultField, err := client.CatalogService.CreateDatasetField(datasetID, &testField)
	require.Nil(t, err)
	require.NotEmpty(t, resultField)
	return resultField
}

// assertDatasetKind - Helper to assert that the kind for the Dataset matches model associated with that kind
func assertDatasetKind(t *testing.T, dataset catalog.Dataset) {
	switch dataset.GetKind() {
	case string(catalog.Index):
		ds, ok := dataset.(catalog.IndexDataset)
		assert.True(t, ok)
		assert.NotEmpty(t, ds.ID)
	case string(catalog.View):
		ds, ok := dataset.(catalog.ViewDataset)
		assert.True(t, ok)
		assert.NotEmpty(t, ds.ID)
	case string(catalog.Lookup):
		ds, ok := dataset.(catalog.LookupDataset)
		assert.True(t, ok)
		assert.NotEmpty(t, ds.ID)
	case string(catalog.Import):
		ds, ok := dataset.(catalog.ImportDataset)
		assert.True(t, ok)
		assert.NotEmpty(t, ds.ID)
	case string(catalog.Job):
		ds, ok := dataset.(catalog.JobDataset)
		assert.True(t, ok)
		assert.NotEmpty(t, ds.ID)
	case string(catalog.Metric):
		ds, ok := dataset.(catalog.MetricDataset)
		assert.True(t, ok)
		assert.NotEmpty(t, ds.ID)
	case string(catalog.KvCollection):
		ds, ok := dataset.(catalog.KVCollectionDataset)
		assert.True(t, ok)
		assert.NotEmpty(t, ds.ID)
	// These are known kinds but are not supported in the spec:
	case "catalog":
	case "splv1sink":
		ds, ok := dataset.(catalog.DatasetBase)
		assert.True(t, ok)
		assert.NotEmpty(t, ds.ID)
	// Anything here is not a known kind and should potentially be on our radar:
	default:
		ds, ok := dataset.(catalog.DatasetBase)
		assert.True(t, ok)
		assert.NotEmpty(t, ds.ID)
		fmt.Printf("WARNING: catalog dataset found with unknown kind, support may be missing for this kind: %s\n", ds.Kind)
	}
}

// Test CreateIndexDataset
func CreateIndexDataset(t *testing.T) {
	indexds, err := createIndexDataset(t, makeDSName("crix"))
	require.Nil(t, err)
	defer cleanupDataset(t, indexds.ID)
	require.NotNil(t, indexds)
	require.Equal(t, catalog.Index, indexds.GetKind())
}

// Test CreateImportDataset
func TestCreateImportDataset(t *testing.T) {
	indexds, err := createIndexDataset(t, makeDSName("crix"))
	require.Nil(t, err)
	defer cleanupDataset(t, indexds.ID)
	importds, err := createImportDatasetByID(t, makeDSName("crim"), indexds.ID)
	require.Nil(t, err)
	defer cleanupDataset(t, importds.ID)
	require.NotNil(t, importds)
	require.Equal(t, string(catalog.Import), importds.Kind)
}

// Test CreateKVCollectionDataset
func TestKVCollectionDataset(t *testing.T) {
	kvds, err := createKVCollectionDataset(t, makeDSName("crkv"))
	require.Nil(t, err)
	defer cleanupDataset(t, kvds.ID)
	require.NotNil(t, kvds)
	require.Equal(t, string(catalog.KvCollection), kvds.Kind)
}

// Test CreateLookupDataset
func TestLookupDataset(t *testing.T) {
	lookupds, err := createLookupDataset(t, makeDSName("crlk"))
	require.Nil(t, err)
	defer cleanupDataset(t, lookupds.ID)
	require.NotNil(t, lookupds)
	require.Equal(t, string(catalog.Lookup), lookupds.Kind)
}

// Test CreateMetricDataset
func TestMetricDataset(t *testing.T) {
	metricds, err := createMetricDataset(t, makeDSName("crmx"))
	require.Nil(t, err)
	defer cleanupDataset(t, metricds.ID)
	require.NotNil(t, metricds)
	require.Equal(t, string(catalog.Metric), metricds.Kind)
}

// Test CreateViewDataset
func TestViewDataset(t *testing.T) {
	viewds, err := createViewDataset(t, makeDSName("crvw"))
	require.Nil(t, err)
	defer cleanupDataset(t, viewds.ID)
	require.NotNil(t, viewds)
	require.Equal(t, string(catalog.View), viewds.Kind)
}

// Test CreateDataset for 409 DatasetInfo already present error
func TestCreateDatasetDataAlreadyPresentError(t *testing.T) {
	// create dataset
	ds, err := createLookupDataset(t, makeDSName("409"))
	require.Nil(t, err)
	defer cleanupDataset(t, ds.ID)
	_, err = createLookupDataset(t, makeDSName("409"))
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 409, httpErr.HTTPStatusCode)
}

// Test CreateDataset for 401 Unauthorized operation error
func TestCreateDatasetUnauthorizedOperationError(t *testing.T) {
	name := makeDSName("401")
	createView := &catalog.CreateViewDataset{
		CreateDatasetBase: catalog.NewCreateDatasetBaseByName(name, catalog.View, testutils.TestModule),
		ViewProperties:    catalog.NewViewProperties(searchString),
	}
	ds, err := getInvalidClient(t).CatalogService.CreateViewDataset(createView)
	if ds != nil {
		defer cleanupDataset(t, ds.ID)
	}
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 401, httpErr.HTTPStatusCode)
	assert.Equal(t, "Error validating request", httpErr.Message)
}

// Test CreateDataset for 400 Invalid DatasetInfo error
func TestCreateDatasetInvalidDatasetInfoError(t *testing.T) {
	ds, err := getSdkClient(t).CatalogService.CreateDataset(&catalog.DatasetBase{Name: makeDSName("400"), Kind: "lookup", CreatedBy: "thisisnotvalid"})
	if ds != nil {
		dsb, ok := ds.(catalog.DatasetBase)
		require.True(t, ok)
		defer cleanupDataset(t, dsb.ID)
	}
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 400, httpErr.HTTPStatusCode)
}

// Test ListDatasets
func TestListAllDatasets(t *testing.T) {
	ds, err := createLookupDataset(t, makeDSName("getall"))
	require.Nil(t, err)
	defer cleanupDataset(t, ds.ID)

	datasets, err := getSdkClient(t).CatalogService.ListDatasets(nil)
	require.Nil(t, err)
	assert.NotZero(t, len(datasets))

	// We should be able assert that the kinds are known and relate to their associated model
	for i := 0; i < len(datasets); i++ {
		assertDatasetKind(t, datasets[i])
	}
}

// Test TestListDatasetsFilter
func TestListDatasetsFilter(t *testing.T) {
	ds, err := createLookupDataset(t, makeDSName("fil"))
	require.Nil(t, err)
	defer cleanupDataset(t, ds.ID)

	values := make(url.Values)
	values.Set("filter", filter) // kind=="lookup"

	datasets, err := getSdkClient(t).CatalogService.ListDatasets(values)
	require.Nil(t, err)
	assert.NotZero(t, len(datasets))
}

// Test TestListDatasetsCount
func TestListDatasetsCount(t *testing.T) {
	// Create three datasets
	ds1, err := createLookupDataset(t, makeDSName("cnt1"))
	require.Nil(t, err)
	defer cleanupDataset(t, ds1.ID)
	ds2, err := createLookupDataset(t, makeDSName("cnt2"))
	require.Nil(t, err)
	defer cleanupDataset(t, ds2.ID)
	ds3, err := createLookupDataset(t, makeDSName("cnt3"))
	require.Nil(t, err)
	defer cleanupDataset(t, ds3.ID)

	values := make(url.Values)
	values.Set("count", "3")

	// There should be at least three
	datasets, err := getSdkClient(t).CatalogService.ListDatasets(values)
	require.Nil(t, err)
	require.Equal(t, 3, len(datasets))
}

// Test TestListDatasetsOrderBy
func TestListDatasetsOrderBy(t *testing.T) {
	ds, err := createViewDataset(t, makeDSName("orby"))
	require.Nil(t, err)
	defer cleanupDataset(t, ds.ID)

	values := make(url.Values)
	values.Set("orderby", "id Ascending")

	datasets, err := getSdkClient(t).CatalogService.ListDatasets(values)
	assert.Nil(t, err)
	assert.NotZero(t, len(datasets))

	// We should be able assert that the kinds are known and relate to their associated model
	for i := 0; i < len(datasets); i++ {
		assertDatasetKind(t, datasets[i])
	}
}

// Test TestListDatasetsAll with filter, count, and orderby
func TestListDatasetsAll(t *testing.T) {
	ds, err := createViewDataset(t, makeDSName("fco"))
	require.Nil(t, err)
	defer cleanupDataset(t, ds.ID)

	values := make(url.Values)
	values.Set("filter", `kind=="kvcollection"`)
	values.Set("count", "1")
	values.Set("orderby", "id Descending")

	datasets, err := getSdkClient(t).CatalogService.ListDatasets(values)
	assert.Nil(t, err)
	assert.NotZero(t, len(datasets))
}

// Test GetDataset by ID
func TestGetDatasetByID(t *testing.T) {
	ds, err := createLookupDataset(t, makeDSName("cnt"))
	require.Nil(t, err)
	defer cleanupDataset(t, ds.ID)

	datasetByID, err := getSdkClient(t).CatalogService.GetDataset(ds.ID)

	require.Nil(t, err)
	assert.NotNil(t, datasetByID)
}

// Test GetDataset for 404 DatasetInfo not found error
func TestGetDatasetByIDDatasetNotFoundError(t *testing.T) {
	_, err := getSdkClient(t).CatalogService.GetDataset("idonotexist")
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
}

// Test UpdateIndexDataset
func TestUpdateIndexDataset(t *testing.T) {
	client := getSdkClient(t)

	indexds, err := createIndexDataset(t, makeDSName("uix"))
	require.Nil(t, err)
	defer cleanupDataset(t, indexds.ID)
	require.NotNil(t, indexds)
	uidx := &catalog.UpdateIndexDataset{
		IndexProperties: catalog.NewIndexProperties(!disabled, newftime),
	}
	newindexds, err := client.CatalogService.UpdateIndexDataset(uidx, indexds.ID)
	require.Nil(t, err)
	assert.Equal(t, !disabled, *newindexds.Disabled)
	assert.Equal(t, newftime, *newindexds.FrozenTimePeriodInSecs)
}

// Test UpdateMetricDataset
func TestUpdateMetricDataset(t *testing.T) {
	client := getSdkClient(t)

	metricds, err := createMetricDataset(t, makeDSName("umx"))
	require.Nil(t, err)
	defer cleanupDataset(t, metricds.ID)
	require.NotNil(t, metricds)
	// Update the metrics dataset, including name and module
	umx := &catalog.UpdateMetricDataset{
		UpdateDatasetBase: catalog.NewUpdateDatasetBase(newname, newmod, newowner),
		MetricProperties:  catalog.NewMetricProperties(!disabled, newftime),
	}
	newmetricsds, err := client.CatalogService.UpdateMetricDataset(umx, metricds.ID)
	require.Nil(t, err)
	assert.Equal(t, newname, newmetricsds.Name)
	assert.Equal(t, newmod, newmetricsds.Module)
	assert.Equal(t, newowner, newmetricsds.Owner)
	assert.Equal(t, !disabled, *newmetricsds.Disabled)
	assert.Equal(t, newftime, *newmetricsds.FrozenTimePeriodInSecs)
}

// NOTE: UpdateImportDataset is not supported at this time
// Test UpdateImportDataset
// func TestUpdateImportDataset(t *testing.T) {
// 	// TODO: create newmetricsds and newimportds
// 	client := getSdkClient(t)
// 	importds, err := createImportDatasetByName(t, makeDSName("uim"), newmetricsds.Name, newmetricsds.Module)
// 	require.Nil(t, err)
// 	defer cleanupDataset(t, importds.ID)
// 	require.NotNil(t, importds)
// 	uim := &catalog.UpdateImportDataset{
// 		ImportProperties: &catalog.ImportProperties{
// 			SourceName:   &newindexds.Name,
// 			SourceModule: &newindexds.Module,
// 		},
// 	}
// 	newimportds, err := client.CatalogService.UpdateImportDataset(uim, importds.ID)
// 	require.Nil(t, err)
// 	assert.Equal(t, newindexds.Name, *newimportds.SourceName)
// 	assert.Equal(t, newindexds.Module, *newimportds.SourceModule)
// }

// Test UpdateJobDataset
func TestUpdateJobDataset(t *testing.T) {
	client := getSdkClient(t)

	// Create a search job to ensure at least one job exists
	searchjobReq := &search.CreateJobRequest{Query: "| from index:main | head 1"}
	searchjob, err := client.SearchService.CreateJob(searchjobReq)
	require.Nil(t, err)
	values := make(url.Values)
	values.Set("filter", fmt.Sprintf(`sid=="%s"`, searchjob.ID))
	values.Set("count", "1")
	datasets, err := getSdkClient(t).CatalogService.ListDatasets(values)
	require.Nil(t, err)
	require.NotZero(t, len(datasets))
	fmt.Printf("%+v", datasets)
	parsedds, err := catalog.ParseRawDataset(datasets[0])
	require.Nil(t, err)
	jobds, ok := parsedds.(catalog.JobDataset)
	require.True(t, ok)
	newstatus := string(search.JobCanceled)
	// This job should not be canceled since it was just created
	require.NotEqual(t, newstatus, *jobds.Status)
	ujb := &catalog.UpdateJobDataset{
		Status: &newstatus,
	}
	newjobds, err := client.CatalogService.UpdateJobDataset(ujb, jobds.ID)
	require.Nil(t, err)
	assert.Equal(t, newstatus, *newjobds.Status)
}

// Test UpdateLookupDataset
func TestUpdateLookupDataset(t *testing.T) {
	client := getSdkClient(t)

	lookupds, err := createLookupDataset(t, makeDSName("ulk"))
	require.Nil(t, err)
	defer cleanupDataset(t, lookupds.ID)
	require.NotNil(t, lookupds)
	notcasematch := !caseMatch
	newxname := "newxname"
	ulk := &catalog.LookupDataset{
		ID:                 "cantchangethis",
		Kind:               "cantchangethat",
		Owner:              "test1@splunk.com",
		CaseSensitiveMatch: &notcasematch,
		ExternalKind:       string(catalog.KvCollection),
		ExternalName:       &newxname,
		Filter:             `kind=="lookup"`,
	}
	newlookupds, err := client.CatalogService.UpdateLookupDataset(ulk, lookupds.ID)
	require.Nil(t, err)
	assert.NotEqual(t, "cantchangethis", *newlookupds.ID)
	assert.NotEqual(t, "cantchangethat", *newlookupds.Kind)
	assert.Equal(t, "test1@splunk.com", *newlookupds.Owner)
	assert.Equal(t, !caseMatch, *newlookupds.CaseSensitiveMatch)
	assert.Equal(t, newxname, *newlookupds.ExternalName)
	assert.Equal(t, newfilter, *newlookupds.Filter)
}

// Test UpdateViewDataset
func TestUpdateViewDataset(t *testing.T) {
	client := getSdkClient(t)

	viewds, err := createViewDataset(t, makeDSName("uvw"))
	require.Nil(t, err)
	defer cleanupDataset(t, viewds.ID)
	require.NotNil(t, viewds)
	newname = fmt.Sprintf("newvw%d", testutils.TimeSec)
	uvw := &catalog.UpdateViewDataset{
		UpdateDatasetBase: catalog.NewUpdateDatasetBase(newname, newmod, newowner),
	}
	newviewds, err := client.CatalogService.UpdateViewDataset(uvw, viewds.ID)
	require.Nil(t, err)
	assert.Equal(t, newname, newviewds.Name)
	assert.Equal(t, newmod, newviewds.Module)
	assert.Equal(t, newowner, newviewds.Owner)
}

// Test UpdateDataset for 404 Datasetnot found error
func TestUpdateExistingDatasetDataNotFoundError(t *testing.T) {
	uvw := &catalog.UpdateViewDataset{
		ViewProperties: &catalog.ViewProperties{
			Search: &searchString,
		},
	}
	_, err := getSdkClient(t).CatalogService.UpdateViewDataset(uvw, "idonotexist")
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
}

// Test DeleteDataset
func TestDeleteDataset(t *testing.T) {
	client := getSdkClient(t)

	ds, err := createViewDataset(t, makeDSName("delv"))
	require.Nil(t, err)

	err = client.CatalogService.DeleteDataset(ds.ID)
	require.Nil(t, err)

	_, err = client.CatalogService.GetDataset(ds.ID)
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

// todo (Parul): 405 DatasetInfo cannot be deleted because of dependencies error case

// Test CreateRules
func TestCreateRules(t *testing.T) {
	client := getSdkClient(t)

	// create rule
	ruleName := makeRuleName("crmat")
	rule, err := client.CatalogService.CreateRule(catalog.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch})
	require.Nil(t, err)
	defer cleanupRule(t, rule.ID)
	assert.Equal(t, ruleName, rule.Name)
	assert.Equal(t, ruleMatch, rule.Match)
}

// Test CreateRule for 409 Rule already present error
func TestCreateRuleDataAlreadyPresent(t *testing.T) {
	client := getSdkClient(t)

	// create rule
	ruleName := makeRuleName("409")
	rule, err := client.CatalogService.CreateRule(catalog.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch})
	require.Nil(t, err)
	defer cleanupRule(t, rule.ID)
	assert.Equal(t, ruleName, rule.Name)
	assert.Equal(t, ruleMatch, rule.Match)

	_, err = client.CatalogService.CreateRule(catalog.Rule{ID: rule.ID, Name: ruleName, Module: ruleModule, Match: ruleMatch})
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 409, httpErr.HTTPStatusCode)
}

// Test CreateRule for 401 Unauthorized operation error
func TestCreateRuleUnauthorizedOperationError(t *testing.T) {
	// create rule
	ruleName := makeRuleName("401")
	rule, err := getInvalidClient(t).CatalogService.CreateRule(catalog.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch})
	if rule != nil {
		defer cleanupRule(t, rule.ID)
	}
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 401, httpErr.HTTPStatusCode)
	assert.Equal(t, "Error validating request", httpErr.Message)
}

// Test GetRules
func TestGetAllRules(t *testing.T) {
	client := getSdkClient(t)

	// create rule
	ruleName := makeRuleName("getall")
	rule, err := client.CatalogService.CreateRule(catalog.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch})
	require.Nil(t, err)
	defer cleanupRule(t, rule.ID)

	rules, err := client.CatalogService.GetRules()
	require.Nil(t, err)
	assert.NotZero(t, len(rules))
}

// Test GetRule By ID
func TestGetRuleByID(t *testing.T) {
	client := getSdkClient(t)

	// create rule
	ruleName := makeRuleName("getid")
	rule, err := client.CatalogService.CreateRule(catalog.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch})
	require.Nil(t, err)
	require.NotEmpty(t, rule.ID)

	ruleByID, err := client.CatalogService.GetRule(rule.ID)
	assert.Nil(t, err)
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
func TestDeleteRule(t *testing.T) {
	client := getSdkClient(t)

	// create rule
	ruleName := makeRuleName("delid")
	rule, err := client.CatalogService.CreateRule(catalog.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch})
	require.Nil(t, err)
	require.NotEmpty(t, rule.ID)

	err = client.CatalogService.DeleteRule(rule.ID)
	assert.Nil(t, err)
}

// Test DeleteRule for 404 Rule not found error
func TestDeleteRulebyIDRuleNotFoundError(t *testing.T) {
	err := getSdkClient(t).CatalogService.DeleteRule("idonotexist")
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
}

// Test GetDatasetField
func TestGetDatasetFields(t *testing.T) {
	client := getSdkClient(t)

	// Create dataset
	dsname := makeDSName("getfs")
	ds, err := createLookupDataset(t, dsname)
	require.Nil(t, err)
	defer cleanupDataset(t, ds.ID)

	// create new fields in the dataset
	testField1 := catalog.Field{Name: "integ_test_field1", DataType: "S", FieldType: "D", Prevalence: "A"}
	testField2 := catalog.Field{Name: "integ_test_field2", DataType: "N", FieldType: "U", Prevalence: "S"}
	f1, err := client.CatalogService.CreateDatasetField(ds.ID, &testField1)
	f2, err := client.CatalogService.CreateDatasetField(ds.ID, &testField2)

	// Validate the creation of new dataset fields
	results, err := client.CatalogService.GetDatasetFields(ds.ID, nil)
	require.Nil(t, err)
	assert.True(t, len(results) >= 2)
	res1, err := client.CatalogService.GetDatasetField(ds.ID, f1.ID)
	require.Nil(t, err)
	assert.NotNil(t, res1)
	res2, err := client.CatalogService.GetDatasetField(ds.ID, f2.ID)
	require.Nil(t, err)
	assert.NotNil(t, res2)
}

// Test GetDatasetFields based on filter
func TestGetDatasetFieldsOnFilter(t *testing.T) {
	client := getSdkClient(t)

	// Create dataset
	dsname := makeDSName("getffil")
	ds, err := createLookupDataset(t, dsname)
	require.Nil(t, err)
	defer cleanupDataset(t, ds.ID)

	// create new fields in the dataset
	testField1 := catalog.Field{Name: "integ_test_field1", DataType: "S", FieldType: "D", Prevalence: "A"}
	testField2 := catalog.Field{Name: "integ_test_field2", DataType: "N", FieldType: "U", Prevalence: "S"}
	_, err = client.CatalogService.CreateDatasetField(ds.ID, &testField1)
	_, err = client.CatalogService.CreateDatasetField(ds.ID, &testField2)

	filter := make(url.Values)
	filter.Add("filter", `name=="integ_test_field1"`)

	// Validate the filter returned one result (testField1, not testField2)
	result, err := client.CatalogService.GetDatasetFields(ds.ID, filter)
	require.Nil(t, err)
	assert.NotEmpty(t, result)
	assert.Emptyf(t, err, "Error retrieving dataset fields: %s", err)
	assert.Equal(t, 1, len(result))
}

// Test CreateDatasetField
func TestCreateDatasetField(t *testing.T) {
	client := getSdkClient(t)

	// Create dataset
	dsname := makeDSName("crf")
	ds, err := createLookupDataset(t, dsname)
	require.Nil(t, err)
	defer cleanupDataset(t, ds.ID)

	// Create a new dataset field
	testField := catalog.Field{Name: "integ_test_field", DataType: "S", FieldType: "D", Prevalence: "A"}
	resultField, err := client.CatalogService.CreateDatasetField(ds.ID, &testField)
	require.Nil(t, err)
	require.NotEmpty(t, resultField)
	assert.Equal(t, "integ_test_field", resultField.Name)
	// TODO: catalog.String, Dimension, and All do not match "S", "D", "A" - why is this?
	assert.Equal(t, catalog.String, resultField.DataType)
	assert.Equal(t, catalog.Dimension, resultField.FieldType)
	assert.Equal(t, catalog.All, resultField.Prevalence)
	assert.Emptyf(t, err, "Error creating dataset field: %s", err)

	// Validate the creation of a new dataset field
	resultField, err = client.CatalogService.GetDatasetField(ds.ID, resultField.ID)
	require.Nil(t, err)
	assert.NotEmpty(t, resultField)
}

// Test PatchDatasetField
func TestPatchDatasetField(t *testing.T) {
	client := getSdkClient(t)

	// Create dataset
	dsname := makeDSName("patf")
	ds, err := createLookupDataset(t, dsname)
	require.Nil(t, err)
	defer cleanupDataset(t, ds.ID)

	// Create a new dataset field
	resultField := createDatasetField(ds.ID, client, t)

	// Update the existing dataset field
	resultField, err = client.CatalogService.UpdateDatasetField(ds.ID, resultField.ID, &catalog.Field{DataType: "O"})
	require.Nil(t, err)
	require.NotNil(t, resultField)
	assert.Equal(t, "integ_test_field", resultField.Name)
	assert.Equal(t, catalog.ObjectID, resultField.DataType)
	assert.Equal(t, catalog.Dimension, resultField.FieldType)
	assert.Equal(t, catalog.All, resultField.Prevalence)
}

// Test DeleteDatasetField
func TestDeleteDatasetField(t *testing.T) {
	client := getSdkClient(t)

	// Create dataset
	dsname := makeDSName("delf")
	ds, err := createLookupDataset(t, dsname)
	require.Nil(t, err)
	defer cleanupDataset(t, ds.ID)

	// Create a new dataset field
	resultField := createDatasetField(ds.ID, client, t)

	// Delete dataset field
	err = client.CatalogService.DeleteDatasetField(ds.ID, resultField.ID)
	require.Nil(t, err)

	// Validate the deletion of the dataset field
	_, err = client.CatalogService.GetDatasetField(ds.ID, resultField.ID)
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
	ds, err := createLookupDataset(t, dsname)
	require.Nil(t, err)
	defer cleanupDataset(t, ds.ID)

	// Create a new dataset field
	testField := catalog.Field{Name: "integ_test_field", DataType: "N", FieldType: "U", Prevalence: "S"}
	_, err = invalidClient.CatalogService.CreateDatasetField(ds.ID, &testField)
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
	ds, err := createLookupDataset(t, dsname)
	require.Nil(t, err)
	defer cleanupDataset(t, ds.ID)

	// Create a new dataset field
	createDatasetField(ds.ID, client, t)

	// Post an already created dataset field
	duplicateTestField := catalog.Field{Name: "integ_test_field", DataType: "S", FieldType: "D", Prevalence: "A"}
	_, err = client.CatalogService.CreateDatasetField(ds.ID, &duplicateTestField)
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 409, httpErr.HTTPStatusCode)
}

// Test CreateDatasetField for 400 error
func TestCreateDatasetFieldInvalidDataFormat(t *testing.T) {
	client := getSdkClient(t)

	// Create dataset
	dsname := makeDSName("f400")
	ds, err := createLookupDataset(t, dsname)
	require.Nil(t, err)
	defer cleanupDataset(t, ds.ID)

	// Create a new dataset field but with no data
	testField := catalog.Field{}
	_, err = client.CatalogService.CreateDatasetField(ds.ID, &testField)
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 400, httpErr.HTTPStatusCode)
}

// Test PatchDatasetField for 404 error
func TestPatchDatasetFieldDataNotFound(t *testing.T) {
	client := getSdkClient(t)

	// Create dataset
	dsname := makeDSName("fp404")
	ds, err := createLookupDataset(t, dsname)
	require.Nil(t, err)
	defer cleanupDataset(t, ds.ID)

	// Update non-existent dataset field
	_, err = client.CatalogService.UpdateDatasetField(ds.ID, "idonotexist", &catalog.Field{DataType: "O"})
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
	require.Nil(t, err)
	defer cleanupDataset(t, ds.ID)

	// Delete dataset field
	err = client.CatalogService.DeleteDatasetField(ds.ID, "idonotexist")
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
	require.Nil(t, err)
	defer cleanupDataset(t, ds.ID)

	// create new field in the dataset
	field := createDatasetField(ds.ID, client, t)

	// Create rule and rule action
	ruleName := makeRuleName("acts")
	rule, err := client.CatalogService.CreateRule(catalog.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch})
	require.Nil(t, err)
	defer cleanupRule(t, rule.ID)

	// Create rule action
	action1, err := client.CatalogService.CreateRuleAction(rule.ID, catalog.NewAliasAction(field.Name, "myfieldalias", ""))
	require.Nil(t, err)
	defer cleanupRuleAction(t, rule.ID, action1.ID)

	//update rule action
	tmpstr := "newaliasi"
	updateact, err := client.CatalogService.UpdateRuleAction(rule.ID, action1.ID, catalog.NewUpdateAliasAction(nil, &tmpstr))
	require.NotNil(t, updateact)
	assert.Equal(t, tmpstr, updateact.Alias)

	action2, err := client.CatalogService.CreateRuleAction(rule.ID, catalog.NewAutoKVAction("auto", "owner1"))
	require.Nil(t, err)
	defer cleanupRuleAction(t, rule.ID, action2.ID)

	//update rule action
	tmpstr = "auto"
	updateact, err = client.CatalogService.UpdateRuleAction(rule.ID, action2.ID, catalog.NewUpdateAutoKVAction(&tmpstr))
	require.NotNil(t, updateact)
	assert.Equal(t, tmpstr, updateact.Mode)

	action3, err := client.CatalogService.CreateRuleAction(rule.ID, catalog.NewEvalAction(field.Name, "now()", ""))
	require.Nil(t, err)
	defer cleanupRuleAction(t, rule.ID, action3.ID)

	//update rule action
	tmpstr = "newField"
	updateact, err = client.CatalogService.UpdateRuleAction(rule.ID, action3.ID, catalog.NewUpdateEvalAction(&tmpstr, nil))
	require.NotNil(t, updateact)
	assert.Equal(t, tmpstr, updateact.Field)

	action4, err := client.CatalogService.CreateRuleAction(rule.ID, catalog.NewLookupAction("myexpression2", ""))
	require.Nil(t, err)
	defer cleanupRuleAction(t, rule.ID, action4.ID)

	//update rule action
	tmpstr = "newexpr"
	updateact, err = client.CatalogService.UpdateRuleAction(rule.ID, action4.ID, catalog.NewUpdateLookupAction(&tmpstr))
	require.NotNil(t, updateact)
	assert.Equal(t, tmpstr, updateact.Expression)

	limit := 5
	action5, err := client.CatalogService.CreateRuleAction(rule.ID, catalog.NewRegexAction(field.Name, `field=myfield "From: (?<from>.*) To: (?<to>.*)"`, &limit, ""))
	require.Nil(t, err)
	assert.Equal(t, 5, *action5.Limit)
	defer cleanupRuleAction(t, rule.ID, action5.ID)

	action6, err := client.CatalogService.CreateRuleAction(rule.ID, catalog.NewRegexAction(field.Name, `field=myfield "From: (?<from>.*) To: (?<to>.*)"`, nil, ""))
	require.Nil(t, err)
	assert.Equal(t, (*int)(nil), action6.Limit)
	defer cleanupRuleAction(t, rule.ID, action6.ID)

	//update rule action
	tmpstr = `field=myotherfield "To: (?<to>.*) From: (?<from>.*)"`
	limit = 9
	updateact, err = client.CatalogService.UpdateRuleAction(rule.ID, action6.ID, catalog.NewUpdateRegexAction(nil, &tmpstr, &limit))
	require.NotNil(t, updateact)
	assert.Equal(t, tmpstr, updateact.Pattern)
	assert.Equal(t, limit, *updateact.Limit)

	//Get rule actions
	actions, err := client.CatalogService.GetRuleActions(rule.ID)
	require.NotNil(t, actions)
	assert.Equal(t, 6, len(actions))

	action7, err := client.CatalogService.GetRuleAction(rule.ID, actions[0].ID)
	require.NotNil(t, action7)

	// Delete action
	action8, err := client.CatalogService.CreateRuleAction(rule.ID, catalog.NewRegexAction(field.Name, `field=myfield "From: (?<from>.*) To: (?<to>.*)"`, &limit, ""))
	require.Nil(t, err)
	err = client.CatalogService.DeleteRuleAction(rule.ID, action8.ID)
	require.Nil(t, err)
}

// Test list modules
func TestGetModules(t *testing.T) {
	client := getSdkClient(t)

	// test using NO filter
	modules, err := client.CatalogService.GetModules(nil)
	require.Nil(t, err)
	assert.True(t, len(modules) > 0)

	// test using filter
	filter := make(url.Values)
	filter.Add("filter", `module==""`)
	modules, err = client.CatalogService.GetModules(filter)
	require.Nil(t, err)
	assert.Equal(t, 1, len(modules))
	assert.Equal(t, "", modules[0].Name)
}

/*
/ Currently unable to generate a bad rule
func TestCreateRuleInvalidRuleError(t *testing.T)  {
	defer cleanupRules(t)

	client := getSdkClient()

	// testing CreateRule for 400 Invalid Rule error
	ruleName := "goSdkTestrRule1"
	_, err := client.CatalogService.CreateRule(catalog.Rule{Name: ruleName})
	assert.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
  assert.True(t, httpErr.Status == 400, "Expected error code 400")
}*/

// todo (Parul): 405 Rule cannot be deleted because of dependencies error case
