package prismacloud

import (
	"fmt"
	"testing"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/compliance/standard"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccComplianceStandard(t *testing.T) {
	var o standard.Standard
	cs := fmt.Sprintf("tf%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccComplianceStandardDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComplianceStandardConfig(cs, "cs acctest first desc"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComplianceStandardExists("prismacloud_compliance_standard.test", &o),
					testAccCheckComplianceStandardAttributes(&o, cs, "cs acctest first desc"),
				),
			},
			{
				Config: testAccComplianceStandardConfig(cs, "cs acctest second desc"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComplianceStandardExists("prismacloud_compliance_standard.test", &o),
					testAccCheckComplianceStandardAttributes(&o, cs, "cs acctest second desc"),
				),
			},
		},
	})
}

func testAccCheckComplianceStandardExists(n string, o *standard.Standard) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Resource not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Object label ID is not set")
		}

		client := testAccProvider.Meta().(*pc.Client)
		csId := rs.Primary.ID
		lo, err := standard.Get(client, csId)
		if err != nil {
			return fmt.Errorf("Error in get %s: %s", rs.Primary.ID, err)
		}

		*o = lo

		return nil
	}
}

func testAccCheckComplianceStandardAttributes(o *standard.Standard, name, desc string) resource.TestCheckFunc {
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

func testAccComplianceStandardDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*pc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "prismacloud_compliance_standard" {
			continue
		}

		if rs.Primary.ID != "" {
			csId := rs.Primary.ID
			if _, err := standard.Get(client, csId); err == nil {
				return fmt.Errorf("Object %q still exists", rs.Primary.ID)
			}
		}
		return nil
	}

	return nil
}

func testAccComplianceStandardConfig(cs, desc string) string {
	return fmt.Sprintf(`
resource "prismacloud_compliance_standard" "test" {
    name = %q
    description = %q
}
`, cs, desc)
}
