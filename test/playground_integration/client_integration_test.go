package playgroundintegration

import (
	"fmt"
	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/testutils"
	"testing"
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

	client, err := service.NewClient(testutils.TestTenantID, testutils.TestAuthenticationToken, url, testutils.TestTimeOut)
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}

	return client
}

func getInvalidClient() *service.Client {
	var url= testutils.TestURLProtocol + "://" + testutils.TestSSCHost

	client, _ := service.NewClient(testutils.TestTenantID, testutils.TestInvalidAuthenticationToken, url, testutils.TestTimeOut)
	return client
}

// TODO?
