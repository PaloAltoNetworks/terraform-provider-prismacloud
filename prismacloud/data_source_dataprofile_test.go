package prismacloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataprofile(t *testing.T) {
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataprofileConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_dataprofile.test", "profile_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_dataprofile.test", "name"),
				),
			},
		},
	})
}

func testAccDataprofileConfig(name string) string {
	return fmt.Sprintf(`
data "prismacloud_dataprofile" "test" {
    profile_id = prismacloud_dataprofile.x.profile_id
}

resource "prismacloud_dataprofile" "x" {
    name = %q
    description = "Data profile Made by terraform"
}
`, name)
}
