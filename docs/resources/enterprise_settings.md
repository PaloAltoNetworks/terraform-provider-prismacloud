---
page_title: "Prisma Cloud: prismacloud_enterprise_settings"
---

# prismacloud_enterprise_settings

Manages enterprise settings config.

## Example Usage

```hcl
resource "prismacloud_enterprise_settings" "example" {
    access_key_max_validity = 30
    session_timeout = 60
    default_policies_enabled = {
        "high": true,
        "medium": true,
        "low": false,
    }
}
```

## Argument Reference

* `session_timeout` - (int) Browser session timeout.
* `access_key_max_validity` - (int) Access Keys maximum validity in days.
* `user_attribution_in_notification` - (bool) User attribution in notification.
* `require_alert_dismissal_note` - (bool) Require alert dismissal note.
* `default_policies_enabled` - (Map of bools) Default policies enabled.
* `apply_default_policies_enabled` - (bool) Apply default policies enabled.
* `alarm_enabled` - (bool) Alarms enabled (Default : `true`). Alarms are Prisma Cloud Platform health notifications which are generated to notify users of system level issues/errors. Disabling alarms will delete all existing alarms which were previously generated.
* `named_users_access_keys_expiry_notifications_enabled` - (bool) Named users access keys expiry notifications enabled.
* `service_users_access_keys_expiry_notifications_enabled` - (bool) Service users access keys expiry notifications enabled.
* `notification_threshold_access_keys_expiry` - (int) Notification threshold access keys expiry.
* `audit_log_siem_intgr_ids` - List of integration ids.
* `audit_logs_enabled` - (bool) Enable audit logs.