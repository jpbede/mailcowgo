package mailcow

import (
	"errors"
	"github.com/jpbede/mailcowgo/internal/transport"
	"net/http"
)

// New creates a new Client with username, password and context
func New(apiHost, apiKey string) (*Client, error) {
	return NewWithOptions(WithHost(apiHost), WithAPIKey(apiKey))
}

// NewWithClient creates a new Client with a given http.Client
func NewWithClient(httpClient *http.Client, apiHost, apiKey string) (*Client, error) {
	return NewWithOptions(WithHost(apiHost), WithHTTPClient(httpClient), WithAPIKey(apiKey))
}

// NewWithOptions creates a new Client with given options
func NewWithOptions(options ...ClientOption) (*Client, error) {
	c := &Client{}

	// always create a base transport it can be overwritten with options
	c.transport = transport.New("")

	// run given options
	for _, option := range options {
		option(c)
	}

	// check if there are credentials
	if !c.transport.HasCredentials() {
		return nil, errors.New("no api credentials supplied")
	}

	if c.transport.BaseURL == "" {
		return nil, errors.New("no api host supplied")
	}

	return c, nil
}
