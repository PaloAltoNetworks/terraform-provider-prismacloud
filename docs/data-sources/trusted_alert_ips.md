---
page_title: "Prisma Cloud: prismacloud_trusted_alert_ips"
---

# prismacloud_trusted_alert_ips

Retrieves a list of trusted alert ips.

## Example Usage

```hcl
data "prismacloud_trusted_alert_ips" "example" {

}
```

## Attribute Reference

* `total` - (int) Total number of trusted alert ips.
* `listing` - List of trusted alert ips returned, as defined [below](#listing).

### Listing

Each trusted alert ip has the following attributes:

* `name` - Name of the trusted alert ip.
* `uuid` - UUID.
* `cidrs` - List of CIDRs, as defined [below](#CIDR).
* `cidr_count` - CIDR count.

### CIDR

* `cidr` - (string) Ip address.
* `description` - Description.
* `uuid` - UUID for cidr.
* `created_on` - (int) Created on.