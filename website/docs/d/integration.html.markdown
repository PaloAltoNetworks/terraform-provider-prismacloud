---
layout: "prismacloud"
page_title: "Prisma Cloud: prismacloud_integration"
description: |-
  Retrieves integration information.
---

# prismacloud_integration

Retrieves integration information.

## Example Usage

```hcl
data "prismacloud_integration" "example" {
    name = "myIntegration"
}
```

## Argument Reference

One of the following must be specified:

* `name` - Name of the integration.
* `integration_id` - Integration ID.

## Attribute Reference

* `description` - Description.
* `integration_type` - Integration type.
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

* `queue_url` - The Queue URL you used when you configured Prisma Cloud in Amazon SQS
* `login` - (Qualys/ServiceNow) Login.
* `base_url` - Qualys Security Operations Center server API URL (without "http(s)")
* `password` - (Qualys/ServiceNow) Password
* `host_url` - ServiceNow URL.
* `tables` - (Map of bools) Key/value pairs that identify the ServiceNow module tables with which to integrate (e.g. - incident, sn_si_incident, or em_event).
* `version` - ServiceNow release version.
* `url` - Webhook URL.
* `auth_token` - (Webhook/PagerDuty) The authentication token for the event collector.
* `integration_key` - PagerDuty integration key.
