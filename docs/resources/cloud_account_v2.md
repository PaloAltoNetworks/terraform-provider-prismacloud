---
page_title: "Prisma Cloud: prismacloud_cloud_account_v2"
---

# prismacloud_cloud_account_v2

Manage a cloud account on the Prisma Cloud platform.

## **Example Usage 1**: AWS cloud account onboarding
### `Step 1`: Fetch the supported features. Refer **[Supported features readme](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloud/latest/docs/data-sources/cloud_account_supported_features)** for more details.
```hcl
data "prismacloud_account_supported_features" "prismacloud_supported_features" {
    cloud_type = "aws"
    account_type = "account"
}
```
```hcl
output "features_supported" {
    value = data.prismacloud_account_supported_features.prismacloud_supported_features.supported_features
}
```

### `Step 2`: Fetch the AWS CFT s3 presigned url based on required features. Refer **[AWS CFT generator Readme](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloud/latest/docs/data-sources/aws_cft_generator_external_id)** for more details.
```hcl
data "prismacloud_aws_cft_generator" "prismacloud_account_cft" {
    account_type = "account"
    account_id = "<account-id>"
    features = data.prismacloud_account_supported_features.prismacloud_supported_features.supported_features
}
```
```hcl
output "s3_presigned_cft_url" {
    value = data.prismacloud_aws_cft_generator.prismacloud_account_cft.s3_presigned_cft_url
}
```

### `Step 3`: Create the IAM Role AWS CloudFormation Stack using S3 presigned cft url from above step2
To create the IAM role using terraform, the aws official terraform aws_cloudformation_stack resource can be used. Refer https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudformation_stack for more details

Example:
```
resource "aws_cloudformation_stack" "prismacloud_iam_role_stack" {
  name = "PrismaCloudApp" // change if needed
  capabilities = ["CAPABILITY_NAMED_IAM"]
#   parameters { // optional
#     PrismaCloudRoleName="<change-if-needed>" 
#   }
  template_url = data.prismacloud_aws_cft_generator.prismacloud_account_cft.s3_presigned_cft_url
}

output "iam_role" {
    value = aws_cloudformation_stack.prismacloud_iam_role_stack.outputs.PrismaCloudRoleARN
}
```

### `Step 4`: Onboard the cloud account onto prisma cloud platform

```hcl
# Single AWS account type.
resource "prismacloud_cloud_account_v2" "aws_account_onboarding_example" {
    disable_on_destroy = true
    aws {
        name = "myAwsAccountName" // should be unique for each account
        account_id = "<account-id>"
        group_ids = [
            data.prismacloud_account_group.existing_account_group_id.group_id,// To use existing Account Group
            // prismacloud_account_group.new_account_group.group_id, // To create new Account group
        ]
        role_arn = "${aws_cloudformation_stack.prismacloud_iam_role_stack.outputs.PrismaCloudRoleARN}" // IAM role arn from step 3
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
data "prismacloud_account_group" "existing_account_group_id" {
    name = "Default Account Group" // Change the account group name, if you already have an account group that you wish to map the account. 
}

// To create a new account group, if required
# resource "prismacloud_account_group" "new_account_group" {
#     name = "MyNewAccountGroup" // Account group name to be creatd
# }

```
### **Consolidated code snippet for all the above steps**

```
data "prismacloud_account_supported_features" "prismacloud_supported_features" {
    cloud_type = "aws"
    account_type = "account"
}

data "prismacloud_aws_cft_generator" "prismacloud_account_cft" {
    account_type = "account"
    account_id = "<account-id>"
    features = data.prismacloud_account_supported_features.prismacloud_supported_features.supported_features
}

resource "aws_cloudformation_stack" "prismacloud_iam_role_stack" {
  name = "PrismaCloudApp" // change if needed
  capabilities = ["CAPABILITY_NAMED_IAM"]
#   parameters { // optional
#     PrismaCloudRoleName="<change-if-needed>" 
#   }
  template_url = data.prismacloud_aws_cft_generator.prismacloud_account_cft.s3_presigned_cft_url
}

resource "prismacloud_cloud_account_v2" "aws_account_onboarding_example" {
    disable_on_destroy = true
    aws {
        name = "myAwsAccountName" // should be unique for each account
        account_id = "<account-id>"
        group_ids = [
            data.prismacloud_account_group.existing_account_group_id.group_id,// To use existing Account Group
            // prismacloud_account_group.new_account_group.group_id, // To create new Account group
        ]
        role_arn = "${aws_cloudformation_stack.prismacloud_iam_role_stack.outputs.PrismaCloudRoleARN}" // IAM role arn from prismacloud_iam_role_stack resource
        features {              // feature names from prismacloud_supported_features data source
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
data "prismacloud_account_group" "existing_account_group_id" {
    name = "Default Account Group" // If you already have an account group that you wish to map the account then change the account group name, 
}

// To create a new account group
# resource "prismacloud_account_group" "new_account_group" {
#     name = "MyNewAccountGroup" // Account group name to be creatd
# }

```

## **Example Usage 2**: Bulk AWS cloud accounts onboarding
### `Prerequisite Step`: Steps 1, 2, 3 mentioned in 'Example Usage 1' should be completed for each of the account and have IAM roles created.

/*
You can also create cloud accounts from a CSV file using native Terraform
HCL and looping.  Assume you have a CSV file of AWS accounts that looks like this (with
"||" separating account group IDs from each other):

accountId,groupIDs,name,roleArn
123456789,Default Account Group ID||AWS Account Group ID,123466789,arn:aws:iam::123456789:role/RedlockReadWriteRole
213456789,Default Account Group ID||AWS Account Group ID,213456789,arn:aws:iam::213456789:role/RedlockReadWriteRole
321466019,Default Account Group ID||AWS Account Group ID,321466019,arn:aws:iam::321466019:role/RedlockReadWriteRole

*/
```
locals {
    instances = csvdecode(file("aws.csv"))
}
// Now specify the cloud account resource with a loop like so:

resource "prismacloud_cloud_account_v2" "aws_account_bulk_onboarding_example" {
    for_each = { for inst in local.instances : inst.name => inst }
    
    aws {
        name = each.value.name
        account_id = each.value.accountId
        group_ids = split("||", each.value.groupIDs)
        role_arn = each.value.roleArn
    }
}
```

## Prerequisite

Before onboarding the aws cloud account. `external_id` for account must be generated using `prismacloud_aws_cft_generator`. Otherwise, you will encounter `error 412 : external_id_empty_or_not_generated`. Refer **[AWS CFT generator Readme](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloud/latest/docs/data-sources/aws_cft_generator_external_id)** for more details.

## **Example Usage 3**: Azure cloud account onboarding

### `Step 1`: Fetch the supported features. Refer **[Supported features readme](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloud/latest/docs/data-sources/cloud_account_supported_features)** for more details.

```hcl
data "prismacloud_account_supported_features" "prismacloud_supported_features" {
  cloud_type      = "azure"
  account_type    = "account"
  deployment_type = "azure"
}
```

```hcl
output "features_supported" {
  value = data.prismacloud_account_supported_features.prismacloud_supported_features.supported_features
}
```

### `Step 2`: Fetch the Azure template based on required features. Refer **[Azure template generator Readme](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloud/latest/docs/data-sources/azure_template)** for more details.

```hcl
data "prismacloud_azure_template" "prismacloud_azure_template" {
  file_name       = "<file-name>" //Provide filename along with path to store azure template
  account_type    = "account"
  tenant_id       = "<tenant-id>"
  deployment_type = "azure"
  subscription_id = "<subscription-id>"
  features        = data.prismacloud_account_supported_features.prismacloud_supported_features.supported_features
}
```

### `Step 3`: Execute the generated terraform file <terraform-file>.tf.json in the above step in the Azure Portal to create app registration and roles. Copy the details from the script output

### `Step 4`: Onboard the cloud account onto prisma cloud platform

```hcl
# Single Azure account type.
resource "prismacloud_cloud_account_v2" "azure_account_onboarding_example" {
  disable_on_destroy = true
  azure {
    client_id    = "<client-id>"
    account_id   = "<account-id>"
    account_type = "account"
    enabled      = false
    name         = "test azure account" //Should be unique for each account
    group_ids    = [
      data.prismacloud_account_group.existing_account_group_id.group_id, //To use existing Account Group
      //prismacloud_account_group.new_account_group.group_id, // To create new Account group
    ]
    key                  = "<secret-id>"
    monitor_flow_logs    = true
    service_principal_id = "<service-principal-id>"
    tenant_id            = "<tenant-id>"
    features {
      name  = "Agentless Scanning" //To enable 'Agentless Scanning' feature if required.
      state = "enabled"
    }
    features {
      name  = "Remediation"  //To enable Remediation also known as Monitor and Protect
      state = "enabled"
    }
  }
}

// Retrive existing account group name id
data "prismacloud_account_group" "existing_account_group_id" {
  name = "Default Account Group"
  // Change the account group name, if you already have an account group that you wish to map the account. 
}

// To create a new account group, if required
# resource "prismacloud_account_group" "new_account_group" {
#     name = "MyNewAccountGroup" // Account group name to be created
# }

```

### **Consolidated code snippet for all the above steps**

```
data "prismacloud_account_supported_features" "prismacloud_supported_features" {
  cloud_type      = "azure"
  account_type    = "account"
  deployment_type = "azure"
}

data "prismacloud_azure_template" "prismacloud_azure_template" {
  file_name       = "<file-name>" //Provide filename along with path to store azure template
  account_type    = "account"
  tenant_id       = "<tenant-id>"
  deployment_type = "azure"
  subscription_id = "<subscription-id>"
  features        = data.prismacloud_account_supported_features.prismacloud_supported_features.supported_features
}

resource "prismacloud_cloud_account_v2" "azure_account_onboarding_example" {
  disable_on_destroy = true
  azure {
    client_id    = "<client-id>"
    account_id   = "<account-id>"
    account_type = "account"
    enabled      = false
    name         = "test azure account" //Should be unique for each account
    group_ids    = [
      data.prismacloud_account_group.existing_account_group_id.group_id, //To use existing Account Group
      //prismacloud_account_group.new_account_group.group_id, //To create new Account group
    ]
    key                  = "<secret-id>"
    monitor_flow_logs    = true
    service_principal_id = "<service-principal-id>"
    tenant_id            = "<tenant-id>"
    features {
      name  = "Agentless Scanning" //To enable 'Agentless Scanning' feature if required.
      state = "enabled"
    }
    features {
      name  = "Remediation"  //To enable Remediation also known as Monitor and Protect
      state = "enabled"
    }
  }
}

// Retrive existing account group name id
data "prismacloud_account_group" "existing_account_group_id" {
  name = "Default Account Group"
  //Change the account group name, if you already have an account group that you wish to map the account. 
}

// To create a new account group, if required
# resource "prismacloud_account_group" "new_account_group" {
#     name = "MyNewAccountGroup" //Account group name to be created
# }

```

## **Example Usage 4**: Bulk Azure cloud accounts onboarding

### `Prerequisite Step`: Steps 1, 2, 3 mentioned in 'Example Usage 3' should be completed for each of the account.

/*
You can also create cloud accounts from a CSV file using native Terraform
HCL and looping. Assume you have a CSV file of Azure accounts that looks like this (with
"||" separating account group IDs from each other):

accountId,groupIDs,name,clientId,key,tenantId,servicePrincipalId
123456789,Default Account Group ID||Azure Account Group ID,123456789,6543256,0xJ8Q~,456189,86e43yuhbjc
213456789,Default Account Group ID||Azure Account Group ID,213456789,5541253,0yJ9Q,356780,78e43yuhbbn
321466019,Default Account Group ID||Azure Account Group ID,321466019,4543250,1xJ8Q~,256783,65e43iuhbjc

*/

```
locals {
    instances = csvdecode(file("azure.csv"))
}
// Now specify the cloud account resource with a loop like so:

resource "prismacloud_cloud_account_v2" "azure_account_bulk_onboarding_example" {
    for_each = { for inst in local.instances : inst.name => inst }
    
    azure {
        account_id = each.value.accountId
        group_ids = split("||", each.value.groupIDs)
        name = each.value.name
        client_id=each.value.clientId
        key=each.value.key
        tenant_id=each.value.tenantId
        service_principal_id=each.value.servicePrincipalId 
    }
}
```

## Prerequisite

Before onboarding the azure cloud account. `azure_template` for account must be generated using `prismacloud_azure_template`. Refer **[Azure template generator Readme](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloud/latest/docs/data-sources/azure_template)** for more details.

## Argument Reference

The type of cloud account to add.

* `disable_on_destroy` - (Optional, bool) To disable cloud account instead of deleting when calling Terraform destroy (default: `false`).
* `aws` - AWS account type spec, defined [below](#aws).
* `azure` - Azure account type spec, defined [below](#azure).

### AWS

* `account_id` - (Required) AWS account ID.
* `enabled` - (Optional, bool) Whether the account is enabled (default: `true`).
* `group_ids` - (Optional) List of account IDs to which you are assigning this account.
* `name` - (Required) Name to be used for the account on the Prisma Cloud platform (must be unique).
* `role_arn` - (Required) Unique identifier for an AWS resource (ARN).
* `account_type` - (Optional) Defaults to `account` if not specified. Valid values : `account` and `organization`.
* `features` - (Optional, List) Features list.

### Azure

* `account_id` - (Required) Azure account ID.
* `enabled` - (Optional, bool) Whether the account is enabled (default: `true`).
* `group_ids` - (Required) List of account IDs to which you are assigning this account.
* `name` - (Required) Name to be used for the account on the Prisma Cloud platform (must be unique).
* `client_id` - (Required) Application ID registered with Active Directory.
* `key` - (Required) Application ID key.
* `monitor_flow_logs` - (Optional, bool) Automatically ingest flow logs.
* `tenant_id` - (Required) Active Directory ID associated with Azure.
* `service_principal_id` - (Required) Unique ID of the service principal object associated with the Prisma Cloud application that you create.
* `account_type` - (Optional) Defaults to `account` if not specified. Valid values: `account` or `tenant`.
* `features` - (Optional, List) Features applicable for azure account, defined [below](#features).
* `environment_type` - (Optional) Defaults to `azure`.Valid values are `azure`,`azure_gov` or `azure_china` for azure subscription account.

## Attribute Reference

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
* `environment_type` - `azure`,`azure_gov` or `azure_china` for azure subscription account.
* `parent_id` - Parent id.
* `customer_name` - Prisma customer name.
* `created_epoch_millis` - Account created epoch time.
* `last_modified_by` - Last modified by.
* `last_modified_epoch_millis` - Last modified at epoch millis.
* `deleted` - (bool) Whether the account is deleted or not.
* `template_url` - Template URL.
* `deployment_type` - `az` for azure account.
* `deployment_type_description` - Deployment type description. Valid values: `Commercial` or `Government`.

#### FEATURES

* `name` - Feature name. Refer **[Supported features readme](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloud/latest/docs/data-sources/cloud_account_supported_features)** for more details.
* `state` - Feature state. Whether the feature to `enabled` or `disabled`.


## Import

Resources can be imported using the cloud type and the ID:

```
$ terraform import prismacloud_cloud_account_v2.example accountIdHere
```
