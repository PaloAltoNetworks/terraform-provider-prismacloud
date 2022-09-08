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

* `name` - (Required) AWS account name.
* `protection_mode` - (Required) Protection mode. Valid values : `MONITOR` or `MONITOR_AND_PROTECT`.
* `account_id` - AWS account ID.
* `storage_scan_enabled` - (Required, bool) Whether the storage_scan is enabled.
* `external_id` -  (Required) AWS account external ID.

## Attribute Reference


* `cft_path` - AWS account cft path.
* `cloud_formation_url` - AWS account cloud formation url.
