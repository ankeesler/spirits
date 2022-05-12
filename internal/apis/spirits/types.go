package spirits

type Script struct {
	APIVersion string `json:"apiVersion"`
	Text       string `json:"text"`
}

type HTTP struct {
	URL                      string `json:"url"`
	CertificateAuthorityData []byte `json:"certificateAuthorityData"`
}
