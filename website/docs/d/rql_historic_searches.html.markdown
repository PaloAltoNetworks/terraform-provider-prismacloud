---
layout: "prismacloud"
page_title: "Prisma Cloud: prismacloud_rql_historic_searches"
description: |-
  Retrieve a list of historic RQL searches.
---

# prismacloud_rql_historic_searches

Retrieve a list of historic RQL searches.

## Example Usage

```hcl
data "prismacloud_rql_historic_searches" "example" {}
```

## Argument Reference

* `filter` - (Optional) Filter for historic RQL searches.  Valid values are `saved` (default) or `recent`.
* `limit` - (Optional, int) Max number of historic RQL searches to return (default: `1000`).

## Attribute Reference

* `total` - (int) Total number of RQL historic searches.
* `listing` - List of historic RQL searches, as defined [below](#listing).

### Listing

Each result in the listing has the following attributes:

* `created_by` - Created by
* `last_modified_by` - Last modified by
* `search_id` - Historic RQL search ID
* `name` - Name
* `search_type` - Search type
* `saved` - (bool) If this is a saved search
