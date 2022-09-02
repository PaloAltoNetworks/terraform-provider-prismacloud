package prismacloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDsComplianceStandardRequirements(t *testing.T) {
	cs := fmt.Sprintf("tf%s", acctest.RandString(6))
	csr := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsComplianceStandardRequirementsConfig(cs, csr),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_compliance_standard_requirements.test", "total"),
				),
			},
		},
	})
}

func testAccDsComplianceStandardRequirementsConfig(cs, csr string) string {
	return fmt.Sprintf(`
data "prismacloud_compliance_standard_requirements" "test" {
    cs_id = prismacloud_compliance_standard.x.cs_id
}

resource "prismacloud_compliance_standard" "x" {
    name = %q
    description = "csr listing data source acctest"
}

resource "prismacloud_compliance_standard_requirement" "x" {
    cs_id = prismacloud_compliance_standard.x.cs_id
    name = %q
    description = "csr listing data source acctest"
}
`, cs, csr)
}
