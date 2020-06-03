package prismacloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDsPolicies(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsPoliciesConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_policies.test", "total"),
				),
			},
		},
	})
}

func testAccDsPoliciesConfig() string {
	return `
data "prismacloud_policies" "test" {}
`
}
