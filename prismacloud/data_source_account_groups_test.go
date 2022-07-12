package prismacloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDsAccountGroups(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsAccountGroupsConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_account_groups.test", "total"),
				),
			},
		},
	})
}

func testAccDsAccountGroupsConfig() string {
	return `
data "prismacloud_account_groups" "test" {}
`
}
