---
page_title: "Prisma Cloud: prismacloud_dataprofiles"
---

# prismacloud_dataprofiles

Retrieve a list of data profiles.

## Example Usage

```hcl
data "prismacloud_dataprofiles" "example" {}
```

## Attribute Reference

* `total` - (int) Total number of data profiles.
* `listing` - List of data profiles returned, as defined [below](#listing).

### Listing

Each data profile has the following attributes:

* `profile_id` - Profile ID.
* `name` - Profile Name.
* `profile_status` - Profile status (active or disabled).
* `updated_at` - Updated at (unix time).
* `updated_by` - Updated by.
* `created_by` - Created by.
* `type` - Type (basic or advance).
