package mailcow

import (
	"github.com/jpbede/mailcowgo/apis/alias"
	"github.com/jpbede/mailcowgo/apis/domain"
	"github.com/jpbede/mailcowgo/apis/mailbox"
	"github.com/jpbede/mailcowgo/internal/transport"
)

// Client represents the main client
type Client struct {
	transport *transport.Client

	alias   alias.Client
	domain  domain.Client
	mailbox mailbox.Client
}

// Alias returns a client for domain specific operations
func (c *Client) Alias() alias.Client {
	if c.alias == nil {
		c.alias = alias.New(c.transport)
	}

	return c.alias
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
