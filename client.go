package mailcow

import (
	"go.bnck.me/mailcow/apis/alias"
	"go.bnck.me/mailcow/apis/domain"
	"go.bnck.me/mailcow/apis/mailbox"
	"go.bnck.me/mailcow/internal/transport"
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
