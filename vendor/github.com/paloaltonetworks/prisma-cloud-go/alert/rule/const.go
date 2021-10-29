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
	TypeMicrosoftTeams       = "microsoft_teams"
	TypeWebhook              = "webhook"
	TypeAwsSecurityHub       = "aws_security_hub"
	TypeGoogleCscc           = "google_cscc"
	TypeServiceNow           = "service_now"
	TypePagerDuty            = "pager_duty"
	TypeDemisto              = "demisto"
	TypeAzureServiceBusQueue = "azure_service_bus_queue"
	TypeSnowFlake            = "snowflake"
	TypeAwsS3                = "aws_s3"
)

const (
	singular = "alert rule"
	plural   = "alert rules"
)
