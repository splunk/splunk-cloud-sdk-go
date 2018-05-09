package test_utils

import (
	"time"

	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/util"
)

func GetSplunkClient(local ...bool) *service.Client {
	if len(local) > 0 {
		return service.NewClient(util.TestTenantID, util.TestToken, util.TestStubbySchme+"://"+util.TestSubbyLocalHost, time.Second*5)
	}
	return service.NewClient(util.TestTenantID, util.TestToken, util.TestStubbySchme+"://"+util.TestStubbyHost, time.Second*5)
}
