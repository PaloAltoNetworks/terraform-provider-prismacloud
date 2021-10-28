package policy

type Policy struct {
	PolicyId               string               `json:"policyId,omitempty"`
	Name                   string               `json:"name"`
	PolicyType             string               `json:"policyType"`
	PolicySubTypes         []string             `json:"policySubTypes,omitempty"`
	SystemDefault          bool                 `json:"systemDefault,omitempty"`
	PolicyUpi              string               `json:"policyUpi,omitempty"`
	Description            string               `json:"description,omitempty"`
	Severity               string               `json:"severity"`
	Rule                   Rule                 `json:"rule"`
	Recommendation         string               `json:"recommendation,omitempty"`
	CloudType              string               `json:"cloudType,omitempty"`
	ComplianceMetadata     []ComplianceMetadata `json:"complianceMetadata"`
	Remediation            Remediation          `json:"remediation,omitempty"`
	Labels                 []string             `json:"labels,omitempty"` // unordered
	Enabled                bool                 `json:"enabled"`
	CreatedOn              int                  `json:"createdOn,omitempty"`
	CreatedBy              string               `json:"createdBy,omitempty"`
	LastModifiedOn         int                  `json:"lastModifiedOn,omitempty"`
	LastModifiedBy         string               `json:"lastModifiedBy,omitempty"`
	RuleLastModifiedOn     int                  `json:"ruleLastModifiedOn,omitempty"`
	Overridden             bool                 `json:"overridden,omitempty"`
	Deleted                bool                 `json:"deleted,omitempty"`
	RestrictAlertDismissal bool                 `json:"restrictAlertDismissal,omitempty"`
	OpenAlertsCount        int                  `json:"openAlertsCount,omitempty"`
	Owner                  string               `json:"owner,omitempty"`
	PolicyMode             string               `json:"policyMode,omitempty"`
	PolicyCategory         string               `json:"policyCategory,omitempty"`
	PolicyClass            string               `json:"policyClass,omitempty"`
	Remediable             bool                 `json:"remediable,omitempty"`
}

/*
Rule is the rule object.

Due to 05befc8b-c78a-45e9-98dc-c7fbaef580e7, criteria has to be
an interface{}.
*/
type Rule struct {
	Name           string            `json:"name"`
	CloudType      string            `json:"cloudType,omitempty"`
	CloudAccount   string            `json:"cloudAccount,omitempty"`
	ResourceType   string            `json:"resourceType,omitempty"`
	ApiName        string            `json:"apiName,omitempty"`
	ResourceIdPath string            `json:"resourceIdPath,omitempty"`
	Criteria       interface{}       `json:"criteria,omitempty"`
	DataCriteria   DataCriteria      `json:"dataCriteria,omitempty"`
	Parameters     map[string]string `json:"parameters"`
	Type           string            `json:"type"`
}

type ComplianceMetadata struct {
	StandardName           string `json:"standardName,omitempty"`
	StandardDescription    string `json:"standardDescription,omitempty"`
	RequirementId          string `json:"requirementId,omitempty"`
	RequirementName        string `json:"requirementName,omitempty"`
	RequirementDescription string `json:"requirementDescription,omitempty"`
	SectionId              string `json:"sectionId,omitempty"`
	SectionDescription     string `json:"sectionDescription,omitempty"`
	PolicyId               string `json:"policyId,omitempty"`
	ComplianceId           string `json:"complianceId,omitempty"`
	SectionLabel           string `json:"sectionLabel,omitempty"`
	CustomAssigned         bool   `json:"customAssigned,omitempty"`
}

type Remediation struct {
	TemplateType        string      `json:"templateType,omitempty"`
	Description         string      `json:"description,omitempty"`
	CliScriptTemplate   string      `json:"cliScriptTemplate,omitempty"`
	CliScriptJsonSchema interface{} `json:"cliScriptJsonSchema,omitempty"`
}

type DataCriteria struct {
	ClassificationResult string   `json:"classificationResult,omitempty"`
	Exposure             string   `json:"exposure,omitempty"`
	Extension            []string `json:"extension,omitempty"`
}
