/*
 * Copyright 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/splunk/splunk-cloud-sdk-go/v2/services/ml"
)

const (
	MLServiceVersion = "v2beta1"
)

var createMachineLearningService = func() *ml.Service {
	return apiClient().MachineLearningService
}

type MachineLearningCommand struct {
	machineLearningService *ml.Service
}

func newMachineLearningCommand() *MachineLearningCommand {
	return &MachineLearningCommand{
		machineLearningService: createMachineLearningService(),
	}
}

func (cmd *MachineLearningCommand) Dispatch(args []string) (result interface{}, err error) {
	arg, args := head(args)
	switch arg {
	case "":
		eusage("too few arguments")
	case "create-workflow":
		result, err = cmd.createWorkflow(args)
	case "create-workflow-build":
		result, err = cmd.createWorkflowBuild(args)
	case "create-workflow-run":
		result, err = cmd.createWorkflowBuildRun(args)
	case "delete-workflow":
		err = cmd.deleteWorkflow(args)
	case "delete-workflow-build":
		err = cmd.deleteWorkflowBuild(args)
	case "delete-workflow-run":
		err = cmd.deleteWorkflowBuildRun(args)
	case "help":
		err = help("ml.txt")
	case "get-spec-json":
		result, err = cmd.getSpecJSON(args)
	case "get-spec-yaml":
		result, err = cmd.getSpecYaml(args)
	case "get-workflow":
		result, err = cmd.getWorkflow(args)
	case "get-workflow-build":
		result, err = cmd.getWorkflowBuild(args)
	case "get-workflow-run":
		result, err = cmd.getWorkflowBuildRun(args)
	case "list-workflows":
		result, err = cmd.listWorkflows(args)
	case "list-workflow-builds":
		result, err = cmd.listWorkflowBuilds(args)
	case "list-workflow-runs":
		result, err = cmd.listWorkflowBuildRuns(args)
	default:
		fatal("unknown command: '%s'", arg)
	}
	return
}

func readFile(args []string) []byte {
	file := head1(args)
	fileIn, err := os.Open(file)
	if err != nil {
		fatal(err.Error())
	}
	defer fileIn.Close()
	payload, err := ioutil.ReadAll(fileIn)
	if err != nil {
		fatal(err.Error())
	}
	return payload
}

func (cmd *MachineLearningCommand) createWorkflow(args []string) (interface{}, error) {
	payload := readFile(args)

	var workflow ml.Workflow
	err := json.Unmarshal(payload, &workflow)
	if err != nil {
		fatal(err.Error())
	}
	return cmd.machineLearningService.CreateWorkflow(workflow)
}

func (cmd *MachineLearningCommand) createWorkflowBuild(args []string) (interface{}, error) {
	workflowID, args := head(args)
	payload := readFile(args)

	var workflowBuild ml.WorkflowBuild
	err := json.Unmarshal(payload, &workflowBuild)
	if err != nil {
		fatal(err.Error())
	}
	return cmd.machineLearningService.CreateWorkflowBuild(workflowID, workflowBuild)
}
func (cmd *MachineLearningCommand) createWorkflowBuildRun(args []string) (interface{}, error) {
	workflowID, args := head(args)
	buildID, args := head(args)
	payload := readFile(args)

	var workflowRun ml.WorkflowRun
	err := json.Unmarshal(payload, &workflowRun)
	if err != nil {
		fatal(err.Error())
	}
	return cmd.machineLearningService.CreateWorkflowRun(workflowID, buildID, workflowRun)
}

func (cmd *MachineLearningCommand) deleteWorkflow(argv []string) error {
	workflowID := head1(argv)
	return cmd.machineLearningService.DeleteWorkflow(workflowID)
}

func (cmd *MachineLearningCommand) deleteWorkflowBuild(argv []string) error {
	workflowID, buildID := head2(argv)
	return cmd.machineLearningService.DeleteWorkflowBuild(workflowID, buildID)
}

func (cmd *MachineLearningCommand) deleteWorkflowBuildRun(argv []string) error {
	workflowID, argv := head(argv)
	buildID, runID := head2(argv)
	return cmd.machineLearningService.DeleteWorkflowRun(workflowID, buildID, runID)
}

func (cmd *MachineLearningCommand) getWorkflow(argv []string) (interface{}, error) {
	workflowID := head1(argv)
	return cmd.machineLearningService.GetWorkflow(workflowID)
}

func (cmd *MachineLearningCommand) getWorkflowBuild(argv []string) (interface{}, error) {
	workflowID, buildID := head2(argv)
	return cmd.machineLearningService.GetWorkflowBuild(workflowID, buildID)
}

func (cmd *MachineLearningCommand) getWorkflowBuildRun(argv []string) (interface{}, error) {
	workflowID, argv := head(argv)
	buildID, runID := head2(argv)
	return cmd.machineLearningService.GetWorkflowRun(workflowID, buildID, runID)
}

func (cmd *MachineLearningCommand) listWorkflows(argv []string) (interface{}, error) {
	return cmd.machineLearningService.ListWorkflows()
}

func (cmd *MachineLearningCommand) listWorkflowBuilds(argv []string) (interface{}, error) {
	workflowID := head1(argv)
	return cmd.machineLearningService.ListWorkflowBuilds(workflowID)
}

func (cmd *MachineLearningCommand) listWorkflowBuildRuns(argv []string) (interface{}, error) {
	workflowID, buildID := head2(argv)
	return cmd.machineLearningService.ListWorkflowRuns(workflowID, buildID)
}

func (cmd *MachineLearningCommand) getSpecJSON(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return GetSpecJSON("api", MLServiceVersion, "ml", cmd.machineLearningService.Client)
}

func (cmd *MachineLearningCommand) getSpecYaml(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return GetSpecYaml("api", MLServiceVersion, "ml", cmd.machineLearningService.Client)
}
