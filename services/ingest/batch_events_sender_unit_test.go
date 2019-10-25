/*
 * Copyright 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package ingest

import (
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/v2/services"
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
	assert.Equal(t, collector.PayLoadBytes, 1040000)
}

func TestBatchEventsSenderInitializationWithNonZeroPayloadSize(t *testing.T) {
	client, err := NewService(&services.Config{Token: "EXAMPLE_AUTHENTICATION_TOKEN"})
	require.Nil(t, err, "error creating ingest service client")
	collector, _ := client.NewBatchEventsSender(5, 1000, 1000)
	assert.Equal(t, collector.PayLoadBytes, 1000)
}

func TestBatchEventsSenderState(t *testing.T) {
	client, err := NewService(&services.Config{Token: "EXAMPLE_AUTHENTICATION_TOKEN"})
	require.Nil(t, err, "error creating ingest service client")

	collector, err := client.NewBatchEventsSender(5, 1000, 20)
	assert.Nil(t, err)

	// Validate initial values
	assert.Equal(t, 0, len(collector.EventsQueue))
	assert.Equal(t, 5, cap(collector.EventsQueue))
	assert.Equal(t, 0, len(collector.EventsChan))
	assert.Equal(t, 5, cap(collector.EventsChan))
	assert.Equal(t, 0, len(collector.QuitChan))
	assert.Equal(t, 1, cap(collector.QuitChan))
	assert.Equal(t, 5, collector.BatchSize)
	assert.Equal(t, 20, collector.PayLoadBytes)
}
