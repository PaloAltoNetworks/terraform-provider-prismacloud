---
page_title: "Prisma Cloud: prismacloud_integration"
---

# prismacloud_integration

Retrieves integration information.

## Example Usage

```hcl
data "prismacloud_integration" "example" {
    name = "myIntegration"
    integration_type = "amazon_sqs"
}
```

## Argument Reference

* `integration_type` - (Required) Integration type.

One of the following must be specified:

* `name` - Name of the integration.
* `integration_id` - Integration ID.

## Attribute Reference

* `description` - Description.
* `enabled` - (bool) Enabled.
* `created_by` - Created by.
* `created_ts` - (int) Created on.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.
* `status` - Status.
* `valid` - (bool) Valid.
* `reason` - Model for the integration status details, as defined [below](#reason).
* `integration_config` - Integration configuration, the values depend on the integration type, as defined [below](#integration-config).

### Reason

* `last_updated` - (int) Last updated.
* `error_type` - Error type.
* `message` - Message.
* `details` - Model for message details, as defined [below](#details).

### Details

* `status_code` - (int) Status code.
* `subject` - Subject.
* `message` - Internationalization key.

### Integration Config

**1. Azure Service Bus Queue**

* `queue_url` - The URL configured in the Azure Service Bus queue where Prisma cloud sends alerts.
* `account_id` - Azure account ID with service principal to which the Azure Service Bus queue belongs.
* `connection_string` - Azure Shared Access Signature connection string.

**2. Amazon SQS**

* `queue_url` - The Queue URL you used when you configured Prisma Cloud in Amazon SQS.
* `more_info` - (bool) Whether specific IAM credentials are specified for SQS queue access.
* `access_key` - AWS access key belonging to AWS IAM credentials meant for SQS queue access.
* `secret_key` - AWS secret key for the given access key.
* `role_arn` - Role ARN associated with the IAM role on Prisma Cloud
* `external_id` - External ID associated with the IAM role on Prisma Cloud.

**3. Qualys**

* `login` - Qualys Login Username.
* `base_url` - Qualys Security Operations Center server API URL.
* `password` - Qualys Password.

**4. ServiceNow**

* `host_url` - ServiceNow URL.
* `login` - ServiceNow Login Username.
* `password` - ServiceNow password for login.
* `tables` - (Map of bools) Key/value pairs that identify the ServiceNow module tables with which to integrate.

**5. Webhook**

* `url` - Webhook URL.
* `headers` - Webhook headers, as defined [below](#headers).

**6. PagerDuty**

* `integration_key` - PagerDuty integration key.

**7. Slack**

* `webhook_url` - Slack webhook URL starting with `https://hooks.slack.com/`.

**8. Splunk**

* `auth_token` - Splunk authentication token for the event collector.
* `url` - Splunk HTTP event collector URL.
* `source_type` - Splunk source type.

**9. Microsoft Teams**

* `url` - Webhook URL.

**10. Cortex XSOAR**

* `host_url` - The Cortex XSOAR instance FQDN/IP — either the name or the IP address of the instance.
* `api_key` - The consumer key you configured when you created the Prisma Cloud application access in your Cortex XSOAR environment.
* `version` - Cortex release version. 

**11. Tenable**

* `secret_key` - Secret key from Tenable.io.
* `access_key` - Access key from Tenable.io.

**12. Google Cloud SCC**

* `source_id` - GCP source ID for the service account you used to onboard your GCP organization to Prisma Cloud.
* `org_id` - GCP organization ID.

**13. Okta**

* `domain` - Okta domain name.
* `api_token` - The authentication API token for Okta.

**14. Amazon S3**

* `s3_uri` - Amazon S3 bucket URI.
* `region` - AWS region where the S3 bucket resides.
* `role_arn` - Role ARN associated with the IAM role on Prisma Cloud.
* `external_id` - External ID associated with the IAM role on Prisma Cloud.
* `roll_up_interval` - (int) Time in minutes at which batching of Prisma Cloud alerts would roll up.

**15. AWS Security Hub**

* `account_id` - AWS account ID to which you assigned AWS Security Hub read-only access.
* `regions` - List of AWS regions, as defined [below](#regions).

**16. Snowflake**

* `host_url` - Snowflake Account URL.
* `user_name` - Snowflake Username.
* `staging_integration_id` - Existing Amazon S3 integration ID.
* `pipe_name` - Snowpipe Name.
* `private_key` - Private Key.
* `pass_phrase` - PassPhrase for private key.
* `roll_up_interval` - (int) Time in minutes at which batching of Prisma Cloud alerts would roll up.

#### Headers

* `key` - Header name.
* `value` - Header value.
* `secure` - (bool) Secure.
* `read_only` - (bool) Read-only.

#### Regions

* `name` - AWS region name.
* `api_identifier` - AWS region code.
* `cloud_type` - Cloud Type.
