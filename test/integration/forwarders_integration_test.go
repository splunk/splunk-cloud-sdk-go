// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/services/forwarders"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test ListCertificates
func TestListCertificates(t *testing.T) {
	// Certificate 1
	// Choose a name to identify the forwarder, create and generate a certificate
	forwarder := "forwarder_01"
	err, pemFile1, certificateInfo1 := createAndUploadCertificate(t, forwarder)
	require.Nil(t, err)
	require.NotNil(t, pemFile1)
	require.NotNil(t, certificateInfo1)

	defer deletePemFile(t, pemFile1)
	defer cleanupCertificate(t, certificateInfo1.Slot)

	// Certificate 2
	// Choose a name to identify the forwarder, create and generate a certificate
	forwarder = "forwarder_02"
	err, pemFile2, certificateInfo2 := createAndUploadCertificate(t, forwarder)
	require.Nil(t, err)
	require.NotNil(t, pemFile2)
	require.NotNil(t, certificateInfo2)

	defer deletePemFile(t, pemFile2)
	defer cleanupCertificate(t, certificateInfo2.Slot)

	// Check if the test certificates were created as expected
	certificates, err := getSdkClient(t).ForwardersService.ListCertificates()
	require.Nil(t, err)
	assert.NotZero(t, len(certificates))
}

// Test DeleteCertificates
func TestDeleteCertificates(t *testing.T) {
	// Certificate 1
	// Choose a name to identify the forwarder, create and generate a certificate
	forwarder := "forwarder_01"
	err, pemFile1, certificateInfo1 := createAndUploadCertificate(t, forwarder)
	require.Nil(t, err)
	require.NotNil(t, pemFile1)
	require.NotNil(t, certificateInfo1)

	defer deletePemFile(t, pemFile1)
	defer cleanupCertificate(t, certificateInfo1.Slot)

	// Certificate 2
	// Choose a name to identify the forwarder, create and generate a certificate
	forwarder = "forwarder_02"
	err, pemFile2, certificateInfo2 := createAndUploadCertificate(t, forwarder)
	require.Nil(t, err)
	require.NotNil(t, pemFile2)
	require.NotNil(t, certificateInfo2)

	defer deletePemFile(t, pemFile2)
	defer cleanupCertificate(t, certificateInfo2.Slot)

	// Delete all the certificates
	err = getSdkClient(t).ForwardersService.DeleteCertificates()
	require.Nil(t, err)

	// Check if all the test certificates have been deleted
	certificates, err := getSdkClient(t).ForwardersService.ListCertificates()
	require.Nil(t, err)
	assert.Zero(t, len(certificates))
}

// Cleans up the created certificate after the test has run
func cleanupCertificate(t *testing.T, slot int) {
	err := getSdkClient(t).ForwardersService.DeleteCertificate(slot)
	assert.Emptyf(t, err, "Error deleting certificate: %s", err)
}

// Test CreateCertificate
func TestCreateCertificate(t *testing.T) {
	// Choose a name to identify the forwarder, create and generate a certificate
	forwarder := "forwarder_01"
	err, pemFile1, certificateInfo1 := createAndUploadCertificate(t, forwarder)
	require.Nil(t, err)
	require.NotNil(t, pemFile1)
	require.NotNil(t, certificateInfo1)

	defer deletePemFile(t, pemFile1)
	defer cleanupCertificate(t, certificateInfo1.Slot)
}

// Test DeleteCertificate
func TestDeleteCertificateIncorrectSlot(t *testing.T) {
	// Choose a name to identify the forwarder, create and generate a certificate
	forwarder := "forwarder_01"
	err, pemFile1, certificateInfo1 := createAndUploadCertificate(t, forwarder)
	require.Nil(t, err)
	require.NotNil(t, pemFile1)
	require.NotNil(t, certificateInfo1)

	defer deletePemFile(t, pemFile1)
	defer cleanupCertificate(t, certificateInfo1.Slot)

	// Delete the certificate. Provide a negative slot number
	invalidSlot := -1
	err = getSdkClient(t).ForwardersService.DeleteCertificate(invalidSlot)
	assert.NotEmpty(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 400, httpErr.HTTPStatusCode)
	assert.Equal(t, "Must specify a slot between 1 and max supported certificates", httpErr.Message)
	assert.Equal(t, "BAD_CERTIFICATE", httpErr.Code)
}

// Create and Upload a test certificate to Splunk Forwarders Service
func createAndUploadCertificate(t *testing.T, forwarder string) (error, string, *forwarders.CertificateInfo) {
	// Generate the certificate
	err := generateCertificate(forwarder)
	require.Nil(t, err)

	// Generated certificate is in an external .pem file (deleted after test run finishes)
	pemFile := fmt.Sprintf("%v.pem", forwarder)

	// Upload the certificate to Splunk Forwarders service
	certificateInfo, err := getSdkClient(t).ForwardersService.CreateCertificate(pemFile)

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
