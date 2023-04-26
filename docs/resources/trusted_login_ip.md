---
page_title: "Prisma Cloud: prismacloud_trusted_login_ip"
---

# prismacloud_trusted_login_ip

Manage a Login IP Allow List.

## Example Usage

```hcl
resource "prismacloud_trusted_login_ip" "example" {
  name = "My new List"
  description = "Made by Terraform"
  cidr = [
    "1.1.1.0/24"
  ]
}
```

## Argument Reference

* `name` - (Required) Unique name for CIDR (IP addresses) allow list.
* `description` - (Optional) Description.
* `cidr` - (Required) List of CIDRs to Allow List for login access. You can include from 1 to 10 CIDRs.

## Attribute Reference

* `name` - Name of the list of CIDR blocks that are in allow list for access.
* `trusted_login_ip_id` - Login IP allow list ID
* `last_modified_ts` - (int) Timestamp for last modification of CIDR block list.
* `cidr` - List of CIDR blocks (IP addresses) from which access is allowed when Login IP Allow List is enabled.
* `description` - Description

## Import

Resources can be imported using the trusted_login_ip ID:

```
$ terraform import prismacloud_trusted_login_ip.example 11111111-2222-3333-4444-555555555555
```
