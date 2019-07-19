
import unittest

import test
import time


def streams(*args):
    return test.scloud("streams", *args)


class TestStreams(unittest.TestCase):
    def setUp(self):
        # retrieve the selected tenant name
        code, self.tname, _ = test.scloud("get", "tenant")
        self.assertEqual(0, code)
        self.assertIsNotNone(self.tname)

    def test_compile_dsl(self):
        code, result, _ = streams("compile-dsl", "--dsl-file", "pass-through-v2.dsl")
        print "\n--Compile-dsl-Test--"
        print "Validate Success Code"
        self.assertEqual(0, code)

    def test_create_pipeline(self):
        pipelinename = "integrationtestpipeline_" + str(int(round(time.time())))
        code, result, _ = streams("create-pipeline", "-name", pipelinename, "-description", "Pipeline for integration tests", "-bypass-validation", "true", "-data-file", "ps.upl")
        print "\n--Create-Pipeline-Test--"
        print "\n1.Create-Pipeline"
        print "Validate Success Code"
        id = result["id"]
        self.assertEqual(0, code)

        code, result, _ = streams("delete-pipeline", id)
        print "\n2.Delete-Pipeline"
        print "Validate Success Code"
        self.assertEqual(0, code)

    def test_get_pipeline(self):
        print "\n--Get-Pipeline-Tests--"
        pipelinename = "integrationtestpipeline_" + str(int(round(time.time())))
        code, result, _ = streams("create-pipeline", "-name", pipelinename, "-description", "Pipeline for integration test", "-bypass-validation", "true", "-data-file", "ps.upl")
        print "\n1.Create-Pipeline"
        print "Validate Success Code"
        id = result["id"]
        self.assertEqual(0, code)

        code, result, _ = streams("get-pipeline", id)
        print "\n2.Get-Pipeline"
        print "Validate Success Code"
        self.assertEqual(0, code)

        code, result, _ = streams("delete-pipeline", id)
        print "\n3.Delete-Pipeline"
        print "Validate Success Code"
        self.assertEqual(0, code)

    def test_replace_pipeline(self):
        print "\n--Update-Pipeline-Tests--"
        pipelinename = "integrationtestpipeline_" + str(int(round(time.time())))
        code, result, _ = streams("create-pipeline", "-name", pipelinename, "-description", "Pipeline for integration test", "-bypass-validation", "true", "-data-file", "ps.upl")
        print "\n1.Create-Pipeline"
        print "Validate Success Code"
        id = result["id"]
        self.assertEqual(0, code)

        code, result, _ = streams("replace-pipeline", id, "-name", pipelinename, "-description", "Updated Pipeline for integration test", "-bypass-validation", "true", "-data-file", "ps.upl")
        print "\n2.Update-Pipeline"
        print "Validate Success Code"
        self.assertEqual(0, code)

        code, result, _ = streams("delete-pipeline", id)
        print "\n3.Delete-Pipeline"
        print "Validate Success Code"
        self.assertEqual(0, code)

    def test_list_pipelines(self):
        pipelineName = "integrationtestpipeline_" + str(int(round(time.time())))
        code, result, _ = streams("create-pipeline", "-name", pipelineName, "-description", "Pipeline for integration tests", "-bypass-validation", "true", "-data-file", "ps.upl")
        id = result["id"]
        print "\n--Create-Pipeline-Test--"
        print "\n1.Create-Pipeline"
        print "Validate Success Code"
        self.assertEqual(0, code)

        print "\n--List-Pipelines-Tests--"
        code, result, _ = streams("list-pipelines")
        print "\n2.--List Pipelines"
        print "Validate Success Code"
        print code
        print "\n3.Validate at least one pipeline is found"
        cnt = len(result)
        self.assertTrue(cnt >= 1)

        code, result, _ = streams("delete-pipeline", id)
        print "\n4.Delete-Pipeline"
        print "Validate Success Code"
        print code
        self.assertEqual(0, code)

    def test_get_pipeline_status(self):
        print "\n--Get-Pipeline-Status-Tests--"
        pipelinename = "integrationtestpipeline_" + str(int(round(time.time())))
        code, result, _ = streams("create-pipeline", "-name", pipelinename, "-description", "Pipeline for integration test", "-bypass-validation", "true", "-data-file", "ps.upl")
        print "\n1.Create-Pipeline"
        print "Validate Success Code"
        id = result["id"]
        name = result["name"]
        self.assertEqual(0, code)

        code, result, _ = streams("get-pipeline-status", "-name", name)
        print "\n 1.Get-Pipeline-Status"
        print "Validate Success Code"
        self.assertEqual(0, code)

        code, result, _ = streams("delete-pipeline", id)
        print "\n2.Delete-Pipeline--"
        print "Validate Success Code"
        self.assertEqual(0, code)


    # def test_activate_pipelines(self):
    #     print "\n--Activate-Pipelines-Tests--"
    #     pipelinename = "integrationtestpipeline_" + str(int(round(time.time())))
    #     code, result, _ = streams("create-pipeline", "-name", pipelinename, "-description", "Pipeline for integration test", "-bypass-validation", "true", "-data-file", "ps.upl")
    #     print "\n1.Create-Pipeline"
    #     print "Validate Success Code"
    #     id = result["id"]
    #     self.assertEqual(0, code)
    #
    #     time.sleep( 5 )
    #     code, result, _ = streams("activate-pipelines", id)
    #     print "\n2.Activate-Pipeline"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)
    #
    #     code, result, _ = streams("delete-pipeline", id)
    #     print "\n3.Delete-Pipeline"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)
    #
    # def test_deactivate_pipelines(self):
    #     print "\n--Deactivate-Pipelines-Tests--"
    #     pipelinename = "integrationtestpipeline_" + str(int(round(time.time())))
    #     code, result, _ = streams("create-pipeline", "-name", pipelinename, "-description", "Pipeline for integration test", "-bypass-validation", "true", "-data-file", "ps.upl")
    #     print "\n1.Create-Pipeline"
    #     print "Validate Success Code"
    #     id = result["id"]
    #     self.assertEqual(0, code)
    #
    #     code, result, _ = streams("activate-pipelines", id)
    #     print "\n2.Activate-Pipeline"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)
    #     time.sleep( 5 )
    #
    #     code, result, _ = streams("deactivate-pipelines", id)
    #     print "\n3.Deactivate-Pipeline"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)
    #
    #     code, result, _ = streams("delete-pipeline", id)
    #     print "\n4.Delete-Pipeline"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)

    def test_merge_pipelines(self):
        print "\n--Merge-Pipelines-Tests--"
        code, result, _ = streams("merge-pipelines", "343aaa0e-b99b-43a4-878e-bbd1fbdba359", "input", "-input-tree-file", "ps_main.upl", "-main-tree-file", "ps.upl")
        print "\n1.Merge-Pipelines"
        print "Validate Success Code"
        self.assertEqual(0, code)

    # def test_validate_upl(self):
    #     print "\n--Validate-Upl-Tests--"
    #     code, result, _ = streams("validate-upl", "-upl-file", "ps.upl")
    #     print "\n1.Validate-Upl"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)
    #
    # def test_get_input_schema(self):
    #     print "\n--Get-Input-Schema-Test--"
    #     code, result, _ = streams("get-input-schema", "343aaa0e-b99b-43a4-878e-bbd1fbdba359", "input", "-upl-file", "ps.upl")
    #     print "\n1.Get-Input-Schema"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)
    #
    # def test_get_output_schema(self):
    #     print "\n--Get-Output-Schema-Test--"
    #     code, result, _ = streams("get-output-schema", "0909d75b-8518-42a8-987c-7216cc6c6f73", "output", "-upl-file", "ps.upl")
    #     print "\n1.Get-Output-Schema"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)

    # Temporarily commenting out template related tests since the service is returning a 500 error for these endpoints
    def test_create_template(self):
        print "\n--Create-Template-Tests--"
        templateName = "integrationtesttemplate_" + str(int(round(time.time())))
        code, result, _ = streams("create-template", "-name", templateName, "-description", "Template for integration tests", "-data-file", "ps.upl")
        print 'TEMPLATE'
        print result

        print "\n 1.Create-Template"
        print "Validate Success Code"
        # id = result["templateId"]
        # self.assertEqual(0, code)

        # code, result, _ = streams("delete-template", id)
        # print "\n2.Delete-Template"
        # print "Validate Success Code"
        # self.assertEqual(0, code)

    def test_list_templates(self):
        print "\n--List-Templates-Tests--"
        code, result, _ = streams("list-templates")
        print "Validate Success Code"
        print code
        self.assertEqual(0, code)

    # def test_get_template(self):
    #     print "\n--Get-Template-Tests--"
    #     templateName = "integrationtesttemplate_" + str(int(round(time.time())))
    #     code, result, _ = streams("create-template", "-name", templateName, "-description", "Template for integration test", "-data-file", "ps.upl")
    #     print "\n1.Create-Template"
    #     print "Validate Success Code"
    #     id = result["templateId"]
    #     self.assertEqual(0, code)
    #
    #     code, result, _ = streams("get-template", id, "-version", "1")
    #     print "\n2.Get-Template"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)
    #
    #     code, result, _ = streams("delete-template", id)
    #     print "\n3.Delete-Template"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)


    def test_update_template(self):
        print "\n--Update-Template-Tests--"
        templateName = "integrationtesttemplate_" + str(int(round(time.time())))
        code, result, _ = streams("create-template", "-name", templateName, "-description", "Template for integration test", "-data-file", "ps.upl")
        print "\n1.Create-Template"
        print "Validate Success Code"
        id = result["templateId"]
        self.assertEqual(0, code)

        code, result, _ = streams("update-template", id, "-name", templateName, "-description", "Updated Template for integration test", "-data-file", "ps.upl")
        print "\n2.Update-Template"
        print "Validate Success Code"
        self.assertEqual(0, code)

        code, result, _ = streams("delete-template", id)
        print "\n3.Delete-Template"
        print "Validate Success Code"
        self.assertEqual(0, code)

    def test_update_template_partially(self):
        print "\n--Update-Partial-Template-Tests--"
        templateName = "integrationtesttemplate_" + str(int(round(time.time())))
        code, result, _ = streams("create-template", "-name", templateName, "-description", "Template for integration test", "-data-file", "ps.upl")
        print "\n1.Create-Template"
        print "Validate Success Code"
        id = result["templateId"]
        self.assertEqual(0, code)

        code, result, _ = streams("update-template-partially", id, "-description", "Updated Template for integration test")
        print "\n2.Update-Template-Partially"
        print "Validate Success Code"
        self.assertEqual(0, code)

        code, result, _ = streams("delete-template", id)
        print "\n3.Delete-Template"
        print "Validate Success Code"
        self.assertEqual(0, code)

    # TODO: enable the tests once SSC-10204 is resolved
    # def test_start_preview_session(self):
    #     print "\n--Start-Preview-Session-Tests--"
    #     code, result, _ = streams("start-preview-session", "-upl-file", "ps.upl")
    #     print "\n1.Start-Preview-Session"
    #     print "Validate Success Code"
    #     id = result["previewId"]
    #     self.assertEqual(0, code)

    #     code, result, _ = streams("delete-preview-session", str(id))
    #     print "\n2.Delete-Preview-Session"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)

    # def test_get_preview_session(self):
    #     print "\n--Get-Preview-Session-Tests--"
    #     code, result, _ = streams("start-preview-session", "-upl-file", "ps.upl")
    #     print "\n1.Start-Preview-Session"
    #     print "Validate Success Code"
    #     id = result["previewId"]
    #     self.assertEqual(0, code)

    #     code, result, _ = streams("get-preview-session", str(id))
    #     print "\n2.Get-Preview-Session"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)

    #     code, result, _ = streams("delete-preview-session", str(id))
    #     print "\n2.Delete-Preview-Session"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)

    # def test_get_preview_data(self):
    #     print "\n--Get-Preview-Data-Tests--"
    #     code, result, _ = streams("start-preview-session", "-upl-file", "ps.upl")
    #     print "\n 1.Start-Preview-Session"
    #     print "Validate Success Code"
    #     id = result["previewId"]
    #     self.assertEqual(0, code)

    #     code, result, _ = streams("get-preview-data", str(id))
    #     print "\n2.Get-Preview-Data"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)

    #     code, result, _ = streams("delete-preview-session", str(id))
    #     print "\n2.Delete-Preview-Session"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)

    # def test_list_latest_preview_metrics(self):
    #     print "\n--Get-Latest-Preview-Metrics-Tests--"
    #     code, result, _ = streams("start-preview-session", "-upl-file", "ps.upl")
    #     print "\n1.Start-Preview-Session"
    #     print "Validate Success Code"
    #     id = result["previewId"]
    #     self.assertEqual(0, code)

    #     code, result, _ = streams("list-latest-preview-metrics", str(id))
    #     print "\n2.Get-latest-Preview-metrics"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)

    #     code, result, _ = streams("delete-preview-session", str(id))
    #     print "\n3.Delete-Preview-Session"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)

    # Commenting out since streams service in staging is currently flaky
    # def test_list_latest_pipeline_metrics(self):
    #     print "\n--Get-latest-Pipeline-metrics-Tests--"
    #     pipelinename = "integrationtestpipeline_" + str(int(round(time.time())))
    #     code, result, _ = streams("create-pipeline", "-name", pipelinename, "-description",
    #                               "Pipeline for integration test", "-bypass-validation", "true", "-data-file", "ps.upl")
    #     print "\n1.Create-Pipeline"
    #     print "Validate Success Code"
    #     id = result["id"]
    #     self.assertEqual(0, code)
    #
    #     time.sleep( 5 )
    #     code, result, _ = streams("activate-pipelines", id)
    #     print "\n2.Activate-Pipeline"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)
    #
    #     code, result, _ = streams("list-latest-pipeline-metrics", id)
    #     print "\n2.Get-latest-Pipeline-metrics"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)
    #
    #     code, result, _ = streams("delete-pipeline", id)
    #     print "\n3.Delete-Pipeline"
    #     print "Validate Success Code"
    #     self.assertEqual(0, code)

    def test_get_connectors(self):
        print "\n--Get-Connectors-Test--"
        code, result, _ = streams("get-connectors")
        print "Validate Success Code"
        print code
        self.assertEqual(0, code)

    def test_get_connections(self):
        print "\n1.Get connections for a specific Connector"
        code, result, _ = streams("list-connections", "-connector-id", "debug-connector")
        print "Validate Success Code"
        self.assertEqual(0, code)

    def test_get_registry(self):
        print "\n--Get-Registry-Tests-"
        func_name = None
        code, result, _ = streams("get-registry")
        print "\n1.Get-Registry"
        print "Validate Success Code"
        self.assertEqual(0, code)
        print "\n2.Validate function exists in registry results"
        functions = result['functions']
        for key in functions:
            attributes = key['attributes']
            application = attributes['application']
            for key1 in application :
                if key1 == 'name' and application['name'] == 'Receive from Ingest REST API':
                    func_name = 'Receive from Ingest REST API'

        self.assertIsNotNone(func_name)

    def test_get_groups(self):
        print "\n--Get-Groups-Tests--"
        group_id = None
        application_name = None
        code, result, _ = streams("get-registry")
        print "\n1.Get-Registry and Get-Group-Id"
        print "Validate Success Code"
        print code
        self.assertEqual(0, code)

        functions = result['functions']
        for key in functions:
            attributes = key['attributes']
            application = attributes['application']
            for key1 in application :
                if key1 == 'name' and application['name'] == 'Receive from Ingest REST API':
                    application_name = 'Receive from Ingest REST API'
                    group_id = application['groupId']

        self.assertIsNotNone(application_name)
        self.assertIsNotNone(group_id)
        #ToDo: Test failure investigation pending, service is currently returning a 404 for this endpoint
        # code, result, _ = streams("get-group", group_id)
        # print "\n2.Get-Group"
        # print "Validate Success Code"
        # self.assertEqual(0, code)
        # mappings = result["mappings"]
        # group_function_id = mappings[0]["function_id"]
        # print "\n3.Validate group function id exists"
        # self.assertIsNotNone(group_function_id)



    # def test_expanded_group(self):
    #     print "\n--Create-Expand-Groups-Tests--"
    #     group_id = None
    #     group_function_id = None
    #     version = None
    #     code, result, _ = streams("get-registry")
    #     print "\n1.Get-Registry and GroupId"
    #     print "Validate Success Code"
    #     print code
    #     self.assertEqual(0, code)
    #     functions = result['functions']
    #     for key in functions:
    #         if key['name'] == "receive-from-ingest-rest-api":
    #             attributes = key['attributes']
    #             application = attributes['application']
    #             for key1 in application :
    #                 if key1 == 'groupId':
    #                     group_id = application['groupId']
    #
    #     self.assertIsNotNone(group_id)
    #     #ToDo: Test failure investigation pending, service is currently returning a 404 for this endpoint
    #     # code, result, _ = streams("get-group", group_id)
    #     # print "\n2.Get-Group and get Group function id"
    #     # print "Validate Success Code"
    #     # print code
    #     # self.assertEqual(0, code)
    #     # mappings = result["mappings"]
    #     # group_function_id = mappings[0]["function_id"]
    #     # arguments = "{\"function_arg\": \"right\", \"group_arg\": \"connection\"}"
    #     # self.assertIsNotNone(group_function_id)
    #     #ToDo: Test failure investigation pending, works manually
    #     #
    #     # code, result, _ = streams("create-expanded-group", "-groupId", group_id, "-groupFunctionId", group_function_id, "-arguments", arguments)
    #     # print "\n3.Create-Expand-Group"
    #     # print "Validate Success Code"
    #     # self.assertEqual(0, code)
    #     # print "\n4.Validate version exists in expanded group result"
    #     # version = result["version"]
    #     # self.assertIsNotNone(version)
