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
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/argx"
	"github.com/splunk/splunk-cloud-sdk-go/services/streams"
)

const (
	StreamsServiceVersion = "v1"
)

// StreamsCommand contains info about the streams services
type StreamsCommand struct {
	streamsService *streams.Service
}

func newStreamsCommand() *StreamsCommand {
	return &StreamsCommand{
		streamsService: apiClient().StreamsService,
	}
}

// Dispatch calls respective streams functions based on the command line keyword received
func (streamsCommand *StreamsCommand) Dispatch(argv []string) (result interface{}, err error) {
	arg, argv := head(argv)
	switch arg {
	case "":
		eusage("too few arguments")
	case "activate-pipelines":
		result, err = streamsCommand.activatePipelines(argv)
	case "compile-dsl":
		result, err = streamsCommand.compileDSL(argv)
	case "create-pipeline":
		result, err = streamsCommand.createPipeline(argv)
	case "deactivate-pipelines":
		result, err = streamsCommand.deactivatePipelines(argv)
	case "delete-pipeline":
		result, err = streamsCommand.deletePipeline(argv)
	case "delete-preview-session":
		result, err = streamsCommand.deletePreviewSession(argv)
	case "get-input-schema":
		result, err = streamsCommand.getInputSchema(argv)
	case "list-latest-pipeline-metrics":
		result, err = streamsCommand.getLatestPipelineMetrics(argv)
	case "list-latest-preview-metrics":
		result, err = streamsCommand.getLatestPreviewMetrics(argv)
	case "get-output-schema":
		result, err = streamsCommand.getOutputSchema(argv)
	case "get-pipeline":
		result, err = streamsCommand.getPipeline(argv)
	case "get-pipeline-status":
		result, err = streamsCommand.getPipelineStatus(argv)
	case "list-pipelines":
		result, err = streamsCommand.getPipelines(argv)
	case "get-preview-data":
		result, err = streamsCommand.getPreviewData(argv)
	case "get-preview-session":
		result, err = streamsCommand.getPreviewSession(argv)
	case "get-registry":
		result, err = streamsCommand.getRegistry(argv)
	case "get-spec-json":
		result, err = streamsCommand.getSpecJSON(argv)
	case "get-spec-yaml":
		result, err = streamsCommand.getSpecYaml(argv)
	case "help":
		err = help("streams.txt")
	case "merge-pipelines":
		result, err = streamsCommand.mergePipelines(argv)
	case "reactivate-pipeline":
		result, err = streamsCommand.reactivatePipeline(argv)
	case "replace-pipeline":
		result, err = streamsCommand.replacePipeline(argv)
	case "start-preview-session":
		result, err = streamsCommand.startPreviewSession(argv)
	case "validate-upl":
		result, err = streamsCommand.validateUpl(argv)
	case "get-connectors":
		result, err = streamsCommand.getConnectors(argv)
	case "list-connections":
		result, err = streamsCommand.getConnections(argv)
	case "create-template":
		result, err = streamsCommand.createTemplate(argv)
	case "update-template":
		result, err = streamsCommand.replaceTemplate(argv)
	case "update-template-partially":
		result, err = streamsCommand.updateTemplatePartially(argv)
	case "list-templates":
		result, err = streamsCommand.getTemplates(argv)
	case "get-template":
		result, err = streamsCommand.getTemplate(argv)
	case "delete-template":
		err = streamsCommand.deleteTemplate(argv)
	case "create-expanded-group":
		result, err = streamsCommand.createExpandedGroup(argv)
	case "get-group":
		result, err = streamsCommand.getGroupbyID(argv)
	default:
		fatal("unknown command: '%s'", arg)
	}
	return
}

type activatePipelinesRequestArgs struct {
	ActivateLatestVersion string `arg:"activate-latest-version"`
	AllowNonRestoredState string `arg:"allow-non-restored-state"`
	SkipRestoreState      string `arg:"skip-restorestate"`
}

func (streamsCommand *StreamsCommand) activatePipelines(argv []string) (interface{}, error) {
	pipelineID, argv := head(argv)

	var args activatePipelinesRequestArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		return nil, err
	}
	activatePipelineRequestQueryParams := streams.ActivatePipelineRequest{}
	if args.AllowNonRestoredState != "" {
		activatePipelineRequestQueryParams.AllowNonRestoredState = streamsCommand.parseBooleanOrFail(args.AllowNonRestoredState, "activate-latest-version")
	}
	if args.ActivateLatestVersion != "" {
		activatePipelineRequestQueryParams.ActivateLatestVersion = streamsCommand.parseBooleanOrFail(args.ActivateLatestVersion, "allow-non-restored-state")
	}
	if args.SkipRestoreState != "" {
		activatePipelineRequestQueryParams.SkipRestoreState = streamsCommand.parseBooleanOrFail(args.SkipRestoreState, "skip-restorestate")
	}

	return streamsCommand.streamsService.ActivatePipeline(pipelineID, activatePipelineRequestQueryParams)
}

func argToInt32Ptr(argValue string) int {
	intValue, err := strconv.Atoi(argValue)
	if err != nil {
		fatal(err.Error())
	}
	result := int(intValue)
	return result
}

type compileDslArgs struct {
	Dsl     string `arg:"dsl"`
	DslFile string `arg:"dsl-file"`
}

func (streamsCommand *StreamsCommand) compileDSL(argv []string) (interface{}, error) {
	var args compileDslArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		return nil, err
	}

	fmt.Println(args.Dsl)
	dsl := args.Dsl

	if dsl == "" && args.DslFile != "" {
		dslBytes, err := ioutil.ReadFile(args.DslFile)
		if err != nil {
			fatal(err.Error())
		}
		dsl = string(dslBytes)
	}
	if len(dsl) == 0 {
		fatal("either -dsl or -dsl-file must be set")
	}

	dslCompilationRequest := streams.DslCompilationRequest{
		Dsl: dsl,
	}
	return streamsCommand.streamsService.CompileDSL(dslCompilationRequest)
}

type pipelineArgs struct {
	Name                     string `arg:"name"`
	Data                     string `arg:"data"`
	DataFile                 string `arg:"data-file"`
	Description              string `arg:"description"`
	BypassValidation         bool   `arg:"bypass-validation"`
	Activated                string `arg:"activated"`
	CreateUserID             string `arg:"create-user-id"`
	StreamingConfigurationID int64  `arg:"streaming-configuration-id"`
}

func (streamsCommand *StreamsCommand) createPipeline(argv []string) (interface{}, error) {
	pipelineRequest := streamsCommand.parseForCreate(argv)
	return streamsCommand.streamsService.CreatePipeline(pipelineRequest)
}

type deactivatePipelinesRequestArgs struct {
	SkipSavePoint string `arg:"skip-savepoint"`
}

func (streamsCommand *StreamsCommand) deactivatePipelines(argv []string) (interface{}, error) {
	pipelineID, argv := head(argv)

	var args deactivatePipelinesRequestArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		return nil, err
	}
	deactivatePipelineRequestQueryParams := streams.DeactivatePipelineRequest{}
	if args.SkipSavePoint != "" {
		deactivatePipelineRequestQueryParams.SkipSavepoint = streamsCommand.parseBooleanOrFail(args.SkipSavePoint, "skip-savepoint")
	}

	return streamsCommand.streamsService.DeactivatePipeline(pipelineID, deactivatePipelineRequestQueryParams)
}

func (streamsCommand *StreamsCommand) deletePipeline(argv []string) (interface{}, error) {
	id := head1(argv)
	return streamsCommand.streamsService.DeletePipeline(id)
}

func (streamsCommand *StreamsCommand) deletePreviewSession(argv []string) (interface{}, error) {
	id := head1(argv)
	sessionID, err := strconv.Atoi(id)
	if err != nil {
		fatal(err.Error())
	}
	return streamsCommand.streamsService.StopPreview(int64(sessionID))
}

type getSchemaArgs struct {
	NodeUUID string `arg:"0"`
	PortName string `arg:"1"`
	UplJSON  string `arg:"upl-json"`
	UplFile  string `arg:"upl-file"`
}

func (streamsCommand *StreamsCommand) getInputSchema(argv []string) (interface{}, error) {
	var args getSchemaArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		return nil, err
	}
	uplPipeline := streamsCommand.loadUplFromJSONOrFile(args.UplFile, args.UplJSON)

	return streamsCommand.streamsService.GetInputSchema(streams.GetInputSchemaRequest{NodeUuid: args.NodeUUID, TargetPortName: args.PortName, UplJson: uplPipeline})
}

func (streamsCommand *StreamsCommand) getLatestPipelineMetrics(argv []string) (interface{}, error) {
	pipelineID := head1(argv)
	return streamsCommand.streamsService.GetPipelineLatestMetrics(pipelineID)
}

func (streamsCommand *StreamsCommand) getLatestPreviewMetrics(argv []string) (interface{}, error) {
	previewSessionID := head1(argv)
	previewID, err := strconv.Atoi(previewSessionID)
	if err != nil {
		fatal(err.Error())
	}
	return streamsCommand.streamsService.GetPreviewSessionLatestMetrics(int64(previewID))
}

func (streamsCommand *StreamsCommand) getOutputSchema(argv []string) (interface{}, error) {
	var args getSchemaArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		return nil, err
	}
	uplPipeline := streamsCommand.loadUplFromJSONOrFile(args.UplFile, args.UplJSON)

	return streamsCommand.streamsService.GetOutputSchema(streams.GetOutputSchemaRequest{NodeUuid: &args.NodeUUID, SourcePortName: &args.PortName, UplJson: uplPipeline})
}

type getPipelineArgs struct {
	Version string `arg:"version"`
}

func (streamsCommand *StreamsCommand) getPipeline(argv []string) (interface{}, error) {
	id, argv := head(argv)
	var args getPipelineArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		fatal(err.Error())
	}
	return streamsCommand.streamsService.GetPipeline(id, &streams.GetPipelineQueryParams{Version: args.Version})
}

// GetPipelineStatusArgs contains the query parameters provided by the user through command line to fetch specific pipeline job statuses
type GetPipelineStatusArgs struct {
	// This struct is overloaded to be used for two purposes: parsing command
	// line args, and serializing the request into queryParams. The "arg" tag
	// is needed for the former, "key" for the latter.
	Offset       int32  `arg:"offset" key:"offset"`
	PageSize     int32  `arg:"page-size" key:"pageSize"`
	SortField    string `arg:"sort-field" key:"sortField"`
	SortDir      string `arg:"sort-dir" key:"sortDir"`
	Activated    bool   `arg:"activated" key:"activated"`
	CreateUserID string `arg:"create-user-id" key:"createUserID"`
	Name         string `arg:"name" key:"name"`
}

func (streamsCommand *StreamsCommand) getPipelineStatus(argv []string) (interface{}, error) {
	var args GetPipelineStatusArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		return nil, err
	}
	pipelineStatusQueryParams := streams.GetPipelinesStatusQueryParams{}
	if args.Offset != 0 {
		pipelineStatusQueryParams.Offset = &args.Offset
	}
	if args.PageSize != 0 {
		pipelineStatusQueryParams.PageSize = &args.PageSize
	}
	if args.SortField != "" {
		pipelineStatusQueryParams.SortField = args.SortField
	}
	if args.SortDir != "" {
		pipelineStatusQueryParams.SortDir = args.SortDir
	}
	if args.Activated {
		pipelineStatusQueryParams.Activated = &args.Activated
	}
	if args.CreateUserID != "" {
		pipelineStatusQueryParams.CreateUserId = args.CreateUserID
	}
	if args.Name != "" {
		pipelineStatusQueryParams.Name = args.Name
	}
	return streamsCommand.streamsService.GetPipelinesStatus(&pipelineStatusQueryParams)
}

type getPipelinesArgs struct {
	Offset       int32  `arg:"offset"`
	PageSize     int32  `arg:"page-size"`
	SortField    string `arg:"sort-field"`
	SortDir      string `arg:"sort-dir"`
	Activated    string `arg:"activated"`
	CreateUserID string `arg:"create-user-id"`
	Name         string `arg:"name"`
	IncludeData  string `arg:"include-data"`
}

func (streamsCommand *StreamsCommand) getPipelines(argv []string) (interface{}, error) {
	var args getPipelinesArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		return nil, err
	}
	pipelineQueryParams := streams.ListPipelinesQueryParams{}
	if args.SortField != "" {
		pipelineQueryParams.SortField = args.SortField
	}
	if args.SortDir != "" {
		pipelineQueryParams.SortDir = args.SortDir
	}
	if args.CreateUserID != "" {
		pipelineQueryParams.CreateUserId = args.CreateUserID
	}
	if args.Name != "" {
		pipelineQueryParams.Name = args.Name
	}
	if args.Offset != 0 {
		pipelineQueryParams.Offset = &args.Offset
	}
	if args.PageSize != 0 {
		pipelineQueryParams.PageSize = &args.PageSize
	}
	if args.Activated != "" {
		pipelineQueryParams.Activated = streamsCommand.parseBooleanOrFail(args.Activated, "activated")
	}
	if args.IncludeData != "" {
		pipelineQueryParams.IncludeData = streamsCommand.parseBooleanOrFail(args.IncludeData, "include-data")
	}
	return streamsCommand.streamsService.ListPipelines(&pipelineQueryParams)
}

func (streamsCommand *StreamsCommand) getPreviewData(argv []string) (interface{}, error) {
	arg1 := head1(argv)
	previewID, err := strconv.Atoi(arg1)
	if err != nil {
		fatal(err.Error())
	}
	return streamsCommand.streamsService.GetPreviewData(int64(previewID))
}

func (streamsCommand *StreamsCommand) getPreviewSession(argv []string) (interface{}, error) {
	arg1 := head1(argv)
	sessionID, err := strconv.Atoi(arg1)
	if err != nil {
		fatal(err.Error())
	}
	return streamsCommand.streamsService.GetPreviewSession(int64(sessionID))
}

type getRegistryArgs struct {
	Local string `arg:"local"`
}

func (streamsCommand *StreamsCommand) getRegistry(argv []string) (interface{}, error) {
	var args getRegistryArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		fatal(err.Error())
	}
	registryQueryParams := streams.GetRegistryQueryParams{}
	if args.Local != "" {
		registryQueryParams.Local = streamsCommand.parseBooleanOrFail(args.Local, "local")
	}
	return streamsCommand.streamsService.GetRegistry(&registryQueryParams)
}

func (streamsCommand *StreamsCommand) loadUplFromJSONOrFile(uplFile string, uplJSONString string) streams.UplPipeline {
	var uplPipeline streams.UplPipeline
	var UplJSON []byte
	var err error

	if uplFile != "" {
		UplJSON, err = ioutil.ReadFile(uplFile)
		if err != nil {
			fatal(err.Error())
		}
	} else {
		UplJSON = []byte(uplJSONString)
	}
	if len(UplJSON) == 0 {
		fatal("either -upl-json or -upl-file must be set")
	}
	err = json.Unmarshal([]byte(UplJSON), &uplPipeline)
	if err != nil {
		fatal("json unmarshal error: %s", err.Error())
	}
	return uplPipeline
}

func (streamsCommand *StreamsCommand) parseBooleanOrFail(includeData string, flagName string) *bool {
	helper, err := strconv.ParseBool(includeData)
	if err != nil {
		flag.Usage()
		fatal(fmt.Sprintf("bad value for %v: '%v'. Boolean expected.", flagName, includeData))
	}
	return &helper
}

func (streamsCommand *StreamsCommand) parseForReplace(argv []string) streams.PipelinePatchRequest {
	var args pipelineArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		fatal(err.Error())
	}
	pipelineRequest := streams.PipelinePatchRequest{}
	if args.Name != "" {
		pipelineRequest.Name = &args.Name
	}
	if args.Description != "" {
		pipelineRequest.Description = &args.Description
	}
	if args.CreateUserID != "" {
		pipelineRequest.CreateUserId = &args.CreateUserID
	}
	if args.BypassValidation {
		pipelineRequest.BypassValidation = &args.BypassValidation
	}
	if args.Data != "" {
		var uplPipeline streams.UplPipeline
		err = json.Unmarshal([]byte(args.Data), &uplPipeline)
		if err != nil {
			fatal("json unmarshal error: %s", err.Error())
		}
		pipelineRequest.Data = &uplPipeline
	} else if args.DataFile != "" {
		fileContents, err := ioutil.ReadFile(args.DataFile)
		if err != nil {
			fatal(err.Error())
		}
		var uplPipeline streams.UplPipeline
		err = json.Unmarshal(fileContents, &uplPipeline)
		if err != nil {
			fatal("json unmarshal error: %s", err.Error())
		}
		pipelineRequest.Data = &uplPipeline
	}
	return pipelineRequest
}

func (streamsCommand *StreamsCommand) parseForCreate(argv []string) streams.PipelineRequest {
	var args pipelineArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		fatal(err.Error())
	}
	pipelineRequest := streams.PipelineRequest{}
	if args.Name != "" {
		pipelineRequest.Name = args.Name
	}
	if args.Description != "" {
		pipelineRequest.Description = &args.Description
	}
	if args.BypassValidation {
		pipelineRequest.BypassValidation = &args.BypassValidation
	}
	if args.Data != "" {
		var uplPipeline streams.UplPipeline
		err = json.Unmarshal([]byte(args.Data), &uplPipeline)
		if err != nil {
			fatal("json unmarshal error: %s", err.Error())
		}
		pipelineRequest.Data = uplPipeline
	} else if args.DataFile != "" {
		fileContents, err := ioutil.ReadFile(args.DataFile)
		if err != nil {
			fatal(err.Error())
		}
		var uplPipeline streams.UplPipeline
		err = json.Unmarshal(fileContents, &uplPipeline)
		if err != nil {
			fatal("json unmarshal error: %s", err.Error())
		}
		pipelineRequest.Data = uplPipeline
	}
	return pipelineRequest
}

type mergePipelinesArgs struct {
	TargetNode    string `arg:"0"`
	TargetPort    string `arg:"1"`
	InputTreeJSON string `arg:"input-tree-json"`
	InputTreeFile string `arg:"input-tree-file"`
	MainTreeJSON  string `arg:"main-tree-json"`
	MainTreeFile  string `arg:"main-tree-file"`
}

func (streamsCommand *StreamsCommand) mergePipelines(argv []string) (interface{}, error) {
	var args mergePipelinesArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		fatal(err.Error())
	}

	inputPipeline := streamsCommand.loadUplFromJSONOrFile(args.InputTreeFile, args.InputTreeJSON)
	mainPipeline := streamsCommand.loadUplFromJSONOrFile(args.MainTreeFile, args.MainTreeJSON)
	request := &streams.PipelinesMergeRequest{
		TargetNode: args.TargetNode,
		TargetPort: args.TargetPort,
		InputTree:  inputPipeline,
		MainTree:   mainPipeline,
	}
	return streamsCommand.streamsService.MergePipelines(*request)
}

func (streamsCommand *StreamsCommand) reactivatePipeline(argv []string) (interface{}, error) {
	id := head1(argv)
	return streamsCommand.streamsService.ReactivatePipeline(id)
}

func (streamsCommand *StreamsCommand) replacePipeline(argv []string) (interface{}, error) {
	id, argv := head(argv)
	pipelineRequest := streamsCommand.parseForReplace(argv)
	return streamsCommand.streamsService.UpdatePipeline(id, pipelineRequest)
}

type startPreviewSessionArgs struct {
	RecordsLimit       string `arg:"records-limit"`
	SessionLifetimeMs  string `arg:"session-lifetime-ms"`
	RecordsPerPipeline string `arg:"records-per-pipeline"`
	UplJSON            string `arg:"upl-json"`
	UplFile            string `arg:"upl-file"`
	UseNewData         string `arg:"use-new-data"`
}

func (streamsCommand *StreamsCommand) startPreviewSession(argv []string) (interface{}, error) {
	var args startPreviewSessionArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		fatal(err.Error())
	}

	uplPipeline := streamsCommand.loadUplFromJSONOrFile(args.UplFile, args.UplJSON)
	request := &streams.PreviewSessionStartRequest{Upl: uplPipeline}
	if args.RecordsLimit != "" {
		temp := int32(argToInt32Ptr(args.RecordsLimit))
		request.RecordsLimit = &temp
	}
	if args.SessionLifetimeMs != "" {
		temp := int64(argToInt32Ptr(args.SessionLifetimeMs))
		request.SessionLifetimeMs = &temp
	}
	if args.RecordsPerPipeline != "" {
		temp := int32(argToInt32Ptr(args.RecordsPerPipeline))
		request.RecordsPerPipeline = &temp
	}
	if args.UseNewData != "" {
		value, err := strconv.ParseBool(args.UseNewData)
		if err != nil {
			fatal(err.Error())
		}
		request.UseNewData = &value
	}

	return streamsCommand.streamsService.StartPreview(*request)
}

type validateUplArgs struct {
	UplJSON string `arg:"upl-json"`
	UplFile string `arg:"upl-file"`
}

func (streamsCommand *StreamsCommand) validateUpl(argv []string) (interface{}, error) {
	var args validateUplArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		fatal(err.Error())
	}

	uplPipeline := streamsCommand.loadUplFromJSONOrFile(args.UplFile, args.UplJSON)

	return streamsCommand.streamsService.ValidatePipeline(streams.ValidateRequest{Upl: uplPipeline})
}

func (streamsCommand *StreamsCommand) getSpecJSON(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return GetSpecJSON("api", StreamsServiceVersion, "streams", streamsCommand.streamsService.Client)
}

func (streamsCommand *StreamsCommand) getSpecYaml(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return GetSpecYaml("api", StreamsServiceVersion, "streams", streamsCommand.streamsService.Client)
}

func (streamsCommand *StreamsCommand) getConnectors(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return streamsCommand.streamsService.ListConnectors()
}

func (streamsCommand *StreamsCommand) getConnections(argv []string) (interface{}, error) {
	var functionID string
	var createUserID string
	var connectorID string
	var name string
	var offset string
	var pageSize string
	var sortField string
	var sortDir string
	var secretName string
	var offsetIntVal int
	var offsetVal int32
	var pageSizeIntVal int
	var pageSizeVal int32

	flags := flag.NewFlagSet("list-connections", flag.ExitOnError)
	flags.StringVar(&functionID, "function-id", "", "Group function id")
	flags.StringVar(&createUserID, "create-user-id", "", "Filter by create user id")
	flags.StringVar(&connectorID, "connector-id", "", "Connection id")
	flags.StringVar(&name, "name", "", "Connection name")
	flags.StringVar(&offset, "offset", "", "Offset")
	flags.StringVar(&pageSize, "page-size", "", "Page size")
	flags.StringVar(&sortField, "sort-field", "", "Keys to sort by")
	flags.StringVar(&sortDir, "sort-dir", "", "Sort dir")
	flags.StringVar(&secretName, "show-secret-name", "", "Show secret name")
	err := flags.Parse(argv) //nolint:errcheck
	if err != nil {
		return nil, err
	}

	fsArgs := flags.Args()
	if len(fsArgs) > 0 {
		fatal("unexpected argument(s): %s", fsArgs)
	}
	if offset != "" {
		offsetIntVal, err = strconv.Atoi(offset)
		offsetVal = int32(offsetIntVal)
		if err != nil {
			fatal(err.Error())
		}
	}
	if pageSize != "" {
		pageSizeIntVal, err = strconv.Atoi(pageSize)
		pageSizeVal = int32(pageSizeIntVal)
		if err != nil {
			fatal(err.Error())
		}
	}
	return streamsCommand.streamsService.ListConnections(&streams.ListConnectionsQueryParams{FunctionId: functionID, CreateUserId: createUserID, ConnectorId: connectorID, Name: name, Offset: &offsetVal, PageSize: &pageSizeVal, SortField: sortField, SortDir: sortDir, ShowSecretNames: secretName})
}

// TemplateArgs contains the create/update template request data passed by user as command line arguments
type TemplateArgs struct {
	Name        string `arg:"name"`
	Data        string `arg:"data"`
	DataFile    string `arg:"data-file"`
	Description string `arg:"description"`
}

func (streamsCommand *StreamsCommand) createTemplate(argv []string) (interface{}, error) {
	templateRequest := streamsCommand.parseForCreateTemplate(argv)
	return streamsCommand.streamsService.CreateTemplate(templateRequest)
}

func (streamsCommand *StreamsCommand) replaceTemplate(argv []string) (interface{}, error) {
	id, argv := head(argv)
	templateRequest := streamsCommand.parseForPutTemplate(argv)
	return streamsCommand.streamsService.PutTemplate(id, templateRequest)
}

func (streamsCommand *StreamsCommand) updateTemplatePartially(argv []string) (interface{}, error) {
	id, argv := head(argv)
	templateRequest := streamsCommand.parseForUpdateTemplatePartially(argv)
	return streamsCommand.streamsService.UpdateTemplate(id, templateRequest)
}

func (streamsCommand *StreamsCommand) getTemplates(argv []string) (interface{}, error) {
	var args ListTemplatesQueryParams
	_, err := argx.Parse(argv, &args)
	if err != nil {
		fatal(err.Error())
	}
	listTemplateQueryParams := streams.ListTemplatesQueryParams{}
	if args.Offset != 0 {
		listTemplateQueryParams.Offset = &args.Offset
	}
	if args.PageSize != 0 {
		listTemplateQueryParams.PageSize = &args.PageSize
	}
	if args.SortDir != "" {
		listTemplateQueryParams.SortDir = args.SortDir
	}

	if args.SortField != "" {
		listTemplateQueryParams.SortField = args.SortField
	}

	return streamsCommand.streamsService.ListTemplates(&listTemplateQueryParams)
}

type getTemplateQueryParamsArgs struct {
	Version int64 `arg:"version"`
}

type ListTemplatesQueryParams struct {
	Offset    int32  `arg:"offset"`
	PageSize  int32  `arg:"page-size"`
	SortDir   string `arg:"sort-dir"`
	SortField string `arg:"sort-field"`
}

func (streamsCommand *StreamsCommand) getTemplate(argv []string) (interface{}, error) {
	id, argv := head(argv)
	var args getTemplateQueryParamsArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		fatal(err.Error())
	}
	return streamsCommand.streamsService.GetTemplate(id, &streams.GetTemplateQueryParams{Version: &args.Version})
}

func (streamsCommand *StreamsCommand) deleteTemplate(argv []string) error {
	id := head1(argv)
	return streamsCommand.streamsService.DeleteTemplate(id)
}

// GroupExpandArgs contains data to create the expanded version of group, provided by user as command line arguments
type GroupExpandArgs struct {
	Arguments       string `arg:"arguments"`
	ArgumentsFile   string `arg:"arguments-file"`
	GroupID         string `arg:"group-id"`
	GroupFunctionID string `arg:"group-function-id"`
}

func (streamsCommand *StreamsCommand) createExpandedGroup(argv []string) (interface{}, error) {
	var args GroupExpandArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		fatal(err.Error())
	}
	type group map[string]interface{}
	arguments := group{}

	if args.Arguments != "" {
		var arguments map[string]interface{}
		err = json.Unmarshal([]byte(args.Arguments), &arguments)
		if err != nil {
			fatal("json unmarshal error: %s", err.Error())
		}
	} else if args.ArgumentsFile != "" {
		fileContents, err := ioutil.ReadFile(args.ArgumentsFile)
		if err != nil {
			fatal(err.Error())
		}
		var arguments map[string]interface{}
		err = json.Unmarshal(fileContents, &arguments)
		if err != nil {
			fatal("json unmarshal error: %s", err.Error())
		}
	}

	return streamsCommand.streamsService.ExpandGroup(args.GroupID, streams.GroupExpandRequest{Arguments: arguments, Id: args.GroupFunctionID})
}

func (streamsCommand *StreamsCommand) getGroupbyID(argv []string) (interface{}, error) {
	id := head1(argv)
	return streamsCommand.streamsService.GetGroup(id)
}

func (streamsCommand *StreamsCommand) parseForCreateTemplate(argv []string) streams.TemplateRequest {
	var args TemplateArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		fatal(err.Error())
	}
	templateRequest := streams.TemplateRequest{}
	if args.Name != "" {
		templateRequest.Name = args.Name
	}
	if args.Description != "" {
		templateRequest.Description = args.Description
	}
	if args.Data != "" {
		var uplPipeline streams.UplPipeline
		err = json.Unmarshal([]byte(args.Data), &uplPipeline)
		if err != nil {
			fatal("json unmarshal error: %s", err.Error())
		}
		templateRequest.Data = uplPipeline
	} else if args.DataFile != "" {
		fileContents, err := ioutil.ReadFile(args.DataFile)
		if err != nil {
			fatal(err.Error())
		}
		var uplPipeline streams.UplPipeline
		err = json.Unmarshal(fileContents, &uplPipeline)
		if err != nil {
			fatal("json unmarshal error: %s", err.Error())
		}
		templateRequest.Data = uplPipeline
	}
	return templateRequest
}

func (streamsCommand *StreamsCommand) parseForPutTemplate(argv []string) streams.TemplatePutRequest {
	var args TemplateArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		fatal(err.Error())
	}
	templateRequest := streams.TemplatePutRequest{}
	if args.Name != "" {
		templateRequest.Name = args.Name
	}
	if args.Description != "" {
		templateRequest.Description = args.Description
	}
	if args.Data != "" {
		var uplPipeline streams.UplPipeline
		err = json.Unmarshal([]byte(args.Data), &uplPipeline)
		if err != nil {
			fatal("json unmarshal error: %s", err.Error())
		}
		templateRequest.Data = uplPipeline
	} else if args.DataFile != "" {
		fileContents, err := ioutil.ReadFile(args.DataFile)
		if err != nil {
			fatal(err.Error())
		}
		var uplPipeline streams.UplPipeline
		err = json.Unmarshal(fileContents, &uplPipeline)
		if err != nil {
			fatal("json unmarshal error: %s", err.Error())
		}
		templateRequest.Data = uplPipeline
	}
	return templateRequest
}

func (streamsCommand *StreamsCommand) parseForUpdateTemplatePartially(argv []string) streams.TemplatePatchRequest {
	var args TemplateArgs
	_, err := argx.Parse(argv, &args)
	if err != nil {
		fatal(err.Error())
	}
	templateRequest := streams.TemplatePatchRequest{}

	if args.Description != "" {
		templateRequest.Description = &args.Description
	}
	if args.Name != "" {
		templateRequest.Name = &args.Name
	}

	if args.Data != "" {
		var uplPipeline streams.UplPipeline
		err = json.Unmarshal([]byte(args.Data), &uplPipeline)
		if err != nil {
			fatal("json unmarshal error: %s", err.Error())
		}
		templateRequest.Data = &uplPipeline
	} else if args.DataFile != "" {
		fileContents, err := ioutil.ReadFile(args.DataFile)
		if err != nil {
			fatal(err.Error())
		}
		var uplPipeline streams.UplPipeline
		err = json.Unmarshal(fileContents, &uplPipeline)
		if err != nil {
			fatal("json unmarshal error: %s", err.Error())
		}
		templateRequest.Data = &uplPipeline
	}
	return templateRequest
}
