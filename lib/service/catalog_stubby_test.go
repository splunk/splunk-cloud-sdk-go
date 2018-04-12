package service

import (
	"github.com/splunk/ssc-client-go/lib/model"
	"testing"
	"unsafe"
)

// Stubby test for DeleteRule() catalog service endpoint
func TestDeleteRuleWithStubby(t *testing.T) {
	response, err := getSplunkClient().CatalogService.DeleteRule("rule1")
	if response.Status != "200 OK" {
		t.Errorf("FAIL: %s, Expected: %s, Actual: %s ", "Failed to delete rule by making a call to DeleteRule() catalog endpoint", "200 OK", response.Status)
	}
	if err != nil {
		t.Errorf("FAIL: %s, Expected: %s, Actual: %v ", "Failed to delete rule, exception raised", "nil", err)
	}
}

// Stubby test for GetRules() catalog service endpoint
func TestGetRulesWithStubby(t *testing.T) {
	result, err := getSplunkClient().CatalogService.GetRules()
	if result == nil || (result[0].Name != "rule1" || result[0].Match != "newrule" || result[0].Priority != 7 || result[0].Description != "first rule" || len(result[0].Actions) != 3) {
		t.Errorf("FAIL: %s, Expected: %s, Actual: %v ", "Failed to retrieve rules by making a call to GetRules() catalog endpoint", "Rules should be displayed", result)
	}
	if err != nil {
		t.Errorf("FAIL: %s, Expected: %s, Actual: %v ", "Failed to retrieve all rules, exception raised", "nil", err)
	}
}

// Stubby test for PostRule() catalog service endpoint
func TestPostRuleWithStubby(t *testing.T) {
	var actions [3]model.Action
	actions[0] = CreateAction("AUTOKV", "", "", true, "NONE", "", "", "", 0, "")
	actions[1] = CreateAction("EVAL", "", "", false, "", "string", "", "", 0, "string")
	actions[2] = CreateAction("LOOKUP", "", "", false, "", "string", "", "", 0, "")
	result, err := getSplunkClient().CatalogService.PostRule(CreateRule("rule1", "newrule", 7, "first rule", actions[:]))
	if unsafe.Sizeof(result) == 0 || (result.Name != "rule4" || result.Match != "newrule" || result.Priority != 7 || result.Description != "first rule" || len(result.Actions) != 3) {
		t.Errorf("FAIL: %s, Expected: %s, Actual: %v ", "Failed to post rule by making a call to PostRules() catalog endpoint", "Rules should be displayed", result)
	}
	if err != nil {
		t.Errorf("FAIL: %s, Expected: %s, Actual: %v ", "Failed to post rule, exception raised", "nil", err)
	}
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
