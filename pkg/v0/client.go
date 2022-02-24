package v0

import "net/http"

// Client is a type that can talk with the http.Handler returned by New().
type Client interface {
	CreateManifest(*Manifest) (*Manifest, error)
}

type client struct {
	baseURL    string
	httpClient *http.Client
}

// Dial establishes a connection to a server running at baseURL.
func Dial(baseURL string, httpClient *http.Client) (Client, error) {
	return &client{baseURL: baseURL, httpClient: httpClient}, nil
}

func (c *client) CreateManifest(*Manifest) (*Manifest, error) {
	return nil, nil
}
