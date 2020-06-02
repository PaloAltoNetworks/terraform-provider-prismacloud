package prismacloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDsAccountGroup(t *testing.T) {
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsAccountGroupConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_account_group.test", "group_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_account_group.test", "last_modified_ts"),
				),
			},
		},
	})
}

func testAccDsAccountGroupConfig(name string) string {
	return fmt.Sprintf(`
resource "prismacloud_account_group" "x" {
    name = %q
    description = "data source account group acc test"
}

data "prismacloud_account_group" "test" {
    group_id = prismacloud_account_group.x.group_id
}
`, name)
}
