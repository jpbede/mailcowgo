package dkim

import (
	"context"
	"github.com/jpbede/mailcowgo/internal/transport"
)

func (c *client) Get(ctx context.Context, domain string) (*DKIM, error) {
	var out *DKIM
	if err := c.transport.Get(ctx, "/get/dkim/"+domain, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *client) Generate(ctx context.Context, req GenerateRequest) error {
	if err := c.transport.Post(ctx, "/add/dkim", nil, transport.WithJSONRequestBody(req)); err != nil {
		return err
	}
	return nil
}

func (c *client) Duplicate(ctx context.Context, fromDomain string, toDomain string) error {
	req := DuplicateRequest{
		FromDomain: fromDomain,
		ToDomain:   toDomain,
	}
	if err := c.transport.Post(ctx, "/add/dkim", nil, transport.WithJSONRequestBody(req)); err != nil {
		return err
	}
	return nil
}

func (c *client) Delete(ctx context.Context, dkimKeys []string) error {
	if err := c.transport.Delete(ctx, "/delete/dkim", nil, transport.WithJSONRequestBody(dkimKeys)); err != nil {
		return err
	}
	return nil
}
