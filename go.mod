module github.com/terraform-providers/terraform-provider-prismacloud

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.4.3
	github.com/paloaltonetworks/prisma-cloud-go v0.5.3
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
)

//replace github.com/paloaltonetworks/prisma-cloud-go => ../prisma-cloud-go

go 1.13
