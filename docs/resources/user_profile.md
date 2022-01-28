---
page_title: "Prisma Cloud: prismacloud_user_profile"
---

# prismacloud_user_profile

Manage a user profile.

## Example Usage

```hcl
resource "prismacloud_user_profile" "example" {
    first_name = "Firstname"
    last_name = "Lastname"
    email = "user@email.com"
    username = "user@email.com"
    role_ids = [
        "11111111-2222-3333-4444-555555555555",
        "12345678-2222-3333-4444-555555555555"
    ]
    time_zone = "America/Los_Angeles"
    default_role_id = "11111111-2222-3333-4444-555555555555"
}

/*
You can also create user profiles from a CSV file using native Terraform
HCL and looping.  Assume you have a CSV file (user_profiles.csv) of user profiles that looks like this (with
"||" separating user role IDs from each other):

"first_name","last_name","email","role_ids","access_keys_allowed","time_zone","default_role_id"
"FirstName1","LastName1","test1@email.com","11111111-2222-3333-4444-555555555555||12345678-2222-3333-4444-555555555555",true,"Asia/Calcutta","12345678-2222-3333-4444-555555555555"
"FirstName2","LastName2","test2@email.com","11111111-2222-3333-4444-555555555555||12345678-2222-3333-4444-555555555555",true,"America/Los_Angeles","12345678-2222-3333-4444-555555555555"
Here's how you would do this:
*/
locals {
    user_profiles = csvdecode(file("user_profiles.csv"))
}

// Now specify the user profile resource with a loop like this:
resource "prismacloud_user_profile" "example" {
    for_each = { for inst in local.user_profiles : inst.email => inst }

    first_name = each.value.first_name
    last_name = each.value.last_name
    email = each.value.email
    username = each.value.email
    role_ids = split("||", each.value.role_ids)
    access_keys_allowed = each.value.access_keys_allowed
    time_zone = each.value.time_zone
    default_role_id = each.value.default_role_id
}
```

## Argument Reference

* `account_type` - (Optional) Account Type. Valid values are `USER_ACCOUNT`, or `SERVICE_ACCOUNT`. (default: `USER_ACCOUNT`)
* `username` - (Required) User email or service account name.
* `first_name` - (Required if `account_type` is `USER_ACCOUNT`) First name.
* `last_name` - (Required if `account_type` is `USER_ACCOUNT`) Last name.
* `email` - (Required if `account_type` is `USER_ACCOUNT`) Email ID.
* `access_key_expiration` - (Optional, int) Access key expiration timestamp in milliseconds for `SERVICE_ACCOUNT`.
* `access_key_name` - (Required if `account_type` is `SERVICE_ACCOUNT`) Access key name.
* `enable_key_expiration` - (Optional, bool) Enable access key expiration. (default: `false`)
* `role_ids` - (Required) List of Role IDs. (default: `false`)
* `default_role_id` - (Required) Default Role ID, must be present in `role_ids`.
* `time_zone` - (Required) Time zone (e.g. America/Los_Angeles).
* `enabled` - (Optional, bool) Is account enabled. (default: `true`)
* `access_keys_allowed` - (Optional, bool) Access keys allowed. (For `USER_ACCOUNT` default value is `true` if `role_ids` contain `System Admin` role)

## Attribute Reference

* `profile_id` - Profile ID (email or username).
* `display_name` - Display name.
* `access_key_id` - Access key ID generated for `SERVICE_ACCOUNT`.
* `secret_key` - Access key secret generated for `SERVICE_ACCOUNT`.
* `last_login_ts` - (int) Last login timestamp.
* `last_modified_by` - Last modified by
* `last_modified_ts` - (int) Last modified timestamp.
* `access_keys_count` - (int) Access keys count.
* `roles` - List of User Profile Roles Details. Each item has role information, as defined [below](#roles).

### Roles
* `role_id` - User Role ID.
* `name` - User Role Name.
* `only_allow_ci_access` - (bool) Allow only CI Access for Build and Deploy security roles.
* `only_allow_compute_access` - (bool) Allow only Compute Access for reduced system admin roles.
* `only_allow_read_access` - (bool) Allow only read access.
* `role_type` - User Role Type.


## Import

Resources can be imported using the username/email:

```
$ terraform import prismacloud_user_profile.example user@email.com
```