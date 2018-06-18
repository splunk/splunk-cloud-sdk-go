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
	Direction *int64 `json:"direction"`

	// The name of the field to index
	Field *string `json:"field"`
}

// IndexDescription index description
type IndexDescription struct {

	// The collection name
	Collection string `json:"collection,omitempty"`

	// fields
	Fields []*IndexFieldDefinition `json:"fields"`

	// The name of the index
	Name string `json:"name,omitempty"`

	// The namespace containing the collection
	Namespace string `json:"namespace,omitempty"`
}

// CollectionDescription collection description
type CollectionDescription struct {

	// The list of indexes on this collection
	Indexes []*IndexDescription `json:"indexes"`

	// The collection name
	Name string `json:"name,omitempty"`

	// The namespace containing the collection
	Namespace string `json:"namespace,omitempty"`
}

// NamespaceDescription namespace description
type NamespaceDescription struct {

	// The list of collections
	Collections []*CollectionDescription `json:"collections"`

	// The name of the namespace
	Name string `json:"name,omitempty"`
}

// TenantDescription tenant description
type TenantDescription struct {

	// The name of the tenant
	Name string `json:"name,omitempty"`

	// The list of namespaces
	Namespaces []*NamespaceDescription `json:"namespaces"`
}

// LookupValue Value tuple used for lookup
type LookupValue []interface{}

// Key to identify a record in a collection
type Key struct {
	Key string `json:"_key"`
}

// Record is a JSON document entity contained in collections
type Record map[string]interface{}