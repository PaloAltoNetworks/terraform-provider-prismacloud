---
page_title: "Prisma Cloud: prismacloud_notification_template"
---

# prismacloud_notification_template

Retrieve information on a specific notification template.

## Example Usage

```hcl
data "prismacloud_notification_template" "example" {
   id = "<notification-template-id>"
}
```

## Argument Reference

You must specify:

* `id` - Notification template ID.

## Attribute Reference

* `integration_id` - Integration ID.
* `created_ts` - (int) Creation Unix timestamp in milliseconds.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.
* `integration_type` - Integration type.
* `created_by` - User who created the notification template.
* `name` - Name to be used for the template on the Prisma Cloud platform.
* `integration_name` - Integration name.
* `customer_id` - (int) Prisma customer ID.
* `module` - Module.
* `template_type` - Type of notification template.
* `enabled` - (bool) Whether the template is enabled.
* `template_config` - Template config spec, as defined [below](#template_config).

## Template Config

* `basic_config` - This field includes additional attributes that can be used to customize the notification, as defined [below](#config_params).
* `open` - Provide config to map the `open` alert state to `jira`/`service_now` state and configure the `jira`/`service_now` fields. This field includes additional attributes, as defined [below](#config_params). 
* `resolved` - Provide config to map the `resolved` alert state to `jira`/`service_now` state and configure the `jira`/`service_now` fields. This field includes additional attributes, as defined [below](#config_params). 
* `dismissed` - Provide config to map the `dismissed` alert state to `service_now` state and configure the `service_now` fields. This field includes additional attributes, as defined [below](#config_params). 
* `snoozed` - This field represents the notification status when the user has chosen to temporarily delay or "snooze" the notification. This field includes additional attributes, as defined [below](#config_params).

### Config Params

* `field_name` - Field name.
* `display_name` - Display name.
* `redlock_mapping` - (bool) Prisma Cloud will provide the field value for notification.
* `required` - (bool) Required.
* `type_ahead_uri` - URL used to query suggestions for field value.
* `type` - Type of field.
* `value` - Field value.
* `alias_field` - Alias field.
* `max_length` - (int) Maximum length.
* `options` - Options, as defined [below](#options).

#### Options

* `name` - Field option name.
* `key` - Field option key.
* `id` - Field option ID.