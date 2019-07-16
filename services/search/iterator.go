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
	"errors"
)

// QueryFunc is the function to be executed in each Next call of the iterator
type QueryFunc func(step, start int) (*ListSearchResultsResponse, error)

// Iterator is the result of a search query. Its cursor starts at 0 index
// of the result set. Use Next() to advance through the rows:
//
//  search, _ := client.SearchService.SubmitSearch(&PostJobsRequest{Search: "search index=main | head 5"})
// 	pages, _ := search.QueryResults(2, 0, &FetchResultsRequest{Count: 5})
// 	defer pages.Close()
// 	for pages.Next() {
// 		values, err := pages.Value()
//      ...
//
// 	}
// 	err := pages.Err() // get any error encountered during iteration
//  ...
type Iterator struct {
	value    *ListSearchResultsResponse // stores current value
	max      int                        // max number of results
	start    int                        // index to start with fetching results, same concept as "offset"
	offset   int                        // offset value to start iterator with. e.g. offset=5 means iterator will skip the first 5 results
	batch    int                        // batch size of results in each Next call
	err      error                      // error encountered during iteration
	fn       QueryFunc                  // function to be executed in each Next call
	isClosed bool                       // signal indicating status of the iterator
}

// NewIterator creates a new reference to the iterator object
func NewIterator(batch, offset, max int, fn QueryFunc) *Iterator {
	return &Iterator{
		start:    offset,
		batch:    batch,
		max:      max,
		fn:       fn,
		isClosed: false,
		offset:   offset,
	}
}

// Value returns value in current iteration or error out if iterator is closed
func (i *Iterator) Value() (*ListSearchResultsResponse, error) {
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
func (i *Iterator) Next() bool {
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
func (i *Iterator) Close() {
	if i.isClosed != true {
		i.isClosed = true
	}
}

// Err returns error encountered during iteration
func (i *Iterator) Err() error {
	return i.err
}
