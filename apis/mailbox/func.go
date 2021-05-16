package mailbox

import (
	"context"
	"go.bnck.me/mailcow/internal/transport"
)

func (c *client) All(ctx context.Context) ([]*Mailbox, error) {
	var out []*Mailbox
	if err := c.transport.Get(ctx, "/get/mailbox/all", &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) Create(ctx context.Context, createRequest CreateRequest) error {
	if err := c.transport.Post(ctx, "/add/mailbox", nil, transport.WithJSONRequestBody(createRequest)); err != nil {
		return err
	}
	return nil
}

func (c *client) Update(ctx context.Context, updateRequest UpdateRequest) error {
	if err := c.transport.Post(ctx, "/edit/mailbox", nil, transport.WithJSONRequestBody(updateRequest)); err != nil {
		return err
	}
	return nil
}

func (c *client) Get(ctx context.Context, mailbox string) (*Mailbox, error) {
	var out *Mailbox
	if err := c.transport.Get(ctx, "/get/mailbox/"+mailbox, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) Delete(ctx context.Context, mailboxes []string) error {
	if err := c.transport.Delete(ctx, "/delete/mailbox", nil, transport.WithJSONRequestBody(mailboxes)); err != nil {
		return err
	}
	return nil
}
