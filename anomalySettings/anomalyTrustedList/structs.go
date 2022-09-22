package anomalyTrustedList

type AnomalyTrustedListRequest struct {
	Name               string             `json:"name"`
	Description        string             `json:"description"`
	ApplicablePolicies []string           `json:"applicablePolicies"`
	TrustedListType    string             `json:"trustedListType"`
	AccountId          string             `json:"accountID"`
	VPC                string             `json:"vpc"`
	TrustedListEntries []TrustedListEntry `json:"trustedListEntries"`
}

type AnomalyTrustedList struct {
	Atl_Id             int                `json:"id"`
	Name               string             `json:"name"`
	Description        string             `json:"description"`
	ApplicablePolicies []string           `json:"applicablePolicies"`
	TrustedListType    string             `json:"trustedListType"`
	AccountId          string             `json:"accountID"`
	VPC                string             `json:"vpc"`
	TrustedListEntries []TrustedListEntry `json:"trustedListEntries"`
	CreatedBy          string             `json:"createdBy"`
	CreatedOn          int                `json:"createdOn"`
}

type TrustedListEntry struct {
	ImageID    string `json:"imageID"`
	IpCIDR     string `json:"ipCIDR"`
	Port       string `json:"port"`
	ResourceID string `json:"resourceID"`
	Service    string `json:"service"`
	Subject    string `json:"subject"`
	TagKey     string `json:"tagKey"`
	TagValue   string `json:"tagValue"`
	Domain     string `json:"domain"`
	Protocol   string `json:"protocol"`
}
