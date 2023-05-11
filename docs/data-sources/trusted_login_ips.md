---
page_title: "Prisma Cloud: prismacloud_trusted_login_ips"
---

# prismacloud_trusted_login_ips

List Login IP Allow Lists.

## Example Usage

```hcl
data "prismacloud_trusted_login_ips" "example" {}
```

## Attribute Reference

* `total` - (int) Total number of trusted login ips.
* `listing` - List of trusted login Ips, as defined [below](#listing).

### Listing

* `name` - Name of the list of CIDR blocks that are in allow list for access.
* `trusted_login_ip_id` - Login IP allow list ID
* `last_modified_ts` - Timestamp for last modification of CIDR block list.
* `cidr` - List of CIDR blocks (IP addresses) from which access is allowed when Login IP Allow List is enabled.
* `description` - Description
