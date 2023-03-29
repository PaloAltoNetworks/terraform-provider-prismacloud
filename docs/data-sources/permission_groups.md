---
page_title: "Prisma Cloud: prismacloud_permission_groups"
---

# prismacloud_permission_groups

Retrieve a list of permission groups.

## Example Usage

```hcl
data "prismacloud_permission_groups" "example" {}
```

## Attribute Reference

* `total` - (int) Total number of permission groups.
* `listing` - List of permission groups returned, as defined [below](#listing).

### Listing

Each permission group has the following attributes:

* `id` - Permission group id.
* `name` - Name of the permission group.
* `description` - Description of permission group.
* `permission_group_type` - Permission group type.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - Last modified timestamp.
* `accept_account_groups` - (bool) Accept account groups.
* `accept_resource_lists` - (bool) Accept resource lists.
* `accept_code_repositories` - (bool) Accept code repositories.
* `custom` - (bool) Boolean value signifying whether this is a custom (i.e. user-defined) permission group.