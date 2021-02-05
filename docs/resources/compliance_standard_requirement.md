---
page_title: "Prisma Cloud: prismacloud_compliance_standard_requirement"
---

# prismacloud_compliance_standard_requirement

Manage a compliance standard requirement.

## Example Usage

```hcl
resource "prismacloud_compliance_standard_requirement" "example" {
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

* `cs_id` - (Required) Compliance standard ID.
* `name` - (Required) Compliance standard requirement name
* `description` - Description
* `requirement_id` - (Required) Compliance requirement number
* `view_order` - (Computed, int) View order

## Attribute Reference

* `csr_id` - Compliance standard requirement ID
* `created_by` - Created by
* `created_on` - (int) Created on
* `last_modified_by` - Last modified by
* `last_modified_on` - (int) Last modified on
* `system_default` - (bool) System default
* `policies_assigned_count` - (int) Number of assigned policies
* `standard_name` - Compliance standard name

## Import

Resources can be imported using the cs_id and the csr_id:

```
$ terraform import prismacloud_compliance_standard_requirement.example 11111111-2222-3333-4444-555555555555:11111111-2222-3333-4444-555555555555
```
