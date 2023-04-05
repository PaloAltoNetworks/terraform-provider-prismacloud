---
page_title: "Prisma Cloud: prismacloud_azure_template"
---

# prismacloud_azure_template

Retrieve information about () for azure account.

## Example Usage

```hcl
data "prismacloud_azure_template" "example" {
  account_type = "tenant"
  tenant_id    = "<tenant-id>"
}
```

## Example usage to output s3_presigned_cft_url for update account and organization scenario

```hcl
output "s3_presigned_cft_url" {
  value = data.prismacloud_aws_cft_generator.example.s3_presigned_cft_url
}
```

## Example usage to output external_id

```hcl
output "external_id" {
  value = data.prismacloud_aws_cft_generator.example.external_id
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

## Attribute Reference

