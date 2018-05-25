// +build !integration

package service

import (
	"testing"
	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
	// "strings"
	"fmt"
)

func cleanupDatasets(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	result, err := client.CatalogService.GetDatasets()
	assert.Nil(t, err)

	for _, item := range result {
		err = client.CatalogService.DeleteDataset(item.ID)
		assert.Nil(t, err)
	}
}

func cleanupRules(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	result, err := client.CatalogService.GetRules()
	assert.Nil(t, err)

	for _, item := range result {
		err := client.CatalogService.DeleteRule(item.ID)
		assert.Nil(t, err)
	}
}

func TestIntegrationCRUDDatasets(t *testing.T) {
	//defer cleanupDatasets(t)

	client := getSplunkClientForPlaygroundTests()
	invalidClient := getInvalidSplunkClientForPlaygroundTests()

	// create dataset
	datasetName := "integ_dataset_1"
	datasetOwner := "Splunk"
	datasetCapabilities := "1101-00000:11010"
	isDatasetDisabled := false
	dataset, err := client.CatalogService.CreateDataset(
		model.DatasetInfo{Name: datasetName, Kind: model.INDEX, Owner: datasetOwner, Capabilities: datasetCapabilities, Disabled: isDatasetDisabled})

	assert.Nil(t, err)
	assert.Equal(t, datasetName, dataset.Name)
	assert.Equal(t, model.INDEX, dataset.Kind)
	// assert.Equal(t, []string{"somerule"}, dataset.Rules)
	//todo field is not returned from playground, bug in catalog service
	//assert.Equal(t, "todos", dataset.Todo)
	_, err = client.CatalogService.CreateDataset(
		model.DatasetInfo{Name: "integ_dataset_2", Kind: model.INDEX, Owner: datasetOwner, Capabilities: datasetCapabilities, Disabled: false})
	assert.Nil(t, err)
	_, err = client.CatalogService.CreateDataset(
		model.DatasetInfo{Name: "integ_dataset_3", Kind: model.INDEX, Owner: datasetOwner, Capabilities: datasetCapabilities, Disabled: false})
	assert.Nil(t, err)

	_, err = client.CatalogService.CreateDataset(
		model.DatasetInfo{Name: "goSdkDataset2", Kind: model.INDEX, Owner: datasetOwner, Capabilities: datasetCapabilities, Disabled: isDatasetDisabled})
	assert.NotNil(t, err)

	_, err = invalidClient.CatalogService.CreateDataset(
		model.DatasetInfo{Name: "goSdkDataset2", Kind: model.INDEX, Owner: datasetOwner, Capabilities: datasetCapabilities, Disabled: isDatasetDisabled})
	assert.NotNil(t, err)

	_, err = client.CatalogService.CreateDataset(
		model.DatasetInfo{Name: "goSdkDataset4", Kind: model.INDEX})
	assert.NotNil(t, err)

	//get datasets
	datasets, err := client.CatalogService.GetDatasets()
	assert.Nil(t, err)
	assert.NotNil(t, len(datasets))

	//get dataset
	datasetByID, getDatasetErr := client.CatalogService.GetDataset("5b0704135ef8cf000a2cd55a")
	assert.Nil(t, getDatasetErr)
	fmt.Println(datasetByID)

	updatedDataset, updateDatasetErr := client.CatalogService.UpdateDataset(model.PartialDatasetInfo{Name: "goSdkDataset6", Kind: model.INDEX, Owner: datasetOwner, Capabilities: datasetCapabilities, Disabled: isDatasetDisabled, Version: 6}, "5b06fce15ef8cf000a2cd557")
	assert.Nil(t, updateDatasetErr)
	fmt.Println(updatedDataset)

	//delete dataset
	err = client.CatalogService.DeleteDataset(datasetByID.ID)

	//delete dataset
	// cleanupDatasets(t)
}

/*func TestIntegrationDatasetsErrors(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSplunkClientForPlaygroundTests()

	// create dataset
	datasetName := "goSdkDataset1"
	_, err := client.CatalogService.CreateDataset(
		model.Dataset{Name: datasetName, Kind: model.VIEW, Rules: []string{"somerule"}, Todo: "todos"})
	assert.Nil(t, err)

	// create duplicated dataset should return 409
	_, err = client.CatalogService.CreateDataset(
		model.Dataset{Name: datasetName, Kind: model.VIEW})
	assert.True(t, strings.Contains(err.Error(), "409 Conflict"))

	//delete dataset
	cleanupDatasets(t)
}*/

func TestIntegrationCRUDRules(t *testing.T) {
	cleanupRules(t)
	defer cleanupRules(t)

	client := getSplunkClientForPlaygroundTests()

	//create rule
	ruleName := "goSdkTestrRule1"
	ruleModule := "catalog"
	ruleMatch := "integration_test_match"
	owner := "splunk"
	rule, err := client.CatalogService.CreateRule(
		model.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	assert.Nil(t, err)
	assert.Equal(t, ruleName, rule.Name)
	assert.Equal(t, ruleMatch, rule.Match)

	_, err = client.CatalogService.CreateRule(
		model.Rule{Name: "anotherone", Module: ruleModule, Owner: owner})
	assert.Nil(t, err)

	_, err = client.CatalogService.CreateRule(
		model.Rule{Name: "thirdone", Module: ruleModule, Owner: owner})
	assert.Nil(t, err)

	//get rules
	rules, err := client.CatalogService.GetRules()
	assert.Nil(t, err)
	assert.NotNil(t, len(rules))

/*	ruleById, getRuleErr := client.CatalogService.GetRule(rule.ID)
	assert.Nil(t, getRuleErr)
	fmt.Println(ruleById)*/

	deleteRuleErr := client.CatalogService.DeleteRule(rule.ID)
	assert.Nil(t, deleteRuleErr)

	//delete rules
	 cleanupRules(t)
}

/*func TestIntegrationRulessErrors(t *testing.T) {
	defer cleanupRules(t)

	client := getSplunkClientForPlaygroundTests()

	//create rule
	ruleName := "goSdkTestrRule1"
	_, err := client.CatalogService.CreateRule(
		model.Rule{Name: ruleName, Priority: 8})
	assert.Nil(t, err)

	// create duplicated rule should return 409
	_, err = client.CatalogService.CreateRule(
		model.Rule{Name: ruleName, Priority: 8})
	assert.True(t, strings.Contains(err.Error(), "409 Conflict"))

	//delete rules
	cleanupRules(t)
}*/
