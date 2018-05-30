// +build !integration

package playgroundintegration

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
	// "strings"
	"strings"
	"fmt"
)

func cleanupDatasets(t *testing.T) {
	client := getClient()
	result, err := client.CatalogService.GetDatasets()
	assert.Nil(t, err)

	for _, item := range result {
		err = client.CatalogService.DeleteDataset(item.ID)
		assert.Nil(t, err)
	}
}

func cleanupRules(t *testing.T) {
	client := getClient()
	result, err := client.CatalogService.GetRules()
	assert.Nil(t, err)

	for _, item := range result {
		err := client.CatalogService.DeleteRule(item.ID)
		assert.Nil(t, err)
	}
}

func TestIntegrationCRUDDatasets(t *testing.T) {
	// cleanupDatasets(t)

	client := getClient()
	invalidClient := getInvalidClient()

	// create dataset
	datasetName := "integ_dataset_1000"
	datasetOwner := "Splunk"
	datasetCapabilities := "1101-00000:11010"

	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: datasetName, Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	assert.Nil(t, err)
	assert.Equal(t, datasetName, dataset.Name)
	assert.Equal(t, model.LOOKUP, dataset.Kind)
	// assert.Equal(t, []string{"somerule"}, dataset.Rules)
	_, err = client.CatalogService.CreateDataset(
		model.DatasetInfo{Name: "integ_dataset_2000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})
	assert.Nil(t, err)
	_, err = client.CatalogService.CreateDataset(
		model.DatasetInfo{Name: "integ_dataset_3000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})
	assert.Nil(t, err)

	// testing CreateDataset for 409 DatasetInfo already present error
	_, err = client.CatalogService.CreateDataset(
		model.DatasetInfo{ID: dataset.ID, Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "409"))

	// testing CreateDataset for 401 Unauthorized operation error
	_, err = invalidClient.CatalogService.CreateDataset(
		model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "401 Unauthorized"))

	// testing CreateDataset for 400 Invalid DatasetInfo error
	_, err = client.CatalogService.CreateDataset(
		model.DatasetInfo{Name: "integ_dataset_4000", Kind: model.LOOKUP})
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "400"))

	// get datasets
	datasets, err := client.CatalogService.GetDatasets()
	assert.Nil(t, err)
	assert.NotNil(t, len(datasets))

	// testing GetDatasets for 401 Unauthorized operation error
	_, err = invalidClient.CatalogService.GetDatasets()
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "401 Unauthorized"))

	// get dataset
	datasetByID, err := client.CatalogService.GetDataset(dataset.ID)
	assert.Nil(t, err)

	// testing GetDataset for 401 Unauthorized operation error
	_, err = invalidClient.CatalogService.GetDataset(dataset.ID)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "401 Unauthorized"))

	// testing GetDataset for 404 DatasetInfo not found error
	_, err = client.CatalogService.GetDataset("123")
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "404"))

	// update an existing dataset
	updatedDataset, err := client.CatalogService.UpdateDataset(model.PartialDatasetInfo{Name: datasetName, Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName", Version: 6}, dataset.ID)
	assert.Nil(t, err)
	assert.NotNil(t, updatedDataset)

	// testing UpdateDataset for 404 DatasetInfo not found error
	_, err = client.CatalogService.UpdateDataset(model.PartialDatasetInfo{Name: "goSdkDataset6", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName", Version: 2}, "123")
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "404"))

	// delete dataset
	err = client.CatalogService.DeleteDataset(datasetByID.ID)
	assert.Nil(t, err)

	// testing DeleteDataset for 401 Unauthorized operation error
	err = invalidClient.CatalogService.DeleteDataset(dataset.ID)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "401 Unauthorized"))

	// testing DeleteDataset for 404 DatasetInfo not found error
	err = client.CatalogService.DeleteDataset("123")
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "404"))

	// todo (Parul): 405 DatasetInfo cannot be deleted because of dependencies error case

	// cleanupDatasets(t)
}

func TestIntegrationCRUDRules(t *testing.T) {
	//defer cleanupRules(t)

	client := getClient()
	invalidClient := getInvalidClient()

	// create rule
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

	// testing CreateRule for 409 Rule already present error
	_, err = client.CatalogService.CreateRule(model.Rule{ID: rule.ID, Name: ruleName, Module: ruleModule, Owner: owner})
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "409"))

	// testing CreateRule for 401 Unauthorized operation error
	_, err = invalidClient.CatalogService.CreateRule(model.Rule{Name: ruleName, Module: ruleModule, Owner: owner})
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "401 Unauthorized"))

	// testing CreateRule for 400 Invalid Rule error
	/*	_, err = client.CatalogService.CreateRule(model.Rule{Name: ruleName})
		assert.NotNil(t, err)
		assert.True(t, strings.Contains(err.Error(), "400 Invalid"))*/

	// get all the rules
	rules, err := client.CatalogService.GetRules()
	assert.Nil(t, err)
	assert.NotNil(t, len(rules))

	// testing GetRules for 401 Unauthorized operation error
	_, err = invalidClient.CatalogService.GetRules()
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "401 Unauthorized"))

	//get a rule by ID
	ruleByID, err := client.CatalogService.GetRule(rule.ID)
	assert.Nil(t, err)
	fmt.Println(ruleByID)

	// testing GetRules for 401 Unauthorized operation error
	_, err = invalidClient.CatalogService.GetRule(rule.ID)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "401 Unauthorized"))

	// testing GetRules for 404 Rule not found error
	_, err = client.CatalogService.GetRule("123")
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "404"))

	// delete a rule by ID
	err = client.CatalogService.DeleteRule(rule.ID)
	assert.Nil(t, err)

	// testing DeleteRule for 401 Unauthorized operation error
	err = invalidClient.CatalogService.DeleteRule(rule.ID)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "401 Unauthorized"))

	// testing DeleteRule for 404 Rule not found error
	err = client.CatalogService.DeleteRule("123")
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "404"))

	// todo (Parul): 405 Rule cannot be deleted because of dependencies error case

	cleanupRules(t)
}
