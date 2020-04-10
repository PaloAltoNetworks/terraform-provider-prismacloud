package policy

type Policy struct {
	PolicyId               string               `json:"policyId,omitempty"`
	Name                   string               `json:"name"`
	PolicyType             string               `json:"policyType"`
	SystemDefault          bool                 `json:"systemDefault,omitempty"`
	Description            string               `json:"description"`
	Severity               string               `json:"severity"`
	Rule                   Rule                 `json:"rule"`
	Recommendation         string               `json:"recommendation"`
	CloudType              string               `json:"cloudType"`
	ComplianceMetadata     []ComplianceMetadata `json:"complianceMetadata"`
	Remediation            Remediation          `json:"remediation"`
	Labels                 []string             `json:"labels"`
	Enabled                bool                 `json:"enabled"`
	CreatedOn              int                  `json:"createdOn,omitempty"`
	CreatedBy              string               `json:"createdBy,omitempty"`
	LastModifiedOn         int                  `json:"lastModifiedOn,omitempty"`
	LastModifiedBy         string               `json:"lastModifiedBy,omitempty"`
	RuleLastModifiedOn     int                  `json:"ruleLastModifiedOn,omitempty"`
	Overridden             bool                 `json:"overridden"`
	Deleted                bool                 `json:"deleted"`
	RestrictAlertDismissal bool                 `json:"restrictAlertDismissal"`
	OpenAlertsCount        int                  `json:"openAlertsCount,omitempty"`
	Owner                  string               `json:"owner"`
	PolicyMode             string               `json:"policyMode"`
	Remediable             bool                 `json:"remediable,omitempty"`
}

type Rule struct {
	Name           string            `json:"name"`
	CloudType      string            `json:"cloudType"`
	CloudAccount   string            `json:"cloudAccount"`
	ResourceType   string            `json:"resourceType"`
	ApiName        string            `json:"apiName"`
	ResourceIdPath string            `json:"resourceIdPath"`
	Criteria       string            `json:"criteria"`
	Parameters     map[string]string `json:"parameters"`
	Type           string            `json:"type"`
}

type ComplianceMetadata struct {
	StandardName           string `json:"standardName"`
	StandardDescription    string `json:"standardDescription"`
	RequirementId          string `json:"requirementId"`
	RequirementName        string `json:"requirementName"`
	RequirementDescription string `json:"requirementDescription"`
	SectionId              string `json:"sectionId"`
	SectionDescription     string `json:"sectionDescription"`
	PolicyId               string `json:"policyId"`
	ComplianceId           string `json:"complianceId"`
	SectionLabel           string `json:"sectionLabel"`
	CustomAssigned         bool   `json:"customAssigned"`
}

type Remediation struct {
	TemplateType        string      `json:"templateType"`
	CliScriptTemplate   string      `json:"cliScriptTemplate"`
	CliScriptJsonSchema interface{} `json:"cliScriptJsonSchema"`
}
