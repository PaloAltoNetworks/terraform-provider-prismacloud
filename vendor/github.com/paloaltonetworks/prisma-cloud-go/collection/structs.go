package collection

type Collection struct {
	Id             string      `json:"id,omitempty"`
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	CreatedBy      string      `json:"createdBy"`
	CreatedTs      int64       `json:"createdTs"`
	LastModifiedBy string      `json:"lastModifiedBy"`
	LastModifiedTs int64       `json:"lastModifiedTs"`
	AssetGroups    AssetGroups `json:"assetGroups"`
}

type AssetGroups struct {
	AccountGroupIds []string `json:"accountGroupIds,omitempty"`
	AccountIds      []string `json:"accountIds,omitempty"`
	RepositoryIds   []string `json:"repositoryIds,omitempty"`
}

type Response struct {
	Value         []Collection `json:"value"`
	NextPageToken string       `json:"nextPageToken"`
}

// To Create A new Collection
type CollectionRequest struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	AssetGroups AssetGroups `json:"assetGroups"`
}
