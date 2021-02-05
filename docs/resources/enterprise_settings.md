---
page_title: "Prisma Cloud: prismacloud_enterprise_settings"
---

# prismacloud_enterprise_settings

Manages enterprise settings config.

## Example Usage

```hcl
resource "prismacloud_enterprise_settings" "example" {
    session_timeout = 60
    default_policies_enabled = {
        "high": True,
        "medium": True,
        "low": False,
    }
}
```

## Argument Reference

* `session_timeout` - (int) Browser session timeout.
* `anomaly_training_model_threshold` - Anomaly training model threshold (`low`, `medium`, or `high`).
* `anomaly_alert_disposition` - Anomaly alert disposition (`low`, `medium`, or `high`).
* `user_attribution_in_notification` - (bool) User attribution in notification.
* `require_alert_dismissal_note` - (bool) Require alert dismissal note.
* `default_policies_enabled` - (Map of bools) Default policies enabled.
* `apply_default_policies_enabled` - (bool) Apply default policies enabled.
