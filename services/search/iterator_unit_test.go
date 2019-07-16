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

package search

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewIterator(t *testing.T) {
	iterator := NewIterator(10, 0, 100,
		func(count, offset int) (*ListSearchResultsResponse, error) {
			return &ListSearchResultsResponse{}, nil
		})
	assert.Equal(t, iterator.batch, 10)
	assert.Equal(t, iterator.offset, 0)
	assert.Equal(t, iterator.max, 100)
	assert.Equal(t, iterator.start, 0)
	assert.Equal(t, iterator.isClosed, false)

}

func TestClose(t *testing.T) {
	iterator := NewIterator(10, 0, 100,
		func(count, offset int) (*ListSearchResultsResponse, error) {
			return &ListSearchResultsResponse{}, nil
		})
	assert.Equal(t, iterator.isClosed, false)
	iterator.Close()
	assert.Equal(t, iterator.isClosed, true)
	_, err := iterator.Value()
	assert.Equal(t, "failed to retrieve values on a closed iterator", err.Error())
}

func TestErr(t *testing.T) {
	iterator := NewIterator(10, 0, 100,
		func(count, offset int) (*ListSearchResultsResponse, error) {
			return &ListSearchResultsResponse{}, nil
		})
	iterator.err = errors.New("test error")
	assert.Error(t, iterator.Err(), "test error")
}

func TestNextSuccess(t *testing.T) {
	var hasNext bool
	batch := 2
	max := 5

	byt := []byte(`[{"Raw":"test"},{"Raw":"test2"},{"Raw":"test3"},{"Raw":"test4"},{"Raw":"test5"}]`)
	var results []map[string]interface{}

	if err := json.Unmarshal(byt, &results); err != nil {
		require.Nil(t, err)
	}

	var searchResults *ListSearchResultsResponse

	iterator := NewIterator(batch, 0, max,
		func(count, offset int) (*ListSearchResultsResponse, error) {
			if offset < count {
				return &ListSearchResultsResponse{Results: results[offset:count]}, nil
			} else if offset+count < max {
				return &ListSearchResultsResponse{Results: results[offset : offset+count]}, nil
				// } else if offset-count <= max {
				// 	return &Results{Results: results[offset-count : max]}, nil
			} else {
				return &ListSearchResultsResponse{Results: results[offset:max]}, nil
			}
		})
	defer iterator.Close()
	// first next
	hasNext = iterator.Next()
	assert.True(t, hasNext)
	searchResults, _ = iterator.Value()
	assert.Equal(t, &ListSearchResultsResponse{Results: results[0:2]}, searchResults)
	// second next
	hasNext = iterator.Next()
	assert.True(t, hasNext)
	searchResults, _ = iterator.Value()
	assert.Equal(t, &ListSearchResultsResponse{Results: results[2:4]}, searchResults)
	// third next
	hasNext = iterator.Next()
	assert.True(t, hasNext)
	searchResults, _ = iterator.Value()
	assert.Equal(t, &ListSearchResultsResponse{Results: results[4:5]}, searchResults)
	// fourth next: break out
	hasNext = iterator.Next()
	assert.False(t, hasNext)
	value4, _ := iterator.Value()
	var emptyResults *ListSearchResultsResponse
	assert.ObjectsAreEqualValues(value4, emptyResults)
}

func TestNextOnClose(t *testing.T) {
	iterator := NewIterator(10, 0, 100,
		func(count, offset int) (*ListSearchResultsResponse, error) {
			return &ListSearchResultsResponse{}, nil
		})
	iterator.Close()
	assert.False(t, iterator.Next())
}

func TestNextOnErr(t *testing.T) {
	iterator := NewIterator(10, 0, 100,
		func(count, offset int) (*ListSearchResultsResponse, error) {
			return &ListSearchResultsResponse{}, nil
		})
	defer iterator.Close()
	iterator.err = errors.New("test error")
	assert.False(t, iterator.Next())
}

func TestNextOnZeroBatch(t *testing.T) {
	batch := 0
	max := 5
	var results []map[string]interface{}
	byt := []byte(`[{"Raw":"test"},{"Raw":"test2"},{"Raw":"test3"},{"Raw":"test4"},{"Raw":"test5"}]`)
	if err := json.Unmarshal(byt, &results); err != nil {
		require.Nil(t, err)
	}

	var searchResults *ListSearchResultsResponse

	iterator := NewIterator(batch, 0, max,
		func(count, offset int) (*ListSearchResultsResponse, error) {
			if count == 0 {
				return &ListSearchResultsResponse{Results: results}, nil
			}
			return nil, nil
		})
	assert.False(t, iterator.Next())
	searchResults, _ = iterator.Value()
	assert.Equal(t, &ListSearchResultsResponse{Results: results}, searchResults)
}

func TestNextOnFnErr(t *testing.T) {
	batch := 0
	max := 5
	var searchResults *ListSearchResultsResponse
	iterator := NewIterator(batch, 0, max,
		func(count, offset int) (*ListSearchResultsResponse, error) {
			return nil, errors.New("error")
		})
	assert.False(t, iterator.Next())
	searchResultsActual, _ := iterator.Value()
	assert.Equal(t, searchResults, searchResultsActual)
	assert.Error(t, iterator.Err(), "error")
}

func TestNextNoMoreResults(t *testing.T) {
	batch := 0
	max := 5
	var results []map[string]interface{}

	var searchResults *ListSearchResultsResponse

	iterator := NewIterator(batch, 0, max,
		func(count, offset int) (*ListSearchResultsResponse, error) {
			return &ListSearchResultsResponse{Results: results}, nil
		})
	assert.False(t, iterator.Next())
	searchResultsActual, _ := iterator.Value()
	assert.Equal(t, searchResults, searchResultsActual)
}
