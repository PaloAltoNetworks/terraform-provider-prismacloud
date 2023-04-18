---
page_title: "Prisma Cloud: prismacloud_trusted_alert_ip"
---

# prismacloud_trusted_alert_ip

Retrieves trusted alert ip information.

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

* `cidrs` - List of CIDRs, as defined [below](#CIDR).
* `cidr_count` - CIDR count.

### CIDR

* `cidr` - (string) Ip address.
* `description` - Description.
* `uuid` - UUID for cidr.
* `created_on` - (int) Created on.