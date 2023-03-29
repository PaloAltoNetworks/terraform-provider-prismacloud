package permission_group

type PermissionGroup struct {
	Id                     string            `json:"id,omitempty"`
	Name                   string            `json:"name"`
	Description            string            `json:"description,omitempty"`
	Type                   string            `json:"type"`
	LastModifiedBy         string            `json:"lastModifiedBy,omitempty"`
	LastModifiedTs         int64             `json:"lastModifiedTs,omitempty"`
	AssociatedRoles        map[string]string `json:"associatedRoles"`
	Features               []Features        `json:"features"`
	AcceptAccountGroups    bool              `json:"acceptAccountGroups"`
	AcceptResourceLists    bool              `json:"acceptResourceLists"`
	AcceptCodeRepositories bool              `json:"acceptCodeRepositories"`
	Custom                 bool              `json:"custom"`
	Deleted                bool              `json:"deleted,omitempty"`
}

type Features struct {
	Operations  Operations `json:"operations"`
	FeatureName string     `json:"featureName"`
}

type NameId struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
type Operations struct {
	CREATE bool `json:"CREATE"`
	READ   bool `json:"READ"`
	UPDATE bool `json:"UPDATE"`
	DELETE bool `json:"DELETE"`
}
