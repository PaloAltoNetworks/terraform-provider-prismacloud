---
layout: "prismacloud"
page_title: "Prisma Cloud: prismacloud_cloud_account"
description: |-
  Manage a cloud account on the Prisma Cloud platform.
---

# prismacloud_cloud_account

Manage a cloud account on the Prisma Cloud platform.

## Example Usage

```hcl
# Single AWS account type.
resource "prismacloud_cloud_account" "aws_example" {
    aws {
        name = "myAwsAccountName"
        account_id = "accountIdHere"
        external_id = "eidHere"
        group_ids = []
        role_arn = "some:arn:here"
    }
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

* `aws` - AWS account type spec, defined [below](#aws).
* `azure` - Azure account type spec, defined [below](#azure).
* `gcp` - GCP account type spec, defined [below](#gcp).
* `alibaba_cloud` - Alibaba account type spec, defined [below](#alibaba-cloud).

### AWS

* `account_id` - (Required) AWS account ID.
* `enabled` - (Optional, bool) Whether or not the account is enabled (defualt: `true`).
* `external_id` - (Required) AWS account external ID.
* `group_ids` - (Required) List of account IDs to which you are assigning this account.
* `name` - (Required) Name to be used for the account on the Prisma Cloud platform (must be unique).
* `role_arn` - (Required) Unique identifier for an AWS resource (ARN).

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

### GCP

* `account_id` - (Required) GCP account ID.
* `enabled` - (Optional, bool) Whether or not the account is enabled (defualt: `true`).
* `group_ids` - (Required) List of account IDs to which you are assigning this account.
* `name` - (Required) Name to be used for the account on the Prisma Cloud platform (must be unique).
* `compression_enabled` - (Optional, bool) Enable flow log compression.
* `data_flow_enabled_project` - (Optional) GCP project for flow log compression.
* `flow_log_storage_bucket` - (Optional) GCP Flow logs storage bucket.
* `credentials_json` - (Required) Content of the JSON credentials file (read in using `file()`).

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
