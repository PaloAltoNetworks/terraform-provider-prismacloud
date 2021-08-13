package org

type AccountAndCredentials struct {
	Account OrgAccount `json:"cloudAccount"`
}

type NameTypeId struct {
	Name      string `json:"name"`
	CloudType string `json:"cloudType"`
	AccountId string `json:"id"`
}

type OrgAccount struct {
	Name                  string   `json:"name"`
	CloudType             string   `json:"cloudType"`
	AccountType           string   `json:"accountType"`
	Enabled               bool     `json:"enabled"`
	LastModifiedTs        int      `json:"lastModifiedTs"`
	LastModifiedBy        string   `json:"lastModifiedBy"`
	StorageScanEnabled    bool     `json:"storageScanEnabled"`
	ProtectionMode        string   `json:"protectionMode"`
	IngestionMode         int      `json:"ingestionMode"`
	GroupIds              []string `json:"groupIds"`
	Groups                []Group  `json:"groups"`
	Status                string   `json:"status"`
	NumberOfChildAccounts int      `json:"numberOfChildAccounts"`
	AccountId             string   `json:"accountId,omitempty"`
	AddedOn               int      `json:"addedOn"`
}

type Group struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type AwsOrg struct {
	AccountId        string   `json:"accountId"`
	Enabled          bool     `json:"enabled"`
	ExternalId       string   `json:"externalId"`
	GroupIds         []string `json:"groupIds"`
	Name             string   `json:"name"`
	RoleArn          string   `json:"roleArn"`
	ProtectionMode   string   `json:"protectionMode"`
	AccountType      string   `json:"accountType"`
	MemberRoleName   string   `json:"memberRoleName"`
	MemberExternalId string   `json:"memberExternalId"`
	MemberRoleStatus bool     `json:"memberRoleStatus"`
}

type GcpCloudAccount struct {
	AccountId      string   `json:"accountId"`
	Enabled        bool     `json:"enabled"`
	ProjectId      string   `json:"projectId"`
	GroupIds       []string `json:"groupIds"`
	Name           string   `json:"name"`
	ProtectionMode string   `json:"protectionMode"`
	AccountType    string   `json:"accountType"`
}

type AzureCloudAccount struct {
	AccountId      string   `json:"accountId"`
	Enabled        bool     `json:"enabled"`
	GroupIds       []string `json:"groupIds"`
	Name           string   `json:"name"`
	ProtectionMode string   `json:"protectionMode"`
	AccountType    string   `json:"accountType"`
}

type AzureOrg struct {
	Account            AzureCloudAccount `json:"cloudAccount"`
	ClientId           string            `json:"clientId"`
	Key                string            `json:"key"`
	MonitorFlowLogs    bool              `json:"monitorFlowLogs"`
	TenantId           string            `json:"tenantId"`
	ServicePrincipalId string            `json:"servicePrincipalId"`
}

type GcpOrgCredentials struct {
	Type            string `json:"type"`
	ProjectId       string `json:"project_id"`
	PrivateKeyId    string `json:"private_key_id"`
	PrivateKey      string `json:"private_key"`
	ClientEmail     string `json:"client_email"`
	ClientId        string `json:"client_id"`
	AuthUri         string `json:"auth_uri"`
	TokenUri        string `json:"token_uri"`
	ProviderCertUrl string `json:"auth_provider_x509_cert_url"`
	ClientCertUrl   string `json:"client_x509_cert_url"`
}

type HierarchySelection struct {
	ResourceId    string `json:"resourceId"`
	DisplayName   string `json:"displayName"`
	NodeType      string `json:"nodeType"`
	SelectionType string `json:"selectionType"`
}

type GcpOrg struct {
	Account                  GcpCloudAccount      `json:"cloudAccount"`
	CompressionEnabled       bool                 `json:"compressionEnabled"`
	DataflowEnabledProject   string               `json:"dataflowEnabledProject"`
	FlowLogStorageBucket     string               `json:"flowLogStorageBucket"`
	OrganizationName         string               `json:"organizationName"`
	AccountGroupCreationMode string               `json:"accountGroupCreationMode"`
	Credentials              GcpOrgCredentials    `json:"credentials"`
	HierarchySelection       []HierarchySelection `json:"hierarchySelection"`
}
type Oci struct {
	AccountId             string `json:"accountId"`
	AccountType           string `json:"accountType"`
	Enabled               bool   `json:"enabled"`
	Name                  string `json:"name"`
	DefaultAccountGroupId string `json:"defaultAccountGroupId"`
	GroupName             string `json:"groupName"`
	HomeRegion            string `json:"homeRegion"`
	PolicyName            string `json:"policyName"`
	UserName              string `json:"userName"`
	UserOcid              string `json:"userOcid"`
}
