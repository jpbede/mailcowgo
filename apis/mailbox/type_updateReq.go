package mailbox

type UpdateRequest struct {
	Items []string          `json:"items"`
	Attr  *UpdateAttributes `json:"attr"`
}

type UpdateAttributes struct {
	Name          string   `json:"name"`
	Quota         string   `json:"quota"`
	Password      string   `json:"password"`
	Password2     string   `json:"password2"`
	Active        string   `json:"active"`
	SenderACL     []string `json:"sender_acl"`
	ForcePwUpdate string   `json:"force_pw_update"`
	SogoAccess    string   `json:"sogo_access"`
}
