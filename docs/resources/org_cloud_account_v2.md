---
page_title: "Prisma Cloud: prismacloud_org_cloud_account_v2"
---

# prismacloud_org_cloud_account_v2

Manage a cloud organization on the Prisma Cloud platform.

## **Example Usage 1**: AWS cloud Organization onboarding

### `Step 1`: Fetch the supported features. Refer *

*[Supported features readme](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloud/latest/docs/data-sources/cloud_account_supported_features)
** for more details.

```hcl
data "prismacloud_account_supported_features" "prismacloud_supported_features_organization" {
  cloud_type   = "aws"
  account_type = "organization"
}
```

```hcl
output "features_supported_organization" {
  value = data.prismacloud_account_supported_features.prismacloud_supported_features_organization.supported_features
}
```

### `Step 2`: Fetch the AWS CFT s3 presigned url based on required features. Refer *

*[AWS CFT generator Readme](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloud/latest/docs/data-sources/aws_cft_generator_external_id)
** for more details.

```hcl
data "prismacloud_aws_cft_generator" "prismacloud_organization_cft" {
  account_type = "organization"
  account_id   = "<account-id>"
  features     = data.prismacloud_account_supported_features.prismacloud_supported_features_organization.supported_features
}
```

```hcl
output "s3_presigned_cft_url_org" {
  value = data.prismacloud_aws_cft_generator.prismacloud_organization_cft.s3_presigned_cft_url
}
```

### `Step 3`: Create the IAM Role AWS CloudFormation Stack using S3 presigned cft url from above step2

OrganizationalUnitIds param: The OrganizationalUnitIds should be provided for Organization stack creation for creating
member roles on the specified OrganizationalUnitIds. Provide the organizational root OU ID (prefix r-) to run it for all
the accounts under the Organization, else provide a comma-separated list of OU IDs (prefix ou-).
Refer https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_details.html#orgs_view_root link for
more info.
Example: OrganizationalUnitIds = "r-abcd" // r-abcd is the AWS organizational root OU ID for the ORG account which
indicates that member role should get created on all the accounts the organization has access to.

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
    name                     = "myAwsOrganizationName" // should be unique for each account
    account_id               = "<account-id>"
    account_type             = "organization"
    default_account_group_id = data.prismacloud_account_group.existing_account_group_id_org.group_id
    // To use existing Account Group
    // prismacloud_account_group.new_account_group.group_id, // To create new Account group
    group_ids                = [
      data.prismacloud_account_group.existing_account_group_id_org.group_id, // To use existing Account Group
      // prismacloud_account_group.new_account_group.group_id, // To create new Account group
    ]
    role_arn = "${aws_cloudformation_stack.prismacloud_iam_role_stack_org.outputs.PrismaCloudRoleARN}"
    // IAM role arn from prismacloud_iam_role_stack_org resource
    features {
      // feature names from step 1
      name  = "Remediation" // To enable Remediation also known as Monitor and Protect
      state = "enabled"
    }
    features {
      name  = "Agentless Scanning" // To enable 'Agentless Scanning' feature if required.
      state = "enabled"
    }
  }
}

// Retrive existing account group name id
data "prismacloud_account_group" "existing_account_group_id_org" {
  name = "Default Account Group"
  // Change the account group name, if you already have an account group that you wish to map the account. 
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
HCL and looping. Assume you have a CSV file of AWS accounts that looks like this (with
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

Before onboarding the aws cloud account `external_id` for account must be generated
using `prismacloud_aws_cft_generator`. Otherwise, you will encounter `error 412 : external_id_empty_or_not_generated`.
Refer **[AWS CFT generator Readme](/docs/data-sources/aws_cft_generator_external_id.md)** for more details.

## **Example Usage 3**: Azure cloud Tenant onboarding

### `Step 1`: Fetch the supported features. Refer *

*[Supported features readme](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloud/latest/docs/data-sources/cloud_account_supported_features)
** for more details.

```hcl
data "prismacloud_account_supported_features" "prismacloud_supported_features" {
  cloud_type      = "azure"
  account_type    = "tenant"
  deployment_type = "azure"
}
```

```hcl
output "features_supported_tenant" {
  value = data.prismacloud_account_supported_features.prismacloud_supported_features_tenant.supported_features
}
```

### `Step 2`: Fetch the Azure template based on required features. Refer *

*[Azure template generator Readme](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloud/latest/docs/data-sources/azure_template)
** for more details.

```hcl
data "prismacloud_azure_template" "prismacloud_tenant_azure_template" {
  file_name       = "<file-name>" //name should be specified with path
  account_type    = "tenant"
  tenant_id       = "<tenant-id>"
  deployment_type = "azure"
  subscription_id = "<subscription-id>"
  features        = data.prismacloud_account_supported_features.prismacloud_supported_features_tenant.supported_features
}
```

### `Step 3`: Execute the generated terraform file <terraform-file>.tf.json in the above step in the Azure Portal to create app registration and roles. Copy the details from the script output

### `Step 4`: Onboard the Azure cloud Active Directory Tenant onto prisma cloud platform

```hcl
# Azure Tenant account type.

resource "prismacloud_org_cloud_account_v2" "example1" {
  azure {
    client_id    = "<client-id>"
    account_type = "tenant"
    enabled      = true
    name         = "test azure account"   // should be unique for each account
    group_ids    = [
      data.prismacloud_account_group.existing_account_group_id_org.group_id, // To use existing Account Group
      // prismacloud_account_group.new_account_group.group_id, // To create new Account group
    ]
    environment_type     = "azure" //default:azure
    key                  = "<secret-key>"
    service_principal_id = "<service-principal-id>"
    tenant_id            = "<tenant-id>"
  }
}

// Retrive existing account group name id
data "prismacloud_account_group" "existing_account_group_id_org" {
  name = "Default Account Group"
  // Change the account group name, if you already have an account group that you wish to map the account. 
}

// To create a new account group, if required
# resource "prismacloud_account_group" "new_account_group" {
#     name = "MyNewAccountGroup" // Account group name to be created
# }

```

### Consolidated code snippet for all the above steps

```
data "prismacloud_account_supported_features" "prismacloud_supported_features" {
  cloud_type      = "azure"
  account_type    = "tenant"
  deployment_type = "azure"
}

data "prismacloud_azure_template" "prismacloud_tenant_azure_template" {
  file_name       = "<file-name>" //name should be specified with path
  account_type    = "tenant"
  tenant_id       = "<tenant-id>"
  deployment_type = "azure"
  subscription_id = "<subscription-id>"
  features        = data.prismacloud_account_supported_features.prismacloud_supported_features_tenant.supported_features
}

resource "prismacloud_org_cloud_account_v2" "example1" {
  azure {
    client_id    = "<client-id>"
    account_type = "tenant"
    enabled      = true
    name         = "test azure account"   // should be unique for each account
    group_ids    = [
      data.prismacloud_account_group.existing_account_group_id_org.group_id, // To use existing Account Group
      // prismacloud_account_group.new_account_group.group_id, // To create new Account group
    ]
    environment_type     = "azure" //default:azure
    key                  = "<secret-key>"
    service_principal_id = "<service-principal-id>"
    tenant_id            = "<tenant-id>"
  }
}

// Retrive existing account group name id
data "prismacloud_account_group" "existing_account_group_id_org" {
  name = "Default Account Group"
  // Change the account group name, if you already have an account group that you wish to map the account. 
}

// To create a new account group, if required
# resource "prismacloud_account_group" "new_account_group" {
#     name = "MyNewAccountGroup" // Account group name to be created
# }

```

## **Example Usage 4**: For Bulk Azure cloud Active Directory Tenant accounts onboarding

### `Prerequisite Step`: Steps 1, 2, 3 mentioned in 'Example Usage 3' should be completed for each of the Tenant.

/*
You can also create cloud accounts from a CSV file using native Terraform
HCL and looping. Assume you have a CSV file of Azure accounts that looks like this:

name,clientId,key,tenantId,servicePrincipalId
123456789,6543256,0xJ8Q~,456189,86e43yuhbjc
213456789,5541253,0yJ9Q,356780,78e43yuhbbn
321466019,4543250,1xJ8Q~,256783,65e43iuhbjc

Here's how you would do this (Terraform 0.12 code):
*/

```
locals {
    instances = csvdecode(file("azure.csv"))
}
// Now specify the cloud account resource with a loop like so:

resource "prismacloud_org_cloud_account_v2" "azure_account_bulk_onboarding_example" {
    for_each = { for inst in local.instances : inst.name => inst }
    
    azure {
        name = each.value.name
        client_id=each.value.clientId
        key=each.value.key
        tenant_id=each.value.tenantId
        service_principal_id=each.value.servicePrincipalId 
    }
}
```

### prismacloud_org_cloud_account_v2 resource block example for Azure Cloud Tenant with hierarchy_selection

```
resource "prismacloud_org_cloud_account_v2" "example1" {
  disable_on_destroy = true
  azure {
    client_id = "<client-id>"
    account_type = "tenant"
    enabled     = true
    name        = "test azure account"
    environment_type       = "azure"
    key                   = "<secret-id>"
    monitor_flow_logs       = true
    service_principal_id    = "<service-principal-id>"
    tenant_id              = "<tenant-id>"
    default_account_group_id = "<deafult-account-group-id>" // must be provided for tenant with management groups(tenant)
    root_sync_enabled       = true     // must be true for tenant with management groups(tenant)
    hierarchy_selection {
            display_name = "displayNameHere"
            node_type= "nodeTypeHere"
            resource_id= "resurceIdHere"
            selection_type= "selectionTypeHere"
        }
    features {
      name  = "Agentless Scanning"
      state = "enabled"
    }
    features {
      name  = "Auto Protect"
      state = "disabled"
    }
  }
}
```

## Prerequisite

Before onboarding the azure cloud account. `azure_template` for account must be generated
using `prismacloud_azure_template`.
Refer *
*[Azure template generator Readme](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloud/latest/docs/data-sources/azure_template)
** for more details.

## Argument Reference

The type of cloud account to add.

* `disable_on_destroy` - (Optional, bool) To disable cloud account instead of deleting when calling Terraform destroy (
  default: `false`).
* `aws` - AWS account type spec, defined [below](#aws).
* `azure` - Azure account type spec, defined [below](#azure).

### AWS

* `account_id` - (Required) AWS account ID.
* `enabled` - (Optional, bool) Whether the account is enabled (default: `true`).
* `default_account_group_id` - (Optional, String) *Applicable only for accountType: **organization**.* This is the
  Default Account Group ID for the AWS organization and its member accounts.
* `group_ids` - (Optional) List of account IDs to which you are assigning this account.
* `name` - (Required) Name to be used for the account on the Prisma Cloud platform (must be unique).
* `role_arn` - (Required) Unique identifier for an AWS resource (ARN).
* `account_type` - (Optional) Defaults to `account` if not specified. Valid values : `account` and `organization`.
* `features` - (Optional, List) Features list

### Azure

* `enabled` - (Optional, bool) Whether the account is enabled (default: `true`).
* `group_ids` - (Optional) List of account IDs to which you are assigning this tenant account.
* `name` - (Required) Name to be used for the account on the Prisma Cloud platform (must be unique).
* `client_id` - (Required) Application ID registered with Active Directory.
* `key` - (Required) Application ID key.
* `monitor_flow_logs` - (Optional, bool) Automatically ingest flow logs.Should be `false` for `active directory tenant`.
* `tenant_id` - (Required) Active Directory ID associated with Azure.
* `service_principal_id` - (Required) Unique ID of the service principal object associated with the Prisma Cloud
  application that you create.
* `account_type` - (Optional) Defaults to "account" if not specified. Valid values: `account` or `tenant`.
* `hierarchy_selection` - (Optional) List of hierarchy selection. Each item has resource ID, display name, node type and
  selection type, as defined [below](#hierarchy-selection).
* `default_account_group_id` - (Optional, String) *Applicable only for accountType: **tenant**.* This is the Default
  Account Group ID for the Azure tenant and its member accounts (must be provided for tenant with management
  groups(`tenant`)).
* `features` - (Optional, List) Features list.
* `root_sync_enabled` - (Optional, bool) Azure tenant has children. Must be set to true when azure tenant is onboarded
  with children i.e., for "Tenant with management groups"(`tenant`).
* `environment_type` - (Optional) Defaults to "azure".Valid values are `azure` or `azure_gov` for azure tenant account.

## Attribute Reference

### AWS

* `account_id` - AWS account ID.
* `enabled` - (bool) Whether the account is enabled.
* `group_ids` - List of account IDs to which you are assigning this account.
* `default_account_group_id` - (Optional, String) *Applicable only for accountType: **organization**.* This is the
  Default Account Group ID for the AWS organization and its member accounts.
* `name` - Name to be used for the account on the Prisma Cloud platform.
* `role_arn` - Unique identifier for an AWS resource (ARN).
* `created_epoch_millis` - Account created epoch time.
* `customer_name` - Prisma customer name.
* `deleted` - Whether the account is deleted or not.
* `deployment_type` - `aws` for aws account.
* `eventbridge_rule_name_prefix` - Eventbridge rule name prefix.
* `external_id` - External id for aws account.
* `features` - Features applicable for aws account, defined [below](#features).
* `last_modified_by` - Last modified by.
* `last_modified_epoch_millis` - Last modified at epoch millis.
* `parent_id` - Parent id.
* `protection_mode` - Protection mode of account.

### Azure

* `account_id` - Azure tenant account ID.
* `enabled` - (bool) Whether the account is enabled.
* `group_ids` - List of account IDs to which you are assigning this account.
* `name` - Name to be used for the account on the Prisma Cloud platform (must be unique).
* `client_id` - Application ID registered with Active Directory.
* `key` - Application ID key.
* `monitor_flow_logs` - (bool) Automatically ingest flow logs.
* `tenant_id` - Active Directory ID associated with Azure.
* `service_principal_id` - Unique ID of the service principal object associated with the Prisma Cloud application that
  you create.
* `account_type` - `tenant` for azure account.
* `protection_mode` - Protection mode of account.
* `default_account_group_id` - Account group id to which you are assigning this account.
* `root_sync_enabled` - (bool) Azure tenant has children. Must be set to true when azure tenant is onboarded with
  children i.e., for `Tenant`.
* `hierarchy_selection` - List of hierarchy selection. Each item has resource ID, display name, node type and selection
  type, as defined [below](#hierarchy-selection).
* `features` - Features applicable for azure account, defined [below](#features).
* `environment_type` - `azure` or `azure_gov` for azure account.
* `parent_id` - Parent id.
* `customer_name` - Prisma customer name.
* `created_epoch_millis` - Account created epoch time.
* `last_modified_by` - Last modified by.
* `last_modified_epoch_millis` - Last modified at epoch millis.
* `deleted` - (bool) Whether the account is deleted or not.
* `template_url` - Template URL.
* `deployment_type` - `az` for azure account.
* `deployment_type_description` - Deployment type description. Valid values: Commercial or Government.
* `member_sync_enabled` - (bool) Azure tenant has children. Must be set to true when azure tenant is onboarded with
  children i.e., for `Tenant`.

#### FEATURES

* `name` - Feature name. Refer *
  *[Supported features readme](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloud/latest/docs/data-sources/cloud_account_supported_features)
  ** for more details.
* `state` - Feature state. Whether the feature to `enabled` or `disabled`.

#### Hierarchy Selection

* `resource_id` - Resource ID. For ACCOUNT, OU, ROOT, Tenant or Subscription. Example : `root`.
* `display_name` - Display name for ACCOUNT, OU, ROOT, Tenant or Subscription. Example : `Root`.
* `node_type` - Node type - ORG, OU, ACCOUNT, SUBSCRIPTION, TENANT or MANAGEMENT_GROUP.
* `selection_type` - Selection type - ALL, INCLUDE or EXCLUDE.

## Import

Resources can be imported using the cloud type and the ID:

```
$ terraform import prismacloud_cloud_account_v2.example accountIdHere
```