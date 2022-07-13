module github.com/terraform-providers/terraform-provider-prismacloud

require (
	github.com/hashicorp/terraform-plugin-sdk v1.17.2
	github.com/paloaltonetworks/prisma-cloud-go v0.5.3
)

//replace github.com/paloaltonetworks/prisma-cloud-go => ../prisma-cloud-go
replace github.com/paloaltonetworks/prisma-cloud-go =>  /users/naware/go/src/github.com/terraform-providers/terraform-provider-prismacloud/vendor/github.com/paloaltonetworks/prisma-cloud-go

go 1.13
