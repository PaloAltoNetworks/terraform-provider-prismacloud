package prismacloud

import (
	"fmt"
	"testing"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/compliance/standard/requirement"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccComplianceStandardRequirement(t *testing.T) {
	var o requirement.Requirement
	cs := fmt.Sprintf("tf%s", acctest.RandString(6))
	csr := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccComplianceStandardRequirementDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComplianceStandardRequirementConfig(cs, csr, "first desc"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComplianceStandardRequirementExists("prismacloud_compliance_standard_requirement.test", &o),
					testAccCheckComplianceStandardRequirementAttributes(&o, csr, "first desc"),
				),
			},
			{
				Config: testAccComplianceStandardRequirementConfig(cs, csr, "second desc"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComplianceStandardRequirementExists("prismacloud_compliance_standard_requirement.test", &o),
					testAccCheckComplianceStandardRequirementAttributes(&o, csr, "second desc"),
				),
			},
		},
	})
}

func testAccCheckComplianceStandardRequirementExists(n string, o *requirement.Requirement) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label ID is not set")
		}

		client := testAccProvider.Meta().(*pc.Client)
		_, csrId := IdToTwoStrings(rs.Primary.ID)
		lo, err := requirement.Get(client, csrId)
		if err != nil {
			return fmt.Errorf("Error in get %s: %s", rs.Primary.ID, err)
		}

		*o = lo

		return nil
	}
}

func testAccCheckComplianceStandardRequirementAttributes(o *requirement.Requirement, name, desc string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if o.Name != name {
			return fmt.Errorf("Name is %s, expected %s", o.Name, name)
		}

		if o.Description != desc {
			return fmt.Errorf("Description is %s, expected %s", o.Description, desc)
		}

		return nil
	}
}

func testAccComplianceStandardRequirementDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*pc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "prismacloud_compliance_standard_requirement" {
			continue
		}

		if rs.Primary.ID != "" {
			_, csrId := IdToTwoStrings(rs.Primary.ID)
			if _, err := requirement.Get(client, csrId); err == nil {
				return fmt.Errorf("Object %q still exists", rs.Primary.ID)
			}
		}
		return nil
	}

	return nil
}

func testAccComplianceStandardRequirementConfig(cs, csr, desc string) string {
	return fmt.Sprintf(`
resource "prismacloud_compliance_standard" "x" {
    name = %q
    description = "csr acctest"
}

resource "prismacloud_compliance_standard_requirement" "test" {
    cs_id = prismacloud_compliance_standard.x.cs_id
    name = %q
    description = %q
}
`, cs, csr, desc)
}
