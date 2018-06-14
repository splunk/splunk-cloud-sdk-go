package service

import (
	"encoding/json"
	"github.com/splunk/ssc-client-go/model"
)

type SearchIterator struct {
	value []byte
	max   int
	batch int
	err   error
	svc   *SearchService
	sid   string
}

func NewSearchIterator(sid string, batch, max int) *SearchIterator {
	return &SearchIterator{
		max:   max,
		batch: batch,
		sid:   sid,
	}
}

func (i *SearchIterator) Value() (*model.SearchResults, error) {
	var results model.SearchResults
	if i.err != nil {
		return &results, i.err
	}
	err := json.Unmarshal(i.value, &results)
	if err != nil {
		return &results, err
	}
	return &results, nil
}

func (i *SearchIterator) Next() bool {
	if i.err != nil {
		return false
	}
}
