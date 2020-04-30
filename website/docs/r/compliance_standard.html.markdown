---
layout: "prismacloud"
page_title: "Prisma Cloud: prismacloud_compliance_standard"
description: |-
  Manage a compliance standard.
---

# prismacloud_compliance_standard

Manage a compliance standard.

## Example Usage

```hcl
resource "prismacloud_compliance_standard" "example" {
    name = "Foo"
    description = "Made by Terraform"
}
```

## Argument Reference

* `name` - (Required) Compliance standard name
* `description` - Description

## Attribute Reference

* `cs_id` - Compliance standard ID
* `created_by` - Created by
* `created_on` - (int) Created on
* `last_modified_by` - Last modified by
* `last_modified_on` - (int) Last modified on
* `system_default` - (bool) System default
* `policies_assigned_count` - (int) Number of assigned policies
* `cloud_types` - List of cloud types (determined based on policies assigned)

## Import

Resources can be imported using the compliance standard ID:

```
$ terraform import prismacloud_compliance_standard.example 11111111-2222-3333-4444-555555555555
```
