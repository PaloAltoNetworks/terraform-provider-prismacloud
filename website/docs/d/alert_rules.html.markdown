---
layout: "prismacloud"
page_title: "Prisma Cloud: prismacloud_alert_rules"
description: |-
  Retrieve a list of alert rules.
---

# prismacloud_alert_rules

Retrieve a list of alert rules.

## Example Usage

```hcl
data "prismacloud_alert_rules" "example" {}
```

## Attribute Reference

* `total` - (int) Total number of alert rules.
* `listing` - List of alerts returned, as defined [below](#listing).

### Listing

Each alert rule has the following attributes:

* `policy_scan_config_id` - Policy scan config ID
* `name` - Rule/Scan name
* `description` - Description
* `enabled` - (bool) Rule/Scan is enabled or not
* `scan_all` - (bool) Scan all policies
* `policies` - List of specific policies to scan
* `owner` - Customer
* `open_alerts_count` - (int) Open alerts count
* `read_only` - (bool) Model is read-only
* `deleted` - (bool) Deleted
