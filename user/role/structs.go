package role

type Role struct {
	Id                      string         `json:"id,omitempty"`
	Name                    string         `json:"name"`
	Description             string         `json:"description,omitempty"`
	RoleType                string         `json:"roleType"`
	LastModifiedBy          string         `json:"lastModifiedBy,omitempty"`
	LastModifiedTs          int64          `json:"lastModifiedTs,omitempty"`
	AccountGroupIds         []string       `json:"accountGroupIds"`
	AssociatedUsers         []string       `json:"associatedUsers"`
	RestrictDismissalAccess bool           `json:"restrictDismissalAccess,omitempty"`
	AccountGroups           []AccountGroup `json:"accountGroups,omitempty"`
}

type AccountGroup struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type NameId struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
