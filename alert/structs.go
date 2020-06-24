package alert

import (
	"github.com/paloaltonetworks/prisma-cloud-go/timerange"
)

type Request struct {
	TimeRange timerange.TimeRange `json:"timeRange"`
	Limit     int                 `json:"limit,omitempty"`
	Offset    int                 `json:"offset,omitempty"`
	Detailed  bool                `json:"detailed"`
	PageToken string              `json:"pageToken,omitempty"`
	SortBy    []string            `json:"sortBy,omitempty"`
	Filters   []Filter            `json:"filters,omitempty"`
}

type Filter struct {
	Name     string `json:"name"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type Response struct {
	Total     int     `json:"totalRows"`
	Data      []Alert `json:"items"`
	PageToken string  `json:"nextPageToken"`
}

type Alert struct {
	Id                 string             `json:"id"`
	Status             string             `json:"status"`
	FirstSeen          int                `json:"firstSeen"`
	LastSeen           int                `json:"lastSeen"`
	AlertTime          int                `json:"alertTime"`
	EventOccurred      int                `json:"eventOccurred"`
	TriggeredBy        string             `json:"triggeredBy"`
	AlertCount         int                `json:"alertCount"`
	History            []History          `json:"history"`
	Policy             Policy             `json:"policy"`
	Risk               RiskDetail         `json:"riskDetail"`
	Resource           Resource           `json:"resource"`
	InvestigateOptions InvestigateOptions `json:"investigateOptions"`
}

type History struct {
	Reason     string `json:"reason"`
	Status     string `json:"status"`
	ModifiedBy string `json:"modifiedBy"`
	ModifiedOn int    `json:"modifiedOn"`
}

type Policy struct {
	Id            string `json:"policyId"`
	Type          string `json:"policyType"`
	SystemDefault bool   `json:"systemDefault"`
	Remediable    bool   `json:"remediable"`
}

type RiskDetail struct {
	RiskScore RiskScore `json:"riskScore"`
	Rating    string    `json:"rating"`
	Score     string    `json:"score"`
}

type RiskScore struct {
	Score    int `json:"score"`
	MaxScore int `json:"maxScore"`
}

type Resource struct {
	Rrn                string      `json:"rrn"`
	Id                 string      `json:"id"`
	Name               string      `json:"name"`
	Account            string      `json:"account"`
	AccountId          string      `json:"accountId"`
	CloudAccountGroups []string    `json:"cloudAccountGroups"`
	Region             string      `json:"region"`
	RegionId           string      `json:"regionId"`
	ResourceType       string      `json:"resourceType"`
	ResourceApiName    string      `json:"resourceApiName"`
	Url                string      `json:"url"`
	Data               interface{} `json:"data"`
	Tags               interface{} `json:"resourceTags,omitempty"`
	AlertAttribution   interface{} `json:"alertAttribution"`
	CloudType          string      `json:"cloudType"`
}

type InvestigateOptions struct {
	SearchId string `json:"searchId"`
	StartTs  int    `json:"startTs"`
	EndTs    int    `json:"endTs"`
}
