package prismacloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccUserProfiles(t *testing.T) {
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUserProfile(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_user_profile.test", "profile_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_user_profile.test", "name"),
				),
			},
		},
	})
}

func testAccUserProfile(name string) string {
	return fmt.Sprintf(`
data "prismacloud_user_profile" "test" {
    profile_id = prismacloud_user_profile.x.id
}

resource "prismacloud_user_profile" "x" {
    name = %q
}
`, name)
}
