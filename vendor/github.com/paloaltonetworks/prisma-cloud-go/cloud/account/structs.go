package account

type AccountAndCredentials struct {
	Account Account `json:"cloudAccount"`
}

type NameTypeId struct {
	Name      string `json:"name"`
	CloudType string `json:"cloudType"`
	AccountId string `json:"id"`
}

type Account struct {
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

type Aws struct {
	AccountId      string   `json:"accountId"`
	Enabled        bool     `json:"enabled"`
	ExternalId     string   `json:"externalId"`
	GroupIds       []string `json:"groupIds"`
	Name           string   `json:"name"`
	RoleArn        string   `json:"roleArn"`
	ProtectionMode string   `json:"protectionMode"`
	AccountType    string   `json:"accountType"`
}

type CloudAccount struct {
	AccountId      string   `json:"accountId"`
	Enabled        bool     `json:"enabled"`
	GroupIds       []string `json:"groupIds"`
	Name           string   `json:"name"`
	ProtectionMode string   `json:"protectionMode"`
	AccountType    string   `json:"accountType"`
}

type Azure struct {
	Account            CloudAccount `json:"cloudAccount"`
	ClientId           string       `json:"clientId"`
	Key                string       `json:"key"`
	MonitorFlowLogs    bool         `json:"monitorFlowLogs"`
	TenantId           string       `json:"tenantId"`
	ServicePrincipalId string       `json:"servicePrincipalId"`
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

type Gcp struct {
	Account                CloudAccount   `json:"cloudAccount"`
	CompressionEnabled     bool           `json:"compressionEnabled"`
	DataflowEnabledProject string         `json:"dataflowEnabledProject"`
	FlowLogStorageBucket   string         `json:"flowLogStorageBucket"`
	Credentials            GcpCredentials `json:"credentials"`
}

type Alibaba struct {
	AccountId string   `json:"accountId"`
	GroupIds  []string `json:"groupIds"`
	Name      string   `json:"name"`
	RamArn    string   `json:"ramArn"`
	Enabled   bool     `json:"enabled"`
}
