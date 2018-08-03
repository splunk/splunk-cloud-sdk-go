package model

// CollectionStats collection stats
type CollectionStats struct {

	// Number of records in collection
	Count int64 `json:"count"`

	// Map of index name to index size in bytes
	IndexSizes interface{} `json:"indexSizes"`

	// Number of indexes on collection
	Nindexes int64 `json:"nindexes"`

	// Collection name
	Ns string `json:"ns"`

	// Size in bytes of collection, not including indexes
	Size int64 `json:"size"`

	// Total size of indexes
	TotalIndexSize int64 `json:"totalIndexSize"`
}

// CollectionDefinition collection definition
type CollectionDefinition struct {

	// The collection name
	// Max Length: 45
	// Min Length: 1
	Collection string `json:"collection"`
}

// CreateCollectionResponse create collection response
type CreateCollectionResponse struct {

	// name
	// Max Length: 45
	// Min Length: 1
	Name string `json:"name"`
}

// CreateNamespaceResponse create namespace response
type CreateNamespaceResponse struct {

	// name
	// Max Length: 45
	// Min Length: 1
	Name string `json:"name"`
}

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
	Status string `json:"status"`
}

const (

	// PingOKBodyStatusHealthy captures enum value "healthy"
	PingOKBodyStatusHealthy string = "healthy"

	// PingOKBodyStatusUnhealthy captures enum value "unhealthy"
	PingOKBodyStatusUnhealthy string = "unhealthy"

	// PingOKBodyStatusUnknown captures enum value "unknown"
	PingOKBodyStatusUnknown string = "unknown"
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
