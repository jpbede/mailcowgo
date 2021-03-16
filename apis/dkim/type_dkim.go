package dkim

type DKIM struct {
	Pubkey       string `json:"pubkey"`
	Length       string `json:"length"`
	DkimTxt      string `json:"dkim_txt"`
	DkimSelector string `json:"dkim_selector"`
	Privkey      string `json:"privkey"`
}
