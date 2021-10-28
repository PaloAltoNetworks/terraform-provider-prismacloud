---
page_title: "Prisma Cloud: prismacloud_datapattern"
---

# prismacloud_datapattern

Retrieve information on a specific data pattern.

## Example Usage

```hcl
data "prismacloud_datapattern" "example" {
    name = "My Data Pattern"
}
```

## Argument Reference

You must specify at least one of the following:

* `pattern_id` - Pattern ID.
* `name` - Pattern name.

## Attribute Reference

* `description` - Pattern description.
* `mode` - Pattern mode (predefined or custom).
* `detection_technique` - Detection technique.
* `entity` - Entity value.
* `grammar` - Grammar value.
* `parent_id` - Parent ID for cloned data pattern.
* `proximity_keywords` - List of proximity keywords.
* `regexes` - List of regexes, as defined [below](#regexes).
* `root_type` - Root type (predefined or custom) for cloned data pattern.
* `s3_path` - S3 Path to the grammar.
* `created_by` - Created by.
* `updated_by` - Updated by.
* `updated_at` - (int) Last updated at.
* `is_editable` - (bool) Is editable.

### Regexes

* `regex` - Regular expression (match criteria for the data you want to find within your assets).
* `weight` - (int) Weight to assign a score to a text entry (pattern match occurs when the score threshold is exceeded).
