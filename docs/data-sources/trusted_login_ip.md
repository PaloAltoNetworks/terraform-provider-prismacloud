---
page_title: "Prisma Cloud: prismacloud_trusted_login_ips"
---

# prismacloud_trusted_login_ip

Retrieves list of CIDRs that are in allow list for login access, for the specified login IP allow list ID.

## Example Usage

```hcl
data "prismacloud_trusted_login_ip" "example" {
  trusted_login_ip_id = "Id"
}
```

## Argument Reference

One of the following must be specified:

* `name` - Name of the trusted login ip Allow List.
* `trusted_login_ip_id` - Trusted login ip allow List ID.

## Attribute Reference

* `name` - Name of the list of CIDR blocks that are in allow list for access.
* `trusted_login_ip_id` - Login IP allow list ID
* `last_modified_ts` - Timestamp for last modification of CIDR block list.
* `cidr` - List of CIDR blocks (IP addresses) from which access is allowed when Login IP Allow List is enabled.
* `description` - Description
