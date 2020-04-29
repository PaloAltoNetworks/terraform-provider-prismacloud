package requirement

const (
	singular = "compliance standard requirement"
	plural   = "compliance standard requirements"
)

func ComplianceSuffix(cid string) []string {
	return []string{"compliance", cid, "requirement"}
}

func RequirementSuffix(id string) []string {
	return []string{"compliance", "requirement", id}
}
