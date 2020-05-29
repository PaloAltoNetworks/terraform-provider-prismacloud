---
layout: "prismacloud"
page_title: "Prisma Cloud: prismacloud_integrations"
description: |-
  Retrieves an integration listing.
---

# prismacloud_integrations

Retrieves an integration listing.

## Example Usage

```hcl
data "prismacloud_integrations" "example" {}
```

## Attribute Reference

* `total` - (int) Total number of all integrations.
* `listing` - List of integrations, as defined [below](#listing).

### Listing

* `name` - Name of the integration.
* `integration_id` - Integration ID.
* `description` - Description.
* `integration_type` - Integration type.
* `enabled` - (bool) Enabled.
* `status` - Status.
* `valid` - (bool) Valid.
