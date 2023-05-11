package notification_template

const (
	singular = "notification-template"
	plural   = "notification-templates"
)

var Suffix1 = []string{"api", "v1", "tenant"}
var Suffix2 = []string{"template"}
var LicenseSuffix = []string{"license"}

const (
	Email      = "email"
	Jira       = "jira"
	ServiceNow = "service_now"

	ListType    = "list"
	TextType    = "text"
	ArrayType   = "array"
	BoolType    = "bool"
	IntegerType = "integer"
)
