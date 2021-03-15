package mailbox

import "context"

type Client interface {
	All(ctx context.Context) ([]*Mailbox, error)
	Get(ctx context.Context, mailbox string) (*Mailbox, error)
	Create(ctx context.Context, createRequest CreateRequest) error
	Update(ctx context.Context, updateRequest UpdateRequest) error
	Delete(ctx context.Context, mailboxes []string) error
}
