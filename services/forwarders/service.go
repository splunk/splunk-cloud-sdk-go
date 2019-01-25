// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package forwarders

import (
	"io/ioutil"
	"strconv"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

// catalog service url prefix
const servicePrefix = "forwarders"
const serviceVersion = "v2beta1"
const serviceCluster = "api"

// Service talks to the Splunk Cloud catalog service
type Service services.BaseService

// NewService creates a new forwarders service client from the given Config
func NewService(config *services.Config) (*Service, error) {
	baseClient, err := services.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Service{Client: baseClient}, nil
}

// ListCertificates lists all the certificates for the tenant
func (s *Service) ListCertificates() ([]CertificateInfo, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "certificates")
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
	var result []CertificateInfo
	err = util.ParseResponse(&result, response)
	return result, err
}

// DeleteCertificate deletes a certificate on a particular slot on a tenant
func (s *Service) DeleteCertificate(slot int) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "certificates", strconv.Itoa(slot))
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// DeleteCertificates deletes all the certificates on a tenant
func (s *Service) DeleteCertificates() error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "certificates")
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// CreateCertificate creates and adds a certificate to a vacant slot on the tenant.
func (s *Service) CreateCertificate(certificateFileName string) (*CertificateInfo, error) {
	fileBytes, err := ioutil.ReadFile(certificateFileName)
	if err != nil {
		return nil, err
	}
	pemFile := PemFile{Pem: string(fileBytes)}
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "certificates")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: pemFile})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result CertificateInfo
	err = util.ParseResponse(&result, response)
	return &result, err
}
