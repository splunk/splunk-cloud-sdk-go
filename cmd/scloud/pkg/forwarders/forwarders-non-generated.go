// Package forwarders -- manual implementation
package forwarders

import (
	"io/ioutil"

	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/auth"
	model "github.com/splunk/splunk-cloud-sdk-go/services/forwarders"
)

// AddCertificate
func AddCertificateOverride(filename string) (*model.CertificateInfo, error) {
	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data := string(bytes)

	resp, err := client.ForwardersService.AddCertificate(model.Certificate{Pem: data})
	if err != nil {
		return nil, err
	}
	return resp, nil

}
