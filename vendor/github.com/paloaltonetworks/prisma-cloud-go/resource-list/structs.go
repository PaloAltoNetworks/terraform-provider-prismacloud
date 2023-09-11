package resource_list

type ResourceList struct {
	Id               string        `json:"id,omitempty"`
	Description      string        `json:"description"`
	Name             string        `json:"name"`
	ResourceListType string        `json:"resourceListType"`
	LastModifiedTs   int64         `json:"lastModifiedTs"`
	LastModifiedBy   string        `json:"lastModifiedBy"`
	Members          []interface{} `json:"members"`
}

// To Create A new ResourceList
type ResourceListRequest struct {
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	ResourceListType string        `json:"resourceListType"`
	Members          []interface{} `json:"members"`
}
