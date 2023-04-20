---
page_title: "Prisma Cloud: prismacloud_cloud_account_v2"
---

# prismacloud_cloud_account_v2

Retrieve information on a specific cloud account.

## Example Usage

```hcl
data "prismacloud_cloud_account_v2" "example" {
    cloud_type = "aws"
    name = "My Aws cloud account"
}
```

## Argument Reference

The following are the params that this data source supports.  At least one of the cloud account name and the account ID must be specified.  If one is left blank, it is determined at run time.

* `cloud_type` - (Required) The cloud type. Valid value is `aws`, `azure`, `gcp` or `ibm`.
* `name` - (Optional, computed) Cloud account name; computed if this is not supplied. Applicable only for `aws`, `azure` and `ibm`.
* `account_id` - (Optional, computed) Account ID; computed if this is not supplied.

## Attribute Reference

The cloud type given above determines which of the attributes are populated:

* `disable_on_destroy` - (bool) To disable cloud account instead of deleting when calling Terraform destroy.
* `aws` - AWS account type spec, defined [below](#aws).
* `azure` - Azure account type spec, defined [below](#azure).
* `gcp` - Gcp account type spec, defined [below](#gcp).
* `ibm` - IBM account type spec, defined [below](#ibm).

### AWS

* `account_id` - AWS account ID.
* `enabled` - (bool) Whether the account is enabled.
* `group_ids` - List of account IDs to which you are assigning this account.
* `name` - Name to be used for the account on the Prisma Cloud platform.
* `role_arn` - Unique identifier for an AWS resource (ARN).
* `created_epoch_millis` - Account created epoch time.
* `customer_name` - Prisma customer name.
* `deleted` - Whether the account is deleted or not.
* `deployment_type` - `aws` for aws account.
* `eventbridge_rule_name_prefix` -  Eventbridge rule name prefix.
* `external_id` - External id for aws account.
* `features` - Features applicable for aws account, defined [below](#features).
* `last_modified_by` - Last modified by.
* `last_modified_epoch_millis` - Last modified at epoch millis.
* `parent_id` - Parent id.
* `protection_mode` - Protection mode of account.

### Azure

* `account_id` - Azure account ID.
* `client_id` - Application ID registered with Active Directory.
* `enabled` - (bool) Whether the account is enabled.
* `group_ids` - List of account IDs to which you are assigning this account.
* `name` - Name to be used for the account on the Prisma Cloud platform (must be unique).
* `key` - Application ID key.
* `monitor_flow_logs` - (bool) Automatically ingest flow logs.
* `tenant_id` - Active Directory ID associated with Azure.
* `service_principal_id` - Unique ID of the service principal object associated with the Prisma Cloud application that you create.
* `account_type` - `account` for azure subscription account.
* `protection_mode` - Protection mode of account.
* `features` - Features applicable for azure account, defined [below](#features).
* `environment_type` - Environment type.
* `parent_id` - Parent id.
* `customer_name` - Prisma customer name.
* `created_epoch_millis` - Account created epoch time.
* `last_modified_by` - Last modified by.
* `last_modified_epoch_millis` - Last modified at epoch millis.
* `deleted` - (bool) Whether the account is deleted or not.
* `template_url` - Template URL.
* `deployment_type` - `az` for azure account.
* `deployment_type_description` - Deployment type description. 

### Gcp

* `account_id` - Gcp account ID.
* `account_type` - `account` for gcp project account and `masterServiceAccount` for gcp master service account.
* `enabled` - (bool) Whether the account is enabled.
* `group_ids` - List of account IDs to which you are assigning this account.
* `name` - Name to be used for the account on the Prisma Cloud platform (must be unique).
* `compression_enabled` - (bool) Enable or disable compressed network flow log generation.
* `credentials` - Content of the JSON credentials file.
* `data_flow_enabled_project` - Project ID where the Dataflow API is enabled .
* `features` - Features applicable for gcp account, defined [below](#features).
* `flow_log_storage_bucket` - Cloud Storage Bucket name that is used store the flow logs.
* `protection_mode` - Protection mode of account.
* `parent_id` - Parent ID.
* `customer_name` - Prisma customer name.
* `created_epoch_millis` - Account created epoch time.
* `last_modified_by` - Last modified by.
* `last_modified_epoch_millis` - Last modified at epoch millis.
* `deleted` - (bool) Whether the account is deleted or not.
* `storage_scan_enabled` - (bool) Whether the storage scan is enabled.
* `added_on_ts` - Added on time stamp.
* `deployment_type` - `gcp` for gcp account.
* `deployment_type_description` - Deployment type description.
* `project_id` - Gcp Project ID.
* `service_account_email` - Service account email of gcp account.
* `authentication_type` - Authentication type of gcp account.
* `account_group_creation_mode` - Account group creation mode.
* `default_account_group_id` - Account group id to which you are assigning this account. Must be provided for gcp `masterServiceAccount`.

### IBM

* `account_id` - IBM account ID.
* `account_type` - `account` for IBM account.
* `api_key` - IBM service API key.
* `enabled` - (bool) Whether the account is enabled.
* `group_ids` - List of account IDs to which you are assigning this account.
* `name` - Name to be used for the account on the Prisma Cloud platform (must be unique).
* `svc_id_iam_id` - IBM service ID.
* `added_on_ts` - Added on time stamp.
* `created_epoch_millis` - Account created epoch time.
* `customer_name` - Prisma customer name.
* `deleted` - (bool) Whether the account is deleted or not.
* `deployment_type` - `ibm` for IBM account.
* `deployment_type_description` - Deployment type description.
* `features` - Features applicable for IBM account, defined [below](#features).
* `last_modified_epoch_millis` - Last modified at epoch millis.
* `last_modified_by` - Last modified by.
* `parent_id` - Parent id.
* `protection_mode` - Protection mode of account.
* `storage_scan_enabled` - (bool) Whether the storage scan is enabled.

#### FEATURES

* `name` - Feature name.
* `state` - Feature state. `enabled` or `disabled`.
