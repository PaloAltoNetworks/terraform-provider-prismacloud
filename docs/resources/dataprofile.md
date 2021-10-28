---
page_title: "Prisma Cloud: prismacloud_dataprofile"
---

# prismacloud_dataprofile

Manage a data profile.

## Example Usage

```hcl
resource "prismacloud_dataprofile" "example" {
    name = "test_dataprofile"
    description = "Made by Terraform"
    data_patterns_rule_1 {
        data_pattern_rules {
            name = "Data pattern name"
            confidence_level = "low"
            match_type = "include"
            occurrence_operator_type = "less_than_equal_to"
            occurrence_count = 5
        }
    }
}
```

## Argument Reference

* `name` - (Required) Profile Name.
* `description` - Profile description.
* `types` - Type. Valid values are `basic` (default), or `advance`.
* `profile_type` - Profile Type. Valid values are `custom` (default), or `system`.
* `status` - Status. Valid values are `non_hidden` (default), or `hidden`.
* `profile_status` - Profile status. Valid values are `active` (default), or `disabled`.
* `data_patterns_rule_1` - (Required) Model for DataProfile Rules, as defined [below](#data-patterns-rule-1).

### Data Patterns Rule 1

* `operator_type` - Pattern operator type. Default: `or`.
* `data_pattern_rules` - (Required) List of DataPattern Rules. Each item has data-pattern information, as defined [below](#data-pattern-rules).

#### Data Pattern Rules

* `name` - Pattern name.
* `match_type` - (Required) Match type. Valid values are `include`, or `exclude`.
* `occurrence_operator_type` - (Required) Occurrence operator type. Valid values are `any`, `more_than_equal_to`, `less_than_equal_to`, or `between`.
* `occurrence_count` - (Required if value of `occurrence_operator_type` is `more_than_equal_to` or `less_than_equal_to`) Occurrence count. Value must be a number between `1` and `250`.
* `confidence_level` - (Required) Confidence level.
* `occurrence_high` - (Required if value of `occurrence_operator_type` is `between`) High occurrence value. Value must be a number between `1` and `250`.
* `occurrence_low` - (Required if value of `occurrence_operator_type` is `between`) Low occurrence value. Value must be a number between `1` and `250`.

## Attribute Reference

* `profile_id` - Profile ID.
* `tenant_id` - Tenant ID.
* `created_by` - Created by.
* `updated_by` - Updated by.
* `created_at` - Created at (unix time).
* `updated_at` - Updated at (unix time).

In each `Data Pattern Rules` section, the following attributes are available:

* `pattern_id` - Pattern ID.
* `detection_technique` - Detection technique.
* `supported_confidence_levels` - List of supported confidence levels.

## Import

Resources can be imported using the data profile ID:

```
$ terraform import prismacloud_dataprofile.example 11111111
```
