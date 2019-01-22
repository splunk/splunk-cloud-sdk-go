package forwarders

// CertificateInfo represents a Forwarder certificate
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

// PemFile represents the contents of the .pem certificate file
type PemFile struct {
	Pem string `json:"pem,omitempty"`
}
