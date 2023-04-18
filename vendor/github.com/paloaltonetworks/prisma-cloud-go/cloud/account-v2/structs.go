package accountv2

type AccountAndCredentials struct {
	Account AccountResponse `json:"cloudAccount"`
}

type NameTypeId struct {
	Name      string `json:"name"`
	CloudType string `json:"cloudType"`
	AccountId string `json:"id"`
}

type CloudAccountResp struct {
	AccountId               string      `json:"accountId"`
	Name                    string      `json:"name"`
	CloudType               string      `json:"cloudType"`
	Enabled                 bool        `json:"enabled"`
	ParentId                string      `json:"parentId"`
	AccountType             string      `json:"accountType"`
	Deleted                 bool        `json:"deleted"`
	ProtectionMode          string      `json:"protectionMode"`
	DeploymentType          string      `json:"deploymentType"`
	CustomerName            string      `json:"customerName"`
	CreatedEpochMillis      int         `json:"createdEpochMillis"`
	LastModifiedEpochMillis int         `json:"lastModifiedEpochMillis"`
	LastModifiedBy          string      `json:"lastModifiedBy"`
	Features                []Features1 `json:"features"`
}

type AccountResponse struct {
	CloudAccountResp CloudAccountResp `json:"cloudAccount"`
	RoleArn          string           `json:"roleArn"`
	ExternalId       string           `json:"externalId"`
	HasMemberRole    bool             `json:"hasMemberRole"`
	TemplateUrl      string           `json:"templateUrl"`
	GroupIds         []string         `json:"groupIds"`
}

type Group struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Aws struct {
	AccountId   string     `json:"accountId"`
	AccountType string     `json:"accountType"`
	Enabled     bool       `json:"enabled"`
	Features    []Features `json:"features"`
	GroupIds    []string   `json:"groupIds"`
	Name        string     `json:"name"`
	RoleArn     string     `json:"roleArn"`
}

type AwsV2 struct {
	CloudAccountResp          CloudAccountResp `json:"cloudAccount"`
	Name                      string           `json:"name"`
	RoleArn                   string           `json:"roleArn"`
	ExternalId                string           `json:"externalId"`
	HasMemberRole             bool             `json:"hasMemberRole"`
	TemplateUrl               string           `json:"templateUrl"`
	GroupIds                  []string         `json:"groupIds"`
	EventbridgeRuleNamePrefix string           `json:"eventBridgeRuleNamePrefix"`
}

type Features1 struct {
	Name  string `json:"featureName"`
	State string `json:"featureState"`
}

type Features struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

type CloudAccount struct {
	AccountId      string   `json:"accountId"`
	Enabled        bool     `json:"enabled"`
	GroupIds       []string `json:"groupIds"`
	Name           string   `json:"name"`
	ProtectionMode string   `json:"protectionMode"`
	AccountType    string   `json:"accountType"`
}

//Azurev2

type CloudAccountAzureResp struct {
	AccountId                 string      `json:"accountId"`
	Name                      string      `json:"name"`
	CloudType                 string      `json:"cloudType"`
	Enabled                   bool        `json:"enabled"`
	ParentId                  string      `json:"parentId"`
	AccountType               string      `json:"accountType"`
	Deleted                   bool        `json:"deleted"`
	ProtectionMode            string      `json:"protectionMode"`
	DeploymentType            string      `json:"deploymentType"`
	CustomerName              string      `json:"customerName"`
	CreatedEpochMillis        int         `json:"createdEpochMillis"`
	LastModifiedEpochMillis   int         `json:"lastModifiedEpochMillis"`
	LastModifiedBy            string      `json:"lastModifiedBy"`
	Features                  []Features1 `json:"features"`
	DeploymentTypeDescription string      `json:"deploymentTypeDescription"`
}
type AzureAccountResponse struct {
	CloudAccountAzureResp CloudAccountAzureResp `json:"cloudAccount"`
	TenantId              string                `json:"tenantId"`
	ServicePrincipalId    string                `json:"servicePrincipalId"`
	ClientId              string                `json:"clientId"`
	TemplateUrl           string                `json:"templateUrl"`
	Key                   string                `json:"key"`
	GroupIds              []string              `json:"groupIds"`
	MonitorFlowLogs       bool                  `json:"monitorFlowLogs"`
	EnvironmentType       string                `json:"environmentType"`
}
type AzureV2 struct {
	CloudAccountAzureResp CloudAccountAzureResp `json:"cloudAccount"`
	ClientId              string                `json:"clientId"`
	EnvironmentType       string                `json:"environmentType"`
	Key                   string                `json:"key"`
	MonitorFlowLogs       bool                  `json:"monitorFlowLogs"`
	ServicePrincipalId    string                `json:"servicePrincipalId"`
	GroupIds              []string              `json:"groupIds"`
	TemplateUrl           string                `json:"templateUrl"`
	TenantId              string                `json:"tenantId"`
	DeploymentType        string                `json:"deploymentType"`
}

type Azure struct {
	CloudAccountAzure  CloudAccountAzure `json:"cloudAccount"`
	ClientId           string            `json:"clientId"`
	Key                string            `json:"key"`
	MonitorFlowLogs    bool              `json:"monitorFlowLogs"`
	TenantId           string            `json:"tenantId"`
	ServicePrincipalId string            `json:"servicePrincipalId"`
	EnvironmentType    string            `json:"environmentType"`
	Features           []Features        `json:"features"`
	Enabled            bool              `json:"enabled"`
}

type CloudAccountAzure struct {
	AccountId   string   `json:"accountId"`
	Enabled     bool     `json:"enabled"`
	Name        string   `json:"name"`
	AccountType string   `json:"accountType"`
	GroupIds    []string `json:"groupIds"`
}
