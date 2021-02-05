---
page_title: "Prisma Cloud: prismacloud_enterprise_settings"
---

# prismacloud_enterprise_settings

Retrieves enterprise settings information.

## Example Usage

```hcl
data "prismacloud_enterprise_settings" "example" {}
```

## Attribute Reference

* `session_timeout` - (int) Browser session timeout.
* `anomaly_training_model_threshold` - Anomaly training model threshold.
* `anomaly_alert_disposition` - Anomaly alert disposition.
* `user_attribution_in_notification` - (bool) User attribution in notification.
* `require_alert_dismissal_note` - (bool) Require alert dismissal note.
* `default_policies_enabled` - (Map of bools) Default policies enabled.
* `apply_default_policies_enabled` - (bool) Apply default policies enabled.
