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

* `cloud_type` - (Required) The cloud type. Valid value is `aws` or `azure`.
* `name` - (Optional, computed) Cloud account name; computed if this is not supplied.
* `account_id` - (Optional, computed) Account ID; computed if this is not supplied.

## Attribute Reference

The cloud type given above determines which of the attributes are populated:

* `disable_on_destroy` - (bool) To disable cloud account instead of deleting when calling Terraform destroy.
* `aws` - AWS account type spec, defined [below](#aws).
* `azure` - Azure account type spec, defined [below](#azure).

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

#### FEATURES

* `name` - Feature name.
* `state` - Feature state. `enabled` or `disabled`.
