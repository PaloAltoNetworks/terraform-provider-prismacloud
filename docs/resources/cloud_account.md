---
page_title: "Prisma Cloud: prismacloud_cloud_account"
---

# prismacloud_cloud_account

Manage a cloud account on the Prisma Cloud platform.

## Example Usage

```hcl
# Single AWS account type.
resource "prismacloud_cloud_account" "aws_example" {
    disable_on_destroy = true
    aws {
        name = "myAwsAccountName"
        account_id = "accountIdHere"
        external_id = "eidHere"
        group_ids = [
            prismacloud_account_group.g1.group_id,
        ]
        role_arn = "some:arn:here"
    }
}

resource "prismacloud_account_group" "g1" {
    name = "My group"
}

/*
You can also create cloud accounts from a CSV file using native Terraform
HCL and looping.  Assume you have a CSV file of AWS accounts that looks like this (with
"||" separating account group IDs from each other):

accountId,externalId,groupIDs,name,roleArn
123456789,PrismaExternalId,Default Account Group||AWS Account Group,123466789,arn:aws:iam::123456789:role/RedlockReadWriteRole
213456789,PrismaExternalId,Default Account Group||AWS Account Group,213456789,arn:aws:iam::213456789:role/RedlockReadWriteRole
321466019,PrismaExternalId,Default Account Group||AWS Account Group,321466019,arn:aws:iam::321466019:role/RedlockReadWriteRole

Here's how you would do this (Terraform 0.12 code):
*/
locals {
    instances = csvdecode(file("aws.csv"))
}

// Now specify the cloud account resource with a loop like so:
resource "prismacloud_cloud_account" "csv" {
    for_each = { for inst in local.instances : inst.name => inst }

    aws {
        name = each.value.name
        account_id = each.value.accountId
        external_id = each.value.externalId
        group_ids = split("||", each.value.groupIDs)
        role_arn = each.value.roleArn
    }
}
```

## Argument Reference

The type of cloud account to add.  You need to specify one and only one of these cloud types.

* `disable_on_destroy` - (Optional, bool) To disable cloud account instead of deleting when calling Terraform destroy (default: `false`).
* `aws` - AWS account type spec, defined [below](#aws).
* `azure` - Azure account type spec, defined [below](#azure).
* `gcp` - GCP account type spec, defined [below](#gcp).
* `alibaba_cloud` - Alibaba account type spec, defined [below](#alibaba-cloud).

### AWS

> **Notice**
> #### No Change in existing behavior of AWS cloud accounts onboarding
> *In version 1.2.7 of Prismacloud Terraform Provider, we published a lookahead notice regarding the confused deputy attack changes for AWS cloud account onboarding.
> *The upstream API is no longer moving ahead with the previous approach, so the corresponding changes to the Terraform Provider will not be released at this time. Subsequent upstream API changes related to cloud account onboarding will be incorporated into the Terraform Provider through a similar approach as they are made available.

* `account_id` - (Required) AWS account ID.
* `enabled` - (Optional, bool) Whether or not the account is enabled (default: `true`).
* `external_id` - (Required) AWS account external ID.
* `group_ids` - (Required) List of account IDs to which you are assigning this account.
* `name` - (Required) Name to be used for the account on the Prisma Cloud platform (must be unique).
* `role_arn` - (Required) Unique identifier for an AWS resource (ARN).
* `account_type` - (Optional) Defaults to "account" if not specified
* `protection_mode` - (Optional) Defaults to "MONITOR".Valid values : `MONITOR` or `MONITOR_AND_PROTECT`

### Azure


* `account_id` - (Required) Azure account ID.
* `enabled` - (Optional, bool) Whether or not the account is enabled (defualt: `true`).
* `group_ids` - (Required) List of account IDs to which you are assigning this account.
* `name` - (Required) Name to be used for the account on the Prisma Cloud platform (must be unique).
* `client_id` - (Required) Application ID registered with Active Directory.
* `key` - (Required) Application ID key.
* `monitor_flow_logs` - (Optional, bool) Automatically ingest flow logs.
* `tenant_id` - (Required) Active Directory ID associated with Azure.
* `service_principal_id` - (Required) Unique ID of the service principal object associated with the Prisma Cloud application that you create.
* `account_type` - (Optional) Defaults to "account" if not specified
* `protection_mode` - (Optional) Defaults to "MONITOR". Valid values : `MONITOR` or `MONITOR_AND_PROTECT`

### GCP

!> The Prisma Cloud API returns a series of asterisks for the private key of the `credentials_json` field instead of the configured value.  Because of this, the provider cannot detect configuration drift on the private key within the `credentials_json` param.

* `account_id` - (Required) GCP project ID.
* `enabled` - (Optional, bool) Whether or not the account is enabled (defualt: `true`).
* `group_ids` - (Required) List of account IDs to which you are assigning this account.
* `name` - (Required) Name to be used for the account on the Prisma Cloud platform (must be unique).
* `compression_enabled` - (Optional, bool) Enable flow log compression.
* `dataflow_enabled_project` - (Optional) GCP project for flow log compression.
* `flow_log_storage_bucket` - (Optional) GCP Flow logs storage bucket.
* `credentials_json` - (Required) Content of the JSON credentials file (read in using `file()`).
* `account_type` - (Optional) Defaults to "account" if not specified
* `protection_mode` - (Optional) Defaults to "MONITOR". Valid values : `MONITOR` or `MONITOR_AND_PROTECT`

### Alibaba Cloud

* `account_id` - (Required) Alibaba account ID.
* `group_ids` - (Required) List of account IDs to which you are assigning this account.
* `name` - (Required) Name to be used for the account on the Prisma Cloud platform (must be unique).
* `ram_arn` - (Required) Unique identifier for an Alibaba RAM role resource.

## Import

Resources can be imported using the cloud type (`aws`, `azure`, `gcp`, or `alibaba_cloud`) and the ID:

```
$ terraform import prismacloud_cloud_account.aws_example aws:accountIdHere
```
