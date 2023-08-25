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

* `id` - (string) ID of resource list.
* `name` - (string) Name of resouce list.
* `description` - (string) Description of the resource list.
* `resource_list_type` - (string) Type of resource list.
* `last_modified_by` - (string) Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.

