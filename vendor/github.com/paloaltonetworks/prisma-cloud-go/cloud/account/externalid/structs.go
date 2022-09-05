package externalid

type ExternalIdReq struct {
	AccountId          string `json:"accountId"`
	Name               string `json:"name"`
	StorageScanEnabled bool   `json:"storageScanEnabled"`
	ProtectionMode     string `json:"protectionMode"`
	AwsPartition       string `json:"awsPartition"`
}

type ExternalId struct {
	AccountId          string `json:"accountId"`
	Name               string `json:"name"`
	ExternalId         string `json:"externalId"`
	ProtectionMode     string `json:"protectionMode"`
	AwsPartition       string `json:"awsPartition"`
	CftPath            string `json:"cftPath"`
	CloudFormationUrl  string `json:"cloudFormationUrl"`
	StorageScanEnabled bool   `json:"storageScanEnabled"`
}
