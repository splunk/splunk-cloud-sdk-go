package service

import (
	"errors"
	"fmt"
	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSearchIterator(t *testing.T) {
	iterator := NewSearchIterator(10, 0, 100,
		func(count, offset int) (*model.SearchResults, error) {
			return &model.SearchResults{}, nil
		})
	assert.Equal(t, iterator.batch, 10)
	assert.Equal(t, iterator.offset, 0)
	assert.Equal(t, iterator.max, 100)
	assert.Equal(t, iterator.start, 0)
	assert.Equal(t, iterator.isClosed, false)

}

func TestClose(t *testing.T) {
	iterator := NewSearchIterator(10, 0, 100,
		func(count, offset int) (*model.SearchResults, error) {
			return &model.SearchResults{}, nil
		})
	assert.Equal(t, iterator.isClosed, false)
	iterator.Close()
	assert.Equal(t, iterator.isClosed, true)
	_, err := iterator.Value()
	assert.Equal(t, "failed to retrieve values on a closed iterator", err.Error())
}

func TestErr(t *testing.T) {
	iterator := NewSearchIterator(10, 0, 100,
		func(count, offset int) (*model.SearchResults, error) {
			return &model.SearchResults{}, nil
		})
	iterator.err = errors.New("test error")
	assert.Error(t, iterator.Err(), "test error")
}

func TestNextSuccess(t *testing.T) {
	var hasNext bool
	batch := 2
	max := 5
	results := make([]*model.Result, max)
	var searchResults *model.SearchResults
	for i := 0; i < max; i++ {
		results = append(results, &model.Result{Raw: fmt.Sprintf("test data %v", i)})
	}
	iterator := NewSearchIterator(batch, 0, max,
		func(count, offset int) (*model.SearchResults, error) {
			if offset < count {
				return &model.SearchResults{Results: results[offset:count]}, nil
			} else if offset+count < max {
				return &model.SearchResults{Results: results[offset : offset+count]}, nil
				// } else if offset-count <= max {
				// 	return &model.SearchResults{Results: results[offset-count : max]}, nil
			} else {
				return &model.SearchResults{Results: results[offset:max]}, nil
			}
		})
	defer iterator.Close()
	// first next
	hasNext = iterator.Next()
	assert.True(t, hasNext)
	searchResults, _ = iterator.Value()
	assert.Equal(t, &model.SearchResults{Results: results[0:2]}, searchResults)
	// second next
	hasNext = iterator.Next()
	assert.True(t, hasNext)
	searchResults, _ = iterator.Value()
	assert.Equal(t, &model.SearchResults{Results: results[2:4]}, searchResults)
	// third next
	hasNext = iterator.Next()
	assert.True(t, hasNext)
	searchResults, _ = iterator.Value()
	assert.Equal(t, &model.SearchResults{Results: results[4:5]}, searchResults)
	// fourth next: break out
	hasNext = iterator.Next()
	assert.False(t, hasNext)
	value4, _ := iterator.Value()
	var emptyResults *model.SearchResults
	assert.ObjectsAreEqualValues(value4, emptyResults)
}

func TestNextOnClose(t *testing.T) {
	iterator := NewSearchIterator(10, 0, 100,
		func(count, offset int) (*model.SearchResults, error) {
			return &model.SearchResults{}, nil
		})
	iterator.Close()
	assert.False(t, iterator.Next())
}

func TestNextOnErr(t *testing.T) {
	iterator := NewSearchIterator(10, 0, 100,
		func(count, offset int) (*model.SearchResults, error) {
			return &model.SearchResults{}, nil
		})
	defer iterator.Close()
	iterator.err = errors.New("test error")
	assert.False(t, iterator.Next())
}

func TestNextOnZeroBatch(t *testing.T) {
	batch := 0
	max := 5
	results := make([]*model.Result, max)
	var searchResults *model.SearchResults
	for i := 0; i < max; i++ {
		results = append(results, &model.Result{Raw: fmt.Sprintf("test data %v", i)})
	}
	iterator := NewSearchIterator(batch, 0, max,
		func(count, offset int) (*model.SearchResults, error) {
			if count == 0 {
				return &model.SearchResults{Results: results}, nil
			}
			return nil, nil
		})
	assert.False(t, iterator.Next())
	searchResults, _ = iterator.Value()
	assert.Equal(t, &model.SearchResults{Results: results}, searchResults)
}

func TestNextOnFnErr(t *testing.T) {
	batch := 0
	max := 5
	var searchResults *model.SearchResults
	iterator := NewSearchIterator(batch, 0, max,
		func(count, offset int) (*model.SearchResults, error) {
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
	// results := make([]*model.Result, max)
	var searchResults *model.SearchResults
	// for i := 0; i < max; i++ {
	// 	results = append(results, &model.Result{Raw: fmt.Sprintf("test data %v", i)})
	// }
	iterator := NewSearchIterator(batch, 0, max,
		func(count, offset int) (*model.SearchResults, error) {
			return &model.SearchResults{Results: []*model.Result{}}, nil
		})
	assert.False(t, iterator.Next())
	searchResultsActual, _ := iterator.Value()
	assert.Equal(t, searchResults, searchResultsActual)
}