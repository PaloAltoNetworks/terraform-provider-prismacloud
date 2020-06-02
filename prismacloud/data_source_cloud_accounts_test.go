package prismacloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDsCloudAccounts(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsCloudAccountsConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_cloud_accounts.test", "total"),
				),
			},
		},
	})
}

func testAccDsCloudAccountsConfig() string {
	return `
data "prismacloud_cloud_accounts" "test" {}
`
}
