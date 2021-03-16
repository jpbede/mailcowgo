package dkim

type GenerateRequest struct {
	Domains  []string `json:"domains"`
	Selector string   `json:"dkim_selector"`
	KeySize  int      `json:"key_size"`
}
