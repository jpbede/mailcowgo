package alias

import (
	"context"
	"go.bnck.me/mailcow/internal/transport"
)

func (c *client) All(ctx context.Context) ([]*Alias, error) {
	var out []*Alias
	if err := c.transport.Get(ctx, "/get/alias/all", &out); err != nil {
		return nil, err
	}
	return out, nil
}

/*func (c *client) Create(ctx context.Context, createRequest CreateRequest) error {
	if err := c.transport.Post(ctx, "/add/alias", nil, transport.WithJSONRequestBody(createRequest)); err != nil {
		return err
	}
	return nil
}

func (c *client) Update(ctx context.Context, updateRequest UpdateRequest) error {
	if err := c.transport.Post(ctx, "/edit/alias", nil, transport.WithJSONRequestBody(updateRequest)); err != nil {
		return err
	}
	return nil
}*/

func (c *client) Get(ctx context.Context, mailbox string) (*Alias, error) {
	var out *Alias
	if err := c.transport.Get(ctx, "/get/alias/"+mailbox, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) Delete(ctx context.Context, mailboxes []string) error {
	if err := c.transport.Delete(ctx, "/delete/alias", nil, transport.WithJSONRequestBody(mailboxes)); err != nil {
		return err
	}
	return nil
}
