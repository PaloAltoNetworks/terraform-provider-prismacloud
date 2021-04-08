package history

import (
	"github.com/paloaltonetworks/prisma-cloud-go/timerange"
)

type NameId struct {
	CreatedBy      string `json:"createdBy"`
	LastModifiedBy string `json:"lastModifiedBy"`
	Model          Query  `json:"searchModel"`
}

type Query struct {
	Id          string              `json:"id,omitempty"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	SearchType  string              `json:"searchType"`
	CloudType   string              `json:"cloudType,omitempty"`
	Query       string              `json:"query"`
	Saved       bool                `json:"saved"`
	TimeRange   timerange.TimeRange `json:"timeRange"`
}

// SavedSearch is to create a saved search.  Always have Saved as false.
type SavedSearch struct {
	Name        string              `json:"name"`
	Id          string              `json:"id"`
	TimeRange   timerange.TimeRange `json:"timeRange"`
	Query       string              `json:"query"`
	Saved       bool                `json:"saved"`
	Description string              `json:"description,omitempty"`
}
