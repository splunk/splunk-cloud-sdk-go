# kvstore
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/kvstore"


## Usage

#### type IndexDefinition

```go
type IndexDefinition struct {

	// The name of the index
	Name string `json:"name,omitempty"`

	// fields
	Fields []IndexFieldDefinition `json:"fields"`
}
```

IndexDefinition index field definition

#### type IndexDescription

```go
type IndexDescription struct {

	// The collection name
	Collection string `json:"collection,omitempty"`

	// fields
	Fields []IndexFieldDefinition `json:"fields"`

	// The name of the index
	Name string `json:"name,omitempty"`
}
```

IndexDescription index description

#### type IndexFieldDefinition

```go
type IndexFieldDefinition struct {

	// The sort direction for the indexed field
	Direction int64 `json:"direction"`

	// The name of the field to index
	Field string `json:"field"`
}
```

IndexFieldDefinition index field definition

#### type Key

```go
type Key struct {
	Key string `json:"_key"`
}
```

Key to identify a record in a collection

#### type LookupValue

```go
type LookupValue []interface{}
```

LookupValue Value tuple used for lookup

#### type PingOKBody

```go
type PingOKBody struct {

	// If database is not healthy, detailed error message
	ErrorMessage string `json:"errorMessage,omitempty"`

	// Database status
	// Enum: [healthy unhealthy unknown]
	Status PingOKBodyStatus `json:"status"`
}
```

PingOKBody ping ok body

#### type PingOKBodyStatus

```go
type PingOKBodyStatus string
```

PingOKBodyStatus used to force type expectation for KVStore Ping endpoint
response

```go
const (
	// PingOKBodyStatusHealthy captures enum value "healthy"
	PingOKBodyStatusHealthy PingOKBodyStatus = "healthy"

	// PingOKBodyStatusUnhealthy captures enum value "unhealthy"
	PingOKBodyStatusUnhealthy PingOKBodyStatus = "unhealthy"

	// PingOKBodyStatusUnknown captures enum value "unknown"
	PingOKBodyStatusUnknown PingOKBodyStatus = "unknown"
)
```

#### type Record

```go
type Record map[string]interface{}
```

Record is a JSON document entity contained in collections

#### type Service

```go
type Service services.BaseService
```

Service talks to kvstore service

#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new kvstore service client from the given Config

#### func (*Service) CreateIndex

```go
func (s *Service) CreateIndex(collectionName string, index IndexDefinition) (*IndexDescription, error)
```
CreateIndex posts a new index to be added to the collection.

#### func (*Service) DeleteIndex

```go
func (s *Service) DeleteIndex(collectionName string, indexName string) error
```
DeleteIndex deletes the specified index in a given collection

#### func (*Service) DeleteRecordByKey

```go
func (s *Service) DeleteRecordByKey(collectionName string, keyValue string) error
```
DeleteRecordByKey deletes a particular record present in a given collection
based on the key value provided by the user.

#### func (*Service) DeleteRecords

```go
func (s *Service) DeleteRecords(collectionName string, values url.Values) error
```
DeleteRecords deletes records present in a given collection based on the
provided query.

#### func (*Service) GetRecordByKey

```go
func (s *Service) GetRecordByKey(collectionName string, keyValue string) (Record, error)
```
GetRecordByKey queries a particular record present in a given collection based
on the key value provided by the user.

#### func (*Service) GetServiceHealthStatus

```go
func (s *Service) GetServiceHealthStatus() (*PingOKBody, error)
```
GetServiceHealthStatus returns Service Health Status

#### func (*Service) InsertRecord

```go
func (s *Service) InsertRecord(collectionName string, record Record) (map[string]string, error)
```
InsertRecord - Create a new record in the tenant's specified collection

#### func (*Service) InsertRecords

```go
func (s *Service) InsertRecords(collectionName string, records []Record) ([]string, error)
```
InsertRecords posts new records to the collection.

#### func (*Service) ListIndexes

```go
func (s *Service) ListIndexes(collectionName string) ([]IndexDefinition, error)
```
ListIndexes retrieves all the indexes in a given collection

#### func (*Service) ListRecords

```go
func (s *Service) ListRecords(collectionName string, filters map[string][]string) ([]map[string]interface{}, error)
```
ListRecords - List the records created for the tenant's specified collection
TODO: include count, offset and orderBy

#### func (*Service) PutRecord

```go
func (s *Service) PutRecord(collectionName, keyValue string, record Record) (map[string]string, bool, error)
```
PutRecord - Inserts or replaces a record in the tenant's specified collection
with the specified key

#### func (*Service) QueryRecords

```go
func (s *Service) QueryRecords(collectionName string, values url.Values) ([]Record, error)
```
QueryRecords queries records present in a given collection.

#### type Servicer

```go
type Servicer interface {
	// GetServiceHealthStatus returns Service Health Status
	GetServiceHealthStatus() (*PingOKBody, error)
	// CreateIndex posts a new index to be added to the collection.
	CreateIndex(collectionName string, index IndexDefinition) (*IndexDescription, error)
	// ListIndexes retrieves all the indexes in a given collection
	ListIndexes(collectionName string) ([]IndexDefinition, error)
	// DeleteIndex deletes the specified index in a given collection
	DeleteIndex(collectionName string, indexName string) error
	// InsertRecords posts new records to the collection.
	InsertRecords(collectionName string, records []Record) ([]string, error)
	// QueryRecords queries records present in a given collection.
	QueryRecords(collectionName string, values url.Values) ([]Record, error)
	// GetRecordByKey queries a particular record present in a given collection based on the key value provided by the user.
	GetRecordByKey(collectionName string, keyValue string) (Record, error)
	// DeleteRecords deletes records present in a given collection based on the provided query.
	DeleteRecords(collectionName string, values url.Values) error
	// DeleteRecordByKey deletes a particular record present in a given collection based on the key value provided by the user.
	DeleteRecordByKey(collectionName string, keyValue string) error
	// ListRecords - List the records created for the tenant's specified collection TODO: include count, offset and orderBy
	ListRecords(collectionName string, filters map[string][]string) ([]map[string]interface{}, error)
	// InsertRecord - Create a new record in the tenant's specified collection
	InsertRecord(collectionName string, record Record) (map[string]string, error)
	// PutRecord - Inserts or replaces a record in the tenant's specified collection with the specified key
	PutRecord(collectionName, keyValue string, record Record) (map[string]string, bool, error)
}
```

Servicer ...
