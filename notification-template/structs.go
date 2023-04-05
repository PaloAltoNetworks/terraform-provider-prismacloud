package notification_template

type NotificationTemplate struct {
	IntegrationId   string                      `json:"integrationId"`
	IntegrationType string                      `json:"integrationType"`
	Name            string                      `json:"name"`
	TemplateConfig  map[string][]TemplateConfig `json:"templateConfig"`
}
type TemplateConfig struct {
	AliasField     string   `json:"aliasField"`
	DisplayName    string   `json:"displayName"`
	FieldName      string   `json:"fieldName"`
	MaxLength      int      `json:"maxLength"`
	Options        []Option `json:"options"`
	RedlockMapping bool     `json:"redlockMapping"`
	Required       bool     `json:"required"`
	Type           string   `json:"type"`
	TypeaheadUri   string   `json:"typeaheadUri"`
	Value          string   `json:"value"`
}

type Option struct {
	Id   string `json:"id"`
	Key  string `json:"key"`
	Name string `json:"name"`
}
