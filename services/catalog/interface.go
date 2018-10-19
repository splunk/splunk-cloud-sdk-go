// AUTO GENERATED. DO NOT EDIT!
package catalog

import (
	"net/url"
)

type CatalogIface interface {
	// ListDatasets returns all Datasets with optional filter, count, or orderby params
	ListDatasets(values url.Values) ([]DatasetInfo, error)
	// GetDatasets returns all Datasets
	// Deprecated: v0.6.1 - Use ListDatasets instead
	GetDatasets() ([]DatasetInfo, error)
	// GetDataset returns the Dataset by resourceName or ID
	GetDataset(resourceNameOrID string) (*DatasetInfo, error)
	// CreateDataset creates a new Dataset
	CreateDataset(dataset *DatasetCreationPayload) (*DatasetInfo, error)
	// UpdateDataset updates an existing Dataset with the specified resourceName or ID
	UpdateDataset(dataset *UpdateDatasetInfoFields, resourceNameOrID string) (*DatasetInfo, error)
	// DeleteDataset implements delete Dataset endpoint with the specified resourceName or ID
	DeleteDataset(resourceNameOrID string) error
	// DeleteRule deletes the rule and its dependencies with the specified rule id or resourceName
	DeleteRule(resourceNameOrID string) error
	// GetRules returns all the rules.
	GetRules() ([]Rule, error)
	// GetRule returns rule by the specified resourceName or ID.
	GetRule(resourceNameOrID string) (*Rule, error)
	// CreateRule posts a new rule.
	CreateRule(rule Rule) (*Rule, error)
	// UpdateRule updates the rule with the specified resourceName or ID
	UpdateRule(resourceNameOrID string, rule *RuleUpdateFields) (*Rule, error)
	// GetDatasetFields returns all the fields belonging to the specified dataset
	GetDatasetFields(datasetID string, values url.Values) ([]Field, error)
	// GetDatasetField returns the field belonging to the specified dataset with the id datasetFieldID
	GetDatasetField(datasetID string, datasetFieldID string) (*Field, error)
	// CreateDatasetField creates a new field in the specified dataset
	CreateDatasetField(datasetID string, datasetField *Field) (*Field, error)
	// UpdateDatasetField updates an already existing field in the specified dataset
	UpdateDatasetField(datasetID string, datasetFieldID string, datasetField *Field) (*Field, error)
	// DeleteDatasetField deletes the field belonging to the specified dataset with the id datasetFieldID
	DeleteDatasetField(datasetID string, datasetFieldID string) error
	// GetFields returns a list of all Fields on Catalog
	GetFields() ([]Field, error)
	// GetField returns the Field corresponding to fieldid
	GetField(fieldID string) (*Field, error)
	// CreateRuleAction creates a new Action on the rule specified
	CreateRuleAction(ruleID string, action *Action) (*Action, error)
	// GetRuleActions returns a list of all actions belonging to the specified rule
	GetRuleActions(ruleID string) ([]Action, error)
	// GetRuleAction returns the action of specified belonging to the specified rule
	GetRuleAction(ruleID string, actionID string) (*Action, error)
	// DeleteRuleAction deletes the action of specified belonging to the specified rule
	DeleteRuleAction(ruleID string, actionID string) error
	// UpdateRuleAction updates the action with the specified id for the specified Rule
	UpdateRuleAction(ruleID string, actionID string, action *Action) (*Action, error)
	// GetModules returns a list of a list of modules that match a filter query if it is given, otherwise return all modules
	GetModules(filter url.Values) ([]Module, error)
}
