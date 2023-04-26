package ip_address

type IdList struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type LoginIpAllow struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Cidr []string `json:"cidr"`
	Description string `json:"description"`
	LastModifiedTs int `json:"lastModifiedTs"`
}
