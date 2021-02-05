---
page_title: "Prisma Cloud: prismacloud_alerts"
---

# prismacloud_alerts

Data source to return information on current alerts in Prisma Cloud.

## Example Usage

```hcl
data "prismacloud_alerts" "info" {
    limit = 2
    time_range {
        relative {
            amount = 48
            unit = "hour"
        }
    }
}

output "alerts" {
    value = data.prismacloud_alerts.info.listing
}
```

## Argument Reference

* `time_range` - (Required) The time range spec, as defined [below](#time-range).
* `limit` - (Optional, int) Max number of alerts to return (default: `10000`).
* `filters` - (Optional) Filtering parameters spec, as defined [below](#filters).
* `sort_by` - (Optional) Array of sort properties. Append :asc or :desc to the key to sort by ascending or descending order respectively.

### Time Range

The `time_range` block allows you to specify one of multiple supported time ranges.  Only one time range can be specified.

* `absolute` - An absolute time range spec, as defined [below](#absolute-time-range).
* `relative` - A relative time range spec, as defined [below](#relative-time-range).
* `to_now` - A to-now time range spec, as defined [below](#to-now-time-range).

### Absolute Time Range

* `start` - (Required, int) Start time.
* `end` - (Required, int) End time.

### Relative Time Range

* `amount` - (Required, int) The time number.
* `unit` - (Required) The time unit.  Valid values are `hour`, `day`, `week`, `month`, or `year`.

### To Now Time Range

From some time in the past until now.

* `unit` - (Required) The time unit.  Valid values are `login`, `epoch`, `day`, `week`, `month`, or `year`.

### Filters

Filtering parameters.  This block can be specified multiple times to add more filters.

* `name` - (Required) Param name to filter on.
* `operator` - (Optional) Operator between the name and value params (default: `=`).
* `value` - (Required) Param value for the filter.

## Attributes Reference

* `page_token` - The next page token returned.
* `total` - (int) Total number of alerts returned.
* `listing` - Alert listing, as defined [below](#listing).

### Listing

The alert information.

* `alert_id` - Alert ID.
* `status` - Alert status.
* `first_seen` - (int) First seen.
* `last_seen` - (int) Last seen.
* `alert_time` - (int) Alert time.
* `event_occurred` - (int) Event occurred.
* `triggered_by` - Triggered by.
* `alert_count` - (int) Alert count.
