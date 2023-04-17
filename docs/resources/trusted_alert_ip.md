---
page_title: "Prisma Cloud: prismacloud_trusted_alert_ip"
---

# prismacloud_trusted_alert_ip

Manage an trusted alert ip.

## Example Usage

```hcl
resource "prismacloud_trusted_alert_ip" "example" {
    name = "My new group"
    description = "Made by Terraform"
    account_ids = [
    prismacloud_cloud_account1,
    ]
}
```

## Argument Reference

One of the following must be specified:

* `name` - Name of the trusted alert ip.
* `cidrs` - CIDRs, as defined [below](#CIDRs).



## Attribute Reference

* `uuid` - UUID.
* `cidr_count` - CIDR count.


### CIDRs

* `cidr` - (int) Last modified timestamp.
* `description` - Description.
* `uuid` - UUID.
* `created_on` - Created_on.

## Import

Resources can be imported using the uuid:

```
$ terraform import prismacloud_trusted_alert_ip.example 11111111-2222-3333-4444-555555555555
```
