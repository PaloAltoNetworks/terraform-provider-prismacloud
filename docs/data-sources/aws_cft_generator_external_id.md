---
page_title: "Prisma Cloud: prismacloud_aws_cft_generator"
---

# prismacloud_aws_cft_generator

Retrieve information about external id and cft for aws account.

## Example Usage

```hcl
data "prismacloud_aws_cft_generator" "example" {
    account_type = "organization"
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

* `account_type` - (Required) AWS account type. `account` or `organization`.
* `account_id` - (Required) AWS account ID.
* `aws_partition` - *Applicable only for Prisma Government Stack(**app.gov.prismacloud.io**) and given if the Cloud account Global Deployment option is enabled.<br />* - **us-east-1** -  AWS Commercial/Global account.<br /> - **us-gov-west-1** - AWS GovCloud account.
* `features` - (Optional) List of features. If features key/field is not passed, then the default features will be applicable.

## Attribute Reference

* `external_id` -  AWS account external ID.
* `create_stack_link_with_s3_presigned_url` - AWS account create stack link.
* `event_bridge_rule_name_prefix` - AWS account event bridge rule name prefix.
* `s3_presigned_cft_url` - AWS CFT S3 Presigned Unencoded URL.