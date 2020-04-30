---
layout: "prismacloud"
page_title: "Prisma Cloud: prismacloud_compliance_standards"
description: |-
  Retrieve a list of compliance standards.
---

# prismacloud_compliance_standards

Retrieve a list of compliance standards.

## Example Usage

```hcl
data "prismacloud_compliance_standards" "example" {}
```

## Attribute Reference

* `standard_count` - (int) Number of system supported and custom compliance standards
* `standards` - List of system supported and custom compliance standards, as defined [below](#standards)

### Standards

Each standard has the following attributes:

* `cs_id` - Compliance standard ID
* `description` - Description
* `created_by` - Created by
* `created_on` - (int) Created on
* `last_modified_by` - Last modified by
* `last_modified_on` - (int) Last modified on
* `system_default` - (bool) System default
* `policies_assigned_count` - (int) Number of assigned policies
* `name` - Compliance standard name
* `cloud_types` - List of cloud types (determined based on policies assigned)
