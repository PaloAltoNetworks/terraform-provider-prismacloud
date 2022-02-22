package integration

const (
	singular = "integration"
	plural   = "integrations"
)

var Suffix = []string{"integration"}
var v1Suffix = []string{"api", "v1", "tenant"}
var LicenseSuffix = []string{"license"}

var InboundIntegrations = []string{"okta_idp", "qualys", "tenable"}

var OutboundIntegrations = []string{"slack", "splunk", "amazon_sqs", "webhook", "microsoft_teams", "azure_service_bus_queue",
	"service_now", "pager_duty", "demisto", "google_cscc", "aws_security_hub", "aws_s3", "snowflake"}
