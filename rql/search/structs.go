package search

import (
	"github.com/paloaltonetworks/prisma-cloud-go/timerange"
)

type ConfigRequest struct {
	Id               string              `json:"id,omitempty"`
	TimeRange        timerange.TimeRange `json:"timeRange"`
	Query            string              `json:"query"`
	Limit            int                 `json:"limit,omitempty"`
	WithResourceJson bool                `json:"withResourceJson,omitempty"`
}

type ConfigResponse struct {
	GroupBy     []string   `json:"groupBy"`
	Id          string     `json:"id"`
	CloudType   string     `json:"cloudType"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	SearchType  string     `json:"searchType"`
	Data        ConfigData `json:"data"`
}

type ConfigData struct {
	Items []ConfigItem `json:"items"`
}

type ConfigItem struct {
	StateId                  string `json:"stateId"`
	Name                     string `json:"name"`
	Url                      string `json:"url"`
	AccountId                string `json:"accountId"`
	AccountName              string `json:"accountName"`
	AccountGroupName         string `json:"accountGroupName"`
	CloudType                string `json:"cloudType"`
	RegionId                 string `json:"regionId"`
	RegionName               string `json:"regionName"`
	Service                  string `json:"service"`
	ResourceType             string `json:"resourceType"`
	InsertTs                 int    `json:"insertTs"`
	Deleted                  bool   `json:"deleted"`
	VpcId                    string `json:"vpcId"`
	VpnName                  string `json:"vpcName"`
	RiskGrade                string `json:"riskGrade"`
	HasNetwork               bool   `json:"hasNetwork"`
	HasAlert                 bool   `json:"hasAlert"`
	HasExternalFinding       bool   `json:"hasExternalFinding"`
	HasExternalIntegration   bool   `json:"hasExternalIntegration"`
	AllowDrillDown           bool   `json:"allowDrillDown"`
	HasExtFindingRiskFactors bool   `json:"hasExtFindingRiskFactors"`
}

type EventRequest struct {
	Id               string              `json:"id,omitempty"`
	TimeRange        timerange.TimeRange `json:"timeRange"`
	Query            string              `json:"query"`
	Limit            int                 `json:"limit,omitempty"`
	WithResourceJson bool                `json:"withResourceJson,omitempty"`
}

type EventResponse struct {
	GroupBy     []string  `json:"groupBy"`
	Id          string    `json:"id"`
	CloudType   string    `json:"cloudType"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SearchType  string    `json:"searchType"`
	Data        EventData `json:"data"`
}

type EventData struct {
	Items []EventItem `json:"items"`
}

type EventItem struct {
	Account             string `json:"account"`
	RegionId            int    `json:"regionId"`
	RegionApiIdentifier string `json:"regionApiIdentifier"`
	EventTs             int    `json:"eventTs"`
	IngestionTs         int    `json:"ingestionTs"`
	Subject             string `json:"subject"`
	Type                string `json:"type"`
	Source              string `json:"source"`
	Name                string `json:"name"`
	Id                  int    `json:"id"`
	Ip                  string `json:"ip"`
	AccessKey           string `json:"accessKey"`
	AnomalyId           string `json:"anomalyId"`
	AccessKeyUsed       bool   `json:"accessKeyUsed"`
	SubjectType         string `json:"subjectType"`
	Role                string `json:"role"`
	ReasonIds           []int  `json:"reasonIds"`
	FlaggedFeature      string `json:"flaggedFeature"`
	CityId              int    `json:"cityId"`
	CityName            string `json:"cityName"`
	StateId             int    `json:"stateId"`
	StateName           string `json:"stateName"`
	CountryId           int    `json:"countryId"`
	Success             bool   `json:"success"`
	Internal            bool   `json:"internal"`
	Location            string `json:"location"`
	Os                  string `json:"os"`
	Browser             string `json:"browser"`
	AccountName         string `json:"accountName"`
	RegionName          string `json:"regionName"`
}

type NetworkRequest struct {
	Id               string              `json:"id,omitempty"`
	TimeRange        timerange.TimeRange `json:"timeRange"`
	Query            string              `json:"query"`
	Limit            int                 `json:"limit,omitempty"`
	WithResourceJson bool                `json:"withResourceJson,omitempty"`
}

type NetworkResponse struct {
	GroupBy     []string    `json:"groupBy"`
	Id          string      `json:"id"`
	CloudType   string      `json:"cloudType"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	SearchType  string      `json:"searchType"`
	Data        NetworkData `json:"data"`
}

type NetworkData struct {
	Items []NetworkItem `json:"items"`
}

type NetworkItem struct {
	Account     string `json:"account"`
	RegionId    string `json:"regionId"`
	AccountName string `json:"accountName"`
	Ip          string `json:"ip"`
	Name        string `json:"name"`
	StateId     int    `json:"stateId"`
	StateName   string `json:"stateName"`
	Source      string `json:"source"`
}
