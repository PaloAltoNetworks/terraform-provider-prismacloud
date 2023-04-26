package gcpTemplate

type GcpTemplateReq struct {
	ProjectId            string   `json:"projectId"`
	OrgId                string   `json:"orgId"`
	AccountType          string   `json:"accountType"`
	AuthenticationType   string   `json:"authenticationType"`
	Name                 string   `json:"name"`
	Features             []string `json:"features"`
	FileName             string   `json:"fileName"`
	FlowLogStorageBucket string   `json:"flowLogStorageBucket"`
}

type GcpTemplate struct {
	ProjectId            string   `json:"projectId"`
	OrgId                string   `json:"orgId"`
	AccountType          string   `json:"accountType"`
	AuthenticationType   string   `json:"authenticationType"`
	Name                 string   `json:"name"`
	Features             []string `json:"features"`
	FileName             string   `json:"fileName"`
	FlowLogStorageBucket string   `json:"flowLogStorageBucket"`
}
