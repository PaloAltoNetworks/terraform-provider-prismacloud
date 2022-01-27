package profile

type Profile struct {
	Id                  string   `json:"id,omitempty"`
	AccountType         string   `json:"type"`
	Username            string   `json:"username"`
	FirstName           string   `json:"firstName,omitempty"`
	LastName            string   `json:"lastName,omitempty"`
	DisplayName         string   `json:"displayName,omitempty"`
	Email               string   `json:"email,omitempty"`
	AccessKeysAllowed   bool     `json:"accessKeysAllowed"`
	AccessKeyExpiration int      `json:"accessKeyExpiration,omitempty"`
	AccessKeyName       string   `json:"accessKeyName,omitempty"`
	DefaultRoleId       string   `json:"defaultRoleId"`
	EnableKeyExpiration bool     `json:"enableKeyExpiration,omitempty"`
	RoleIds             []string `json:"roleIds"`
	TimeZone            string   `json:"timeZone"`
	Enabled             bool     `json:"enabled"`
	LastLoginTs         int      `json:"lastLoginTs,omitempty"`
	LastModifiedBy      string   `json:"lastModifiedBy,omitempty"`
	LastModifiedTs      int      `json:"lastModifiedTs,omitempty"`
	AccessKeysCount     int      `json:"accessKeysCount,omitempty"`
	Roles               []Role   `json:"roles,omitempty"`
}

type AccessKeyResponse struct {
	AccessKeyId string `json:"id"`
	SecretKey   string `json:"secretKey"`
}

type Role struct {
	RoleId                 string `json:"id"`
	Name                   string `json:"name"`
	OnlyAllowCIAccess      bool   `json:"onlyAllowCIAccess"`
	OnlyAllowComputeAccess bool   `json:"onlyAllowComputeAccess"`
	OnlyAllowReadAccess    bool   `json:"onlyAllowReadAccess"`
	RoleType               string `json:"type"`
}
