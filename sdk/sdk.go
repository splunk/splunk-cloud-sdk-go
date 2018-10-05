package sdk

import (
	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/services/action"
	"github.com/splunk/splunk-cloud-sdk-go/services/catalog"
	"github.com/splunk/splunk-cloud-sdk-go/services/identity"
	"github.com/splunk/splunk-cloud-sdk-go/services/ingest"
	"github.com/splunk/splunk-cloud-sdk-go/services/kvstore"
	"github.com/splunk/splunk-cloud-sdk-go/services/search"
	"github.com/splunk/splunk-cloud-sdk-go/services/streams"
)

// Client to communicate with Splunk Cloud service endpoints
type Client struct {
	*services.Client
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
	// StreamsService talks to SSC streams service
	StreamsService *streams.Service
}

// NewClient returns a Splunk Cloud client for communicating with any service
func NewClient(config *services.Config) (*Client, error) {
	client, err := services.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Client{
		Client:          client,
		ActionService:   action.NewService(client),
		CatalogService:  catalog.NewService(client),
		IdentityService: identity.NewService(client),
		IngestService:   ingest.NewService(client),
		KVStoreService:  kvstore.NewService(client),
		SearchService:   search.NewService(client),
		StreamsService:  streams.NewService(client),
	}, nil
}
