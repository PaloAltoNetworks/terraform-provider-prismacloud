---
page_title: "Prisma Cloud: prismacloud_integration"
---

# prismacloud_integration

Manages an integration.

## Example Usage

```hcl
resource "prismacloud_integration" "example" {
    name = "SQS"
    integration_type = "amazon_sqs"
    description = "Made by Terraform"
    enabled = true
    integration_config {
        queue_url = "https://sqs.us-east-1.amazonaws.com/12345/url"
    }
}
```

## Argument Reference

* `name` - (Required) Name of the integration.
* `description` - Description.
* `integration_type` - (Required) Integration type. Valid values are : `okta_idp`, `qualys`, `tenable`, `slack`, `splunk`, `amazon_sqs`, `webhook`, `microsoft_teams`, `azure_service_bus_queue`, `service_now`, `pager_duty`, `demisto`, `google_cscc`, `aws_security_hub`, `aws_s3`, `snowflake`.
* `enabled` - (bool) Enabled. Default: `true` (For outbound integrations (i.e. all integrations except `okta_idp`, `qualys`, `tenable`) this will always be `true` while creating, can be changed to `false` only while updating).
* `integration_config` - (Required) Integration configuration, the values depend on the integration type, as defined [below](#integration-config).

### Integration Config

Refer to the [Prisma Cloud integration documentation](https://prisma.pan.dev/api/cloud/api-integration-config/) if you need more information on a specific integration.

**1. Azure Service Bus Queue**

* `queue_url` - (Required) The URL configured in the Azure Service Bus queue where Prisma cloud sends alerts.
* `account_id` - (Required if you want to use the service principal-based access provided when the Azure cloud account was onboarded to Prisma Cloud) Azure account ID with service principal to which the Azure Service Bus queue belongs.
* `connection_string` - (Required if you want to use a role with limited permissions) Azure Shared Access Signature connection string.

**2. Amazon SQS**

* `queue_url` - (Required) The Queue URL you used when you configured Prisma Cloud in Amazon SQS.
* `more_info` - (Optional, bool) Whether specific IAM credentials are specified for SQS queue access. Set it to `true` while configuring additional IAM information like `role_arn` and `external_id` or `secret_key` and `access_key`.
* `access_key` - (Required if you want to use IAM access keys) AWS access key belonging to AWS IAM credentials meant for SQS queue access.
* `secret_key` - (Required when you are using IAM access keys) AWS secret key for the given access key.
* `role_arn` - (Required if you want to use the IAM Role associated with Prisma Cloud) Role ARN associated with the IAM role on Prisma Cloud
* `external_id` - (Required when you are using the IAM Role associated with Prisma Cloud) External ID associated with the IAM role on Prisma Cloud. New or updated value must be a unique 128-bit UUID.

**3. Qualys**

* `login` - (Required) Qualys Login Username.
* `base_url` - (Required) Qualys Security Operations Center server API URL (without http(s)).
* `password` - (Required) Qualys Password.

**4. ServiceNow**

* `host_url` - (Required) ServiceNow URL.
* `login` - (Required) ServiceNow Login Username.
* `password` - (Required) ServiceNow password for login.
* `tables` - (Required, Map of bools) Key/value pairs that identify the ServiceNow module tables with which to integrate. The possible keys are: `incident`, `sn_si_incident`, `em_event`. The possible values for each key are: `true`, `false`.

**5. Webhook**

* `url` - (Required) Webhook URL.
* `headers` - (Optional) Webhook headers, as defined [below](#headers).

**6. PagerDuty**

* `integration_key` - (Required) PagerDuty integration key.

**7. Slack**

* `webhook_url` - (Required) Slack webhook URL starting with `https://hooks.slack.com/`.

**8. Splunk**

* `auth_token` - (Required) Splunk authentication token for the event collector.
* `url` - (Required) Splunk HTTP event collector URL.
* `source_type` - (Optional) Splunk source type.

**9. Microsoft Teams**

* `url` - (Required) Webhook URL.

**10. Cortex XSOAR**

* `host_url` - (Required) The Cortex XSOAR instance FQDN/IP — either the name or the IP address of the instance.
* `api_key` - (Required) The consumer key you configured when you created the Prisma Cloud application access in your Cortex XSOAR environment.

**11. Tenable**

* `secret_key` - (Required) Secret key from Tenable.io.
* `access_key` - (Required) Access key from Tenable.io.

**12. Google Cloud SCC**

* `source_id` - (Required) GCP source ID for the service account you used to onboard your GCP organization to Prisma Cloud.
* `org_id` - (Required) GCP organization ID.

**13. Okta**

* `domain` - (Required) Okta domain name.
* `api_token` - (Required) The authentication API token for Okta. The token must be of type Read-Only Admin.

**14. Amazon S3**

* `s3_uri` - (Required) Amazon S3 bucket URI.
* `region` - (Required) AWS region where the S3 bucket resides.
* `role_arn` - (Required) Role ARN associated with the IAM role on Prisma Cloud.
* `external_id` - (Required) External ID associated with the IAM role on Prisma Cloud. Any new or updated value must be a unique 128-bit UUID.
* `roll_up_interval` - (Required, int) Time in minutes at which batching of Prisma Cloud alerts would roll up. Valid values are `15`, `30`, `60`, or `180`.

**15. AWS Security Hub**

* `account_id` - (Required) AWS account ID to which you assigned AWS Security Hub read-only access.
* `regions` - (Required) List of AWS regions, as defined [below](#regions).

**16. Snowflake**

* `host_url` - (Required) Snowflake Account URL. Format should be 'YOURACCOUNTNAME.snowflakecomputing.com'.
* `user_name` - (Required) Snowflake Username.
* `staging_integration_id` - (Required) Existing Amazon S3 integration ID.
* `pipe_name` - (Required) Snowpipe Name. Format should be '<db_name>.<schema_name>.<pipe_name>'.
* `private_key` - (Required) Private Key.
* `pass_phrase` - (Required if you are using encrypted key) PassPhrase for private key.
* `roll_up_interval` - (Required, int) Time in minutes at which batching of Prisma Cloud alerts would roll up. Valid values are `15`, `30`, `60`, or `180`.

#### Headers

* `key` - (Required) Header name.
* `value` - (Required) Header value.
* `secure` - (bool) Secure.
* `read_only` - (bool) Read-only.

#### Regions

* `name` - AWS region name e.g. `AWS California`.
* `api_identifier` - AWS region code e.g. `us-west-1`.
* `cloud_type` - Cloud Type (default: `aws`).

## Attribute Reference

* `integration_id` - Integration ID.
* `status` - Status.
* `valid` - (bool) Valid.
* `created_by` - Created by.
* `created_ts` - (int) Created timestamp.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.
* `reason` - Model for the integration status details, as defined [below](#reason).

### Reason

* `last_updated` - (int) Last updated.
* `error_type` - Error type.
* `message` - Message.
* `details` - Model for message details, as defined [below](#details).

### Details

* `status_code` - (int) Status code.
* `subject` - Subject.
* `message` - Internationalization key.

In `integration_config` section, the following attributes are available:

* `version` - Cortex release version.