---
page_title: "Prisma Cloud: prismacloud_org_cloud_accounts"
---

# prismacloud_org_cloud_accounts

Retrieve a list of cloud accounts onboarded onto the Prisma Cloud platform.

## Example Usage

```hcl
data "prismacloud_org_cloud_accounts" "example" {}
```

## Attribute Reference

* `total` - (int) Total number of cloud accounts.
* `listing` - List of cloud accounts, defined [below](#listing).

### Listing

Each cloud account has the following attributes:

* `name` - Account name
* `cloud_type` - Cloud type
* `account_id` - Account ID.
