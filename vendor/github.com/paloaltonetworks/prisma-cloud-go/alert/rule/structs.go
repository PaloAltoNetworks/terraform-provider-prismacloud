package rule

type Rule struct {
	PolicyScanConfigId   string               `json:"policyScanConfigId,omitempty"`
	Name                 string               `json:"name"`
	Description          string               `json:"description"`
	Enabled              bool                 `json:"enabled"`
	ScanAll              bool                 `json:"scanAll"`
	Policies             []string             `json:"policies"`
	PolicyLabels         []string             `json:"policyLabels"`
	ExcludedPolicies     []string             `json:"excludedPolicies"`
	Target               Target               `json:"target"`
	LastModifiedOn       int                  `json:"lastModifiedOn,omitempty"`
	LastModifiedBy       string               `json:"lastModifiedBy,omitempty"`
	NotificationConfig   []NotificationConfig `json:"alertRuleNotificationConfig,omitempty"`
	AllowAutoRemediate   bool                 `json:"allowAutoRemediate"`
	DelayNotificationMs  int                  `json:"delayNotificationMs"`
	NotifyOnOpen         bool                 `json:"notifyOnOpen"`
	NotifyOnSnoozed      bool                 `json:"notifyOnSnoozed,omitempty"`
	NotifyOnDismissed    bool                 `json:"notifyOnDismissed,omitempty"`
	NotifyOnResolved     bool                 `json:"notifyOnResolved,omitempty"`
	Owner                string               `json:"owner,omitempty"`
	NotificationChannels []string             `json:"notificationChannels,omitempty"`
	OpenAlertsCount      int                  `json:"openAlertsCount,omitempty"`
	ReadOnly             bool                 `json:"readOnly,omitempty"`
	Deleted              bool                 `json:"deleted,omitempty"`
}

type Target struct {
	AccountGroups    []string `json:"accountGroups"`
	ExcludedAccounts []string `json:"excludedAccounts,omitempty"`
	Regions          []string `json:"regions,omitempty"`
	Tags             []Tag    `json:"tags,omitempty"`
}

type Tag struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type NotificationConfig struct {
	Id                 string   `json:"id,omitempty"`
	Frequency          string   `json:"frequency,omitempty"`
	Enabled            bool     `json:"enabled"`
	Recipients         []string `json:"recipients"`
	DetailedReport     bool     `json:"detailedReport"`
	WithCompression    bool     `json:"withCompression"`
	IncludeRemediation bool     `json:"includeRemediation"`
	LastUpdated        int      `json:"lastUpdated"`
	LastSentTs         int      `json:"last_send_ts"`
	Type               string   `json:"type"`
	TemplateId         string   `json:"templateId,omitempty"`
	TimezoneId         string   `json:"timezone,omitempty"`
	DayOfMonth         int      `json:"dayOfMonth,omitempty"`
	RruleSchedule      string   `json:"rruleSchedule,omitempty"`
	FrequencyFromRrule string   `json:"frequencyFromRRule,omitempty"`
	HourOfDay          int      `json:"hourOfDay,omitempty"`
	DaysOfWeek         []Day    `json:"daysOfWeek,omitempty"`
}

type Day struct {
	Day    string `json:"day"`
	Offset int    `json:"offset"`
}
