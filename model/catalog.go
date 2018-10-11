// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package model

import (
	"github.com/splunk/splunk-cloud-sdk-go/services/catalog"
)

// DatasetInfoKind is Deprecated: please use services/catalog.DatasetInfoKind
type DatasetInfoKind = catalog.DatasetInfoKind

const (
	// LOOKUP is Deprecated: please use services/catalog.Lookup
	LOOKUP DatasetInfoKind = catalog.Lookup
	// KVCOLLECTION is Deprecated: please use services/catalog.KvCollection
	KVCOLLECTION DatasetInfoKind = catalog.KvCollection
	// INDEX is Deprecated: please use services/catalog.Index
	INDEX DatasetInfoKind = catalog.Index
)

// DatasetInfo is Deprecated: please use services/catalog.DatasetInfo
type DatasetInfo = catalog.DatasetInfo

// DatasetCreationPayload is Deprecated: please use services/catalog.DatasetCreationPayload
type DatasetCreationPayload = catalog.DatasetCreationPayload

// UpdateDatasetInfoFields is Deprecated: please use services/catalog.UpdateDatasetInfoFields
type UpdateDatasetInfoFields = catalog.UpdateDatasetInfoFields

// Field is Deprecated: please use services/catalog.Field
type Field = catalog.Field

// PrevalenceType is Deprecated: please use services/catalog.PrevalenceType
type PrevalenceType = catalog.PrevalenceType

const (
	// ALL is Deprecated: please use services/catalog.All
	ALL PrevalenceType = catalog.All
	// SOME is Deprecated: please use services/catalog.Some
	SOME PrevalenceType = catalog.Some
	// PREVALANCEUNKNOWN is Deprecated: please use services/catalog.PrevalenceUnknown
	PREVALANCEUNKNOWN PrevalenceType = catalog.PrevalenceUnknown
)

// DataType is Deprecated: please use services/catalog.DataType
type DataType = catalog.DataType

const (
	// DATE is Deprecated: please use services/catalog.Date
	DATE DataType = catalog.Date
	// NUMBER is Deprecated: please use services/catalog.Number
	NUMBER DataType = catalog.Number
	// OBJECTID is Deprecated: please use services/catalog.ObjectID
	OBJECTID DataType = catalog.ObjectID
	// STRING is Deprecated: please use services/catalog.String
	STRING DataType = catalog.String
	// DATATYPEUNKNOWN is Deprecated: please use services/catalog.DataTypeUnknown
	DATATYPEUNKNOWN DataType = catalog.DataTypeUnknown
)

// FieldType is Deprecated: please use services/catalog.FieldType
type FieldType = catalog.FieldType

const (
	// DIMENSION is Deprecated: please use services/catalog.Dimension
	DIMENSION FieldType = catalog.Dimension
	// MEASURE is Deprecated: please use services/catalog.Measure
	MEASURE FieldType = catalog.Measure
	// FIELDTYPEUNKNOWN is Deprecated: please use services/catalog.FieldTypeUnknown
	FIELDTYPEUNKNOWN = catalog.FieldTypeUnknown
)

// CatalogActionKind is Deprecated: please use services/catalog.ActionKind
type CatalogActionKind = catalog.ActionKind

const (
	// ALIAS is Deprecated: please use services/catalog.Alias
	ALIAS CatalogActionKind = catalog.Alias
	// AUTOKV is Deprecated: please use services/catalog.AutoKV
	AUTOKV CatalogActionKind = catalog.AutoKV
	// REGEX is Deprecated: please use services/catalog.Regex
	REGEX CatalogActionKind = catalog.Regex
	// EVAL is Deprecated: please use services/catalog.Eval
	EVAL CatalogActionKind = catalog.Eval
	// LOOKUPACTION is Deprecated: please use services/catalog.LookupAction
	LOOKUPACTION CatalogActionKind = catalog.LookupAction
)

// Rule is Deprecated: please use services/catalog.Rule
type Rule = catalog.Rule

// RuleUpdateFields is Deprecated: please use services/catalog.RuleUpdateFields
type RuleUpdateFields = catalog.RuleUpdateFields

// CatalogAction is Deprecated: please use services/catalog.Action
type CatalogAction = catalog.Action

// Module is Deprecated: please use services/catalog.Module
type Module = catalog.Module

// NewAliasAction is Deprecated: please use services/catalog.NewAliasAction
func NewAliasAction(field string, alias string, owner string) *CatalogAction {
	return catalog.NewAliasAction(field, alias, owner)
}

// NewAutoKVAction is Deprecated: please use services/catalog.NewAutoKVAction
func NewAutoKVAction(mode string, owner string) *CatalogAction {
	return catalog.NewAutoKVAction(mode, owner)
}

// NewEvalAction is Deprecated: please use services/catalog.NewEvalAction
func NewEvalAction(field string, expression string, owner string) *CatalogAction {
	return catalog.NewEvalAction(field, expression, owner)
}

// NewLookupAction is Deprecated: please use services/catalog.NewLookupAction
func NewLookupAction(expression string, owner string) *CatalogAction {
	return catalog.NewLookupAction(expression, owner)
}

// NewRegexAction is Deprecated: please use services/catalog.NewRegexAction
func NewRegexAction(field string, pattern string, limit *int, owner string) *CatalogAction {
	return catalog.NewRegexAction(field, pattern, limit, owner)
}

// NewUpdateAliasAction is Deprecated: please use services/catalog.NewUpdateAliasAction
func NewUpdateAliasAction(field *string, alias *string) *CatalogAction {
	return catalog.NewUpdateAliasAction(field, alias)
}

// NewUpdateAutoKVAction is Deprecated: please use services/catalog.NewUpdateAutoKVAction
func NewUpdateAutoKVAction(mode *string) *CatalogAction {
	return catalog.NewUpdateAutoKVAction(mode)
}

// NewUpdateEvalAction is Deprecated: please use services/catalog.NewUpdateEvalAction
func NewUpdateEvalAction(field *string, expression *string) *CatalogAction {
	return catalog.NewUpdateEvalAction(field, expression)
}

// NewUpdateLookupAction is Deprecated: please use services/catalog.NewUpdateLookupAction
func NewUpdateLookupAction(expression *string) *CatalogAction {
	return catalog.NewUpdateLookupAction(expression)
}

// NewUpdateRegexAction is Deprecated: please use services/catalog.NewUpdateRegexAction
func NewUpdateRegexAction(field *string, pattern *string, limit *int) *CatalogAction {
	return catalog.NewUpdateRegexAction(field, pattern, limit)
}
