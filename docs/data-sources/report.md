---
page_title: "Prisma Cloud: prismacloud_report"
---

# prismacloud_report

Retrieve information on a specific alert report or compliance report.

## Example Usage

```hcl
data "prismacloud_report" "example" {
    name = "My Report"
}
```

## Argument Reference

You must specify at least one of the following:

* `report_id` - Report ID
* `name` - Report name

## Attribute Reference

* `report_type` - Report type
* `cloud_type` - Cloud type
* `created_on` - (int) Created on
* `created_by` - Created by
* `last_modified_on` - (int) Last modified on
* `last_modified_by` - Last modified by
* `compliance_standard_id` - Compliance Standard ID
* `status` - Report status
* `next_schedule` - (int) Next schedule
* `last_scheduled` - (int) Last scheduled
* `total_instance_count` - (int) Total instance count
* `target` - Model for report target, as defined [below](#target)
* `counts` - Model for compliance aggregate count, as defined [below](#counts).

### Target

* `account_groups` - List of cloud account groups
* `accounts` - List of cloud accounts
* `regions` - List of regions
* `compliance_standard_ids` - List of compliance standard IDs
* `resource_groups` - List of resource groups
* `notify_to` - List of email addresses to receive notification
* `compression_enabled` - (bool) Business unit detailed report compression enabled
* `download_now` - (bool) True = download now
* `schedule_enabled` - (bool) Report scheduling enabled
* `schedule` - Recurring report schedule in RRULE format
* `notification_template_id` - Notification template id
* `time_range` - (Required) The time range spec, as defined [below](#time-range).

### Time Range

* `absolute` - An absolute time range spec, as defined [below](#absolute-time-range)
* `relative` - A relative time range spec, as defined [below](#relative-time-range)
* `to_now` - A "To Now" time range spec, as defined [below](#to-now-time-range)

### Absolute Time Range

* `start` - (int) Start time
* `end` - (int) End time

### Relative Time Range

* `amount` - (int) The time number
* `unit` - The time unit

### To Now Time Range

* `unit` - The time unit

### Counts

* `failed` - (int) Failed
* `high_severity_failed` - (int) Number of high-severity failures
* `low_severity_failed` - (int) Number of low-severity failures
* `medium_severity_failed` - (int) Number of medium-severity failures
* `passed` - (int) Passed
* `total` - (int) Total
