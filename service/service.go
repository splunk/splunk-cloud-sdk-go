// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

// Package service contains Splunk Cloud client, services, and helpers.
//
// Deprecated: v0.6.1 - these client, services and related helpers have been moved to their respective sdk/sdk.go, services/*.gom or services/<service>/service.go files. See below for details for each client/service/helper.
package service

import (
	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/services/action"
	"github.com/splunk/splunk-cloud-sdk-go/services/catalog"
	"github.com/splunk/splunk-cloud-sdk-go/services/identity"
	"github.com/splunk/splunk-cloud-sdk-go/services/ingest"
	"github.com/splunk/splunk-cloud-sdk-go/services/kvstore"
	"github.com/splunk/splunk-cloud-sdk-go/services/search"
	"github.com/splunk/splunk-cloud-sdk-go/services/streams"
)

//
// Deprecated: Client
//

// Client is Deprecated: please use sdk.Client
type Client = sdk.Client

// Request is Deprecated: please use services.Request
type Request = services.Request

// Config is Deprecated: please use services.Config
type Config = services.Config

// RequestParams is Deprecated: please use services.RequestParams
type RequestParams = services.RequestParams

// NewClient is Deprecated: please use sdk.NewClient
func NewClient(config *Config) (*Client, error) {
	return sdk.NewClient(config)
}

//
// Deprecated: Services
//

// ActionService is Deprecated: please use services/action.Service
type ActionService = action.Service

// CatalogService is Deprecated: please use services/catalog.Service
type CatalogService = catalog.Service

// IdentityService is Deprecated: please use services/identity.Service
type IdentityService = identity.Service

// IngestService is Deprecated: please use services/ingest.Service
type IngestService = ingest.Service

// KVStoreService is Deprecated: please use services/kvstore.Service
type KVStoreService = kvstore.Service

// SearchService is Deprecated: please use services/search.Service
type SearchService = search.Service

// StreamsService is Deprecated: please use services/streams.Service
type StreamsService = streams.Service

//
// Deprecated: Service helpers
//

// UserErrHandler is Deprecated: please use services/ingest.UserErrHandler
type UserErrHandler = ingest.UserErrHandler

// BatchEventsSender is Deprecated: please use services/ingest.BatchEventsSender
type BatchEventsSender = ingest.BatchEventsSender

// ResponseHandler is Deprecated: please use services.ResponseHandler
type ResponseHandler = services.ResponseHandler

// AuthnResponseHandler is Deprecated: please use services.AuthnResponseHandler
type AuthnResponseHandler = services.AuthnResponseHandler

// SearchIterator is Deprecated: please use services/search.Iterator
type SearchIterator = search.Iterator

// NewSearchIterator is Deprecated: please use services/search.NewIterator
func NewSearchIterator(batch, offset, max int, fn search.QueryFunc) *SearchIterator {
	return search.NewIterator(batch, offset, max, fn)
}
