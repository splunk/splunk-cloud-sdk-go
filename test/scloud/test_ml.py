import unittest
import test
import os

SDK_ML_TENANT = "testsdksml"

ML_WORKFLOW_FILE = os.path.join(os.path.abspath(os.path.dirname(__file__)), "01_ml_workflow.json")
ML_WORKFLOW_BUILD_FILE = os.path.join(os.path.abspath(os.path.dirname(__file__)), "02_ml_build.json")
ML_WORKFLOW_RUN_FILE = os.path.join(os.path.abspath(os.path.dirname(__file__)), "03_ml_run.json")


def ml(*args):
    return test.scloud("ml", *args)


class TestML(unittest.TestCase):
    def setUp(self):
        # retrieve the selected tenant name
        code, self.tname, _ = test.scloud("get", "tenant")
        self.assertEqual(0, code)
        self.assertIsNotNone(self.tname)
        test.scloud("set", "tenant", SDK_ML_TENANT)

    def tearDown(self):
        test.scloud("set", "tenant", self.tname)

    def test_help(self):
        code, result, _ = ml("help")
        self.assertEqual(0, code)

    def test_too_few_args(self):
        code, result, _ = ml("")
        self.assertEqual(1, code)

    # POST tests

    def test_01_create_ml_workflow(self):
        code, result, _ = ml("create-workflow", ML_WORKFLOW_FILE)
        self.assertEqual(0, code)
        self.assertTrue('id' in result)
        self.created_workflow = result["id"]
        self.assertTrue('creationTime' in result)
        self.assertTrue('tasks' in result)

        global created_workflow
        created_workflow = result["id"]

    def test_02_create_ml_workflow_build(self):
        code, result, _ = ml("create-workflow-build", globals().get("created_workflow"), ML_WORKFLOW_BUILD_FILE)
        self.assertEqual(0, code)

        global created_workflow_build
        created_workflow_build = result["id"]

    # GET tests

    def test_04_get_workflow(self):
        code, result, _ = ml("get-workflow", globals().get("created_workflow"))
        self.assertEqual(0, code)
        self.assertTrue('tasks' in result)

    def test_05_get_workflow_build(self):
        code, result, _ = ml("get-workflow-build",
                             globals().get("created_workflow"),
                             globals().get("created_workflow_build"))
        self.assertEqual(0, code)
        self.assertTrue('status' in result)

    # LIST tests

    def test_07_list_workflows(self):
        code, result, _ = ml("list-workflows")
        self.assertEqual(0, code)
        self.assertIsInstance(result, list)
        self.assertTrue(len(result) >= 0)

    # DELETE tests

    def test_12_delete_workflow(self):
        code, result, _ = ml("delete-workflow", globals().get("created_workflow"))
        self.assertEqual(0, code)

    # Spec tests

    def test_get_spec_json(self):
        code, result, _ = ml("get-spec-json")
        self.assertEqual(0, code)
        self.assertTrue(result)

    def test_get_spec_yaml(self):
        code, result, _ = ml("get-spec-yaml")
        self.assertEqual(0, code)
        self.assertTrue(result)


if __name__ == "__main__":
    unittest.main()
