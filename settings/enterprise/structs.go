package enterprise

type Config struct {
	SessionTimeout                int             `json:"sessionTimeout,omitempty"`
	AnomalyTrainingModelThreshold string          `json:"anomalyTrainingModelThreshold,omitempty"`
	AnomalyAlertDisposition       string          `json:"anomalyAlertDisposition,omitempty"`
	UserAttributionInNotification bool            `json:"userAttributionInNotification"`
	RequireAlertDismissalNote     bool            `json:"requireAlertDismissalNote"`
	DefaultPoliciesEnabled        map[string]bool `json:"defaultPoliciesEnabled"`
	ApplyDefaultPoliciesEnabled   bool            `json:"applyDefaultPoliciesEnabled"`
}
