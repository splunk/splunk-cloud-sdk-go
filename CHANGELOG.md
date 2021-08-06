# Splunk Cloud SDK for Go Changelog

## Version 1.12.0-beta.5


#### Non-Breaking Changes 

* KVstore service :  `DeleteRecordsQueryParams` struct has added a new  parameter of `EnableMvl`;   `QueryRecordsQueryParams` struc has  added  new optional input parameters of `EnableMvl` and `Shared`


* Streams service: new api of `DeleteSource()` is added

#### Features
* Add connection timeouts to default retry functionality in GO sdk client

## SCloud v8.0.0-beta.5

#### Non-Breaking Changes
* KVstore command: 
    - New flag `enable-mvl` is added to `delete-records` command;   New flags `enable-mvl` and `shared` is added to `query-records` command

* Streams command: New command of `delete-source` is added


## Version 1.12.0-beta.4
#### Breaking Changes
* Ingest Service (v1beta2):
    - Model `HECResponse` and `InlineObject` have been removed

* KVStore Service (v1beta1):
    - Model `Key` renamed to `Record` and has an additional required field `User`
    - Endpoints `InsertRecord` and `PutRecord` have returnType Record (earlier it was `Key`)

#### Non-Breaking Changes
* Identity Service (v3):
    - Models `ResetPasswordBody`, `UpdateGroupBody`, `UpdatePasswordBody` and `UpdateRoleBody` have been added
    - Endpoints `ResetPassword`, `UpdatePassword`, `UpdateGroup` and `UpdateRole` have been added
    - Parameters `Description` and `DisplayName` have been added to models `CreateGroupBody`, `CreateRoleBody`, `Role`, and `Group`.

* Streams Service (v3beta1):
    - Models `PipelineReactivateResponseAsync`, `PipelineReactivationStatus`, `UpgradePipelineRequest` and `ValidateConnectionRequest` have been added. 
    - Parameter `SkipValidation` has been added to `CreateConnection` endpoint
    - Parameter `CreateUserId` has been added to `ListTemplates` endpoint
    - Endpoints `ReactivationStatus`, `ValidateConnection`, and `UpgradePipeline` have been added
    - Parameter `Metadata`  has been added to `ConnectorResponse`
    - Parameter `Labels` has been added to `PipelinePatchRequest`, `PipelineRequest` and `PipelineResponse` models
    - Parameter `UberJarSha256` has been added to `PipelineResponse` model
    - `ACTIVATING` and `DEACTIVATING` added to model `PipelineResponse` status enum
    - Model `Source` has changed with properties `Node` and `PipelineVersion` deleted and 18 new properties added
    - Parameter `LearnMoreLocation` has been added to model `TemplateResponse` 

## SCloud v8.0.0-beta.4

#### Non-Breaking Changes
* Identity command: 
    - New commands `reset-password`, `update-group`, `update-password` and `update-role` were added 
    - New flags `description` and `display-name` added for `create-group`, `create-role`, `update-group` and `update-role` commands
* Streams command: 
    - New commands `reactivation-status`, `upgrade-pipeline` and `validate-connection` were added
    - New flag `skip-validation` added for `create-connection` command
    - New flag `labels` added for `create-pipeline`, `patch-pipeline` and `update-pipeline` commands
    - New flag `create-user-id` added for `list-templates` command

## Version 1.12.0-beta.3

#### Non-Breaking Changes
* Modified ingestSearch example to not fail on receiving 429 or 500 http error

## SCloud v8.0.0-beta.3

#### Non-Breaking Changes
* Support SCS environments gstage and prod:
   New environment `playground-scs` can be specified to target `api.playground.scs.splunk.com` environment out-of-the-box


## Version 1.12.0-beta.2

#### Breaking Changes
* Identity service: `UseDefaultIdp` optional property was removed from `Tenant` in model
* Ingest service: `PostCollectorRaw` and `PostCollectorRawV1` APIs were removed
* Streams service: `CreateDataStream`, `DeleteDataStream`, `DescribeDataStream`, `ListDataStreams`, and `UpdateDataStream` APIs were removed

#### Non-Breaking Changes
* Search service: updated to target v2 endpoints (previously v2beta1) including new endpoint: `ExportResults`
* Identity service: added new APIs of `CreateIdentityProvider`, `DeleteIdentityProvider`, `GetIdentityProvider`, `ListIdentityProvider`, and `UpdateIdentityProvider`; 
New optional property `AcceptTos` for `CreatePrincipalBody` was added in model
  
* Streams service: Added a new API of `UploadLookupFile` 

## SCloud v8.0.0-beta.2

#### Breaking Changes
* Ingest command: removed `post-collector-raw` and `post-collector-raw-v-1` commands
* Streams command: removed commands of `create-data-stream`, `delete-data-stream`, `describe-data-stream`, `list-data-streams`, `update-data-stream`

#### Non-Breaking Changes
* Search command: updated to v2 endpoints (previously v2beta1) including new commmand: `export-results`
* Identity commmand: new commands of `create-identity-provider`, `delete-identity-provider`, `get-identity-provider`, `update-identity-provider` were added; new `--accept-tos` flag for identity `create-principal` commmand
* Streams command: new command of  `upload-lookup-file` was added


## Version 1.12.0-beta.1
* Added support for multipart/form data endpoints
* Multi-cell support:
	- Added support for tenant/region scoped hostnames to invoke tenant and system based api, auth domain endpoints 
 	- New config settings to create SDK/Auth client -> tenantScoped bool, region, tenant
	- BREAKING change: IDP TokenRetrievers such as NewClientCredentialsRetriever now accept hostURL config which consists of tenantScoped, region, tenant parameters
	- To enable tenantScoped(multi-cell), set tenantScoped to True
  	- Enabling tenantScoped setting in the auth client will only generate tenant scoped tokens

## SCloud v8.0.0-beta.1
* Multi-cell support:
    - New flag --region to specify region for system scoped api calls
    - New flag --tenant-scoped to enable tenant/region scoping of the hostnames to support multi-celll (by default tenantScoped is set to False currently in scloud)
* Support SCS environments gstage and prod:
	New environments staging-scs and prod-scs can be specified to target gstage and gprod environments out-of-the-box
* Support Device flow authentication:
	- scloud login by default points to device-flow for the staging-scs and prod-scs environments
	- To login using device-flow 
		scloud config set -—key username -—value <user>
		scloud config set -—key tenant -—value <tenantName>
		scloud config set -—key env -—value staging-scs
		scloud login -—use-device or scloud login
	- To login using pkce-flow
		scloud login -—use-pkce
	- To login using refresh-flow (once logged in via one of the above flows)
		scloud login -—use-refresh-token
* BREAKING change: Added required flag `tenant` to the context set command
    - To set a context
        scloud context set --key access_token --value <token> --tenant <tenantName>
* Added support to list contexts
    - To list all the contexts
        scloud context list
    - To list context specific to a given tenant
        scloud context list --tenant <tenantName>

## Version 1.11.1

### Bugfix
 * Bugfix in scloud

## SCloud v7.1.0

### Bugfix
 * Fixed scloud panic when ~/.scloud_context file is missing

## Version 1.11.0
### Services

#### Breaking Changes

##### Features

* Identity v2beta1:
    * Version `v2beta1` replaced by new version `v3`

* Provisioner v1beta1
    * Models `CreateProvisionJobBody`, `ProvisionJobInfo`, `ProvisionJobInfoErrors`, `ProvisionJobInfoErrors` and 'ProvisionJobs' removed
    * Endpoints `CreateProvisionJob`, `GetProvisionJob` and `ListProvisionJobs` removed

* Streams v3beta1:
    * Models `CollectJobPatchRequest`, `CollectJobRequest`, `CollectJobResponse`, `CollectJobStartStopResponse`, `EntitlementRequest`, `EntitlementResponse`, `PaginatedResponseOfCollectJobResponse`, `PaginatedResponseOfPlugin`, `PaginatedResponseOfRulesResponse`,  `PaginatedResponseOfRuleKind`, `Plugin`, `PluginPatchRequest`, `PluginRequest`, `PluginResponse`, `RulesRequest` and `RulesResponse` removed

    * Model `UploadFile` renamed to `UploadFileResponse`

    * Endpoints `CreateCollectJob`, `CreateRulesPackage`, `DeleteCollectJobs`, `DeleteCollectJob`, `DeleteEntitlements`, `DeletePlugin`, `DeleteRulesPackage`, `GetCollectJob`, `GetEntitlements`, `GetPlugins`, `GetRulesPackageById`, 
    `ListCollectJobs`, `ListRulesKinds`, `ListRulesPackages`, `PatchPlugin`, `RegisterPlugin`, `ReleaseInfo`, `SetEntitlements`, `StartCollectJob`, `StopCollectJob`, `UpdateCollectJob`, `UpdatePlugin` and `UpdateRulesPackageById`   

#### Non-Breaking Changes

##### Features

* Added Device flow token retriever support

* Ingest v1beta2: 
    * New endpoints `PostCollectorRaw` and `PostCollectorRawV1` added

* Streams v2beta1:
    * New property `Attributes` added to `ConnectorResponse`
    * New property `StatusDescription` added to `PipelineReactivateResponse`
    * New parameter `functionOp` added to `listConnections`

* Streams v3beta1
    * New model `UploadFileResponse` added
    * New endpoint `DeleteLookupFile`, `GetLookupFileMetadata` and `GetLookupFilesMetadata` added

## SCloud v7.0.0
### Breaking Changes
* Provisioner Command
    * Removed :
        * create-provision-job
        * get-provision-job
        * list-provision-job
* Streams Command
    * Removed :
        * create-collect-job
        * create-rules-package
        * delete-collect-job
        * delete-entitlements
        * delete-plugin
        * delete-rules-package
        * get-collect-job
        * get-entitlements
        * get-plugins
        * get-rules-package-by-id
        * list-collect-job
        * list-rules-kinds
        * list-rules-packages
        * patch-plugin
        * register-plugin
        * release-info
        * set-entitlements
        * start-collect-job
        * stop-collect-job
        * update-collect-job
        * update-plugin
        * update-rules-package-by-id
        * upload-plugin

### Non-Breaking Changes
* Added support for Keycloak authorization flow to enable scloud to get and set session token from the cookies

* Identity Command
    * Added :
        * create-principal
    * New parameters orderby, page-size and page-token added :
        * list-group-members 
        * list-group-roles
        * list-groups
        * list-member-groups
        * list-member-permissions
        * list-member-roles
        * list-principals
        * list-role-groups
        * list-role-permissions
        * list-roles
    * New parameters kind, orderby, page-size and page-token added :
        * list-members

* Ingest Command
    * Added :
        * post-collector-raw
        * post-collector-raw-v-1

* Streams Command
    * Added :
        * delete-lookup-file
        * get-lookup-file-metadata
        * get-lookup-files-metadata

## Version 1.10.0
#### Breaking Changes
##### Features
* Catalog v2beta1:
    * `CreateDatasetImport` returns datatype of `Dataset` (replaced 'ImportDataset')
    * `CreateDatasetImportById` returns `DatasetImportedby` (replaced 'ImportDataset')
    * `DatasetImportedBy` has a new property `Owner` and property `Name` is now optional

* Provisioner v1beta1:
    * Model `ECStackName` renamed to `EcStackName`

* Search v2beta1:
    * Model `ListSearchResultsResponseFields` renamed to `ListPreiviewResultsResponseFields`

* Search v3alpha1:
    * Model `ListSearchResultsResponseFields` renamed to `ListPreiviewResultsResponseFields`

* Stream v3beta1: 
    * Model `RulesSourcetypesResponse` renamed to `RulesPackageSourcetypes`
    * Model `RulesActionsResponse` renamed to `RulesPackageActions`
    
#### Non-Breaking Changes
##### Features

* Auth 
    * `ServicePrincipalAuthManager` added to Auth service

* Identity v2beta1:
    * New model `AddInvisibleMemberBody` added
    * New endpoints `AddInvisibleMember`, `GetMemberAdmin` and `RemoveMemberAdmin` added
    * New properties `ExpiresAt` and `Visible` added to `Member` model
    * New models `DeviceAuthInfo` and `UpdateRoleBody` added

* Identity v3alpha1: 
    * New version introduced

* Ingest v1beta2: 
    * New models `UploadSuccessResponse` and `FileUploadDetails`  added

* KVStore v1beta1:
    * New endpoint `TruncateRecords` added

* Search v3alpha1:
    * New models `SearchModule`, `StatementDispatchStatus`, and `SingleSatatementQueryParamters` added
    * New endpoints `createMultiSearchMethod` and `createSearchStatements` added

* Streams v2beta1:
    * New property `Messages` added to model `ConnectionSaveResponse`
    * New property `Complexity` added to model `PipelineResponse`
    * New property `ActivateLatestVersion` added to model `ReactivatePipelineRequest`

* Streams v3beta1:
    * New models `CollectJobPatchRequest`, `DataStream`, `DataStreamRequest`, `DataStreamResponse`, `EntitlementRequest`,
     `EntitlementResponse`, `PaginatedResponseOfRuleKind`, `RulesKind` and `PluginResponse`
    * New endpoints `CreateDataStream`, `DeleteCollectJob`, `Deletedatastream`, `DeleteEntitlements`, `DeleteRulesPackage`, 
     `DescribeDataStream`, `GetEntitlements`, `GetRulesPackageById`, `ListDataStreams`, `ListRuleKinds`, `ReleaseInfo`, 
     `SetEntitlements`, `UpdateCollectJob`, `UpdateDataStream` and `UpdateRulesPackageById` added
 
## SCloud v6.0.0
### Breaking Changes
* Action Command
    * create-action-email-action replaced by  create-action-email
    * create-action-webhook-action replaced by  create-action-webhook
    * update-action-email-action-mutable replaced by update-action-email-mutable
    * update-action-webhook-action-mutable replaced by update-action-webhook-mutable
* AppRegistry Command
    * create-app-native-app replaced by create-app-native
    * create-app-service-app replaced by  create-app-service
* Catalog Command
    * create-action-for-rule-alias-action replaced by create-action-for-rule-alias
    * create-action-for-rule-auto-kv-action  replaced by create-action-for-rule-auto-kv
    * create-action-for-rule-by-id-alias-action replaced by  create-action-for-rule-by-id-alias
    * create-action-for-rule-by-id-auto-kv-action  replaced by create-action-for-rule-by-id-auto-kv
    * create-action-for-rule-by-id-eval-action  replaced by create-action-for-rule-by-id-eval
    * create-action-for-rule-by-id-lookup-action  replaced by create-action-for-rule-by-id-lookup
    * create-action-for-rule-by-id-regex-action  replaced by create-action-for-rule-by-id-regex
    * create-action-for-rule-eval-action  replaced by create-action-for-rule-eval
    * create-action-for-rule-lookup-action  replaced by create-action-for-rule-lookup
    * create-action-for-rule-regex-action  replaced by create-action-for-rule-regex
    * create-dataset-index-dataset replaced by create-dataset-index
    * create-dataset-kv-collection-dataset replaced by create-dataset-kv-collection
    * create-dataset-lookup-dataset replaced by create-dataset-lookup
    * create-dataset-metric-dataset  replaced by create-dataset-metric
    * create-dataset-view-dataset replaced by create-dataset-view
    * update-action-by-id-for-rule-alias-action replaced by update-action-by-id-for-rule-alias
    * update-action-by-id-for-rule-auto-kv-action replaced by update-action-by-id-for-rule-auto-kv
    * update-action-by-id-for-rule-by-id-alias-action replaced by update-action-by-id-for-rule-by-id-alias
    * update-action-by-id-for-rule-by-id-auto-kv-action replaced by update-action-by-id-for-rule-by-id-auto-kv
    * update-action-by-id-for-rule-by-id-eval-action replaced by update-action-by-id-for-rule-by-id-eval
    * Update-action-by-id-for-rule-by-id-lookup-action replaced by update-action-by-id-for-rule-by-id-lookup
    * update-action-by-id-for-rule-by-id-regex-action replaced by update-action-by-id-for-rule-by-id-regex
    * update-action-by-id-for-rule-eval-action replaced by update-action-by-id-for-rule-eval
    * update-action-by-id-for-rule-lookup-action replaced by update-action-by-id-for-rule-lookup
    * update-action-by-id-for-rule-regex-action  replaced by update-action-by-id-for-rule-regex
    * update-dataset-by-id-import-dataset  replaced by update-dataset-by-id-import
    * update-dataset-by-id-index-dataset replaced by update-dataset-by-id-index
    * update-dataset-by-id-kv-collection-dataset replaced by update-dataset-by-id-kv-collection
    * update-dataset-by-id-lookup-dataset replaced by update-dataset-by-id-lookup
    * update-dataset-by-id-metric-dataset replaced by update-dataset-by-id-metric
    * update-dataset-by-id-view-dataset replaced by update-dataset-by-id-view
    * update-dataset-import-dataset replaced by update-dataset-import
    * update-dataset-index-dataset replaced by update-dataset-index
    * update-dataset-kv-collection-dataset replaced by update-dataset-kv-collection
    * update-dataset-lookup-dataset replaced by update-dataset-lookup
    * update-dataset-metric-dataset replaced by update-dataset-metric
    * update-dataset-view-dataset replaced by update-dataset-view
    * create-dataset-import 
        * requires "owner"
        * name is NOT required
    * create-dataset-import-by-id 
        * requires "owner" 
        * name is NOT required
    * create-dataset-import-by-idv-1 
        * requires "owner" 
        * name NOT required
    * create-dataset-importv-1
        * requires "owner" 
        * name NOT required
* Streams
    * get-rules-package is replaced by get-rules-package-by-id
* create-collect-job
    *  requires new parameters: cron & workers
    
### Non-Breaking Changes
* Identity Command
    * Added :
        * add-invisible-member 
        * get-member-admin
        * remove-member-admin
* Ingest Command
    * post-collector-tokens
        * New parameters
            * ack-enabled 
            * allow-query-string-auth
            * disabled 
            * index 
    * put-collector-token
        * New parameters
            * ark-enabled
            * allow-query-string-auth
            * description
            * disabled
            * indexes
* KVStore
    * Added
        * truncate-records
* Streams
    * Added
        * create-data-stream
        * delete-collect-job (all jobs)
        * delete-data-stream
        * delete-entitlements
        * delete-rules-packages
        * describe-data-stream
        * get-entitlements
        * list-data-streams
        * list-rules-kinds
        * release-info
        * set-entitlements
        * update-collect-job
        * update-data-stream
        * update-rules-package-by-id
    * Modified
        * create-rules-package - `arguments` and `kind` are no longer used

## Version 1.9.0
## Go SDK v1.9.0

### Breaking Changes
*  AppRegistry 
   *  Removed support for creating app of type `WebApp`, all models and properties pertaining to WebApp have been removed.
*  Catalog
   *  Model `JobDataset` updated to `JobDatasetGet` which includes new field `ExtractFields`
*  Service client creation
   * Service client creation now accepts a Client object instead of a SDK specific config object
### Features
*  Search
   * Model `SearchJob` updated with new field `ExtractFields`
*  Streams
   * New endpoints and models added for `CreateRulesPackage`, `GetRulesPackage`, `ListRulesPackage`, `ListCollectJobs`

## SCloud v5.0.0
### Breaking Changes
*  Remove AppRegistry command `create-app-web-app` to create `WebApp `
### Features
*  New field `extract-fields` added to Search command `create-job`
*  New commands for `CreateRulesPackage`,` GetRulesPackage`, `ListRulesPackage`, `ListCollectJobs`
*  Help update - Usage information will not be returned for service specific errors. Short usage information is returned for incorrect command / arguments that are CLI specific, long usage will be reserved for --help

## Version 1.8.0
**Breaking Changes**

`Catalog service`: GetDataset,  GetDatasetById and ListDatasets endpoints now return data type of DatasetGet instead of Dataset

`Identify service`: remove endpoint SetPrincipalPublicKeys

`Kvstore service`:  InsertRecords endpoint has a new parameter of InsertRecordsQueryParams

`Stream service`: ReactivatePipeline has a new parameter of ReactivatePipelineRequest

**Non-Breaking Changes**

`Identify service`: 
- new endpoints are added: AddPrincipalPublicKey, GetPrincipalPublicKey, GetPrincipalPublicKeys, DeletePrincipalPublicKey, and UpdatePrincipalPublicKey 

`Stream service`: 
- new endpoints are added: CreateCollectJob, DeleteCollectJob, DeletePlugin, GetCollectJob, GetPlugins, PatchPlugin, RegisterPlugin,StartCollectJob,StopCollectJob,UpdatePlugin


## Version 1.7.0
## Go SDK v1.7.0

### Breaking Changes
*  Identity
   *  Added  `query` parameter with type `ListMemberPermissionsQueryParams` to  `ListMemberPermissions` endpoint 
   *  New `ListMemberPermissionsQueryParams` model with `ScopeFilter` property
*  Search
   *  Modified `MaxTime` property to be typed int32  in `SearchJob`
   *  Modified `Duration` property to be typed `float64` in `FieldSummary`
   *  Modified `RequiredFreshness` property to be typed `int32` in `FieldSummary` 
   *  Modified `Mean` and `Stddev` properties to typed `float64` in `SingleFieldSummary`
   *  Modified `Duration` and `EarliestTime` properties to typed `float64` in `SingleTimeBucket`
   *  Modified `Count` and `Offset` properties to be typed `*int32` in `ListEventsSummaryQueryParams` model
   *  Modified `Count` property to be typed `*int32` in `ListJobsQueryParams` model
   *  Modified `Count` and `Offset` properties to be typed `*int32` in `ListPreviewResultsQueryParams` model
   *  Modified `Count` and `Offset` properties to be typed `*int32` in `ListResultsQueryParams` model
*  Streams
   * Modified `ConnectorId` property to `[]string` in `ListConnectionsQueryParams` model
### Non-Breaking Changes
*  Catalog
   * `AppClientIdProperties` model has been added
* Identity
   * New `SetPrincipalPublicKeys` endpoint added
*  Streams
   * Added `NodeId` property to ValidateResponse Model
   * Removed `JsonNode` Model
   * Added `GetFileMetadata` endpoint

## SCloud v4.0.0
### Breaking Changes
*  SCloud forwarders AddCertificate endpoint accepts inputFile
### Non-Breaking Changes
*  New `Context list` cmd added
*  New `Context set --key <key> --value <value>` cmd added
*  New `Identity set-principal-public-keys --principal <principal name>` cmd added
*  New `--use-refresh-token` flag added to trigger refresh authentication flow
*  Increased client timeout to 60 seconds
### BugFix
* Fixed system level crash when `--env` flag is used

## Version 1.6.0
### BREAKING CHANGES

* Streams - Includes support for new Streams version v3beta1 and support for v2Beta1 has been removed. As a result, here are the endpoints that are either removed or updated
	* `CompileDSL` is not longer supported, substituted by Compile which leverages SPL instead of DSL to produce streams JSON object
	*  CRUD on `Group`  endpoints have been removed and all models corresponding to Groups have been removed.
	* `ExpandGroup` which creates and returns the expanded version of a group has been removed.
	* `UpdatePipeline` endpoint accepts a request body of type PipelineRequest instead of PipelinePatchRequest.
	* `UplPipeline` model replaced by `Pipeline` model
	* `UplRegistry` model replaced by `RegistryModel` model
	* `MergePipelines` support has been removed


### FEATURES

* Search
	* New `DeleteSearchJob` creates a search job that deletes events from an index.
* Streams
	* New `UploadFiles` endpoint uploads files to streams
	* New `GetLookupTable` endpoint returns lookup table results
	* New `Decompile` endpoint decompiles UPL and returns SPL 
	* New `DeleteFile` endpoint deletes a file give a file-id
	* New `GetFilesMetaData` endpoint returns files metadata

## Version 1.5.0
Release v1.5.0

## Version 1.5.0

### Services

#### Breaking Changes

- Provisioner
	- Removed endpoints: `CreateEntitlementsJob` and `GetEntitlementsJob`

#### Features

- Ingest
  - Support for new operations: `deleteAllCollectorTokens`, `listCollectorTokens`, `postCollectorTokens`, `deleteCollectorToken`, `getCollectorToken`, `putCollectorToken`
- Search
  - Support for new operation: `deleteJob`

## Version 1.4.0
### Breaking Changes
* Scloud v2.0 is released: all new commands and formats. Check https://dev.splunk.com/scs/docs/overview/tools/tools_scloud/ to learn more about Scloud v2.0

* The type of model search.searchjob.AllowSideEffects was changed from interface{} to *bool

### Deprecated code
* Scloud v1.x.x will not be supported any more. Older version of scloud binaries can still be found at previous releases

## Version 1.3.0
### BREAKING CHANGES

* AppRegistry
	* Model `UpdateAppRequest` has been refactored from a discriminator-based app-kind-specific model to a single model.
* Catalog
	* `JobDatasetPATCH` and `JobDatasetPOST` have been removed.
* Collect
	* `executionPatch` model now requires `status` property.
* Forwarders
	* `Certificates` model now requires `pem` property.

### FEATURES

* Identity
	* `ListGroups` now allows passing a query  to filter by access permission
	* `ListMemberPermissions` returns new `max-age` header describing how long member permission can be cached
	* New `RevokePrincipalAuthTokens` revokes all tokens for a principal
* Provisioner
	* Support for new endpoints: `CreateEntitlementsJob` and `GetEntitlementsJob`

## Version 1.2.1
### Breaking Changes
*Appregistry models UpdateAppRequest, CreateAppRequest, AppResponseCreateUpdate,AppResponseGetList have been refactored from single model encompassing app related properties to discriminator based using app kind specific models - NativeApp, Webapp, ServiceApp models
### Non-breaking Changes
*Collect service has added support for new endpoints - CreateExecution, GetExecution, PatchExecution for scheduled jobs
*Identity service has new Enum value for TenantStatus - tombstones
*Search service has additional filter parameter in ListJobsQueryParams

## Version 1.1.1
### Non-breaking Changes
* Update PKCE auth flow to read the CSRF token from the response cookie returned from the /csrfToken endpoint to mitigate security bug SCP-16944

## Version 1.1.0
### Breaking Changes
* The type for model of ValidateTokeninclude in Identity service is changed from []string to []ValidateTokenincludeEnum 
* model renamed
### Bug Fixes
* Parse array types for map values in url param correctly

## Version 1.0.1
### Breaking Changes
* `CreateUserId` in PipelineRequest model in Stream service has been removed

### Non-breaking Changes
* SearchJob model in Search service has a new field `RequiredFreshness`

## Version 1.0.0
### Breaking Changes
* The `TriggerEvent model` in action service is changed from required to optional field
* The `catalog.GetDataset()` and `catalog.GetDatasetById()` are changed to require a new query parameter arguments 
* The `ScalePolicy` model in collect service has changed its type from `map[string]interface{}` to `*ScalePolicy`
* The property `include` in `ValidateTokenQueryParams` model in identity service has changed type from []string to new type of `ValidateTokeninclude`

### Non-breaking Changes
* Add `delete-certificate` cmd in Scloud 
* Add `ListAnnotations()` api in catalog service
* Add `CreateWorkflowStreamDeployment()`, `DeleteWorkflowStreamDeployment()`, `GetWorkflowStreamDeployment()`  in ml service
* Add `ListPreviewResults()` in search service

### Bug Fixes
* Fix event size calculation in batch_event_sender

## Version 1.0.0-beta.4
### Breaking Changes
* Update to KVStore PutRecord API call to remove Record version identifier
* Update to Catalog Annotations CRUD API calls to use a generic requestBody map[string]interface  in lieu of AnnotationPost for specifying AnnotationProperties
* Update Login flow to use /csrfToken endpoint - IDP client now accepts a csrfTokenPath as well.

### Non-breaking Changes
* Updates corresponding to latest service versions

### Bug Fixes
* Response body not closed in the api calls to service

## Version 1.0.0-beta.3
### Breaking Changes
* Removed -host, -port, -scheme flags in favor of -host-url in scloud
### Non-breaking Changes
* Incorporated updates to services for latest versions
* GoSDK/Scloud support for ingest via file endpoint
* GoSDK support for collect service
* Allow idp client in gosdk to disable TLS cert validation with -insecure
* scloud appreg list-subscriptions command does not require kind as a required param
* Update scloud streams list templates command to include sort-dir and sort- field
* Update scloud provisioner create-provision-job command, apps flag to be app instead and still be multivalued
* Allow auth-urls to be specified in scloud
### Bug fix
* Fix issue with auth error handling which prevented errors from IACs token endpoint from being surfaced via sdk

## Version 1.0.0-beta.2
### Non-breaking Changes
* Updated Provisioner service with spec changes for invites
* Updated KVStore service with spec changes
* Updated catalog service with spec changes
* Updated Scloud to include commands for provisioner invites

## Version 1.0.0-beta.1
### Breaking Changes
* Updated Catalog service and models with spec changes
### Non-breaking Changes
* Updated Identity service with spec changes
* Updated Provisioner service with spec changes
* Moved Scloud source code to this repository

## Version 1.0.0-beta.0
* Splunk Cloud SDK For GO v1.0.0-beta.0 release
