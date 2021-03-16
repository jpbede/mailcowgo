package mailbox

import "github.com/jpbede/mailcowgo/internal/transport"

type Mailbox struct {
	MaxNewQuota  int64              `json:"max_new_quota"`
	Username     string             `json:"username"`
	Rl           bool               `json:"rl"`
	IsRelayed    int                `json:"is_relayed"`
	Name         string             `json:"name"`
	Active       int                `json:"active"`
	Domain       string             `json:"domain"`
	LocalPart    string             `json:"local_part"`
	Quota        int64              `json:"quota"`
	Attributes   *Attributes        `json:"attributes"`
	QuotaUsed    int                `json:"quota_used"`
	PercentInUse transport.StrInt64 `json:"percent_in_use"`
	Messages     int                `json:"messages"`
	SpamAliases  int                `json:"spam_aliases"`
	PercentClass string             `json:"percent_class"`
}

type Attributes struct {
	ForcePwUpdate          string `json:"force_pw_update"`
	TLSEnforceIn           string `json:"tls_enforce_in"`
	TLSEnforceOut          string `json:"tls_enforce_out"`
	SogoAccess             string `json:"sogo_access"`
	MailboxFormat          string `json:"mailbox_format"`
	QuarantineNotification string `json:"quarantine_notification"`
}
