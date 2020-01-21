package test

import (
	"testing"

	test_util "github.com/splunk/splunk-cloud-sdk-go/scloud_generated/test/utils"
)

// Tests in this file are recording test results. They should only be executed when updating test results

var testhook_arg = "--testhook-dryrun" //use --testhook to run against service; use --testhook-dryrun to record test results only

func Test_record_action(t *testing.T) {
	test_util.Record_test_result("testcases/action_test", testhook_arg, t)
}

//func Test_record_identity(t *testing.T) {
//	test_util.Record_test_result("testcases/identity_test", t)
//}
