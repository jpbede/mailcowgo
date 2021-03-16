package alias

import "context"

type Client interface {
	All(ctx context.Context) ([]*Alias, error)
	Get(ctx context.Context, mailbox string) (*Alias, error)
	//Create(ctx context.Context, createRequest CreateRequest) error
	//Update(ctx context.Context, updateRequest UpdateRequest) error
	Delete(ctx context.Context, mailboxes []string) error
}
