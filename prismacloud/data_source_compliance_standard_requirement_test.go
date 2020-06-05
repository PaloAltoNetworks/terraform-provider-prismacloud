package prismacloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDsComplianceStandardRequirement(t *testing.T) {
	cs := fmt.Sprintf("tf%s", acctest.RandString(6))
	csr := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsComplianceStandardRequirementConfig(cs, csr),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_compliance_standard_requirement.test", "csr_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_compliance_standard_requirement.test", "name"),
				),
			},
		},
	})
}

func testAccDsComplianceStandardRequirementConfig(cs, csr string) string {
	return fmt.Sprintf(`
data "prismacloud_compliance_standard_requirement" "test" {
    cs_id = prismacloud_compliance_standard_requirement.x.cs_id
    csr_id = prismacloud_compliance_standard_requirement.x.csr_id
}

resource "prismacloud_compliance_standard" "x" {
    name = %q
    description = "csr data source acctest"
}

resource "prismacloud_compliance_standard_requirement" "x" {
    cs_id = prismacloud_compliance_standard.x.cs_id
    name = %q
    description = "csr data source acctest"
}
`, cs, csr)
}
