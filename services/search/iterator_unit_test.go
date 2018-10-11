// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

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
		func(count, offset int) (*Results, error) {
			return &Results{}, nil
		})
	assert.Equal(t, iterator.batch, 10)
	assert.Equal(t, iterator.offset, 0)
	assert.Equal(t, iterator.max, 100)
	assert.Equal(t, iterator.start, 0)
	assert.Equal(t, iterator.isClosed, false)

}

func TestClose(t *testing.T) {
	iterator := NewIterator(10, 0, 100,
		func(count, offset int) (*Results, error) {
			return &Results{}, nil
		})
	assert.Equal(t, iterator.isClosed, false)
	iterator.Close()
	assert.Equal(t, iterator.isClosed, true)
	_, err := iterator.Value()
	assert.Equal(t, "failed to retrieve values on a closed iterator", err.Error())
}

func TestErr(t *testing.T) {
	iterator := NewIterator(10, 0, 100,
		func(count, offset int) (*Results, error) {
			return &Results{}, nil
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

	var searchResults *Results

	iterator := NewIterator(batch, 0, max,
		func(count, offset int) (*Results, error) {
			if offset < count {
				return &Results{Results: results[offset:count]}, nil
			} else if offset+count < max {
				return &Results{Results: results[offset : offset+count]}, nil
				// } else if offset-count <= max {
				// 	return &Results{Results: results[offset-count : max]}, nil
			} else {
				return &Results{Results: results[offset:max]}, nil
			}
		})
	defer iterator.Close()
	// first next
	hasNext = iterator.Next()
	assert.True(t, hasNext)
	searchResults, _ = iterator.Value()
	assert.Equal(t, &Results{Results: results[0:2]}, searchResults)
	// second next
	hasNext = iterator.Next()
	assert.True(t, hasNext)
	searchResults, _ = iterator.Value()
	assert.Equal(t, &Results{Results: results[2:4]}, searchResults)
	// third next
	hasNext = iterator.Next()
	assert.True(t, hasNext)
	searchResults, _ = iterator.Value()
	assert.Equal(t, &Results{Results: results[4:5]}, searchResults)
	// fourth next: break out
	hasNext = iterator.Next()
	assert.False(t, hasNext)
	value4, _ := iterator.Value()
	var emptyResults *Results
	assert.ObjectsAreEqualValues(value4, emptyResults)
}

func TestNextOnClose(t *testing.T) {
	iterator := NewIterator(10, 0, 100,
		func(count, offset int) (*Results, error) {
			return &Results{}, nil
		})
	iterator.Close()
	assert.False(t, iterator.Next())
}

func TestNextOnErr(t *testing.T) {
	iterator := NewIterator(10, 0, 100,
		func(count, offset int) (*Results, error) {
			return &Results{}, nil
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

	var searchResults *Results

	iterator := NewIterator(batch, 0, max,
		func(count, offset int) (*Results, error) {
			if count == 0 {
				return &Results{Results: results}, nil
			}
			return nil, nil
		})
	assert.False(t, iterator.Next())
	searchResults, _ = iterator.Value()
	assert.Equal(t, &Results{Results: results}, searchResults)
}

func TestNextOnFnErr(t *testing.T) {
	batch := 0
	max := 5
	var searchResults *Results
	iterator := NewIterator(batch, 0, max,
		func(count, offset int) (*Results, error) {
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

	var searchResults *Results

	iterator := NewIterator(batch, 0, max,
		func(count, offset int) (*Results, error) {
			return &Results{Results: results}, nil
		})
	assert.False(t, iterator.Next())
	searchResultsActual, _ := iterator.Value()
	assert.Equal(t, searchResults, searchResultsActual)
}
