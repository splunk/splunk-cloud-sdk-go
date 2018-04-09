package model

type DatasetKind string

const (
	VIEW    DatasetKind = "VIEW"
	INDEX   DatasetKind = "INDEX"
	KVSTORE DatasetKind = "KVSTORE"
	EXTERN  DatasetKind = "EXTERN"
	TOPIC   DatasetKind = "TOPIC"
	CATALOG DatasetKind = "CATALOG"
)

type Dataset struct {
	Id    string      `json:"id"`
	Name  string      `json:"name"`
	Kind  DatasetKind `json:"kind"`
	Rules []string    `json:"rules"`
	Todo  string      `json:"todo"`
}

type Datasets []Dataset

type ActionKind string

const (
	ALIAS  ActionKind = "ALIAS"
	AUTOKV ActionKind = "AUTOKV"
	REGEX  ActionKind = "REGEX"
	EVAL   ActionKind = "EVAL"
	LOOKUP ActionKind = "LOOKUP"
)

type Rule struct {
	Name        string       `json:"name"`
	Action      []ActionKind `json:"actions"`
	Match       string       `json:"match"`
	Priority    int          `json:"priority"`
	Description string       `json:"description"`
}
type Rules []Rule

