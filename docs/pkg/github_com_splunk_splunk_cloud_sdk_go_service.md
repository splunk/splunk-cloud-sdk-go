# service
--
    import "github.com/splunk/splunk-cloud-sdk-go/service"

Package service contains Splunk Cloud client, services, and helpers.

Deprecated: v0.6.1 - these client, services and related helpers have been moved
to their respective sdk/sdk.go, services/*.gom or services/<service>/service.go
files. See below for details for each client/service/helper.

## Usage

#### type ActionService

```go
type ActionService = action.Service
```

ActionService is Deprecated: please use services/action.Service

#### type AuthnResponseHandler

```go
type AuthnResponseHandler = services.AuthnResponseHandler
```

AuthnResponseHandler is Deprecated: please use services.AuthnResponseHandler

#### type BatchEventsSender

```go
type BatchEventsSender = ingest.BatchEventsSender
```

BatchEventsSender is Deprecated: please use services/ingest.BatchEventsSender

#### type CatalogService

```go
type CatalogService = catalog.Service
```

CatalogService is Deprecated: please use services/catalog.Service

#### type Client

```go
type Client = sdk.Client
```

Client is Deprecated: please use sdk.Client

#### func  NewClient

```go
func NewClient(config *Config) (*Client, error)
```
NewClient is Deprecated: please use sdk.NewClient

#### type Config

```go
type Config = services.Config
```

Config is Deprecated: please use services.Config

#### type IdentityService

```go
type IdentityService = identity.Service
```

IdentityService is Deprecated: please use services/identity.Service

#### type IngestService

```go
type IngestService = ingest.Service
```

IngestService is Deprecated: please use services/ingest.Service

#### type KVStoreService

```go
type KVStoreService = kvstore.Service
```

KVStoreService is Deprecated: please use services/kvstore.Service

#### type Request

```go
type Request = services.Request
```

Request is Deprecated: please use services.Request

#### type RequestParams

```go
type RequestParams = services.RequestParams
```

RequestParams is Deprecated: please use services.RequestParams

#### type ResponseHandler

```go
type ResponseHandler = services.ResponseHandler
```

ResponseHandler is Deprecated: please use services.ResponseHandler

#### type SearchIterator

```go
type SearchIterator = search.Iterator
```

SearchIterator is Deprecated: please use services/search.Iterator

#### func  NewSearchIterator

```go
func NewSearchIterator(batch, offset, max int, fn search.QueryFunc) *SearchIterator
```
NewSearchIterator is Deprecated: please use services/search.NewIterator

#### type SearchService

```go
type SearchService = search.Service
```

SearchService is Deprecated: please use services/search.Service

#### type StreamsService

```go
type StreamsService = streams.Service
```

StreamsService is Deprecated: please use services/streams.Service

#### type UserErrHandler

```go
type UserErrHandler = ingest.UserErrHandler
```

UserErrHandler is Deprecated: please use services/ingest.UserErrHandler
