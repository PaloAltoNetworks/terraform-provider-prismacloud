module github.com/terraform-providers/terraform-provider-prismacloud

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.0
	github.com/paloaltonetworks/prisma-cloud-go v0.0.0-20200327180535-f648dce6b742
)

go 1.13

replace github.com/paloaltonetworks/prisma-cloud-go => ../prisma-cloud-go
