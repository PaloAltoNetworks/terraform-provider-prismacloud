---
layout: "prismacloud"
page_title: "Prisma Cloud: prismacloud_compliance_standard_requirement_section"
description: |-
  Manage a compliance standard requirement section.
---

# prismacloud_compliance_standard_requirement_section

Manage a compliance standard requirement section.

## Example Usage

```hcl
resource "prismacloud_compliance_standard_requirement_section" "example" {
    csr_id = prismacloud_compliance_standard_requirement.y.csr_id
    section_id = "Section 1"
    description = "Section description"
}

resource "prismacloud_compliance_standard_requirement" "y" {
    cs_id = prismacloud_compliance_standard.x.cs_id
    name = "My first req"
    description = "Also made by Terraform"
    requirement_id = "1.007"
}

resource "prismacloud_compliance_standard" "x" {
    name = "My Terraform Standard"
    description = "Made by Terraform"
}
```

## Argument Reference

* `csr_id` - (Required) Compliance standard ID.
* `section_id` - (Required) Compliance section ID
* `description` - Description
* `view_order` - (Computed, int) View order

## Attribute Reference

* `csrs_id` - Compliance standard requirement section ID
* `created_by` - Created by
* `created_on` - (int) Created on
* `last_modified_by` - Last modified by
* `last_modified_on` - (int) Last modified on
* `system_default` - (bool) System default
* `policies_assigned_count` - (int) Number of assigned policies
* `standard_name` - Compliance standard name
* `requirement_name` - Compliance requirement name
* `label` - Section label
* `associated_policy_ids` - List of associated policy IDs

## Import

Resources can be imported using the csr_id and the csrs_id:

```
$ terraform import prismacloud_compliance_standard_requirement_section.example 11111111-2222-3333-4444-555555555555:11111111-2222-3333-4444-555555555555
```
