---
page_title: "Prisma Cloud: prismacloud_notification_template"
---

# prismacloud_notification_template

Manage a notification template.

## Example Usage for Email

```hcl
resource "prismacloud_notification_template" "example-EMAIL" {
  integration_type = "email"
  name             = "Test Terraform Template EMAIL"
  template_config {
    basic_config {
      display_name = "Email template Created by terraform-11"
      field_name   = "custom_note"
      type         = "text"
      value        = "Test Terraform Template for testing purpose"
    }
  }
}
```

## Example Usage for Jira

```hcl
resource "prismacloud_notification_template" "example-JIRA" {
  integration_type = "jira"
  name             = "Terraform Test Template JIRA"
  integration_id   = "<integration-id>"
  template_config {
    basic_config {
      field_name      = "project"
      display_name    = "Project"
      type            = "list"
      redlock_mapping = false
      required        = false
      options {
        id   = "RED"
        key  = "RED"
        name = "RedLock"
      }
      options {
        id   = "BLUE"
        key  = "BLUE"
        name = "RedLock"
      }
      value = "RedLock"
    }
    basic_config {
      field_name      = "issueType"
      display_name    = "Issue Type"
      type            = "list"
      redlock_mapping = false
      required        = true
      options {
        id   = "10002"
        name = "Task"
      }
      value = "Task"
    }
    open {
      display_name    = "State"
      field_name      = "state"
      type            = "list"
      redlock_mapping = false
      required        = true
      options {
        name = "In Review"
        id   = "10001"
      }
      value = "In Review"
    }
    open {
      display_name    = "Summary"
      field_name      = "summary"
      type            = "text"
      redlock_mapping = true
      required        = true
      value           = "AccountId <$AccountId>"
      options {}
    }
    open {
      display_name    = "Description"
      field_name      = "description"
      type            = "text"
      redlock_mapping = true
      required        = true
      value           = "PolicyDescription <$PolicyDescription>"
      options {}
    }
    open {
      display_name    = "Labels"
      field_name      = "labels"
      type            = "array"
      redlock_mapping = false
      type_ahead_uri    = "<type-ahead-uri>"
      required        = true
      options {
        name = "test"
        id   = "test"
      }
      value = "test"
    }
  }
}
```

## Example Usage for Service Now

```hcl
resource "prismacloud_notification_template" "example-SERVICENOW" {

  integration_type = "service_now"
  name             = "Terraform Test Template SERVICENOW"
  integration_id   = "<integration-id>"
  template_config {
    basic_config {
      field_name      = "incidentType"
      display_name    = "Incident Type"
      type            = "list"
      redlock_mapping = false
      required        = true
      options {
        key  = "incident"
        name = "Incident"
      }
      value = "Incident"
    }
    basic_config {
      field_name = "createIncidentOnAlertReOpen"
      value      = true
    }
    dismissed {
      display_name    = "State"
      field_name      = "state"
      type            = "list"
      redlock_mapping = false
      max_length      = 40
      required        = true
      options {
        name = "Canceled"
        key  = "8"
      }
      value = "Canceled"
    }
    resolved {
      display_name    = "State"
      field_name      = "state"
      type            = "list"
      redlock_mapping = false
      max_length      = 40
      required        = true
      options {
        name = "Resolved"
        key  = "6"
      }
      value = "Resolved"
    }
    open {
      display_name    = "State"
      field_name      = "state"
      type            = "list"
      redlock_mapping = false
      max_length      = 40
      required        = true
      options {
        name = "New"
        key  = "1"
      }
      value = "New"
    }
  }
}
```

## Argument Reference

* `integration_id` - (Optional) Integration ID.
* `integration_type` - (Required) Integration type. Valid values are `email`, `jira` or `service_now`.
* `name` - (Required) Template name.
* `enabled` - (Optional, bool) Whether the notification template is enabled (default: `true`).
* `template_type` - (Optional) Type of notification template.
* `template_config` - (Required) Template config spec, as defined [below](#template_config).

### Template Config

* `basic_config` - (Optional) This field includes additional attributes that can be used to customize the notification, as defined [below](#config_params).
* `open` - (Optional) This field represents the notification status when it is first created and has not yet been `resolved`, `dismissed`, or `snoozed`. This field includes additional attributes, as defined [below](#config_params). Applicable only for integration type `jira` and `service_now`.
* `resolved` - (Optional) This field represents the notification status when the issue or event that triggered the notification has been resolved or addressed. This field includes additional attributes, as defined [below](#config_params). Applicable only for integration type `jira` and `service_now`.
* `dismissed` - (Optional) This field represents the notification status when the user has dismissed or ignored the notification. This field includes additional attributes, as defined [below](#config_params). Applicable only for integration type `service_now`.
* `snoozed` - (Optional) This field represents the notification status when the user has chosen to temporarily delay or "snooze" the notification. This field includes additional attributes, as defined [below](#config_params).

#### Status

* `field_name` - (Optional) Field name.
* `display_name` - (Optional) Display name.
* `redlock_mapping` - (Optional,bool) Prisma Cloud will provide the field value for notification.
* `required` - (Optional,bool) Required.
* `type_ahead_uri` - (Optional) URL used to query suggestions for field value.
* `type` - (Optional) Type of field.
* `value` - (Optional) Field value.
* `alias_field` - (Optional) Alias field.
* `max_length` - (Optional,int) Maximum length.
* `options` - (Optional) Options, as defined [below](#options).

##### Options

* `name` - (Optional) Field option name.
* `key` - (Optional) Field option key.
* `id` - (Optional) Field option ID.

## Attribute Reference

* `id` - Notification template id.
* `created_ts` - (int) Creation Unix timestamp in milliseconds.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.
* `integration_name` - Integration name.
* `created_by` - Created by.
* `module` - Module.
* `customer_id` - (int) Customer Id.
