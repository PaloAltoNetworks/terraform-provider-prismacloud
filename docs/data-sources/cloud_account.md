---
page_title: "Prisma Cloud: prismacloud_cloud_account"
---

# prismacloud_cloud_account

Retrieve information on a specific cloud account.

## Example Usage

```hcl
data "prismacloud_cloud_account" "example" {
    cloud_type = "azure"
    name = "My Azure cloud account"
}
```

## Argument Reference

The following are the params that this data source supports.  At least one of the cloud account name and the account ID must be specified.  If one is left blank, it is determined at run time.

* `cloud_type` - (Required) The cloud type.  Valid values are `aws`, `azure`, `gcp`, or `alibaba_cloud`.
* `name` - (Optional, computed) Cloud account name; computed if this is not supplied.
* `account_id` - (Optional, computed) Account ID; computed if this is not supplied.

## Attribute Reference

The cloud type given above determines which of the attributes are populated:

* `aws` - AWS account type spec, defined [below](#aws).
* `azure` - Azure account type spec, defined [below](#azure).
* `gcp` - GCP account type spec, defined [below](#gcp).
* `alibaba_cloud` - Alibaba account type spec, defined [below](#alibaba-cloud).

### AWS

* `account_id` - AWS account ID.
* `enabled` - (bool) Whether or not the account is enabled.
* `external_id` - AWS account external ID.
* `group_ids` - List of account IDs to which you are assigning this account.
* `name` - Name to be used for the account on the Prisma Cloud platform.
* `role_arn` - Unique identifier for an AWS resource (ARN).

### Azure

* `account_id` - Azure account ID.
* `enabled` - (bool) Whether or not the account is enabled.
* `group_ids` - List of account IDs to which you are assigning this account.
* `name` - Name to be used for the account on the Prisma Cloud platform.
* `client_id` - Application ID registered with Active Directory.
* `key` - Application ID key.
* `monitor_flow_logs` - (bool) Automatically ingest flow logs.
* `tenant_id` - Active Directory ID associated with Azure.
* `service_principal_id` - Unique ID of the service principal object associated with the Prisma Cloud application that you create.

### GCP

* `account_id` - GCP account ID.
* `enabled` - (bool) Whether or not the account is enabled.
* `group_ids` - List of account IDs to which you are assigning this account.
* `name` - Name to be used for the account on the Prisma Cloud platform
* `compression_enabled` - (bool) Enable flow log compression.
* `data_flow_enabled_project` - GCP project for flow log compression.
* `flow_log_storage_bucket` - GCP Flow logs storage bucket.
* `credentials_json` - Content of the JSON credentials file.

### Alibaba Cloud

* `account_id` - Alibaba account ID.
* `group_ids` - List of account IDs to which you are assigning this account.
* `name` - Name to be used for the account on the Prisma Cloud platform.
* `ram_arn` - Unique identifier for an Alibaba RAM role resource.
