package prismacloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSamlBypassUsers(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSamlBypassUsersConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_saml_bypass_users.test", "usernames"),
				),
			},
		},
	})
}

func testAccDataSourceSamlBypassUsersConfig() string {
	return `
provider "prismacloud" {
  # Provider configuration would go here
}

data "prismacloud_saml_bypass_users" "test" {}
`
}
