---
page_title: "Prisma Cloud: prismacloud_external_id"
---

# prismacloud_external_id

Retrieve information about external id and cft for aws account.

## Example Usage

```hcl
data "prismacloud_external_id" "example" {
    account_type = "organization"
}
```

## Argument Reference

The following are the params that this data source supports:

* `account_type` - (Required) AWS account type. `account` or `organization`.
* `account_id` - (Required) AWS account ID.
* `aws_partition` - (Optional) The aws cloud account partition. Valid values are `us-gov-west-1`, `cn-north-1`,  or `us-east-1`.
* `features` - (Optional) List of features. If features key/field is not passed, then it defaults to all the features applicable for the combination of `licenseType`,`cloudType`,`accountType`,`deploymentType`,`stack`.

## Attribute Reference

* `external_id` -  AWS account external ID.
* `create_stack_link_with_s3_presigned_url` - AWS account cft link.
* `event_bridge_rule_name_prefix` - AWS account event bridge rule name prefix.