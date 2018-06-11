package model

// CollectionStats collection stats
type CollectionStats struct {

	// Number of records in collection
	Count *int64 `json:"count"`

	// Map of index name to index size in bytes
	IndexSizes interface{} `json:"indexSizes"`

	// Number of indexes on collection
	Nindexes *int64 `json:"nindexes"`

	// Collection name
	Ns *string `json:"ns"`

	// Size in bytes of collection, not including indexes
	Size *int64 `json:"size"`

	// Total size of indexes
	TotalIndexSize *int64 `json:"totalIndexSize"`
}

// PingOKBody ping ok body
type PingOKBody struct {

	// If database is not healthy, detailed error message
	ErrorMessage string `json:"errorMessage,omitempty"`

	// Database status
	// Enum: [healthy unhealthy unknown]
	Status *string `json:"status"`
}

const (

	// PingOKBodyStatusHealthy captures enum value "healthy"
	PingOKBodyStatusHealthy string = "healthy"

	// PingOKBodyStatusUnhealthy captures enum value "unhealthy"
	PingOKBodyStatusUnhealthy string = "unhealthy"

	// PingOKBodyStatusUnknown captures enum value "unknown"
	PingOKBodyStatusUnknown string = "unknown"
)

// LookupValue Value tuple used for lookup
type LookupValue []interface{}
