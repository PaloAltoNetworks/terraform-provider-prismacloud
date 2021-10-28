package prismacloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDatapattern(t *testing.T) {
	name := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDatapatternConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_datapattern.test", "pattern_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_datapattern.test", "name"),
				),
			},
		},
	})
}

func testAccDatapatternConfig(name string) string {
	return fmt.Sprintf(`
data "prismacloud_datapattern" "test" {
    pattern_id = prismacloud_datapattern.x.pattern_id
}

resource "prismacloud_datapattern" "x" {
    name = %q
    description = "Data pattern Made by terraform"
}
`, name)
}
