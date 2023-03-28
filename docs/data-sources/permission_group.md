---
page_title: "Prisma Cloud: prismacloud_permission_group"
---

# prismacloud_permission_group

Retrieve information on a specific permission group.

## Example Usage

```hcl
data "prismacloud_permission_group" "example" {
    name = "My Permission Group"
}
```

## Argument Reference

You must specify at least one of the following:

* `id` - Permission group id
* `name` - Name of the permission group.

## Attribute Reference

* `description` - Description.
* `permission_group_type` - Permission group type.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.
* `associated_roles` - List of associated user roles which cannot exist in the system without the permission group.
* `accept_account_groups` - (bool) Accept account groups.
* `accept_resource_lists` - (bool) Accept resource lists.
* `accept_code_repositories` - (bool) Accept code repositories.
* `custom` - (bool) Boolean value signifying whether this is a custom (i.e. user-defined) permission group.
* `features` - Collection of permitted features associated with the role, as defined [below](#features).

### Features

* `feature_name` - Prisma Cloud Feature Name.
* `operations` - A mapping of operations and a boolean value representing whether the privilege to perform the operation needs to be granted, as defined [below](#operations).

#### Operations
* `create`- (bool) Create operation.
* `read`- (bool) Read operation.
* `update`- (bool) Update operation.
* `delete`- (bool) Delete operation.


