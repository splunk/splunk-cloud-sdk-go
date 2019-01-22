// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package ingest

import (
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBatchEventsSenderInitializationWithZeroBatchSizeAndZeroIntervalParameters(t *testing.T) {
	client, err := NewService(&services.Config{Token: "EXAMPLE_AUTHENTICATION_TOKEN"})
	require.Nil(t, err, "error creating ingest service client")
	_, err = client.NewBatchEventsSender(0, 0, 10)
	assert.EqualError(t, err, "batchSize cannot be 0")
}

func TestBatchEventsSenderInitializationWithZeroBatchSize(t *testing.T) {
	client, err := NewService(&services.Config{Token: "EXAMPLE_AUTHENTICATION_TOKEN"})
	require.Nil(t, err, "error creating ingest service client")
	_, err = client.NewBatchEventsSender(0, 1000, 10)
	assert.EqualError(t, err, "batchSize cannot be 0")
}

func TestBatchEventsSenderInitializationWithZeroInterval(t *testing.T) {
	client, err := NewService(&services.Config{Token: "EXAMPLE_AUTHENTICATION_TOKEN"})
	require.Nil(t, err, "error creating ingest service client")
	_, err = client.NewBatchEventsSender(5, 0, 10)
	assert.EqualError(t, err, "interval cannot be 0")
}

func TestBatchEventsSenderInitializationWithZeroPayloadSize(t *testing.T) {
	client, err := NewService(&services.Config{Token: "EXAMPLE_AUTHENTICATION_TOKEN"})
	require.Nil(t, err, "error creating ingest service client")
	collector, _ := client.NewBatchEventsSender(5, 10000, 0)
	assert.Equal(t, collector.PayLoadSize, 1000000)
}

func TestBatchEventsSenderInitializationWithNonZeroPayloadSize(t *testing.T) {
	client, err := NewService(&services.Config{Token: "EXAMPLE_AUTHENTICATION_TOKEN"})
	require.Nil(t, err, "error creating ingest service client")
	collector, _ := client.NewBatchEventsSender(5, 1000, 1000)
	assert.Equal(t, collector.PayLoadSize, 1000)
}
