module github.com/terraform-providers/terraform-provider-prismacloud

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.0
	github.com/paloaltonetworks/prisma-cloud-go v0.4.1-0.20210716024753-6249da101d38
)

//replace github.com/paloaltonetworks/prisma-cloud-go => ../prisma-cloud-go

go 1.13
