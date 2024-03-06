---
page_title: "Prisma Cloud: prismacloud_collection"
---

# prismacloud_collection

Retrieves information about a specific collection.

## Example Usage

```hcl
data "prismacloud_collection" "example" {
    id = "collection_id"
}
```

## Argument Reference

You must specify:

* `id` - ID of the collection.

## Attribute Reference

* `name` - The name of the collection.
* `description` - Description.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.
* `created_by` - Created by.
* `created_ts` - The timestamp when the collection was created.
* `asset_groups` - List of asset groups contained within the collection as defined [below](#asset_groups)

### Asset Groups

* `account_group_ids` - A list of account group IDs associated with the collection.
* `account_ids` - A list of cloud account IDs associated with the collection.
* `repository_ids` - A list of repository IDs associated with the collection.
