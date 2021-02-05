---
page_title: "Prisma Cloud: prismacloud_policies"
---

# prismacloud_policies

Retrieve a list of policies.

## Example Usage

```hcl
data "prismacloud_policies" "example" {
    filters = {
        "policy.severity": "high",
        "policy.type": "network",
    }
}
```

## Argument Reference

* `filters` - (Optional, map) Filters to limit policies returned.  Filter options can be found [here](https://api.docs.prismacloud.io/reference#get-policies-v2).

## Attribute Reference

* `total` - (int) Total number of policies.
* `listing` - List of policies returned, as defined [below](#listing).

### Listing

Each policy has the following attributes:

* `policy_id` - Policy ID
* `name` - Policy name
* `policy_type` - Policy type
* `system_default` - (bool) If the policy is a system default for Prisma Cloud
* `description` - Description
* `severity` - Severity
* `recommendation` - Remediation recommendation
* `cloud_type` - Cloud type
* `labels` - List of labels
* `enabled` - (bool) Enabled
* `overridden` - (bool) Overridden
* `deleted` - (bool) Deleted
* `open_alerts_count` - (int) Open alerts count
* `policy_mode` - Policy mode
* `remediable` - (bool) Remediable
