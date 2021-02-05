---
page_title: "Prisma Cloud: prismacloud_rql_historic_search"
---

# prismacloud_rql_historic_search

Retrieve a specific historic RQL search.

## Example Usage

```hcl
data "prismacloud_rql_historic_search" "example" {
    name = 
}
```

## Argument Reference

You must specify at least one of the following:

* `search_id` - Historic RQL search ID
* `name` - Historic RQL search name

## Attribute Reference

* `description` - Description
* `search_type` - Search type
* `cloud_type` - Cloud type
* `query` - RQL query
* `saved` - (bool) If this is a saved search
* `time_range` - The RQL time range spec, as defined [below](#time-range)

### Time Range

Only one of these will be defined:

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
