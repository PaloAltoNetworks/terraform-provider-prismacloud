---
page_title: "Prisma Cloud: prismacloud_azure_template"
---

# prismacloud_azure_template

Retrieve information about azure template for azure account.

## Example Usage for Azure Subscription

```hcl
data "prismacloud_azure_template" "example" {
  file_name       = "<file-name>" //filename along with path should be specified
  account_type    = "account"
  subscription_id = "<subscription-id>"
  tenant_id       = "<tenant_id>"
}
```

## Example Usage for Azure Active Directory

```hcl
data "prismacloud_azure_template" "example" {
  file_name    = "<file-name>"  //filename along with path should be specified
  account_type = "tenant"
  tenant_id    = "<tenant-id>"
}
```

## Example Usage for Azure Tenant

```hcl
data "prismacloud_azure_template" "example" {
  file_name         = "<file-name>" //filename along with path should be specified
  account_type      = "tenant"
  tenant_id         = "<tenant-id>"
  root_sync_enabled = true
}
```

## Example usage to output azure_template

```hcl
output "azure_template" {
  value = data.prismacloud_azure_template.example.azure_template
}
```

## Argument Reference

The following are the params that this data source supports:

* `account_type` - (Required) Azure account type. `account` or `tenant`.
* `tenant_id` - (Required) Azure tenant ID.
* `file_name` - (Required) File name to store azure template (complete path should be specified).
* `subscription_id` - (Optional) Azure subscription ID.
* `root_sync_enabled` - (Optional) Azure tenant has children. Must be set to true `account_type` is `tenant`.
* `deployment_type` - (Optional) `azure` or `azure_gov` for azure account.
* `features` - (Optional) List of features. If features key/field is not passed, then the default features will be
  applicable. Refer : *
  *[Supported features readme](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloud/latest/docs/data-sources/cloud_account_supported_features)
  ** for more details.


