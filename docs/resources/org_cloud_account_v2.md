---
page_title: "Prisma Cloud: prismacloud_org_cloud_account_v2"
---

# prismacloud_org_cloud_account_v2

Manage a cloud organization on the Prisma Cloud platform.

## **Example Usage 1**: AWS cloud Organization onboarding
### `Step 1`: Fetch the supported features. Refer **[Supported features readme](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloud/latest/docs/data-sources/cloud_account_supported_features)** for more details.
```hcl
data "prismacloud_account_supported_features" "prismacloud_supported_features_organization" {
    cloud_type = "aws"
    account_type = "organization"
}
```
```hcl
output "features_supported_organization" {
    value = data.prismacloud_account_supported_features.prismacloud_supported_features_organization.supported_features
}
```

### `Step 2`: Fetch the AWS CFT s3 presigned url based on required features. Refer **[AWS CFT generator Readme](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloud/latest/docs/data-sources/aws_cft_generator_external_id)** for more details.
```hcl
data "prismacloud_aws_cft_generator" "prismacloud_organization_cft" {
    account_type = "organization"
    account_id = "<account-id>"
    features = data.prismacloud_account_supported_features.prismacloud_supported_features_organization.supported_features
}
```
```hcl
output "s3_presigned_cft_url_org" {
    value = data.prismacloud_aws_cft_generator.prismacloud_organization_cft.s3_presigned_cft_url
}
```

### `Step 3`: Create the IAM Role AWS CloudFormation Stack using S3 presigned cft url from above step2
To create the IAM role using terraform, the aws official terraform aws_cloudformation_stack resource can be used. Refer https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudformation_stack for more details

OrganizationalUnitIds param: The OrganizationalUnitIds should be provided for Organization stack creation for creating member roles on the specified OrganizationalUnitIds. Provide the organizational root OU ID (prefix r-) to run it for all the accounts under the Organization, else provide a comma-separated list of OU IDs (prefix ou-). Refer https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_details.html#orgs_view_root link for more info.
Example: OrganizationalUnitIds = "r-abcd" // r-abcd is the AWS organizational root OU ID for the ORG account which indicates that member role should get created on all the accounts the organization has access to.

```
resource "aws_cloudformation_stack" "prismacloud_iam_role_stack_org" {
  name = "PrismaCloudOrgApp" // change if needed
  capabilities = ["CAPABILITY_NAMED_IAM"]
  parameters = {
    OrganizationalUnitIds = "OrganizationalUnitIds" 
    # PrismaCloudRoleName = "change-if-needed" // [Optional] A Default PrismaCloudRoleName will be present in CFT
  }
  template_url = data.prismacloud_aws_cft_generator.prismacloud_organization_cft.s3_presigned_cft_url
}

output "iam_role_org" {
    value = aws_cloudformation_stack.prismacloud_iam_role_stack_org.outputs.PrismaCloudRoleARN
}
```

### `Step 4`: Onboard the AWS cloud Organization onto prisma cloud platform

```hcl
# AWS Organization account type.
resource "prismacloud_org_cloud_account_v2" "aws_organization_onboarding_example" {
    disable_on_destroy = true
    aws {
        name = "myAwsOrganizationName" // should be unique for each account
        account_id = "<account-id>"
        account_type = "organization"
        default_account_group_id = data.prismacloud_account_group.existing_account_group_id_org.group_id// To use existing Account Group
            // prismacloud_account_group.new_account_group.group_id, // To create new Account group
        group_ids = [
            data.prismacloud_account_group.existing_account_group_id_org.group_id,// To use existing Account Group
            // prismacloud_account_group.new_account_group.group_id, // To create new Account group
        ]
        role_arn = "${aws_cloudformation_stack.prismacloud_iam_role_stack_org.outputs.PrismaCloudRoleARN}" // IAM role arn from prismacloud_iam_role_stack_org resource
        features {              // feature names from step 1
            name = "Remediation" // To enable Remediation also known as Monitor and Protect
            state = "enabled"
        }
        features {
            name = "Agentless Scanning" // To enable 'Agentless Scanning' feature if required.
            state = "enabled"
        }
    }
}

// Retrive existing account group name id
data "prismacloud_account_group" "existing_account_group_id_org" {
    name = "Default Account Group" // Change the account group name, if you already have an account group that you wish to map the account. 
}

// To create a new account group, if required
# resource "prismacloud_account_group" "new_account_group" {
#     name = "MyNewAccountGroup" // Account group name to be creatd
# }

```

### Consolidated code snippet for all the above steps

```
data "prismacloud_account_supported_features" "prismacloud_supported_features_organization" {
    cloud_type = "aws"
    account_type = "organization"
}

data "prismacloud_aws_cft_generator" "prismacloud_organization_cft" {
    account_type = "organization"
    account_id = "<account-id>"
    features = data.prismacloud_account_supported_features.prismacloud_supported_features_organization.supported_features
}

resource "aws_cloudformation_stack" "prismacloud_iam_role_stack_org" {
  name = "PrismaCloudOrgApp" // change if needed
  capabilities = ["CAPABILITY_NAMED_IAM"]
  parameters = {
    OrganizationalUnitIds = "<OrganizationalUnitIds>" // 
    # PrismaCloudRoleName = "<change-if-needed>" // [Optional] A Default PrismaCloudRoleName will be present in CFT
  }
  template_url = data.prismacloud_aws_cft_generator.prismacloud_organization_cft.s3_presigned_cft_url
}

output "iam_role_org" {
    value = aws_cloudformation_stack.prismacloud_iam_role_stack_org.outputs.PrismaCloudRoleARN
}

resource "prismacloud_org_cloud_account_v2" "aws_organization_onboarding_example" {
    disable_on_destroy = true
    aws {
        name = "myAwsOrganizationName" // should be unique for each account
        account_id = "<account-id>"
        account_type = "organization"
        default_account_group_id = data.prismacloud_account_group.existing_account_group_id_org.group_id// To use existing Account Group
            // prismacloud_account_group.new_account_group.group_id, // To create new Account group
        group_ids = [
            data.prismacloud_account_group.existing_account_group_id_org.group_id,// To use existing Account Group
            // prismacloud_account_group.new_account_group.group_id, // To create new Account group
        ]
        role_arn = "${aws_cloudformation_stack.prismacloud_iam_role_stack_org.outputs.PrismaCloudRoleARN}" // IAM role arn from step 3
        features {              // feature names from step 1
            name = "Remediation" // To enable Remediation also known as Monitor and Protect
            state = "enabled"
        }
        features {
            name = "Agentless Scanning" // To enable 'Agentless Scanning' feature if required.
            state = "enabled"
        }
    }
}

// Retrive existing account group name id
data "prismacloud_account_group" "existing_account_group_id_org" {
    name = "Default Account Group" // Change the account group name, if you already have an account group that you wish to map the account. 
}

// To create a new account group, if required
# resource "prismacloud_account_group" "new_account_group" {
#     name = "MyNewAccountGroup" // Account group name to be creatd
# }

```


## **Example Usage 2**: For Bulk AWS cloud Organization accounts onboarding

### `Prerequisite Step`: Steps 1, 2, 3 mentioned in 'Example Usage 1' should be completed for each of the Organization and have IAM roles created.

/*
You can also create cloud accounts from a CSV file using native Terraform
HCL and looping.  Assume you have a CSV file of AWS accounts that looks like this (with
"||" separating account group IDs from each other):

accountId,groupIDs,name,roleArn
123456789,Default Account Group ID||AWS Account Group ID,123466789,arn:aws:iam::123456789:role/RedlockReadWriteRole
213456789,Default Account Group ID||AWS Account Group ID,213456789,arn:aws:iam::213456789:role/RedlockReadWriteRole
321466019,Default Account Group ID||AWS Account Group ID,321466019,arn:aws:iam::321466019:role/RedlockReadWriteRole

Here's how you would do this (Terraform 0.12 code):
*/
```
locals {
    instances = csvdecode(file("aws.csv"))
}
// Now specify the cloud account resource with a loop like so:

resource "prismacloud_org_cloud_account_v2" "aws_account_bulk_onboarding_example" {
    for_each = { for inst in local.instances : inst.name => inst }
    
    aws {
        name = each.value.name
        account_type = "organization"
        account_id = each.value.accountId
        group_ids = split("||", each.value.groupIDs)
        role_arn = each.value.roleArn
    }
}
```


### prismacloud_org_cloud_account_v2 resource block example for AWS Cloud Organization with hierarchy_selection

```
resource "prismacloud_org_cloud_account_v2" "aws_organization_onboarding_example_with_hierarchy_selection" {
    disable_on_destroy = true
    aws {
        name = "myAwsOrganizationName" // should be unique for each account
        account_id = "<account-id>"
        account_type = "organization"
        default_account_group_id = data.prismacloud_account_group.existing_account_group_id_org.group_id// To use existing Account Group
            // prismacloud_account_group.new_account_group.group_id, // To create new Account group
        group_ids = [
            data.prismacloud_account_group.existing_account_group_id_org.group_id,// To use existing Account Group
            // prismacloud_account_group.new_account_group.group_id, // To create new Account group
        ]
        role_arn = "${aws_cloudformation_stack.prismacloud_iam_role_stack_org.outputs.PrismaCloudRoleARN}" // IAM role arn from step 3
        features {              // feature names from step 1
            name = "Remediation" // To enable Remediation also known as Monitor and Protect
            state = "enabled"
        }
        features {
            name = "Agentless Scanning" // To enable 'Agentless Scanning' feature if required.
            state = "enabled"
        }
        hierarchy_selection {
            display_name = "displayNameHere"
            node_type= "nodeTypeHere"
            resource_id= "resurceIdHere"
            selection_type= "selectionTypeHere"
        }
    }
}
```

## Prerequisite

Before onboarding the aws cloud account `external_id` for account must be generated using `prismacloud_aws_cft_generator`. Otherwise, you will encounter `error 412 : external_id_empty_or_not_generated`. Refer **[AWS CFT generator Readme](/docs/data-sources/aws_cft_generator_external_id.md)** for more details.

## Argument Reference

The type of cloud account to add.

* `disable_on_destroy` - (Optional, bool) To disable cloud account instead of deleting when calling Terraform destroy (default: `false`).
* `aws` - AWS account type spec, defined [below](#aws).

### AWS

* `account_id` - (Required) AWS account ID.
* `enabled` - (Optional, bool) Whether the account is enabled (default: `true`).
* `default_account_group_id` - (Optional, String) *Applicable only for accountType: **organization**.* This is the Default Account Group ID for the AWS organization and its member accounts.
* `group_ids` - (Optional) List of account IDs to which you are assigning this account.
* `name` - (Required) Name to be used for the account on the Prisma Cloud platform (must be unique).
* `role_arn` - (Required) Unique identifier for an AWS resource (ARN).
* `account_type` - (Optional) Defaults to `account` if not specified. Valid values : `account` and `organization`.
* `features` - (Optional, List) Features list

## Attribute Reference

### AWS

* `account_id` - AWS account ID.
* `enabled` - (bool) Whether the account is enabled.
* `group_ids` - List of account IDs to which you are assigning this account.
* `default_account_group_id` - (Optional, String) *Applicable only for accountType: **organization**.* This is the Default Account Group ID for the AWS organization and its member accounts.
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

#### FEATURES

* `name` - Feature name. Refer **[Supported features readme](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloud/latest/docs/data-sources/cloud_account_supported_features)** for more details.
* `state` - Feature state. Whether the feature to `enabled` or `disabled`.

#### Hierarchy Selection

* `resource_id` - Resource ID. For  ACCOUNT, OU or ROOT. Example : `root`.
* `display_name` - Display name for ACCOUNT, OU or ROOT. Example : `Root`.
* `node_type` - Node type - ORG, OU, ACCOUNT.
* `selection_type` - Selection type - ALL, INCLUDE or EXCLUDE.

## Import

Resources can be imported using the cloud type and the ID:

```
$ terraform import prismacloud_cloud_account_v2.aws_example aws:accountIdHere
```