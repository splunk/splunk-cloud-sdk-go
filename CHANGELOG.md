# Splunk Cloud SDK for Go Changelog

## Version 1.0.0-beta.0
### Breaking Changes
* Update all services with latest spec changes
* refactor the services apis and parameter setter functions

## Version 0.11.0
### Non-breaking Changes
* Update the default URL from splunkbeta.com to scp.splunk.com 
### Bug Fixes
* Return detailed TOS error

## Version 0.10.0
### Breaking Changes
* Update App Registry service api and models to latest version v1beta2 that match the spec files using auto generated code
* Update Ingest service api and models  to latest version v1beta2 that match the spec files using auto generated code
* Update KVStore service api and models  to latest version v1beta1 that match the spec files using auto generated cod
* Update ML service api and models  to latest version v2beta1 that match the spec files using auto generated code
* Update Identity service api and models to latest version v2beta1 that match the spec files using auto generated code
* Update Search service apis and models to latest version v2beta1 
* Update Catalog service api and models to latest version v2beta1 
* Update auth flow 
* Update Stream service api and models to version v2beta1 that match the spec files using auto generated code

### Bug Fixes
* Fix batch event sender code

## Version 0.9.1
### Breaking Changes
* Update Forwarder service apis and models to latest version v2beta1 that match the spec files using auto generated code
* Update Action service api and models to latest version v1beta2 that match the spec files using auto generated code
* Update Catalog service api and models  to latest version v2alpha2 that match the spec files using auto generated code
* Update Search service api and models to latest version v2alpha2 that match the spec files using auto generated code

### Non-breaking Changes
* Update Provisioner service apis and models using auto generated code

## Version 0.9.0
We now have auto generate generated models from spec files for all services and auto generated API signatures from spec files for Search (v2alpha2) and Identity (v2beta1) services. If you are using old SDK, you may need to update your client code as we have many breaking changes.

### Breaking Changes
* For all services the models are changed to match in the spec files
* For Search and Identity services the API signatures are changed to match in the spec files
* Remove the deprecated service package, Client is not supported any more and should use sdk.Client instead
* Remove the deprecated model package
* Update Auth flow using IAC endpoints

## Version 0.8.3
### Non-Breaking Changes
* Remove 401 AuthnHandler from SDK client that automatically gets a new token from idp. Instead, a StartTime field is added to track the token expiry time. Every request creation is preceded by checking the validity of the token, if its set to expire, retrieve a new token and update the context, all other concurrent requests will wait on the lock set during token renewal and further on use the new token.

### Bug Fixes
* Allow http request body to be nil

<!-- NOTE: DO NOT CHANGE THIS FILE MANUALLY (except to make minor corrections). The first line and the format between versions (## Version x.y.z) are used in preparing/publishing our releases, this file should automatically be populated during the prerelease stage using the prerelease tag message -->