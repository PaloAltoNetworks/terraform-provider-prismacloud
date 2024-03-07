package group

type NameId struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Group struct {
	Id             string      `json:"id,omitempty"`
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	LastModifiedBy string      `json:"lastModifiedBy"`
	LastModifiedTs int64       `json:"lastModifiedTs"`
	AccountIds     []string    `json:"accountIds"`
	Accounts       []Account   `json:"accounts,omitempty"`
	AlertRules     []AlertRule `json:"alertRules,omitempty"`
	ChildGroupIds  []string    `json:"childGroupIds"`
	ParentInfo     ParentInfo  `json:"parentInfo,omitempty"`
}

type Account struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type AlertRule struct {
	Id   string `json:"alertId"`
	Name string `json:"alertName"`
}

type ParentInfo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	AutoCreated bool   `json:"autoCreated"`
}
