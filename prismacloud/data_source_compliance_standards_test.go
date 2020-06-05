package prismacloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDsComplianceStandards(t *testing.T) {
	cs := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsComplianceStandardsConfig(cs),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_compliance_standards.test", "total"),
				),
			},
		},
	})
}

func testAccDsComplianceStandardsConfig(cs string) string {
	return fmt.Sprintf(`
data "prismacloud_compliance_standards" "test" {}

resource "prismacloud_compliance_standard" "x" {
    name = %q
    description = "csr listing data source acctest"
}
`, cs)
}
