package enterprise

type Config struct {
	SessionTimeout                                   int             `json:"sessionTimeout,omitempty"`
	UserAttributionInNotification                    bool            `json:"userAttributionInNotification"`
	RequireAlertDismissalNote                        bool            `json:"requireAlertDismissalNote"`
	DefaultPoliciesEnabled                           map[string]bool `json:"defaultPoliciesEnabled"`
	ApplyDefaultPoliciesEnabled                      bool            `json:"applyDefaultPoliciesEnabled"`
	AccessKeyMaxValidity                             int             `json:"accessKeyMaxValidity"`
	AlarmEnabled                                     bool            `json:"alarmEnabled"`
	NamedUsersAccessKeysExpiryNotificationsEnabled   bool            `json:"namedUsersAccessKeysExpiryNotificationsEnabled"`
	ServiceUsersAccessKeysExpiryNotificationsEnabled bool            `json:"serviceUsersAccessKeysExpiryNotificationsEnabled"`
	NotificationThresholdAccessKeysExpiry            int             `json:"notificationThresholdAccessKeysExpiry"`
}
