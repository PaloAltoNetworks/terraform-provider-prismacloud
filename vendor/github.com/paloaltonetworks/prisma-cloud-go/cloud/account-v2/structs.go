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

type AwsStorageScan struct {
	AccountId         string            `json:"accountId"`
	AccountType       string            `json:"accountType"`
	Enabled           bool              `json:"enabled"`
	Features          []Features        `json:"features"`
	GroupIds          []string          `json:"groupIds"`
	Name              string            `json:"name"`
	RoleArn           string            `json:"roleArn"`
	StorageScanConfig StorageScanConfig `json:"storageScanConfig"`
	StorageUUID       string            `json:"storageUUID"`
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
	CloudAccountResp          CloudAccountResp  `json:"cloudAccount"`
	Name                      string            `json:"name"`
	RoleArn                   string            `json:"roleArn"`
	ExternalId                string            `json:"externalId"`
	HasMemberRole             bool              `json:"hasMemberRole"`
	TemplateUrl               string            `json:"templateUrl"`
	GroupIds                  []string          `json:"groupIds"`
	EventbridgeRuleNamePrefix string            `json:"eventBridgeRuleNamePrefix"`
	StorageScanConfig         StorageScanConfig `json:"storageScanConfig"`
	StorageUUID               string            `json:"storageUUID"`
}

type StorageScanConfig struct {
	Buckets     Buckets `json:"buckets"`
	ScanOption  string  `json:"scanOption"`
	SnsTopicArn string  `json:"snsTopicArn"`
}

type Buckets struct {
	Backward []string `json:"backward"`
	Forward  []string `json:"forward"`
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

//GCPV2

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
	Features                  []Features1 `json:"features"`
	AddedOnTs                 int         `json:"addedOnTs"`
	DeploymentType            string      `json:"deploymentType"`
	DeploymentTypeDescription string      `json:"deploymentTypeDescription"`
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
type GcpV2 struct {
	CloudAccountGcpResp      CloudAccountGcpResp `json:"cloudAccount"`
	Credentials              GcpCredentials      `json:"credentials"`
	CompressionEnabled       bool                `json:"compressionEnabled"`
	GroupIds                 []string            `json:"groupIds"`
	DataflowEnabledProject   string              `json:"dataflowEnabledProject"`
	FlowLogStorageBucket     string              `json:"flowLogStorageBucket"`
	ProjectId                string              `json:"projectId"`
	ServiceAccountEmail      string              `json:"serviceAccountEmail"`
	DefaultAccountGroupId    string              `json:"defaultAccountGroupId"`
	AuthenticationType       string              `json:"authenticationType"`
	AccountGroupCreationMode string              `json:"accountGroupCreationMode"`
}

type Gcp struct {
	CloudAccountGcp        CloudAccountGcp `json:"cloudAccount"`
	CompressionEnabled     bool            `json:"compressionEnabled"`
	DataflowEnabledProject string          `json:"dataflowEnabledProject"`
	FlowLogStorageBucket   string          `json:"flowLogStorageBucket"`
	DefaultAccountGroupId  string          `json:"defaultAccountGroupId"`
	Credentials            GcpCredentials  `json:"credentials"`
	Features               []Features      `json:"features"`
}

type CloudAccountGcp struct {
	AccountId   string   `json:"accountId"`
	Enabled     bool     `json:"enabled"`
	GroupIds    []string `json:"groupIds"`
	Name        string   `json:"name"`
	AccountType string   `json:"accountType"`
	ProjectId   string   `json:"projectId"`
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

// IBM V2
type Ibm struct {
	AccountId   string   `json:"accountId"`
	AccountType string   `json:"accountType"`
	ApiKey      string   `json:"apiKey"`
	Enabled     bool     `json:"enabled"`
	GroupIds    []string `json:"groupIds"`
	Name        string   `json:"name"`
	SvcIdIamId  string   `json:"svcIdIamId"`
}

type IbmV2 struct {
	CloudAccountIbmResp CloudAccountIbmResp `json:"cloudAccount"`
	GroupIds            []string            `json:"groupIds"`
	SvcIdIamId          string              `json:"svcIdIamId"`
	ApiKey              string              `json:"apiKey"`
}

type CloudAccountIbmResp struct {
	AccountId                 string      `json:"accountId"`
	AccountType               string      `json:"accountType"`
	AddedOnTs                 int         `json:"addedOnTs"`
	CloudType                 string      `json:"cloudType"`
	CreatedEpochMillis        int         `json:"createdEpochMillis"`
	CustomerName              string      `json:"customerName"`
	Deleted                   bool        `json:"deleted"`
	DeploymentType            string      `json:"deploymentType"`
	DeploymentTypeDescription string      `json:"deploymentTypeDescription"`
	Enabled                   bool        `json:"enabled"`
	LastModifiedEpochMillis   int         `json:"lastModifiedEpochMillis"`
	LastModifiedBy            string      `json:"lastModifiedBy"`
	Name                      string      `json:"name"`
	ParentId                  string      `json:"parentId"`
	ProtectionMode            string      `json:"protectionMode"`
	StorageScanEnabled        bool        `json:"storageScanEnabled"`
	Features                  []Features1 `json:"features"`
}

type IbmAccountResponse struct {
	AccountId                 string      `json:"accountId"`
	AccountType               string      `json:"accountType"`
	AddedOnTs                 int         `json:"addedOnTs"`
	CloudType                 string      `json:"cloudType"`
	CreatedEpochMillis        int         `json:"createdEpochMillis"`
	CustomerName              string      `json:"customerName"`
	Deleted                   bool        `json:"deleted"`
	DeploymentType            string      `json:"deploymentType"`
	DeploymentTypeDescription string      `json:"deploymentTypeDescription"`
	Enabled                   bool        `json:"enabled"`
	LastModifiedEpochMillis   int         `json:"lastModifiedEpochMillis"`
	LastModifiedBy            string      `json:"lastModifiedBy"`
	Name                      string      `json:"name"`
	ParentId                  string      `json:"parentId"`
	ProtectionMode            string      `json:"protectionMode"`
	StorageScanEnabled        bool        `json:"storageScanEnabled"`
	Features                  []Features1 `json:"features"`
}

// Alibaba
type Alibaba struct {
	AccountId      string   `json:"accountId"`
	DeploymentType string   `json:"deploymentType"`
	Enabled        bool     `json:"enabled"`
	GroupIds       []string `json:"groupIds"`
	Name           string   `json:"name"`
	RamArn         string   `json:"ramArn"`
}

type AlibabaV2 struct {
	Name               string             `json:"name"`
	CloudAccountStatus CloudAccountStatus `json:"cloudAccountStatus"`
	AccountType        string             `json:"accountType"`
	AddedOn            int                `json:"addedOn"`
	Enabled            bool               `json:"enabled"`
	LastModifiedTs     int                `json:"lastModifiedTs"`
	LastModifiedBy     string             `json:"lastModifiedBy"`
	ProtectionMode     string             `json:"protectionMode"`
	StorageScanEnabled bool               `json:"storageScanEnabled"`
	DeploymentType     string             `json:"deploymentType"`
	GroupIds           []string           `json:"groupIds"`
	RamArn             string             `json:"ramArn"`
}

type CloudAccountStatus struct {
	AccountId        string `json:"accountId"`
	LastUpdated      int    `json:"lastUpdated"`
	LastFullSnapshot int    `json:"lastFullSnapshot"`
	IngestionEndTime int    `json:"ingestionEndTime"`
	CloudType        string `json:"cloudType"`
}
type AlibabaAccountResponse struct {
	Name                  string   `json:"name"`
	AccountId             string   `json:"accountId"`
	CloudType             string   `json:"cloudType"`
	AccountType           string   `json:"accountType"`
	AddedOn               int      `json:"addedOn"`
	Enabled               bool     `json:"enabled"`
	LastModifiedTs        int      `json:"lastModifiedTs"`
	LastModifiedBy        string   `json:"lastModifiedBy"`
	ProtectionMode        string   `json:"protectionMode"`
	StorageScanEnabled    bool     `json:"storageScanEnabled"`
	DeploymentType        string   `json:"deploymentType"`
	Groups                []Groups `json:"groups"`
	GroupIds              []string `json:"groupIds"`
	NumberOfChildAccounts int      `json:"numberOfChildAccounts"`
}
type Groups struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
