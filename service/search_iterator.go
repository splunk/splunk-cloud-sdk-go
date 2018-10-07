// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package service

import (
	"github.com/splunk/splunk-cloud-sdk-go/services/search"
)

// SearchIterator is DEPRECATED, please use services/search.SearchIterator
type SearchIterator = search.SearchIterator

// NewSearchIterator is DEPRECATED, please use services/search.NewSearchIterator
func NewSearchIterator(batch, offset, max int, fn search.QueryFunc) *SearchIterator {
	return search.NewSearchIterator(batch, offset, max, fn)
}
