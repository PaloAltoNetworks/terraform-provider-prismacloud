---
page_title: "Prisma Cloud: prismacloud_trusted_alert_ip"
---

# prismacloud_trusted_alert_ip

Manage an trusted alert ip.

## Example Usage

```hcl
resource "prismacloud_trusted_alert_ip" "example" {
    name = "My new group"
    cidrs {
        cidr = "1.1.1.1/32"
        description = "ip address description"
    }
}
```

## Argument Reference

* `name` - Name of the trusted alert ip.
* `cidrs` - CIDRs, as defined [below](#CIDR).

## Attribute Reference

* `uuid` - UUID.
* `cidr_count` - CIDR count.

### CIDR

* `cidr` - (string) Ip address.
* `description` - Description.
* `uuid` - UUID for cidr.
* `created_on` - (int) Created on.

## Import

Resources can be imported using the uuid:

```
$ terraform import prismacloud_trusted_alert_ip.example 11111111-2222-3333-4444-555555555555
```
