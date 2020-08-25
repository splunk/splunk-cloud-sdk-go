# Splunk Cloud SDK for Go Changelog

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