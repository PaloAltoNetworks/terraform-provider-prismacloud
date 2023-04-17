---
page_title: "Prisma Cloud: prismacloud_trusted_alert_ip"
---

# prismacloud_trusted_alert_ip

Retrieves trusted alert ip information

## Example Usage

```hcl
data "prismacloud_trusted_alert_ip" "example" {
    name = "trusted alert ip name"
}
```

## Argument Reference

One of the following must be specified:

* `name` - Name of the trusted alert ip.
* `uuid` - UUID.

## Attribute Reference

* `cidrs` - CIDRs, as defined [below](#CIDRs).
* `cidr_count` - CIDR count.


### CIDRs

* `cidr` - (int) Last modified timestamp.
* `description` - Description.
* `uuid` - List of cloud account IDs.
* `created_on` - List of child account group IDs.