# Splunk Cloud SDK for Go Changelog

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