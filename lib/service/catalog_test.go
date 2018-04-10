package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/splunk/ssc-client-go/lib/model"
	"fmt"
)

func getSplunkClient() *Client {

	return NewSplunkdClient("",
		[2]string{"admin", "changeme"},
		"localhost:32769", "http", nil)
}

func Test_getDataset(t *testing.T) {

	result, err := getSplunkClient().CatalogService.GetDataset("ds1")
	assert.Empty(t, err)
	assert.NotEmpty(t, result.Id)
	assert.Equal(t, "ds1", result.Name)
	assert.Equal(t, model.VIEW, result.Kind)
}

func Test_getDatasets(t *testing.T) {

	result, err := getSplunkClient().CatalogService.GetDatasets()
	assert.Empty(t, err)
	assert.Equal(t, 2, len(result))
}

func Test_postDataset(t *testing.T) {

	//dataset := model.Dataset_post{"ds1", model.VIEW, []string{"string"}, "string"}
	result, err := getSplunkClient().CatalogService.PostDataset(
		getSplunkClient().CatalogService.CreateDataset("ds1", model.VIEW, []string{"string"}, "string"))
	assert.Empty(t, err)
	assert.NotEmpty(t, result.Id)
	assert.Equal(t, "ds1", result.Name)
	assert.Equal(t, model.VIEW, result.Kind)
	assert.Equal(t, []string{"string"}, result.Rules)
}

// TODO Move the stubby test to a separate folder
// Stubby test for DeleteRule() catalog service endpoint
func TestDeleteRule(t *testing.T) {
	result, err := getSplunkClient().CatalogService.deleteRule("rule1")
	fmt.Println(err)
	fmt.Println(result)
}

// Stubby test for GetRules() catalog service endpoint
func TestGetRules(t *testing.T) {
	result, err := getSplunkClient().CatalogService.GetRules()
	fmt.Println(err)
	fmt.Println(result)

}

// Stubby test for PostRule() catalog service endpoint
func TestPostRule(t *testing.T) {
	var actions[3]model.Action
	actions[0] = CreateAction("AUTOKV","" ,"",true, "NONE", "", "", "", 0, "")
	actions[1] = CreateAction("EVAL","" ,"", false, "", "string", "", "", 0, "string")
	actions[2] = CreateAction("LOOKUP","" ,"",false, "", "string", "", "", 0, "")
	result, err := getSplunkClient().CatalogService.PostRule(CreateRule("rule1","newrule",7,"first rule", actions[:]))
	fmt.Println(err)
	fmt.Println(result)
}

// creates a rule to post
func CreateRule(name string, match string, priority int, description string, actions []model.Action) model.Rule {
	return model.Rule{
		Name:         name,
		Match:        match,
		Priority:     priority,
		Description:  description,
		Actions:      actions,
	}
}

// creates an action for rule to post
func CreateAction(kind model.ActionKind, field string, alias string, trim bool, mode model.AutoMode, expression string, pattern string, format string, limit int, result string) model.Action {
	return model.Action{
		Kind:        kind,
		Field:       field,
		Alias:       alias,
		Trim:        trim,
		Mode:        mode,
		Expression:  expression,
		Pattern:     pattern,
		Format:      format,
		Limit:       limit,
		Result:      result,
	}
}