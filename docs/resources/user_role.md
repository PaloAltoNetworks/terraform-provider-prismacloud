---
page_title: "Prisma Cloud: prismacloud_user_role"
---

# prismacloud_user_role

Manage an user role.

## Example Usage

```hcl
resource "prismacloud_user_role" "example" {
    name = "my user role"
    description = "Made by Terraform"
    role_type = "Account Group Admin"
}

/*
You can also create user roles from a CSV file using native Terraform
HCL and looping.  Assume you have a CSV file (user_roles.csv) of user roles that looks like this (with
"||" separating account group IDs from each other):

"name","description","roletype","account_group_ids","restrict_dismissal_access","only_allow_ci_access","only_allow_read_access","has_defender_permissions","only_allow_compute_access"
"test_role1","Made by Terraform","System Admin",,false,false,false,true,false
"test_role2","Made by Terraform","Account Group Admin","11111111-2222-3333-4444-555555555555||12345678-2222-3333-4444-555555555555",false,false,false,false,false
"test_role3","Made by Terraform","Account Group Read Only","12345678-2222-3333-4444-555555555555",true,false,true,false,false
"test_role4","Made by Terraform","Cloud Provisioning Admin","12345678-2222-3333-4444-555555555555",true,false,false,true,false
"test_role5","Made by Terraform","Build and Deploy Security",,true,false,false,false,false
"test_role6","Made by Terraform","Account and Cloud Provisioning Admin","12345678-2222-3333-4444-555555555555",false,false,false,false,false

Here's how you would do this:
*/
locals {
    user_roles = csvdecode(file("user_roles.csv"))
}

// Now specify the user role resource with a loop like this:
resource "prismacloud_user_role" "example" {
    for_each = { for inst in local.user_roles : inst.name => inst }

    name = each.value.name
    description = each.value.description
    role_type = each.value.roletype
    restrict_dismissal_access = each.value.restrict_dismissal_access
    account_group_ids = (each.value.roletype == "System Admin" || each.value.roletype == "Build and Deploy Security") ? [] : split("||", each.value.account_group_ids)
    additional_attributes {
        only_allow_ci_access = each.value.only_allow_ci_access
        only_allow_read_access = each.value.only_allow_read_access
        has_defender_permissions = each.value.has_defender_permissions
        only_allow_compute_access = each.value.only_allow_compute_access
    }  
}
```

## Argument Reference

* `name` - (Required) Name of the role.
* `description` - (Optional) Description.
* `role_type` - (Required) User role type.  Valid values are `System Admin`, `Account Group Admin`, `Account Group Read Only`, `Cloud Provisioning Admin`, `Account and Cloud Provisioning Admin`, or `Build and Deploy Security`.
* `account_group_ids` - (Optional) List of accessible account group IDs. (Can't be set if `role_type` is `System Admin` or `Build and Deploy Security`)
* `resource_list_ids` - (Optional) List of resource list IDs.
* `code_repository_ids` - (Optional) List of code repository IDs.
* `restrict_dismissal_access` - (Optional, bool) Restrict dismissal access.
* `additional_attributes` - (Optional) An Additional attributes spec, as defined [below](#additional-attributes).

## Additional Attributes

* `only_allow_ci_access` - (Optional, bool) - Allows only CI Access
* `only_allow_read_access` - (Optional, bool) - Allow read only access (True for Account Group Read Only user role)
* `has_defender_permissions`- (Optional, bool) - Has defender Permissions (True for Cloud Provisioning Admin user Role)
* `only_allow_compute_access`- (Optional, bool) - Access to only Compute tab and Access keys

## Attribute Reference

* `role_id` - Role UUID.
* `last_modified_by` - Last modified by
* `last_modified_ts` - (int) Last modified timestamp.
* `associated_users` - List of associated application users which cannot exist in the system without the user role.
* `account_groups` - List of account groups, as defined [below](#account-groups).

### Account Groups

Each account group has the following attributes.

* `group_id` - The group ID.
* `name` - Group name.

## Import

Resources can be imported using the role ID:

```
$ terraform import prismacloud_user_role.example 11111-22-33
```
