---
page_title: "Prisma Cloud: prismacloud_resource_lists"
---

# prismacloud_resource_lists

Retrieves list of resource lists.

## Example Usage

```hcl
data "prismacloud_resource_lists" "example" {}
```

## Attribute Reference

* `total` - (int) Total number of resource lists.
* `listing` - List of resource lists, as defined [below](#resource-lists).

### Resource Lists

Each resource list has the following attributes.

* `id` - ID of resource list.
* `name` - Name of resource list.
* `description` - Description of the resource list.
* `resource_list_type` - Type of resource list.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - Last modified timestamp.

