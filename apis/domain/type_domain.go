package domain

import "go.bnck.me/mailcow/internal/transport"

type Domain struct {
	MaxNewMailboxQuota       int64              `json:"max_new_mailbox_quota"`
	DefNewMailboxQuota       int64              `json:"def_new_mailbox_quota"`
	QuotaUsedInDomain        string             `json:"quota_used_in_domain"`
	BytesTotal               transport.StrInt64 `json:"bytes_total"`
	MessagesTotal            transport.StrInt64 `json:"msgs_total"`
	MailboxesInDomain        int                `json:"mboxes_in_domain"`
	MailboxesLeft            int                `json:"mboxes_left"`
	DomainName               string             `json:"domain_name"`
	Description              string             `json:"description"`
	MaxNumAliasesForDomain   int                `json:"max_num_aliases_for_domain"`
	MaxNumMailboxesForDomain int                `json:"max_num_mboxes_for_domain"`
	DefQuotaForMailbox       int64              `json:"def_quota_for_mbox"`
	MaxQuotaForMailbox       int64              `json:"max_quota_for_mbox"`
	MaxQuotaForDomain        int64              `json:"max_quota_for_domain"`
	RelayHost                string             `json:"relayhost"`
	BackupMX                 int                `json:"backupmx"`
	Gal                      int                `json:"gal"`
	Lang                     string             `json:"lang"`
	Rl                       bool               `json:"rl"`
	Active                   transport.StrInt64 `json:"active"`
	RelayAllRecipients       transport.StrInt64 `json:"relay_all_recipients"`
	AliasesInDomain          int                `json:"aliases_in_domain"`
	AliasesLeft              int                `json:"aliases_left"`
}
