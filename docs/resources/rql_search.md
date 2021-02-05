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
* `time_range` - (Required) The time range spec, as defined below.
* `limit` - (int) Limit rules (default: `10`).

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
