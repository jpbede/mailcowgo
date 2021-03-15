package mailcow

import (
	"github.com/jpbede/mailcowgo/apis/domain"
	"github.com/jpbede/mailcowgo/apis/mailbox"
	"github.com/jpbede/mailcowgo/internal/transport"
)

// Client represents the main client
type Client struct {
	transport *transport.Client

	domain  domain.Client
	mailbox mailbox.Client
}

// Domain returns a client for domain specific operations
func (c *Client) Domain() domain.Client {
	if c.domain == nil {
		c.domain = domain.New(c.transport)
	}

	return c.domain
}

// Mailbox returns a client for mailbox specific operations
func (c *Client) Mailbox() mailbox.Client {
	if c.mailbox == nil {
		c.mailbox = mailbox.New(c.transport)
	}

	return c.mailbox
}
