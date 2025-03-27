---
page_title: "Prisma Cloud: prismacloud_account_supported_features"
---

# prismacloud_account_supported_features

Retrieve information about Supported Features For Cloud Type.

## Example Usage

```hcl
data "prismacloud_account_supported_features" "prismacloud_supported_features" {
    cloud_type = "aws"
    account_type = "account"
}
```

## Example usage to output supported features for AWS CFT template generator and account, organization onboarding scenario

```hcl
output "features_supported" {
    value = data.prismacloud_account_supported_features.prismacloud_supported_features.supported_features
}
```

## Argument Reference

The following are the params that this data source supports:

* `cloud_type` - (Required) Cloud type. `aws`, `azure`, or `gcp`.
* `account_type` - (Required) Cloud account type. `account`, `organization`, `masterServiceAccount`, `tenant` or `workspaceDomain`. Supported values based on cloud_type are given below. <br /> - account, organization - cloud_type: **aws**<br /> - account, organization, masterServiceAccount, workspaceDomain - cloud_type: **gcp** <br /> - account, tenant - cloud_type: **azure**
* `deployment_type` - (Optional) *Applicable only for cloud_type: **azure***. Possible values: `azure`, `azure_gov`, or `azure_china`. <br /> - **azure** -  Account type is commercial<br /> - **azure_gov** - Account type is Government on Prisma Commercial and Government stacks.<br /> - **azure_china** - Prisma China Stack.
* `aws_partition` - (Optional) *Applicable only for Prisma Government Stack(**app.gov.prismacloud.io**) and given if the Cloud account Global Deployment option is enabled.<br />* - **us-east-1** -  AWS Commercial/Global account.<br /> - **us-gov-west-1** - AWS GovCloud account.
* `root_sync_enabled` - (Optional) *Applicable only for cloud_type: **azure** and accountType: **tenant***.<br />  In order to get supported features for accountType **tenant** and its associated **management groups** and **subscriptions**, rootSyncEnabled must be set to true.

## Attribute Reference

* `cloud_type` -  Cloud type.
* `deployment_type` - Cloud Account deployment type.
* `account_type` - Cloud account type.
* `license_type` - Customer License type.
* `supported_features` - List of supported feature names.
