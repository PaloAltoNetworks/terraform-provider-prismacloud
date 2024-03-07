---
page_title: "Prisma Cloud: prismacloud_account_group"
---

# prismacloud_account_group

Retrieves account group information.

## Example Usage

```hcl
data "prismacloud_account_group" "example" {
    name = "myGroup"
}
```

## Argument Reference

One of the following must be specified:

* `name` - Name of the account group.
* `group_id` - Account group ID.

## Attribute Reference

* `description` - Description.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.
* `account_ids` - List of cloud account IDs.
* `child_group_ids` - List of child account group IDs.