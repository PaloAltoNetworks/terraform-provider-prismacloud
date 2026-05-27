# prismacloud_saml_bypass_users Data Source

The `prismacloud_saml_bypass_users` data source retrieves the list of users that are allowed to bypass SAML authentication.

## Example Usage

```hcl
provider "prismacloud" {
  # Provider configuration
}

data "prismacloud_saml_bypass_users" "example" {}

output "bypass_users" {
  value = data.prismacloud_saml_bypass_users.example.usernames
}
```

## Argument Reference

No arguments are required for this data source.

## Attributes Reference

- `usernames` - (Computed) List of usernames that are allowed to bypass SAML authentication.