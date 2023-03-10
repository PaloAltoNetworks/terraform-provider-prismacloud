package supportedFeatures

type SupportedFeaturesReq struct {
	CloudType       string `json:"cloudType"`
	AccountType     string `json:"accountType"`
	DeploymentType  string `json:"deploymentType"`
	AwsPartition    string `json:"awsPartition"`
	RootSyncEnabled bool   `json:"rootSyncEnabled"`
}

type SupportedFeatures struct {
	AccountType       string   `json:"accountType"`
	DeploymentType    string   `json:"deploymentType"`
	AwsPartition      string   `json:"awsPartition"`
	RootSyncEnabled   bool     `json:"rootSyncEnabled"`
	CloudType         string   `json:"cloudType"`
	LicenseType       string   `json:"licenseType"`
	SupportedFeatures []string `json:"supportedFeatures"`
}
