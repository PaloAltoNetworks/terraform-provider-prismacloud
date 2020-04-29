package section

const (
	singular = "compliance standard requirement section"
	plural   = "compliance standard requirement sections"
)

func RequirementSuffix(id string) []string {
	return []string{"compliance", id, "section"}
}

func SectionSuffix(id string) []string {
	return []string{"compliance", "requirement", "section", id}
}
