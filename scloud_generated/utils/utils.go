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

var sdkclient *sdk.Client

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
func GetClient() (*sdk.Client,error) {
	if sdkclient == nil {
		glog.CopyStandardLogTo("INFO")

		load()

		client := apiClient()
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

// Prints an error message and exits.
func Fatal(msg string, args ...interface{}) {
	fatal(msg,args)
}
