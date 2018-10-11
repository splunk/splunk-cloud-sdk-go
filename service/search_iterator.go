// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package service

import (
	"github.com/splunk/splunk-cloud-sdk-go/services/search"
)

// SearchIterator is Deprecated: please use services/search.Iterator
type SearchIterator = search.Iterator

// NewSearchIterator is Deprecated: please use services/search.NewIterator
func NewSearchIterator(batch, offset, max int, fn search.QueryFunc) *SearchIterator {
	return search.NewIterator(batch, offset, max, fn)
}
