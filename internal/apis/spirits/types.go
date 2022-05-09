package spirits

type Script struct {
	Text string `json:"text"`
}

type HTTP struct {
	URL                      string `json:"url"`
	CertificateAuthorityData []byte `json:"certificateAuthorityData"`
}
