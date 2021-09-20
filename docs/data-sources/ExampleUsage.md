---
page_title: "Prisma Cloud: Prismacloud Example Usage"
---

# prismacloud Example Usage 

`Create a terraform file` : touch filename.tf


## Example Usage

```hcl
provider "prismacloud" {
    url = "api.prismacloud.io"
    username= *******
    password= *********
    customer_name= *****
    skip_ssl_cert_verification= true
    logging= {"action": true, "send": true, "receive": true}
}

terraform {
  required_providers {
    prismacloud = {
      source = "my-host/my-namespace/prismacloud"
      version = "1.0.0"
    }
  }
}

data "prismacloud_account_group" "example" {
     name = "account_group_example"
}

output "account_group_info" {
  value = data.prismacloud_account_group.example
}
```

## terraform init

```hcl
Initializing the backend...

Initializing provider plugins...
- Finding my-host/my-namespace/prismacloud versions matching "1.0.0"...
- Installing my-host/my-namespace/prismacloud v1.0.0...
- Installed my-host/my-namespace/prismacloud v1.0.0 (unauthenticated)

Terraform has created a lock file .terraform.lock.hcl to record the provider
selections it made above. Include this file in your version control repository
so that Terraform can guarantee to make the same selections by default when
you run "terraform init" in the future.

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.
```


## terraform apply

```hcl

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:

Terraform will perform the following actions:

Plan: 0 to add, 0 to change, 0 to destroy.

Changes to Outputs:
  ~ account_group_info = {
      ~ account_ids      = null -> [
          + "53905******", 
          + "398411*****",
        ]
      ~ description      = null -> "Made by Terraform"
      ~ group_id         = null -> *********************
      ~ id               = null -> *********************
      ~ last_modified_by = null -> ******
      ~ last_modified_ts = null -> ******
      ~ name             = null -> "account_group_example"
    }

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes


Apply complete! Resources: 0 added, 0 changed, 0 destroyed.

Outputs:

account_group_info = {
  "account_ids" = toset([
     "53905******", 
     "398411*****"
  ])
  "description" = "Made by Terraform"
  "group_id" = **************
  "id" = ************
  "last_modified_by" = ******
  "last_modified_ts" = *******
  "name" = "account_group_example"
}
account_group_example
```

## terraform destroy

```hcl
An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:

Terraform will perform the following actions:

Plan: 0 to add, 0 to change, 0 to destroy.

Changes to Outputs:
  - account_group_info = {
      - account_ids      = [
          - "53905******", 
          - "398411*****",
        ]
      - description      = "Made by Terraform"
      - group_id         = ***************************
      - id               = ****************************
      - last_modified_by = *********
      - last_modified_ts = *********
      - name             = "account_group_example"
    } -> null

Do you really want to destroy all resources?
  Terraform will destroy all your managed infrastructure, as shown above.
  There is no undo. Only 'yes' will be accepted to confirm.

  Enter a value: yes


Destroy complete! Resources: 0 destroyed.
```
