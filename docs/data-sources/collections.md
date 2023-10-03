---
page_title: "Prisma Cloud: prismacloud_collections"
---

# prismacloud_collections

Lists collections.

## Example Usage

```hcl
data "prismacloud_collections" "example" {}
```

## Attribute Reference

* `total` - (int) Total number of collections.
* `listing` - List of collections, as defined [below](#listing).

### Listing

* `id` - ID of the collection.
* `name` - The name of the collection.
* `description` - Description.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.
* `created_by` - Created by.
* `created_ts` - The timestamp when the collection was created.
