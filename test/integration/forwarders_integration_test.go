package integration

import (
	"fmt"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// Test ListCertificates
func TestListCertificates(t *testing.T) {
	// Create a certificate
	certificateInfo1, err := getSdkClient(t).ForwardersService.CreateCertificate("forwarder_certificate_01.pem")
	require.Nil(t, err)
	require.NotNil(t, certificateInfo1)

	defer cleanupCertificate(t, certificateInfo1.Slot)

	// Create another certificate
	certificateInfo2, err := getSdkClient(t).ForwardersService.CreateCertificate("forwarder_certificate_02.pem")
	require.Nil(t, err)
	require.NotNil(t, certificateInfo2)

	defer cleanupCertificate(t, certificateInfo2.Slot)

	// Check if the test certificates were created as expected
	certificates, err := getSdkClient(t).ForwardersService.ListCertificates()
	require.Nil(t, err)
	assert.NotZero(t, len(certificates))
	fmt.Println(certificates)
}

/*// Test DeleteCertificates
func TestDeleteCertificates(t *testing.T) {
	// Create a certificate
	certificateInfo1, err := getSdkClient(t).ForwardersService.CreateCertificate("forwarder_certificate_01.pem")
	require.Nil(t, err)
	require.NotNil(t, certificateInfo1)

	defer cleanupCertificate(t, certificateInfo1.Slot)

	// Create another certificate
	certificateInfo2, err := getSdkClient(t).ForwardersService.CreateCertificate("forwarder_certificate_02.pem")
	require.Nil(t, err)
	require.NotNil(t, certificateInfo2)

	defer cleanupCertificate(t, certificateInfo2.Slot)

	// Delete all the certificates
	err = getSdkClient(t).ForwardersService.DeleteCertificates()
	require.Nil(t, err)

	// Check if all the test certificates have been deleted
	certificates, err := getSdkClient(t).ForwardersService.ListCertificates()
	require.Nil(t, err)
	assert.Zero(t, len(certificates))
}*/

// Cleans up the created certificate after the test has run
func cleanupCertificate(t *testing.T, slot int) {
	err := getSdkClient(t).ForwardersService.DeleteCertificate(slot)
	assert.Emptyf(t, err, "Error deleting certificate: %s", err)
}

// Test CreateCertificate
func TestCreateCertificate(t *testing.T) {
	client := getSdkClient(t)

	// Create a certificate
	certificateInfo, err := client.ForwardersService.CreateCertificate("forwarder_certificate_01.pem")
	require.Nil(t, err)
	require.NotNil(t, certificateInfo)

	defer cleanupCertificate(t, certificateInfo.Slot)
}

// Test DeleteCertificate
func TestDeleteCertificateIncorrectSlot(t *testing.T) {
	client := getSdkClient(t)

	// Create a certificate
	certificateInfo, err := client.ForwardersService.CreateCertificate("forwarder_certificate_01.pem")
	require.Nil(t, err)
	require.NotNil(t, certificateInfo)

	defer cleanupCertificate(t, certificateInfo.Slot)

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