package section

type Section struct {
	Id                    string   `json:"id,omitempty"`
	Description           string   `json:"description,omitempty"`
	CreatedBy             string   `json:"createdBy,omitempty"`
	CreatedOn             int      `json:"createdOn,omitempty"`
	LastModifiedBy        string   `json:"lastModifiedBy,omitempty"`
	LastModifiedOn        int      `json:"lastModifiedOn,omitempty"`
	SystemDefault         bool     `json:"systemDefault,omitempty"`
	PoliciesAssignedCount int      `json:"policiesAssignedCount,omitempty"`
	StandardName          string   `json:"standardName"`
	RequirementId         string   `json:"-"`
	RequirementName       string   `json:"requirementName,omitempty"`
	SectionId             string   `json:"sectionId,omitempty"`
	Label                 string   `json:"label,omitempty"`
	ViewOrder             int      `json:"viewOrder,omitempty"`
	AssociatedPolicyIds   []string `json:"associatedPolicyIds,omitempty"`
}
