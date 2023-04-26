package org

const (
	singular = "org cloud account"
	plural   = "org cloud accounts"
)

var DeleteSuffix = []string{"cloud"}

var Suffix = []string{"cas", "v1"}

var ListSuffix = []string{"v1", "cloudAccounts"}

var ListSuffixAws = []string{"v1", "cloudAccounts", "awsAccounts"}

var ListSuffixAzure = []string{"v1", "cloudAccounts", "azureAccounts"}

var ListSuffixGcp = []string{"cas", "v1", "cloud_account", "gcp"}

const (
	TypeAwsOrg   = "aws"
	TypeAzureOrg = "azure"
	TypeGcpOrg   = "gcp"
)
