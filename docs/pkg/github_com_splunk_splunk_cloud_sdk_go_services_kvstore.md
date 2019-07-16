# kvstore
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/kvstore"


## Usage

#### type DeleteRecordsQueryParams

```go
type DeleteRecordsQueryParams struct {
	// Query : Query JSON expression.
	Query string `key:"query"`
}
```

DeleteRecordsQueryParams represents valid query parameters for the DeleteRecords
operation For convenience DeleteRecordsQueryParams can be formed in a single
statement, for example:

    `v := DeleteRecordsQueryParams{}.SetQuery(...)`

#### func (DeleteRecordsQueryParams) SetQuery

```go
func (q DeleteRecordsQueryParams) SetQuery(v string) DeleteRecordsQueryParams
```

#### type ErrorResponse

```go
type ErrorResponse struct {
	// Internal status code of the error.
	Code string `json:"code"`
	// Detailed error message.
	Message string `json:"message"`
}
```


#### type IndexDefinition

```go
type IndexDefinition struct {
	Fields []IndexFieldDefinition `json:"fields"`
	// The name of the index.
	Name string `json:"name"`
}
```


#### type IndexDescription

```go
type IndexDescription struct {
	// The collection name.
	Collection *string                `json:"collection,omitempty"`
	Fields     []IndexFieldDefinition `json:"fields,omitempty"`
	// The name of the index.
	Name *string `json:"name,omitempty"`
}
```


#### type IndexFieldDefinition

```go
type IndexFieldDefinition struct {
	// The sort direction for the indexed field.
	Direction int32 `json:"direction"`
	// The name of the field to index.
	Field string `json:"field"`
}
```


#### type Key

```go
type Key struct {
	// Key of the inserted document.
	Key string `json:"_key"`
}
```


#### type ListRecordsQueryParams

```go
type ListRecordsQueryParams struct {
	// Count : Maximum number of records to return.
	Count *int32 `key:"count"`
	// Fields : Comma-separated list of fields to include or exclude.
	Fields []string `key:"fields"`
	// Offset : Number of records to skip from the start.
	Offset *int32 `key:"offset"`
	// Orderby : Sort order. Format is &#x60;&lt;field&gt;:&lt;sort order&gt;&#x60;. Valid sort orders are 1 for ascending, -1 for descending.
	Orderby []string `key:"orderby"`
}
```

ListRecordsQueryParams represents valid query parameters for the ListRecords
operation For convenience ListRecordsQueryParams can be formed in a single
statement, for example:

    `v := ListRecordsQueryParams{}.SetCount(...).SetFields(...).SetOffset(...).SetOrderby(...)`

#### func (ListRecordsQueryParams) SetCount

```go
func (q ListRecordsQueryParams) SetCount(v int32) ListRecordsQueryParams
```

#### func (ListRecordsQueryParams) SetFields

```go
func (q ListRecordsQueryParams) SetFields(v []string) ListRecordsQueryParams
```

#### func (ListRecordsQueryParams) SetOffset

```go
func (q ListRecordsQueryParams) SetOffset(v int32) ListRecordsQueryParams
```

#### func (ListRecordsQueryParams) SetOrderby

```go
func (q ListRecordsQueryParams) SetOrderby(v []string) ListRecordsQueryParams
```

#### type PingResponse

```go
type PingResponse struct {
	// Database status.
	Status PingResponseStatus `json:"status"`
	// If database is not healthy, detailed error message.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}
```


#### type PingResponseStatus

```go
type PingResponseStatus string
```

PingResponseStatus : Database status.

```go
const (
	PingResponseStatusHealthy   PingResponseStatus = "healthy"
	PingResponseStatusUnhealthy PingResponseStatus = "unhealthy"
	PingResponseStatusUnknown   PingResponseStatus = "unknown"
)
```
List of PingResponseStatus

#### type QueryRecordsQueryParams

```go
type QueryRecordsQueryParams struct {
	// Count : Maximum number of records to return.
	Count *int32 `key:"count"`
	// Fields : Comma-separated list of fields to include or exclude.
	Fields []string `key:"fields"`
	// Offset : Number of records to skip from the start.
	Offset *int32 `key:"offset"`
	// Orderby : Sort order. Format is &#x60;&lt;field&gt;:&lt;sort order&gt;&#x60;. Valid sort orders are 1 for ascending, -1 for descending.
	Orderby []string `key:"orderby"`
	// Query : Query JSON expression.
	Query string `key:"query"`
}
```

QueryRecordsQueryParams represents valid query parameters for the QueryRecords
operation For convenience QueryRecordsQueryParams can be formed in a single
statement, for example:

    `v := QueryRecordsQueryParams{}.SetCount(...).SetFields(...).SetOffset(...).SetOrderby(...).SetQuery(...)`

#### func (QueryRecordsQueryParams) SetCount

```go
func (q QueryRecordsQueryParams) SetCount(v int32) QueryRecordsQueryParams
```

#### func (QueryRecordsQueryParams) SetFields

```go
func (q QueryRecordsQueryParams) SetFields(v []string) QueryRecordsQueryParams
```

#### func (QueryRecordsQueryParams) SetOffset

```go
func (q QueryRecordsQueryParams) SetOffset(v int32) QueryRecordsQueryParams
```

#### func (QueryRecordsQueryParams) SetOrderby

```go
func (q QueryRecordsQueryParams) SetOrderby(v []string) QueryRecordsQueryParams
```

#### func (QueryRecordsQueryParams) SetQuery

```go
func (q QueryRecordsQueryParams) SetQuery(v string) QueryRecordsQueryParams
```

#### type Service

```go
type Service services.BaseService
```


#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new kvstore service client from the given Config

#### func (*Service) CreateIndex

```go
func (s *Service) CreateIndex(collection string, indexDefinition IndexDefinition, resp ...*http.Response) (*IndexDescription, error)
```
CreateIndex - Creates an index on a collection. Parameters:

    collection: The name of the collection.
    indexDefinition
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteIndex

```go
func (s *Service) DeleteIndex(collection string, index string, resp ...*http.Response) error
```
DeleteIndex - Removes an index from a collection. Parameters:

    collection: The name of the collection.
    index: The name of the index.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteRecordByKey

```go
func (s *Service) DeleteRecordByKey(collection string, key string, resp ...*http.Response) error
```
DeleteRecordByKey - Deletes a record with a given key. Parameters:

    collection: The name of the collection.
    key: The key of the record.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteRecords

```go
func (s *Service) DeleteRecords(collection string, query *DeleteRecordsQueryParams, resp ...*http.Response) error
```
DeleteRecords - Removes records in a collection that match the query.
Parameters:

    collection: The name of the collection.
    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetRecordByKey

```go
func (s *Service) GetRecordByKey(collection string, key string, resp ...*http.Response) (*map[string]interface{}, error)
```
GetRecordByKey - Returns a record with a given key. Parameters:

    collection: The name of the collection.
    key: The key of the record.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) InsertRecord

```go
func (s *Service) InsertRecord(collection string, body map[string]interface{}, resp ...*http.Response) (*Key, error)
```
InsertRecord - Inserts a record into a collection. Parameters:

    collection: The name of the collection.
    body: Record to add to the collection, formatted as a JSON object.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) InsertRecords

```go
func (s *Service) InsertRecords(collection string, requestBody []map[string]interface{}, resp ...*http.Response) ([]string, error)
```
InsertRecords - Inserts multiple records in a single request. Parameters:

    collection: The name of the collection.
    requestBody: Array of records to insert.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListIndexes

```go
func (s *Service) ListIndexes(collection string, resp ...*http.Response) ([]IndexDefinition, error)
```
ListIndexes - Returns a list of all indexes on a collection. Parameters:

    collection: The name of the collection.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListRecords

```go
func (s *Service) ListRecords(collection string, query *ListRecordsQueryParams, resp ...*http.Response) ([]map[string]interface{}, error)
```
ListRecords - Returns a list of records in a collection with basic filtering,
sorting, pagination and field projection. Use key-value query parameters to
filter fields. Fields are implicitly ANDed and values for the same field are
implicitly ORed. Parameters:

    collection: The name of the collection.
    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) Ping

```go
func (s *Service) Ping(resp ...*http.Response) (*PingResponse, error)
```
Ping - Returns the health status from the database. Parameters:

    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) PutRecord

```go
func (s *Service) PutRecord(collection string, key string, body map[string]interface{}, resp ...*http.Response) (*Key, error)
```
PutRecord - Updates the record with a given key, either by inserting or
replacing the record. Parameters:

    collection: The name of the collection.
    key: The key of the record.
    body: Record to add to the collection, formatted as a JSON object.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) QueryRecords

```go
func (s *Service) QueryRecords(collection string, query *QueryRecordsQueryParams, resp ...*http.Response) ([]map[string]interface{}, error)
```
QueryRecords - Returns a list of query records in a collection. Parameters:

    collection: The name of the collection.
    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### type Servicer

```go
type Servicer interface {
	/*
		CreateIndex - Creates an index on a collection.
		Parameters:
			collection: The name of the collection.
			indexDefinition
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateIndex(collection string, indexDefinition IndexDefinition, resp ...*http.Response) (*IndexDescription, error)
	/*
		DeleteIndex - Removes an index from a collection.
		Parameters:
			collection: The name of the collection.
			index: The name of the index.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteIndex(collection string, index string, resp ...*http.Response) error
	/*
		DeleteRecordByKey - Deletes a record with a given key.
		Parameters:
			collection: The name of the collection.
			key: The key of the record.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteRecordByKey(collection string, key string, resp ...*http.Response) error
	/*
		DeleteRecords - Removes records in a collection that match the query.
		Parameters:
			collection: The name of the collection.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteRecords(collection string, query *DeleteRecordsQueryParams, resp ...*http.Response) error
	/*
		GetRecordByKey - Returns a record with a given key.
		Parameters:
			collection: The name of the collection.
			key: The key of the record.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetRecordByKey(collection string, key string, resp ...*http.Response) (*map[string]interface{}, error)
	/*
		InsertRecord - Inserts a record into a collection.
		Parameters:
			collection: The name of the collection.
			body: Record to add to the collection, formatted as a JSON object.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	InsertRecord(collection string, body map[string]interface{}, resp ...*http.Response) (*Key, error)
	/*
		InsertRecords - Inserts multiple records in a single request.
		Parameters:
			collection: The name of the collection.
			requestBody: Array of records to insert.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	InsertRecords(collection string, requestBody []map[string]interface{}, resp ...*http.Response) ([]string, error)
	/*
		ListIndexes - Returns a list of all indexes on a collection.
		Parameters:
			collection: The name of the collection.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListIndexes(collection string, resp ...*http.Response) ([]IndexDefinition, error)
	/*
		ListRecords - Returns a list of records in a collection with basic filtering, sorting, pagination and field projection.
		Use key-value query parameters to filter fields. Fields are implicitly ANDed and values for the same field are implicitly ORed.
		Parameters:
			collection: The name of the collection.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListRecords(collection string, query *ListRecordsQueryParams, resp ...*http.Response) ([]map[string]interface{}, error)
	/*
		Ping - Returns the health status from the database.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	Ping(resp ...*http.Response) (*PingResponse, error)
	/*
		PutRecord - Updates the record with a given key, either by inserting or replacing the record.
		Parameters:
			collection: The name of the collection.
			key: The key of the record.
			body: Record to add to the collection, formatted as a JSON object.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	PutRecord(collection string, key string, body map[string]interface{}, resp ...*http.Response) (*Key, error)
	/*
		QueryRecords - Returns a list of query records in a collection.
		Parameters:
			collection: The name of the collection.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	QueryRecords(collection string, query *QueryRecordsQueryParams, resp ...*http.Response) ([]map[string]interface{}, error)
}
```

Servicer represents the interface for implementing all endpoints for this
service
