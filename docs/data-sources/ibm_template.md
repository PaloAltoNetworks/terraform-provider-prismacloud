---
page_title: "Prisma Cloud: prismacloud_ibm_template"
---

# prismacloud_ibm_template

Retrieve information about ibm template for IBM account.

## Example Usage for IBM Account

```hcl
data "prismacloud_ibm_template" "example" {
  file_name    = "<file-name>" //Provide filename along with path to store ibm template
  account_type = "account"
}
```

## Argument Reference

The following are the params that this data source supports:

* `account_type` - (Required) IBM account type.
* `file_name` - (Required) File name to store ibm template (Provide filename along with path to store ibm template).