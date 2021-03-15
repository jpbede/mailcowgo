package mailcow

import (
	"github.com/rs/zerolog"
	"net/http"
)

// ClientOption options for creating the client
type ClientOption func(*Client)

// WithHost supplies the base api host url
func WithHost(host string) ClientOption {
	return func(c *Client) {
		c.transport.BaseURL = host
	}
}

// WithHTTPClient supplies a already created http.Client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.transport.HTTPClient = httpClient
	}
}

// WithLogger supplies a logger for the transport client
func WithLogger(logger *zerolog.Logger) ClientOption {
	return func(c *Client) {
		c.transport.SetLogger(logger)
	}
}

// WithAPIKey supplies the api credentials
func WithAPIKey(apiKey string) ClientOption {
	return func(c *Client) {
		c.transport.APIKey = apiKey
	}
}
