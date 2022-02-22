module github.com/terraform-providers/terraform-provider-prismacloud

require (
	github.com/hashicorp/terraform-plugin-sdk v1.17.2
	github.com/paloaltonetworks/prisma-cloud-go v0.5.1
)

//replace github.com/paloaltonetworks/prisma-cloud-go => ../prisma-cloud-go

go 1.13
