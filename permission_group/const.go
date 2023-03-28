package permission_group

var Suffix = []string{"authz", "v1", "permission_group"}

const (
	singular = "permission group"
	plural   = "permission groups"
)

// Valid values for permission_group.PermissionGroup.
const (
	TypeDefault  = "Default"
	TypeCustom   = "Custom"
	TypeInternal = "Internal"
)
