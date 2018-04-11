package model

//DatasetKind defines that the kind in a dataset
type DatasetKind string

const (
	// VIEW object
	VIEW    DatasetKind = "VIEW"
	// INDEX object
	INDEX   DatasetKind = "INDEX"
	// KVSTORE object
	KVSTORE DatasetKind = "KVSTORE"
	// EXTERN object
	EXTERN  DatasetKind = "EXTERN"
	// TOPIC object
	TOPIC   DatasetKind = "TOPIC"
	// CATALOG object
	CATALOG DatasetKind = "CATALOG"
)

// Dataset represent a knowledge object in Splunk
type Dataset struct {
	ID    string      `json:"id"`
	Name  string      `json:"name"`
	Kind  DatasetKind `json:"kind"`
	Rules []string    `json:"rules"`
	Todo  string      `json:"todo"`
}

// Datasets is a set of Dataset
type Datasets []Dataset

// ActionKind type: define the action kind of a rule
type ActionKind string

const (
	// ALIAS action
	ALIAS  ActionKind = "ALIAS"
	// AUTOKV kv action
	AUTOKV ActionKind = "AUTOKV"
	// REGEX action
	REGEX  ActionKind = "REGEX"
	// EVAL action
	EVAL   ActionKind = "EVAL"
	// LOOKUP action
	LOOKUP ActionKind = "LOOKUP"
)

// Rule type
type Rule struct {
	Name        string       `json:"name"`
	Action      []ActionKind `json:"actions"`
	Match       string       `json:"match"`
	Priority    int          `json:"priority"`
	Description string       `json:"description"`
}

// Rules is a set of rule
type Rules []Rule

