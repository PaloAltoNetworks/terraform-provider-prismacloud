package ibmTemplate

type IbmTemplateReq struct {
	AccountType string `json:"accountType"`
	FileName    string `json:"fileName"`
}

type IbmTemplate struct {
	AccountType string `json:"accountType"`
	FileName    string `json:"fileName"`
}
