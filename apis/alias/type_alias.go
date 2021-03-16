package alias

import (
	"github.com/jpbede/mailcowgo/internal/transport"
)

type Alias struct {
	InPrimaryDomain string             `json:"in_primary_domain"`
	ID              int                `json:"id"`
	Domain          string             `json:"domain"`
	PublicComment   string             `json:"public_comment"`
	PrivateComment  string             `json:"private_comment"`
	Goto            string             `json:"goto"`
	Address         string             `json:"address"`
	IsCatchAll      transport.StrInt64 `json:"is_catch_all"`
	Active          transport.StrInt64 `json:"active"`
	Created         string             `json:"created"`
	Modified        string             `json:"modified"`
}
