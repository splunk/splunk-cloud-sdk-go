package service

import (
	"github.com/splunk/ssc-client-go/lib/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getDataset(t *testing.T) {

	result, err := getSplunkClient().CatalogService.GetDataset("ds1")
	assert.Empty(t, err)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, "ds1", result.Name)
	// TODO: uncomment this once the stubby server response is fixed
	//assert.Equal(t, model.VIEW, result.Kind)
}

func Test_getDatasets(t *testing.T) {

	result, err := getSplunkClient().CatalogService.GetDatasets()
	assert.Empty(t, err)
	assert.Equal(t, 2, len(result))
}

func Test_postDataset(t *testing.T) {
	result, err := getSplunkClient().CatalogService.PostDataset(
		getSplunkClient().CatalogService.CreateDataset("ds1", model.VIEW, []string{"string"}, "string"))
	assert.Empty(t, err)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, "ds1", result.Name)
	assert.Equal(t, model.VIEW, result.Kind)
	assert.Equal(t, []string{"string"}, result.Rules)
}

func Test_deleteDataset(t *testing.T) {
	err := getSplunkClient().CatalogService.DeleteDataset("ds1")
	assert.Empty(t, err)
}

// TODO No dfference between stubby tests and unit tests (Should one be removed?)
// Stubby test for DeleteRule() catalog service endpoint
func TestDeleteRule(t *testing.T) {
	response, err := getSplunkClient().CatalogService.DeleteRule("rule1")
	assert.Empty(t, err)
	assert.Equal(t, "200 OK", response.Status)
}

// Stubby test for GetRules() catalog service endpoint
func TestGetRules(t *testing.T) {
	result, err := getSplunkClient().CatalogService.GetRules()
	assert.Empty(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, 3, len(result[0].Actions))
}

// Stubby test for PostRule() catalog service endpoint
func TestPostRule(t *testing.T) {
	var actions[3]model.Action
	actions[0] = CreateAction("AUTOKV","" ,"",true, "NONE", "", "", "", 0, "")
	actions[1] = CreateAction("EVAL","" ,"", false, "", "string", "", "", 0, "string")
	actions[2] = CreateAction("LOOKUP","" ,"",false, "", "string", "", "", 0, "")
	result, err := getSplunkClient().CatalogService.PostRule(CreateRule("rule1","newrule",7,"first rule", actions[:]))
	assert.Empty(t, err)
	//assert.NotEmpty(t, result.Id)
	assert.Equal(t, "rule4", result.Name)
	assert.Equal(t, "newrule", result.Match)
	assert.Equal(t, 3, len(result.Actions))
}