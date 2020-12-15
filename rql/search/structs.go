package search

import (
	"github.com/paloaltonetworks/prisma-cloud-go/timerange"
)

type ConfigRequest struct {
	TimeRange        timerange.TimeRange `json:"timeRange"`
	Query            string              `json:"query"`
	Limit            int                 `json:"limit,omitempty"`
	WithResourceJson bool                `json:"withResourceJson,omitempty"`
}

type ConfigResponse struct {
	GroupBy     []string `json:"groupBy"`
	AlertId     string   `json:"alertId"`
	Id          string   `json:"id"`
	CloudType   string   `json:"cloudType"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	SearchType  string   `json:"searchType"`
}
