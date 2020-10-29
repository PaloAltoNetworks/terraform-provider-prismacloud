package integration

type Integration struct {
	Id                string            `json:"id,omitempty"`
	Name              string            `json:"name"`
	IntegrationType   string            `json:"integrationType"`
	IntegrationConfig IntegrationConfig `json:"integrationConfig"`
	Description       string            `json:"description"`
	Enabled           bool              `json:"enabled"`
	CreatedBy         string            `json:"createdBy,omitempty"`
	CreatedTs         int               `json:"createdTs,omitempty"`
	LastModifiedBy    string            `json:"lastModifiedBy,omitempty"`
	LastModifiedTs    int64             `json:"lastModifiedTs,omitempty"`
	Status            string            `json:"status,omitempty"`
	Reason            *Reason           `json:"reason"`
	Valid             bool              `json:"valid,omitempty"`
	AlertRules        []AlertRule       `json:"alertRules,omitempty"`
}

type IntegrationConfig struct {
	// Amazon SQS.
	QueueUrl string `json:"queueUrl,omitempty"`

	// Qualys.
	Login    string `json:"login,omitempty"`
	BaseUrl  string `json:"baseUrl,omitempty"`
	Password string `json:"password,omitempty"`

	// Service Now.
	HostUrl string `json:"hostUrl,omitempty"`
	// Login
	// Password
	Tables  []map[string]bool `json:"tables,omitempty"`
	Version string            `json:"version,omitempty"`

	// Webhook.
	Url     string   `json:"url,omitempty"`
	Headers []Header `json:"headers,omitempty"`

	// PagerDuty.
	IntegrationKey string `json:"integrationKey,omitempty"`
	AuthToken      string `json:"authToken,omitempty"`

	// Slack
	WebHookUrl string `json:"webhookUrl,omitempty"`
}

type Header struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Secure   bool   `json:"secure"`
	ReadOnly bool   `json:"readOnly"`
}

type Reason struct {
	LastUpdated int      `json:"lastUpdated,omitempty"`
	ErrorType   string   `json:"errorType,omitempty"`
	Message     string   `json:"message,omitempty"`
	Details     *Details `json:"details"`
}

type Details struct {
	StatusCode int    `json:"statusCode,omitempty"`
	Subject    string `json:"subject,omitempty"`
	Message    string `json:"i18nKey,omitempty"`
}

type AlertRule struct {
	PolicyScanConfigId string `json:"policyScanConfigId"`
	Name               string `json:"name"`
}
