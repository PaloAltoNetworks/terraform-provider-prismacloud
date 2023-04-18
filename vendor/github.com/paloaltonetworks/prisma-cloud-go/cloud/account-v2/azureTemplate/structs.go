package azureTemplate

type AzureTemplateReq struct {
	SubscriptionId  string   `json:"subscriptionId"`
	TenantId        string   `json:"tenantId"`
	AccountType     string   `json:"accountType"`
	DeploymentType  string   `json:"deploymentType"`
	RootSyncEnabled bool     `json:"rootSyncEnabled"`
	Features        []string `json:"features"`
	FileName        string   `json:"fileName"`
}

type AzureTemplate struct {
	SubscriptionId  string   `json:"subscriptionId"`
	TenantId        string   `json:"tenantId"`
	AccountType     string   `json:"accountType"`
	DeploymentType  string   `json:"deploymentType"`
	RootSyncEnabled bool     `json:"rootSyncEnabled"`
	Features        []string `json:"features"`
	FileName        string   `json:"fileName"`
}
