# prismacloud_saml_bypass_subset_users Resource

The `prismacloud_saml_bypass_subset_users` resource allows you to manage a subset of users that are allowed to bypass SAML authentication. Unlike `prismacloud_saml_bypass_users` which manages the complete list, this resource only adds or removes the specified users while preserving other users that may be managed outside of Terraform.

## Example Usage

```hcl
provider "prismacloud" {
  # Provider configuration
}

resource "prismacloud_saml_bypass_subset_users" "example" {
  usernames = ["admin", "backup-user"]
}
```

## Argument Reference

- `usernames` - (Required) List of usernames that are allowed to bypass SAML authentication. This resource will add these users to the existing list and remove them when they are no longer in the configuration, but will preserve other users not managed by this resource.

## Import

The resource can be imported using a simple identifier:

```shell
hcl terraform import prismacloud_saml_bypass_subset_users.example saml_bypass_subset_users
```

## Notes

- This resource is designed for scenarios where you want to manage only a subset of SAML bypass users while allowing other users to be managed through other means (e.g., manually or by other Terraform configurations).
- When users are removed from this resource's configuration, they will be removed from the SAML bypass list, but other users not managed by this resource will remain.
- For managing the complete list of SAML bypass users, use the `prismacloud_saml_bypass_users` resource instead.