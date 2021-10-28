---
page_title: "Prisma Cloud: prismacloud_datapatterns"
---

# prismacloud_datapatterns

Retrieve a list of data patterns.

## Example Usage

```hcl
data "prismacloud_datapatterns" "example" {}
```

## Attribute Reference

* `total` - (int) Total number of data patterns.
* `listing` - List of data patterns returned, as defined [below](#listing).

### Listing

Each data pattern has the following attributes:

* `pattern_id` - Pattern ID.
* `name` - Pattern Name.
* `mode` - Pattern mode (predefined or custom).
* `detection_technique` - Detection technique.
* `updated_at` - (int) Last updated at.
* `updated_by` - Updated by.
* `created_by` - Created by.
* `is_editable` - (bool) Is editable.
