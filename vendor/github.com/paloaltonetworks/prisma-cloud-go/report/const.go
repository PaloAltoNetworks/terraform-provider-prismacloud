package report

const (
	singular = "report"
	plural   = "reports"
)

//Valid values for alert report types
const (
	CloudSecurityAssessment    = "RIS"
	BusinessUnitReport         = "INVENTORY_OVERVIEW"
	DetailedBusinessUnitReport = "INVENTORY_DETAIL"
)

var alertReportTypes = []string{CloudSecurityAssessment, BusinessUnitReport, DetailedBusinessUnitReport}

var Suffix = []string{"report"}
