package mailcow_test

import (
	"github.com/stretchr/testify/assert"
	"go.bnck.me/mailcow"
	"testing"
)

func TestClient_Domain(t *testing.T) {
	c, err := mailcow.New("abc", "abc-123")
	assert.NoError(t, err)

	domainAPI := c.Domain()
	assert.NotNil(t, domainAPI)
}

func TestClient_Mailbox(t *testing.T) {
	c, err := mailcow.New("abc", "abc-123")
	assert.NoError(t, err)

	mailboxAPI := c.Mailbox()
	assert.NotNil(t, mailboxAPI)
}
