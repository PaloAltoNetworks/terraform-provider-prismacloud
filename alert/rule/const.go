package rule

var Suffix = []string{"alert", "rule"}

// Valid values for NotificationConfig.Frequency.
const (
	FrequencyAsItHappens = "as_it_happens"
	FrequencyDaily       = "daily"
	FrequencyWeekly      = "weekly"
	FrequencyMonthly     = "monthly"
)

// valid values for NotificationConfig.Type.
const (
	TypeEmail                = "email"
	TypeSlack                = "slack"
	TypeSplunk               = "splunk"
	TypeAmazonSqs            = "amazon_sqs"
	TypeJira                 = "jira"
	TypeMicrosoftTeams       = "microsoft_teams"
	TypeWebhook              = "webhook"
	TypeAwsSecurityHub       = "aws_security_hub"
	TypeGoogleCscc           = "google_cscc"
	TypeServiceNow           = "service_now"
	TypePagerDuty            = "pager_duty"
	TypeDemisto              = "demisto"
	TypeAzureServiceBusQueue = "azure_service_bus_queue"
	TypeTenable              = "tenable"
	TypeOkta                 = "okta_idp"
	TypeSnowFlake            = "snowflake"
)

const (
	singular = "alert rule"
	plural   = "alert rules"
)
