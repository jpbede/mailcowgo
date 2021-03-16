package dkim

type DuplicateRequest struct {
	FromDomain string `json:"from_domain"`
	ToDomain   string `json:"to_domain"`
}
