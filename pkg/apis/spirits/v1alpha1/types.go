package v1alpha1

// Script holds the source lines of a script to be used in the context of the spirits API
type Script struct {
	// APIVersion is the API version that that script source expects for serializing script input/output
	APIVersion string `json:"apiVersion"`

	// Text holds the text of the script
	Text string `json:"text"`
}

// HTTP holds information about an HTTP endpoint
type HTTP struct {
	// URL holds the location of the HTTP endpoint; the endpoint must use HTTPS
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern=`^https://`
	URL string `json:"url"`

	// CertificateAuthorityData holds the optional CA bundle data used to validate TLS connections
	// to the URL
	// +optional
	CertificateAuthorityData []byte `json:"certificateAuthorityData"`
}
