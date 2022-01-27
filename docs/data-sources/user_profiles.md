---
page_title: "Prisma Cloud: prismacloud_user_profiles"
---

# prismacloud_user_profiles

Retrieve a list of user profiles.

## Example Usage

```hcl
data "prismacloud_user_profiles" "example" {}
```

## Attribute Reference

* `total` - (int) Total number of user profiles.
* `listing` - List of user profiles returned, as defined [below](#listing).

### Listing

Each user profile has the following attributes:

* `profile_id` - Profile ID (email or username).
* `account_type` - Account Type (USER_ACCOUNT or SERVICE_ACCOUNT).
* `username` - User email or service account name.
* `display_name` - Display name.
* `default_role_id` - Default User Role ID.
* `role_ids` - List of Role IDs.
* `time_zone` - Time zone (e.g. America/Los_Angeles).
* `enabled` - (bool) Enabled.
* `last_login_ts` - (int) Last login timestamp.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.