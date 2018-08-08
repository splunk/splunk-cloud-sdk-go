// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package model

// RequestParams contains all the optional request URL parameters
type RequestParams struct {
	Body    interface{}
	Headers map[string]string
}
