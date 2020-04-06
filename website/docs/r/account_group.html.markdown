---
layout: "prismacloud"
page_title: "Prisma Cloud: prismacloud_account_group"
description: |-
  Manage an account group.
---

# prismacloud_account_group

Manage an account group.

## Example Usage

```hcl
resource "prismacloud_account_group" "example" {
    name = "My new group"
    description = "Made by Terraform"
}
```

## Argument Reference

* `name` - (Required) name of the group.
* `description` - (Optional) Description.
* `account_ids` - (Optional) List of cloud account IDs.

## Attribute Reference

* `group_id` - Account group ID.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.
* `accounts` - Associated cloud accounts spec, as defined [below](#accounts).
* `alert_rules` - Singly associated rules which cannot exist in the system without the account group spec, as defined [below](#alert-rules).

### Accounts

* `account_id` - Associated cloud account ID.
* `name` - Associated cloud account name.
* `account_type` - Associated cloud account type.

### Alert Rules

* `alert_id` - The alert ID.
* `name` - Alert name.

## Import

Resources can be imported using the group ID:

```
$ terraform import prismacloud_account_group.example 11111111-2222-3333-4444-555555555555
```
