---
page_title: "Prisma Cloud: prismacloud_notification_templates"
---

# prismacloud_notification_templates

Retrieve a list of notification templates.

## Example Usage

```hcl
data "prismacloud_notification_templates" "example" {}
```

## Attribute Reference

* `total` - (int) Total number of notification templates.
* `listing` - List of notification templates returned, as defined [below](#listing).

### Listing

Each notification template has the following attributes:

* `integration_id` - Integration ID.
* `id` - Notification template id.
* `created_ts` - (int) Creation Unix timestamp in milliseconds.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.
* `created_by` - User who created the notification template.
* `integration_type` - Integration type.
* `name` - Name to be used for the template on the Prisma Cloud platform.
* `integration_name` - Integration name.
* `customer_id` - (int) Prisma customer id.
* `module` - Module.
* `template_type` - Type of notification template.
* `enabled` - (bool) Whether the template is enabled.