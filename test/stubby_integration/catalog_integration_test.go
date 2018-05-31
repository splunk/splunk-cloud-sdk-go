package stubbyintegration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/splunk/ssc-client-go/model"
)

// Stubby test for GetDataset() catalog service endpoint
func TestGetDataset(t *testing.T) {
	result, err := getClient(t).CatalogService.GetDataset("ds1")

	assert.Empty(t, err)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, "ds1", result.Name)
	assert.Equal(t, model.VIEW, result.Kind)
}

// Stubby test for GetDatasets() catalog service endpoint
func TestGetDatasets(t *testing.T) {
	result, err := getClient(t).CatalogService.GetDatasets()

	assert.Empty(t, err)
	assert.Equal(t, 2, len(result))
}

// Stubby test for CreateDataset() catalog service endpoint
func TestPostDataset(t *testing.T) {
	testDataset := model.Dataset{Name: "ds1", Kind: model.VIEW, Rules: []string{"string"}, Todo: "string"}
	result, err := getClient(t).CatalogService.CreateDataset(testDataset)

	assert.Empty(t, err)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, "ds1", result.Name)
	assert.Equal(t, model.VIEW, result.Kind)
	assert.Equal(t, []string{"string"}, result.Rules)
}

// Stubby test for DeleteDataset() catalog service endpoint
func TestDeleteDataset(t *testing.T) {
	err := getClient(t).CatalogService.DeleteDataset("ds1")
	assert.Empty(t, err)
}

// Stubby test for DeleteRule() catalog service endpoint
func TestDeleteRule(t *testing.T) {
	err := getClient(t).CatalogService.DeleteRule("rule1")
	assert.Empty(t, err)
}

// Stubby test for GetRules() catalog service endpoint
func TestGetRules(t *testing.T) {
	result, err := getClient(t).CatalogService.GetRules()

	assert.Empty(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, 3, len(result[0].Actions))
}

// Stubby test for CreateRule() catalog service endpoint
func TestPostRule(t *testing.T) {
	var actions [3]model.Action
	actions[0] = CreateAction("AUTOKV", "", "", true, "NONE", "", "", "", 0, "")
	actions[1] = CreateAction("EVAL", "", "", false, "", "string", "", "", 0, "string")
	actions[2] = CreateAction("LOOKUP", "", "", false, "", "string", "", "", 0, "")
	result, err := getClient(t).CatalogService.CreateRule(CreateRule("rule1", "newrule", 7, "first rule", actions[:]))

	assert.Empty(t, err)
	assert.Equal(t, "rule4", result.Name)
	assert.Equal(t, "newrule", result.Match)
	assert.Equal(t, 3, len(result.Actions))
}

// creates a rule to post
func CreateRule(name string, match string, priority int, description string, actions []model.Action) model.Rule {
	return model.Rule{
		Name:        name,
		Match:       match,
		Priority:    priority,
		Description: description,
		Actions:     actions,
	}
}

// creates an action for rule to post
func CreateAction(kind model.ActionKind, field string, alias string, trim bool, mode model.AutoMode, expression string, pattern string, format string, limit int, result string) model.Action {
	return model.Action{
		Kind:       kind,
		Field:      field,
		Alias:      alias,
		Trim:       trim,
		Mode:       mode,
		Expression: expression,
		Pattern:    pattern,
		Format:     format,
		Limit:      limit,
		Result:     result,
	}
}
