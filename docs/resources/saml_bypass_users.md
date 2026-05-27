# prismacloud_saml_bypass_users Resource

The `prismacloud_saml_bypass_users` resource allows you to manage the list of users that are allowed to bypass SAML authentication.

## Example Usage

```hcl
provider "prismacloud" {
  # Provider configuration
}

resource "prismacloud_saml_bypass_users" "example" {
  usernames = ["admin", "backup-user", "support-user"]
}
```

## Argument Reference

- `usernames` - (Required) List of usernames that are allowed to bypass SAML authentication.

## Import

The resource can be imported using a simple identifier:

```shell
hcl terraform import prismacloud_saml_bypass_users.example saml_bypass_users
```