package test

import (
	"os"
	"testing"

	test_engine "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/test/utils"
)

// Tests in this file are recording test results. They should only be executed when updating test results

var testhook_arg = "--testhook-dryrun" //use --testhook to run against service; use --testhook-dryrun to record test results only

func Test_record_action(t *testing.T) {
	skipCI(t)
	test_engine.Record_test_result("testcases/action_test", testhook_arg, t)
}

func Test_record_appreg(t *testing.T) {
	skipCI(t)
	test_engine.Record_test_result("testcases/appreg_test", testhook_arg, t)
}

func Test_record_collect(t *testing.T) {
	skipCI(t)
	test_engine.Record_test_result("testcases/collect_test", testhook_arg, t)
}

func Test_record_identity(t *testing.T) {
	skipCI(t)
	test_engine.Record_test_result("testcases/identity_test", testhook_arg, t)
}

func Test_record_ingest(t *testing.T) {
	skipCI(t)
	test_engine.Record_test_result("testcases/ingest_test", testhook_arg, t)
}


func Test_record_kvstore(t *testing.T) {
	skipCI(t)
	test_engine.Record_test_result("testcases/kvstore_test", testhook_arg, t)
}

func Test_record_provisioner(t *testing.T) {
	skipCI(t)
	test_engine.Record_test_result("testcases/provisioner_test", testhook_arg, t)
}

func Test_record_search(t *testing.T) {
	skipCI(t)
	test_engine.Record_test_result("testcases/search_test", testhook_arg, t)
}

func Test_record_streams(t *testing.T) {
	skipCI(t)
	test_engine.Record_test_result("testcases/streams_test", testhook_arg, t)
}

func skipCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skip testing in CI environment")
	}
}
