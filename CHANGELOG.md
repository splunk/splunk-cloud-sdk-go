# Splunk Cloud SDK for Go Changelog

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