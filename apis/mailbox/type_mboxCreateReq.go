package mailbox

type CreateRequest struct {
	LocalPart     string `json:"local_part"`
	Domain        string `json:"domain"`
	Name          string `json:"name"`
	Quota         string `json:"quota"`
	Password      string `json:"password"`
	Password2     string `json:"password2"`
	Active        string `json:"active"`
	ForcePwUpdate string `json:"force_pw_update"`
	TLSEnforceIn  string `json:"tls_enforce_in"`
	TLSEnforceOut string `json:"tls_enforce_out"`
}
