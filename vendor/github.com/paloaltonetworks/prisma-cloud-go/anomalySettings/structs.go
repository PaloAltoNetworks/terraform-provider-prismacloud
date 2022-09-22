package anomalySettings

type Response struct {
	AnomalySettingsResponse struct {
		AlertDisposition            string                      `json:"alertDisposition"`
		AlertDispositionDescription AlertDispositionDescription `json:"alertDispositionDescription"`
		PolicyDescription           string                      `json:"policyDescription"`
		PolicyName                  string                      `json:"policyName,omitempty"`
		TrainingModelDescription    TrainingModelDescription    `json:"trainingModelDescription"`
		TrainingModelThreshold      string                      `json:"trainingModelThreshold"`
	}
}

type AnomalySettingsResponse struct {
	AlertDisposition            string                      `json:"alertDisposition"`
	AlertDispositionDescription AlertDispositionDescription `json:"alertDispositionDescription"`
	PolicyDescription           string                      `json:"policyDescription"`
	PolicyName                  string                      `json:"policyName,omitempty"`
	TrainingModelDescription    TrainingModelDescription    `json:"trainingModelDescription"`
	TrainingModelThreshold      string                      `json:"trainingModelThreshold"`
}

type AnomalySettings struct {
	PolicyId                    string                      `json:"id"`
	AlertDisposition            string                      `json:"alertDisposition"`
	AlertDispositionDescription AlertDispositionDescription `json:"alertDispositionDescription"`
	PolicyDescription           string                      `json:"policyDescription"`
	PolicyName                  string                      `json:"policyName,omitempty"`
	TrainingModelDescription    TrainingModelDescription    `json:"trainingModelDescription"`
	TrainingModelThreshold      string                      `json:"trainingModelThreshold"`
	Type                        string                      `json:"type"`
}

type AlertDispositionDescription struct {
	Aggressive   string `json:"aggressive"`
	Moderate     string `json:"moderate"`
	Conservative string `json:"conservative"`
}

type TrainingModelDescription struct {
	Low    string `json:"low"`
	Medium string `json:"medium"`
	High   string `json:"high"`
}
