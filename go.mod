module github.com/terraform-providers/terraform-provider-prismacloud

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.0
	github.com/paloaltonetworks/prisma-cloud-go v0.4.0
)

//replace github.com/paloaltonetworks/prisma-cloud-go => ../prisma-cloud-go

go 1.13