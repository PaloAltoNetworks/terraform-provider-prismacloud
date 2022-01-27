package role

const (
	singular = "user role"
	plural   = "user roles"
)

// Valid values for Role.RoleType.
const (
	TypeSystemAdmin                      = "System Admin"
	TypeAccountGroupAdmin                = "Account Group Admin"
	TypeAccountGroupReadOnly             = "Account Group Read Only"
	TypeCloudProvisioningAdmin           = "Cloud Provisioning Admin"
	TypeAccountAndCloudProvisioningAdmin = "Account and Cloud Provisioning Admin"
	TypeBuildAndDeploySecurity           = "Build and Deploy Security"
)

var Suffix = []string{"user", "role"}
