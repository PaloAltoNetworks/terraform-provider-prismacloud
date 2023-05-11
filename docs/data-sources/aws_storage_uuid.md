---
page_title: "Prisma Cloud: prismacloud_aws_storage_uuid"
---

# prismacloud_aws_storage_uuid

Retrieve information about Storage UUID. Required if you are onboarding aws account with `Data Security` feature.

## Example Usage

```hcl
data "prismacloud_aws_storage_uuid" "example" {
    account_id = "aws account id"
    external_id = "external id"
    role_arn = "aws role arn"
}
```

## Argument Reference

The following are the params that this data source supports:

* `account_id` - (Required) AWS account ID.
* `role_arn` - (Required) Unique identifier for an AWS resource (ARN).
* `external_id` - (Required) External id for aws account.

## Attribute Reference

* `storage_uuid` -  Storage UUID for aws account.
