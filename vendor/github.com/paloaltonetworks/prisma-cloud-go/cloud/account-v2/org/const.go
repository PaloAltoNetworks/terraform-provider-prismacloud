package org

const (
	singular = "org cloud account"
	plural   = "org cloud accounts"
)

var DeleteSuffix = []string{"cloud"}

var Suffix = []string{"cas", "v1", "aws_account"}

var ListSuffix = []string{"v1", "cloudAccounts", "awsAccounts"}

const (
	TypeAwsOrg   = "aws"
	TypeAzureOrg = "azure"
	TypeGcpOrg   = "gcp"
	TypeOci      = "oci"
)
