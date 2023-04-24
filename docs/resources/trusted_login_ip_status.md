---
page_title: "Prisma Cloud: prismacloud_trusted_login_ip_status"
---

# prismacloud_trusted_login_ip_status

Manage a Trusted Login IP Status.

## Example Usage

```hcl
resource "prismacloud_trusted_login_ip_status" "example" {
  enabled = true
}
```

## Argument Reference

* `enabled` - (Required, bool) Enable or disable the login IP allow list.

## Import

Resources can be imported using the trusted_login_ip_status ID:

```
$ terraform import trusted_login_ip_status.example login_ip_status
```
