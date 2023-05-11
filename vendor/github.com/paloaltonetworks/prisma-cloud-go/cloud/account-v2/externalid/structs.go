package externalid

type ExternalIdReq struct {
	AccountId    string   `json:"accountId"`
	AccountType  string   `json:"accountType"`
	AwsPartition string   `json:"awsPartition"`
	Features     []string `json:"features"`
}

type ExternalId struct {
	AccountId                         string    `json:"accountId"`
	AccountType                       string    `json:"accountType"`
	AwsPartition                      string    `json:"awsPartition"`
	Features                          Features1 `json:"features"`
	ExternalId                        string    `json:"externalId"`
	CreateStackLinkWithS3PresignedUrl string    `json:"createStackLinkWithS3PresignedUrl"`
	EventBridgeRuleNamePrefix         string    `json:"eventBridgeRuleNamePrefix"`
}

type Features1 struct {
	Name  string `json:"featureName"`
	State string `json:"featureState"`
}

type StorageUUID struct {
	AccountId   string `json:"accountId"`
	ExternalId  string `json:"externalId"`
	RoleArn     string `json:"roleArn"`
	StorageUUID string `json:"storageUUID"`
}
