/*
 * Copyright 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package main

import (
	"encoding/json"
	"flag"
	"strconv"
	"strings"

	"fmt"

	"net/url"

	"github.com/splunk/splunk-cloud-sdk-go/services/catalog"
)

//
// ./scloud catalog <command> [command-options]
//

const (
	CatalogServiceVersion = "v2beta1"
)

type SharedDatasetFlags struct {
	// Flags used when creating or updating datasets
	Capabilities string
	ReadRoles    []string
	WriteRoles   []string

	ExternalKind       string
	ExternalName       string
	CaseSensitiveMatch bool
	Filter             string
	MaxMatches         int
	MinMatches         int
	DefaultMatch       string

	Datatype string
	Disabled bool
	Search   string
}

var createCatalogService = func() *catalog.Service {
	return apiClient().CatalogService
}

type CatalogCommand struct {
	catalogService *catalog.Service
}

func newCatalogCommand() *CatalogCommand {
	return &CatalogCommand{
		catalogService: createCatalogService(),
	}
}

func (cmd *CatalogCommand) parse(args []string) []string {
	flags := flag.NewFlagSet("catalog command", flag.ExitOnError)
	flags.Parse(args) //nolint:errcheck
	return flags.Args()
}

func (cmd *CatalogCommand) Dispatch(args []string) (result interface{}, err error) {
	arg, args := head(args)
	switch arg {
	case "":
		eusage("too few arguments")
	case "create-dataset":
		result, err = cmd.createDataset(args)
	case "create-field":
		result, err = cmd.createDatasetField(args)
	case "create-rule":
		result, err = cmd.createRule(args)
	case "create-action":
		result, err = cmd.createRuleAction(args)
	case "delete-dataset":
		err = cmd.deleteDataset(args)
	case "delete-field":
		err = cmd.deleteDatasetField(args)
	case "delete-rule":
		err = cmd.deleteRule(args)
	case "delete-action":
		err = cmd.deleteRuleAction(args)
	case "get-dataset":
		result, err = cmd.getDataset(args)
	case "list-datasets":
		result, err = cmd.getDatasets(args)
	case "get-dataset-field":
		result, err = cmd.getDatasetField(args)
	case "list-dataset-fields":
		result, err = cmd.getDatasetFields(args)
	case "get-field":
		result, err = cmd.getField(args)
	case "list-fields":
		result, err = cmd.getFields(args)
	case "list-modules":
		result, err = cmd.getModules(args)
	case "get-rule":
		result, err = cmd.getRule(args)
	case "get-action":
		result, err = cmd.getRuleAction(args)
	case "list-actions":
		result, err = cmd.getRuleActions(args)
	case "list-rules":
		result, err = cmd.getRules(args)
	case "get-spec-json":
		result, err = cmd.getSpecJSON(args)
	case "get-spec-yaml":
		result, err = cmd.getSpecYaml(args)
	case "help":
		err = help("catalog.txt")
	case "update-dataset":
		result, err = cmd.updateDataset(args)
	case "update-field":
		result, err = cmd.updateDatasetField(args)
	case "update-action":
		result, err = cmd.updateRuleAction(args)
	default:
		fatal("unknown command: '%s'", arg)
	}
	return
}

func booleanPointer(input string) (bool, error) {
	value, err := strconv.ParseBool(input)
	if err != nil {
		return value, err
	}
	return value, nil
}

func (cmd *CatalogCommand) createDataset(args []string) (interface{}, error) {
	// Required args
	name, args := head(args)
	kindString, args := head(args)

	if name == "" || kindString == "" {
		etoofew()
	}

	// Optional flags
	flags := flag.NewFlagSet("create-dataset", flag.ExitOnError)
	var fieldsJSON, module string
	var externalKind catalog.LookupDatasetExternalKind

	flags.StringVar(&fieldsJSON, "fields", "[]", "JSON specifying the fields in this dataset")
	flags.StringVar(&module, "module", "", "Module for the dataset")
	sharedDatasetFlags, err := parseDatasetFlags(flags, args)
	if err != nil {
		fatal(err.Error())
	}
	fields := make([]catalog.Field, 0)
	err = json.Unmarshal([]byte(fieldsJSON), &fields)
	if err != nil {
		fatal(err.Error())
	}
	if sharedDatasetFlags.ExternalKind == "kvcollection" {
		externalKind = catalog.LookupDatasetExternalKindKvcollection
	}

	switch kindString {
	case "lookup":
		kind := catalog.LookupDatasetKindLookup
		var fieldPost []catalog.FieldPost
		for _, field := range fields {
			fieldPost = append(fieldPost, catalog.FieldPost{Name: field.Name, Datatype: &field.Datatype, Fieldtype: &field.Fieldtype, Prevalence: &field.Prevalence})
		}
		catalogLookupDatasetPost := catalog.LookupDatasetPost{CaseSensitiveMatch: &sharedDatasetFlags.CaseSensitiveMatch, ExternalKind: externalKind, ExternalName: sharedDatasetFlags.ExternalName, Kind: kind, Fields: fieldPost, Module: &module, Name: name, Filter: &sharedDatasetFlags.Filter}
		return cmd.catalogService.CreateDataset(catalog.MakeDatasetPostFromLookupDatasetPost(catalogLookupDatasetPost))
	case "index":
		kind := catalog.IndexDatasetKindIndex
		var fieldPost []catalog.FieldPost
		for _, field := range fields {
			fieldPost = append(fieldPost, catalog.FieldPost{Name: field.Name, Datatype: &field.Datatype, Fieldtype: &field.Fieldtype, Prevalence: &field.Prevalence})
		}

		catalogIndexDatasetPost := catalog.IndexDatasetPost{Disabled: sharedDatasetFlags.Disabled, Kind: kind, Fields: fieldPost, Module: &module, Name: name}
		return cmd.catalogService.CreateDataset(catalog.MakeDatasetPostFromIndexDatasetPost(catalogIndexDatasetPost))
	case "metric":
		kind := catalog.MetricDatasetKindMetric
		var fieldPost []catalog.FieldPost
		for _, field := range fields {
			fieldPost = append(fieldPost, catalog.FieldPost{Name: field.Name, Datatype: &field.Datatype, Fieldtype: &field.Fieldtype, Prevalence: &field.Prevalence})
		}

		catalogMetricDatasetPost := catalog.MetricDatasetPost{Disabled: sharedDatasetFlags.Disabled, Kind: kind, Fields: fieldPost, Module: &module, Name: name}
		return cmd.catalogService.CreateDataset(catalog.MakeDatasetPostFromMetricDatasetPost(catalogMetricDatasetPost))
	case "kvcollection":
		kind := catalog.KvCollectionDatasetKindKvcollection
		var fieldPost []catalog.FieldPost
		for _, field := range fields {
			fieldPost = append(fieldPost, catalog.FieldPost{Name: field.Name, Datatype: &field.Datatype, Fieldtype: &field.Fieldtype, Prevalence: &field.Prevalence})
		}
		catalogKvCollectionDatasetPost := catalog.KvCollectionDatasetPost{Kind: kind, Fields: fieldPost, Module: &module, Name: name}
		return cmd.catalogService.CreateDataset(catalog.MakeDatasetPostFromKvCollectionDatasetPost(catalogKvCollectionDatasetPost))
	default:
		msg := fmt.Sprintf("'%v' was passed, use subcommand 'index', 'metric', 'kvcollection', or 'lookup'", kindString)
		fatal(msg)

	}
	return catalog.Dataset{}, fmt.Errorf("catalog create dataset failed. please refer to help text for usage")
}

func (cmd *CatalogCommand) createDatasetField(args []string) (interface{}, error) {
	datasetName, args := head(args)
	name, args := head(args)

	// Optional flags
	flags := flag.NewFlagSet("create-field", flag.ExitOnError)
	var dataType, fieldType, prevalence string
	flags.StringVar(&dataType, "data-type", "STRING", "data-type in this dataset")
	flags.StringVar(&fieldType, "field-type", "UNKNOWN", "field-type in this dataset")
	flags.StringVar(&prevalence, "prevalence", "UNKNOWN", "prevalence in this dataset")
	err := flags.Parse(args) //nolint:errcheck
	if err != nil {
		return nil, err
	}

	var fieldDataType catalog.FieldDataType
	var fieldTypeC catalog.FieldType
	var fieldPrevalence catalog.FieldPrevalence

	switch dataType {
	case "STRING":
		fieldDataType = catalog.FieldDataTypeString
	case "DATE":
		fieldDataType = catalog.FieldDataTypeDate
	case "OBJECT_ID":
		fieldDataType = catalog.FieldDataTypeObjectId
	case "NUMBER":
		fieldDataType = catalog.FieldDataTypeNumber
	case "UNKNOWN":
		fieldDataType = catalog.FieldDataTypeUnknown
	default:
		fieldDataType = catalog.FieldDataTypeString

	}
	switch fieldType {
	case "DIMENSION":
		fieldTypeC = catalog.FieldTypeDimension
	case "MEASURE":
		fieldTypeC = catalog.FieldTypeMeasure
	case "UNKNOWN":
		fieldTypeC = catalog.FieldTypeUnknown
	default:
		fieldTypeC = catalog.FieldTypeUnknown
	}
	switch fieldPrevalence {
	case "ALL":
		fieldPrevalence = catalog.FieldPrevalenceAll
	case "SOME":
		fieldPrevalence = catalog.FieldPrevalenceSome
	case "UNKNOWN":
		fieldPrevalence = catalog.FieldPrevalenceUnknown
	default:
		fieldPrevalence = catalog.FieldPrevalenceUnknown
	}

	datasetField := catalog.FieldPost{Name: name, Datatype: &fieldDataType, Fieldtype: &fieldTypeC, Prevalence: &fieldPrevalence}
	return cmd.catalogService.CreateFieldForDataset(datasetName, datasetField)
}

func (cmd *CatalogCommand) createRule(args []string) (interface{}, error) {
	name, args := head(args)
	module, args := head(args)
	match, args := head(args)

	var actionsJSON string
	actions := make([]catalog.ActionPost, 0)

	flags := flag.NewFlagSet("create-rule", flag.ExitOnError)
	flags.StringVar(&actionsJSON, "actions", "[]", "JSON index definition")
	err := flags.Parse(args) //nolint:errcheck
	if err != nil {
		return nil, err
	}
	errUnmarshall := json.Unmarshal([]byte(actionsJSON), &actions)
	if errUnmarshall != nil {
		fatal(errUnmarshall.Error())
	}

	rule := catalog.RulePost{
		Name:    name,
		Module:  &module,
		Match:   match,
		Actions: actions,
	}

	return cmd.catalogService.CreateRule(rule)
}

func (cmd *CatalogCommand) createRuleAction(argv []string) (interface{}, error) {
	// Required args
	ruleName, argv := head(argv)

	// Optional flags
	flags := flag.NewFlagSet("create-action", flag.ExitOnError)
	var field, alias, mode, kindString, expression, pattern string
	var limit int64

	flags.StringVar(&field, "field", "", "Name of the field to be aliased or added, modified by the EVAL expression")
	flags.StringVar(&kindString, "kind", "alias", "Action kind")
	flags.StringVar(&alias, "alias", "", "Action alias name")
	flags.StringVar(&mode, "mode", "", "Action autokv action mode")
	flags.StringVar(&expression, "expression", "", "EVAL expression that calculates the field")
	flags.StringVar(&pattern, "pattern", "", "Regular expression that includes named capture groups for the purpose of field extraction")
	flags.Uint64("limit", 1, "maximum number of times per event to attempt to match fields with the regular expression")

	err := flags.Parse(argv) //nolint:errcheck
	if err != nil {
		return nil, err
	}

	switch kindString {
	case "alias":
		aliasPost := catalog.AliasActionPost{Field: field, Alias: alias, Kind: catalog.AliasActionKindAlias}
		return cmd.catalogService.CreateActionForRule(ruleName, catalog.MakeActionPostFromAliasActionPost(aliasPost))
	case "autokv":
		autoKvPost := catalog.AutoKvActionPost{Mode: mode, Kind: catalog.AutoKvActionKindAutokv}
		return cmd.catalogService.CreateActionForRule(ruleName, catalog.MakeActionPostFromAutoKvActionPost(autoKvPost))
	case "eval":
		evalPost := catalog.EvalActionPost{Field: field, Expression: expression, Kind: catalog.EvalActionKindEval}
		return cmd.catalogService.CreateActionForRule(ruleName, catalog.MakeActionPostFromEvalActionPost(evalPost))
	case "lookup":
		lookupPost := catalog.LookupActionPost{Expression: expression, Kind: catalog.LookupActionKindLookup}
		return cmd.catalogService.CreateActionForRule(ruleName, catalog.MakeActionPostFromLookupActionPost(lookupPost))
	case "regex":
		limit32 := int32(limit)
		regexPost := catalog.RegexActionPost{Field: expression, Limit: &limit32, Pattern: pattern, Kind: catalog.RegexActionKindRegex}
		return cmd.catalogService.CreateActionForRule(ruleName, catalog.MakeActionPostFromRegexActionPost(regexPost))
	default:
		msg := fmt.Sprintf("'%v' was passed, use subcommand 'alias', 'autokv', 'eval', 'lookup','eval', 'regex'", kindString)
		fatal(msg)

	}
	return catalog.Action{}, fmt.Errorf("catalog create dataset failed. please refer to help text for usage")
}

func (cmd *CatalogCommand) deleteDataset(args []string) error {
	datasetName := head1(args)
	return cmd.catalogService.DeleteDataset(datasetName)
}

func (cmd *CatalogCommand) deleteDatasetField(args []string) error {
	datasetName, datasetFieldID := head2(args)
	return cmd.catalogService.DeleteFieldByIdForDataset(datasetFieldID, datasetName)
}

func (cmd *CatalogCommand) deleteRule(args []string) error {
	ruleName := head1(args)
	return cmd.catalogService.DeleteRule(ruleName)
}

func (cmd *CatalogCommand) deleteRuleAction(args []string) error {
	ruleName, actionID := head2(args)
	return cmd.catalogService.DeleteActionByIdForRule(ruleName, actionID)
}

func (cmd *CatalogCommand) getModules(args []string) (interface{}, error) {
	var filter url.Values
	filterBy, _ := head(args)
	if filterBy != "" {
		filter = make(url.Values)
		filter.Set("filter", filterBy)
	}
	var modulesQueryParams = catalog.ListModulesQueryParams{Filter: filterBy}

	return cmd.catalogService.ListModules(&modulesQueryParams)
}

func (cmd *CatalogCommand) getDataset(args []string) (interface{}, error) {
	datasetName := head1(args)
	return cmd.catalogService.GetDataset(datasetName)
}

func (cmd *CatalogCommand) getDatasets(args []string) (interface{}, error) {
	var count string
	var countInt int
	var countInt32 int32
	var filter string
	var orderBy multiFlags

	flags := flag.NewFlagSet("list-datasets", flag.ExitOnError)
	flags.StringVar(&count, "count", "10", "Total number of datasets to return")
	flags.StringVar(&filter, "filter", "", "A filter to apply to the datasets")
	flags.Var(&orderBy, "order-by", "keys to sort by")
	err := flags.Parse(args) //nolint:errcheck
	if err != nil {
		return nil, err
	}

	fsArgs := flags.Args()
	if len(fsArgs) > 0 {
		fatal("unexpected argument(s): %s", fsArgs)
	}

	if count != "" {
		countInt, err = strconv.Atoi(count)
		countInt32 = int32(countInt)
		if err != nil {
			fatal(err.Error())
		}
	}

	return cmd.catalogService.ListDatasets(&catalog.ListDatasetsQueryParams{Filter: filter, Count: &countInt32, Orderby: orderBy})
}

func (cmd *CatalogCommand) getDatasetField(args []string) (interface{}, error) {
	datasetID, datasetFieldID := head2(args)
	return cmd.catalogService.GetFieldByIdForDataset(datasetFieldID, datasetID)
}

func (cmd *CatalogCommand) getDatasetFields(args []string) (interface{}, error) {
	datasetName, args := head(args)
	var urlValues = make(url.Values)

	var filter string
	flags := flag.NewFlagSet("list-dataset-fields", flag.ExitOnError)
	flags.StringVar(&filter, "filter", "", "A filter to apply to the datasets")
	err := flags.Parse(args) //nolint:errcheck
	if err != nil {
		return nil, err
	}

	if len(filter) > 0 {
		urlValues.Set("filter", filter)
	}

	var listDatasetFieldsQueryParams = catalog.ListFieldsForDatasetQueryParams{Filter: filter}

	return cmd.catalogService.ListFieldsForDataset(datasetName, &listDatasetFieldsQueryParams)
}

func (cmd *CatalogCommand) getField(argv []string) (interface{}, error) {
	fieldID := head1(argv)
	return cmd.catalogService.GetFieldById(fieldID)
}

func (cmd *CatalogCommand) getFields(argv []string) (interface{}, error) {
	var filter url.Values
	filterBy, _ := head(argv)
	if filterBy != "" {
		filter = make(url.Values)
		filter.Set("filter", filterBy)
	}
	var listFieldsQueryParams = catalog.ListFieldsQueryParams{Filter: filterBy}

	return cmd.catalogService.ListFields(&listFieldsQueryParams)
}

func (cmd *CatalogCommand) getRule(args []string) (interface{}, error) {
	ruleName := head1(args)
	return cmd.catalogService.GetRule(ruleName)
}

func (cmd *CatalogCommand) getRuleAction(argv []string) (interface{}, error) {
	ruleName, actionID := head2(argv)
	return cmd.catalogService.GetActionByIdForRule(ruleName, actionID)
}

func (cmd *CatalogCommand) getRuleActions(argv []string) (interface{}, error) {
	ruleName := head1(argv)
	var filterBy url.Values
	var filter string
	flags := flag.NewFlagSet("list-actions", flag.ExitOnError)
	flags.StringVar(&filter, "filter", "", "A filter to apply to the actions")
	err := flags.Parse(argv) //nolint:errcheck
	if err != nil {
		return nil, err
	}

	if len(filter) > 0 {
		filterBy.Set("filter", filter)
	}
	var listActionRulesQueryParams = catalog.ListActionsForRuleQueryParams{Filter: filter}
	return cmd.catalogService.ListActionsForRule(ruleName, &listActionRulesQueryParams)
}

func (cmd *CatalogCommand) getRules(args []string) (interface{}, error) {
	var filter url.Values
	filterBy, _ := head(args)
	if filterBy != "" {
		filter = make(url.Values)
		filter.Set("filter", filterBy)
	}
	var listRulesQueryParams = catalog.ListRulesQueryParams{Filter: filterBy}
	return cmd.catalogService.ListRules(&listRulesQueryParams)
}

func parseDatasetFlags(flags *flag.FlagSet, args []string) (*SharedDatasetFlags, error) {
	sharedDatasetFlags := &SharedDatasetFlags{}

	var caseSensitiveMatch, disabled, readRolesString, writeRolesString string
	flags.StringVar(&sharedDatasetFlags.Capabilities, "capabilities", "", "Capabilities for the dataset")
	flags.StringVar(&sharedDatasetFlags.ExternalKind, "external-kind", "",
		"The kind of the external lookup, for lookup datasets")
	flags.StringVar(&sharedDatasetFlags.ExternalName, "external-name", "",
		"The name of the external lookup, for lookup datasets")
	flags.StringVar(&readRolesString, "read-roles", "", "CSV of read roles for the dataset")
	flags.StringVar(&writeRolesString, "write-roles", "", "CSV of write roles for the dataset")
	flags.StringVar(&caseSensitiveMatch, "case-sensitive-match", "",
		"Match case-sensitively against the lookup. Applicable for lookup datasets.")
	flags.StringVar(&sharedDatasetFlags.Filter, "filter", "",
		"Filter results from the lookup before returning data.")
	flags.IntVar(&sharedDatasetFlags.MaxMatches, "max-matches", -1,
		"The maximum number of matches that should be returned. Applicable for lookup datasets.")
	flags.IntVar(&sharedDatasetFlags.MinMatches, "min-matches", -1,
		"The minimum number of matches that should be returned. Applicable for lookup datasets.")
	flags.StringVar(&sharedDatasetFlags.DefaultMatch, "default-match", "", "The default match. If minMatches > 0 and no matches are found, this value is returned instead. "+
		"Applicable for lookup datasets")
	flags.StringVar(&sharedDatasetFlags.Datatype, "data-type", "",
		"The type of data in this dataset")
	flags.StringVar(&disabled, "disabled", "",
		"Whether or not the Splunk index is disabled. Applicable for index datasets")
	flags.StringVar(&sharedDatasetFlags.Search, "search", "",
		"A valid SPL-defined search")

	err := flags.Parse(args)
	if err != nil {
		return nil, err
	}

	if caseSensitiveMatch != "" {
		sharedDatasetFlags.CaseSensitiveMatch, err = booleanPointer(caseSensitiveMatch)
		if err != nil {
			return nil, err
		}
	}

	if disabled != "" {
		sharedDatasetFlags.Disabled, err = booleanPointer(disabled)
		if err != nil {
			return nil, err
		}
	}

	if len(readRolesString) > 0 {
		sharedDatasetFlags.ReadRoles = strings.Split(readRolesString, ",")
	}
	if len(writeRolesString) > 0 {
		sharedDatasetFlags.WriteRoles = strings.Split(writeRolesString, ",")
	}

	return sharedDatasetFlags, nil
}

func (cmd *CatalogCommand) updateDataset(args []string) (interface{}, error) {
	// Required args
	name, args := head(args)

	// Optional flags
	flags := flag.NewFlagSet("update-dataset", flag.ExitOnError)
	var fieldsJSON, module string
	var externalKind catalog.LookupDatasetExternalKind

	flags.StringVar(&fieldsJSON, "fields", "[]", "JSON specifying the fields in this dataset")
	flags.StringVar(&module, "module", "", "Module for the dataset")
	sharedDatasetFlags, err := parseDatasetFlags(flags, args)
	if err != nil {
		fatal(err.Error())
	}
	fields := make([]catalog.Field, 0)
	err = json.Unmarshal([]byte(fieldsJSON), &fields)
	if err != nil {
		fatal(err.Error())
	}
	if sharedDatasetFlags.ExternalKind == "kvcollection" {
		externalKind = catalog.LookupDatasetExternalKindKvcollection
	}

	dataset, err := cmd.catalogService.GetDataset(name)

	if dataset.IsLookupDataset() {
		catalogLookupDatasetPatch := catalog.LookupDatasetPatch{CaseSensitiveMatch: &sharedDatasetFlags.CaseSensitiveMatch, ExternalKind: &externalKind, ExternalName: &sharedDatasetFlags.ExternalName, Module: &module, Name: &name, Filter: &sharedDatasetFlags.Filter}
		return cmd.catalogService.UpdateDataset(name, catalog.MakeDatasetPatchFromLookupDatasetPatch(catalogLookupDatasetPatch))
	}
	if dataset.IsIndexDataset() {
		catalogIndexDatasetPatch := catalog.IndexDatasetPatch{Disabled: &sharedDatasetFlags.Disabled, Module: &module, Name: &name}
		return cmd.catalogService.UpdateDataset(name, catalog.MakeDatasetPatchFromIndexDatasetPatch(catalogIndexDatasetPatch))
	}
	if dataset.IsKvCollectionDataset() {
		catalogKvCollectionDatasetPatch := catalog.KvCollectionDatasetPatch{Module: &module, Name: &name}
		return cmd.catalogService.UpdateDataset(name, catalog.MakeDatasetPatchFromKvCollectionDatasetPatch(catalogKvCollectionDatasetPatch))
	}
	return catalog.Dataset{}, fmt.Errorf("catalog create dataset failed. please refer to help text for usage")
}

func (cmd *CatalogCommand) updateDatasetField(args []string) (interface{}, error) {
	// Required args
	datasetName, args := head(args)
	datasetFieldID, args := head(args)

	if datasetName == "" || datasetFieldID == "" {
		etoofew()
	}

	// Optional flags
	flags := flag.NewFlagSet("update-field", flag.ExitOnError)
	var dataType, fieldType, prevalence, name string
	flags.StringVar(&dataType, "data-type", "STRING", "data-type in this dataset")
	flags.StringVar(&fieldType, "field-type", "UNKNOWN", "field-type in this dataset")
	flags.StringVar(&prevalence, "prevalence", "UNKNOWN", "prevalence in this dataset")
	flags.StringVar(&name, "name", "", "field name")
	err := flags.Parse(args) //nolint:errcheck
	if err != nil {
		return nil, err
	}

	var fieldDataType catalog.FieldDataType
	var fieldTypeC catalog.FieldType
	var fieldPrevalence catalog.FieldPrevalence

	switch dataType {
	case "STRING":
		fieldDataType = catalog.FieldDataTypeString
	case "DATE":
		fieldDataType = catalog.FieldDataTypeDate
	case "OBJECT_ID":
		fieldDataType = catalog.FieldDataTypeObjectId
	case "NUMBER":
		fieldDataType = catalog.FieldDataTypeNumber
	case "UNKNOWN":
		fieldDataType = catalog.FieldDataTypeUnknown
	default:
		fieldDataType = catalog.FieldDataTypeString
	}
	switch fieldType {
	case "DIMENSION":
		fieldTypeC = catalog.FieldTypeDimension
	case "MEASURE":
		fieldTypeC = catalog.FieldTypeMeasure
	case "UNKNOWN":
		fieldTypeC = catalog.FieldTypeUnknown
	default:
		fieldTypeC = catalog.FieldTypeUnknown
	}
	switch fieldPrevalence {
	case "ALL":
		fieldPrevalence = catalog.FieldPrevalenceAll
	case "SOME":
		fieldPrevalence = catalog.FieldPrevalenceSome
	case "UNKNOWN":
		fieldPrevalence = catalog.FieldPrevalenceUnknown
	default:
		fieldPrevalence = catalog.FieldPrevalenceUnknown
	}

	datasetField := catalog.FieldPatch{Name: &name, Datatype: &fieldDataType, Fieldtype: &fieldTypeC, Prevalence: &fieldPrevalence}

	return cmd.catalogService.UpdateFieldByIdForDataset(datasetFieldID, datasetName, datasetField)
}

func (cmd *CatalogCommand) updateRuleAction(argv []string) (interface{}, error) {
	// Required args
	ruleName, args := head(argv)
	actionID, args := head(args)

	if ruleName == "" || actionID == "" {
		etoofew()
	}

	// Optional flags
	flags := flag.NewFlagSet("update-action", flag.ExitOnError)
	var field, alias, mode, expression, pattern string
	var limit int64

	flags.StringVar(&field, "field", "", "Name of the field to be aliased or added, modified by the EVAL expression or matched by the regular expression")
	flags.StringVar(&alias, "alias", "", "Action alias name")
	flags.StringVar(&mode, "mode", "", "Action autokv action mode")
	flags.StringVar(&expression, "expression", "", "EVAL expression that calculates the field")
	flags.StringVar(&pattern, "pattern", "", "Regular expression that includes named capture groups for the purpose of field extraction")
	flags.Uint64("limit", 1, "maximum number of times per event to attempt to match fields with the regular expression")

	err := flags.Parse(args) //nolint:errcheck
	if err != nil {
		return nil, err
	}

	action, err := cmd.catalogService.GetActionByIdForRuleById(ruleName, actionID)
	if err != nil {
		return nil, err
	}

	if action.IsAliasAction() {

		catalogActionPatch := catalog.MakeActionPatchFromAliasActionPatch(catalog.AliasActionPatch{Alias: &alias, Field: &field})
		return cmd.catalogService.UpdateActionByIdForRule(ruleName, actionID, catalogActionPatch)
	}
	if action.IsAutoKvAction() {
		catalogActionPatch := catalog.MakeActionPatchFromAutoKvActionPatch(catalog.AutoKvActionPatch{Mode: &mode})
		return cmd.catalogService.UpdateActionByIdForRule(ruleName, actionID, catalogActionPatch)
	}
	if action.IsEvalAction() {
		catalogActionPatch := catalog.MakeActionPatchFromEvalActionPatch(catalog.EvalActionPatch{Expression: &expression, Field: &field})
		return cmd.catalogService.UpdateActionByIdForRule(ruleName, actionID, catalogActionPatch)
	}
	if action.IsLookupAction() {
		catalogActionPatch := catalog.MakeActionPatchFromLookupActionPatch(catalog.LookupActionPatch{Expression: &expression})
		return cmd.catalogService.UpdateActionByIdForRule(ruleName, actionID, catalogActionPatch)
	}
	if action.IsRegexAction() {
		limit32 := int32(limit)
		catalogActionPatch := catalog.MakeActionPatchFromRegexActionPatch(catalog.RegexActionPatch{Field: &field, Limit: &limit32, Pattern: &pattern})
		return cmd.catalogService.UpdateActionByIdForRule(ruleName, actionID, catalogActionPatch)
	}
	return catalog.Action{}, fmt.Errorf("catalog create dataset failed. please refer to help text for usage")
}

func (cmd *CatalogCommand) getSpecJSON(args []string) (interface{}, error) {
	checkEmpty(args)
	return GetSpecJSON("api", CatalogServiceVersion, "catalog", cmd.catalogService.Client)
}

func (cmd *CatalogCommand) getSpecYaml(args []string) (interface{}, error) {
	checkEmpty(args)
	return GetSpecYaml("api", CatalogServiceVersion, "catalog", cmd.catalogService.Client)
}
