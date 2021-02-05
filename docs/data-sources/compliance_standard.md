---
page_title: "Prisma Cloud: prismacloud_compliance_standard"
---

# prismacloud_compliance_standard

Retrieve info for a compliance standard.

## Example Usage

```hcl
data "prismacloud_compliance_standard" "example" {
    name = "Foo"
}
```

## Argument Reference

You must specify at least one of the following:

* `cs_id` - Compliance standard ID
* `name` - Compliance standard name

## Attribute Reference

* `description` - Description
* `created_by` - Created by
* `created_on` - (int) Created on
* `last_modified_by` - Last modified by
* `last_modified_on` - (int) Last modified on
* `system_default` - (bool) System default
* `policies_assigned_count` - (int) Number of assigned policies
* `cloud_types` - List of cloud types (determined based on policies assigned)
