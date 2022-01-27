---
page_title: "Prisma Cloud: prismacloud_user_role"
---

# prismacloud_user_role

Retrieve information on a specific user role.

## Example Usage

```hcl
data "prismacloud_user_role" "example" {
    name = "My Role"
}
```

## Argument Reference

You must specify at least one of the following:

* `role_id` - Role Id
* `name` - Name of the role.

## Attribute Reference

* `description` - Description.
* `role_type` - User role type.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.
* `account_group_ids` - List of accessible account group IDs.
* `resource_list_ids` - List of resource list IDs.
* `code_repository_ids` - List of code repository IDs.
* `restrict_dismissal_access` - (bool) Restrict dismissal access.
* `associated_users` - List of associated application users which cannot exist in the system without the user role.
* `additional_attributes` - An Additional attributes spec, as defined [below](#additional-attributes).

## Additional Attributes

* `only_allow_ci_access` - (bool) - Allows only CI Access.
* `only_allow_read_access` - (bool) - Allow read only access.
* `has_defender_permissions`- (bool) - Has defender Permissions.
* `only_allow_compute_access`- (bool) - Access to only Compute tab and Access keys.

