---
page_title: "Prisma Cloud: prismacloud_user_profile"
---

# prismacloud_user_profile

Retrieve information on a specific user profile.

## Example Usage

```hcl
data "prismacloud_user_profile" "example" {
    profile_id = "user@email.com"
}
```

## Argument Reference

* `profile_id` - (Required) Profile ID (email or username).

## Attribute Reference

* `account_type` - Account Type (USER_ACCOUNT or SERVICE_ACCOUNT).
* `username` - User email or service account name.
* `first_name` - First name.
* `last_name` - Last name.
* `display_name` - Display name.
* `email` - Email ID.
* `access_keys_allowed` - (bool) Access keys allowed.
* `default_role_id` - Default User Role ID.
* `role_ids` - List of Role IDs.
* `time_zone` - Time zone (e.g. America/Los_Angeles).
* `enabled` - (bool) Enabled.
* `last_login_ts` - (int) Last login timestamp.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.
* `access_keys_count` - (int) Access key count.
* `roles` - List of User Profile Roles Details. Each item has role information, as defined [below](#roles).

### Roles
* `role_id` - User Role ID.
* `name` - User Role Name.
* `only_allow_ci_access` - (bool) Allow only CI Access for Build and Deploy security roles.
* `only_allow_compute_access` - (bool) Allow only Compute Access for reduced system admin roles.
* `only_allow_read_access` - (bool) Allow only read access.
* `role_type` - User Role Type.

