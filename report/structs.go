package report

import "github.com/paloaltonetworks/prisma-cloud-go/timerange"

type Report struct {
	Id                   string `json:"id,omitempty"`
	Name                 string `json:"name"`
	Type                 string `json:"type"`
	CloudType            string `json:"cloudType"`
	ComplianceStandardId string `json:"complianceStandardId,omitempty"`
	Target               Target `json:"target"`
	Status               string `json:"status,omitempty"`
	CreatedOn            int    `json:"createdOn,omitempty"`
	CreatedBy            string `json:"createdBy,omitempty"`
	LastModifiedOn       int    `json:"lastModifiedOn,omitempty"`
	LastModifiedBy       string `json:"lastModifiedBy,omitempty"`
	NextSchedule         int    `json:"nextSchedule,omitempty"`
	LastScheduled        int    `json:"lastScheduled,omitempty"`
	TotalInstanceCount   int    `json:"totalInstanceCount,omitempty"`
	Counts               Counts `json:"counts,omitempty"`
}

type Target struct {
	AccountGroups          []string            `json:"accountGroups"`
	Accounts               []string            `json:"accounts"`
	Regions                []string            `json:"regions"`
	ComplianceStandardIds  []string            `json:"complianceStandardIds"`
	CompressionEnabled     bool                `json:"compressionEnabled"`
	DownloadNow            bool                `json:"downloadNow"`
	NotifyTo               []string            `json:"notifyTo"`
	ResourceGroups         []string            `json:"resourceGroups"`
	Schedule               string              `json:"schedule"`
	ScheduleEnabled        bool                `json:"scheduleEnabled"`
	NotificationTemplateId string              `json:"notificationTemplateId"`
	TimeRange              timerange.TimeRange `json:"timeRange"`
}

type Counts struct {
	Failed               int `json:"failed"`
	HighSeverityFailed   int `json:"highSeverityFailed"`
	LowSeverityFailed    int `json:"lowSeverityFailed"`
	MediumSeverityFailed int `json:"mediumSeverityFailed"`
	Passed               int `json:"passed"`
	Total                int `json:"total"`
}
