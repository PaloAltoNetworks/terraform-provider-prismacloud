package prismacloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestOrgAccDsCloudAccounts(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testOrgAccDsCloudAccountsConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_org_cloud_accounts.test", "total"),
				),
			},
		},
	})
}

func testOrgAccDsCloudAccountsConfig() string {
	return `
data "prismacloud_org_cloud_accounts" "test" {}
`
}
