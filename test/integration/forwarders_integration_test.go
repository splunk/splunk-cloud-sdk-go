/*
 * Copyright 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package integration

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/services/forwarders"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test ListCertificates
func TestListCertificates(t *testing.T) {
	err := getSdkClient(t).ForwardersService.DeleteCertificates()
	require.NoError(t, err)

	// Certificate 1
	// Choose a name to identify the forwarder, create and generate a certificate
	forwarder := "forwarder_01"
	err, pemFile1, certificateInfo1 := createAndUploadCertificate(t, forwarder)
	require.NoError(t, err)
	require.NotNil(t, pemFile1)
	require.NotNil(t, certificateInfo1)

	defer deletePemFile(t, pemFile1)
	defer getSdkClient(t).ForwardersService.DeleteCertificate(strconv.FormatInt(*certificateInfo1.Slot, 10))

	// Certificate 2
	// Choose a name to identify the forwarder, create and generate a certificate
	forwarder = "forwarder_02"
	err, pemFile2, certificateInfo2 := createAndUploadCertificate(t, forwarder)
	require.NoError(t, err)
	require.NotNil(t, pemFile2)
	require.NotNil(t, certificateInfo2)

	defer deletePemFile(t, pemFile2)

	// Check if the test certificates were created as expected
	var resp http.Response
	certificates, err := getSdkClient(t).ForwardersService.ListCertificates(&resp)
	require.NoError(t, err)
	assert.NotZero(t, len(certificates))
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

// Test DeleteCertificates
func TestDeleteCertificates(t *testing.T) {
	// Certificate 1
	// Choose a name to identify the forwarder, create and generate a certificate
	forwarder := "forwarder_01"
	err, pemFile1, certificateInfo1 := createAndUploadCertificate(t, forwarder)
	require.NoError(t, err)
	require.NotNil(t, pemFile1)
	require.NotNil(t, certificateInfo1)

	defer deletePemFile(t, pemFile1)
	defer getSdkClient(t).ForwardersService.DeleteCertificates()

	// Certificate 2
	// Choose a name to identify the forwarder, create and generate a certificate
	forwarder = "forwarder_02"
	err, pemFile2, certificateInfo2 := createAndUploadCertificate(t, forwarder)
	require.NoError(t, err)
	require.NotNil(t, pemFile2)
	require.NotNil(t, certificateInfo2)

	defer deletePemFile(t, pemFile2)

	// Delete all the certificates
	err = getSdkClient(t).ForwardersService.DeleteCertificates()
	require.NoError(t, err)

	// Check if all the test certificates have been deleted
	certificates, err := getSdkClient(t).ForwardersService.ListCertificates()
	require.NoError(t, err)
	assert.Zero(t, len(certificates))
}

// Test CreateCertificate
func TestCreateAndDeleteCertificate(t *testing.T) {
	// Choose a name to identify the forwarder, create and generate a certificate
	forwarder := "forwarder_01"
	err, pemFile1, certificateInfo1 := createAndUploadCertificate(t, forwarder)
	require.NoError(t, err)
	require.NotNil(t, pemFile1)
	require.NotNil(t, certificateInfo1)

	defer deletePemFile(t, pemFile1)
	defer getSdkClient(t).ForwardersService.DeleteCertificates()

	//Delete certificate
	err = getSdkClient(t).ForwardersService.DeleteCertificate(strconv.FormatInt(*certificateInfo1.Slot, 10))
	assert.NoError(t, err)
}

// Create and Upload a test certificate to Splunk Forwarders Service
func createAndUploadCertificate(t *testing.T, forwarder string) (error, string, *forwarders.CertificateInfo) {
	// Generate the certificate
	err := generateCertificate(forwarder)
	require.NoError(t, err)

	// Generated certificate is in an external .pem file (deleted after test run finishes)
	pemFile := fmt.Sprintf("%v.pem", forwarder)

	fileBytes, err := ioutil.ReadFile(pemFile)
	if err != nil {
		return err, "", nil
	}

	str := string(fileBytes)
	cert := forwarders.Certificate{Pem: str}

	// Upload the certificate to Splunk Forwarders service
	certificateInfo, err := getSdkClient(t).ForwardersService.AddCertificate(cert)

	return err, pemFile, certificateInfo
}

// Generate the test certificate with the provided forwarder name
func generateCertificate(forwarderName string) error {
	// Generating a private key to sign the certificate
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey

	// Use the private key to generate a Certificate Signing Request (CSR), and then sign it with the private key
	certificateTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:   forwarderName,
			Country:      []string{"US"},
			Province:     []string{"CA"},
			Organization: []string{"Splunk"},
		},
		NotAfter: time.Now().AddDate(5, 0, 0),
	}
	certificate, err := x509.CreateCertificate(rand.Reader, certificateTemplate, certificateTemplate, publicKey, privateKey)
	if err != nil {
		return err
	}

	// Create the .pem certificate file and output the certificate created into .pem file
	pemFileName := fmt.Sprintf("%v.pem", forwarderName)
	pemFile, err := os.Create(pemFileName)
	if err != nil {
		return err
	}
	pem.Encode(pemFile, &pem.Block{Type: "CERTIFICATE", Bytes: certificate})
	pemFile.Close()

	return nil
}

// Deletes the test certificate .pem file after the test run finishes
func deletePemFile(t *testing.T, pemFileName string) {
	err := os.Remove(pemFileName)
	assert.Emptyf(t, err, "Error deleting the certificate file (.pem): %s", err)
}
