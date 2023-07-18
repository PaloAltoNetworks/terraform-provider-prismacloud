---
page_title: "Prisma Cloud: prismacloud_saved_search"
---

# prismacloud_saved_search

Manage a saved search on the Prisma Cloud Platform.

## Example Usage

NOTE: In scenarios where a use case necessitates updating the resource `prismacloud_rql_search`, it is advised to first remove the corresponding `prismacloud_rql_search` resources from the Terraform state file.

```hcl
resource "prismacloud_saved_search" "example" {
    name = "Made by Terraform"
    description = "made by terraform"
    search_id = prismacloud_rql_search.x.search_id
    query = prismacloud_rql_search.x.query
    time_range {
        relative {
            unit = prismacloud_rql_search.x.time_range.0.relative.0.unit
            amount = prismacloud_rql_search.x.time_range.0.relative.0.amount
        }
    }
}

resource "prismacloud_rql_search" "x" {
    search_type = "config"
    skip_result = true
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

* `query` - (Required) The RQL query.
* `search_id` - (Required) The search ID.
* `name` - (Required) Name.
* `description` - Description.
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

* `saved` - (bool) This is set to `true` when the saved search is created.

## Import

Resources can be imported using the saved-search ID:

```
$ terraform import prismacloud_saved_search.example 11111111-2222-3333-4444-555555555555
```