package prismacloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceSamlBypassSubsetUsers(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSamlBypassSubsetUsersConfig("user1", "user2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("prismacloud_saml_bypass_subset_users.test", "usernames.#", "2"),
					resource.TestCheckResourceAttr("prismacloud_saml_bypass_subset_users.test", "usernames.0", "user1"),
					resource.TestCheckResourceAttr("prismacloud_saml_bypass_subset_users.test", "usernames.1", "user2"),
				),
			},
			{
				Config: testAccResourceSamlBypassSubsetUsersConfig("user1", "user2", "user3"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("prismacloud_saml_bypass_subset_users.test", "usernames.#", "3"),
					resource.TestCheckResourceAttr("prismacloud_saml_bypass_subset_users.test", "usernames.0", "user1"),
					resource.TestCheckResourceAttr("prismacloud_saml_bypass_subset_users.test", "usernames.1", "user2"),
					resource.TestCheckResourceAttr("prismacloud_saml_bypass_subset_users.test", "usernames.2", "user3"),
				),
			},
			{
				Config: testAccResourceSamlBypassSubsetUsersConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("prismacloud_saml_bypass_subset_users.test", "usernames.#", "0"),
				),
			},
		},
	})
}

func testAccResourceSamlBypassSubsetUsersConfig(usernames ...string) string {
	config := `
provider "prismacloud" {
  # Provider configuration would go here
}

resource "prismacloud_saml_bypass_subset_users" "test" {
  usernames = [
`

	for i, username := range usernames {
		config += `	"` + username + `"`
		if i < len(usernames)-1 {
			config += `,
`
		}
	}

	config += `
  ]
}
`

	return config
}