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

## Attribute Reference

* `account_ids` - List of cloud account IDs.
* `group_id` - Account group ID.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.

## Import

Resources can be imported using the group ID:

```
$ terraform import prismacloud_account_group.example 11111111-2222-3333-4444-555555555555
```
