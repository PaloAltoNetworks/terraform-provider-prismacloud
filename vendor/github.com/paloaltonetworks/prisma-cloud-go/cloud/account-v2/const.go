package accountv2

const (
	singular = "cloud account"
	plural   = "cloud accounts"
)

var DeleteSuffix = []string{"cloud"}

var Suffix = []string{"cas", "v1"}

var ListSuffix = []string{"v1", "cloudAccounts"}

var ListSuffixAws = []string{"v1", "cloudAccounts", "awsAccounts"}

var ListSuffixAzure = []string{"v1", "cloudAccounts", "azureAccounts"}

const (
	TypeAws   = "aws"
	TypeAzure = "azure"
)
