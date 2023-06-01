package notification_template

// To Create A new Template
type NotificationTemplateRequest struct {
	IntegrationType string               `json:"integrationType"`
	IntegrationId   string               `json:"integrationId,omitempty"`
	TemplateType    string               `json:"templateType,omitempty"`
	Enabled         bool                 `json:"enabled,omitempty"`
	Name            string               `json:"name"`
	TemplateConfig  TemplateConfigStruct `json:"templateConfig"`
}

type TemplateConfigStruct struct {
	Open        []Config `json:"open,omitempty"`
	Resolved    []Config `json:"resolved,omitempty"`
	Dismissed   []Config `json:"dismissed,omitempty"`
	BasicConfig []Config `json:"basic_config,omitempty"`
	Snoozed     []Config `json:"snoozed,omitempty"`
}
type Config struct {
	AliasField     string   `json:"aliasField,omitempty"`
	DisplayName    string   `json:"displayName,omitempty"`
	FieldName      string   `json:"fieldName,omitempty"`
	MaxLength      int      `json:"maxLength,omitempty"`
	Options        []Option `json:"options"`
	RedlockMapping bool     `json:"redlockMapping,omitempty"`
	Required       bool     `json:"required,omitempty"`
	Type           string   `json:"type,omitempty"`
	TypeaheadUri   string   `json:"typeaheadUri,omitempty"`
	Value          string   `json:"value,omitempty"`
}

type Option struct {
	Id   string `json:"id,omitempty"`
	Key  string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`
}

type NotificationTemplate struct {
	Id              string               `json:"id"`
	IntegrationId   string               `json:"integrationId"`
	CreatedTs       int64                `json:"createdTs"`
	IntegrationType string               `json:"integrationType"`
	Name            string               `json:"name"`
	LastModifiedBy  string               `json:"lastModifiedBy"`
	LastModifiedTs  int64                `json:"LastModifiedTs"`
	IntegrationName string               `json:"integrationName"`
	CreatedBy       string               `json:"createdBy"`
	CustomerId      int32                `json:"customerId"`
	Enabled         bool                 `json:"enabled"`
	Module          string               `json:"module"`
	TemplateType    string               `json:"templateType"`
	TemplateConfig  TemplateConfigStruct `json:"templateConfig"`
}

type LicenseInfo struct {
	PrismaId string `json:"prismaId"`
}
