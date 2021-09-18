---
page_title: "Prisma Cloud: prismacloud_report"
---

# prismacloud_report

Manage alert reports and compliance reports.

## Example Usage

```hcl
resource "prismacloud_report" "example" {
    name = "test_report"
    report_type = "RIS"
    cloud_type = "aws"
    target{
        time_range {
            relative {
                unit = "hour"
                amount = 24
            }
        }
    }    
}
```

## Argument Reference

* `name` - (Required) Report Name
* `report_type` - (Required) Report type. Valid values are `RIS` (for Cloud Security Assessment report)
  , `INVENTORY_OVERVIEW` (for Business Unit report), `INVENTORY_DETAIL` (for Detailed Business Unit report), or name of
  a compliance standard (for Compliance report)
* `cloud_type` - Cloud type
* `target` - (Required) Model for report target, as defined [below](#target)

### Target

There should be one and only one target block:

* `account_groups` - List of cloud account groups
* `accounts` - List of cloud accounts
* `regions` - List of regions
* `compliance_standard_ids` - List of compliance standard IDs (For Business Unit Report and Detailed Business Unit
  Report)
* `resource_groups` - List of resource groups
* `notify_to` - List of email addresses to receive notification (not supported for Cloud Security Assessment Report)
* `compression_enabled` - (bool) Business unit detailed report compression enabled (For Detailed Business Unit Report)
* `download_now` - (bool) True = download now
* `schedule_enabled` - (bool) Report scheduling enabled (not supported for Cloud Security Assessment Report)
* `schedule` - Recurring report schedule in RRULE format (not supported for Cloud Security Assessment Report)
* `notification_template_id` - Notification template id (not supported for Cloud Security Assessment Report)
* `time_range` - (Required) The time range spec, as defined [below](#time-range).

### Time Range

Only one of these can be defined:

* `absolute` - An absolute time range spec, as defined [below](#absolute-time-range).
* `relative` - A relative time range spec, as defined [below](#relative-time-range).
* `to_now` - A "To Now" time range spec, as defined [below](#to-now-time-range).

### Absolute Time Range

* `start` - (Required, int) Start time.
* `end` - (Required, int) End time.

### Relative Time Range

* `amount` - (Required, int) The time number.
* `unit` - (Required) The time unit.

### To Now Time Range

* `unit` - (Required) The time unit.

## Attribute Reference

* `report_id` - Report ID
* `created_on` - (int) Created on
* `created_by` - Created by
* `last_modified_on` - (int) Last modified on
* `last_modified_by` - Last modified by
* `compliance_standard_id` - Compliance Standard ID
* `status` - Report status
* `next_schedule` - (int) Next schedule
* `last_scheduled` - (int) Last scheduled
* `total_instance_count` - (int) Total instance count
* `counts` - Model for compliance aggregate count, as defined [below](#counts).

### Counts

* `failed` - (int) Failed
* `high_severity_failed` - (int) Number of high-severity failures
* `low_severity_failed` - (int) Number of low-severity failures
* `medium_severity_failed` - (int) Number of medium-severity failures
* `passed` - (int) Passed
* `total` - (int) Total

## Import

Resources can be imported using the report ID:

```
$ terraform import prismacloud_report.example 11111111-2222-3333-4444-555555555555
```
