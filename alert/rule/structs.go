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
	CreatedOn            int                  `json:"createdOn,omitempty"`
	CreatedBy            string               `json:"createdBy,omitempty"`
	LastModifiedOn       int                  `json:"lastModifiedOn,omitempty"`
	LastModifiedBy       string               `json:"lastModifiedBy,omitempty"`
	NotificationConfig   []NotificationConfig `json:"alertRuleNotificationConfig"`
	AllowAutoRemediate   bool                 `json:"allowAutoRemediate"`
	DelayNotificationMs  int                  `json:"delayNotificationMs"`
	NotifyOnOpen         bool                 `json:"notifyOnOpen"`
	NotifyOnSnoozed      bool                 `json:"notifyOnSnoozed"`
	NotifyOnDismissed    bool                 `json:"notifyOnDismissed"`
	NotifyOnResolved     bool                 `json:"notifyOnResolved"`
	Owner                string               `json:"owner,omitempty"`
	NotificationChannels []string             `json:"notificationChannels"`
	OpenAlertsCount      int                  `json:"openAlertsCount,omitempty"`
	ReadOnly             bool                 `json:"readOnly,omitempty"`
	Deleted              bool                 `json:"deleted,omitempty"`
}

type Target struct {
	AccountGroups    []string `json:"accountGroups"`
	ExcludedAccounts []string `json:"excludedAccounts"`
	Regions          []string `json:"regions"`
	Tags             []Tag    `json:"tags"`
}

type Tag struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type NotificationConfig struct {
	Id                 string   `json:"id"`
	Frequency          string   `json:"frequency"`
	Enabled            bool     `json:"enabled"`
	Recipients         []string `json:"recipients"`
	DetailedReport     bool     `json:"detailedReport"`
	WithCompression    bool     `json:"withCompression"`
	IncludeRemediation bool     `json:"includeRemediation"`
	LastUpdated        int      `json:"lastUpdated"`
	LastSentTs         int      `json:"last_send_ts"`
	Type               string   `json:"type"`
	TemplateId         string   `json:"templateId"`
	Timezone           Timezone `json:"timezone"`
	DayOfMonth         int      `json:"dayOfMonth"`
	RruleSchedule      string   `json:"rruleSchedule"`
	FrequencyFromRrule string   `json:"frequencyFromRRule"`
	HourOfDay          int      `json:"hourOfDay"`
	DaysOfWeek         []Day    `json:"daysOfWeek"`
}

type Timezone struct {
	Id string `json:"id"`
}

type Day struct {
	Day    string `json:"day"`
	Offset int    `json:"offset"`
}
