---
layout: "prismacloud"
page_title: "Prisma Cloud: prismacloud_compliance_standard_requirements"
description: |-
  Retrieve a list of compliance standard requirements.
---

# prismacloud_compliance_standard_requirements

Retrieve a list of compliance standard requirements.

## Example Usage

```hcl
data "prismacloud_compliance_standard_requirements" "example" {
    cs_id = "11111111-2222-3333-4444-555555555555"
}
```

## Argument Reference

* `cs_id` - (Required) Compliance standard ID.

## Attribute Reference

* `total` - (int) Total number of system supported and custom compliance standard requirements
* `listing` - List of compliance requirements, as defined [below](#listing)

### Listing

Each requirement has the following attributes:

* `csr_id` - Compliance standard requirement ID
* `name` - Compliance standard requirement name
* `description` - Description
* `created_by` - Created by
* `created_on` - (int) Created on
* `last_modified_by` - Last modified by
* `last_modified_on` - (int) Last modified on
* `system_default` - (bool) System default
* `policies_assigned_count` - (int) Number of assigned policies
* `standard_name` - Compliance standard name
* `requirement_id` - Compliance requirement number
* `view_order` - (int) View order
