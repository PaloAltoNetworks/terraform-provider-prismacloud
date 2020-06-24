package prismacloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDsPolicy(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsPolicyConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_policy.test", "policy_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_policy.test", "name"),
				),
			},
		},
	})
}

func testAccDsPolicyConfig() string {
	return `
data "prismacloud_policies" "x" {
    filters = {
        "policy.severity": "high",
        "policy.type": "network",
    }
}

locals {
    default_policies = [
        for inst in data.prismacloud_policies.x.listing :
        inst.policy_id if inst.system_default
    ]
}

data "prismacloud_policy" "test" {
    policy_id = local.default_policies[0]
}
`
}
