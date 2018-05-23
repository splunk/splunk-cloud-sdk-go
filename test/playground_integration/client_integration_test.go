package playgroundintegration

import (
	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/testutils"
)

func getClient() *service.Client {
	var url = testutils.TestURLProtocol + "://" + testutils.TestSSCHost

	//fmt.Printf("=================================================================")
	//fmt.Printf("CREATING A CLIENT WITH THESE SETTINGS")
	//fmt.Printf("=================================================================")
	//fmt.Printf("Authentication Token: " + testutils.TestAuthenticationToken + "\n")
	//fmt.Printf("SSC Host API: " + testutils.TestSSCHost + "\n")
	//fmt.Printf("Tenant ID: " + testutils.TestTenantID + "\n")
	//fmt.Printf("URL Protocol: " + testutils.TestURLProtocol + "\n")
	//fmt.Printf("Fully Qualified URL: " + url + "\n")

	return service.NewClient(testutils.TestTenantID, testutils.TestAuthenticationToken, url, testutils.TestTimeOut)
}

// TODO?
