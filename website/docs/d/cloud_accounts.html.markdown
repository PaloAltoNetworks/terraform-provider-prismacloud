---
layout: "prismacloud"
page_title: "Prisma Cloud: prismacloud_cloud_accounts"
description: |-
  Retrieve a list of cloud accounts onboarded onto the Prisma Cloud platform.
---

# prismacloud_cloud_accounts

Retrieve a list of cloud accounts onboarded onto the Prisma Cloud platform.

## Example Usage

```hcl
data "prismacloud_cloud_accounts" "example" {}
```

## Attribute Reference

* `accounts` - List of cloud accounts, defined [below](#accounts).

### Accounts

Each account has the following attributes:

* `name` - Account name
* `cloud_type` - Cloud type
* `account_id` - Account ID.
