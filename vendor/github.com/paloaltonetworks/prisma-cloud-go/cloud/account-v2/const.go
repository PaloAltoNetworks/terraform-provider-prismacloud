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

var ListSuffixGcp = []string{"cas", "v1", "cloud_account", "gcp"}

var ListSuffixIbm = []string{"cas", "v1", "cloud_account", "ibm"}

const (
	TypeAws   = "aws"
	TypeAzure = "azure"
	TypeGcp   = "gcp"
	TypeIbm   = "ibm"
)
