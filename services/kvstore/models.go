// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package kvstore

// Error error reason
type Error struct {

	// The reason of the error
	Code int64 `json:"code"`
	// Error message
	Message string `json:"message"`
	// State Storage error code
	SsCode int64 `json:"ssCode"`
}

// AuthError auth error reason
type AuthError struct {

	// The reason of the auth error
	Reason string `json:"reason"`
}

// PingOKBody ping ok body
type PingOKBody struct {

	// If database is not healthy, detailed error message
	ErrorMessage string `json:"errorMessage,omitempty"`

	// Database status
	// Enum: [healthy unhealthy unknown]
	Status PingOKBodyStatus `json:"status"`
}

// PingOKBodyStatus used to force type expectation for KVStore Ping endpoint response
type PingOKBodyStatus string

const (
	// PingOKBodyStatusHealthy captures enum value "healthy"
	PingOKBodyStatusHealthy PingOKBodyStatus = "healthy"

	// PingOKBodyStatusUnhealthy captures enum value "unhealthy"
	PingOKBodyStatusUnhealthy PingOKBodyStatus = "unhealthy"

	// PingOKBodyStatusUnknown captures enum value "unknown"
	PingOKBodyStatusUnknown PingOKBodyStatus = "unknown"
)

// IndexFieldDefinition index field definition
type IndexFieldDefinition struct {

	// The sort direction for the indexed field
	Direction int64 `json:"direction"`

	// The name of the field to index
	Field string `json:"field"`
}

// IndexDefinition index field definition
type IndexDefinition struct {

	// The name of the index
	Name string `json:"name,omitempty"`

	// fields
	Fields []IndexFieldDefinition `json:"fields"`
}

// IndexDescription index description
type IndexDescription struct {

	// The collection name
	Collection string `json:"collection,omitempty"`

	// fields
	Fields []IndexFieldDefinition `json:"fields"`

	// The name of the index
	Name string `json:"name,omitempty"`
}

// LookupValue Value tuple used for lookup
type LookupValue []interface{}

// Key to identify a record in a collection
type Key struct {
	Key string `json:"_key"`
}

// Record is a JSON document entity contained in collections
type Record map[string]interface{}
