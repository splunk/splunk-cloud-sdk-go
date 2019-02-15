// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package appreg

import (
	"github.com/splunk/splunk-cloud-sdk-go/services"

	"github.com/splunk/splunk-cloud-sdk-go/util"
)

// action service url prefix
const servicePrefix = "app-registry"
const serviceVersion = "v1beta2"
const serviceCluster = "api"

// Service - A service the receives incoming notifications and uses
// pre-defined templates to turn those notifications into meaningful actions
type Service services.BaseService

// NewService creates a new action service client from the given Config
func NewService(config *services.Config) (*Service, error) {
	baseClient, err := services.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Service{Client: baseClient}, nil
}

/*
CreateApp
Create a new application.
* @param createAppRequest Create a new application.
@return AppResponse
*/
func (s *Service) CreateApp(createAppRequest *CreateAppRequest) (*AppResponse, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "apps")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: createAppRequest})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result AppResponse
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}


/*
DeleteApp
Delete an application.
* @param appName Application name.
*/
func (s *Service) DeleteApp(appName string) (error){
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "apps", appName)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}
	return nil
}

/*
GetApp
Retrieve the metadata of an application.
* @param appName Application name.
@return AppResponse
*/
func (s *Service) GetApp(appName string) (*AppResponse,  error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "apps", appName)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result AppResponse
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}



///*
//GetAppSubscriptions
//Retrieve the collection of subscriptions to an app.
//* @param tenant The tenant issuing the request.
//* @param appName Application name.
//@return []Subscription
//*/
//func (s *Service) GetAppSubscriptions(tenant string, appName ResourceName) ([]Subscription, *http.Response, error)

/*
ListApps
List applications.
@return []AppResponse
*/
func (s *Service) ListApps() ([]AppResponse, error){
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "apps")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var results []AppResponse
	err = util.ParseResponse(&results, response)
	if err != nil {
		return nil, err
	}
	return results, nil
}


///*
//RotateSecret
//Rotate the client secret for the application.
//* @param tenant The tenant issuing the request.
//* @param appName Application name.
//@return AppResponse
//*/
//func (s *Service) RotateSecret(tenant string, appName ResourceName) (AppResponse, *http.Response, error)
///*
//UpdateApp
//Update an application.
//* @param tenant The tenant issuing the request.
//* @param appName Application name.
//* @param updateAppRequest Updated app contents.
//@return AppResponse
//*/
//func (s *Service) UpdateApp(tenant string, appName ResourceName, updateAppRequest UpdateAppRequest) (AppResponse, *http.Response, error)
