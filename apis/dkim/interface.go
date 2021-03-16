package dkim

import "context"

type Client interface {
	Get(ctx context.Context, domain string) (*DKIM, error)
	Generate(ctx context.Context, request GenerateRequest) error
	Duplicate(ctx context.Context, fromDomain string, toDomain string) error
	Delete(ctx context.Context, mailboxes []string) error
}
