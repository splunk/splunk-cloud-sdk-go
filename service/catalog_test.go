package service

import (
	"testing"
	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
)

// Stubby test for GetDataset() catalog service endpoint
func TestGetDataset(t *testing.T) {
	result, err := getSplunkClient().CatalogService.GetDataset("ds1")
	assert.Empty(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, "ds1", result.ID)
	assert.Equal(t, model.INDEX, result.Kind)
}

// Stubby test for GetDatasets() catalog service endpoint
func TestGetDatasets(t *testing.T) {
	result, err := getSplunkClient().CatalogService.GetDatasets()
	assert.Empty(t, err)
	assert.Equal(t, 2, len(result))
}

// Stubby test for CreateDataset() catalog service endpoint
func TestPostDataset(t *testing.T) {
	result, err := getSplunkClient().CatalogService.CreateDataset(
		model.DatasetInfo{Name: "stubby_dataset_1", Kind: model.INDEX, Owner: "Splunk", Capabilities: "1101-00000:11010", Disabled: true})
	assert.Empty(t, err)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, "stubby_dataset_1", result.Name)
	assert.Equal(t, model.INDEX, result.Kind)
}

// TODO: UpdateDataset

// Stubby test for DeleteDataset() catalog service endpoint
func TestDeleteDataset(t *testing.T) {
	err := getSplunkClient().CatalogService.DeleteDataset("ds1")
	assert.Empty(t, err)
}

// Stubby test for DeleteRule() catalog service endpoint
func TestDeleteRule(t *testing.T) {
	err := getSplunkClient().CatalogService.DeleteRule("rule1")
	assert.Empty(t, err)
}

// Stubby test for GetRules() catalog service endpoint
func TestGetRules(t *testing.T) {
	result, err := getSplunkClient().CatalogService.GetRules()
	assert.Empty(t, err)
	assert.Equal(t, 2, len(result))
}

func TestGetRule(t *testing.T) {
	result, err := getSplunkClient().CatalogService.GetRule("rule1")
	assert.Empty(t, err)
	assert.NotNil(t, "rule1",result.ID)
	assert.Equal(t, "_internal", result.Name)
}

// Stubby test for CreateRule() catalog service endpoint
func TestPostRule(t *testing.T) {
	var actions [3]model.Action
	actions[0] = CreateAction("AUTOKV", "Splunk", 0,"", model.NONE, "", "", "", 0)
	actions[1] = CreateAction("EVAL", "Splunk",0,"Splunk", "", "string", "", "",0)
	actions[2] = CreateAction("LOOKUP", "Splunk", 0, "", "", "string", "", "", 0)
	result, err := getSplunkClient().CatalogService.CreateRule(CreateRule("_internal", "test_match", "splunk", "Splunk", actions[:]))
	assert.Empty(t, err)
	assert.Equal(t, "_internal", result.Name)
	assert.Equal(t, "test_match", result.Match)
	assert.Equal(t, 3, len(result.Actions))
}

// creates a rule to post
func CreateRule(name string, match string, module string, owner string, actions []model.Action) model.Rule {
	return model.Rule{
		Name:    name,
		Match:   match,
		Module:  module,
		Owner:   owner,
		Actions: actions,
	}
}

// creates an action for rule to post
func CreateAction(kind model.ActionKind, owner string, version int, field string, mode model.AutoMode, expression string, pattern string, alias string, limit int) model.Action {
	return model.Action{
		Kind:       kind,
		Owner:      owner,
		Version:    version,
		Field:      field,
		Mode:       mode,
		Expression: expression,
		Pattern:    pattern,
		Alias:      alias,
		Limit:      limit,
	}
}
