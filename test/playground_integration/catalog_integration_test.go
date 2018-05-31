package playgroundintegration

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/splunk/ssc-client-go/model"
)

func cleanupDatasets(t *testing.T) {
	client := getClient(t)
	result, err := client.CatalogService.GetDatasets()
	assert.Nil(t, err)

	for _, item := range result {
		err = client.CatalogService.DeleteDataset(item.Name)
		assert.Nil(t, err)
	}
}

func cleanupRules(t *testing.T) {
	client := getClient(t)
	result, err := client.CatalogService.GetRules()
	assert.Nil(t, err)

	for _, item := range result {
		err := client.CatalogService.DeleteRule(item.Name)
		assert.Nil(t, err)
	}
}

func TestIntegrationCRUDDatasets(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// create dataset
	datasetName := "goSdkDataset1"
	testDataset := model.Dataset{Name: datasetName, Kind: model.VIEW, Rules: []string{"somerule"}, Todo: "todos"}
	dataset, err := client.CatalogService.CreateDataset(testDataset)

	assert.Nil(t, err)
	assert.Equal(t, datasetName, dataset.Name)
	assert.Equal(t, model.VIEW, dataset.Kind)
	assert.Equal(t, []string{"somerule"}, dataset.Rules)
	//todo field is not returned from playground_integration, bug in catalog service
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
	defer cleanupDatasets(t)

	client := getClient(t)

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
	defer cleanupRules(t)

	client := getClient(t)

	//create rule
	ruleName := "goSdkTestrRule1"
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
	defer cleanupRules(t)

	client := getClient(t)

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
}
