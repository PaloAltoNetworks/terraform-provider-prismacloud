---
page_title: "Prisma Cloud: prismacloud_compliance_standard_requirement_section"
---

# prismacloud_compliance_standard_requirement_section

Retrieve information on a compliance standard requirement section.

## Example Usage

```hcl
data "prismacloud_compliance_standard_requirement_section" "example" {
    csr_id = "11111111-2222-3333-4444-555555555555"
    section_id = "1.007"
}
```

## Argument Reference

You must specify at least one of the following:

* `csrs_id` - Compliance standard requirement section ID
* `section_id` - Compliance section ID

Additional arguments:

* `csr_id` - (Required) Compliance standard ID.

## Attribute Reference

* `description` - Description
* `created_by` - Created by
* `created_on` - (int) Created on
* `last_modified_by` - Last modified by
* `last_modified_on` - (int) Last modified on
* `system_default` - (bool) System default
* `policies_assigned_count` - (int) Number of assigned policies
* `standard_name` - Compliance standard name
* `requirement_name` - Compliance requirement name
* `label` - Section label
* `view_order` - (int) View order
* `associated_policy_ids` - List of associated policy IDs
