---
page_title: "Prisma Cloud: prismacloud_datapattern"
---

# prismacloud_datapattern

Manage a data pattern.

## Example Usage

```hcl
resource "prismacloud_datapattern" "example" {
    name = "test_datapattern"
    description = "Made by Terraform"
    proximity_keywords = ["terraform", "prisma"]
    regexes {
        regex = "prisma"
        weight = 2
    }
}
```

## Argument Reference

* `name` - (Required) Pattern name.
* `description` - Pattern description.
* `detection_technique` - Detection technique (default: `regex`).
* `proximity_keywords` - List of proximity keywords.
* `regexes` - (Required) List of regexes, as defined [below](#regexes).

### Regexes

* `regex` - (Required) Regular expression (match criteria for the data you want to find within your assets).
* `weight` - (int) Weight to assign a score to a text entry (pattern match occurs when the score threshold is exceeded). Default: `1`.

## Attribute Reference

* `pattern_id` - Pattern ID.
* `mode` - Pattern mode (predefined or custom).
* `entity` - Entity value.
* `grammar` - Grammar value.
* `parent_id` - Parent ID for cloned data pattern.
* `root_type` - Root type (predefined or custom) for cloned data pattern.
* `s3_path` - S3 Path to the grammar.
* `created_by` - Created by.
* `updated_by` - Updated by.
* `updated_at` - (int) Last updated at.
* `is_editable` - (bool) Is editable.

## Import

Resources can be imported using the data pattern ID:

```
$ terraform import prismacloud_datapattern.example 111111111111111111111111
```
