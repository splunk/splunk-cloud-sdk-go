// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package model

// Event defines raw event send to event endpoint
type Event struct {
	// Specifies a JSON object that contains explicit custom fields to be defined at index time.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// JSON object for the event.
	Body interface{} `json:"body"`
	// Epoch time in milliseconds.
	Timestamp int64 `json:"timestamp,omitempty"`
	// Optional nanoseconds part of the timestamp.
	Nanos int32 `json:"nanos,omitempty"`
	// The source value to assign to the event data. For example, if you are sending data from an app that you are developing,
	// set this key to the name of the app.
	Source string `json:"source,omitempty"`
	// The sourcetype value assigned to the event data.
	Sourcetype string `json:"sourcetype,omitempty"`
	// The host value assigned to the event data. Typically, this is the hostname of the client from which you are sending data.
	Host string `json:"host,omitempty"`
	// An optional ID that uniquely identifies the metric data. It is used to deduplicate the data if same data is set multiple times.
	// If ID is not specified, it will be assigned by the system.
	ID string `json:"id,omitempty"`
}

// MetricEvent define event send to metric endpoint
type MetricEvent struct {
	// Specify multiple related metrics e.g. Memory, CPU etc.
	Body []Metric `json:"body"`
	// Epoch time in milliseconds.
	Timestamp int64 `json:"timestamp,omitempty"`
	// Optional nanoseconds part of the timestamp.
	Nanos int32 `json:"nanos,omitempty"`
	// The source value to assign to the event data. For example, if you're sending data from an app you're developing,
	// you could set this key to the name of the app.
	Source string `json:"source,omitempty"`
	// The sourcetype value to assign to the event data.
	Sourcetype string `json:"sourcetype,omitempty"`
	// The host value to assign to the event data. This is typically the hostname of the client from which you're sending data.
	Host string `json:"host,omitempty"`
	// Optional ID uniquely identifies the metric data. It is used to deduplicate the data if same data is set multiple times.
	// If ID is not specified, it will be assigned by the system.
	ID string `json:"id,omitempty"`
	// Default attributes for the metric data.
	Attributes MetricAttribute `json:"attributes,omitempty"`
}

// Metric defines individual metric data.
type Metric struct {
	// Name of the metric e.g. CPU, Memory etc.
	Name string `json:"name"`
	// Value of the metric.
	Value float64 `json:"value"`
	// Dimensions allow metrics to be classified e.g. {"Server":"nginx", "Region":"us-west-1", ...}
	Dimensions map[string]string `json:"dimensions,omitempty"`
	// Type of metric. Default is g for gauge.
	Type string `json:"type,omitempty"`
	// Unit of the metric e.g. percent, megabytes, seconds etc.
	Unit string `json:"unit,omitempty"`
}

// MetricAttribute defines default attributes for the metric.
type MetricAttribute struct {
	// Optional. If set, individual Metrics will inherit these dimensions and can override any/all of them.
	DefaultDimensions map[string]string `json:"defaultDimensions,omitempty"`
	// Optional. If set, individual Metrics will inherit this type and can optionally override.
	DefaultType string `json:"defaultType,omitempty"`
	// Optional. If set, individual Metrics will inherit this unit and can optionally override.
	DefaultUnit string `json:"defaultUnit,omitempty"`
}
