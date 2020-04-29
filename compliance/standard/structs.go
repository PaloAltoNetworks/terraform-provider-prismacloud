package standard

type Standard struct {
	Name                  string   `json:"name"`
	Id                    string   `json:"id,omitempty"`
	Description           string   `json:"description,omitempty"`
	CreatedBy             string   `json:"createdBy,omitempty"`
	CreatedOn             int      `json:"createdOn,omitempty"`
	LastModifiedBy        string   `json:"lastModifiedBy,omitempty"`
	LastModifiedOn        int      `json:"lastModifiedOn,omitempty"`
	SystemDefault         bool     `json:"systemDefault,omitempty"`
	PoliciesAssignedCount int      `json:"policiesAssignedCount,omitempty"`
	CloudTypes            []string `json:"cloudType,omitempty"`
}
