# forwarders
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/forwarders"


## Usage

#### type CertificateInfo

```go
type CertificateInfo struct {
	Content string `json:"content,omitempty"`

	Hash string `json:"hash,omitempty"`

	Issuer string `json:"issuer,omitempty"`

	LastUpdate string `json:"lastUpdate,omitempty"`

	NotAfter string `json:"notAfter,omitempty"`

	NotBefore string `json:"notBefore,omitempty"`

	Slot int `json:"slot,omitempty"`

	Subject string `json:"subject,omitempty"`
}
```

CertificateInfo represents a Forwarder certificate

#### type PemFile

```go
type PemFile struct {
	Pem string `json:"pem,omitempty"`
}
```

PemFile represents the contents of the .pem certificate file

#### type Service

```go
type Service services.BaseService
```

Service talks to the Splunk Cloud catalog service

#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new forwarders service client from the given Config

#### func (*Service) CreateCertificate

```go
func (s *Service) CreateCertificate(certificateFileName string) (*CertificateInfo, error)
```
CreateCertificate creates and adds a certificate to a vacant slot on the tenant.

#### func (*Service) DeleteCertificate

```go
func (s *Service) DeleteCertificate(slot int) error
```
DeleteCertificate deletes a certificate on a particular slot on a tenant

#### func (*Service) DeleteCertificates

```go
func (s *Service) DeleteCertificates() error
```
DeleteCertificates deletes all the certificates on a tenant

#### func (*Service) ListCertificates

```go
func (s *Service) ListCertificates() ([]CertificateInfo, error)
```
ListCertificates lists all the certificates for the tenant

#### type Servicer

```go
type Servicer interface {
	// ListCertificates lists all the certificates for the tenant
	ListCertificates() ([]CertificateInfo, error)
	// DeleteCertificate deletes a certificate on a particular slot on a tenant
	DeleteCertificate(slot int) error
	// DeleteCertificates deletes all the certificates on a tenant
	DeleteCertificates() error
	// CreateCertificate creates and adds a certificate to a vacant slot on the tenant.
	CreateCertificate(certificateFileName string) (*CertificateInfo, error)
}
```

Servicer ...
