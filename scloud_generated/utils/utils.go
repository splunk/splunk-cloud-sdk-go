package utils

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"os"
	"strings"
)


func Pprint(value interface{}) {
	if value == nil {
		return
	}
	switch vt := value.(type) {
	case string:
		fmt.Print(vt)
		if !strings.HasSuffix(vt, "\n") {
			fmt.Println()
		}
	default:
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "    ")
		err := encoder.Encode(value)
		if err != nil {
			fatal("json pprint error: %s", err.Error())
		}
	}
}
func GetClient() (*sdk.Client,error){
	//client,err:=sdk.NewClient(&services.Config{
	//	Token:       "[REDACTED]",
	//	Host:         "staging.scp.splunk.com",
	//	Tenant:       "testsdks",
	//})


	glog.CopyStandardLogTo("INFO")

	load()


	client:=apiClient()


	return client, nil
}


// Authenticate, using the selected app profile.
func Login(args []string) (*idp.Context, error) {
	load()
	return login(args)
}


func Head1(items []string) string {
	return head1(items)
}