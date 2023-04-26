package trustedalertip

type TrustedAlertIP struct {
	UUID      string  `json:"uuid"`
	Name      string  `json:"name"`
	CidrCount int     `json:"cidrCount"`
	CIDRS     []CIDRS `json:"cidrs"`
}

type CIDRS struct {
	CIDR        string `json:"cidr"`
	UUID        string `json:"uuid"`
	CreatedOn   int    `json:"createdOn"`
	Description string `json:"description"`
}
