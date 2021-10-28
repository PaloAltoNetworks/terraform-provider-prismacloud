---
page_title: "Prisma Cloud: prismacloud_account_group"
---

# prismacloud_account_group

Manage an account group.

## Example Usage

```hcl
resource "prismacloud_account_group" "example" {
    name = "My new group"
    description = "Made by Terraform"
    account_ids = [
    prismacloud_cloud_account1,
    ]
}
```

## Argument Reference

* `name` - (Required) name of the group.
* `description` - (Optional) Description.
* `account_ids` - (Optional) List of cloud account IDs.
* `child_group_ids` - (Optional) List of child account group IDs.

## Attribute Reference

* `group_id` - Account group ID.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.

## Import

Resources can be imported using the group ID:

```
$ terraform import prismacloud_account_group.example 11111111-2222-3333-4444-555555555555
```
