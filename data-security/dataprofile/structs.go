package dataprofile

type Profile struct {
	Id                   string               `json:"id,omitempty"`
	Name                 string               `json:"name"`
	Description          string               `json:"description"`
	Types                string               `json:"types"`
	ProfileType          string               `json:"profileType"`
	TenantId             string               `json:"tenantId,omitempty"`
	Status               string               `json:"status"`
	ProfileStatus        string               `json:"profileStatus"`
	DataPatternsRulesOne DataPatternsRulesOne `json:"dataPatternsRulesOne,omitempty"`
	DataPatternsRule1    DataPatternsRule1    `json:"dataPatternsRule1"`
	CreatedBy            string               `json:"createdBy,omitempty"`
	UpdatedBy            string               `json:"updatedBy,omitempty"`
	CreatedAt            string               `json:"createdAt,omitempty"`
	UpdatedAt            string               `json:"updatedAt,omitempty"`
}

type DataPatternsRulesOne struct {
	OperatorType     string               `json:"operatorType,omitempty"`
	DataPatternRules []DataPatternRuleOne `json:"dataPatternRules,omitempty"`
}

type DataPatternsRule1 struct {
	OperatorType     string             `json:"operatorType"`
	DataPatternRules []DataPatternRule1 `json:"dataPatternRules"`
}

type DataPatternRule1 struct {
	Id                        string   `json:"id,omitempty"`
	Name                      string   `json:"name"`
	DetectionTechnique        string   `json:"detectionTechnique"`
	MatchType                 string   `json:"matchType"`
	OccurrenceOperatorType    string   `json:"occurrenceOperatorType"`
	OccurrenceCount           int      `json:"occurrenceCount,omitempty"`
	ConfidenceLevel           string   `json:"confidenceLevel"`
	SupportedConfidenceLevels []string `json:"supportedConfidenceLevels,omitempty"`
	OccurrenceHigh            int      `json:"occurrenceHigh,omitempty"`
	OccurrenceLow             int      `json:"occurrenceLow,omitempty"`
}

type DataPatternRuleOne struct {
	Id                        string   `json:"id"`
	Name                      string   `json:"name"`
	DetectionTechnique        string   `json:"detectionTechnique"`
	MatchType                 string   `json:"matchType"`
	OccurrenceOperatorType    string   `json:"occurrenceOperatorType"`
	OccurrenceCount           int      `json:"occurrenceCount"`
	ConfidencLevel            string   `json:"confidencLevel"`
	SupportedConfidenceLevels []string `json:"supportedConfidenceLevels"`
	OccurrenceHigh            int      `json:"occurrenceHigh"`
	OccurrenceLow             int      `json:"occurrenceLow"`
}

type ListBody struct {
	Profiles []Profile `json:"profiles"`
}
