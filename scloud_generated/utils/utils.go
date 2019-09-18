package utils

import (
	"github.com/golang/glog"
	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/splunk/splunk-cloud-sdk-go/sdk"
)

var sdkclient *sdk.Client

func Pprint(value interface{}) {
	pprint(value)
}

func GetClient() (*sdk.Client,error) {
	if sdkclient == nil {
		glog.CopyStandardLogTo("INFO")

		load()

		client := apiClient()
		return client, nil
	}

	return sdkclient, nil
}

func GetClientSystemTenant() (*sdk.Client,error) {
	if sdkclient == nil {
		glog.CopyStandardLogTo("INFO")

		load()

		client := apiClientWithTenant("system")
		return client, nil
	}

	return sdkclient, nil
}


// Authenticate, using the selected app profile.
func Login(args []string) (*idp.Context, error) {
	load()
	return login(args)
}


func Head1(items []string) string {
	return head1(items)
}

func Head(items []string) (string, []string) {
	return head(items)
}


func Head2(items []string) (string, string) {
	return head2(items)
}

// Prints an error message and exits.
func Fatal(msg string, args ...interface{}) {
	fatal(msg,args)
}

func CheckEmpty(items []string) {
	 checkEmpty(items)
}

func GetTenantName() string {
	return getTenantName()
}
