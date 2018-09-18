// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package service

import (
	"errors"
	"github.com/splunk/splunk-cloud-sdk-go/model"
)

type queryFunc func(step, start int) (*model.SearchResults, error)

// SearchIterator is the result of a search query. Its cursor starts at 0 index
// of the result set. Use Next() to advance through the rows:
//
//  search, _ := client.SearchService.SubmitSearch(&model.PostJobsRequest{Search: "search index=main | head 5"})
// 	pages, _ := search.QueryResults(2, 0, &model.FetchResultsRequest{Count: 5})
// 	defer pages.Close()
// 	for pages.Next() {
// 		values, err := pages.Value()
//      ...
//
// 	}
// 	err := pages.Err() // get any error encountered during iteration
//  ...
type SearchIterator struct {
	value    *model.SearchResults // stores current value
	max      int                  // max number of results
	start    int                  // index to start with fetching results, same concept as "offset"
	offset   int                  // offset value to start iterator with. e.g. offset=5 means iterator will skip the first 5 results
	batch    int                  // batch size of results in each Next call
	err      error                // error encountered during iteration
	fn       queryFunc            // function to be executed in each Next call
	isClosed bool                 // signal indicating status of the iterator
}

// NewSearchIterator creates a new reference to the iterator object
func NewSearchIterator(batch, offset, max int, fn queryFunc) *SearchIterator {
	return &SearchIterator{
		start:    offset,
		batch:    batch,
		max:      max,
		fn:       fn,
		isClosed: false,
		offset:   offset,
	}
}

// Value returns value in current iteration or error out if iterator is closed
func (i *SearchIterator) Value() (*model.SearchResults, error) {
	if i.isClosed == true {
		return nil, errors.New("failed to retrieve values on a closed iterator")
	}
	return i.value, nil
}

// Next prepares the next result set for reading with the Value method. It
// returns true on success, or false if there is no next result row or an error
// occurred while preparing it.
//
// Every call to Value, even the first one, must be preceded by a call to Next.
func (i *SearchIterator) Next() bool {
	if i.start > i.max || i.isClosed == true || i.err != nil {
		return false
	}
	results, err := i.fn(i.batch, i.start)
	if err != nil {
		i.err = err
		return false
	}
	// No more results
	if len(results.Results) == 0 {
		return false
	}
	i.value = results
	// Return all results, therefore no longer need to iterate
	if i.batch == 0 {
		return false
	}
	i.start += i.batch
	return true
}

// Close checks the status and closes iterator if it's not already. After Close, no results can be retrieved
func (i *SearchIterator) Close() {
	if i.isClosed != true {
		i.isClosed = true
	}
}

// Err returns error encountered during iteration
func (i *SearchIterator) Err() error {
	return i.err
}
