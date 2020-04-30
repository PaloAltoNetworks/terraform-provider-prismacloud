---
layout: "prismacloud"
page_title: "Prisma Cloud: prismacloud_compliance_standard_requirement"
description: |-
  Retrieve info on a compliance standard requirement.
---

# prismacloud_compliance_standard_requirement

Retrieve info on a compliance standard requirement.

## Example Usage

```hcl
data "prismacloud_compliance_standard_requirement" "example" {
    cs_id = "11111111-2222-3333-4444-555555555555"
    name = "My requirement name"
}
```

## Argument Reference

You must specify at least one of the following:

* `csr_id` - Compliance standard requirement ID
* `name` - Compliance standard requirement name

Additional arguments:

* `cs_id` - (Required) Compliance standard ID.

## Attribute Reference

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
