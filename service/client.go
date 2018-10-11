// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

/*
Package service implements a service client which is used to communicate
with Search Service endpoints
*/
package service

import (
	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services"
)

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
