package sdk

import (
	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/services/action"
	"github.com/splunk/splunk-cloud-sdk-go/services/catalog"
	"github.com/splunk/splunk-cloud-sdk-go/services/forwarders"
	"github.com/splunk/splunk-cloud-sdk-go/services/identity"
	"github.com/splunk/splunk-cloud-sdk-go/services/ingest"
	"github.com/splunk/splunk-cloud-sdk-go/services/kvstore"
	"github.com/splunk/splunk-cloud-sdk-go/services/search"
	"github.com/splunk/splunk-cloud-sdk-go/services/streams"
	"github.com/splunk/splunk-cloud-sdk-go/services/appreg"
)

// Client to communicate with Splunk Cloud service endpoints
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
	// AppRegService talks to the Splunk Cloud app registry service
	AppRegService *appreg.Service
}

// NewClient returns a Splunk Cloud client for communicating with any service
func NewClient(config *services.Config) (*Client, error) {
	client, err := services.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Client{
		BaseClient:        client,
		ActionService:     &action.Service{Client: client},
		CatalogService:    &catalog.Service{Client: client},
		IdentityService:   &identity.Service{Client: client},
		IngestService:     &ingest.Service{Client: client},
		KVStoreService:    &kvstore.Service{Client: client},
		SearchService:     &search.Service{Client: client},
		StreamsService:    &streams.Service{Client: client},
		ForwardersService: &forwarders.Service{Client: client},
	}, nil
}

// NewBatchEventsSenderWithMaxAllowedError is Deprecated: please use client.IngestService.NewBatchEventsSenderWithMaxAllowedError
func (c *Client) NewBatchEventsSenderWithMaxAllowedError(batchSize int, interval int64, maxErrorsAllowed int) (*ingest.BatchEventsSender, error) {
	return c.IngestService.NewBatchEventsSenderWithMaxAllowedError(batchSize, interval, 0, maxErrorsAllowed)
}

// NewBatchEventsSender is Deprecated: please use client.IngestService.NewBatchEventsSender
func (c *Client) NewBatchEventsSender(batchSize int, interval int64) (*ingest.BatchEventsSender, error) {
	return c.IngestService.NewBatchEventsSender(batchSize, interval, 0)
}
