---
page_title: "Prisma Cloud: prismacloud_org_cloud_account"
---

# prismacloud_org_cloud_account

Retrieve information on a specific cloud account.

## Example Usage

```hcl
data "prismacloud_org_cloud_account" "example" {
    cloud_type = "azure"
    name = "My Azure cloud account"
}
```

## Argument Reference

The following are the params that this data source supports.  At least one of the cloud account name and the account ID must be specified.  If one is left blank, it is determined at run time.

* `cloud_type` - (Required) The cloud type.  Valid values are `aws`, `azure`, `gcp`, or `oci`.
* `name` - (Optional, computed) Cloud account name; computed if this is not supplied.
* `account_id` - (Optional, computed) Account ID; computed if this is not supplied.

## Attribute Reference

The cloud type given above determines which of the attributes are populated:

* `disable_on_destroy` - (bool) To disable cloud account instead of deleting when calling Terraform destroy (default: `false`).
* `aws` - AWS org account type spec, defined [below](#aws).
* `azure` - Azure org account type spec, defined [below](#azure).
* `gcp` - GCP org account type spec, defined [below](#gcp).
* `oci` - Oci account type spec, defined [below](#OCI).

### AWS

* `account_id` - AWS account ID.
* `enabled` - (bool) Whether or not the account is enabled.
* `external_id` - AWS account external ID.
* `group_ids` - List of account IDs to which you are assigning this account.
* `name` - Name to be used for the account on the Prisma Cloud platform (must be unique).
* `role_arn` - Unique identifier for an AWS resource (ARN).
* `account_type` - Defaults to "account" if not specified.
* `protection_mode` - Defaults to "MONITOR".

### Azure

* `account_id` - Azure org account ID.
* `enabled` - (bool) Whether or not the account is enabled.
* `group_ids` - List of account IDs to which you are assigning this account.
* `name` - Name to be used for the account on the Prisma Cloud platform (must be unique).
* `client_id` - Application ID registered with Active Directory.
* `key` - Application ID key.
* `monitor_flow_logs` - (bool) Automatically ingest flow logs.
* `tenant_id` - Active Directory ID associated with Azure.
* `service_principal_id` - Unique ID of the service principal object associated with the Prisma Cloud application that you create.
* `account_type` - Defaults to "tenant" if not specified.
* `protection_mode` - Defaults to "MONITOR".

### GCP

* `account_id` - GCP org project ID.
* `enabled` - (bool) Whether or not the account is enabled.
* `group_ids` - List of account IDs to which you are assigning this account.
* `name` - Name to be used for the account on the Prisma Cloud platform.
* `compression_enabled` - (bool) Enable flow log compression.
* `data_flow_enabled_project` - GCP project for flow log compression.
* `flow_log_storage_bucket` - GCP Flow logs storage bucket.
* `credentials_json` - Content of the JSON credentials file.
* `account_type` - Account type - organization, or account.
* `protection_mode` - Protection Mode - Monitor, or Monitor and Protect.
* `organization_name` - GCP org organization name.
* `account_group_creation_mode` - Cloud account group creation mode - MANUAL, AUTO, or RECURSIVE.
* `hierarchy_selection` - List of hierarchy selection. Each item has resource ID, display name, node type and selection type, as defined [below](#hierarchy-selection).

#### Hierarchy Selection

* `resource_id` - Resource ID. For folders, format is folders/{folder ID}. For projects, format is {project number}. For orgs, format is organizations/{org ID}.
* `display_name` - Display name for folder, project, or organization.
* `node_type` - Node type - FOLDER, PROJECT, or ORG.
* `selection_type` - Selection type - INCLUDE, EXCLUDE, or ALL.

### OCI

* `account_id` - Oci account ID.
* `enabled` - (bool) Whether or not the account is enabled.
* `group_name` - (Required) OCI identity group name that you define. Can be an existing group.
* `group_ids` - account ID to which you are assigning this account.
* `name` - Name to be used for the account on the Prisma Cloud platform (must be unique).
* `group_name` - OCI identity group name that you define. Can be an existing group
* `ram_arn` - Unique identifier for an Alibaba RAM role resource.
* `account_type` - Account type - account or tenant
* `default_account_group_id` - (Required)  account ID to which you are assigning this account.
* `home_region` - OCI tenancy home region
* `policy_name` - OCI identity policy name that you define. Can be an existing policy that has the right policy statements
* `user_name` - OCI identity user name that you define. Can be an existing user that has the right privileges
* `user_ocid` - OCI identity user Ocid that you define. Can be an existing user that has the right privileges
