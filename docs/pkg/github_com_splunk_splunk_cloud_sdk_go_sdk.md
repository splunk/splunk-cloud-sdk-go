# sdk
--
    import "github.com/splunk/splunk-cloud-sdk-go/sdk"


## Usage

#### type Client

```go
type Client struct {
	*services.BaseClient
	// ActionService talks to Splunk Cloud action service
	ActionService *action.Service
	// CatalogService talks to the Splunk Cloud catalog service
	CatalogService *catalog.Service
	// IdentityService talks to Splunk Cloud IAC service
	IdentityService *identity.Service
	// IngestService talks to the Splunk Cloud ingest service
	IngestService *ingest.Service
	// KVStoreService talks to Splunk Cloud kvstore service
	KVStoreService *kvstore.Service
	// SearchService talks to the Splunk Cloud search service
	SearchService *search.Service
	// StreamsService talks to the Splunk Cloud streams service
	StreamsService *streams.Service
	// ForwardersService talks to the Splunk Cloud forwarders service
	ForwardersService *forwarders.Service
}
```

Client to communicate with Splunk Cloud service endpoints

#### func  NewClient

```go
func NewClient(config *services.Config) (*Client, error)
```
NewClient returns a Splunk Cloud client for communicating with any service

#### func (*Client) NewBatchEventsSender

```go
func (c *Client) NewBatchEventsSender(batchSize int, interval int64) (*ingest.BatchEventsSender, error)
```
NewBatchEventsSender is Deprecated: please use
client.IngestService.NewBatchEventsSender

#### func (*Client) NewBatchEventsSenderWithMaxAllowedError

```go
func (c *Client) NewBatchEventsSenderWithMaxAllowedError(batchSize int, interval int64, maxErrorsAllowed int) (*ingest.BatchEventsSender, error)
```
NewBatchEventsSenderWithMaxAllowedError is Deprecated: please use
client.IngestService.NewBatchEventsSenderWithMaxAllowedError
