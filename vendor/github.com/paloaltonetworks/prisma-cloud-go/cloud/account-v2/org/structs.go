package org

type AccountAndCredentials struct {
	Account OrgAccount `json:"cloudAccount"`
}

type NameTypeId struct {
	Name      string `json:"name"`
	CloudType string `json:"cloudType"`
	AccountId string `json:"id"`
}

type CloudAccountResp struct {
	AccountId               string      `json:"accountId"`
	Name                    string      `json:"name"`
	AccountTypeId           int         `json:"accountTypeId"`
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
	CloudAccountResp        CloudAccountResp     `json:"cloudAccount"`
	RoleArn                 string               `json:"roleArn"`
	ExternalId              string               `json:"externalId"`
	HasMemberRole           bool                 `json:"hasMemberRole"`
	TemplateUrl             string               `json:"templateUrl"`
	GroupIds                []string             `json:"groupIds"`
	HierarchySelection      []HierarchySelection `json:"hierarchySelection"`
	DefaultAccountGroupId   string               `json:"defaultAccountGroupId"`
	DefaultAccountGroupName string               `json:"defaultAccountGroupName"`
	MemberRoleName          string               `json:"memberRoleName"`
	MemberExternalId        string               `json:"memberExternalId"`
	MemberTemplateUrl       string               `json:"memberTemplateUrl"`
}

type AwsOrg struct {
	AccountId             string               `json:"accountId"`
	Enabled               bool                 `json:"enabled"`
	DefaultAccountGroupId string               `json:"defaultAccountGroupId"`
	Name                  string               `json:"name"`
	RoleArn               string               `json:"roleArn"`
	AccountType           string               `json:"accountType"`
	GroupIds              []string             `json:"groupIds"`
	HierarchySelection    []HierarchySelection `json:"hierarchySelection"`
	Features              []Features           `json:"features"`
}

type AwsOrgV2 struct {
	CloudAccountResp          CloudAccountResp     `json:"cloudAccount"`
	Name                      string               `json:"name"`
	RoleArn                   string               `json:"roleArn"`
	ExternalId                string               `json:"externalId"`
	HasMemberRole             bool                 `json:"hasMemberRole"`
	TemplateUrl               string               `json:"templateUrl"`
	GroupIds                  []string             `json:"groupIds"`
	EventbridgeRuleNamePrefix string               `json:"eventbridgeRuleNamePrefix"`
	DefaultAccountGroupId     string               `json:"defaultAccountGroupId"`
	HierarchySelection        []HierarchySelection `json:"hierarchySelection"`
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

type HierarchySelection struct {
	ResourceId    string `json:"resourceId"`
	DisplayName   string `json:"displayName"`
	NodeType      string `json:"nodeType"`
	SelectionType string `json:"selectionType"`
}

type Features struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

type Features1 struct {
	Name  string `json:"featureName"`
	State string `json:"featureState"`
}

//AZUREORG

type CloudAccountAzureResp struct {
	AccountId                 string      `json:"accountId"`
	Name                      string      `json:"name"`
	CloudType                 string      `json:"cloudType"`
	Enabled                   bool        `json:"enabled"`
	ParentId                  string      `json:"parentId"`
	AccountType               string      `json:"accountType"`
	Deleted                   bool        `json:"deleted"`
	ProtectionMode            string      `json:"protectionMode"`
	CustomerName              string      `json:"customerName"`
	CreatedEpochMillis        int         `json:"createdEpochMillis"`
	LastModifiedEpochMillis   int         `json:"lastModifiedEpochMillis"`
	LastModifiedBy            string      `json:"lastModifiedBy"`
	Features                  []Features1 `json:"features"`
	DeploymentType            string      `json:"deploymentType"`
	DeploymentTypeDescription string      `json:"deploymentTypeDescription"`
}
type AzureAccountResponse struct {
	CloudAccountAzureResp CloudAccountAzureResp `json:"cloudAccount"`
	TenantId              string                `json:"tenantId"`
	ServicePrincipalId    string                `json:"servicePrincipalId"`
	ClientId              string                `json:"clientId"`
	TemplateUrl           string                `json:"templateUrl"`
	HierarchySelection    []HierarchySelection  `json:"hierarchySelection"`
	DefaultAccountGroupId string                `json:"defaultAccountGroupId"`
	Key                   string                `json:"key"`
	GroupIds              []string              `json:"groupIds"`
	MonitorFlowLogs       bool                  `json:"monitorFlowLogs"`
	EnvironmentType       string                `json:"environmentType"`
	MemberSyncEnabled     bool                  `json:"memberSyncEnabled"`
}
type AzureOrgV2 struct {
	CloudAccountAzureResp CloudAccountAzureResp `json:"cloudAccount"`
	ClientId              string                `json:"clientId"`
	EnvironmentType       string                `json:"environmentType"`
	Key                   string                `json:"key"`
	MonitorFlowLogs       bool                  `json:"monitorFlowLogs"`
	ServicePrincipalId    string                `json:"servicePrincipalId"`
	TemplateUrl           string                `json:"templateUrl"`
	HierarchySelection    []HierarchySelection  `json:"hierarchySelection"`
	DefaultAccountGroupId string                `json:"defaultAccountGroupId"`
	RootSyncEnabled       bool                  `json:"rootSyncEnabled"`
	GroupIds              []string              `json:"groupIds"`
	MemberSyncEnabled     bool                  `json:"memberSyncEnabled"`
	TenantId              string                `json:"tenantId"`
}

type AzureOrg struct {
	OrgAccountAzure       OrgAccountAzure      `json:"cloudAccount"`
	Enabled               bool                 `json:"enabled"`
	ClientId              string               `json:"clientId"`
	HierarchySelection    []HierarchySelection `json:"hierarchySelection"`
	DefaultAccountGroupId string               `json:"defaultAccountGroupId"`
	Key                   string               `json:"key"`
	MonitorFlowLogs       bool                 `json:"monitorFlowLogs"`
	TenantId              string               `json:"tenantId"`
	ServicePrincipalId    string               `json:"servicePrincipalId"`
	Features              []Features           `json:"features"`
	RootSyncEnabled       bool                 `json:"rootSyncEnabled"`
}

type OrgAccountAzure struct {
	AccountId      string   `json:"accountId"`
	Enabled        bool     `json:"enabled"`
	GroupIds       []string `json:"groupIds"`
	Name           string   `json:"name"`
	ProtectionMode string   `json:"protectionMode"`
	AccountType    string   `json:"accountType"`
}

//GCPORG

type CloudAccountGcpResp struct {
	AccountId                 string      `json:"accountId"`
	Name                      string      `json:"name"`
	CloudType                 string      `json:"cloudType"`
	Enabled                   bool        `json:"enabled"`
	ParentId                  string      `json:"parentId"`
	AccountType               string      `json:"accountType"`
	Deleted                   bool        `json:"deleted"`
	ProtectionMode            string      `json:"protectionMode"`
	CustomerName              string      `json:"customerName"`
	CreatedEpochMillis        int         `json:"createdEpochMillis"`
	LastModifiedEpochMillis   int         `json:"lastModifiedEpochMillis"`
	LastModifiedBy            string      `json:"lastModifiedBy"`
	AddedOnTs                 int         `json:"addedOnTs"`
	DeploymentType            string      `json:"deploymentType"`
	DeploymentTypeDescription string      `json:"deploymentTypeDescription"`
	Features                  []Features1 `json:"features"`
	StorageScanEnabled        bool        `json:"storageScanEnabled"`
}
type GcpAccountResponse struct {
	AccountId                 string `json:"accountId"`
	Name                      string `json:"name"`
	CloudType                 string `json:"cloudType"`
	Enabled                   bool   `json:"enabled"`
	ParentId                  string `json:"parentId"`
	AccountType               string `json:"accountType"`
	Deleted                   bool   `json:"deleted"`
	ProtectionMode            string `json:"protectionMode"`
	CustomerName              string `json:"customerName"`
	CreatedEpochMillis        int    `json:"createdEpochMillis"`
	LastModifiedEpochMillis   int    `json:"lastModifiedEpochMillis"`
	LastModifiedBy            string `json:"lastModifiedBy"`
	AddedOn                   int    `json:"addedOn"`
	DeploymentType            string `json:"deploymentType"`
	DeploymentTypeDescription string `json:"deploymentTypeDescription"`
	StorageScanEnabled        bool   `json:"storageScanEnabled"`
}

type GcpOrgV2 struct {
	CloudAccountGcpResp      CloudAccountGcpResp  `json:"cloudAccount"`
	Credentials              GcpCredentials       `json:"credentials"`
	CompressionEnabled       bool                 `json:"compressionEnabled"`
	DataflowEnabledProject   string               `json:"dataflowEnabledProject"`
	FlowLogStorageBucket     string               `json:"flowLogStorageBucket"`
	ProjectId                string               `json:"projectId"`
	ServiceAccountEmail      string               `json:"serviceAccountEmail"`
	AuthenticationType       string               `json:"authenticationType"`
	GroupIds                 []string             `json:"groupIds"`
	AccountGroupCreationMode string               `json:"accountGroupCreationMode"`
	HierarchySelection       []HierarchySelection `json:"hierarchySelection"`
	DefaultAccountGroupId    string               `json:"defaultAccountGroupId"`
	OrganizationName         string               `json:"organizationName"`
}

type GcpOrg struct {
	OrgAccountGcp            OrgAccountGcp        `json:"cloudAccount"`
	CompressionEnabled       bool                 `json:"compressionEnabled"`
	AccountGroupCreationMode string               `json:"accountGroupCreationMode"`
	DataflowEnabledProject   string               `json:"dataflowEnabledProject"`
	FlowLogStorageBucket     string               `json:"flowLogStorageBucket"`
	Credentials              GcpCredentials       `json:"credentials"`
	Features                 []Features           `json:"features"`
	HierarchySelection       []HierarchySelection `json:"hierarchySelection"`
	DefaultAccountGroupId    string               `json:"defaultAccountGroupId"`
	OrganizationName         string               `json:"organizationName"`
}

type OrgAccountGcp struct {
	AccountId   string `json:"accountId"`
	Enabled     bool   `json:"enabled"`
	ProjectId   string `json:"projectId"`
	Name        string `json:"name"`
	AccountType string `json:"accountType"`
}

type GcpCredentials struct {
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
