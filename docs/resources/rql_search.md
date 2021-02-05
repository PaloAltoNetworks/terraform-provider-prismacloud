---
page_title: "Prisma Cloud: prismacloud_rql_search"
---

# prismacloud_rql_search

Manage a RQL search on the Prisma Cloud Platform.

NOTE:  Prisma Cloud does not currently support deleting RQL searches, so
`terraform destroy` is a noop.

## Example Usage

```hcl
resource "prismacloud_rql_search" "example" {
    search_type = "config"
    query = "config from cloud.resource where api.name = 'aws-ec2-describe-instances'"
    time_range {
        relative {
            unit = "hour"
            amount = 24
        }
    }
}
```

## Argument Reference

The following arguments are supported:

* `search_type` - (Required) The search type.  Valid values are `config`
  (default) and `event`.
* `query` - (Required) The RQL query.
* `limit` - (int) Limit rules (default: `10`).
* `time_range` - (Required) The RQL time range spec, as defined [below](#time-range).

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

The following attributes are supported:

* `search_id` - The search ID returned from a successful RQL query.
* `group_by` - (list) Group by.
* `cloud_type` - The cloud type.
* `name` - Name.
* `description` - Description.
* `config_data` - (for `search_type="config"`, list) List of config_data specs,
  as defined below.
* `event_data` - (For `search_type="event"`, list) List of event_data specs,
  as defined below.

`config_data` supports the following attributes:

* `state_id` - The state ID.
* `name` - Name.
* `url` - The URL.

`event_data` supports the following attributes:

* `account` - Account.
* `region_id` - (int) Region ID.
* `region_api_identifier` - Region API identifier.
