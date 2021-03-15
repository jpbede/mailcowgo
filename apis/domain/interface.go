package domain

import "context"

type Client interface {
	All(ctx context.Context) ([]*Domain, error)
	Get(ctx context.Context, domain string) (*Domain, error)
	//Create(ctx context.Context)
	//Update(ctx context.Context)
	Delete(ctx context.Context, domains []string) error
}
