---
page_title: "Prisma Cloud: prismacloud_user_roles"
---

# prismacloud_user_roles

Retrieve a list of user roles.

## Example Usage

```hcl
data "prismacloud_user_roles" "example" {}
```

## Attribute Reference

* `total` - (int) Total number of user roles.
* `listing` - List of user roles returned, as defined [below](#listing).

### Listing

Each user role has the following attributes:

* `role_id` - Role Id
* `name` - Name of the role.
* `role_type` - User role type.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.
* `associated_users` - List of associated application users which cannot exist in the system without the user role.
* `restrict_dismissal_access` - (bool) Restrict dismissal access.
* `account_groups` - List of associated account groups, as defined [below](#account-groups).
* `additional_attributes` - An Additional attributes spec, as defined [below](#additional-attributes).

#### Account Groups

Each account group has the following attributes.

* `group_id` - The group ID.
* `name` - Group name.

#### Additional Attributes

* `only_allow_ci_access` - (bool) - Allows only CI Access.
* `only_allow_read_access` - (bool) - Allow read only access.
* `has_defender_permissions`- (bool) - Has defender Permissions.
* `only_allow_compute_access`- (bool) - Access to only Compute tab and Access keys.