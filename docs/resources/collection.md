---
page_title: "Prisma Cloud: prismacloud_collection"
---

# prismacloud_collection

Manage a collection.

## Example Usage 

```hcl
resource "prismacloud_collection" "example" {
  name        = "test terraform collection"
  description = "Made by terraform"
  asset_groups {
    account_group_ids = ["account_group_ids"]
    account_ids      = ["account_ids"]          //["*"] for including all existing and newly created cloud accounts
    repository_ids   = ["repository_ids"]       //["*"] for including all existing and newly created repositories
  }
}
```

## Argument Reference

* `name` - (Required) Name of the collection.
* `description` - (Optional) Description of the collection.
* `asset_groups` - (Required) List of asset groups contained within the collection as defined [below](#asset_groups)

## Attribute Reference

* `id` - ID of the collection.
* `created_by` - Created by.
* `created_ts` - (int) The timestamp when the collection was created.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.

### Asset Groups

* `account_group_ids` - (Optional) A list of account group IDs associated with the collection.
* `account_ids` - (Optional) A list of cloud account IDs associated with the collection.
* `repository_ids` - (Optional) A list of repository IDs associated with the collection.
