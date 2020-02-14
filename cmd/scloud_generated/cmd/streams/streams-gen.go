// Package streams -- generated by scloudgen
// !! DO NOT EDIT !!
//
package streams

import (
	"github.com/spf13/cobra"
	impl "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/pkg/streams"
)

// activatePipeline -- Activates an existing pipeline.
var activatePipelineCmd = &cobra.Command{
	Use:   "activate-pipeline",
	Short: "Activates an existing pipeline.",
	RunE:  impl.ActivatePipeline,
}

// compileDSL -- Compiles the Streams DSL and returns Streams JSON.
var compileDSLCmd = &cobra.Command{
	Use:   "compile-dsl",
	Short: "Compiles the Streams DSL and returns Streams JSON.",
	RunE:  impl.CompileDSL,
}

// compileSPL -- Compiles SPL2 and returns Streams JSON.
var compileSPLCmd = &cobra.Command{
	Use:   "compile-spl",
	Short: "Compiles SPL2 and returns Streams JSON.",
	RunE:  impl.CompileSPL,
}

// createConnection -- Create a new DSP connection.
var createConnectionCmd = &cobra.Command{
	Use:   "create-connection",
	Short: "Create a new DSP connection.",
	RunE:  impl.CreateConnection,
}

// createGroup -- Create a new group function by combining the Streams JSON of two or more functions.
var createGroupCmd = &cobra.Command{
	Use:   "create-group",
	Short: "Create a new group function by combining the Streams JSON of two or more functions.",
	RunE:  impl.CreateGroup,
}

// createPipeline -- Creates a pipeline.
var createPipelineCmd = &cobra.Command{
	Use:   "create-pipeline",
	Short: "Creates a pipeline.",
	RunE:  impl.CreatePipeline,
}

// createTemplate -- Creates a template for a tenant.
var createTemplateCmd = &cobra.Command{
	Use:   "create-template",
	Short: "Creates a template for a tenant.",
	RunE:  impl.CreateTemplate,
}

// deactivatePipeline -- Deactivates an existing pipeline.
var deactivatePipelineCmd = &cobra.Command{
	Use:   "deactivate-pipeline",
	Short: "Deactivates an existing pipeline.",
	RunE:  impl.DeactivatePipeline,
}

// deleteConnection -- Delete all versions of a connection by its id.
var deleteConnectionCmd = &cobra.Command{
	Use:   "delete-connection",
	Short: "Delete all versions of a connection by its id.",
	RunE:  impl.DeleteConnection,
}

// deleteGroup -- Removes an existing group.
var deleteGroupCmd = &cobra.Command{
	Use:   "delete-group",
	Short: "Removes an existing group.",
	RunE:  impl.DeleteGroup,
}

// deletePipeline -- Removes a pipeline.
var deletePipelineCmd = &cobra.Command{
	Use:   "delete-pipeline",
	Short: "Removes a pipeline.",
	RunE:  impl.DeletePipeline,
}

// deleteTemplate -- Removes a template with a specific ID.
var deleteTemplateCmd = &cobra.Command{
	Use:   "delete-template",
	Short: "Removes a template with a specific ID.",
	RunE:  impl.DeleteTemplate,
}

// expandGroup -- Creates and returns the expanded version of a group.
var expandGroupCmd = &cobra.Command{
	Use:   "expand-group",
	Short: "Creates and returns the expanded version of a group.",
	RunE:  impl.ExpandGroup,
}

// expandPipeline -- Returns the entire Streams JSON, including the expanded Streams JSON of any group functions in the pipeline.
var expandPipelineCmd = &cobra.Command{
	Use:   "expand-pipeline",
	Short: "Returns the entire Streams JSON, including the expanded Streams JSON of any group functions in the pipeline.",
	RunE:  impl.ExpandPipeline,
}

// getGroup -- Returns the full Streams JSON of a group.
var getGroupCmd = &cobra.Command{
	Use:   "get-group",
	Short: "Returns the full Streams JSON of a group.",
	RunE:  impl.GetGroup,
}

// getInputSchema -- Returns the input schema for a function in a pipeline.
var getInputSchemaCmd = &cobra.Command{
	Use:   "get-input-schema",
	Short: "Returns the input schema for a function in a pipeline.",
	RunE:  impl.GetInputSchema,
}

// getOutputSchema -- Returns the output schema for a specified function in a pipeline. If no function ID is  specified, the request returns the output schema for all functions in a pipeline.
var getOutputSchemaCmd = &cobra.Command{
	Use:   "get-output-schema",
	Short: "Returns the output schema for a specified function in a pipeline. If no function ID is  specified, the request returns the output schema for all functions in a pipeline.",
	RunE:  impl.GetOutputSchema,
}

// getPipeline -- Returns an individual pipeline by version.
var getPipelineCmd = &cobra.Command{
	Use:   "get-pipeline",
	Short: "Returns an individual pipeline by version.",
	RunE:  impl.GetPipeline,
}

// getPipelineLatestMetrics -- Returns the latest metrics for a single pipeline.
var getPipelineLatestMetricsCmd = &cobra.Command{
	Use:   "get-pipeline-latest-metrics",
	Short: "Returns the latest metrics for a single pipeline.",
	RunE:  impl.GetPipelineLatestMetrics,
}

// getPipelinesStatus -- Returns the status of pipelines from the underlying streaming system.
var getPipelinesStatusCmd = &cobra.Command{
	Use:   "get-pipelines-status",
	Short: "Returns the status of pipelines from the underlying streaming system.",
	RunE:  impl.GetPipelinesStatus,
}

// getPreviewData -- Returns the preview data for a session.
var getPreviewDataCmd = &cobra.Command{
	Use:   "get-preview-data",
	Short: "Returns the preview data for a session.",
	RunE:  impl.GetPreviewData,
}

// getPreviewSession -- Returns information from a preview session.
var getPreviewSessionCmd = &cobra.Command{
	Use:   "get-preview-session",
	Short: "Returns information from a preview session.",
	RunE:  impl.GetPreviewSession,
}

// getPreviewSessionLatestMetrics -- Returns the latest metrics for a preview session.
var getPreviewSessionLatestMetricsCmd = &cobra.Command{
	Use:   "get-preview-session-latest-metrics",
	Short: "Returns the latest metrics for a preview session.",
	RunE:  impl.GetPreviewSessionLatestMetrics,
}

// getRegistry -- Returns all functions in JSON format.
var getRegistryCmd = &cobra.Command{
	Use:   "get-registry",
	Short: "Returns all functions in JSON format.",
	RunE:  impl.GetRegistry,
}

// getTemplate -- Returns an individual template by version.
var getTemplateCmd = &cobra.Command{
	Use:   "get-template",
	Short: "Returns an individual template by version.",
	RunE:  impl.GetTemplate,
}

// listConnections -- Returns a list of connections (latest versions only) by tenant ID.
var listConnectionsCmd = &cobra.Command{
	Use:   "list-connections",
	Short: "Returns a list of connections (latest versions only) by tenant ID.",
	RunE:  impl.ListConnections,
}

// listConnectors -- Returns a list of the available connectors.
var listConnectorsCmd = &cobra.Command{
	Use:   "list-connectors",
	Short: "Returns a list of the available connectors.",
	RunE:  impl.ListConnectors,
}

// listPipelines -- Returns all pipelines.
var listPipelinesCmd = &cobra.Command{
	Use:   "list-pipelines",
	Short: "Returns all pipelines.",
	RunE:  impl.ListPipelines,
}

// listTemplates -- Returns a list of all templates.
var listTemplatesCmd = &cobra.Command{
	Use:   "list-templates",
	Short: "Returns a list of all templates.",
	RunE:  impl.ListTemplates,
}

// mergePipelines -- Combines two Streams JSON programs.
var mergePipelinesCmd = &cobra.Command{
	Use:   "merge-pipelines",
	Short: "Combines two Streams JSON programs.",
	RunE:  impl.MergePipelines,
}

// putConnection -- Modifies an existing DSP connection.
var putConnectionCmd = &cobra.Command{
	Use:   "put-connection",
	Short: "Modifies an existing DSP connection.",
	RunE:  impl.PutConnection,
}

// putGroup -- Update a group function combining the Streams JSON of two or more functions.
var putGroupCmd = &cobra.Command{
	Use:   "put-group",
	Short: "Update a group function combining the Streams JSON of two or more functions.",
	RunE:  impl.PutGroup,
}

// putTemplate -- Updates an existing template.
var putTemplateCmd = &cobra.Command{
	Use:   "put-template",
	Short: "Updates an existing template.",
	RunE:  impl.PutTemplate,
}

// reactivatePipeline -- Reactivate a pipeline
var reactivatePipelineCmd = &cobra.Command{
	Use:   "reactivate-pipeline",
	Short: "Reactivate a pipeline",
	RunE:  impl.ReactivatePipeline,
}

// startPreview -- Creates a preview session for a pipeline.
var startPreviewCmd = &cobra.Command{
	Use:   "start-preview",
	Short: "Creates a preview session for a pipeline.",
	RunE:  impl.StartPreview,
}

// stopPreview -- Stops a preview session.
var stopPreviewCmd = &cobra.Command{
	Use:   "stop-preview",
	Short: "Stops a preview session.",
	RunE:  impl.StopPreview,
}

// updateConnection -- Partially modifies an existing DSP connection.
var updateConnectionCmd = &cobra.Command{
	Use:   "update-connection",
	Short: "Partially modifies an existing DSP connection.",
	RunE:  impl.UpdateConnection,
}

// updateGroup -- Modify a group function by combining the Streams JSON of two or more functions.
var updateGroupCmd = &cobra.Command{
	Use:   "update-group",
	Short: "Modify a group function by combining the Streams JSON of two or more functions.",
	RunE:  impl.UpdateGroup,
}

// updatePipeline -- Partially modifies an existing pipeline.
var updatePipelineCmd = &cobra.Command{
	Use:   "update-pipeline",
	Short: "Partially modifies an existing pipeline.",
	RunE:  impl.UpdatePipeline,
}

// updateTemplate -- Partially modifies an existing template.
var updateTemplateCmd = &cobra.Command{
	Use:   "update-template",
	Short: "Partially modifies an existing template.",
	RunE:  impl.UpdateTemplate,
}

// validatePipeline -- Verifies whether the Streams JSON is valid.
var validatePipelineCmd = &cobra.Command{
	Use:   "validate-pipeline",
	Short: "Verifies whether the Streams JSON is valid.",
	RunE:  impl.ValidatePipeline,
}

func init() {
	streamsCmd.AddCommand(activatePipelineCmd)

	var activatePipelineId string
	activatePipelineCmd.Flags().StringVar(&activatePipelineId, "id", "", "This is a required parameter. id of the pipeline to activate")
	activatePipelineCmd.MarkFlagRequired("id")

	var activatePipelineActivateLatestVersion string
	activatePipelineCmd.Flags().StringVar(&activatePipelineActivateLatestVersion, "activate-latest-version", "false", "Set to true to activate the latest version of the pipeline. Set to false to use the previously activated version of the pipeline. Defaults to true.")

	var activatePipelineAllowNonRestoredState string
	activatePipelineCmd.Flags().StringVar(&activatePipelineAllowNonRestoredState, "allow-non-restored-state", "false", "Set to true to allow the pipeline to ignore any unused progress states. In some cases, when a data pipeline is changed, the progress state will be stored for functions that no longer exist, so this must be set to activate a pipeline in this state. Defaults to false.")

	var activatePipelineSkipRestoreState string
	activatePipelineCmd.Flags().StringVar(&activatePipelineSkipRestoreState, "skip-restore-state", "false", "Set to true to start reading from the latest input rather than from where the pipeline's previous run left off, which can cause data loss. Defaults to false.")

	streamsCmd.AddCommand(compileDSLCmd)

	var compileDSLInputDatafile string
	compileDSLCmd.Flags().StringVar(&compileDSLInputDatafile, "input-datafile", "", "The input data file.")

	streamsCmd.AddCommand(compileSPLCmd)

	var compileSPLSpl string
	compileSPLCmd.Flags().StringVar(&compileSPLSpl, "spl", "", "This is a required parameter. The SPL2 representation of a pipeline or function parameters.")
	compileSPLCmd.MarkFlagRequired("spl")

	var compileSPLSyntax string
	compileSPLCmd.Flags().StringVar(&compileSPLSyntax, "syntax", "", "The parse parameters as arguments to this SPL2 command can accept values UPL, DSL, SPL, EVAL, WHERE, TIMECHART, FIELDS, MVEXPAND, REX, BIN, RENAME, STATS, STATS_BY, SELECT, EXPRESSION, FUNCTION, LITERAL, UNKNOWN")

	streamsCmd.AddCommand(createConnectionCmd)

	var createConnectionConnectorId string
	createConnectionCmd.Flags().StringVar(&createConnectionConnectorId, "connector-id", "", "This is a required parameter. The ID of the parent connector.")
	createConnectionCmd.MarkFlagRequired("connector-id")

	var createConnectionData string
	createConnectionCmd.Flags().StringVar(&createConnectionData, "data", "", "This is a required parameter. The key-value pairs of configurations for this connection. Connectors may have some configurations that are required, which all connections must provide values for. For configuration values of type BYTES, the provided values must be Base64 encoded.")
	createConnectionCmd.MarkFlagRequired("data")

	var createConnectionDescription string
	createConnectionCmd.Flags().StringVar(&createConnectionDescription, "description", "", "This is a required parameter. The description of the connection.")
	createConnectionCmd.MarkFlagRequired("description")

	var createConnectionName string
	createConnectionCmd.Flags().StringVar(&createConnectionName, "name", "", "This is a required parameter. The name of the connection.")
	createConnectionCmd.MarkFlagRequired("name")

	streamsCmd.AddCommand(createGroupCmd)

	var createGroupArguments string
	createGroupCmd.Flags().StringVar(&createGroupArguments, "arguments", "", "This is a required parameter. The group function arguments list.")
	createGroupCmd.MarkFlagRequired("arguments")

	var createGroupAttributes string
	createGroupCmd.Flags().StringVar(&createGroupAttributes, "attributes", "", "This is a required parameter. The attributes map for function.")
	createGroupCmd.MarkFlagRequired("attributes")

	var createGroupCategories []int
	createGroupCmd.Flags().IntSliceVar(&createGroupCategories, "categories", nil, "This is a required parameter. The categories for this function.")
	createGroupCmd.MarkFlagRequired("categories")

	var createGroupEdges string
	createGroupCmd.Flags().StringVar(&createGroupEdges, "edges", "", "This is a required parameter. A list of links or connections between the output of one pipeline function and the input of another pipeline function")
	createGroupCmd.MarkFlagRequired("edges")

	var createGroupMappings string
	createGroupCmd.Flags().StringVar(&createGroupMappings, "mappings", "", "This is a required parameter. The group function mappings list.")
	createGroupCmd.MarkFlagRequired("mappings")

	var createGroupName string
	createGroupCmd.Flags().StringVar(&createGroupName, "name", "", "This is a required parameter. The group function name.")
	createGroupCmd.MarkFlagRequired("name")

	var createGroupNodes string
	createGroupCmd.Flags().StringVar(&createGroupNodes, "nodes", "", "This is a required parameter. The functions (or nodes) in your entire pipeline, including each function's operations, attributes, and properties")
	createGroupCmd.MarkFlagRequired("nodes")

	var createGroupOutputType string
	createGroupCmd.Flags().StringVar(&createGroupOutputType, "output-type", "", "This is a required parameter. The data type of the function's output.")
	createGroupCmd.MarkFlagRequired("output-type")

	var createGroupRootNode []string
	createGroupCmd.Flags().StringSliceVar(&createGroupRootNode, "root-node", nil, "This is a required parameter. The UUIDs of all sink functions in a given pipeline")
	createGroupCmd.MarkFlagRequired("root-node")

	var createGroupScalar string
	createGroupCmd.Flags().StringVar(&createGroupScalar, "scalar", "false", "")

	var createGroupVariadic string
	createGroupCmd.Flags().StringVar(&createGroupVariadic, "variadic", "false", "")

	streamsCmd.AddCommand(createPipelineCmd)

	var createPipelineInputDatafile string
	createPipelineCmd.Flags().StringVar(&createPipelineInputDatafile, "input-datafile", "", "The input data file.")

	streamsCmd.AddCommand(createTemplateCmd)

	var createTemplateInputDatafile string
	createTemplateCmd.Flags().StringVar(&createTemplateInputDatafile, "input-datafile", "", "The input data file.")

	streamsCmd.AddCommand(deactivatePipelineCmd)

	var deactivatePipelineId string
	deactivatePipelineCmd.Flags().StringVar(&deactivatePipelineId, "id", "", "This is a required parameter. id of the pipeline to deactivate")
	deactivatePipelineCmd.MarkFlagRequired("id")

	var deactivatePipelineSkipSavepoint string
	deactivatePipelineCmd.Flags().StringVar(&deactivatePipelineSkipSavepoint, "skip-savepoint", "false", "Set to true to skip saving the state of a deactivated pipeline. When the pipeline is later activated, it will start with the newest data and skip any data that arrived after this deactivation, which can cause data loss. Defaults to false.")

	streamsCmd.AddCommand(deleteConnectionCmd)

	var deleteConnectionConnectionId string
	deleteConnectionCmd.Flags().StringVar(&deleteConnectionConnectionId, "connection-id", "", "This is a required parameter. ID of the connection")
	deleteConnectionCmd.MarkFlagRequired("connection-id")

	streamsCmd.AddCommand(deleteGroupCmd)

	var deleteGroupGroupId string
	deleteGroupCmd.Flags().StringVar(&deleteGroupGroupId, "group-id", "", "This is a required parameter. The group function's ID from the function registry")
	deleteGroupCmd.MarkFlagRequired("group-id")

	streamsCmd.AddCommand(deletePipelineCmd)

	var deletePipelineId string
	deletePipelineCmd.Flags().StringVar(&deletePipelineId, "id", "", "This is a required parameter. id of the pipeline to delete")
	deletePipelineCmd.MarkFlagRequired("id")

	streamsCmd.AddCommand(deleteTemplateCmd)

	var deleteTemplateTemplateId string
	deleteTemplateCmd.Flags().StringVar(&deleteTemplateTemplateId, "template-id", "", "This is a required parameter. ID of the template to delete")
	deleteTemplateCmd.MarkFlagRequired("template-id")

	streamsCmd.AddCommand(expandGroupCmd)

	var expandGroupArguments string
	expandGroupCmd.Flags().StringVar(&expandGroupArguments, "arguments", "", "This is a required parameter. Function arguments for the given id. Overrides default values.")
	expandGroupCmd.MarkFlagRequired("arguments")

	var expandGroupGroupId string
	expandGroupCmd.Flags().StringVar(&expandGroupGroupId, "group-id", "", "This is a required parameter. The group function's ID from the function registry")
	expandGroupCmd.MarkFlagRequired("group-id")

	var expandGroupId string
	expandGroupCmd.Flags().StringVar(&expandGroupId, "id", "", "This is a required parameter. The ID associated with your group function in the pipeline Streams JSON")
	expandGroupCmd.MarkFlagRequired("id")

	streamsCmd.AddCommand(expandPipelineCmd)

	var expandPipelineInputDatafile string
	expandPipelineCmd.Flags().StringVar(&expandPipelineInputDatafile, "input-datafile", "", "The input data file.")

	streamsCmd.AddCommand(getGroupCmd)

	var getGroupGroupId string
	getGroupCmd.Flags().StringVar(&getGroupGroupId, "group-id", "", "This is a required parameter. The group function's ID from the function registry")
	getGroupCmd.MarkFlagRequired("group-id")

	streamsCmd.AddCommand(getInputSchemaCmd)

	var getInputSchemaEdges string
	getInputSchemaCmd.Flags().StringVar(&getInputSchemaEdges, "edges", "", "This is a required parameter. A list of links or connections between the output of one pipeline function and the input of another pipeline function")
	getInputSchemaCmd.MarkFlagRequired("edges")

	var getInputSchemaNodeUuid string
	getInputSchemaCmd.Flags().StringVar(&getInputSchemaNodeUuid, "node-uuid", "", "This is a required parameter. The function ID.")
	getInputSchemaCmd.MarkFlagRequired("node-uuid")

	var getInputSchemaNodes string
	getInputSchemaCmd.Flags().StringVar(&getInputSchemaNodes, "nodes", "", "This is a required parameter. The functions (or nodes) in your entire pipeline, including each function's operations, attributes, and properties")
	getInputSchemaCmd.MarkFlagRequired("nodes")

	var getInputSchemaRootNode []string
	getInputSchemaCmd.Flags().StringSliceVar(&getInputSchemaRootNode, "root-node", nil, "This is a required parameter. The UUIDs of all sink functions in a given pipeline")
	getInputSchemaCmd.MarkFlagRequired("root-node")

	var getInputSchemaTargetPortName string
	getInputSchemaCmd.Flags().StringVar(&getInputSchemaTargetPortName, "target-port-name", "", "This is a required parameter. The name of the input port.")
	getInputSchemaCmd.MarkFlagRequired("target-port-name")

	streamsCmd.AddCommand(getOutputSchemaCmd)

	var getOutputSchemaEdges string
	getOutputSchemaCmd.Flags().StringVar(&getOutputSchemaEdges, "edges", "", "This is a required parameter. A list of links or connections between the output of one pipeline function and the input of another pipeline function")
	getOutputSchemaCmd.MarkFlagRequired("edges")

	var getOutputSchemaNodes string
	getOutputSchemaCmd.Flags().StringVar(&getOutputSchemaNodes, "nodes", "", "This is a required parameter. The functions (or nodes) in your entire pipeline, including each function's operations, attributes, and properties")
	getOutputSchemaCmd.MarkFlagRequired("nodes")

	var getOutputSchemaRootNode []string
	getOutputSchemaCmd.Flags().StringSliceVar(&getOutputSchemaRootNode, "root-node", nil, "This is a required parameter. The UUIDs of all sink functions in a given pipeline")
	getOutputSchemaCmd.MarkFlagRequired("root-node")

	var getOutputSchemaNodeUuid string
	getOutputSchemaCmd.Flags().StringVar(&getOutputSchemaNodeUuid, "node-uuid", "", "The function ID. If omitted, returns the output schema for all functions.")

	var getOutputSchemaSourcePortName string
	getOutputSchemaCmd.Flags().StringVar(&getOutputSchemaSourcePortName, "source-port-name", "", "The name of the output port. Deprecated.")

	streamsCmd.AddCommand(getPipelineCmd)

	var getPipelineId string
	getPipelineCmd.Flags().StringVar(&getPipelineId, "id", "", "This is a required parameter. id of the pipeline to get")
	getPipelineCmd.MarkFlagRequired("id")

	var getPipelineVersion string
	getPipelineCmd.Flags().StringVar(&getPipelineVersion, "version", "", "version")

	streamsCmd.AddCommand(getPipelineLatestMetricsCmd)

	var getPipelineLatestMetricsId string
	getPipelineLatestMetricsCmd.Flags().StringVar(&getPipelineLatestMetricsId, "id", "", "This is a required parameter. ID of the pipeline to get metrics for")
	getPipelineLatestMetricsCmd.MarkFlagRequired("id")

	streamsCmd.AddCommand(getPipelinesStatusCmd)

	var getPipelinesStatusActivated string
	getPipelinesStatusCmd.Flags().StringVar(&getPipelinesStatusActivated, "activated", "false", "activated")

	var getPipelinesStatusCreateUserId string
	getPipelinesStatusCmd.Flags().StringVar(&getPipelinesStatusCreateUserId, "create-user-id", "", "createUserId")

	var getPipelinesStatusName string
	getPipelinesStatusCmd.Flags().StringVar(&getPipelinesStatusName, "name", "", "name")

	var getPipelinesStatusOffset int32
	getPipelinesStatusCmd.Flags().Int32Var(&getPipelinesStatusOffset, "offset", 0, "offset")

	var getPipelinesStatusPageSize int32
	getPipelinesStatusCmd.Flags().Int32Var(&getPipelinesStatusPageSize, "page-size", 0, "pageSize")

	var getPipelinesStatusSortDir string
	getPipelinesStatusCmd.Flags().StringVar(&getPipelinesStatusSortDir, "sort-dir", "", "sortDir")

	var getPipelinesStatusSortField string
	getPipelinesStatusCmd.Flags().StringVar(&getPipelinesStatusSortField, "sort-field", "", "sortField")

	streamsCmd.AddCommand(getPreviewDataCmd)

	var getPreviewDataPreviewSessionId int64
	getPreviewDataCmd.Flags().Int64Var(&getPreviewDataPreviewSessionId, "preview-session-id", 0, "This is a required parameter. ID of the preview session")
	getPreviewDataCmd.MarkFlagRequired("preview-session-id")

	streamsCmd.AddCommand(getPreviewSessionCmd)

	var getPreviewSessionPreviewSessionId int64
	getPreviewSessionCmd.Flags().Int64Var(&getPreviewSessionPreviewSessionId, "preview-session-id", 0, "This is a required parameter. ID of the preview session")
	getPreviewSessionCmd.MarkFlagRequired("preview-session-id")

	streamsCmd.AddCommand(getPreviewSessionLatestMetricsCmd)

	var getPreviewSessionLatestMetricsPreviewSessionId int64
	getPreviewSessionLatestMetricsCmd.Flags().Int64Var(&getPreviewSessionLatestMetricsPreviewSessionId, "preview-session-id", 0, "This is a required parameter. ID of the preview session")
	getPreviewSessionLatestMetricsCmd.MarkFlagRequired("preview-session-id")

	streamsCmd.AddCommand(getRegistryCmd)

	var getRegistryLocal string
	getRegistryCmd.Flags().StringVar(&getRegistryLocal, "local", "false", "local")

	streamsCmd.AddCommand(getTemplateCmd)

	var getTemplateTemplateId string
	getTemplateCmd.Flags().StringVar(&getTemplateTemplateId, "template-id", "", "This is a required parameter. ID of the template")
	getTemplateCmd.MarkFlagRequired("template-id")

	var getTemplateVersion int64
	getTemplateCmd.Flags().Int64Var(&getTemplateVersion, "version", 0, "version of the template")

	streamsCmd.AddCommand(listConnectionsCmd)

	var listConnectionsConnectorId string
	listConnectionsCmd.Flags().StringVar(&listConnectionsConnectorId, "connector-id", "", "")

	var listConnectionsCreateUserId string
	listConnectionsCmd.Flags().StringVar(&listConnectionsCreateUserId, "create-user-id", "", "")

	var listConnectionsFunctionId string
	listConnectionsCmd.Flags().StringVar(&listConnectionsFunctionId, "function-id", "", "")

	var listConnectionsName string
	listConnectionsCmd.Flags().StringVar(&listConnectionsName, "name", "", "")

	var listConnectionsOffset int32
	listConnectionsCmd.Flags().Int32Var(&listConnectionsOffset, "offset", 0, "")

	var listConnectionsPageSize int32
	listConnectionsCmd.Flags().Int32Var(&listConnectionsPageSize, "page-size", 0, "")

	var listConnectionsShowSecretNames string
	listConnectionsCmd.Flags().StringVar(&listConnectionsShowSecretNames, "show-secret-names", "", "")

	var listConnectionsSortDir string
	listConnectionsCmd.Flags().StringVar(&listConnectionsSortDir, "sort-dir", "", "Specify either ascending ('asc') or descending ('desc') sort order for a given field (sortField), which must be set for sortDir to apply. Defaults to 'asc'.")

	var listConnectionsSortField string
	listConnectionsCmd.Flags().StringVar(&listConnectionsSortField, "sort-field", "", "")

	streamsCmd.AddCommand(listConnectorsCmd)

	streamsCmd.AddCommand(listPipelinesCmd)

	var listPipelinesActivated string
	listPipelinesCmd.Flags().StringVar(&listPipelinesActivated, "activated", "false", "activated")

	var listPipelinesCreateUserId string
	listPipelinesCmd.Flags().StringVar(&listPipelinesCreateUserId, "create-user-id", "", "createUserId")

	var listPipelinesIncludeData string
	listPipelinesCmd.Flags().StringVar(&listPipelinesIncludeData, "include-data", "false", "includeData")

	var listPipelinesName string
	listPipelinesCmd.Flags().StringVar(&listPipelinesName, "name", "", "name")

	var listPipelinesOffset int32
	listPipelinesCmd.Flags().Int32Var(&listPipelinesOffset, "offset", 0, "offset")

	var listPipelinesPageSize int32
	listPipelinesCmd.Flags().Int32Var(&listPipelinesPageSize, "page-size", 0, "pageSize")

	var listPipelinesSortDir string
	listPipelinesCmd.Flags().StringVar(&listPipelinesSortDir, "sort-dir", "", "sortDir")

	var listPipelinesSortField string
	listPipelinesCmd.Flags().StringVar(&listPipelinesSortField, "sort-field", "", "sortField")

	streamsCmd.AddCommand(listTemplatesCmd)

	var listTemplatesOffset int32
	listTemplatesCmd.Flags().Int32Var(&listTemplatesOffset, "offset", 0, "offset")

	var listTemplatesPageSize int32
	listTemplatesCmd.Flags().Int32Var(&listTemplatesPageSize, "page-size", 0, "pageSize")

	var listTemplatesSortDir string
	listTemplatesCmd.Flags().StringVar(&listTemplatesSortDir, "sort-dir", "", "sortDir")

	var listTemplatesSortField string
	listTemplatesCmd.Flags().StringVar(&listTemplatesSortField, "sort-field", "", "sortField")

	streamsCmd.AddCommand(mergePipelinesCmd)

	var mergePipelinesInputDatafile string
	mergePipelinesCmd.Flags().StringVar(&mergePipelinesInputDatafile, "input-datafile", "", "The input data file.")

	streamsCmd.AddCommand(putConnectionCmd)

	var putConnectionConnectionId string
	putConnectionCmd.Flags().StringVar(&putConnectionConnectionId, "connection-id", "", "This is a required parameter. ID of the connection")
	putConnectionCmd.MarkFlagRequired("connection-id")

	var putConnectionData string
	putConnectionCmd.Flags().StringVar(&putConnectionData, "data", "", "This is a required parameter. The key-value pairs of configurations for this connection. Connectors may have some configurations that are required, which all connections must provide values for. For configuration values of type BYTES, the provided values must be Base64 encoded.")
	putConnectionCmd.MarkFlagRequired("data")

	var putConnectionDescription string
	putConnectionCmd.Flags().StringVar(&putConnectionDescription, "description", "", "This is a required parameter. The description of the connection.")
	putConnectionCmd.MarkFlagRequired("description")

	var putConnectionName string
	putConnectionCmd.Flags().StringVar(&putConnectionName, "name", "", "This is a required parameter. The name of the connection.")
	putConnectionCmd.MarkFlagRequired("name")

	streamsCmd.AddCommand(putGroupCmd)

	var putGroupGroupId string
	putGroupCmd.Flags().StringVar(&putGroupGroupId, "group-id", "", "This is a required parameter. The group function's ID from the function registry")
	putGroupCmd.MarkFlagRequired("group-id")

	var putGroupInputDatafile string
	putGroupCmd.Flags().StringVar(&putGroupInputDatafile, "input-datafile", "", "The input data file.")

	streamsCmd.AddCommand(putTemplateCmd)

	var putTemplateTemplateId string
	putTemplateCmd.Flags().StringVar(&putTemplateTemplateId, "template-id", "", "This is a required parameter. ID of the template")
	putTemplateCmd.MarkFlagRequired("template-id")

	var putTemplateInputDatafile string
	putTemplateCmd.Flags().StringVar(&putTemplateInputDatafile, "input-datafile", "", "The input data file.")

	streamsCmd.AddCommand(reactivatePipelineCmd)

	var reactivatePipelineId string
	reactivatePipelineCmd.Flags().StringVar(&reactivatePipelineId, "id", "", "This is a required parameter. Pipeline UUID to reactivate")
	reactivatePipelineCmd.MarkFlagRequired("id")

	streamsCmd.AddCommand(startPreviewCmd)

	var startPreviewInputDatafile string
	startPreviewCmd.Flags().StringVar(&startPreviewInputDatafile, "input-datafile", "", "The input data file.")

	var startPreviewRecordsLimit int32
	startPreviewCmd.Flags().Int32Var(&startPreviewRecordsLimit, "records-limit", 0, "The maximum number of events per function. Defaults to 100.")

	var startPreviewRecordsPerPipeline int32
	startPreviewCmd.Flags().Int32Var(&startPreviewRecordsPerPipeline, "records-per-pipeline", 0, "The maximum number of events per pipeline. Defaults to 10000.")

	var startPreviewSessionLifetimeMs int64
	startPreviewCmd.Flags().Int64Var(&startPreviewSessionLifetimeMs, "session-lifetime-ms", 0, "The maximum lifetime of a session, in milliseconds. Defaults to 300,000.")

	var startPreviewStreamingConfigurationId int64
	startPreviewCmd.Flags().Int64Var(&startPreviewStreamingConfigurationId, "streaming-configuration-id", 0, "Deprecated. Must be null if set.")

	var startPreviewUseNewData string
	startPreviewCmd.Flags().StringVar(&startPreviewUseNewData, "use-new-data", "false", "Deprecated. Must be true if set.")

	streamsCmd.AddCommand(stopPreviewCmd)

	var stopPreviewPreviewSessionId int64
	stopPreviewCmd.Flags().Int64Var(&stopPreviewPreviewSessionId, "preview-session-id", 0, "This is a required parameter. ID of the preview session")
	stopPreviewCmd.MarkFlagRequired("preview-session-id")

	streamsCmd.AddCommand(updateConnectionCmd)

	var updateConnectionConnectionId string
	updateConnectionCmd.Flags().StringVar(&updateConnectionConnectionId, "connection-id", "", "This is a required parameter. ID of the connection")
	updateConnectionCmd.MarkFlagRequired("connection-id")

	var updateConnectionData string
	updateConnectionCmd.Flags().StringVar(&updateConnectionData, "data", "", "The key-value pairs of configurations for this connection. Connectors may have some configurations that are required, which all connections must provide values for. For configuration values of type BYTES, the provided values must be Base64 encoded.")

	var updateConnectionDescription string
	updateConnectionCmd.Flags().StringVar(&updateConnectionDescription, "description", "", "The description of the connection.")

	var updateConnectionName string
	updateConnectionCmd.Flags().StringVar(&updateConnectionName, "name", "", "The name of the connection.")

	streamsCmd.AddCommand(updateGroupCmd)

	var updateGroupGroupId string
	updateGroupCmd.Flags().StringVar(&updateGroupGroupId, "group-id", "", "This is a required parameter. The group function's ID from the function registry")
	updateGroupCmd.MarkFlagRequired("group-id")

	var updateGroupInputDatafile string
	updateGroupCmd.Flags().StringVar(&updateGroupInputDatafile, "input-datafile", "", "The input data file.")

	streamsCmd.AddCommand(updatePipelineCmd)

	var updatePipelineId string
	updatePipelineCmd.Flags().StringVar(&updatePipelineId, "id", "", "This is a required parameter. id of the pipeline to update")
	updatePipelineCmd.MarkFlagRequired("id")

	var updatePipelineInputDatafile string
	updatePipelineCmd.Flags().StringVar(&updatePipelineInputDatafile, "input-datafile", "", "The input data file.")

	streamsCmd.AddCommand(updateTemplateCmd)

	var updateTemplateTemplateId string
	updateTemplateCmd.Flags().StringVar(&updateTemplateTemplateId, "template-id", "", "This is a required parameter. ID of the template")
	updateTemplateCmd.MarkFlagRequired("template-id")

	var updateTemplateDescription string
	updateTemplateCmd.Flags().StringVar(&updateTemplateDescription, "description", "", "Template description")

	var updateTemplateInputDatafile string
	updateTemplateCmd.Flags().StringVar(&updateTemplateInputDatafile, "input-datafile", "", "The input data file.")

	var updateTemplateName string
	updateTemplateCmd.Flags().StringVar(&updateTemplateName, "name", "", "Template name")

	streamsCmd.AddCommand(validatePipelineCmd)

	var validatePipelineInputDatafile string
	validatePipelineCmd.Flags().StringVar(&validatePipelineInputDatafile, "input-datafile", "", "The input data file.")

}
