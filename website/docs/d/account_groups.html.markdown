---
layout: "prismacloud"
page_title: "Prisma Cloud: prismacloud_account_groups"
description: |-
  Lists account groups.
---

# prismacloud_account_groups

Lists account groups.

## Example Usage

```hcl
data "prismacloud_account_groups" "example" {}
```

## Attribute Reference

* `name` - Name of the account group.
* `group_id` - Account group ID.
* `description` - Description.
* `account_ids` - List of cloud account IDs.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.
* `accounts` - Associated cloud accounts spec, as defined [below](#accounts).
* `alert_rules` - Singly associated rules which cannot exist in the system without the account group spec, as defined [below](#alert-rules).

### Accounts

Each account has the following attributes.

* `account_id` - Associated cloud account ID.
* `name` - Associated cloud account name.
* `account_type` - Associated cloud account type.

### Alert Rules

Each alert rule has the following attributes.

* `alert_id` - The alert ID.
* `name` - Alert name.
