package forwarders

// Servicer ...
type Servicer interface {
	// ListCertificates lists all the certificates for the tenant
	ListCertificates() ([]CertificateInfo, error)
	// CreateCertificate creates and adds a certificate to a vacant slot on the tenant
	CreateCertificate(certificateFileName string) (*CertificateInfo, error)
	// DeleteCertificate deletes a certificate on a particular slot on a tenant
	DeleteCertificate(slot int) error
	// DeleteCertificates deletes all the certificates on a tenant
	DeleteCertificates() error
}
