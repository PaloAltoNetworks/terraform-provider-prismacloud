package prismacloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDsComplianceStandard(t *testing.T) {
	cs := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsComplianceStandardConfig(cs),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_compliance_standard.test", "cs_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_compliance_standard.test", "name"),
				),
			},
		},
	})
}

func testAccDsComplianceStandardConfig(cs string) string {
	return fmt.Sprintf(`
data "prismacloud_compliance_standard" "test" {
    cs_id = prismacloud_compliance_standard.x.cs_id
}

resource "prismacloud_compliance_standard" "x" {
    name = %q
    description = "cs data source acctest"
}
`, cs)
}
