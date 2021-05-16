package mailcow_test

import (
	"github.com/stretchr/testify/assert"
	"go.bnck.me/mailcow"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	c, err := mailcow.New("abc", "abc-123")
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestNewWithClient(t *testing.T) {
	c, err := mailcow.NewWithClient(&http.Client{}, "abc", "abc-123")
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestNewWithOptions(t *testing.T) {
	c, err := mailcow.NewWithOptions(mailcow.WithHost("https://mx.awesome.email"), mailcow.WithAPIKey("abc"))
	assert.NoError(t, err)
	assert.NotNil(t, c)
}
