# forwarders
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/forwarders"


## Usage

#### type Certificate

```go
type Certificate struct {
	Pem *string `json:"pem,omitempty"`
}
```


#### type CertificateInfo

```go
type CertificateInfo struct {
	Content    *string `json:"content,omitempty"`
	Hash       *string `json:"hash,omitempty"`
	Issuer     *string `json:"issuer,omitempty"`
	LastUpdate *string `json:"lastUpdate,omitempty"`
	NotAfter   *string `json:"notAfter,omitempty"`
	NotBefore  *string `json:"notBefore,omitempty"`
	Slot       *int64  `json:"slot,omitempty"`
	Subject    *string `json:"subject,omitempty"`
}
```


#### type Error

```go
type Error struct {
	Code    *string                `json:"code,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
	Message *string                `json:"message,omitempty"`
}
```


#### type Service

```go
type Service services.BaseService
```


#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new forwarders service client from the given Config

#### func (*Service) AddCertificate

```go
func (s *Service) AddCertificate(certificate Certificate, resp ...*http.Response) (*CertificateInfo, error)
```
AddCertificate - Adds a certificate to a vacant slot on a tenant. Each tenant
can have up to five certificates. Parameters:

    certificate
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteCertificate

```go
func (s *Service) DeleteCertificate(slot string, resp ...*http.Response) error
```
DeleteCertificate - Removes a certificate on a particular slot on a tenant.
Parameters:

    slot
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteCertificates

```go
func (s *Service) DeleteCertificates(resp ...*http.Response) error
```
DeleteCertificates - Removes all certificates on a tenant. Parameters:

    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListCertificates

```go
func (s *Service) ListCertificates(resp ...*http.Response) ([]CertificateInfo, error)
```
ListCertificates - Returns a list of all certificates for a tenant. Parameters:

    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### type Servicer

```go
type Servicer interface {
	/*
		AddCertificate - Adds a certificate to a vacant slot on a tenant.
		Each tenant can have up to five certificates.
		Parameters:
			certificate
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	AddCertificate(certificate Certificate, resp ...*http.Response) (*CertificateInfo, error)
	/*
		DeleteCertificate - Removes a certificate on a particular slot on a tenant.
		Parameters:
			slot
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteCertificate(slot string, resp ...*http.Response) error
	/*
		DeleteCertificates - Removes all certificates on a tenant.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteCertificates(resp ...*http.Response) error
	/*
		ListCertificates - Returns a list of all certificates for a tenant.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListCertificates(resp ...*http.Response) ([]CertificateInfo, error)
}
```

Servicer represents the interface for implementing all endpoints for this
service
