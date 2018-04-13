package model

//DatasetKind enumerates the kinds of datasets known to the system.
type DatasetKind string

const (
	// VIEW represents a view over base data in some other dataset.
	// The view  consists of a Splunk query (with at least a generating command)
	// and an optional collection of search time transformation rules.
	VIEW DatasetKind = "VIEW"
	// INDEX represents a Splunk events or metrics index
	INDEX DatasetKind = "INDEX"
	// KVSTORE represents an instance of the KV storage service as a dataset
	KVSTORE DatasetKind = "KVSTORE"
	// EXTERN represents an extern REST API based dataset
	EXTERN DatasetKind = "EXTERN"
	// TOPIC represents a message bus topic as a dataset.
	TOPIC DatasetKind = "TOPIC"
	// CATALOG represents the metadata catalog as a dataset
	CATALOG DatasetKind = "CATALOG"
)

// Dataset represents the sources of data that can be serched by Splunk
type Dataset struct {
	ID    string      `json:"id"`
	Name  string      `json:"name"`
	Kind  DatasetKind `json:"kind"`
	Rules []string    `json:"rules"`
	Todo  string      `json:"todo"`
}


// ActionKind enumerates the kinds of search time transformation action known by the service.
type ActionKind string

const (
	// ALIAS action
	ALIAS ActionKind = "ALIAS"
	// AUTOKV action
	AUTOKV ActionKind = "AUTOKV"
	// REGEX action
	REGEX ActionKind = "REGEX"
	// EVAL action
	EVAL ActionKind = "EVAL"
	// LOOKUP action
	LOOKUP ActionKind = "LOOKUP"
)

// Rule represents a rule for transforming results at search time.
// A rule consits of a `match` clause and a collection of transformation actions
type Rule struct {
	Name        string       `json:"name"`
	Action      []ActionKind `json:"actions"`
	Match       string       `json:"match"`
	Priority    int          `json:"priority"`
	Description string       `json:"description"`
}
