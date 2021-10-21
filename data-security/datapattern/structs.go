package datapattern

type Pattern struct {
	Id                 string      `json:"id,omitempty"`
	Name               string      `json:"name"`
	Description        string      `json:"description"`
	Mode               string      `json:"mode,omitempty"`
	DetectionTechnique string      `json:"detectionTechnique"`
	Entity             string      `json:"entity,omitempty"`
	Grammar            string      `json:"grammar,omitempty"`
	ParentId           string      `json:"parentId,omitempty"`
	ProximityKeywords  []string    `json:"proximityKeywords,omitempty"`
	Regexes            []RegexInfo `json:"regexes"`
	RootType           string      `json:"rootType,omitempty"`
	S3Path             string      `json:"s3Path,omitempty"`
	CreatedBy          string      `json:"createdBy,omitempty"`
	UpdatedBy          string      `json:"updatedBy,omitempty"`
	UpdatedAt          int         `json:"updatedAt,omitempty"`
	IsEditable         bool        `json:"isEditable,omitempty"`
}

type RegexInfo struct {
	Regex  string `json:"regex"`
	Weight int    `json:"weight,omitempty"`
}

type ListBody struct {
	ModeFilter          []string  `json:"modeFilter"`
	LastUpdatedByFilter []string  `json:"lastUpdatedByFilter"`
	Patterns            []Pattern `json:"patterns"`
}
