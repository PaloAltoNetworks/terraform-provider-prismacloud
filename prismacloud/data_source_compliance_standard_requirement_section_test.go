package prismacloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDsComplianceStandardRequirementSection(t *testing.T) {
	cs := fmt.Sprintf("tf%s", acctest.RandString(6))
	csr := fmt.Sprintf("tf%s", acctest.RandString(6))
	csrs := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDsComplianceStandardRequirementSectionConfig(cs, csr, csrs),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prismacloud_compliance_standard_requirement_section.test", "csrs_id"),
					resource.TestCheckResourceAttrSet("data.prismacloud_compliance_standard_requirement_section.test", "section_id"),
				),
			},
		},
	})
}

func testAccDsComplianceStandardRequirementSectionConfig(cs, csr, csrs string) string {
	return fmt.Sprintf(`
data "prismacloud_compliance_standard_requirement_section" "test" {
    csr_id = prismacloud_compliance_standard_requirement_section.x.csr_id
    csrs_id = prismacloud_compliance_standard_requirement_section.x.csrs_id
}

resource "prismacloud_compliance_standard" "x" {
    name = %q
    description = "csrs data source acctest"
}

resource "prismacloud_compliance_standard_requirement" "x" {
    cs_id = prismacloud_compliance_standard.x.cs_id
    name = %q
    description = "csrs data source acctest"
}

resource "prismacloud_compliance_standard_requirement_section" "x" {
    csr_id = prismacloud_compliance_standard_requirement.x.csr_id
    section_id = %q
    description = "csrs data source acctest"
}
`, cs, csr, csrs)
}
