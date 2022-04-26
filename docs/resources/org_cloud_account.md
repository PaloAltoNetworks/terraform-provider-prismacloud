---
page_title: "Prisma Cloud: prismacloud_org_cloud_account"
---

# prismacloud_org_cloud_account

Manage a org cloud account on the Prisma Cloud platform.

## Example Usage

```hcl
# Single AWS org account type.
resource "prismacloud_org_cloud_account" "aws_example" {
    disable_on_destroy = true
    aws {
        name = "myAwsOrgAccountName"
        account_id = "accountIdHere"
        external_id = "eidHere"
        member_external_id = "membereidHere"
        group_ids = [
            prismacloud_account_group.g1.group_id,
        ]
        role_arn = "some:arn:here"
        member_role_name = "memberRoleHere"
    }
}

resource "prismacloud_account_group" "g1" {
    name = "My group"
}

/*
You can also create cloud accounts from a CSV file using native Terraform
HCL and looping.  Assume you have a CSV file of AWS accounts that looks like this (with
"||" separating account group IDs from each other):

accountId,externalId,groupIDs,name,roleArn,memberRoleName,memberExternalId
123456789,PrismaExternalId,Default Account Group||AWS Account Group,123466789,arn:aws:iam::123456789:role/RedlockReadWriteRole,PrismaMemberRole,PrismaMemberExternalId 
213456789,PrismaExternalId,Default Account Group||AWS Account Group,213456789,arn:aws:iam::213456789:role/RedlockReadWriteRole,PrismaMemberRole,PrismaMemberExternalId 
321466019,PrismaExternalId,Default Account Group||AWS Account Group,321466019,arn:aws:iam::321466019:role/RedlockReadWriteRole,PrismaMemberRole,PrismaMemberExternalId

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
        member_external_id = each.value.externalId
        group_ids = split("||", each.value.groupIDs)
        role_arn = each.value.roleArn
        member_role_name = each.memberRoleName
    }
}

// Cloud account with hierarchy_selection 
resource "prismacloud_org_cloud_account" "aws_example" {
    disable_on_destroy = true
    aws {
        name = "myAwsOrgAccountName"
        account_id = "accountIdHere"
        external_id = "eidHere"
        member_external_id = "membereidHere"
        group_ids = [
            prismacloud_account_group.g1.group_id,
        ]
        role_arn = "some:arn:here"
        member_role_name = "memberRoleHere"
        hierarchy_selection {
            display_name = "displayNameHere"
            node_type= "nodeTypeHere"
            resource_id= "resurceIdHere"
            selection_type= "selectionTypeHere"
        }
    }
}
```

## Argument Reference

The type of org cloud account to add.  You need to specify one and only one of these cloud types.

* `disable_on_destroy` - (Optional,bool) To disable cloud account instead of deleting when calling Terraform destroy (default: `false`).
* `aws` - AWS org account type spec, defined [below](#aws).
* `azure` - Azure org account type spec, defined [below](#azure).
* `gcp` - GCP org account type spec, defined [below](#gcp).
* `oci` - Oci account type spec, defined [below](#oci).

### AWS

* `account_id` - (Required) AWS Org account ID.
* `enabled` - (Optional, bool) Whether or not the account is enabled (default: `true`).
* `external_id` - (Required) AWS org account external ID.
* `group_ids` - (Required) List of account IDs to which you are assigning this account.
* `name` - (Required) Name to be used for the account on the Prisma Cloud platform (must be unique).
* `role_arn` - (Required) Unique identifier for an AWS org resource (ARN).
* `account_type` - (Optional) Defaults to "organization" if not specified.
* `member_role_name`- (Required) AWS org Member account role name. 
* `member_external_id` - (Required) AWS org Member account role's external ID.
* `member_role_status` - (Optional, bool) - True =  The member role created using stack set exists in all the member accounts. 
                        All the Org accounts will be added. false = Only the master account will be added(Default = False).
* `protection_mode` - (Optional) Defaults to "MONITOR". Valid values : `MONITOR` or `MONITOR_AND_PROTECT`
* `hierarchy_selection` -  (Optional) List of AWS Organization Units (OU), AWS accounts, and AWS Organizations to onboard under this organization. [below](#For-AWS)

### Azure

* `account_id` - (Required) Azure org account ID.
* `enabled` - (Optional, bool) Whether or not the account is enabled (defualt: `true`).
* `group_ids` - (Required) List of account IDs to which you are assigning this account.
* `name` - (Required) Name to be used for the account on the Prisma Cloud platform (must be unique).
* `client_id` - (Required) Application ID registered with Active Directory.
* `key` - (Required) Application ID key.
* `monitor_flow_logs` - (Required, bool) Automatically ingest flow logs.
* `tenant_id` - (Required) Active Directory ID associated with Azure.
* `service_principal_id` - (Required) Unique ID of the service principal object associated with the Prisma Cloud application that you create.
* `account_type` - (Optional) Defaults to "tenant" if not specified.
* `protection_mode` - (Optional) Defaults to "MONITOR". Valid values : `MONITOR`

### GCP

* `account_id` - (Required) GCP org project ID.
* `enabled` - (Optional, bool) Whether or not the account is enabled (defualt: `true`).
* `group_ids` - List of account IDs to which you are assigning this account.
* `name` - (Required) Name to be used for the account on the Prisma Cloud platform (must be unique).
* `compression_enabled` - (Optional, bool) Enable flow log compression.
* `data_flow_enabled_project` - (Optional) GCP project for flow log compression.
* `flow_log_storage_bucket` - (Optional) GCP Flow logs storage bucket.
* `credentials_json` - (Required) Content of the JSON credentials file (read in using `file()`).
* `account_type` - (Optional) Account type. Defaults to `organization` if not specified.
* `protection_mode` - (Optional) Protection Mode. Valid values : `MONITOR` or `MONITOR_AND_PROTECT`. Defaults to `MONITOR` if not specified.
* `organization_name` - (Required) GCP org organization name.
* `account_group_creation_mode` - (Optional) Cloud account group creation mode. Valid values : `MANUAL`: Create account groups manually, `AUTO`: Create high-level account groups based on folders identified, or `RECURSIVE`: Drill down in folder tree to create account groups (default : `MANUAL`). `AUTO` can't be used if `selection_type` in `hierarchy_selection` is `EXCLUDE`. 
* `hierarchy_selection` - (Optional) List of hierarchy selection. Each item has resource ID, display name, node type and selection type, as defined [below](#For-GCP). 

#### Hierarchy Selection

##### For AWS
* `resource_id` - (Required) Resource ID. Valid values are AWS OU ID, AWS account ID, or AWS Organization ID. Note you must escape any double quotes in the resource ID with a backslash.
* `display_name` - (Required) Display name for AWS OU, AWS account, or AWS organization.
* `node_type` - (Required) Valid values: `OU`, `ACCOUNT`, `ORG`.
* `selection_type` - (Required) Valid values: `INCLUDE` to include the specified resource to onboard, `EXCLUDE` to exclude the specified resource and onboard the rest, `ALL` to onboard all resources in the organization.

##### For GCP
* `resource_id` - (Required) Resource ID.  For folders, format is folders/{folder ID}. For projects, format is {project number}. For orgs, format is organizations/{org ID}
* `display_name` - (Required) Display name for folder, project, or organization.
* `node_type` - (Required) Node type. Valid values : `FOLDER`, `PROJECT`, or `ORG`.
* `selection_type` - (Required) Selection type. Valid values: `INCLUDE`, `EXCLUDE`, or `ALL`. If `node_type` is `PROJECT` or `FOLDER`, then a valid value is either `INCLUDE` or `EXCLUDE`.

### Oci

* `account_id` - (Required) OCI account ID.
* `enabled` - (Optional, bool) Whether or not the account is enabled (default: `true`).
* `group_name` - (Required) OCI identity group name that you define. Can be an existing group.
* `group_ids` - (Required)  account ID to which you are assigning this account.
* `name` - (Required) Name to be used for the account on the Prisma Cloud platform (must be unique). 
* `group_name` - (Required) OCI identity group name that you define. Can be an existing group.
* `ram_arn` - (Required) Unique identifier for an Alibaba RAM role resource.
* `account_type` - (Required) Account type - account or tenant.
* `default_account_group_id` - (Required)  account ID to which you are assigning this account.
* `home_region` - (Required) OCI tenancy home region.
* `policy_name` - (Required) OCI identity policy name that you define. Can be an existing policy that has the right policy statements. 
* `user_name` - (Required) OCI identity user name that you define. Can be an existing user that has the right privileges.
* `user_ocid` - (Required) OCI identity user Ocid that you define. Can be an existing user that has the right privileges.

## Import

Resources can be imported using the cloud type (`aws`, `azure`, `gcp`, or `oci`) and the ID:

```
$ terraform import prismacloud_org_cloud_account.aws_example aws:accountIdHere
```
