package domain

import (
	"context"
	"go.bnck.me/mailcow/internal/transport"
)

func (c *client) All(ctx context.Context) ([]*Domain, error) {
	var out []*Domain
	if err := c.transport.Get(ctx, "/get/domain/all", &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) Get(ctx context.Context, domain string) (*Domain, error) {
	var out *Domain
	if err := c.transport.Get(ctx, "/get/domain/"+domain, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) Delete(ctx context.Context, domains []string) error {
	if err := c.transport.Delete(ctx, "/delete/domain", nil, transport.WithJSONRequestBody(domains)); err != nil {
		return err
	}
	return nil
}
