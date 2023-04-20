---
page_title: "Prisma Cloud: prismacloud_gcp_template"
---

# prismacloud_gcp_template

Retrieve information about gcp template for gcp account.

## Example Usage for Gcp Project

```hcl
data "prismacloud_gcp_template" "example"{
  name = "test account"
  account_type = "account"
  project_id = "<project_id>"
  authentication_type = "service_account"
  file_name = "<file-name>" //Provide filename along with path to store gcp template
}
```

## Example Usage for Gcp Master Service Account

```hcl
data "prismacloud_gcp_template" "example"{
  name = "test account"
  account_type = "masterServiceAccount"
  project_id = "<project_id>"
  authentication_type = "service_account"
  file_name = "<file-name>" //Provide filename along with path to store gcp template
}
```

## Example Usage for Gcp Organization

```hcl
data "prismacloud_gcp_template" "example"{
  name = "test account"
  account_type = "organization"
  org_id = "<org_id>"
  authentication_type = "service_account"
  file_name = "<file-name>" //Provide filename along with path to store gcp template
}
```

## Argument Reference

The following are the params that this data source supports:

* `name` - Name to be used for the account on the Prisma Cloud platform (must be unique).
* `account_type` - (Required) Gcp account type.
* `authentication_type` - (Optional) Authentication type of gcp account.
* `project_id` - (Optional) Gcp Project ID. Must be provided for account type `account` and `masterServiceAccount`.
* `org_id` - (Optional) Gcp organization ID. Must be provided for account type `organization`.
* `file_name` - (Required) File name to store gcp template (Provide filename along with path to store gcp template).
* `features` - (Optional) List of features. If features key/field is not passed, then the default features will be applicable. Refer : **[Supported features readme](https://registry.terraform.io/providers/PaloAltoNetworks/prismacloud/latest/docs/data-sources/cloud_account_supported_features) ** for more details.


