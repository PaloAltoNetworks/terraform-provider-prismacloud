---
page_title: "Prisma Cloud: prismacloud_dataprofile"
---

# prismacloud_dataprofile

Retrieve information on a specific data profile.

## Example Usage

```hcl
data "prismacloud_dataprofile" "example" {
    name = "My Data Profile"
}
```

## Argument Reference

You must specify at least one of the following:

* `profile_id` - Profile ID.
* `name` - Profile Name.

## Attribute Reference

* `description` - Profile description.
* `types` - Type (basic or advance).
* `profile_type` - Profile type (custom or system).
* `tenant_id` - Tenant ID.
* `status` - Status (hidden or non_hidden).
* `profile_status` - Profile status (active or disabled).
* `created_by` - Created by.
* `updated_by` - Updated by.
* `created_at` - Created at (unix time).
* `updated_at` - Updated at (unix time).
* `data_patterns_rule_1` - Model for DataProfile Rules, as defined [below](#data-patterns-rule-1).

### Data Patterns Rule 1

* `operator_type` - Pattern operator type.
* `data_pattern_rules` - List of DataPattern Rules. Each item has data-pattern information, as defined [below](#data-pattern-rules).

#### Data Pattern Rules

* `pattern_id` - Pattern ID.
* `name` - Pattern name.
* `detection_technique` - Detection technique.
* `match_type` - Match type.
* `occurrence_operator_type` - Occurrence operator type.
* `occurrence_count` - Occurrence count.
* `confidence_level` - Confidence level.
* `supported_confidence_levels` - List of supported confidence levels.
* `occurrence_high` - High occurrence value.
* `occurrence_low` - Low occurrence value.
