package prismacloud

type AuthResponse struct {
	Token     string         `json:"token"`
	Message   string         `json:"message"`
	Customers []AuthCustomer `json:"customerNames"`
}

type AuthCustomer struct {
	Name        string `json:"customerName"`
	TosAccepted bool   `json:"tosAccepted"`
}
