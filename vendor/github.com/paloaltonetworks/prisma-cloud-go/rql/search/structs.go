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

type IamRequest struct {
	Id    string `json:"id,omitempty"`
	Query string `json:"query"`
	Limit int    `json:"limit,omitempty"`
}

type IamResponse struct {
	Id          string              `json:"id,omitempty"`
	Name        string              `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
	SearchType  string              `json:"searchType,omitempty"`
	Saved       bool                `json:"saved,omitempty"`
	TimeRange   timerange.TimeRange `json:"timeRange,omitempty"`
	Data        IamData             `json:"data,omitempty"`
}

type IamData struct {
	Items []IamItem `json:"items"`
}

type IamItem struct {
	AccessedResourcesCount          int         `json:"accessedResourcesCount"`
	DestCloudAccount                string      `json:"destCloudAccount"`
	DestCloudRegion                 string      `json:"destCloudRegion"`
	DestCloudResourceRrn            string      `json:"destCloudResourceRrn"`
	DestCloudServiceName            string      `json:"destCloudServiceName"`
	DestCloudType                   string      `json:"destCloudType"`
	DestResourceId                  string      `json:"destResourceId"`
	DestResourceName                string      `json:"destResourceName"`
	DestResourceType                string      `json:"destResourceType"`
	EffectiveActionName             string      `json:"effectiveActionName"`
	Exceptions                      []Exception `json:"exceptions"`
	GrantedByCloudEntityId          string      `json:"grantedByCloudEntityId"`
	GrantedByCloudEntityName        string      `json:"grantedByCloudEntityName"`
	GrantedByCloudEntityRrn         string      `json:"grantedByCloudEntityRrn"`
	GrantedByCloudEntityType        string      `json:"grantedByCloudEntityType"`
	GrantedByCloudPolicyId          string      `json:"grantedByCloudPolicyId"`
	GrantedByCloudPolicyName        string      `json:"grantedByCloudPolicyName"`
	GrantedByCloudPolicyRrn         string      `json:"grantedByCloudPolicyRrn"`
	GrantedByCloudPolicyType        string      `json:"grantedByCloudPolicyType"`
	GrantedByCloudType              string      `json:"grantedByCloudType"`
	MessageId                       string      `json:"id"`
	IsWildCardDestCloudResourceName bool        `json:"isWildCardDestCloudResourceName"`
	LastAccessDate                  string      `json:"lastAccessDate"`
	SourceCloudAccount              string      `json:"sourceCloudAccount"`
	SourceCloudRegion               string      `json:"sourceCloudRegion"`
	SourceCloudResourceRrn          string      `json:"sourceCloudResourceRrn"`
	SourceCloudServiceName          string      `json:"sourceCloudServiceName"`
	SourceCloudType                 string      `json:"sourceCloudType"`
	SourceIdpDomain                 string      `json:"sourceIdpDomain"`
	SourceIdpEmail                  string      `json:"sourceIdpEmail"`
	SourceIdpGroup                  string      `json:"sourceIdpGroup"`
	SourceIdpRrn                    string      `json:"sourceIdpRrn"`
	SourceIdpService                string      `json:"sourceIdpService"`
	SourceIdpUsername               string      `json:"sourceIdpUsername"`
	SourcePublic                    bool        `json:"sourcePublic"`
	SourceResourceId                string      `json:"sourceResourceId"`
	SourceResourceName              string      `json:"sourceResourceName"`
	SourceResourceType              string      `json:"sourceResourceType"`
}

type Exception struct {
	MessageCode string `json:"messageCode"`
}
