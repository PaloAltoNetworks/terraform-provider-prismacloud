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
	Id          string   `json:"id"`
	CloudType   string   `json:"cloudType"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	SearchType  string   `json:"searchType"`
}

type EventRequest struct {
	TimeRange        timerange.TimeRange `json:"timeRange"`
	Query            string              `json:"query"`
	Limit            int                 `json:"limit,omitempty"`
	WithResourceJson bool                `json:"withResourceJson,omitempty"`
}

type EventResponse struct {
	GroupBy     []string `json:"groupBy"`
	Id          string   `json:"id"`
	CloudType   string   `json:"cloudType"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	SearchType  string   `json:"searchType"`
}

type NetworkRequest struct {
	TimeRange        timerange.TimeRange `json:"timeRange"`
	Query            string              `json:"query"`
	Limit            int                 `json:"limit,omitempty"`
	WithResourceJson bool                `json:"withResourceJson,omitempty"`
}

type NetworkResponse struct {
	GroupBy     []string `json:"groupBy"`
	Id          string   `json:"id"`
	CloudType   string   `json:"cloudType"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	SearchType  string   `json:"searchType"`
}
