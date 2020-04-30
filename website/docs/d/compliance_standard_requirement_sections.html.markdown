---
layout: "prismacloud"
page_title: "Prisma Cloud: prismacloud_compliance_standard_requirement_sections"
description: |-
  Retrieve a list of compliance standard requirement sections.
---

# prismacloud_compliance_standard_requirement_sections

Retrieve a list of compliance standard requirement sections.

## Example Usage

```hcl
data "prismacloud_compliance_standard_requirement_sections" "example" {
    csr_id = "11111111-2222-3333-4444-555555555555"
}
```

## Argument Reference

* `csr_id` - (Required) Compliance standard ID.

## Attribute Reference

* `section_count` - (int) Number of system supported and custom compliance standard requirement sections
* `sections` - List of compliance requirements, as defined [below](#sections)

### Sections

Each requirement has the following attributes:

* `csrs_id` - Compliance standard requirement section ID
* `description` - Description
* `created_by` - Created by
* `created_on` - (int) Created on
* `last_modified_by` - Last modified by
* `last_modified_on` - (int) Last modified on
* `system_default` - (bool) System default
* `policies_assigned_count` - (int) Number of assigned policies
* `standard_name` - Compliance standard name
* `requirement_name` - Compliance requirement name
* `section_id` - Compliance section ID
* `label` - Section label
* `view_order` - (int) View order
* `associated_policy_ids` - List of associated policy IDs
