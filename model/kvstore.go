// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package model

import (
	"github.com/splunk/splunk-cloud-sdk-go/services/kvstore"
)

// Error is Deprecated: please use services/kvstore.Error
type Error = kvstore.Error

// AuthError is Deprecated: please use services/kvstore.AuthError
type AuthError = kvstore.AuthError

// PingOKBody is Deprecated: please use services/kvstore.PingOKBody
type PingOKBody = kvstore.PingOKBody

// PingOKBodyStatus is Deprecated: please use services/kvstore.PingOKBodyStatus
type PingOKBodyStatus = kvstore.PingOKBodyStatus

const (
	// PingOKBodyStatusHealthy is Deprecated: please use services/kvstore.PingOKBodyStatusHealthy
	PingOKBodyStatusHealthy PingOKBodyStatus = kvstore.PingOKBodyStatusHealthy

	// PingOKBodyStatusUnhealthy is Deprecated: please use services/kvstore.PingOKBodyStatusUnhealthy
	PingOKBodyStatusUnhealthy PingOKBodyStatus = kvstore.PingOKBodyStatusUnhealthy

	// PingOKBodyStatusUnknown is Deprecated: please use services/kvstore.PingOKBodyStatusUnknown
	PingOKBodyStatusUnknown PingOKBodyStatus = kvstore.PingOKBodyStatusUnknown
)

// IndexFieldDefinition is Deprecated: please use services/kvstore.IndexFieldDefinition
type IndexFieldDefinition = kvstore.IndexFieldDefinition

// IndexDefinition is Deprecated: please use services/kvstore.IndexDefinition
type IndexDefinition = kvstore.IndexDefinition

// IndexDescription is Deprecated: please use services/kvstore.IndexDescription
type IndexDescription = kvstore.IndexDescription

// LookupValue is Deprecated: please use services/kvstore.LookupValue
type LookupValue = kvstore.LookupValue

// Key is Deprecated: please use services/kvstore.Key
type Key = kvstore.Key

// Record is Deprecated: please use services/kvstore.Record
type Record = kvstore.Record
