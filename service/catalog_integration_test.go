// +build !integration

package service

import (
	"testing"
	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
	"strings"
)

func cleanupDatasets(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	result, err := client.CatalogService.GetDatasets()
	assert.Nil(t, err)

	for _, item := range result {
		err = client.CatalogService.DeleteDataset(item.Name)
		assert.Nil(t, err)
	}
}

func cleanupRules(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()
	result, err := client.CatalogService.GetRules()
	assert.Nil(t, err)

	for _, item := range result {
		err := client.CatalogService.DeleteRule(item.Name)
		assert.Nil(t, err)
	}
}

func TestIntegrationCRUDDatasets(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()

	cleanupDatasets(t)

	// create dataset
	datasetName := "goSdkDataset1"
	dataset, err := client.CatalogService.CreateDataset(
		model.Dataset{Name: datasetName, Kind: model.VIEW, Rules: []string{"somerule"}, Todo: "todos"})

	assert.Nil(t, err)
	assert.Equal(t, datasetName, dataset.Name)
	assert.Equal(t, model.VIEW, dataset.Kind)
	assert.Equal(t, []string{"somerule"}, dataset.Rules)
	//todo field is not returned from playground, bug in catalog service
	//assert.Equal(t, "todos", dataset.Todo)

	_, err = client.CatalogService.CreateDataset(
		model.Dataset{Name: "anotherone", Kind: model.VIEW, Rules: []string{"somerule"}, Todo: "todos"})
	assert.Nil(t, err)
	_, err = client.CatalogService.CreateDataset(
		model.Dataset{Name: "thirdone", Kind: model.VIEW, Rules: []string{"somerule"}, Todo: "todos"})
	assert.Nil(t, err)

	//get datasets
	datasets, err := client.CatalogService.GetDatasets()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(datasets))

	//delete dataset
	cleanupDatasets(t)
}

func TestIntegrationDatasetsErrors(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()

	cleanupDatasets(t)

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
}

func TestIntegrationCRUDRules(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()

	cleanupRules(t)

	//create rule
	ruleName := "goSdkTestRule1"
	rule, err := client.CatalogService.CreateRule(
		model.Rule{Name: ruleName, Priority: 8})
	assert.Nil(t, err)
	assert.Equal(t, ruleName, rule.Name)
	assert.Equal(t, 8, rule.Priority)

	_, err = client.CatalogService.CreateRule(
		model.Rule{Name: "anotherone"})
	assert.Nil(t, err)

	_, err = client.CatalogService.CreateRule(
		model.Rule{Name: "thirdone"})
	assert.Nil(t, err)

	//get rules
	rules, err := client.CatalogService.GetRules()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(rules))

	//delete rules
	cleanupRules(t)
}

func TestIntegrationRulesErrors(t *testing.T) {
	client := getSplunkClientForPlaygroundTests()

	cleanupRules(t)

	//create rule
	ruleName := "goSdkTestRule1"
	_, err := client.CatalogService.CreateRule(
		model.Rule{Name: ruleName, Priority: 8})
	assert.Nil(t, err)

	// create duplicated rule should return 409
	_, err = client.CatalogService.CreateRule(
		model.Rule{Name: ruleName, Priority: 8})
	assert.True(t, strings.Contains(err.Error(), "409 Conflict"))

	//delete rules
	cleanupRules(t)
}
