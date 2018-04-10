package model

import (
)

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
	Name        string               `json:"name"`
	Actions     []Action             `json:"actions"`
	Match       string               `json:"match"`
	Priority    int                  `json:"priority"`
	Description string               `json:"description"`
}
type Rules []Rule

type Action struct {
	Kind           ActionKind         `json:"kind"`
	Field          string             `json:"field,omitempty"`
	Alias          string             `json:"alias,omitempty"`
	Trim           bool               `json:"trim,omitempty"`
	Mode           AutoMode           `json:"mode,omitempty"`
	Expression     string             `json:"expression,omitempty"`
	Pattern        string             `json:"pattern,omitempty"`
	Format         string             `json:"format,omitempty"`
	Limit          int                `json:"limit,omitempty"`
	Result         string             `json:"result,omitempty"`

}

type AutoMode string