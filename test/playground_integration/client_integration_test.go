/*
 * Copyright © 2018 Splunk Inc.
 * SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
 * without a valid written license from Splunk Inc. is PROHIBITED.
 *
 */

package playgroundintegration

import (
	"testing"

	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/stretchr/testify/require"
)

func getClient(t *testing.T) *service.Client {
	var url = testutils.TestURLProtocol + "://" + testutils.TestSSCHost

	//fmt.Printf("=================================================================")
	//fmt.Printf("CREATING A CLIENT WITH THESE SETTINGS")
	//fmt.Printf("=================================================================")
	//fmt.Printf("Authentication Token: " + testutils.TestAuthenticationToken + "\n")
	//fmt.Printf("SSC Host API: " + testutils.TestSSCHost + "\n")
	//fmt.Printf("Tenant ID: " + testutils.TestTenantID + "\n")
	//fmt.Printf("URL Protocol: " + testutils.TestURLProtocol + "\n")
	//fmt.Printf("Fully Qualified URL: " + url + "\n")
	client, err := service.NewClient(&service.Config{Token: testutils.TestAuthenticationToken, URL: url, TenantID: testutils.TestTenantID, Timeout: testutils.TestTimeOut})
	require.Emptyf(t, err, "Error calling service.NewClient(): %s", err)
	return client
}

func getInvalidClient(t *testing.T) *service.Client {
	var url = testutils.TestURLProtocol + "://" + testutils.TestSSCHost

	client, err := service.NewClient(&service.Config{Token: testutils.TestAuthenticationToken, URL: url, TenantID: testutils.TestInvalidTestTenantID, Timeout: testutils.TestTimeOut})
	require.Emptyf(t, err, "Error calling service.NewClient(): %s", err)
	return client
}
