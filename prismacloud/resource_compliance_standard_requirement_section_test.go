package prismacloud

import (
	"fmt"
	"testing"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/compliance/standard/requirement/section"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccComplianceStandardRequirementSection(t *testing.T) {
	var o section.Section
	cs := fmt.Sprintf("tf%s", acctest.RandString(6))
	csr := fmt.Sprintf("tf%s", acctest.RandString(6))
	csrs := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccComplianceStandardRequirementSectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComplianceStandardRequirementSectionConfig(cs, csr, csrs, "first desc"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComplianceStandardRequirementSectionExists("prismacloud_compliance_standard_requirement_section.test", &o),
					testAccCheckComplianceStandardRequirementSectionAttributes(&o, csrs, "first desc"),
				),
			},
			{
				Config: testAccComplianceStandardRequirementSectionConfig(cs, csr, csrs, "second desc"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComplianceStandardRequirementSectionExists("prismacloud_compliance_standard_requirement_section.test", &o),
					testAccCheckComplianceStandardRequirementSectionAttributes(&o, csrs, "second desc"),
				),
			},
		},
	})
}

func testAccCheckComplianceStandardRequirementSectionExists(n string, o *section.Section) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label ID is not set")
		}

		client := testAccProvider.Meta().(*pc.Client)
		csrId, csrsId := IdToTwoStrings(rs.Primary.ID)
		lo, err := section.GetId(client, csrId, csrsId)
		if err != nil {
			return fmt.Errorf("Error in get %s: %s", rs.Primary.ID, err)
		}

		*o = lo

		return nil
	}
}

func testAccCheckComplianceStandardRequirementSectionAttributes(o *section.Section, sid, desc string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.SectionId != sid {
			return fmt.Errorf("Section ID is %s, expected %s", o.SectionId, sid)
		}

		if o.Description != desc {
			return fmt.Errorf("Description is %s, expected %s", o.Description, desc)
		}

		return nil
	}
}

func testAccComplianceStandardRequirementSectionDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*pc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "prismacloud_compliance_standard_requirement_section" {
			continue
		}

		if rs.Primary.ID != "" {
			csrId, csrsId := IdToTwoStrings(rs.Primary.ID)
			if _, err := section.GetId(client, csrId, csrsId); err == nil {
				return fmt.Errorf("Object %q still exists", rs.Primary.ID)
			}
		}
		return nil
	}

	return nil
}

func testAccComplianceStandardRequirementSectionConfig(cs, csr, csrs, desc string) string {
	return fmt.Sprintf(`
resource "prismacloud_compliance_standard" "x" {
    name = %q
    description = "csrs acctest"
}

resource "prismacloud_compliance_standard_requirement" "x" {
    cs_id = prismacloud_compliance_standard.x.cs_id
    name = %q
    description = "csrs acctest"
}

resource "prismacloud_compliance_standard_requirement_section" "test" {
    csr_id = prismacloud_compliance_standard_requirement.x.csr_id
    section_id = %q
    description = %q
}
`, cs, csr, csrs, desc)
}
