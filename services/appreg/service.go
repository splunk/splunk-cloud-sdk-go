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
const servicePrefix = "appreg"
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

///*
//CreateApp
//Create a new application.
//* @param tenant The tenant issuing the request.
//* @param createAppRequest Create a new application.
//@return AppResponse
//*/
//func (s *Service) CreateApp(tenant string, createAppRequest CreateAppRequest) (AppResponse, *http.Response, error)
//{
//
//}
///*
//DeleteApp
//Delete an application.
//* @param tenant The tenant issuing the request.
//* @param appName Application name.
//*/
//func (s *Service) DeleteApp(tenant string, appName ResourceName) (*http.Response, error)
///*
//GetApp
//Retrieve the metadata of an application.
//* @param tenant The tenant issuing the request.
//* @param appName Application name.
//@return AppResponse
//*/
//func (s *Service) GetApp(tenant string, appName ResourceName) (AppResponse, *http.Response, error)
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
* @param tenant The tenant issuing the request.
@return []AppResponse
*/
func (s *Service) ListApps() ([]AppResponse, error){
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "appreg")
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
