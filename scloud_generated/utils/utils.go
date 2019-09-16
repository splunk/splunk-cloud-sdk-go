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
	//	Token:       "IACMrEUeFq6c30LoA7jcxZYvyV-9SDrjgT2Nq9ukzteLYzsdd19f7UO947Gr51QCtJEQ8x-LWqt7nQn8YsW3o4cyyyy-Ws78nOSMYiofKIUpXn9Lvb8HfxbXnOx4HrdcQjG5MSnXK-b-g_HKwIKdcjvSLZh4H3pu_ADqgOW33SKFVdZ9RgQJ7lsOTY6_L1_29ikz7Fh9lxeCbPtyjfCWRCN0yXQbyAurvZdljpjWfU9yRA2aFihOQwOIaVX14Whj8G4yAxc6nxpn8z1bdg0GqjlVrLTDVFLmfC9Bl8-nsy8s9QvqBRk2zGuTsvi9pIAtPJOEKe0RwCh1Ij3b9rKm-JeQUm5OAKaEiSLKGuNlGwo13GW_RdHnKx3vGKJ_G8N2dYV5FAzEojnib7pkzjen762ROcnP5y244RvsRJtjyPnshz9z7aVAT-sLYPecx-Z1SiidsvZc51THGmMwDmPCkIccMRrqbo8qwaFbqvHouc8RAAYaQuqGeciKYlAodLY0-mKMXDyZul3oiRzYH42BJJjcgu_KEYws0qDIrDumE1jeFtaVrmXCgHNmLG9IlxzkYSdT2lhlWkBAIg8wF3mEQE217FK6v6n5U8mxB8cwn5bv6Jcek8Ne2QA_uo_VGpAeMTd_5irjxk9peOPCTinYKBseu0GVoK56oq7sWRAp-jyCTpNtn2BArOT3sTvE_kyP0h7-IYcbct44v2A5AJl63C26c3aHs6BX8K3PMJTJjV0PFsXUqSW8z-QSTF2BrPHnu1gfxYPixO8LvkH-A3TcsF0PjtEpFRMMKTk6OMd603YGUK5-VNZKodYFYBC0eGvOcFOZ1HwQoYZGWHPjn5B-9TOT3lEXuifM6Bs1VPUSWEq7nCx2HXrZ_K2oXa_AFOTw29r1yOTbMekPIjKDqzOQkR7Kg0ON-PL6LSjpgOZ67-kzhebMFdrl8dRwbM5Z-26r0NyI6fokw7YI8Q5BEjo4g-N-hb25kOhvFnvGAbp0HZoSxgWMBYLYCl-wJY5-3HPAH334QyZyTy7VlKsV_rCBTSSsIFGC9ZJKK4u9MUZQzl1QmTVkSNS6nhqrtlzOR80FhirZBiLyOQNWmt5-3XLp6wzsXnM9eFW-O3vvN1f9uP8Mc6JyJt0r9Pl8PTACk3MvHWUgeY2wIw2Gx2DPml1xSzxrbXL6OVheRF8icwyXvKqS_8ulgXMhYFoHGoKXRrK8PbZd0qRzT61lm5yJdb1jeGVcdxFosDHYWNG6eHqytz-beIfzAMfNUwCMfXuwLJg21XBeeN-IyVDFxs-Iu0zctOORvRU",
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