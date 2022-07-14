package prismacloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDsUserRole(t *testing.T) {
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsUserRole(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_user_role.test", "terraform_user"),
					resource.TestCheckResourceAttrSet("data.prismacloud_user_role.test", "name"),
				),
			},
		},
	})
}

func testAccDsUserRole(name string) string {
	return fmt.Sprintf(`
resource "prismacloud_user_role" "x" {
    name = %q
    description = "User Role for Terraform"
	role_type = "System Admin"
}

data "prismacloud_user_role" "test" {
    role_id = prismacloud_user_role.x.role_id
}
`, name)
}
