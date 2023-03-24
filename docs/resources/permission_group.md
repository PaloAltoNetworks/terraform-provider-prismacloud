---
page_title: "Prisma Cloud: prismacloud_permission_group"
---

# prismacloud_permission_group

Manage a permission group.

## Example Usage

```hcl
resource "prismacloud_permission_group" "example" {
  name = "test permission group"
  description = "Made by Terraform"
  features{
    feature_name= "settingsAuditLogs"
    operations {
      read= true
    }
  }
}
```

## Argument Reference

* `name` - (Required) Name of the permission group.
* `description` - (Optional) Description.
* `permission_group_type` - (Optional) Permission Group type.  Valid values are `Default`, `Custom` or `Internal`.
* `accept_account_groups` - (Optional, bool) Accept account groups.
* `accept_resource_lists` - (Optional, bool) Accept resource lists.
* `accept_code_repositories` - (Optional, bool) Accept code repositories.
* `custom` - (Optional, bool) Boolean value signifying whether this is a custom (i.e. user-defined) permission group.
* `features` - (Required) Collection of permitted features associated with the role, as defined [below](#features).


### Features

* `feature_name` - (Required) Prisma Cloud Feature Name.
* `operations` - (Required) A mapping of operations and a boolean value representing whether the privilege to perform the operation needs to be granted, as defined [below](#operations).

#### Operations
* `create`- (Optional, bool) Create operation.
* `read`- (Optional, bool) Read operation.
* `update`- (Optional, bool) Update operation.
* `delete`- (Optional, bool) Delete operation.


## Attribute Reference

* `id` - Permission group id.
* `last_modified_by` - Last modified by
* `last_modified_ts` - (int) Last modified timestamp.
* `associated_roles` - List of associated user roles which cannot exist in the system without the permission group.



