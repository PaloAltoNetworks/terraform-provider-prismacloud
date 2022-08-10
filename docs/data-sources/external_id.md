---
page_title: "Prisma Cloud: prismacloud_externalid"
---

# prismacloud_externalid

Retrieve information on a specific external id for aws account.

## Example Usage

```hcl
data "prismacloud_external_id" "example" {
    name = "Aws account name"
}
```

## Argument Reference

You must specify all the following:

* `name` - AWS account name.
* `protection_mode` - `Monitor` or `Monitor and Protect`.
* `account_id` - AWS account ID.
* `aws_partition` - The aws cloud account partition.
* `storage_scan_enabled` - (bool) Whether the storage_scan is enabled.

## Attribute Reference

* `external_id` -  AWS account external ID.
* `cft_path` - AWS account cft path.
* `cloud_formation_url` - AWS account cloud formation url.

