package test

import (
	"testing"

	test_util "github.com/splunk/splunk-cloud-sdk-go/scloud_generated/test/utils"
)

func TestCli_action(t *testing.T) {
	test_util.RunTest("testcases/action_test", t)
}

//func TestCli_identit(t *testing.T) {
//	test_util.RunTest("testcases/identity_test", t)
//}
